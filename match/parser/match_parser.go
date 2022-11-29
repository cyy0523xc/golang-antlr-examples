// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Match

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type MatchParser struct {
	*antlr.BaseParser
}

var matchParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func matchParserInit() {
	staticData := &matchParserStaticData
	staticData.literalNames = []string{
		"", "'/'", "'<'", "'<='", "'>'", "'>='", "'='", "'not'", "'and'", "'or'",
		"'near'", "'('", "')'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "WORD", "DIGITS",
		"LINE_COMMENT", "COMMENT", "WS",
	}
	staticData.ruleNames = []string{
		"prog", "expr", "cmpExpr", "nearExpr", "wordExpr", "words", "word",
		"digits", "cmpOp", "notOp", "andOp", "orOp", "nearOp", "leftOp", "rightOp",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 17, 128, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 3, 1, 47, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 5, 1, 57, 8, 1, 10, 1, 12, 1, 60, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 3, 2, 70, 8, 2, 1, 3, 1, 3, 3, 3, 74, 8, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 3, 3, 81, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 87,
		8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 94, 8, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 108, 8, 5, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 0, 1, 2, 15, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 0, 1, 1, 0, 2, 6, 125, 0,
		30, 1, 0, 0, 0, 2, 46, 1, 0, 0, 0, 4, 69, 1, 0, 0, 0, 6, 86, 1, 0, 0, 0,
		8, 93, 1, 0, 0, 0, 10, 107, 1, 0, 0, 0, 12, 109, 1, 0, 0, 0, 14, 111, 1,
		0, 0, 0, 16, 113, 1, 0, 0, 0, 18, 115, 1, 0, 0, 0, 20, 117, 1, 0, 0, 0,
		22, 119, 1, 0, 0, 0, 24, 121, 1, 0, 0, 0, 26, 123, 1, 0, 0, 0, 28, 125,
		1, 0, 0, 0, 30, 31, 3, 2, 1, 0, 31, 32, 5, 0, 0, 1, 32, 1, 1, 0, 0, 0,
		33, 34, 6, 1, -1, 0, 34, 35, 3, 18, 9, 0, 35, 36, 3, 26, 13, 0, 36, 37,
		3, 2, 1, 0, 37, 38, 3, 28, 14, 0, 38, 47, 1, 0, 0, 0, 39, 40, 3, 26, 13,
		0, 40, 41, 3, 2, 1, 0, 41, 42, 3, 28, 14, 0, 42, 47, 1, 0, 0, 0, 43, 47,
		3, 4, 2, 0, 44, 47, 3, 6, 3, 0, 45, 47, 3, 8, 4, 0, 46, 33, 1, 0, 0, 0,
		46, 39, 1, 0, 0, 0, 46, 43, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 46, 45, 1,
		0, 0, 0, 47, 58, 1, 0, 0, 0, 48, 49, 10, 7, 0, 0, 49, 50, 3, 20, 10, 0,
		50, 51, 3, 2, 1, 8, 51, 57, 1, 0, 0, 0, 52, 53, 10, 6, 0, 0, 53, 54, 3,
		22, 11, 0, 54, 55, 3, 2, 1, 7, 55, 57, 1, 0, 0, 0, 56, 48, 1, 0, 0, 0,
		56, 52, 1, 0, 0, 0, 57, 60, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 59, 1,
		0, 0, 0, 59, 3, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 61, 62, 3, 12, 6, 0, 62,
		63, 3, 16, 8, 0, 63, 64, 3, 14, 7, 0, 64, 70, 1, 0, 0, 0, 65, 66, 3, 26,
		13, 0, 66, 67, 3, 4, 2, 0, 67, 68, 3, 28, 14, 0, 68, 70, 1, 0, 0, 0, 69,
		61, 1, 0, 0, 0, 69, 65, 1, 0, 0, 0, 70, 5, 1, 0, 0, 0, 71, 74, 3, 10, 5,
		0, 72, 74, 3, 12, 6, 0, 73, 71, 1, 0, 0, 0, 73, 72, 1, 0, 0, 0, 74, 75,
		1, 0, 0, 0, 75, 76, 3, 24, 12, 0, 76, 77, 5, 1, 0, 0, 77, 80, 3, 14, 7,
		0, 78, 81, 3, 10, 5, 0, 79, 81, 3, 12, 6, 0, 80, 78, 1, 0, 0, 0, 80, 79,
		1, 0, 0, 0, 81, 87, 1, 0, 0, 0, 82, 83, 3, 26, 13, 0, 83, 84, 3, 6, 3,
		0, 84, 85, 3, 28, 14, 0, 85, 87, 1, 0, 0, 0, 86, 73, 1, 0, 0, 0, 86, 82,
		1, 0, 0, 0, 87, 7, 1, 0, 0, 0, 88, 94, 3, 12, 6, 0, 89, 90, 3, 26, 13,
		0, 90, 91, 3, 8, 4, 0, 91, 92, 3, 28, 14, 0, 92, 94, 1, 0, 0, 0, 93, 88,
		1, 0, 0, 0, 93, 89, 1, 0, 0, 0, 94, 9, 1, 0, 0, 0, 95, 96, 3, 12, 6, 0,
		96, 97, 3, 20, 10, 0, 97, 98, 3, 12, 6, 0, 98, 108, 1, 0, 0, 0, 99, 100,
		3, 12, 6, 0, 100, 101, 3, 22, 11, 0, 101, 102, 3, 12, 6, 0, 102, 108, 1,
		0, 0, 0, 103, 104, 3, 26, 13, 0, 104, 105, 3, 10, 5, 0, 105, 106, 3, 28,
		14, 0, 106, 108, 1, 0, 0, 0, 107, 95, 1, 0, 0, 0, 107, 99, 1, 0, 0, 0,
		107, 103, 1, 0, 0, 0, 108, 11, 1, 0, 0, 0, 109, 110, 5, 13, 0, 0, 110,
		13, 1, 0, 0, 0, 111, 112, 5, 14, 0, 0, 112, 15, 1, 0, 0, 0, 113, 114, 7,
		0, 0, 0, 114, 17, 1, 0, 0, 0, 115, 116, 5, 7, 0, 0, 116, 19, 1, 0, 0, 0,
		117, 118, 5, 8, 0, 0, 118, 21, 1, 0, 0, 0, 119, 120, 5, 9, 0, 0, 120, 23,
		1, 0, 0, 0, 121, 122, 5, 10, 0, 0, 122, 25, 1, 0, 0, 0, 123, 124, 5, 11,
		0, 0, 124, 27, 1, 0, 0, 0, 125, 126, 5, 12, 0, 0, 126, 29, 1, 0, 0, 0,
		9, 46, 56, 58, 69, 73, 80, 86, 93, 107,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// MatchParserInit initializes any static state used to implement MatchParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMatchParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MatchParserInit() {
	staticData := &matchParserStaticData
	staticData.once.Do(matchParserInit)
}

// NewMatchParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMatchParser(input antlr.TokenStream) *MatchParser {
	MatchParserInit()
	this := new(MatchParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &matchParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// MatchParser tokens.
const (
	MatchParserEOF          = antlr.TokenEOF
	MatchParserT__0         = 1
	MatchParserT__1         = 2
	MatchParserT__2         = 3
	MatchParserT__3         = 4
	MatchParserT__4         = 5
	MatchParserT__5         = 6
	MatchParserT__6         = 7
	MatchParserT__7         = 8
	MatchParserT__8         = 9
	MatchParserT__9         = 10
	MatchParserT__10        = 11
	MatchParserT__11        = 12
	MatchParserWORD         = 13
	MatchParserDIGITS       = 14
	MatchParserLINE_COMMENT = 15
	MatchParserCOMMENT      = 16
	MatchParserWS           = 17
)

// MatchParser rules.
const (
	MatchParserRULE_prog     = 0
	MatchParserRULE_expr     = 1
	MatchParserRULE_cmpExpr  = 2
	MatchParserRULE_nearExpr = 3
	MatchParserRULE_wordExpr = 4
	MatchParserRULE_words    = 5
	MatchParserRULE_word     = 6
	MatchParserRULE_digits   = 7
	MatchParserRULE_cmpOp    = 8
	MatchParserRULE_notOp    = 9
	MatchParserRULE_andOp    = 10
	MatchParserRULE_orOp     = 11
	MatchParserRULE_nearOp   = 12
	MatchParserRULE_leftOp   = 13
	MatchParserRULE_rightOp  = 14
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MatchParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(MatchParserEOF, 0)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *MatchParser) Prog() (localctx IProgContext) {
	this := p
	_ = this

	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MatchParserRULE_prog)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)
		p.expr(0)
	}
	{
		p.SetState(31)
		p.Match(MatchParserEOF)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MatchParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) NotOp() INotOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INotOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INotOpContext)
}

