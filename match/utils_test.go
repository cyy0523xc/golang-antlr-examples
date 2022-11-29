package main

import (
	"fmt"
	"testing"
)

func TestFmtExpr(t *testing.T) {
	var expr = "'xxx' and ('a' and 'b' and 'c') near/10 ('e' or 'f')"
	newExpr, err := FmtNearStr(expr)
	fmt.Println(newExpr, err)
}
