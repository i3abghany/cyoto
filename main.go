package main

import (
	"Kyoto/pkg/parser"
	"github.com/antlr4-go/antlr/v4"
	"os"
)

func readFile(fileName string) string {
	b, err := os.ReadFile(fileName)
	if err != nil {
		println("Error reading file: ", err)
		return ""
	}
	return string(b)
}

func lex(code string) <-chan antlr.Token {
	inputStream := antlr.NewInputStream(code)
	lexer := parser.NewKyotoGrammarLexer(inputStream)
	ch := make(chan antlr.Token)
	go func() {
		for {
			t := lexer.NextToken()
			if t.GetTokenType() == antlr.TokenEOF {
				break
			}
			ch <- t
		}
		close(ch)
	}()
	return ch
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		println("Usage: ./cyoto <SRC_FILES>")
		return
	}

	for _, file := range args {
		fileContent := readFile(file)

		for token := range lex(fileContent) {
			println(token.GetText())
		}
	}
}