func (s *ExprContext) LeftOp() ILeftOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeftOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeftOpContext)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) RightOp() IRightOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRightOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRightOpContext)
}

func (s *ExprContext) CmpExpr() ICmpExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmpExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmpExprContext)
}

func (s *ExprContext) NearExpr() INearExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INearExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INearExprContext)
}

func (s *ExprContext) WordExpr() IWordExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordExprContext)
}

func (s *ExprContext) AndOp() IAndOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndOpContext)
}

func (s *ExprContext) OrOp() IOrOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrOpContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *MatchParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *MatchParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, MatchParserRULE_expr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(34)
			p.NotOp()
		}
		{
			p.SetState(35)
			p.LeftOp()
		}
		{
			p.SetState(36)
			p.expr(0)
		}
		{
			p.SetState(37)
			p.RightOp()
		}

	case 2:
		{
			p.SetState(39)
			p.LeftOp()
		}
		{
			p.SetState(40)
			p.expr(0)
		}
		{
			p.SetState(41)
			p.RightOp()
		}

	case 3:
		{
			p.SetState(43)
			p.CmpExpr()
		}

	case 4:
		{
			p.SetState(44)
			p.NearExpr()
		}

	case 5:
		{
			p.SetState(45)
			p.WordExpr()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(56)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MatchParserRULE_expr)
				p.SetState(48)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(49)
					p.AndOp()
				}
				{
					p.SetState(50)
					p.expr(8)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MatchParserRULE_expr)
				p.SetState(52)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(53)
					p.OrOp()
				}
				{
					p.SetState(54)
					p.expr(7)
				}

			}

		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// ICmpExprContext is an interface to support dynamic dispatch.
type ICmpExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCmpExprContext differentiates from other interfaces.
	IsCmpExprContext()
}

type CmpExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmpExprContext() *CmpExprContext {
	var p = new(CmpExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MatchParserRULE_cmpExpr
	return p
}

func (*CmpExprContext) IsCmpExprContext() {}

func NewCmpExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmpExprContext {
	var p = new(CmpExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_cmpExpr

	return p
}

func (s *CmpExprContext) GetParser() antlr.Parser { return s.parser }

func (s *CmpExprContext) Word() IWordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordContext)
}

func (s *CmpExprContext) CmpOp() ICmpOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmpOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmpOpContext)
}

func (s *CmpExprContext) Digits() IDigitsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDigitsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDigitsContext)
}

func (s *CmpExprContext) LeftOp() ILeftOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeftOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeftOpContext)
}

func (s *CmpExprContext) CmpExpr() ICmpExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmpExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmpExprContext)
}

func (s *CmpExprContext) RightOp() IRightOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRightOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRightOpContext)
}

func (s *CmpExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmpExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmpExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterCmpExpr(s)
	}
}

func (s *CmpExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitCmpExpr(s)
	}
}

func (p *MatchParser) CmpExpr() (localctx ICmpExprContext) {
	this := p
	_ = this

	localctx = NewCmpExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MatchParserRULE_cmpExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(69)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MatchParserWORD:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)
			p.Word()
		}
		{
			p.SetState(62)
			p.CmpOp()
		}
		{
			p.SetState(63)
			p.Digits()
		}

	case MatchParserT__10:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.LeftOp()
		}
		{
			p.SetState(66)
			p.CmpExpr()
		}
		{
			p.SetState(67)
			p.RightOp()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INearExprContext is an interface to support dynamic dispatch.
type INearExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNearExprContext differentiates from other interfaces.
	IsNearExprContext()
}

type NearExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNearExprContext() *NearExprContext {
	var p = new(NearExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MatchParserRULE_nearExpr
	return p
}

func (*NearExprContext) IsNearExprContext() {}

func NewNearExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NearExprContext {
	var p = new(NearExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_nearExpr

	return p
}

func (s *NearExprContext) GetParser() antlr.Parser { return s.parser }

func (s *NearExprContext) NearOp() INearOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INearOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INearOpContext)
}

func (s *NearExprContext) Digits() IDigitsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDigitsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDigitsContext)
}

