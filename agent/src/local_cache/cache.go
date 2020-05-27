package local_cache

import (
	"errors"
	"github.com/dc-lab/sky/agent/src/common"
	"io"
	"os"
)

func GetCacheFileReader(filePath string) (io.Reader, error) {
	err := errors.New("Cache file not found")
	var reader io.Reader
	if val, err := common.PathExists(filePath, false); val && err != nil {
		reader, err = os.Open(filePath)
	}
	return reader, err
}
