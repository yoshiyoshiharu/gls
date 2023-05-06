package gls

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func DisplayCurrentDir() {
	currentDirFiles, err := CurrentDirFiles()

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range currentDirFiles {
		if file.IsDir() {
			c := color.New(color.FgBlue)

			childFiles, err := ChildFiles(file.Name())
			if err != nil {
				log.Fatal(err)
			}

			c.Printf("%s(%d)\n", file.Name(), len(childFiles))
			recursivelyDisplayDir(file.Name(), 1)
		} else {
			color.White(file.Name())
		}
	}
}

func recursivelyDisplayDir(dir string, nl int) (os.DirEntry, int, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, 0, err
	}

	for _, file := range files {
		path := filepath.Join(dir, file.Name())
		if file.IsDir() {
			c := color.New(color.FgBlue)

			childFiles, err := ChildFiles(path)
			if err != nil {
				log.Fatal(err)
				return nil, 0, err
			}

			c.Printf("%s(%d)\n",  strings.Repeat("--", nl) + file.Name(), len(childFiles))
			recursivelyDisplayDir(path, nl + 1)
		} else {
			color.White(strings.Repeat("--", nl) + file.Name())
		}
	}

	return nil, nl, nil
}
