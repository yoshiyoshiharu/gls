package gls

import (
	"os"
)

func currentDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return dir, nil
}

func DirFiles(dir string) ([]os.DirEntry, error) {
	fileInfos, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	return fileInfos, nil
}

func CurrentDirFiles() ([]os.DirEntry, error) {
	currentDir, err := currentDir()
	if err != nil {
		return nil, err
	}

	files, err := DirFiles(currentDir)
	if err != nil {
		return nil, err
	}

	return files, nil
}
