package common

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func DealWithError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func DieWithError(err error) {
	if err != nil {
		log.Fatalln(err)
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

func CreateDirectory(dirPath string, removeIfExist bool) error {
	if exist, err := PathExists(dirPath, false); exist && removeIfExist {
		DealWithError(err)
		_ = RemoveDirectory(dirPath)
	}
	return os.MkdirAll(dirPath, 0755)
}

func RemoveDirectory(dirPath string) error {
	return os.RemoveAll(dirPath)
}

func ConvertToRelativePath(rootDir string, file string) string {
	newFile, err := filepath.Rel(rootDir, file)
	DealWithError(err)
	return newFile
}

func CurrentTimestampMillisString() string {
	ts := time.Now().UnixNano() / int64(time.Millisecond)
	return strconv.FormatInt(ts, 10)
}
