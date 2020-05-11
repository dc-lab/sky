package common

import (
	"fmt"
	"os"
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

func PathExists(path string, directoryFlag bool) (bool, error) {
	stat, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) || (directoryFlag && !stat.IsDir()) {
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
