package local_cache

import (
	"errors"
	"github.com/dc-lab/sky/internal/agent/src/common"
	"github.com/dc-lab/sky/internal/agent/src/hardware"
	"github.com/dc-lab/sky/internal/agent/src/parser"
	"io"
	"os"
	"path"
	"time"
)

type Cache struct {
	config *parser.Config
}

func NewCache(config *parser.Config) (*Cache, error) {
	return &Cache{config}, nil
}

func (c *Cache) DeleteOldCacheFilesIfNeeded() {
	cacheDir := c.config.LocalCacheDirectory
	cacheSize := hardware.GetDiskUsage(c.config.AgentDirectory)
	for cacheSize > c.config.MaxCacheSize {
		oldestTime := time.Now()
		oldestFile := ""
		consumer := func(fileName string, info os.FileInfo) error {
			if info.ModTime().Before(oldestTime) {
				oldestFile = fileName
			}
			return nil
		}
		err := common.ListChildrenFiles(cacheDir, consumer)
		common.DealWithError(err)
		if val, err := common.PathExists(oldestFile, false); val && err == nil {
			err = os.Remove(oldestFile)
		}
		common.DealWithError(err)
		cacheSize = hardware.GetDiskUsage(cacheDir)
	}
}

func (c *Cache) GetCacheFileReader(fileHash string) (io.Reader, error) {
	c.DeleteOldCacheFilesIfNeeded()
	filePath := path.Join(c.config.LocalCacheDirectory, fileHash)
	err := errors.New("Cache file not found")
	var reader io.Reader
	if val, existsError := common.PathExists(filePath, false); val && existsError == nil {
		reader, err = os.Open(filePath)
	}
	return reader, err
}

func (c *Cache) NewCacheFileWriter(fileHash string) (io.Writer, error) {
	c.DeleteOldCacheFilesIfNeeded()
	filePath := path.Join(c.config.LocalCacheDirectory, fileHash)
	if val, err := common.PathExists(filePath, false); !val && err == nil {
		return common.CreateFile(filePath), err
	}
	return nil, errors.New("File already exists")
}

func (c *Cache) GetDiskUsage() uint64 {
	return hardware.GetDiskUsage(c.config.LocalCacheDirectory)
}
