package main

import (
	"fmt"
	"os"

	"parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func main() {
	println("input file:", os.Args[1])
	input, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		panic(err)
	}

	lexer := parser.NewMatchLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewMatchParser(stream)
	tree := p.Prog()
	antlr.ParseTreeWalkerDefault.Walk(NewTraceListener(p, tree), tree)
}

type TraceListener struct {
	*parser.BaseMatchListener
	p *parser.MatchParser
	t antlr.Tree
}

func NewTraceListener(p *parser.MatchParser, t antlr.Tree) *TraceListener {
	return &TraceListener{
		p: p,
		t: t,
	}
}

func (l *TraceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	// printLevelPrefix(ctx)
	i := ctx.GetRuleIndex()
	ruleName := l.p.RuleNames[i]
	fmt.Printf("==> %s 《 %s 》\n", ruleName, ctx.GetText())
}

func (l *TraceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	// printLevelPrMatchx)
	i := ctx.GetRuleIndex()
	ruleName := l.p.RuleNames[i]
	fmt.Println("<==", ruleName)
}
