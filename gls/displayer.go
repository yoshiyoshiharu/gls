package gls

import (
	"log"
	"os"
	"path/filepath"

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
			recursivelyDisplayDir(file.Name())
		} else {
			color.White(file.Name())
		}
	}
}

func recursivelyDisplayDir(dir string) (os.DirEntry, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		path := filepath.Join(dir, file.Name())
		if file.IsDir() {
			c := color.New(color.FgBlue)
			println("aaaa")

			childFiles, err := ChildFiles(path)
			if err != nil {
				log.Fatal(err)
				return nil, err
			}
			println("aaa")

			c.Printf("%s(%d)\n", file.Name(), len(childFiles))
			recursivelyDisplayDir(path)
		} else {
			color.White("--" + file.Name())
		}
	}

	return nil, nil
}
