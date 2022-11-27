// https://tonybai.com/2022/05/25/an-example-of-implement-dsl-using-antlr-and-go-part2/
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"

	"parser"
)

var (
	// 进入深度
	depth = 0
	// 测试标志
	debug = true
)

type TraceListener struct {
	*parser.BaseMatchListener
	p *parser.MatchParser
	t antlr.Tree
}

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

func ParseExpr(expr string) {
	input := antlr.NewInputStream(expr)
	lexer := parser.NewMatchLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewMatchParser(stream)
	tree := p.Prog()
	antlr.ParseTreeWalkerDefault.Walk(NewTraceListener(p, tree), tree)
}

func NewTraceListener(p *parser.MatchParser, t antlr.Tree) *TraceListener {
	fmt.Printf("--%+v\n", p)
	fmt.Printf("--%+v\n", t)
	return &TraceListener{
		p: p,
		t: t,
	}
}

func (l *TraceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	i := ctx.GetRuleIndex()
	ruleName := l.p.RuleNames[i]
	if debug {
		fmt.Printf("%s==> %s 《 %s 》\n", strings.Repeat(" ", depth*2), ruleName, ctx.GetText())
	}
	depth += 1
}

func (l *TraceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	depth -= 1
	i := ctx.GetRuleIndex()
	ruleName := l.p.RuleNames[i]
	if debug {
		fmt.Printf("%s<== %s\n", strings.Repeat(" ", depth*2), ruleName)
	}
}
