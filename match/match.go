package main

import (
	"fmt"
	"strconv"
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
	AndOpKey    = "andOp"
	OrOpKey     = "orOp"
	NotOpKey    = "notOp"
	WordKey     = "word"
	ExprKey     = "expr"
	NearOpKey   = "nearOp"
	LeftOpKey   = "leftOp"
	RightOpKey  = "rightOp"
)

// 比较表达式: keyword {op} num
// 单个词等价于：keyword >= 1
type CmpExpr struct {
	Word string `json:"word"` // 关键词
	Op   string `json:"op"`   // 比较操作符: <, <=, >, >=, =
	Num  int    `json:"num"`  // 数量
	flag int    // 当前解释的位置
}

// Near表达式：word1 near/10 word2
type NearExpr struct {
	LeftWords  []string `json:"left"`  // near左边的关键词
	RightWords []string `json:"right"` // near右边的关键词
	Dist       int      `json:"dist"`  // 距离
	flag       int      // 当前解释的位置
}

// 节点
type Node struct {
	father *Node  // 父节点
	rule   string // 表达式的规则名

	// 逻辑操作符：and, or, not, cmp, near, 空字符串
	// 其中：cmp和near是叶子节点
	Op       string    `json:"op"`
	Cmp      *CmpExpr  `json:"cmp"`
	Near     *NearExpr `json:"near"`
	Children []*Node   `json:"children"` // 如果是叶子节点，则没有子节点
}

type Stack struct {
	data []bool
}

func (s *Stack) Push(val bool) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() bool {
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}

// stack []bool 是否生成新节点
func (n *Node) Enter(rule string, text string, stack *Stack) (node *Node, err error) {
	if rule == AndOpKey || rule == OrOpKey || rule == NotOpKey {
		// 如：a and b
		// 解释到and的时候，
		node = n
		node.Op = rule[:len(rule)-2]
		fmt.Printf("%+v\n", node)
		fmt.Printf("%+v\n", node.father)
		stack.Push(false)
		return
	} else if rule == WordKey {
		// 词
		if n.rule == WordExprKey || n.rule == CmpExprKey {
			if text[0] == '"' || text[0] == '\'' {
				text = text[1 : len(text)-1]
			}
			n.Cmp = &CmpExpr{
				Word: text,
				Op:   ">=",
				Num:  1,
			}
		} else if n.rule == NearExprKey {
			if text[0] == '"' || text[0] == '\'' {
				text = text[1 : len(text)-1]
			}
			if n.Near == nil {
				n.Near = &NearExpr{
					LeftWords:  []string{},
					RightWords: []string{},
				}
			}
			if n.Near.flag == 1 {
				n.Near.RightWords = append(n.Near.RightWords, text)
			} else {
				n.Near.LeftWords = append(n.Near.LeftWords, text)
			}
		}
		node = n
		stack.Push(false)
		return
	} else if rule == CmpOpKey {
		n.Cmp.Op = text
		node = n
		stack.Push(false)
		return
	} else if rule == "digits" {
		// fmt.Printf("cmp: %+v\n", n)
		if n.Cmp != nil {
			n.Cmp.Num, err = strconv.Atoi(text)
		} else {
			n.Near.Dist, err = strconv.Atoi(text)
		}
		if err != nil {
			return
		}
		node = n
		stack.Push(false)
		return
	} else if rule == NearOpKey {
		fmt.Printf("near: %+v\n", n)
		n.Near.flag = 1
		node = n
		stack.Push(false)
		return
	} else if rule == LeftOpKey {
		node = n
		stack.Push(false)
		return
	} else if rule == RightOpKey {
		node = n
		stack.Push(false)
		return
	}

	// 生成新节点
	node = &Node{
		father: n,
		rule:   rule,
	}
	if n.Children == nil {
		n.Children = []*Node{
			node,
		}
	} else {
		n.Children = append(n.Children, node)
	}
	stack.Push(true)
	if rule == WordExprKey || rule == CmpExprKey {
		// 单个关键或者带比较符号的简单表达式
		node.Op = "cmp"
		return
	} else if rule == NearExprKey {
		node.Op = "near"
		return
	} else if rule == "prog" || rule == "expr" {
		return
	} else if rule == CmpExprKey {
		// 比较表达式
		return
	} else {
		err = fmt.Errorf("[ERROR] rule name = %s", rule)
	}
	return
}

func (n *Node) Exit(rule string, stack *Stack) (node *Node, err error) {
	val := stack.Pop()
	if val {
		node = n.father
	} else {
		node = n
	}
	return
}

func (n *Node) Simple() {
	if n.Children == nil {
		return
	}
	if len(n.Children) == 1 {
		// 只有一个子节点
		// fmt.Println("==================")
		child := n.Children[0]
		n.rule = child.rule
		n.Op = child.Op
		n.Cmp = child.Cmp
		n.Near = child.Near
		n.Children = child.Children
		for _, sun := range child.Children {
			sun.father = n
		}
		child = nil
		n.Simple()
	}

	// 递归处理子节点
	for _, child := range n.Children {
		child.Simple()
	}
}
