package main

import (
	"github.com/yoshiyoshiharu/gls/gls"
	"github.com/fatih/color"
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

			c.Println(file.Name(), gls.CountFiles(file))
		} else {
			color.White(file.Name())
		}
	}
}