func (s *NearExprContext) AllWords() []IWordsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWordsContext); ok {
			len++
		}
	}

	tst := make([]IWordsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWordsContext); ok {
			tst[i] = t.(IWordsContext)
			i++
		}
	}

	return tst
}

func (s *NearExprContext) Words(i int) IWordsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordsContext)
}

func (s *NearExprContext) AllWord() []IWordContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWordContext); ok {
			len++
		}
	}

	tst := make([]IWordContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWordContext); ok {
			tst[i] = t.(IWordContext)
			i++
		}
	}

	return tst
}

func (s *NearExprContext) Word(i int) IWordContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordContext)
}

func (s *NearExprContext) LeftOp() ILeftOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeftOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeftOpContext)
}

func (s *NearExprContext) NearExpr() INearExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INearExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INearExprContext)
}

func (s *NearExprContext) RightOp() IRightOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRightOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRightOpContext)
}

func (s *NearExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NearExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NearExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterNearExpr(s)
	}
}

func (s *NearExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitNearExpr(s)
	}
}

func (p *MatchParser) NearExpr() (localctx INearExprContext) {
	this := p
	_ = this

	localctx = NewNearExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MatchParserRULE_nearExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(71)
				p.Words()
			}

		case 2:
			{
				p.SetState(72)
				p.Word()
			}

		}
		{
			p.SetState(75)
			p.NearOp()
		}
		{
			p.SetState(76)
			p.Match(MatchParserT__0)
		}
		{
			p.SetState(77)
			p.Digits()
		}
		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(78)
				p.Words()
			}

		case 2:
			{
				p.SetState(79)
				p.Word()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(82)
			p.LeftOp()
		}
		{
			p.SetState(83)
			p.NearExpr()
		}
		{
			p.SetState(84)
			p.RightOp()
		}

	}

	return localctx
}

// IWordExprContext is an interface to support dynamic dispatch.
type IWordExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWordExprContext differentiates from other interfaces.
	IsWordExprContext()
}

type WordExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWordExprContext() *WordExprContext {
	var p = new(WordExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MatchParserRULE_wordExpr
	return p
}

func (*WordExprContext) IsWordExprContext() {}

func NewWordExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WordExprContext {
	var p = new(WordExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_wordExpr

	return p
}

func (s *WordExprContext) GetParser() antlr.Parser { return s.parser }

func (s *WordExprContext) Word() IWordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordContext)
}

func (s *WordExprContext) LeftOp() ILeftOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeftOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeftOpContext)
}

func (s *WordExprContext) WordExpr() IWordExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordExprContext)
}

func (s *WordExprContext) RightOp() IRightOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRightOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRightOpContext)
}

func (s *WordExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WordExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WordExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterWordExpr(s)
	}
}

func (s *WordExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitWordExpr(s)
	}
}

func (p *MatchParser) WordExpr() (localctx IWordExprContext) {
	this := p
	_ = this

	localctx = NewWordExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MatchParserRULE_wordExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(93)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MatchParserWORD:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Word()
		}

	case MatchParserT__10:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.LeftOp()
		}
		{
			p.SetState(90)
			p.WordExpr()
		}
		{
			p.SetState(91)
			p.RightOp()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IWordsContext is an interface to support dynamic dispatch.
type IWordsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWordsContext differentiates from other interfaces.
	IsWordsContext()
}

type WordsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWordsContext() *WordsContext {
	var p = new(WordsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MatchParserRULE_words
	return p
}

func (*WordsContext) IsWordsContext() {}

func NewWordsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WordsContext {
	var p = new(WordsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_words

	return p
}

func (s *WordsContext) GetParser() antlr.Parser { return s.parser }

func (s *WordsContext) AllWord() []IWordContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWordContext); ok {
			len++
		}
	}

	tst := make([]IWordContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWordContext); ok {
			tst[i] = t.(IWordContext)
			i++
		}
	}

	return tst
}

func (s *WordsContext) Word(i int) IWordContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordContext)
}

func (s *WordsContext) AndOp() IAndOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndOpContext)
}

func (s *WordsContext) OrOp() IOrOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrOpContext)
}

func (s *WordsContext) LeftOp() ILeftOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeftOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeftOpContext)
}

func (s *WordsContext) Words() IWordsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordsContext)
}

