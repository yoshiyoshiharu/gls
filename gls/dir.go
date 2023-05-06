package gls

import(
	"os"
)

func ChildFiles(dir string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(dir);

	if err != nil {
		return nil, err
	}

	return files, nil
}
