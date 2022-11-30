package main

import "fmt"

// 化简表达式
func (n *Node) Simple() {
	if n.Children == nil {
		return
	}

	if len(n.Children) == 1 {
		// 只有一个子节点: 可以直接将孙节点直接连接到该节点上
		child := n.Children[0]
		isMerge := false
		if n.Op == "" {
			// 父节点为空，保留子节点信息
			n.rule = child.rule
			n.Op = child.Op
			isMerge = true
		} else if child.Op == "" {
			// 子节点为空, 保留父节点的信息
			isMerge = true
		}

		if isMerge {
			n.Cmp = child.Cmp
			n.Near = child.Near
			n.Children = child.Children
			for _, sun := range child.Children {
				sun.father = n
			}
			if debug {
				fmt.Println("=====> Simple op: " + n.Op + " child op: " + child.Op)
			}
			child = nil
			n.Simple()
		}
	}

	// 递归处理子节点
	for _, child := range n.Children {
		child.Simple()
	}
	return
}

// 化简嵌套的逻辑表达式
func (n *Node) SimpleLogic() {
	// a and (b and c) => and(a,b,c)
	// a or (b or c) => or(a,b,c)
	if n.Children == nil {
		return
	}
	if !(n.Op == And || n.Op == Or || n.Op == "") {
		return
	}
	if n.Op == Not {
		return
	}
	hasAnd, hasOr, hasNot, hasChild := false, false, false, false
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
		if child.Op == Not {
			hasNot = true
		}
		if len(child.Children) > 0 {
			hasChild = true
		}
	}
	if !hasChild || (hasAnd && hasOr) || hasNot {
		// 没有子节点，或者同时出现了and和or，这都不能合并
		// 递归处理子节点
		for _, child := range n.Children {
			child.SimpleLogic()
		}
		return
	}

	// 需要进行合并
	if debug {
		fmt.Println("======> simple logic")
	}
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
