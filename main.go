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

	p := Parse(readFile(args[0]))
	i := NewInterpreter()
	ret := i.Interpret(p)
	fmt.Printf("main returned %d\n", ret)
}
