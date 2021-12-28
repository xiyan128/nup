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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 22, 171,
	8, 1, 8, 1, 8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6,
	9, 6, 4, 7, 9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4,
	12, 9, 12, 4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17,
	9, 17, 4, 18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9,
	22, 4, 23, 9, 23, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 7, 4, 57, 10,
	4, 12, 4, 14, 4, 60, 11, 4, 3, 4, 3, 4, 3, 4, 5, 4, 65, 10, 4, 3, 4, 7,
	4, 68, 10, 4, 12, 4, 14, 4, 71, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 6, 7, 89,
	10, 7, 13, 7, 14, 7, 90, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 6, 14, 106, 10, 14, 13, 14, 14,
	14, 107, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 7, 16,
	118, 10, 16, 12, 16, 14, 16, 121, 11, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3,
	18, 6, 18, 128, 10, 18, 13, 18, 14, 18, 129, 3, 18, 3, 18, 6, 18, 134,
	10, 18, 13, 18, 14, 18, 135, 5, 18, 138, 10, 18, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 149, 10, 19, 3, 20, 6, 20,
	152, 10, 20, 13, 20, 14, 20, 153, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 6, 23, 168, 10, 23, 13, 23,
	14, 23, 169, 2, 2, 24, 5, 3, 7, 2, 9, 4, 11, 5, 13, 6, 15, 7, 17, 8, 19,
	9, 21, 10, 23, 11, 25, 12, 27, 13, 29, 14, 31, 15, 33, 16, 35, 2, 37, 17,
	39, 18, 41, 19, 43, 20, 45, 21, 47, 22, 5, 2, 3, 4, 9, 4, 2, 11, 11, 34,
	34, 4, 2, 12, 12, 15, 15, 7, 2, 12, 12, 15, 15, 38, 38, 94, 94, 127, 127,
	5, 2, 11, 12, 14, 15, 34, 34, 4, 2, 36, 36, 94, 94, 6, 2, 50, 59, 67, 92,
	97, 97, 99, 124, 3, 2, 38, 38, 2, 182, 2, 5, 3, 2, 2, 2, 2, 9, 3, 2, 2,
	2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2,
	2, 2, 3, 19, 3, 2, 2, 2, 3, 21, 3, 2, 2, 2, 3, 23, 3, 2, 2, 2, 3, 25, 3,
	2, 2, 2, 3, 27, 3, 2, 2, 2, 3, 29, 3, 2, 2, 2, 3, 31, 3, 2, 2, 2, 3, 33,
	3, 2, 2, 2, 3, 37, 3, 2, 2, 2, 3, 39, 3, 2, 2, 2, 3, 41, 3, 2, 2, 2, 4,
	43, 3, 2, 2, 2, 4, 45, 3, 2, 2, 2, 4, 47, 3, 2, 2, 2, 5, 49, 3, 2, 2, 2,
	7, 53, 3, 2, 2, 2, 9, 58, 3, 2, 2, 2, 11, 72, 3, 2, 2, 2, 13, 77, 3, 2,
	2, 2, 15, 88, 3, 2, 2, 2, 17, 92, 3, 2, 2, 2, 19, 94, 3, 2, 2, 2, 21, 96,
	3, 2, 2, 2, 23, 98, 3, 2, 2, 2, 25, 100, 3, 2, 2, 2, 27, 102, 3, 2, 2,
	2, 29, 105, 3, 2, 2, 2, 31, 111, 3, 2, 2, 2, 33, 115, 3, 2, 2, 2, 35, 124,
	3, 2, 2, 2, 37, 127, 3, 2, 2, 2, 39, 148, 3, 2, 2, 2, 41, 151, 3, 2, 2,
	2, 43, 155, 3, 2, 2, 2, 45, 160, 3, 2, 2, 2, 47, 167, 3, 2, 2, 2, 49, 50,
	7, 94, 2, 2, 50, 51, 3, 2, 2, 2, 51, 52, 8, 2, 2, 2, 52, 6, 3, 2, 2, 2,
	53, 54, 9, 2, 2, 2, 54, 8, 3, 2, 2, 2, 55, 57, 5, 7, 3, 2, 56, 55, 3, 2,
	2, 2, 57, 60, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 64,
	3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 61, 62, 7, 15, 2, 2, 62, 65, 7, 12, 2,
	2, 63, 65, 9, 3, 2, 2, 64, 61, 3, 2, 2, 2, 64, 63, 3, 2, 2, 2, 65, 69,
	3, 2, 2, 2, 66, 68, 5, 7, 3, 2, 67, 66, 3, 2, 2, 2, 68, 71, 3, 2, 2, 2,
	69, 67, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 10, 3, 2, 2, 2, 71, 69, 3,
	2, 2, 2, 72, 73, 7, 38, 2, 2, 73, 74, 7, 38, 2, 2, 74, 75, 3, 2, 2, 2,
	75, 76, 8, 5, 3, 2, 76, 12, 3, 2, 2, 2, 77, 78, 7, 38, 2, 2, 78, 79, 3,
	2, 2, 2, 79, 80, 8, 6, 3, 2, 80, 14, 3, 2, 2, 2, 81, 89, 10, 4, 2, 2, 82,
	83, 7, 94, 2, 2, 83, 89, 7, 94, 2, 2, 84, 85, 7, 94, 2, 2, 85, 89, 7, 127,
	2, 2, 86, 87, 7, 94, 2, 2, 87, 89, 7, 38, 2, 2, 88, 81, 3, 2, 2, 2, 88,
	82, 3, 2, 2, 2, 88, 84, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 89, 90, 3, 2, 2,
	2, 90, 88, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 16, 3, 2, 2, 2, 92, 93,
	7, 127, 2, 2, 93, 18, 3, 2, 2, 2, 94, 95, 7, 46, 2, 2, 95, 20, 3, 2, 2,
	2, 96, 97, 7, 36, 2, 2, 97, 22, 3, 2, 2, 2, 98, 99, 7, 93, 2, 2, 99, 24,
	3, 2, 2, 2, 100, 101, 7, 95, 2, 2, 101, 26, 3, 2, 2, 2, 102, 103, 7, 63,
	2, 2, 103, 28, 3, 2, 2, 2, 104, 106, 9, 5, 2, 2, 105, 104, 3, 2, 2, 2,
	106, 107, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108,
	109, 3, 2, 2, 2, 109, 110, 8, 14, 4, 2, 110, 30, 3, 2, 2, 2, 111, 112,
	7, 125, 2, 2, 112, 113, 3, 2, 2, 2, 113, 114, 8, 15, 5, 2, 114, 32, 3,
	2, 2, 2, 115, 119, 7, 36, 2, 2, 116, 118, 10, 6, 2, 2, 117, 116, 3, 2,
	2, 2, 118, 121, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2,
	120, 122, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 122, 123, 7, 36, 2, 2, 123,
	34, 3, 2, 2, 2, 124, 125, 4, 50, 59, 2, 125, 36, 3, 2, 2, 2, 126, 128,
	4, 50, 59, 2, 127, 126, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 127, 3,
	2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 137, 3, 2, 2, 2, 131, 133, 7, 48, 2,
	2, 132, 134, 4, 50, 59, 2, 133, 132, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2,
	135, 133, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 138, 3, 2, 2, 2, 137,
	131, 3, 2, 2, 2, 137, 138, 3, 2, 2, 2, 138, 38, 3, 2, 2, 2, 139, 140, 7,
	118, 2, 2, 140, 141, 7, 116, 2, 2, 141, 142, 7, 119, 2, 2, 142, 149, 7,
	103, 2, 2, 143, 144, 7, 104, 2, 2, 144, 145, 7, 99, 2, 2, 145, 146, 7,
	110, 2, 2, 146, 147, 7, 117, 2, 2, 147, 149, 7, 103, 2, 2, 148, 139, 3,
	2, 2, 2, 148, 143, 3, 2, 2, 2, 149, 40, 3, 2, 2, 2, 150, 152, 9, 7, 2,
	2, 151, 150, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 151, 3, 2, 2, 2, 153,
	154, 3, 2, 2, 2, 154, 42, 3, 2, 2, 2, 155, 156, 7, 38, 2, 2, 156, 157,
	7, 38, 2, 2, 157, 158, 3, 2, 2, 2, 158, 159, 8, 21, 5, 2, 159, 44, 3, 2,
	2, 2, 160, 161, 7, 38, 2, 2, 161, 162, 3, 2, 2, 2, 162, 163, 8, 22, 5,
	2, 163, 46, 3, 2, 2, 2, 164, 168, 10, 8, 2, 2, 165, 166, 7, 94, 2, 2, 166,
	168, 7, 38, 2, 2, 167, 164, 3, 2, 2, 2, 167, 165, 3, 2, 2, 2, 168, 169,
	3, 2, 2, 2, 169, 167, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170, 48, 3, 2,
	2, 2, 19, 2, 3, 4, 58, 64, 69, 88, 90, 107, 119, 129, 135, 137, 148, 153,
	167, 169, 6, 7, 3, 2, 7, 4, 2, 8, 2, 2, 6, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE", "STRICT", "MATH",
}

