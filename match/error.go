package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

const (
	// 异常编码
	ErrSyntaxCode = 1
	ErrJsonCode   = 2
)

var errMsg = map[int]string{
	ErrSyntaxCode: "表达式解释失败",
	ErrJsonCode:   "生成json字符串异常",
}

type ErrorListener struct {
	root *Node
}

func CreateJsonError(detail string) *Node {
	return &Node{
		Err: &Error{
			Code:   ErrJsonCode,
			Msg:    errMsg[ErrJsonCode],
			Detail: detail,
		},
	}
}

func (d *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	// fmt.Printf("ERROR: %#v\n", recognizer.(*antlr.BaseParser).GetExpectedTokens().String())
	// fmt.Println(column)
	// fmt.Printf("ERROR: %#v, %s\n", e.GetMessage(), msg)
	d.root.Err = &Error{
		Code:   ErrSyntaxCode,
		Column: column,
		Msg:    errMsg[ErrSyntaxCode],
		Detail: msg,
	}
}

func (d *ErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (d *ErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (d *ErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
}
