package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s <SRC_FILES>", os.Args[0])
		return
	}

	p := Parse(ReadFile(args[0]))
	i := NewInterpreter()
	i.Interpret(p)
}