func (s *WordsContext) RightOp() IRightOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRightOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRightOpContext)
}

func (s *WordsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WordsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WordsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterWords(s)
	}
}

func (s *WordsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitWords(s)
	}
}

func (p *MatchParser) Words() (localctx IWordsContext) {
	this := p
	_ = this

	localctx = NewWordsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MatchParserRULE_words)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(95)
			p.Word()
		}
		{
			p.SetState(96)
			p.AndOp()
		}
		{
			p.SetState(97)
			p.Word()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(99)
			p.Word()
		}
		{
			p.SetState(100)
			p.OrOp()
		}
		{
			p.SetState(101)
			p.Word()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(103)
			p.LeftOp()
		}
		{
			p.SetState(104)
			p.Words()
		}
		{
			p.SetState(105)
			p.RightOp()
		}

	}

	return localctx
}

// IWordContext is an interface to support dynamic dispatch.
type IWordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWordContext differentiates from other interfaces.
	IsWordContext()
}

type WordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWordContext() *WordContext {
	var p = new(WordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MatchParserRULE_word
	return p
}

func (*WordContext) IsWordContext() {}

func NewWordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WordContext {
	var p = new(WordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_word

	return p
}

func (s *WordContext) GetParser() antlr.Parser { return s.parser }

func (s *WordContext) WORD() antlr.TerminalNode {
	return s.GetToken(MatchParserWORD, 0)
}

func (s *WordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterWord(s)
	}
}

func (s *WordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitWord(s)
	}
}

func (p *MatchParser) Word() (localctx IWordContext) {
	this := p
	_ = this

	localctx = NewWordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MatchParserRULE_word)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(MatchParserWORD)
	}

	return localctx
}

// IDigitsContext is an interface to support dynamic dispatch.
type IDigitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDigitsContext differentiates from other interfaces.
	IsDigitsContext()
}

type DigitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDigitsContext() *DigitsContext {
	var p = new(DigitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MatchParserRULE_digits
	return p
}

func (*DigitsContext) IsDigitsContext() {}

func NewDigitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DigitsContext {
	var p = new(DigitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_digits

	return p
}

func (s *DigitsContext) GetParser() antlr.Parser { return s.parser }

func (s *DigitsContext) DIGITS() antlr.TerminalNode {
	return s.GetToken(MatchParserDIGITS, 0)
}

func (s *DigitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DigitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DigitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterDigits(s)
	}
}

func (s *DigitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitDigits(s)
	}
}

func (p *MatchParser) Digits() (localctx IDigitsContext) {
	this := p
	_ = this

	localctx = NewDigitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MatchParserRULE_digits)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(MatchParserDIGITS)
	}

	return localctx
}

// ICmpOpContext is an interface to support dynamic dispatch.
type ICmpOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCmpOpContext differentiates from other interfaces.
	IsCmpOpContext()
}

type CmpOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmpOpContext() *CmpOpContext {
	var p = new(CmpOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MatchParserRULE_cmpOp
	return p
}

func (*CmpOpContext) IsCmpOpContext() {}

func NewCmpOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmpOpContext {
	var p = new(CmpOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_cmpOp

	return p
}

func (s *CmpOpContext) GetParser() antlr.Parser { return s.parser }
func (s *CmpOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmpOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmpOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterCmpOp(s)
	}
}

func (s *CmpOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitCmpOp(s)
	}
}

func (p *MatchParser) CmpOp() (localctx ICmpOpContext) {
	this := p
	_ = this

	localctx = NewCmpOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MatchParserRULE_cmpOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&124) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INotOpContext is an interface to support dynamic dispatch.
type INotOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNotOpContext differentiates from other interfaces.
	IsNotOpContext()
}

type NotOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotOpContext() *NotOpContext {
	var p = new(NotOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MatchParserRULE_notOp
	return p
}

func (*NotOpContext) IsNotOpContext() {}

func NewNotOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotOpContext {
	var p = new(NotOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_notOp

	return p
}

func (s *NotOpContext) GetParser() antlr.Parser { return s.parser }
func (s *NotOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterNotOp(s)
	}
}

func (s *NotOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitNotOp(s)
	}
}

func (p *MatchParser) NotOp() (localctx INotOpContext) {
	this := p
	_ = this

	localctx = NewNotOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MatchParserRULE_notOp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(MatchParserT__6)
	}

	return localctx
}

