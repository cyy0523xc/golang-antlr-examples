// https://tonybai.com/2022/05/25/an-example-of-implement-dsl-using-antlr-and-go-part2/
package main

import (
	"encoding/json"
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
	p     *parser.MatchParser
	t     antlr.Tree
	n     *Node
	stack *Stack
}

var rootNode = &Node{}

func main() {
	println("input file:", os.Args[1])
	input, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		panic(err)
	}
	parse(input.InputStream)
}

func ParseExpr(expr string) (jsonStr string, err error) {
	input := antlr.NewInputStream(expr)
	return parse(input)
}

func parse(input *antlr.InputStream) (jsonStr string, err error) {
	lexer := parser.NewMatchLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewMatchParser(stream)
	tree := p.Prog()
	antlr.ParseTreeWalkerDefault.Walk(NewTraceListener(p, tree, rootNode), tree)
	if debug {
		bytes, _ := json.Marshal(rootNode)
		fmt.Println("------Before: " + string(bytes))
	}

	// 化简
	rootNode.Simple()
	bytes, err := json.Marshal(rootNode)
	if err != nil {
		return
	}
	jsonStr = string(bytes)
	if debug {
		fmt.Println("------After: " + jsonStr)
	}
	return
}

func NewTraceListener(p *parser.MatchParser, t antlr.Tree, node *Node) *TraceListener {
	return &TraceListener{
		p:     p,
		t:     t,
		n:     node,
		stack: &Stack{},
	}
}

func (l *TraceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	i := ctx.GetRuleIndex()
	ruleName := l.p.RuleNames[i]
	var err error
	l.n, err = l.n.Enter(ruleName, ctx.GetText(), l.stack)
	if err != nil {
		panic(err)
	}
	if debug {
		fmt.Printf("%s%02d==> %s 《 %s 》\n", strings.Repeat(" ", depth*2), depth, ruleName, ctx.GetText())
		fmt.Println(l.stack)
		bytes, _ := json.Marshal(l.n)
		fmt.Println("------" + string(bytes))
	}
	depth += 1
}

func (l *TraceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	depth -= 1
	i := ctx.GetRuleIndex()
	ruleName := l.p.RuleNames[i]
	l.n, _ = l.n.Exit(ruleName, l.stack)
	if debug {
		fmt.Printf("%s%02d<== %s\n", strings.Repeat(" ", depth*2), depth, ruleName)
		fmt.Println(l.stack)
	}
}
