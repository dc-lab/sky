package local_cache

import (
	"errors"
	"github.com/dc-lab/sky/agent/src/common"
	"github.com/dc-lab/sky/agent/src/hardware"
	"github.com/dc-lab/sky/agent/src/parser"
	"io"
	"os"
	"path"
	"time"
)

func DeleteOldCacheFilesIfNeeded() {
	cacheDir := parser.AgentConfig.LocalCacheDirectory
	cacheSize := hardware.GetDiskUsage(cacheDir)
	for cacheSize > parser.AgentConfig.MaxCacheSize {
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

func GetCacheFileReader(fileHash string) (io.Reader, error) {
	DeleteOldCacheFilesIfNeeded()
	filePath := path.Join(parser.AgentConfig.LocalCacheDirectory, fileHash)
	err := errors.New("Cache file not found")
	var reader io.Reader
	if val, existsError := common.PathExists(filePath, false); val && existsError == nil {
		reader, err = os.Open(filePath)
	}
	return reader, err
}

func NewCacheFileWriter(fileHash string) (io.Writer, error) {
	DeleteOldCacheFilesIfNeeded()
	filePath := path.Join(parser.AgentConfig.LocalCacheDirectory, fileHash)
	if val, err := common.PathExists(filePath, false); !val && err == nil {
		return common.CreateFile(filePath), err
	}
	return nil, errors.New("File already exists")
}
