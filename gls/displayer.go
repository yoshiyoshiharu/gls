package gls

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func DisplayCurrentDir() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	_, _, err = displayRecursivelyDir(dir, 0)
	if err != nil {
		return err
	}

	return nil
}

func displayRecursivelyDir(dir string, nest int) (os.DirEntry, int, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, 0, err
	}

	for _, file := range files {
		path := filepath.Join(dir, file.Name())
		if file.IsDir() {
			c := color.New(color.FgBlue)

			childFiles, err := os.ReadDir(path)
			if err != nil {
				log.Fatal(err)
				return nil, 0, err
			}

			c.Printf("%s(%d)\n",  strings.Repeat("  ", nest) + "└──" + file.Name(), len(childFiles))
			displayRecursivelyDir(path, nest + 1)
		} else {
			color.White(strings.Repeat("  ", nest) + "└──" + file.Name())
		}
	}

	return nil, nest, nil
}
