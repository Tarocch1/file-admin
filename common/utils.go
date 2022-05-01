package common

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func GetRootDir() {
	var dir string
	var err error
	if filepath.IsAbs(FlagDir) {
		dir = FlagDir
	} else {
		dir, err = filepath.Abs(FlagDir)
		if err != nil {
			panic(err)
		}
	}
	isDir := PathIsDir(dir)
	if isDir {
		RootDir = dir
	} else {
		panic(errors.New(dir + " is not a dir"))
	}
}

func GetWorkingPath(path string) (string, error) {
	if strings.HasPrefix(path, "..") {
		return "", errors.New("path is out of root dir")
	}
	return filepath.Abs(filepath.Join(RootDir, path))
}

func PathNotExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	}
	return false
}

func PathIsDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}