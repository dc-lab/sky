package repo

import (
	log "github.com/sirupsen/logrus"

	"database/sql"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/cenkalti/backoff/v4"

	"github.com/dc-lab/sky/data_manager/modeldb"
)

type FilesRepo struct {
	Conn *sql.DB
}

func OpenFilesRepo(driver string, connStr string) (*FilesRepo, error) {
	db, err := sql.Open(driver, connStr)

	if err == nil {
		err = backoff.Retry(db.Ping, backoff.NewExponentialBackOff())
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
	m, err := migrate.NewWithDatabaseInstance("file:///migrations", "postgres", driver)
	if err != nil {
		return err
	}
	err = m.Up()
	if err == migrate.ErrNoChange {
		err = nil
	}
	return err
}

func (s *FilesRepo) Create(file modeldb.File) (modeldb.File, error) {
	err := s.Conn.QueryRow(
		`INSERT INTO files(
			owner,
			name,
			hash,
			tags,
			task_id,
			executable
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING 
			id,
			upload_token`,
		file.Owner, file.Name, file.Hash, file.Tags, file.TaskId, file.Executable,
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

func (s *FilesRepo) Get(id string) (modeldb.File, error) {
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
	return file, err
}

func (s *FilesRepo) Delete(file modeldb.File) error {
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
		`INSERT INTO hash_ref_counts VALUES ($1, 0) ON CONFLICT DO NOTHING`, hash,
	)
	if err != nil {
		return count, err
	}

	err = s.Conn.QueryRow(
		`UPDATE hash_ref_counts
		SET ref_count=ref_count+1
		WHERE hash=$1 RETURNING ref_count`,
		hash,
	).Scan(&count)
	return count, err
}
