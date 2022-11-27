package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

const (
	// 比较操作符
	Gt = ">"
	Ge = ">="
	Lt = "<"
	Le = "<="
	Eq = "="

	// 逻辑运算符
	Not = "not"
	And = "and"
	Or  = "or"

	// near运算符
	Near = "near"

	// 关键词
	WordExprKey = "wordExpr"
	NearExprKey = "nearExpr"
	CmpExprKey  = "cmpExpr"
	CmpOpKey    = "cmpOp"
)

// 比较表达式: keyword {op} num
// 单个词等价于：keyword >= 1
type CmpExpr struct {
	Word string // 关键词
	Op   string // 比较操作符: <, <=, >, >=, =
	Num  int    // 数量
}

// Near表达式：word1 near/10 word2
type NearExpr struct {
	LeftWord  string // near左边的关键词
	RightWord string // near右边的关键词
	Dist      int    // 距离
}

// 节点
type Node struct {
	Op        string // 逻辑操作符：and, or, not, 空字符串
	Father    *Node  // 父节点
	children  []Node // 如果是叶子节点，则没有子节点
	CmpExprs  []CmpExpr
	NearExprs []NearExpr
}

// 树
var Tree = Node{}

func (n *Node) Walk(t antlr.Tree) {
	n.Enter(t)
	for i := 0; i < t.GetChildCount(); i++ {
		child := t.GetChild(i)
		n.Walk(child)
	}
	n.Exit(t)
}

func (n *Node) Enter(t antlr.Tree) (newNode Node, err error) {
	var rule string
	if rule == "prog" || rule == "expr" {
		newNode = Node{}
		n.children = append(n.children, newNode)
		return
	}
	if rule == CmpExprKey {
		// 比较表达式
		newNode = Node{}
		n.children = append(n.children, newNode)
		return
	}
	return
}

func (n *Node) Exit(t antlr.Tree) (err error) {
	return
}
