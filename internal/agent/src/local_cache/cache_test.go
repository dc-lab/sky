package local_cache

import (
	"path"
	"testing"

	agent "github.com/dc-lab/sky/internal/agent/src"
	"github.com/dc-lab/sky/internal/agent/src/common"
	"github.com/dc-lab/sky/internal/agent/src/parser"
	"github.com/docker/distribution/uuid"
	"github.com/stretchr/testify/assert"
)

func InitializeAgentConfig(maxCacheSize uint64) (*parser.Config, error) {
	return parser.InitializeAgentConfigFromCustomFields(map[string]interface{}{
		"LogsDirectory":          "/tmp/agent-logs-test",
		"RunDirectory":           "/tmp/agent_test/",
		"ResourceManagerAddress": "localhost:5051",
		"AgentDirectory":         "/tmp/agent_test",
		"TokenPath":              "/tmp/sample_token",
		"MaxCacheSize":           maxCacheSize,
	})
}

func InitializeFakeAgent(maxCacheSize uint64) (*agent.App, error) {
	config, err := InitializeAgentConfig(maxCacheSize)
	if err != nil {
		return nil, err
	}

	return agent.NewApp(config)
}

func WriteTextToFile(text []byte, filePath string) (int, error) {
	file := common.CreateFile(filePath)
	defer file.Close()
	fileContent := []byte(text)
	return file.Write(fileContent)
}

func TestGetCacheFileReader(t *testing.T) {
	agent, err := InitializeFakeAgent(1024 * 1024)
	assert.NoError(t, err)
	assert.NotNil(t, agent)

	cache, err := NewCache(agent.Config)
	assert.NoError(t, err)
	assert.NotNil(t, cache)

	randString := uuid.Generate().String()
	filePath := path.Join("/tmp", randString)
	_, err = cache.GetCacheFileReader(randString)
	assert.NoError(t, err)

	fileContent := []byte("text")
	writeBytes, err := WriteTextToFile(fileContent, filePath)
	assert.NoError(t, err)

	reader, err := cache.GetCacheFileReader(randString)
	assert.NoError(t, err)
	assert.NotNil(t, reader)
	cachedContent := make([]byte, writeBytes)
	readBytes, err := reader.Read(cachedContent)
	assert.Equal(t, readBytes, writeBytes)
	assert.Equal(t, cachedContent, fileContent)
}

func TestDeleteOldCacheFilesIfNeeded(t *testing.T) {
	agent, err := InitializeFakeAgent(1)
	assert.NoError(t, err)
	assert.NotNil(t, agent)

	cache, err := NewCache(agent.Config)
	assert.NoError(t, err)
	assert.NotNil(t, cache)

	randString := uuid.Generate().String()
	filePath := path.Join(agent.Config.LocalCacheDirectory, randString)
	maxCacheSize := agent.Config.MaxCacheSize
	cacheSize := cache.GetDiskUsage()
	assert.Less(t, cacheSize, maxCacheSize)

	fileContent := []byte("text")
	WriteTextToFile(fileContent, filePath)

	cacheSize = cache.GetDiskUsage()
	assert.Greater(t, cacheSize, maxCacheSize)

	cache.DeleteOldCacheFilesIfNeeded()
	cacheSize = cache.GetDiskUsage()
	assert.Less(t, cacheSize, maxCacheSize)
}
