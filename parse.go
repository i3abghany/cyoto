package main

import (
	"Kyoto/pkg/parser"
	"github.com/antlr4-go/antlr/v4"
)

type KyotoListener struct {
	*parser.BaseKyotoGrammarListener
}

func (s *KyotoListener) VisitErrorNode(node antlr.ErrorNode) {
	println("Error: ", node.GetText())
}

func (s *KyotoListener) EnterProgram(ctx *parser.ProgramContext) {
	println("Program: ", ctx.GetText())
}

func parse(code string) {
	inputStream := antlr.NewInputStream(code)
	lexer := parser.NewKyotoGrammarLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewKyotoGrammarParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(&KyotoListener{}, p.Program())
}
