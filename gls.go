package main

import (
	"fmt"
	"github.com/yoshiyoshiharu/gls/ls"
	"log"
)

func main() {
	currentDirFiles, err := ls.CurrentDirFiles()

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range currentDirFiles {
		fmt.Println(file.Name())
	}
}
