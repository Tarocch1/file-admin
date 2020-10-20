package main

import (
	"errors"
	"os"
	"path/filepath"
)

func getWorkDir() (string, error) {
	if filepath.IsAbs(flagDir) {
		return flagDir, nil
	}
	dirPath, err := filepath.Abs(flagDir)
	if err != nil {
		return "", err
	}
	isDir, err := pathIsDir(dirPath)
	if err != nil {
		return "", err
	}
	if isDir {
		return dirPath, nil
	}
	return "", errors.New(dirPath + " is not a dir.")
}

func getTargetPath(path string) (string, error) {
	return filepath.Abs(filepath.Join(workDir, path))
}

func pathNotExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	}
	return false
}

func pathIsDir(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}
