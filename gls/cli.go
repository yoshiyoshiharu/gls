package gls

import(
	"flag"
)

func Run() {
	nest := flag.Int("n", 0, "Max depth of the tree")
	flag.Parse()

	DisplayCurrentDir(*nest)
}
