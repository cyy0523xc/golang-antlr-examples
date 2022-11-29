// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Match

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseMatchListener is a complete listener for a parse tree produced by MatchParser.
type BaseMatchListener struct{}

var _ MatchListener = &BaseMatchListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMatchListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMatchListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMatchListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMatchListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseMatchListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseMatchListener) ExitProg(ctx *ProgContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseMatchListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseMatchListener) ExitExpr(ctx *ExprContext) {}

// EnterCmpExpr is called when production cmpExpr is entered.
func (s *BaseMatchListener) EnterCmpExpr(ctx *CmpExprContext) {}

// ExitCmpExpr is called when production cmpExpr is exited.
func (s *BaseMatchListener) ExitCmpExpr(ctx *CmpExprContext) {}

// EnterNearExpr is called when production nearExpr is entered.
func (s *BaseMatchListener) EnterNearExpr(ctx *NearExprContext) {}

// ExitNearExpr is called when production nearExpr is exited.
func (s *BaseMatchListener) ExitNearExpr(ctx *NearExprContext) {}

// EnterWordExpr is called when production wordExpr is entered.
func (s *BaseMatchListener) EnterWordExpr(ctx *WordExprContext) {}

// ExitWordExpr is called when production wordExpr is exited.
func (s *BaseMatchListener) ExitWordExpr(ctx *WordExprContext) {}

// EnterWords is called when production words is entered.
func (s *BaseMatchListener) EnterWords(ctx *WordsContext) {}

// ExitWords is called when production words is exited.
func (s *BaseMatchListener) ExitWords(ctx *WordsContext) {}

// EnterWord is called when production word is entered.
func (s *BaseMatchListener) EnterWord(ctx *WordContext) {}

// ExitWord is called when production word is exited.
func (s *BaseMatchListener) ExitWord(ctx *WordContext) {}

// EnterDigits is called when production digits is entered.
func (s *BaseMatchListener) EnterDigits(ctx *DigitsContext) {}

// ExitDigits is called when production digits is exited.
func (s *BaseMatchListener) ExitDigits(ctx *DigitsContext) {}

// EnterCmpOp is called when production cmpOp is entered.
func (s *BaseMatchListener) EnterCmpOp(ctx *CmpOpContext) {}

// ExitCmpOp is called when production cmpOp is exited.
func (s *BaseMatchListener) ExitCmpOp(ctx *CmpOpContext) {}

// EnterNotOp is called when production notOp is entered.
func (s *BaseMatchListener) EnterNotOp(ctx *NotOpContext) {}

// ExitNotOp is called when production notOp is exited.
func (s *BaseMatchListener) ExitNotOp(ctx *NotOpContext) {}

// EnterAndOp is called when production andOp is entered.
func (s *BaseMatchListener) EnterAndOp(ctx *AndOpContext) {}

// ExitAndOp is called when production andOp is exited.
func (s *BaseMatchListener) ExitAndOp(ctx *AndOpContext) {}

// EnterOrOp is called when production orOp is entered.
func (s *BaseMatchListener) EnterOrOp(ctx *OrOpContext) {}

// ExitOrOp is called when production orOp is exited.
func (s *BaseMatchListener) ExitOrOp(ctx *OrOpContext) {}

// EnterNearOp is called when production nearOp is entered.
func (s *BaseMatchListener) EnterNearOp(ctx *NearOpContext) {}

// ExitNearOp is called when production nearOp is exited.
func (s *BaseMatchListener) ExitNearOp(ctx *NearOpContext) {}

// EnterLeftOp is called when production leftOp is entered.
func (s *BaseMatchListener) EnterLeftOp(ctx *LeftOpContext) {}

// ExitLeftOp is called when production leftOp is exited.
func (s *BaseMatchListener) ExitLeftOp(ctx *LeftOpContext) {}

// EnterRightOp is called when production rightOp is entered.
func (s *BaseMatchListener) EnterRightOp(ctx *RightOpContext) {}

// ExitRightOp is called when production rightOp is exited.
func (s *BaseMatchListener) ExitRightOp(ctx *RightOpContext) {}
