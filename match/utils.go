package main

// 移除某个节点
func RemoveNode(nodes []*Node, idx int) []*Node {
	if idx == 0 {
		return nodes[1:]
	}
	if idx == len(nodes)-1 {
		return nodes[:len(nodes)-1]
	}
	return append(nodes[:idx], nodes[idx+1:]...)
}

// 创建Near节点
func CreateNearNode(left, right string, father *Node, dist int) (node *Node) {
	return &Node{
		father: father,
		Op:     Near,
		Near: &NearExpr{
			LeftWords:  []string{left},
			RightWords: []string{right},
			Dist:       dist,
		},
	}
}
