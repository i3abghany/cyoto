package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadFile(fileName string) string {
	b, err := os.ReadFile(fileName)
	if err != nil {
		println("Error reading file: ", err)
		return ""
	}
	return string(b)
}

type TestCase struct {
	Name     string
	Code     string
	Expected int
}

func readTest(name string) []TestCase {

	content := ReadFile(fmt.Sprintf("tests\\inputs\\%s.kyo", name))
	cases := strings.Split(content, "// END")

	var ret = make([]TestCase, len(cases))
	for i, c := range cases {
		c := strings.TrimSpace(c)
		lines := strings.Split(c, "\n")
		_, err1 := fmt.Sscanf(lines[0], "// NAME %s", &ret[i].Name)
		_, err2 := fmt.Sscanf(lines[1], "// RET %d", &ret[i].Expected)
		if err1 != nil || err2 != nil {
			log.Panicf("failed to parse test case: %v", err1)
		}
		ret[i].Code = strings.Join(lines[2:], "\n")
	}

	return ret
}
