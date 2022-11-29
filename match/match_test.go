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
		Title: "两个关键词and组合",
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
		Title: "3个关键词and组合",
		In:    "'keyword' and 'keyword2' >= 3 and \"keyword3\"  < 2",
		Out: Node{
			Op: "and",
			Children: []*Node{
				{
					Op: "cmp",
					Cmp: &CmpExpr{
						Word: "keyword3",
						Op:   "<",
						Num:  2,
					},
				},
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
		Title: "4个关键词and组合",
		In:    "'keyword' and 'keyword2' >= 3 and \"keyword3\" < 2 and 'keyword4'=3",
		Out: Node{
			Op: "and",
			Children: []*Node{
				{
					Op: "cmp",
					Cmp: &CmpExpr{
						Word: "keyword4",
						Op:   "=",
						Num:  3,
					},
				},
				{
					Op: "cmp",
					Cmp: &CmpExpr{
						Word: "keyword3",
						Op:   "<",
						Num:  2,
					},
				},
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
		Title: "关键词and与or组合",
		In:    "'keyword5' or 'keyword1' >= 3 or \"keyword2\" < 2 and 'keyword3'=3 and 'keyword4'>4",
		Out: Node{
			Op: "or",
			Children: []*Node{
				{
					Op: "or",
					Children: []*Node{
						{
							Op: "cmp",
							Cmp: &CmpExpr{
								Word: "keyword5",
								Op:   ">=",
								Num:  1,
							},
						},
						{
							Op: "cmp",
							Cmp: &CmpExpr{
								Word: "keyword1",
								Op:   ">=",
								Num:  3,
							},
						},
					},
				},
				{
					Op: "and",
					Children: []*Node{
						{
							Op: "cmp",
							Cmp: &CmpExpr{
								Word: "keyword4",
								Op:   ">",
								Num:  4,
							},
						},
						{
							Op: "cmp",
							Cmp: &CmpExpr{
								Word: "keyword2",
								Op:   "<",
								Num:  2,
							},
						},
						{
							Op: "cmp",
							Cmp: &CmpExpr{
								Word: "keyword3",
								Op:   "=",
								Num:  3,
							},
						},
					},
				},
			},
		},
	},
	{
		Title: "关键词and与or组合（带括号）",
		In:    "'keyword5' or ('keyword1' >= 3 or \"keyword2\" < 2) and 'keyword3'=3 and 'keyword4'>4",
		Out: Node{
			Op: "or",
			Children: []*Node{
				{
					Op: "cmp",
					Cmp: &CmpExpr{
						Word: "keyword5",
						Op:   ">=",
						Num:  1,
					},
				},
				{
					Op: "and",
					Children: []*Node{
						{
							Op: "cmp",
							Cmp: &CmpExpr{
								Word: "keyword4",
								Op:   ">",
								Num:  4,
							},
						},
						{
							Op: "or",
							Children: []*Node{
								{
									Op: "cmp",
									Cmp: &CmpExpr{
										Word: "keyword1",
										Op:   ">=",
										Num:  3,
									},
								},
								{
									Op: "cmp",
									Cmp: &CmpExpr{
										Word: "keyword2",
										Op:   "<",
										Num:  2,
									},
								},
							},
						},
						{
							Op: "cmp",
							Cmp: &CmpExpr{
								Word: "keyword3",
								Op:   "=",
								Num:  3,
							},
						},
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
	{
		Title: "比较表达式与多组near表达式组合",
		In:    "'keyword' and ('keyword2' near/10 'keyword3') or ('中国' near/20 '1000')",
		Out: Node{
			Op: "or",
			Children: []*Node{
				{
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
				{
					Op: "near",
					Near: &NearExpr{
						LeftWords:  []string{"中国"},
						RightWords: []string{"1000"},
						Dist:       20,
					},
				},
			},
		},
	},
	{
		Title: "near表达式嵌套",
		In:    "('中国' and '美国') near/20 '1000'",
		Out: Node{
			Op: "and",
			Children: []*Node{
				{
					Op: "near",
					Near: &NearExpr{
						LeftWords:  []string{"中国"},
						RightWords: []string{"1000"},
						Dist:       20,
					},
				},
				{
					Op: "near",
					Near: &NearExpr{
						LeftWords:  []string{"美国"},
						RightWords: []string{"1000"},
						Dist:       20,
					},
				},
			},
		},
	},
	{
		Title: "near表达式嵌套and/or",
		In:    "('中国' and '美国') near/20 ('1000' or '08广州')",
		Out: Node{
			Op: "or",
			Children: []*Node{
				{
					Op: "and",
					Children: []*Node{
						{
							Op: "near",
							Near: &NearExpr{
								LeftWords:  []string{"中国"},
								RightWords: []string{"1000"},
								Dist:       20,
							},
						},
						{
							Op: "near",
							Near: &NearExpr{
								LeftWords:  []string{"美国"},
								RightWords: []string{"1000"},
								Dist:       20,
							},
						},
					},
				},
				{
					Op: "and",
					Children: []*Node{
						{
							Op: "near",
							Near: &NearExpr{
								LeftWords:  []string{"中国"},
								RightWords: []string{"08广州"},
								Dist:       20,
							},
						},
						{
							Op: "near",
							Near: &NearExpr{
								LeftWords:  []string{"美国"},
								RightWords: []string{"08广州"},
								Dist:       20,
							},
						},
					},
				},
			},
		},
	},
	{
		Title: "复杂near表达式嵌套",
		In:    "'keyword' and ('keyword2' near/10 'keyword3') or (('中国' and '美国') near/20 '1000')",
		Out: Node{
			Op: "or",
			Children: []*Node{
				{
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
				{
					Op: "and",
					Children: []*Node{
						{
							Op: "near",
							Near: &NearExpr{
								LeftWords:  []string{"中国"},
								RightWords: []string{"1000"},
								Dist:       20,
							},
						},
						{
							Op: "near",
							Near: &NearExpr{
								LeftWords:  []string{"美国"},
								RightWords: []string{"1000"},
								Dist:       20,
							},
						},
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
			fmt.Println(FmtJson([]byte(json2)))
			t.Fatalf("%02d: %s: not pass!\n%s \n%s\n", i, c.Title, json1, json2)
		} else {
			fmt.Printf("\n------------------\n%02d: %s: pass! \n", i, c.Title)
		}
		fmt.Println()
	}
}
