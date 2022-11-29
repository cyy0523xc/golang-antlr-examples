// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Match

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// MatchListener is a complete listener for a parse tree produced by MatchParser.
type MatchListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterCmpExpr is called when entering the cmpExpr production.
	EnterCmpExpr(c *CmpExprContext)

	// EnterNearExpr is called when entering the nearExpr production.
	EnterNearExpr(c *NearExprContext)

	// EnterWordExpr is called when entering the wordExpr production.
	EnterWordExpr(c *WordExprContext)

	// EnterWords is called when entering the words production.
	EnterWords(c *WordsContext)

	// EnterWord is called when entering the word production.
	EnterWord(c *WordContext)

	// EnterDigits is called when entering the digits production.
	EnterDigits(c *DigitsContext)

	// EnterCmpOp is called when entering the cmpOp production.
	EnterCmpOp(c *CmpOpContext)

	// EnterNotOp is called when entering the notOp production.
	EnterNotOp(c *NotOpContext)

	// EnterAndOp is called when entering the andOp production.
	EnterAndOp(c *AndOpContext)

	// EnterOrOp is called when entering the orOp production.
	EnterOrOp(c *OrOpContext)

	// EnterNearOp is called when entering the nearOp production.
	EnterNearOp(c *NearOpContext)

	// EnterLeftOp is called when entering the leftOp production.
	EnterLeftOp(c *LeftOpContext)

	// EnterRightOp is called when entering the rightOp production.
	EnterRightOp(c *RightOpContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitCmpExpr is called when exiting the cmpExpr production.
	ExitCmpExpr(c *CmpExprContext)

	// ExitNearExpr is called when exiting the nearExpr production.
	ExitNearExpr(c *NearExprContext)

	// ExitWordExpr is called when exiting the wordExpr production.
	ExitWordExpr(c *WordExprContext)

	// ExitWords is called when exiting the words production.
	ExitWords(c *WordsContext)

	// ExitWord is called when exiting the word production.
	ExitWord(c *WordContext)

	// ExitDigits is called when exiting the digits production.
	ExitDigits(c *DigitsContext)

	// ExitCmpOp is called when exiting the cmpOp production.
	ExitCmpOp(c *CmpOpContext)

	// ExitNotOp is called when exiting the notOp production.
	ExitNotOp(c *NotOpContext)

	// ExitAndOp is called when exiting the andOp production.
	ExitAndOp(c *AndOpContext)

	// ExitOrOp is called when exiting the orOp production.
	ExitOrOp(c *OrOpContext)

	// ExitNearOp is called when exiting the nearOp production.
	ExitNearOp(c *NearOpContext)

	// ExitLeftOp is called when exiting the leftOp production.
	ExitLeftOp(c *LeftOpContext)

	// ExitRightOp is called when exiting the rightOp production.
	ExitRightOp(c *RightOpContext)
}
