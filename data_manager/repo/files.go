package repo

import (
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/cenkalti/backoff/v4"

	"data_manager/modeldb"
)

type FilesRepo struct {
	Conn *sql.DB
}

func OpenFilesRepo(driver string, connStr string) (*FilesRepo, error) {
	db, err := sql.Open(driver, connStr)

	if err == nil {
		err = backoff.Retry(db.Ping, backoff.NewExponentialBackOff())
	}

	return &FilesRepo{
		Conn: db,
	}, err
}

func (s *FilesRepo) Migrate() error {
	_, err := s.Conn.Exec(`
		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
		CREATE TABLE IF NOT EXISTS files(
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			owner text NOT NULL DEFAULT '',
			name text NOT NULL DEFAULT '',
			tags json,
			task_id text NOT NULL DEFAULT '',
			executable boolean NOT NULL DEFAULT false,
			hash text NOT NULL DEFAULT '',
			content_type text NOT NULL DEFAULT '',
			upload_token uuid DEFAULT uuid_generate_v4()
		);
		CREATE TABLE IF NOT EXISTS hash_ref_counts(
			hash text PRIMARY KEY,
			ref_count integer NOT NULL
		);
	`)

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
			content_type=$8,
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
			content_type,
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
