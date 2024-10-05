package main

import (
	"Kyoto/pkg/parser"
	"github.com/antlr4-go/antlr/v4"
)

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
