package main

import (
	"Kyoto/pkg/parser"
	"github.com/antlr4-go/antlr/v4"
)

func Parse(code string) parser.IProgramContext {
	inputStream := antlr.NewInputStream(code)
	lexer := parser.NewKyotoGrammarLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewKyotoGrammarParser(stream)
	return p.Program()
}
