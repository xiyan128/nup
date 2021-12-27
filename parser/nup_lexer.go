// Code generated from parser/NupLexer.g4 by ANTLR 4.9. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 17, 133,
	8, 1, 8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6,
	4, 7, 9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12,
	9, 12, 4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9,
	17, 4, 18, 9, 18, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 7, 4, 46, 10,
	4, 12, 4, 14, 4, 49, 11, 4, 3, 4, 3, 4, 3, 4, 5, 4, 54, 10, 4, 3, 4, 7,
	4, 57, 10, 4, 12, 4, 14, 4, 60, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 6,
	5, 67, 10, 5, 13, 5, 14, 5, 68, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3,
	9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 6, 12, 84, 10, 12, 13, 12,
	14, 12, 85, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 7,
	14, 96, 10, 14, 12, 14, 14, 14, 99, 11, 14, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 16, 6, 16, 106, 10, 16, 13, 16, 14, 16, 107, 3, 16, 3, 16, 6, 16, 112,
	10, 16, 13, 16, 14, 16, 113, 5, 16, 116, 10, 16, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 127, 10, 17, 3, 18, 6, 18,
	130, 10, 18, 13, 18, 14, 18, 131, 2, 2, 19, 4, 3, 6, 2, 8, 4, 10, 5, 12,
	6, 14, 7, 16, 8, 18, 9, 20, 10, 22, 11, 24, 12, 26, 13, 28, 14, 30, 2,
	32, 15, 34, 16, 36, 17, 4, 2, 3, 8, 4, 2, 11, 11, 34, 34, 4, 2, 12, 12,
	15, 15, 6, 2, 12, 12, 15, 15, 94, 94, 127, 127, 5, 2, 11, 12, 14, 15, 34,
	34, 4, 2, 36, 36, 94, 94, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 2, 142,
	2, 4, 3, 2, 2, 2, 2, 8, 3, 2, 2, 2, 2, 10, 3, 2, 2, 2, 2, 12, 3, 2, 2,
	2, 3, 14, 3, 2, 2, 2, 3, 16, 3, 2, 2, 2, 3, 18, 3, 2, 2, 2, 3, 20, 3, 2,
	2, 2, 3, 22, 3, 2, 2, 2, 3, 24, 3, 2, 2, 2, 3, 26, 3, 2, 2, 2, 3, 28, 3,
	2, 2, 2, 3, 32, 3, 2, 2, 2, 3, 34, 3, 2, 2, 2, 3, 36, 3, 2, 2, 2, 4, 38,
	3, 2, 2, 2, 6, 42, 3, 2, 2, 2, 8, 47, 3, 2, 2, 2, 10, 66, 3, 2, 2, 2, 12,
	70, 3, 2, 2, 2, 14, 72, 3, 2, 2, 2, 16, 74, 3, 2, 2, 2, 18, 76, 3, 2, 2,
	2, 20, 78, 3, 2, 2, 2, 22, 80, 3, 2, 2, 2, 24, 83, 3, 2, 2, 2, 26, 89,
	3, 2, 2, 2, 28, 93, 3, 2, 2, 2, 30, 102, 3, 2, 2, 2, 32, 105, 3, 2, 2,
	2, 34, 126, 3, 2, 2, 2, 36, 129, 3, 2, 2, 2, 38, 39, 7, 94, 2, 2, 39, 40,
	3, 2, 2, 2, 40, 41, 8, 2, 2, 2, 41, 5, 3, 2, 2, 2, 42, 43, 9, 2, 2, 2,
	43, 7, 3, 2, 2, 2, 44, 46, 5, 6, 3, 2, 45, 44, 3, 2, 2, 2, 46, 49, 3, 2,
	2, 2, 47, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 53, 3, 2, 2, 2, 49, 47,
	3, 2, 2, 2, 50, 51, 7, 15, 2, 2, 51, 54, 7, 12, 2, 2, 52, 54, 9, 3, 2,
	2, 53, 50, 3, 2, 2, 2, 53, 52, 3, 2, 2, 2, 54, 58, 3, 2, 2, 2, 55, 57,
	5, 6, 3, 2, 56, 55, 3, 2, 2, 2, 57, 60, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2,
	58, 59, 3, 2, 2, 2, 59, 9, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 61, 67, 10,
	4, 2, 2, 62, 63, 7, 94, 2, 2, 63, 67, 7, 94, 2, 2, 64, 65, 7, 94, 2, 2,
	65, 67, 7, 127, 2, 2, 66, 61, 3, 2, 2, 2, 66, 62, 3, 2, 2, 2, 66, 64, 3,
	2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69,
	11, 3, 2, 2, 2, 70, 71, 7, 127, 2, 2, 71, 13, 3, 2, 2, 2, 72, 73, 7, 46,
	2, 2, 73, 15, 3, 2, 2, 2, 74, 75, 7, 36, 2, 2, 75, 17, 3, 2, 2, 2, 76,
	77, 7, 93, 2, 2, 77, 19, 3, 2, 2, 2, 78, 79, 7, 95, 2, 2, 79, 21, 3, 2,
	2, 2, 80, 81, 7, 63, 2, 2, 81, 23, 3, 2, 2, 2, 82, 84, 9, 5, 2, 2, 83,
	82, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 85, 86, 3, 2, 2,
	2, 86, 87, 3, 2, 2, 2, 87, 88, 8, 12, 3, 2, 88, 25, 3, 2, 2, 2, 89, 90,
	7, 125, 2, 2, 90, 91, 3, 2, 2, 2, 91, 92, 8, 13, 4, 2, 92, 27, 3, 2, 2,
	2, 93, 97, 7, 36, 2, 2, 94, 96, 10, 6, 2, 2, 95, 94, 3, 2, 2, 2, 96, 99,
	3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 100, 3, 2, 2, 2,
	99, 97, 3, 2, 2, 2, 100, 101, 7, 36, 2, 2, 101, 29, 3, 2, 2, 2, 102, 103,
	4, 50, 59, 2, 103, 31, 3, 2, 2, 2, 104, 106, 4, 50, 59, 2, 105, 104, 3,
	2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 107, 108, 3, 2, 2,
	2, 108, 115, 3, 2, 2, 2, 109, 111, 7, 48, 2, 2, 110, 112, 4, 50, 59, 2,
	111, 110, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 113,
	114, 3, 2, 2, 2, 114, 116, 3, 2, 2, 2, 115, 109, 3, 2, 2, 2, 115, 116,
	3, 2, 2, 2, 116, 33, 3, 2, 2, 2, 117, 118, 7, 118, 2, 2, 118, 119, 7, 116,
	2, 2, 119, 120, 7, 119, 2, 2, 120, 127, 7, 103, 2, 2, 121, 122, 7, 104,
	2, 2, 122, 123, 7, 99, 2, 2, 123, 124, 7, 110, 2, 2, 124, 125, 7, 117,
	2, 2, 125, 127, 7, 103, 2, 2, 126, 117, 3, 2, 2, 2, 126, 121, 3, 2, 2,
	2, 127, 35, 3, 2, 2, 2, 128, 130, 9, 7, 2, 2, 129, 128, 3, 2, 2, 2, 130,
	131, 3, 2, 2, 2, 131, 129, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 37, 3,
	2, 2, 2, 16, 2, 3, 47, 53, 58, 66, 68, 85, 97, 107, 113, 115, 126, 131,
	5, 7, 3, 2, 8, 2, 2, 6, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE", "STRICT",
}

