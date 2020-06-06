package local_cache

import (
	"github.com/dc-lab/sky/agent/src/common"
	"github.com/dc-lab/sky/agent/src/hardware"
	"github.com/dc-lab/sky/agent/src/parser"
	"github.com/docker/distribution/uuid"
	"github.com/stretchr/testify/assert"
	"path"
	"testing"
)

func InitializeAgentConfig(maxCacheSize uint64) {
	parser.InitializeAgentConfigFromCustomFields(map[string]interface{}{
		"LogsDirectory":          "/tmp/agent-logs-test",
		"RunDirectory":           "/tmp/agent_test/",
		"ResourceManagerAddress": "localhost:5051",
		"AgentDirectory":         "/tmp/agent_test",
		"TokenPath":              "/tmp/sample_token",
		"MaxCacheSize":           maxCacheSize,
	})
}

func WriteTextToFile(text []byte, filePath string) (int, error) {
	file := common.CreateFile(filePath)
	defer file.Close()
	fileContent := []byte(text)
	return file.Write(fileContent)
}

func TestGetCacheFileReader(t *testing.T) {
	InitializeAgentConfig(1024 * 1024)
	randString := uuid.Generate().String()
	filePath := path.Join("/tmp", randString)
	_, err := GetCacheFileReader(randString)
	assert.NotEqual(t, err, nil)

	fileContent := []byte("text")
	writeBytes, err := WriteTextToFile(fileContent, filePath)
	assert.Equal(t, err, nil)

	reader, err := GetCacheFileReader(randString)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, reader, nil)
	cachedContent := make([]byte, writeBytes)
	readBytes, err := reader.Read(cachedContent)
	assert.Equal(t, readBytes, writeBytes)
	assert.Equal(t, cachedContent, fileContent)
}

func TestDeleteOldCacheFilesIfNeeded(t *testing.T) {
	InitializeAgentConfig(1)
	randString := uuid.Generate().String()
	filePath := path.Join(parser.AgentConfig.LocalCacheDirectory, randString)
	cacheDir := parser.AgentConfig.LocalCacheDirectory
	maxCacheSize := parser.AgentConfig.MaxCacheSize
	cacheSize := hardware.GetDiskUsage(cacheDir)
	assert.Less(t, cacheSize, maxCacheSize)

	fileContent := []byte("text")
	WriteTextToFile(fileContent, filePath)

	cacheSize = hardware.GetDiskUsage(cacheDir)
	assert.Less(t, maxCacheSize, cacheSize)

	DeleteOldCacheFilesIfNeeded()
	cacheSize = hardware.GetDiskUsage(cacheDir)
	assert.Less(t, cacheSize, maxCacheSize)
}
