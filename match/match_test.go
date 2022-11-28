package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Case struct {
	In    string
	Out   Node
	Title string
}

var cases = [...]Case{
	{
		Title: "单个关键词",
		In:    "'keyword'",
		Out: Node{
			Op: "cmp",
			Cmp: &CmpExpr{
				Word: "keyword",
				Op:   ">=",
				Num:  1,
			},
		},
	},
	{
		Title: "两个关键词组合",
		In:    "'keyword' and 'keyword2' >= 3",
		Out: Node{
			Op: "and",
			Children: []*Node{
				{
					Op: "cmp",
					Cmp: &CmpExpr{
						Word: "keyword",
						Op:   ">=",
						Num:  1,
					},
				},
				{
					Op: "cmp",
					Cmp: &CmpExpr{
						Word: "keyword2",
						Op:   ">=",
						Num:  3,
					},
				},
			},
		},
	},
	{
		Title: "比较表达式与near表达式组合",
		In:    "'keyword' and ('keyword2' near/10 'keyword3')",
		Out: Node{
			Op: "and",
			Children: []*Node{
				{
					Op: "cmp",
					Cmp: &CmpExpr{
						Word: "keyword",
						Op:   ">=",
						Num:  1,
					},
				},
				{
					Op: "near",
					Near: &NearExpr{
						LeftWords:  []string{"keyword2"},
						RightWords: []string{"keyword3"},
						Dist:       10,
					},
				},
			},
		},
	},
}

func TestParseExpr(t *testing.T) {
	// debug = false
	for i, c := range cases {
		bytes, err := json.Marshal(c.Out)
		if err != nil {
			t.Errorf(err.Error())
		}
		json1 := string(bytes)

		// 调用解释
		json2, err := ParseExpr(c.In)
		if err != nil {
			t.Errorf(err.Error())
		}
		if json1 != json2 {
			t.Fatalf("%02d: %s: not pass!\n%s \n%s\n", i, c.Title, json1, json2)
		} else {
			fmt.Printf("\n------------------\n%02d: %s: pass! \n", i, c.Title)
		}
		fmt.Println()
	}
}
