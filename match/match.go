package main

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

// stack []bool 是否生成新节点
func (n *Node) Enter(rule string, text string, stack []bool) (node *Node, err error) {
	if rule == AndOpKey || rule == OrOpKey || rule == NotOpKey {
		// 如：a and b
		// 解释到and的时候，
		node = n
		node.father.Op = rule
		stack = append(stack, false)
		return
	}
	if rule == "word" {
		if n.rule == WordExprKey {
			n.Cmp = &CmpExpr{
				Word: text,
			}
		}
		node = n
		stack = append(stack, false)
		return
	}

	// 生成新节点
	node = &Node{
		rule: rule,
	}
	stack = append(stack, true) // 生成新节点
	if rule == WordExprKey {
		// 单个关键或者带比较符号的简单表达式
		node.Op = "cmp"
		return
	}
	if rule == "prog" || rule == "expr" {
		n.Children = append(n.Children, node)
		return
	}
	if rule == CmpExprKey {
		// 比较表达式
		n.Children = append(n.Children, node)
		return
	}
	return
}

func (n *Node) Exit(rule string, stack []bool) (node *Node, err error) {
	if stack[len(stack)-1] {
		node = n.father
		stack = stack[:len(stack)-1]
	}
	return
}
