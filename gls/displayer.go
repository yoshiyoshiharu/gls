package gls

import(
	"log"

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

			childFiles, err := ChildFiles((file))
			if err != nil {
				log.Fatal(err)
			}

			c.Printf("%s(%d)\n", file.Name(), len(childFiles))
		} else {
			color.White(file.Name())
		}
	}
}
