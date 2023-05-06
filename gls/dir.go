package gls

import(
	"os"
	"errors"
)

// TODO
// Dirを構造体にして、childFilesを要素orメソッドで持ちたい

func CountFiles(dir os.DirEntry) (int, error) {
	if !dir.IsDir() { return 0, errors.New("arg is not dir") }

	files, err := os.ReadDir(dir.Name());
	if err != nil {
		return 0, err
	}

	var count int
	for range files {
		count++
	}

	return count, nil
}
