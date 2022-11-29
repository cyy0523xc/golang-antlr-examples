package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
)

const (
	// Near前后多个关键词的分隔符
	NearWordSplit = "__::__"
	// Near表达式正则
	NearExprStr  = "\\s+near\\s*\\/\\s*[0-9]+\\s+"
	NearDistStr  = "\\s+near\\s*\\/\\s*([0-9]+)\\s+"
	NearWordsStr = "\\(.*?\\)"
	NearBothStr  = NearWordsStr + NearExprStr + NearWordsStr
	NearLeftStr  = NearWordsStr + NearExprStr
	NearRightStr = NearExprStr + NearWordsStr
)

var (
	// 编译后的正则
	NearBothRe, _  = regexp.Compile(NearBothStr)
	NearLeftRe, _  = regexp.Compile(NearLeftStr)
	NearRightRe, _ = regexp.Compile(NearRightStr)
	WordsRe, _     = regexp.Compile("\\((.*?)\\)")
	NearDistRe, _  = regexp.Compile(NearDistStr)
)

// 格式Near
// 如: (a and b and c) near/10 (e or f)
// ==>  'and:a__::__b__::__c' near/10 'or:e__::__f'
// go test -run TestFmtExpr -v .
func FmtNearStr(expr string) (newExpr string, err error) {
	// fmt.Println(NearBothRe.FindAllString(expr, -1))
	exprBytes := []byte(expr)
	indexes := NearBothRe.FindAllIndex(exprBytes, -1)
	for _, item := range indexes {
		str := string(exprBytes[item[0]:item[1]])
		fmt.Println("==>", str)
		wordsMatch := WordsRe.FindAllString(str, -1)
		fmt.Println(len(wordsMatch), wordsMatch)
	}
	fmt.Println(indexes)
	return
}

// 格式表达式
// 1. 关键词前后如果没有引号，则加上双引号
// 2. 关键词前后如果是单引号，则改为双引号
// 3. 格式化near表达式
// 异常情况：
// 1. 括号不配对
// 2. 引号不配对
func FmtExpr(expr string) (newExpr string, err error) {
	return
}

// 移除某个子节点
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

// 格式化json字符串
func FmtJson(bs []byte) (out string) {
	var buf bytes.Buffer
	json.Indent(&buf, bs, "", "  ")
	out = buf.String()
	return
}
