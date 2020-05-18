package common

import (
	"log"
	"os"
)

func DealWithError(msg string, err error) {
	if err != nil {
		log.Printf("%s\n%v", msg, err)
	}
}

func DieWithError(msg string, err error) {
	if err != nil {
		log.Fatalf("%s\n%v", msg, err)
	}
}

func PathExists(path string, shouldBeDir bool) (bool, error) {
	stat, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) || (shouldBeDir && !stat.IsDir()) {
		return false, nil
	}
	return true, err
}

func MakeDir(dirPath string, removeIfExist bool) error {
	if exist, err := PathExists(dirPath, false); exist && removeIfExist {
		DealWithError("Unexpected PathError during path check", err)
		_ = RemoveDir(dirPath)
	}
	return os.MkdirAll(dirPath, 0755)
}

func RemoveDir(dirPath string) error {
	return os.RemoveAll(dirPath)
}