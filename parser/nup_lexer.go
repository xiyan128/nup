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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 34, 266,
	8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4,
	4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10,
	4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4,
	16, 9, 16, 4, 17, 9, 17, 4, 18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21,
	9, 21, 4, 22, 9, 22, 4, 23, 9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9,
	26, 4, 27, 9, 27, 4, 28, 9, 28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31,
	4, 32, 9, 32, 4, 33, 9, 33, 4, 34, 9, 34, 4, 35, 9, 35, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 4, 7, 4, 84, 10, 4, 12, 4, 14, 4, 87, 11, 4, 3,
	4, 3, 4, 3, 4, 5, 4, 92, 10, 4, 3, 4, 7, 4, 95, 10, 4, 12, 4, 14, 4, 98,
	11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 6, 9, 128, 10, 9, 13, 9, 14, 9, 129, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18,
	3, 18, 3, 18, 5, 18, 156, 10, 18, 3, 18, 3, 18, 3, 19, 3, 19, 7, 19, 162,
	10, 19, 12, 19, 14, 19, 165, 11, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21,
	6, 21, 172, 10, 21, 13, 21, 14, 21, 173, 3, 21, 3, 21, 6, 21, 178, 10,
	21, 13, 21, 14, 21, 179, 5, 21, 182, 10, 21, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 193, 10, 22, 3, 23, 6, 23, 196,
	10, 23, 13, 23, 14, 23, 197, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25,
	3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 6, 26, 212, 10, 26, 13, 26, 14,
	26, 213, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 5, 28, 223, 10,
	28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 5, 29, 231, 10, 29, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3,
	32, 3, 32, 6, 32, 246, 10, 32, 13, 32, 14, 32, 247, 3, 33, 6, 33, 251,
	10, 33, 13, 33, 14, 33, 252, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35,
	3, 35, 5, 35, 262, 10, 35, 3, 35, 3, 35, 3, 35, 2, 2, 36, 8, 3, 10, 2,
	12, 4, 14, 5, 16, 6, 18, 7, 20, 8, 22, 9, 24, 10, 26, 11, 28, 12, 30, 13,
	32, 14, 34, 15, 36, 16, 38, 17, 40, 18, 42, 19, 44, 2, 46, 20, 48, 21,
	50, 22, 52, 23, 54, 24, 56, 25, 58, 26, 60, 27, 62, 28, 64, 29, 66, 30,
	68, 31, 70, 32, 72, 33, 74, 34, 8, 2, 3, 4, 5, 6, 7, 11, 4, 2, 11, 11,
	34, 34, 4, 2, 12, 12, 15, 15, 9, 2, 12, 12, 15, 15, 38, 38, 94, 94, 98,
	98, 125, 125, 127, 127, 5, 2, 11, 11, 14, 14, 34, 34, 4, 2, 36, 36, 94,
	94, 5, 2, 50, 59, 67, 92, 99, 124, 3, 2, 38, 38, 3, 2, 98, 98, 6, 2, 50,
	59, 67, 92, 97, 97, 99, 124, 2, 282, 2, 8, 3, 2, 2, 2, 2, 12, 3, 2, 2,
	2, 2, 14, 3, 2, 2, 2, 2, 16, 3, 2, 2, 2, 2, 18, 3, 2, 2, 2, 2, 20, 3, 2,
	2, 2, 2, 22, 3, 2, 2, 2, 2, 24, 3, 2, 2, 2, 2, 26, 3, 2, 2, 2, 3, 28, 3,
	2, 2, 2, 3, 30, 3, 2, 2, 2, 3, 32, 3, 2, 2, 2, 3, 34, 3, 2, 2, 2, 3, 36,
	3, 2, 2, 2, 3, 38, 3, 2, 2, 2, 3, 40, 3, 2, 2, 2, 3, 42, 3, 2, 2, 2, 3,
	46, 3, 2, 2, 2, 3, 48, 3, 2, 2, 2, 3, 50, 3, 2, 2, 2, 4, 52, 3, 2, 2, 2,
	4, 54, 3, 2, 2, 2, 4, 56, 3, 2, 2, 2, 5, 58, 3, 2, 2, 2, 5, 60, 3, 2, 2,
	2, 6, 62, 3, 2, 2, 2, 6, 64, 3, 2, 2, 2, 6, 66, 3, 2, 2, 2, 6, 68, 3, 2,
	2, 2, 6, 70, 3, 2, 2, 2, 7, 72, 3, 2, 2, 2, 7, 74, 3, 2, 2, 2, 8, 76, 3,
	2, 2, 2, 10, 80, 3, 2, 2, 2, 12, 85, 3, 2, 2, 2, 14, 99, 3, 2, 2, 2, 16,
	104, 3, 2, 2, 2, 18, 108, 3, 2, 2, 2, 20, 112, 3, 2, 2, 2, 22, 127, 3,
	2, 2, 2, 24, 131, 3, 2, 2, 2, 26, 133, 3, 2, 2, 2, 28, 135, 3, 2, 2, 2,
	30, 137, 3, 2, 2, 2, 32, 139, 3, 2, 2, 2, 34, 143, 3, 2, 2, 2, 36, 145,
	3, 2, 2, 2, 38, 149, 3, 2, 2, 2, 40, 155, 3, 2, 2, 2, 42, 159, 3, 2, 2,
	2, 44, 168, 3, 2, 2, 2, 46, 171, 3, 2, 2, 2, 48, 192, 3, 2, 2, 2, 50, 195,
	3, 2, 2, 2, 52, 199, 3, 2, 2, 2, 54, 204, 3, 2, 2, 2, 56, 211, 3, 2, 2,
	2, 58, 215, 3, 2, 2, 2, 60, 222, 3, 2, 2, 2, 62, 230, 3, 2, 2, 2, 64, 232,
	3, 2, 2, 2, 66, 236, 3, 2, 2, 2, 68, 245, 3, 2, 2, 2, 70, 250, 3, 2, 2,
	2, 72, 254, 3, 2, 2, 2, 74, 261, 3, 2, 2, 2, 76, 77, 7, 94, 2, 2, 77, 78,
	3, 2, 2, 2, 78, 79, 8, 2, 2, 2, 79, 9, 3, 2, 2, 2, 80, 81, 9, 2, 2, 2,
	81, 11, 3, 2, 2, 2, 82, 84, 5, 10, 3, 2, 83, 82, 3, 2, 2, 2, 84, 87, 3,
	2, 2, 2, 85, 83, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 91, 3, 2, 2, 2, 87,
	85, 3, 2, 2, 2, 88, 89, 7, 15, 2, 2, 89, 92, 7, 12, 2, 2, 90, 92, 9, 3,
	2, 2, 91, 88, 3, 2, 2, 2, 91, 90, 3, 2, 2, 2, 92, 96, 3, 2, 2, 2, 93, 95,
	5, 10, 3, 2, 94, 93, 3, 2, 2, 2, 95, 98, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2,
	96, 97, 3, 2, 2, 2, 97, 13, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 99, 100, 7,
	38, 2, 2, 100, 101, 7, 38, 2, 2, 101, 102, 3, 2, 2, 2, 102, 103, 8, 5,
	3, 2, 103, 15, 3, 2, 2, 2, 104, 105, 7, 38, 2, 2, 105, 106, 3, 2, 2, 2,
	106, 107, 8, 6, 4, 2, 107, 17, 3, 2, 2, 2, 108, 109, 7, 98, 2, 2, 109,
	110, 3, 2, 2, 2, 110, 111, 8, 7, 5, 2, 111, 19, 3, 2, 2, 2, 112, 113, 7,
	98, 2, 2, 113, 114, 7, 98, 2, 2, 114, 115, 7, 98, 2, 2, 115, 116, 3, 2,
	2, 2, 116, 117, 8, 8, 6, 2, 117, 21, 3, 2, 2, 2, 118, 128, 10, 4, 2, 2,
	119, 120, 7, 94, 2, 2, 120, 128, 7, 94, 2, 2, 121, 122, 7, 94, 2, 2, 122,
	128, 7, 127, 2, 2, 123, 124, 7, 94, 2, 2, 124, 128, 7, 38, 2, 2, 125, 126,
	7, 94, 2, 2, 126, 128, 7, 98, 2, 2, 127, 118, 3, 2, 2, 2, 127, 119, 3,
	2, 2, 2, 127, 121, 3, 2, 2, 2, 127, 123, 3, 2, 2, 2, 127, 125, 3, 2, 2,
	2, 128, 129, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130,
	23, 3, 2, 2, 2, 131, 132, 7, 127, 2, 2, 132, 25, 3, 2, 2, 2, 133, 134,
	7, 125, 2, 2, 134, 27, 3, 2, 2, 2, 135, 136, 7, 46, 2, 2, 136, 29, 3, 2,
	2, 2, 137, 138, 7, 36, 2, 2, 138, 31, 3, 2, 2, 2, 139, 140, 7, 125, 2,
	2, 140, 141, 3, 2, 2, 2, 141, 142, 8, 14, 7, 2, 142, 33, 3, 2, 2, 2, 143,
	144, 7, 93, 2, 2, 144, 35, 3, 2, 2, 2, 145, 146, 7, 95, 2, 2, 146, 147,
	3, 2, 2, 2, 147, 148, 8, 16, 7, 2, 148, 37, 3, 2, 2, 2, 149, 150, 7, 63,
	2, 2, 150, 39, 3, 2, 2, 2, 151, 156, 9, 5, 2, 2, 152, 153, 7, 15, 2, 2,
	153, 156, 7, 12, 2, 2, 154, 156, 9, 3, 2, 2, 155, 151, 3, 2, 2, 2, 155,
	152, 3, 2, 2, 2, 155, 154, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 158,
	8, 18, 8, 2, 158, 41, 3, 2, 2, 2, 159, 163, 7, 36, 2, 2, 160, 162, 10,
	6, 2, 2, 161, 160, 3, 2, 2, 2, 162, 165, 3, 2, 2, 2, 163, 161, 3, 2, 2,
	2, 163, 164, 3, 2, 2, 2, 164, 166, 3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 166,
	167, 7, 36, 2, 2, 167, 43, 3, 2, 2, 2, 168, 169, 4, 50, 59, 2, 169, 45,
	3, 2, 2, 2, 170, 172, 4, 50, 59, 2, 171, 170, 3, 2, 2, 2, 172, 173, 3,
	2, 2, 2, 173, 171, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 181, 3, 2, 2,
	2, 175, 177, 7, 48, 2, 2, 176, 178, 4, 50, 59, 2, 177, 176, 3, 2, 2, 2,
	178, 179, 3, 2, 2, 2, 179, 177, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180,
	182, 3, 2, 2, 2, 181, 175, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 47, 3,
	2, 2, 2, 183, 184, 7, 118, 2, 2, 184, 185, 7, 116, 2, 2, 185, 186, 7, 119,
	2, 2, 186, 193, 7, 103, 2, 2, 187, 188, 7, 104, 2, 2, 188, 189, 7, 99,
	2, 2, 189, 190, 7, 110, 2, 2, 190, 191, 7, 117, 2, 2, 191, 193, 7, 103,
	2, 2, 192, 183, 3, 2, 2, 2, 192, 187, 3, 2, 2, 2, 193, 49, 3, 2, 2, 2,
	194, 196, 9, 7, 2, 2, 195, 194, 3, 2, 2, 2, 196, 197, 3, 2, 2, 2, 197,
	195, 3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198, 51, 3, 2, 2, 2, 199, 200, 7,
	38, 2, 2, 200, 201, 7, 38, 2, 2, 201, 202, 3, 2, 2, 2, 202, 203, 8, 24,
	7, 2, 203, 53, 3, 2, 2, 2, 204, 205, 7, 38, 2, 2, 205, 206, 3, 2, 2, 2,
	206, 207, 8, 25, 7, 2, 207, 55, 3, 2, 2, 2, 208, 212, 10, 8, 2, 2, 209,
	210, 7, 94, 2, 2, 210, 212, 7, 38, 2, 2, 211, 208, 3, 2, 2, 2, 211, 209,
	3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 211, 3, 2, 2, 2, 213, 214, 3, 2,
	2, 2, 214, 57, 3, 2, 2, 2, 215, 216, 7, 93, 2, 2, 216, 217, 3, 2, 2, 2,
	217, 218, 8, 27, 2, 2, 218, 59, 3, 2, 2, 2, 219, 220, 7, 15, 2, 2, 220,
	223, 7, 12, 2, 2, 221, 223, 9, 3, 2, 2, 222, 219, 3, 2, 2, 2, 222, 221,
	3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224, 225, 8, 28, 7, 2, 225, 226, 8, 28,
	4, 2, 226, 61, 3, 2, 2, 2, 227, 228, 7, 15, 2, 2, 228, 231, 7, 12, 2, 2,
	229, 231, 9, 3, 2, 2, 230, 227, 3, 2, 2, 2, 230, 229, 3, 2, 2, 2, 231,
	63, 3, 2, 2, 2, 232, 233, 7, 98, 2, 2, 233, 234, 3, 2, 2, 2, 234, 235,
	8, 30, 7, 2, 235, 65, 3, 2, 2, 2, 236, 237, 7, 98, 2, 2, 237, 238, 7, 98,
	2, 2, 238, 239, 7, 98, 2, 2, 239, 240, 3, 2, 2, 2, 240, 241, 8, 31, 7,
	2, 241, 67, 3, 2, 2, 2, 242, 246, 10, 9, 2, 2, 243, 244, 7, 94, 2, 2, 244,
	246, 7, 98, 2, 2, 245, 242, 3, 2, 2, 2, 245, 243, 3, 2, 2, 2, 246, 247,
	3, 2, 2, 2, 247, 245, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248, 69, 3, 2,
	2, 2, 249, 251, 9, 10, 2, 2, 250, 249, 3, 2, 2, 2, 251, 252, 3, 2, 2, 2,
	252, 250, 3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253, 71, 3, 2, 2, 2, 254, 255,
	7, 93, 2, 2, 255, 256, 3, 2, 2, 2, 256, 257, 8, 34, 2, 2, 257, 73, 3, 2,
	2, 2, 258, 259, 7, 15, 2, 2, 259, 262, 7, 12, 2, 2, 260, 262, 9, 3, 2,
	2, 261, 258, 3, 2, 2, 2, 261, 260, 3, 2, 2, 2, 262, 263, 3, 2, 2, 2, 263,
	264, 8, 35, 7, 2, 264, 265, 8, 35, 5, 2, 265, 75, 3, 2, 2, 2, 28, 2, 3,
	4, 5, 6, 7, 85, 91, 96, 127, 129, 155, 163, 173, 179, 181, 192, 197, 211,
	213, 222, 230, 245, 247, 252, 261, 9, 7, 3, 2, 7, 5, 2, 7, 4, 2, 7, 6,
	2, 7, 7, 2, 6, 2, 2, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE", "STRICT", "MATH", "MATH_ATTR", "CODE", "CODE_ATTR",
}

