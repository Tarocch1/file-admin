package common

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func GetWorkDir() {
	var dirPath string
	var err error
	if filepath.IsAbs(FlagDir) {
		dirPath = FlagDir
	} else {
		dirPath, err = filepath.Abs(FlagDir)
		if err != nil {
			panic(err)
		}
	}
	isDir := PathIsDir(dirPath)
	if isDir {
		WorkDir = dirPath
	} else {
		panic(errors.New(dirPath + " is not a dir"))
	}
}

func GetTargetPath(path string) (string, error) {
	if strings.HasPrefix(path, "..") {
		return "", errors.New("path is out of root dir")
	}
	return filepath.Abs(filepath.Join(WorkDir, path))
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