// IAndOpContext is an interface to support dynamic dispatch.
type IAndOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAndOpContext differentiates from other interfaces.
	IsAndOpContext()
}

type AndOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndOpContext() *AndOpContext {
	var p = new(AndOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MatchParserRULE_andOp
	return p
}

func (*AndOpContext) IsAndOpContext() {}

func NewAndOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndOpContext {
	var p = new(AndOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_andOp

	return p
}

func (s *AndOpContext) GetParser() antlr.Parser { return s.parser }
func (s *AndOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterAndOp(s)
	}
}

func (s *AndOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitAndOp(s)
	}
}

func (p *MatchParser) AndOp() (localctx IAndOpContext) {
	this := p
	_ = this

	localctx = NewAndOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MatchParserRULE_andOp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(MatchParserT__7)
	}

	return localctx
}

// IOrOpContext is an interface to support dynamic dispatch.
type IOrOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrOpContext differentiates from other interfaces.
	IsOrOpContext()
}

type OrOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrOpContext() *OrOpContext {
	var p = new(OrOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MatchParserRULE_orOp
	return p
}

func (*OrOpContext) IsOrOpContext() {}

func NewOrOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrOpContext {
	var p = new(OrOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_orOp

	return p
}

func (s *OrOpContext) GetParser() antlr.Parser { return s.parser }
func (s *OrOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterOrOp(s)
	}
}

func (s *OrOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitOrOp(s)
	}
}

func (p *MatchParser) OrOp() (localctx IOrOpContext) {
	this := p
	_ = this

	localctx = NewOrOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MatchParserRULE_orOp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(MatchParserT__8)
	}

	return localctx
}

// INearOpContext is an interface to support dynamic dispatch.
type INearOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNearOpContext differentiates from other interfaces.
	IsNearOpContext()
}

type NearOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNearOpContext() *NearOpContext {
	var p = new(NearOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MatchParserRULE_nearOp
	return p
}

func (*NearOpContext) IsNearOpContext() {}

func NewNearOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NearOpContext {
	var p = new(NearOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_nearOp

	return p
}

func (s *NearOpContext) GetParser() antlr.Parser { return s.parser }
func (s *NearOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NearOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NearOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterNearOp(s)
	}
}

func (s *NearOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitNearOp(s)
	}
}

func (p *MatchParser) NearOp() (localctx INearOpContext) {
	this := p
	_ = this

	localctx = NewNearOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MatchParserRULE_nearOp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(MatchParserT__9)
	}

	return localctx
}

// ILeftOpContext is an interface to support dynamic dispatch.
type ILeftOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLeftOpContext differentiates from other interfaces.
	IsLeftOpContext()
}

type LeftOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLeftOpContext() *LeftOpContext {
	var p = new(LeftOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MatchParserRULE_leftOp
	return p
}

func (*LeftOpContext) IsLeftOpContext() {}

func NewLeftOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeftOpContext {
	var p = new(LeftOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_leftOp

	return p
}

func (s *LeftOpContext) GetParser() antlr.Parser { return s.parser }
func (s *LeftOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeftOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LeftOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterLeftOp(s)
	}
}

func (s *LeftOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitLeftOp(s)
	}
}

func (p *MatchParser) LeftOp() (localctx ILeftOpContext) {
	this := p
	_ = this

	localctx = NewLeftOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MatchParserRULE_leftOp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Match(MatchParserT__10)
	}

	return localctx
}

// IRightOpContext is an interface to support dynamic dispatch.
type IRightOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRightOpContext differentiates from other interfaces.
	IsRightOpContext()
}

type RightOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRightOpContext() *RightOpContext {
	var p = new(RightOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MatchParserRULE_rightOp
	return p
}

func (*RightOpContext) IsRightOpContext() {}

func NewRightOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RightOpContext {
	var p = new(RightOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_rightOp

	return p
}

func (s *RightOpContext) GetParser() antlr.Parser { return s.parser }
func (s *RightOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RightOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RightOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterRightOp(s)
	}
}

func (s *RightOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitRightOp(s)
	}
}

func (p *MatchParser) RightOp() (localctx IRightOpContext) {
	this := p
	_ = this

	localctx = NewRightOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MatchParserRULE_rightOp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Match(MatchParserT__11)
	}

	return localctx
}

func (p *MatchParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MatchParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
