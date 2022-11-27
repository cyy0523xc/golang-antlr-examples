package main

import (
	"testing"
)

var cases = [...][2]string{
	{
		"'keyword'",
	},
	{
		"'keyword' and 'keyword2' >= 3",
		"{'op':'and','cmp':[{'word':'keyword','op':'>=','val':1},{'word'}]}",
	},
}

func TestParseExpr(t *testing.T) {

}
