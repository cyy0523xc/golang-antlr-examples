package main

import "fmt"

func (n *Node) OptimNear() {
	if n.Op != Near {
		for _, child := range n.Children {
			child.OptimNear()
		}
		return
	}

	if len(n.Near.LeftWords) == 1 && len(n.Near.RightWords) == 1 {
		return
	}
	if debug {
		fmt.Println("------------ Begin parse near ---------------")
	}
	father := n.father // 父节点
	// 从原父节点中移除该节点
	child_idx := 0
	for i, child := range father.Children {
		if child == n {
			child_idx = i
		}
	}
	if child_idx == 0 {
		father.Children = father.Children[1:]
	} else if child_idx == len(father.Children)-1 {
		father.Children = father.Children[:len(father.Children)-1]
	} else {
		father.Children = append(father.Children[:child_idx], father.Children[child_idx+1:]...)
	}

	// 加入一个节点
	curr := &Node{
		father:   father,
		Children: []*Node{},
	}
	father.Children = append(father.Children, curr)

	if len(n.Near.LeftWords) == 1 {
		// 左侧只有一个词
		curr.Op = n.Near.rightOp
		for _, w := range n.Near.RightWords {
			child := CreateNearNode(n.Near.LeftWords[0], w, curr, n.Near.Dist)
			curr.Children = append(curr.Children, child)
		}
		return
	}
	if len(n.Near.RightWords) == 1 {
		// 右侧只有一个词
		curr.Op = n.Near.leftOp
		for _, w := range n.Near.LeftWords {
			child := CreateNearNode(w, n.Near.RightWords[0], curr, n.Near.Dist)
			curr.Children = append(curr.Children, child)
		}
		return
	}
	if n.Near.leftOp == n.Near.rightOp {
		// 加入一个节点
		// 左右两侧的逻辑操作符一样
		curr.Op = n.Near.leftOp
		for _, wl := range n.Near.LeftWords {
			for _, wr := range n.Near.RightWords {
				child := CreateNearNode(wl, wr, curr, n.Near.Dist)
				curr.Children = append(curr.Children, child)
			}
		}
		return
	}
	if n.Near.leftOp == And && n.Near.rightOp == Or {
		fmt.Println("左侧是and，右侧是or")
		curr.Op = Or
		for _, wr := range n.Near.RightWords {
			andNode := &Node{
				father:   curr,
				Op:       And,
				Children: []*Node{},
			}
			curr.Children = append(curr.Children, andNode)
			for _, wl := range n.Near.LeftWords {
				child := CreateNearNode(wl, wr, andNode, n.Near.Dist)
				andNode.Children = append(andNode.Children, child)
			}
		}
		return
	}
	if n.Near.leftOp == Or && n.Near.rightOp == And {
		fmt.Println("左侧是or，右侧是and")
		curr.Op = Or
		for _, wl := range n.Near.LeftWords {
			andNode := &Node{
				father:   curr,
				Op:       And,
				Children: []*Node{},
			}
			curr.Children = append(curr.Children, andNode)
			for _, wr := range n.Near.RightWords {
				child := CreateNearNode(wl, wr, andNode, n.Near.Dist)
				andNode.Children = append(andNode.Children, child)
			}
		}
		return
	}

	return
}
