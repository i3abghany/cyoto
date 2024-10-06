package main

import (
	"Kyoto/pkg/parser"
	"github.com/antlr4-go/antlr/v4"
	"log"
)

type Variable struct {
	Name  string
	Type  string
	Value interface{}
}

type Interpreter struct {
	V         KyotoVisitor
	Variables map[string]Variable
	Functions map[string]*parser.FunctionDeclarationContext
}

func NewInterpreter() *Interpreter {
	ret := &Interpreter{
		V:         KyotoVisitor{pInterpreter: nil},
		Variables: make(map[string]Variable),
		Functions: make(map[string]*parser.FunctionDeclarationContext),
	}
	ret.V.pInterpreter = ret
	return ret
}

func (i *Interpreter) Interpret(tree antlr.ParseTree) int {
	switch t := tree.(type) {
	case *parser.ProgramContext:
		ret := i.V.VisitProgram(t)
		return ret.(int)
	default:
		log.Panicf("starting node `%T` is not a `program` ", t)
	}

	return -1
}
