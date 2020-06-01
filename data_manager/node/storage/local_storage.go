package storage

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"hash"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/dc-lab/sky/data_manager/node/config"
)

type LocalStorage struct {
	config *config.Config
}

func MakeLocalStorage(config *config.Config) (Storage, error) {
	err := os.Mkdir(config.StorageDir, 0o755)
	if err != nil && !os.IsExist(err) {
		log.WithError(err).WithField("dir", config.StorageDir).Fatalln("Failed to create storage directory")
		return nil, err
	}

	return &LocalStorage{config}, nil
}

const (
	hashEncodedSize   = 64
	hashPrefixDirSize = 2
)

func buildHashFromPath(root string, file string) (string, error) {
	file, err := filepath.Rel(root, file)
	if err != nil {
		return "", err
	}

	hash := ""
	for file != "" {
		var comp string
		file, comp = path.Split(file)
		hash = comp + hash
	}

	if len(hash) != hashEncodedSize {
		return "", errors.New("Invalid hash value: " + hash)
	}
	return hash, nil
}

func buildPathFromHash(root string, hash string) (string, error) {
	// For now we just split after first two bytes
	// abcdefgh -> /ab/cdefgh

	if len(hash) != hashEncodedSize {
		return "", errors.New("Invalid hash value: " + hash)
	}

	prefix := hash[:hashPrefixDirSize]
	suffix := hash[hashPrefixDirSize:]
	dir := path.Join(root, prefix)

	err := os.MkdirAll(dir, 0o777)
	if err != nil {
		return "", err
	}

	return path.Join(dir, suffix), nil
}

func (s *LocalStorage) ListBlobs(f func(hash string) error) error {
	err := filepath.Walk(s.config.StorageDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		hash, err := buildHashFromPath(s.config.StorageDir, path)
		if err != nil {
			return err
		}

		if err := f(hash); err != nil {
			return err
		}

		return nil
	})
	return err
}

func (s *LocalStorage) PutBlob() (BlobWriter, error) {
	return makeLocalBlobWriter(s.config)
}

func (s *LocalStorage) GetBlob(hash string) (int64, io.Reader, error) {
	filePath, err := buildPathFromHash(s.config.StorageDir, hash)
	if err != nil {
		return 0, nil, err
	}
	stat, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		return 0, nil, nil
	} else if err != nil {
		return 0, nil, err
	}

	reader, err := os.Open(filePath)
	if err != nil {
		return 0, nil, err
	}

	return stat.Size(), reader, nil
}

func (s *LocalStorage) GetFreeSpace() (int64, error) {
	panic("Unimplemented")
}

type localBlobWriter struct {
	file   *os.File
	hasher hash.Hash
	writer io.Writer

	config   *config.Config
	hash     string
	fileSize int64
	finished bool
}

func makeLocalBlobWriter(config *config.Config) (*localBlobWriter, error) {
	tempFile, err := ioutil.TempFile(config.StorageDir, "temp_blob_")
	if err != nil {
		return nil, err
	}

	hasher := sha256.New()
	writer := io.MultiWriter(tempFile, hasher)

	return &localBlobWriter{
		tempFile,
		hasher,
		writer,
		config,
		"",
		0,
		false,
	}, nil
}

func (s *localBlobWriter) Write(p []byte) (int, error) {
	return s.writer.Write(p)
}

func (s *localBlobWriter) Close() error {
	s.file.Close()
	s.hash = hex.EncodeToString(s.hasher.Sum(nil))
	return nil
}

func (s *localBlobWriter) GetHash() string {
	return s.hash
}

func (s *localBlobWriter) GetFileSize() int64 {
	return s.fileSize
}

func (s *localBlobWriter) Finish() error {
	newpath, err := buildPathFromHash(s.config.StorageDir, s.hash)
	if err != nil {
		return err
	}
	err = os.Rename(s.file.Name(), newpath)
	if err != nil {
		s.finished = true
	}
	return err
}

func (s *localBlobWriter) Discard() error {
	if !s.finished {
		return os.Remove(s.file.Name())
	}
	return nil
}
