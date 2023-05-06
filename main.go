package main

import (
	"log"

	"github.com/fatih/color"
	"github.com/yoshiyoshiharu/gls/gls"
)

func main() {
	currentDirFiles, err := gls.CurrentDirFiles()

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range currentDirFiles {
		if file.IsDir() {
			c := color.New(color.FgBlue)

			childFiles, err := gls.ChildFiles((file))
			if err != nil {
				log.Fatal(err)
			}

			c.Printf("%s(%d)\n", file.Name(), len(childFiles))
		} else {
			color.White(file.Name())
		}
	}
}
