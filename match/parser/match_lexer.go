// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type MatchLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var matchlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func matchlexerLexerInit() {
	staticData := &matchlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'/'", "'<'", "'<='", "'>'", "'>='", "'='", "'not'", "'and'", "'or'",
		"'near'", "'('", "')'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "WORD", "DIGITS",
		"LINE_COMMENT", "COMMENT", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "WORD", "DIGITS", "LINE_COMMENT", "COMMENT",
		"WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 17, 127, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 12, 1, 12, 4, 12, 72, 8, 12, 11, 12, 12, 12, 73, 1, 12,
		1, 12, 1, 12, 4, 12, 79, 8, 12, 11, 12, 12, 12, 80, 1, 12, 3, 12, 84, 8,
		12, 1, 13, 4, 13, 87, 8, 13, 11, 13, 12, 13, 88, 1, 14, 1, 14, 1, 14, 1,
		14, 5, 14, 95, 8, 14, 10, 14, 12, 14, 98, 9, 14, 1, 14, 3, 14, 101, 8,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 111,
		8, 15, 10, 15, 12, 15, 114, 9, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		16, 4, 16, 122, 8, 16, 11, 16, 12, 16, 123, 1, 16, 1, 16, 4, 73, 80, 96,
		112, 0, 17, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 1, 0, 3, 6,
		0, 46, 46, 48, 57, 65, 90, 92, 92, 97, 122, 19968, 40869, 1, 0, 48, 57,
		3, 0, 9, 10, 13, 13, 32, 32, 134, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0,
		0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0,
		0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 1, 35, 1, 0,
		0, 0, 3, 37, 1, 0, 0, 0, 5, 39, 1, 0, 0, 0, 7, 42, 1, 0, 0, 0, 9, 44, 1,
		0, 0, 0, 11, 47, 1, 0, 0, 0, 13, 49, 1, 0, 0, 0, 15, 53, 1, 0, 0, 0, 17,
		57, 1, 0, 0, 0, 19, 60, 1, 0, 0, 0, 21, 65, 1, 0, 0, 0, 23, 67, 1, 0, 0,
		0, 25, 83, 1, 0, 0, 0, 27, 86, 1, 0, 0, 0, 29, 90, 1, 0, 0, 0, 31, 106,
		1, 0, 0, 0, 33, 121, 1, 0, 0, 0, 35, 36, 5, 47, 0, 0, 36, 2, 1, 0, 0, 0,
		37, 38, 5, 60, 0, 0, 38, 4, 1, 0, 0, 0, 39, 40, 5, 60, 0, 0, 40, 41, 5,
		61, 0, 0, 41, 6, 1, 0, 0, 0, 42, 43, 5, 62, 0, 0, 43, 8, 1, 0, 0, 0, 44,
		45, 5, 62, 0, 0, 45, 46, 5, 61, 0, 0, 46, 10, 1, 0, 0, 0, 47, 48, 5, 61,
		0, 0, 48, 12, 1, 0, 0, 0, 49, 50, 5, 110, 0, 0, 50, 51, 5, 111, 0, 0, 51,
		52, 5, 116, 0, 0, 52, 14, 1, 0, 0, 0, 53, 54, 5, 97, 0, 0, 54, 55, 5, 110,
		0, 0, 55, 56, 5, 100, 0, 0, 56, 16, 1, 0, 0, 0, 57, 58, 5, 111, 0, 0, 58,
		59, 5, 114, 0, 0, 59, 18, 1, 0, 0, 0, 60, 61, 5, 110, 0, 0, 61, 62, 5,
		101, 0, 0, 62, 63, 5, 97, 0, 0, 63, 64, 5, 114, 0, 0, 64, 20, 1, 0, 0,
		0, 65, 66, 5, 40, 0, 0, 66, 22, 1, 0, 0, 0, 67, 68, 5, 41, 0, 0, 68, 24,
		1, 0, 0, 0, 69, 71, 5, 34, 0, 0, 70, 72, 7, 0, 0, 0, 71, 70, 1, 0, 0, 0,
		72, 73, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 74, 75, 1,
		0, 0, 0, 75, 84, 5, 34, 0, 0, 76, 78, 5, 39, 0, 0, 77, 79, 7, 0, 0, 0,
		78, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 80, 78, 1,
		0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 84, 5, 39, 0, 0, 83, 69, 1, 0, 0, 0, 83,
		76, 1, 0, 0, 0, 84, 26, 1, 0, 0, 0, 85, 87, 7, 1, 0, 0, 86, 85, 1, 0, 0,
		0, 87, 88, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 28,
		1, 0, 0, 0, 90, 91, 5, 47, 0, 0, 91, 92, 5, 47, 0, 0, 92, 96, 1, 0, 0,
		0, 93, 95, 9, 0, 0, 0, 94, 93, 1, 0, 0, 0, 95, 98, 1, 0, 0, 0, 96, 97,
		1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0,
		99, 101, 5, 13, 0, 0, 100, 99, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 102,
		1, 0, 0, 0, 102, 103, 5, 10, 0, 0, 103, 104, 1, 0, 0, 0, 104, 105, 6, 14,
		0, 0, 105, 30, 1, 0, 0, 0, 106, 107, 5, 47, 0, 0, 107, 108, 5, 42, 0, 0,
		108, 112, 1, 0, 0, 0, 109, 111, 9, 0, 0, 0, 110, 109, 1, 0, 0, 0, 111,
		114, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 113, 115,
		1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 115, 116, 5, 42, 0, 0, 116, 117, 5, 47,
		0, 0, 117, 118, 1, 0, 0, 0, 118, 119, 6, 15, 0, 0, 119, 32, 1, 0, 0, 0,
		120, 122, 7, 2, 0, 0, 121, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123,
		121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 126,
		6, 16, 0, 0, 126, 34, 1, 0, 0, 0, 9, 0, 73, 80, 83, 88, 96, 100, 112, 123,
		1, 6, 0, 0,
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

// MatchLexerInit initializes any static state used to implement MatchLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewMatchLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func MatchLexerInit() {
	staticData := &matchlexerLexerStaticData
	staticData.once.Do(matchlexerLexerInit)
}

// NewMatchLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewMatchLexer(input antlr.CharStream) *MatchLexer {
	MatchLexerInit()
	l := new(MatchLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &matchlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Match.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MatchLexer tokens.
const (
	MatchLexerT__0         = 1
	MatchLexerT__1         = 2
	MatchLexerT__2         = 3
	MatchLexerT__3         = 4
	MatchLexerT__4         = 5
	MatchLexerT__5         = 6
	MatchLexerT__6         = 7
	MatchLexerT__7         = 8
	MatchLexerT__8         = 9
	MatchLexerT__9         = 10
	MatchLexerT__10        = 11
	MatchLexerT__11        = 12
	MatchLexerWORD         = 13
	MatchLexerDIGITS       = 14
	MatchLexerLINE_COMMENT = 15
	MatchLexerCOMMENT      = 16
	MatchLexerWS           = 17
)
