package gls

import(
	"flag"
)

func Run() {
	maxNest := flag.Int("n", 0, "Max depth of the tree")
	flag.Parse()

	DisplayCurrentDir(*maxNest)
}
