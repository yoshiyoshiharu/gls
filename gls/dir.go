package gls

import(
	"os"
	"errors"
)

func ChildFiles(dir os.DirEntry) ([]os.DirEntry, error) {
	if !dir.IsDir() { return nil, errors.New("arg is not dir") }

	files, err := os.ReadDir(dir.Name());
	if err != nil {
		return nil, err
	}

	return files, nil
}
