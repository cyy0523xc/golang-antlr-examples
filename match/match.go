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
	father    *Node // 父节点
	rule      string
	Op        string     `json:"op"`       // 逻辑操作符：and, or, not, 空字符串
	Children  []*Node    `json:"children"` // 如果是叶子节点，则没有子节点
	CmpExprs  []CmpExpr  `json:"cmp"`
	NearExprs []NearExpr `json:"near"`
}

func (n *Node) Enter(rule string) (newNode *Node, err error) {
	if rule == AndOpKey || rule == OrOpKey || rule == NotOpKey {
		n.father.Op = rule
		return n, err
	}
	newNode = &Node{
		rule: rule,
	}
	if rule == "prog" || rule == "expr" {
		n.Children = append(n.Children, newNode)
		return
	}
	if rule == CmpExprKey {
		// 比较表达式
		n.Children = append(n.Children, newNode)
		return
	}
	return
}

func (n *Node) Exit(rule string) (err error) {
	return
}