var lexerLiteralNames = []string{
	"", "'\\'", "", "", "", "", "'}'", "','", "'\"'", "'['", "']'", "'='",
	"", "'{'",
}

var lexerSymbolicNames = []string{
	"", "BACKSLASH", "NL", "DOLLAR", "DOLLARS", "TEXT", "CLOSE_BRACE", "COMMA",
	"QUOTE", "OPEN_BRACKET", "CLOSE_BRACKET", "EQUALS", "SPACE", "OPEN_BRACE",
	"STR", "NUMBER", "BOOLEAN", "ALPHANUMERIC", "MDOLLAR", "MDOLLARS", "MTEXT",
}

var lexerRuleNames = []string{
	"BACKSLASH", "WS", "NL", "DOLLAR", "DOLLARS", "TEXT", "CLOSE_BRACE", "COMMA",
	"QUOTE", "OPEN_BRACKET", "CLOSE_BRACKET", "EQUALS", "SPACE", "OPEN_BRACE",
	"STR", "DIGIT", "NUMBER", "BOOLEAN", "ALPHANUMERIC", "MDOLLAR", "MDOLLARS",
	"MTEXT",
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
	NupLexerDOLLAR        = 3
	NupLexerDOLLARS       = 4
	NupLexerTEXT          = 5
	NupLexerCLOSE_BRACE   = 6
	NupLexerCOMMA         = 7
	NupLexerQUOTE         = 8
	NupLexerOPEN_BRACKET  = 9
	NupLexerCLOSE_BRACKET = 10
	NupLexerEQUALS        = 11
	NupLexerSPACE         = 12
	NupLexerOPEN_BRACE    = 13
	NupLexerSTR           = 14
	NupLexerNUMBER        = 15
	NupLexerBOOLEAN       = 16
	NupLexerALPHANUMERIC  = 17
	NupLexerMDOLLAR       = 18
	NupLexerMDOLLARS      = 19
	NupLexerMTEXT         = 20
)

// NupLexer modes.
const (
	NupLexerSTRICT = iota + 1
	NupLexerMATH
)
