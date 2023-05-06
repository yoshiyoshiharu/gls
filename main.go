package main

import (
	"github.com/fatih/color"
	"github.com/yoshiyoshiharu/gls/gls"
	"log"
)

func main() {
	currentDirFiles, err := gls.CurrentDirFiles()

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range currentDirFiles {
		if file.IsDir() {
			c := color.New(color.FgBlue)

			file_count, err := gls.CountFiles((file))
			if err != nil {
				log.Fatal(err)
			}

			c.Printf("%s(%d)\n", file.Name(), file_count)
		} else {
			color.White(file.Name())
		}
	}
}
