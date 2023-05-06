package gls

import (
	"os"
)

func CurrentDirFiles() ([]os.DirEntry, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	return files, nil
}
