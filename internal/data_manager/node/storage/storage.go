package storage

import "io"

type Storage interface {
	ListBlobs(func(hash string) error) error
	PutBlob() (BlobWriter, error)
	GetBlob(hash string) (int64, io.Reader, error)
	GetFreeSpace() (int64, error)
}

type BlobWriter interface {
	io.WriteCloser

	GetHash() string
	GetFileSize() int64

	Finish() error
	Discard() error
}
