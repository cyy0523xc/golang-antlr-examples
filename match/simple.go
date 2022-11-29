package main

// 化简表达式
func (n *Node) Simple() {
	if n.Children == nil {
		return
	}
	if len(n.Children) == 1 {
		// 只有一个子节点: 可以直接将孙节点直接连接到该节点上
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

// 化简嵌套的逻辑表达式
func (n *Node) SimpleLogic() {
	// a and (b and c) => and(a,b,c)
	// a or (b or c) => or(a,b,c)
	if n.Children == nil {
		return
	}
	if !(n.Op == "and" || n.Op == "or" || n.Op == "") {
		return
	}
	hasAnd, hasOr, hasChild := false, false, false
	if n.Op == And {
		hasAnd = true
	}
	if n.Op == Or {
		hasOr = true
	}
	for _, child := range n.Children {
		if child.Op == And {
			hasAnd = true
		}
		if child.Op == Or {
			hasOr = true
		}
		if len(child.Children) > 0 {
			hasChild = true
		}
	}
	if !hasChild || (hasAnd && hasOr) {
		// 没有子节点，或者同时出现了and和or，这都不能合并
		// 递归处理子节点
		for _, child := range n.Children {
			child.SimpleLogic()
		}
		return
	}

	// 需要进行合并
	op := n.Op
	if hasAnd {
		op = And
	} else if hasOr {
		op = Or
	}
	n.Op = op
	var dels []int       // 需要删除的节点
	var newNodes []*Node // 新增的节点
	for i, child := range n.Children {
		if child.Children == nil {
			continue
		}
		// 如果子节点的节点不为空，则删除子节点
		dels = append(dels, i)
		for _, sun := range child.Children {
			sun.father = n
			newNodes = append(newNodes, sun)
		}
	}
	// 删除多余的节点
	for i := len(dels) - 1; i >= 0; i-- {
		n.Children = RemoveNode(n.Children, dels[i])
	}
	// 新增节点
	for _, child := range newNodes {
		n.Children = append(n.Children, child)
		child.father = n
	}

	// 还需要继续处理该节点
	n.SimpleLogic()
	return
}
