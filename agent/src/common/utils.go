package common

import (
	"fmt"
	"os"
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
