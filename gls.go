package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var files []os.DirEntry

	dir, err := currentDir()
	if err != nil {
		log.Fatal(err)
	}

	files, err = currentDirFiles(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		fmt.Println(file.Type())
	}
}

func currentDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return dir, nil
}

func currentDirFiles(dir string) ([]os.DirEntry, error) {
	fileInfos, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	return fileInfos, nil
}
