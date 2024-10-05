package main

import "os"

func readFile(fileName string) string {
	b, err := os.ReadFile(fileName)
	if err != nil {
		println("Error reading file: ", err)
		return ""
	}
	return string(b)
}
