module github.com/cyy0523xc/golang-antlr-examples/match

go 1.18

require (
	github.com/antlr/antlr4/runtime/Go/antlr/v4 v4.0.0-20221120182715-47415e33c366
	golang.org/x/exp v0.0.0-20220722155223-a9213eeb770e // indirect
)

require parser v0.0.0

replace parser => ./parser