var lexerLiteralNames = []string{
	"", "'\\'", "", "", "", "", "", "", "'}'", "", "','", "'\"'", "", "", "']'",
	"'='",
}

var lexerSymbolicNames = []string{
	"", "BACKSLASH", "NL", "DOLLARS", "DOLLAR", "GRAVE", "GRAVES", "TEXT",
	"CLOSE_BRACE", "OPEN_BRACE", "COMMA", "QUOTE", "SOPEN_BRACE", "OPEN_BRACKET",
	"CLOSE_BRACKET", "EQUALS", "SPACE", "STR", "NUMBER", "BOOLEAN", "ALPHANUMERIC",
	"MDOLLARS", "MDOLLAR", "MTEXT", "MABRACKET", "MANL", "CNL", "CGRAVE", "CGRAVES",
	"CTEXT", "LANGUAGE", "CABRACKET", "CANL",
}

var lexerRuleNames = []string{
	"BACKSLASH", "WS", "NL", "DOLLARS", "DOLLAR", "GRAVE", "GRAVES", "TEXT",
	"CLOSE_BRACE", "OPEN_BRACE", "COMMA", "QUOTE", "SOPEN_BRACE", "OPEN_BRACKET",
	"CLOSE_BRACKET", "EQUALS", "SPACE", "STR", "DIGIT", "NUMBER", "BOOLEAN",
	"ALPHANUMERIC", "MDOLLARS", "MDOLLAR", "MTEXT", "MABRACKET", "MANL", "CNL",
	"CGRAVE", "CGRAVES", "CTEXT", "LANGUAGE", "CABRACKET", "CANL",
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
	NupLexerDOLLARS       = 3
	NupLexerDOLLAR        = 4
	NupLexerGRAVE         = 5
	NupLexerGRAVES        = 6
	NupLexerTEXT          = 7
	NupLexerCLOSE_BRACE   = 8
	NupLexerOPEN_BRACE    = 9
	NupLexerCOMMA         = 10
	NupLexerQUOTE         = 11
	NupLexerSOPEN_BRACE   = 12
	NupLexerOPEN_BRACKET  = 13
	NupLexerCLOSE_BRACKET = 14
	NupLexerEQUALS        = 15
	NupLexerSPACE         = 16
	NupLexerSTR           = 17
	NupLexerNUMBER        = 18
	NupLexerBOOLEAN       = 19
	NupLexerALPHANUMERIC  = 20
	NupLexerMDOLLARS      = 21
	NupLexerMDOLLAR       = 22
	NupLexerMTEXT         = 23
	NupLexerMABRACKET     = 24
	NupLexerMANL          = 25
	NupLexerCNL           = 26
	NupLexerCGRAVE        = 27
	NupLexerCGRAVES       = 28
	NupLexerCTEXT         = 29
	NupLexerLANGUAGE      = 30
	NupLexerCABRACKET     = 31
	NupLexerCANL          = 32
)

// NupLexer modes.
const (
	NupLexerSTRICT = iota + 1
	NupLexerMATH
	NupLexerMATH_ATTR
	NupLexerCODE
	NupLexerCODE_ATTR
)
