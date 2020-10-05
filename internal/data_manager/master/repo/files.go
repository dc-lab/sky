package repo

import (
	"database/sql"
	"strings"

	_ "github.com/lib/pq"
	"github.com/markbates/pkger"
	log "github.com/sirupsen/logrus"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/cenkalti/backoff/v4"

	"github.com/dc-lab/sky/internal/data_manager/master/modeldb"
)

type FilesRepo struct {
	Conn *sql.DB
}

func OpenFilesRepo(driver string, connStr string) (*FilesRepo, error) {
	db, err := sql.Open(driver, connStr)

	if err == nil {
		err = backoff.Retry(func() error {
			err := db.Ping()
			if err != nil {
				log.WithError(err).Error("Connection to database failed")
			}
			return err
		}, backoff.NewExponentialBackOff())
	}

	repo := &FilesRepo{
		Conn: db,
	}

	if err != nil {
		return repo, err
	}

	err = repo.migrate(connStr)

	if err != nil {
		log.WithError(err).Fatalln("Failed to run migrations")
	}

	return repo, err
}

func (s *FilesRepo) migrate(connStr string) error {
	driver, err := postgres.WithInstance(s.Conn, &postgres.Config{})
	if err != nil {
		return err
	}

	pkger.Include("internal/data_manager/master/migrations")
	m, err := migrate.NewWithDatabaseInstance("pkger://internal/data_manager/master/migrations", "postgres", driver)
	if err != nil {
		return err
	}
	err = m.Up()
	if err == migrate.ErrNoChange {
		err = nil
	}
	return err
}

func (s *FilesRepo) Create(file *modeldb.File) (*modeldb.File, error) {
	err := s.Conn.QueryRow(
		`INSERT INTO files(
			owner,
			name,
			tags,
			task_id,
			executable
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING 
			id,
			upload_token`,
		file.Owner, file.Name, file.Tags, file.TaskId, file.Executable,
	).Scan(&file.Id, &file.UploadToken)
	return file, err
}

func (s *FilesRepo) Update(file modeldb.File) (modeldb.File, error) {
	_, err := s.Conn.Exec(
		`UPDATE files
		SET
			owner=$2,
			name=$3,
			hash=$4,
			tags=$5,
			task_id=$6,
			executable=$7,
			content_type=$8
		WHERE id=$1`,
		file.Id, file.Owner, file.Name, file.Hash, file.Tags, file.TaskId, file.Executable, file.ContentType,
	)
	return file, err
}

func (s *FilesRepo) Get(id string) (*modeldb.File, error) {
	var file modeldb.File

	err := s.Conn.QueryRow(
		`SELECT
			id,
			owner,
			name,
			hash,
			tags,
			task_id,
			executable,
			upload_token,
			content_type
		FROM files
		WHERE id=$1`, id,
	).Scan(&file.Id, &file.Owner, &file.Name, &file.Hash, &file.Tags, &file.TaskId, &file.Executable, &file.UploadToken, &file.ContentType)
	return &file, err
}

func (s *FilesRepo) GetTaskResults(task_id string, path_prefix string) ([]modeldb.File, error) {
	path_prefix = strings.TrimSuffix(path_prefix, "/")
	rows, err := s.Conn.Query(
		`SELECT
			id,
			owner,
			name,
			hash,
			tags,
			task_id,
			executable,
			upload_token,
			content_type
		FROM files
		WHERE task_id=$1 AND (name LIKE $2 || '/%' OR name LIKE $2)`, task_id, path_prefix,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	files := make([]modeldb.File, 0, 16)

	for rows.Next() {
		var file modeldb.File

		err = rows.Scan(&file.Id, &file.Owner, &file.Name, &file.Hash, &file.Tags, &file.TaskId, &file.Executable, &file.UploadToken, &file.ContentType)
		if err != nil {
			return nil, err
		}

		files = append(files, file)
	}
	return files, err
}

func (s *FilesRepo) Delete(file *modeldb.File) error {
	_, err := s.Conn.Exec(
		`DELETE FROM files
		WHERE id=$1`,
		file.Id,
	)
	return err
}

func (s *FilesRepo) IncFileHashRefCount(hash string) (uint, error) {
	var count uint = 0

	_, err := s.Conn.Exec(
		`INSERT INTO hash_counts VALUES ($1, 0) ON CONFLICT DO NOTHING`, hash,
	)
	if err != nil {
		return count, err
	}

	err = s.Conn.QueryRow(
		`UPDATE hash_counts
		SET ref_count=ref_count+1
		WHERE hash=$1 RETURNING ref_count`,
		hash,
	).Scan(&count)

	return count, err
}

func (s *FilesRepo) AddHashTarget(hash string, node string) error {
	_, err := s.Conn.Exec(
		`INSERT INTO hash_targets VALUES ($1, $2) ON CONFLICT DO NOTHING`, hash, node,
	)

	return err
}

func (s *FilesRepo) SetFileHash(file_id string, hash string) (bool, error) {
	resId := ""
	err := s.Conn.QueryRow(
		`UPDATE files
		SET hash=$2
		WHERE id=$1 AND hash=''
		RETURNING id
		`,
		file_id, hash,
	).Scan(&resId)
	return resId == file_id, err
}

func (s *FilesRepo) UpdateNodeReportTimestamp(location string, freeSpace int64) error {
	_, err := s.Conn.Exec(`
		INSERT INTO nodes
		VALUES ($1, $2, CURRENT_TIMESTAMP)
		ON CONFLICT(location) DO UPDATE
		SET free_space=$2,
			report_time=CURRENT_TIMESTAMP
		`, location, freeSpace,
	)

	return err
}

func (s *FilesRepo) UpdateLocations(location string, blobs []string) error {
	tx, err := s.Conn.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec(`DELETE FROM hash_locations WHERE location = $1`, location)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO hash_locations VALUES ($1, $2) ON CONFLICT DO NOTHING`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, hash := range blobs {
		_, err = stmt.Exec(hash, location)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (s *FilesRepo) GetBlobsForLocation(location string) ([]string, error) {
	rows, err := s.Conn.Query(
		`SELECT hash FROM hash_targets WHERE location=$1`, location,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	blobs := make([]string, 0, 16)

	for rows.Next() {
		var hash string

		err = rows.Scan(&hash)
		if err != nil {
			return nil, err
		}

		blobs = append(blobs, hash)
	}

	return blobs, err
}

func (s *FilesRepo) SelectBestUploadStorages() ([]string, error) {
	// TODO: Calculate some availability metric
	// TODO: Select storages by acl
	rows, err := s.Conn.Query(
		`SELECT location FROM nodes ORDER BY free_space DESC LIMIT 5`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	locations := make([]string, 0, 5)

	for rows.Next() {
		var location string

		err = rows.Scan(&location)
		if err != nil {
			return nil, err
		}

		locations = append(locations, location)
	}

	return locations, err
}

func (s *FilesRepo) GetFileLocations(id string) ([]string, error) {
	file, err := s.Get(id)
	if err != nil {
		return nil, err
	}

	rows, err := s.Conn.Query(
		`SELECT location FROM hash_locations WHERE hash=$1 ORDER BY report_time DESC LIMIT 5`, file.Hash,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	locations := make([]string, 0, 5)

	for rows.Next() {
		var location string

		err = rows.Scan(&location)
		if err != nil {
			return nil, err
		}

		locations = append(locations, location)
	}

	return locations, err
}
