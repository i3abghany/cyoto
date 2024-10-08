package main

import (
	"Kyoto/pkg/parser"
	"github.com/antlr4-go/antlr/v4"
	"log"
	"strconv"
)

type KyotoVisitor struct {
	*parser.BaseKyotoGrammarVisitor
	pInterpreter *Interpreter
}

func (v *KyotoVisitor) Visit(ctx antlr.ParseTree) interface{} { return ctx.Accept(v) }

func (v *KyotoVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	for _, c := range ctx.AllFunctionDeclaration() {
		c, ok := c.(*parser.FunctionDeclarationContext)
		if !ok {
			log.Panic("not a function declaration")
		}

		if _, ok := v.pInterpreter.Functions[c.IDENTIFIER().GetText()]; ok {
			log.Panicf("function redifinition: %s", c.IDENTIFIER().GetText())
		} else {
			v.pInterpreter.Functions[c.IDENTIFIER().GetText()] = c
		}
	}

	main := v.pInterpreter.Functions["main"]
	if main == nil {
		log.Panicf("no main function")
	}

	ret := v.VisitFunctionDeclaration(main)
	if ret == nil {
		return 0
	} else {
		return ret
	}
}

func (v *KyotoVisitor) VisitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) interface{} {
	b := ctx.Block()
	res := v.VisitBlock(b.(*parser.BlockContext))
	if res != nil {
		return res
	}
	return nil
}

func (v *KyotoVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	switch t := ctx.GetChild(0).(type) {
	case *parser.VariableDeclarationContext:
		return v.VisitVariableDeclaration(t)
	case *parser.ReturnStatementContext:
		return v.VisitReturnStatement(t)
	case *parser.IfStatementContext:
		return v.VisitIfStatement(t)
	case *parser.BlockContext:
		return v.VisitBlock(t)
	case *parser.ExpressionContext:
		return v.visitExpressionStatement(t)
	default:
		log.Panicf("unsupported statement: %T", t)
	}
	return nil
}

func (v *KyotoVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	for _, s := range ctx.AllStatement() {
		res := v.VisitStatement(s.(*parser.StatementContext))
		if res != nil {
			return res
		}
	}
	return nil
}

func (v *KyotoVisitor) VisitReturnStatement(ctx *parser.ReturnStatementContext) interface{} {
	return v.Visit(ctx.Expression())
}

func (v *KyotoVisitor) VisitVariableDeclaration(ctx *parser.VariableDeclarationContext) interface{} {
	name := ctx.IDENTIFIER().GetText()

	v.pInterpreter.Variables[name] = Variable{
		Name:  name,
		Type:  ctx.Ktype().GetText(),
		Value: v.Visit(ctx.Expression()),
	}
	return nil
}

func (v *KyotoVisitor) VisitLiteralExpr(ctx *parser.LiteralExprContext) interface{} {
	switch t := ctx.GetChild(0).(type) {
	case *parser.IntLiteralContext:
		i, err := strconv.Atoi(t.GetText())
		if err != nil {
			log.Panicf("failed to parse int literal: %v", err)
		}
		return i
	default:
		log.Panicf("unsupported literal: %T", t)
	}
	return nil
}

func (v *KyotoVisitor) VisitVariableExpr(ctx *parser.VariableExprContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	r, ok := v.pInterpreter.Variables[name]
	if !ok {
		log.Panicf("variable not found: %s", name)
	}
	return r.Value
}

func (v *KyotoVisitor) VisitMultiplicativeExpr(ctx *parser.MultiplicativeExprContext) interface{} {
	l := v.Visit(ctx.Expression(0))
	r := v.Visit(ctx.Expression(1))
	op := ctx.MultiplicativeOp().GetText()

	switch op {
	case "*":
		return l.(int) * r.(int)
	case "/":
		return l.(int) / r.(int)
	case "%":
		return l.(int) % r.(int)
	default:
		log.Panicf("unsupported multiplicative operator: %s", op)
		return nil
	}
}

func (v *KyotoVisitor) VisitAdditiveExpr(ctx *parser.AdditiveExprContext) interface{} {
	l := v.Visit(ctx.Expression(0))
	r := v.Visit(ctx.Expression(1))
	op := ctx.AdditiveOp().GetText()

	switch op {
	case "+":
		return l.(int) + r.(int)
	case "-":
		return l.(int) - r.(int)
	default:
		log.Panicf("unsupported additive operator: %s", op)
		return nil
	}
}

func (v *KyotoVisitor) VisitComparisonExpr(ctx *parser.ComparisonExprContext) interface{} {
	l := v.Visit(ctx.Expression(0))
	r := v.Visit(ctx.Expression(1))
	op := ctx.ComparisonOp().GetText()

	res := func() bool {
		switch op {
		case "<":
			return l.(int) < r.(int)
		case "<=":
			return l.(int) <= r.(int)
		case ">":
			return l.(int) > r.(int)
		case ">=":
			return l.(int) >= r.(int)
		case "==":
			return l.(int) == r.(int)
		case "!=":
			return l.(int) != r.(int)
		default:
			log.Panicf("unsupported comparison operator: %s", op)
			return false
		}
	}()

	if res {
		return 1
	} else {
		return 0
	}
}

func (v *KyotoVisitor) VisitUnaryExpr(ctx *parser.UnaryExprContext) interface{} {
	r := v.Visit(ctx.Expression())
	op := ctx.UnaryOp().GetText()

	switch op {
	case "-":
		return -r.(int)
	case "+":
		return r.(int)
	default:
		log.Panicf("unsupported unary operator: %s", op)
		return nil
	}
}

func (v *KyotoVisitor) VisitParenthesizedExpr(ctx *parser.ParenthesizedExprContext) interface{} {
	return v.Visit(ctx.Expression())
}

func (v *KyotoVisitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	cond := v.Visit(ctx.Expression())
	if cond.(int) != 0 {
		return v.Visit(ctx.Statement(0))
	} else {
		if ctx.Statement(1) != nil {
			return v.Visit(ctx.Statement(1))
		}
	}
	return nil
}

func (v *KyotoVisitor) VisitFunctionCallExpr(ctx *parser.FunctionCallExprContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	f := v.pInterpreter.Functions[name]
	if f == nil {
		log.Panicf("function not found: %s", name)
	}
	return v.VisitFunctionDeclaration(f)
}

func (v *KyotoVisitor) visitExpressionStatement(t *parser.ExpressionContext) interface{} {
	return v.Visit(t)
}