var lexerLiteralNames = []string{
	"", "'\\'", "", "", "'}'", "','", "'\"'", "'['", "']'", "'='", "", "'{'",
}

var lexerSymbolicNames = []string{
	"", "BACKSLASH", "NL", "TEXT", "CLOSE_BRACE", "COMMA", "QUOTE", "OPEN_BRACKET",
	"CLOSE_BRACKET", "EQUALS", "SPACE", "OPEN_BRACE", "STR", "NUMBER", "BOOLEAN",
	"ALPHANUMERIC",
}

var lexerRuleNames = []string{
	"BACKSLASH", "WS", "NL", "TEXT", "CLOSE_BRACE", "COMMA", "QUOTE", "OPEN_BRACKET",
	"CLOSE_BRACKET", "EQUALS", "SPACE", "OPEN_BRACE", "STR", "DIGIT", "NUMBER",
	"BOOLEAN", "ALPHANUMERIC",
}

type NupLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewNupLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *NupLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewNupLexer(input antlr.CharStream) *NupLexer {
	l := new(NupLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "NupLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// NupLexer tokens.
const (
	NupLexerBACKSLASH     = 1
	NupLexerNL            = 2
	NupLexerTEXT          = 3
	NupLexerCLOSE_BRACE   = 4
	NupLexerCOMMA         = 5
	NupLexerQUOTE         = 6
	NupLexerOPEN_BRACKET  = 7
	NupLexerCLOSE_BRACKET = 8
	NupLexerEQUALS        = 9
	NupLexerSPACE         = 10
	NupLexerOPEN_BRACE    = 11
	NupLexerSTR           = 12
	NupLexerNUMBER        = 13
	NupLexerBOOLEAN       = 14
	NupLexerALPHANUMERIC  = 15
)

// NupLexerSTRICT is the NupLexer mode.
const NupLexerSTRICT = 1
