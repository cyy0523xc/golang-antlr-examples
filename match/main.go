// https://tonybai.com/2022/05/25/an-example-of-implement-dsl-using-antlr-and-go-part2/
package main

import (
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"

	"parser"
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
	// fmt.Printf("%v\n", tree)
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

func (l *TraceListener) EnterNearBaseExpr(ctx antlr.ParserRuleContext) {
	// i := ctx.GetRuleIndex()
	fmt.Printf("====> %#v 《 %s 》\n", l, ctx.GetText())
}

func (l *TraceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	// printLevelPrefix(ctx)
	i := ctx.GetRuleIndex()
	ruleName := l.p.RuleNames[i]
	fmt.Printf("==> %s 《 %s 》\n", ruleName, ctx.GetText())
	// fmt.Printf("%+v\n", l.p)
}

func (l *TraceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	// printLevelPrefix(ctx)
	i := ctx.GetRuleIndex()
	ruleName := l.p.RuleNames[i]
	fmt.Println("<==", ruleName)
}
