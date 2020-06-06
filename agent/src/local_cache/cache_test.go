package local_cache

import (
	"github.com/dc-lab/sky/agent/src/common"
	"github.com/docker/distribution/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func WriteTextToFile(text []byte, filePath string) (int, error) {
	file := common.CreateFile(filePath)
	defer file.Close()
	fileContent := []byte(text)
	return file.Write(fileContent)
}

func TestGetCacheFileReader(t *testing.T) {
	filePath := "/tmp/" + uuid.Generate().String()
	_, err := GetCacheFileReader(filePath)
	assert.NotEqual(t, err, nil)

	fileContent := []byte("text")
	writeBytes, err := WriteTextToFile(fileContent, filePath)
	assert.Equal(t, err, nil)

	reader, err := GetCacheFileReader(filePath)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, reader, nil)
	cachedContent := make([]byte, writeBytes)
	readBytes, err := reader.Read(cachedContent)
	assert.Equal(t, readBytes, writeBytes)
	assert.Equal(t, cachedContent, fileContent)
}
