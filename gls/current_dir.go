package gls

import (
	"os"
)

func DirFiles(dir string) ([]os.DirEntry, error) {
	fileInfos, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	return fileInfos, nil
}

func CurrentDirFiles() ([]os.DirEntry, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	files, err := DirFiles(currentDir)
	if err != nil {
		return nil, err
	}

	return files, nil
}
