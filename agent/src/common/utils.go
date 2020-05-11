package common

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func DealWithError(err error) {
	if err != nil {
		fmt.Println(err)
		// TODO(glebx777): add logging
	}
}

func CreateFile(filePath string) *os.File {
	stdoutFile, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	return stdoutFile
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func GetChildrenFilePaths(rootDir string) []string {
	var files []string
	err := filepath.Walk(rootDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				files = append(files, path)
			}
			return nil
		})
	DealWithError(err)
	return files
}

func ConvertToRelativePath(rootDir string, file string) string {
	newFile, err := filepath.Rel(rootDir, file)
	DealWithError(err)
	return newFile
}

func GetExecutionDirForTaskId(rootDir string, task_id string) string {
	return path.Join(rootDir, task_id)
}
