package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprint(os.Stderr, "Usage: ./cyoto <SRC_FILES>")
		return
	}

	parse(readFile(args[0]))
}
