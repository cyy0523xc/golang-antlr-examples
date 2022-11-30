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

func main() {
	println("input file:", os.Args[1])
	input, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		panic(err)
	}
	rootNode := &Node{}
	rootNode.parse(input.InputStream)

	// 打印字符串
	bytes, err := json.Marshal(rootNode)
	if err != nil {
		fmt.Println("json处理出错")
	}
	fmt.Println("------> " + FmtJson(bytes))
}

func ParseExpr(expr string) *Node {
	rootNode := &Node{}
	input := antlr.NewInputStream(expr)
	rootNode.parse(input)
	return rootNode
}

func (root *Node) parse(input *antlr.InputStream) {
	lexer := parser.NewMatchLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewMatchParser(stream)
	errorListener := &ErrorListener{
		root: root,
	}
	p.AddErrorListener(errorListener)
	tree := p.Prog() // 可能会产生解释错误，如：line 1:0 token recognition error at: 'k'
	if errorListener.root.Err != nil {
		return
	}

	antlr.ParseTreeWalkerDefault.Walk(NewTraceListener(p, tree, root), tree)
	if debug {
		bytes, _ := json.Marshal(root)
		fmt.Println("------Before: " + FmtJson(bytes))
	}

	// 处理Near
	root.OptimNear()
	if debug {
		bytes, _ := json.Marshal(root)
		fmt.Println("------After optim near: " + FmtJson(bytes))
	}

	// 化简
	root.Simple()
	if debug {
		bytes, _ := json.Marshal(root)
		fmt.Println("------After simple: " + FmtJson(bytes))
	}
	// 逻辑运算化简
	root.SimpleLogic()
	if debug {
		bytes, _ := json.Marshal(root)
		fmt.Println("------After: " + FmtJson(bytes))
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
		// fmt.Println(l.stack)
		bytes, _ := json.Marshal(l.n)
		fmt.Printf("depth: %d, node: %s\n", l.stack.depth, string(bytes))
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
