// Code generated from parser/NupParser.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // NupParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 34, 161,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 3, 3, 3, 6, 3, 33, 10, 3, 13, 3, 14,
	3, 34, 3, 4, 7, 4, 38, 10, 4, 12, 4, 14, 4, 41, 11, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 7, 4, 47, 10, 4, 12, 4, 14, 4, 50, 11, 4, 3, 4, 5, 4, 53, 10, 4,
	3, 5, 3, 5, 5, 5, 57, 10, 5, 3, 5, 5, 5, 60, 10, 5, 6, 5, 62, 10, 5, 13,
	5, 14, 5, 63, 3, 6, 3, 6, 5, 6, 68, 10, 6, 3, 7, 3, 7, 3, 8, 6, 8, 73,
	10, 8, 13, 8, 14, 8, 74, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 5, 11,
	83, 10, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 5, 11, 95, 10, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 6, 11, 107, 10, 11, 13, 11, 14, 11, 108, 3,
	11, 5, 11, 112, 10, 11, 3, 11, 3, 11, 7, 11, 116, 10, 11, 12, 11, 14, 11,
	119, 11, 11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 125, 10, 11, 12, 11, 14,
	11, 128, 11, 11, 5, 11, 130, 10, 11, 3, 11, 7, 11, 133, 10, 11, 12, 11,
	14, 11, 136, 11, 11, 3, 11, 3, 11, 5, 11, 140, 10, 11, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 152, 10, 14,
	12, 14, 14, 14, 155, 11, 14, 5, 14, 157, 10, 14, 3, 14, 3, 14, 3, 14, 2,
	2, 15, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 2, 7, 5, 2, 4, 4,
	27, 27, 34, 34, 5, 2, 9, 9, 25, 25, 31, 31, 5, 2, 15, 15, 26, 26, 33, 33,
	4, 2, 11, 11, 14, 14, 3, 2, 19, 21, 2, 170, 2, 28, 3, 2, 2, 2, 4, 30, 3,
	2, 2, 2, 6, 39, 3, 2, 2, 2, 8, 61, 3, 2, 2, 2, 10, 67, 3, 2, 2, 2, 12,
	69, 3, 2, 2, 2, 14, 72, 3, 2, 2, 2, 16, 76, 3, 2, 2, 2, 18, 78, 3, 2, 2,
	2, 20, 139, 3, 2, 2, 2, 22, 141, 3, 2, 2, 2, 24, 143, 3, 2, 2, 2, 26, 147,
	3, 2, 2, 2, 28, 29, 9, 2, 2, 2, 29, 3, 3, 2, 2, 2, 30, 32, 5, 2, 2, 2,
	31, 33, 5, 2, 2, 2, 32, 31, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 32, 3,
	2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 5, 3, 2, 2, 2, 36, 38, 5, 2, 2, 2, 37,
	36, 3, 2, 2, 2, 38, 41, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2,
	2, 40, 42, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 42, 48, 5, 8, 5, 2, 43, 44,
	5, 4, 3, 2, 44, 45, 5, 8, 5, 2, 45, 47, 3, 2, 2, 2, 46, 43, 3, 2, 2, 2,
	47, 50, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 52, 3,
	2, 2, 2, 50, 48, 3, 2, 2, 2, 51, 53, 5, 4, 3, 2, 52, 51, 3, 2, 2, 2, 52,
	53, 3, 2, 2, 2, 53, 7, 3, 2, 2, 2, 54, 56, 5, 10, 6, 2, 55, 57, 5, 2, 2,
	2, 56, 55, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 59, 3, 2, 2, 2, 58, 60,
	5, 10, 6, 2, 59, 58, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 62, 3, 2, 2, 2,
	61, 54, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 63, 64, 3,
	2, 2, 2, 64, 9, 3, 2, 2, 2, 65, 68, 5, 20, 11, 2, 66, 68, 5, 12, 7, 2,
	67, 65, 3, 2, 2, 2, 67, 66, 3, 2, 2, 2, 68, 11, 3, 2, 2, 2, 69, 70, 9,
	3, 2, 2, 70, 13, 3, 2, 2, 2, 71, 73, 7, 22, 2, 2, 72, 71, 3, 2, 2, 2, 73,
	74, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 15, 3, 2, 2,
	2, 76, 77, 9, 4, 2, 2, 77, 17, 3, 2, 2, 2, 78, 79, 9, 5, 2, 2, 79, 19,
	3, 2, 2, 2, 80, 82, 7, 5, 2, 2, 81, 83, 5, 26, 14, 2, 82, 81, 3, 2, 2,
	2, 82, 83, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 85, 5, 2, 2, 2, 85, 86,
	5, 12, 7, 2, 86, 87, 7, 23, 2, 2, 87, 140, 3, 2, 2, 2, 88, 89, 7, 6, 2,
	2, 89, 90, 5, 12, 7, 2, 90, 91, 7, 24, 2, 2, 91, 140, 3, 2, 2, 2, 92, 94,
	7, 8, 2, 2, 93, 95, 5, 26, 14, 2, 94, 93, 3, 2, 2, 2, 94, 95, 3, 2, 2,
	2, 95, 96, 3, 2, 2, 2, 96, 97, 5, 2, 2, 2, 97, 98, 5, 12, 7, 2, 98, 99,
	7, 30, 2, 2, 99, 140, 3, 2, 2, 2, 100, 101, 7, 7, 2, 2, 101, 102, 5, 12,
	7, 2, 102, 103, 7, 29, 2, 2, 103, 140, 3, 2, 2, 2, 104, 106, 7, 3, 2, 2,
	105, 107, 7, 22, 2, 2, 106, 105, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108,
	106, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 111, 3, 2, 2, 2, 110, 112,
	5, 26, 14, 2, 111, 110, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 113, 3,
	2, 2, 2, 113, 117, 5, 18, 10, 2, 114, 116, 5, 2, 2, 2, 115, 114, 3, 2,
	2, 2, 116, 119, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2,
	118, 129, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 120, 126, 5, 8, 5, 2, 121,
	122, 5, 4, 3, 2, 122, 123, 5, 8, 5, 2, 123, 125, 3, 2, 2, 2, 124, 121,
	3, 2, 2, 2, 125, 128, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 126, 127, 3, 2,
	2, 2, 127, 130, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 129, 120, 3, 2, 2, 2,
	129, 130, 3, 2, 2, 2, 130, 134, 3, 2, 2, 2, 131, 133, 5, 2, 2, 2, 132,
	131, 3, 2, 2, 2, 133, 136, 3, 2, 2, 2, 134, 132, 3, 2, 2, 2, 134, 135,
	3, 2, 2, 2, 135, 137, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 137, 138, 7, 10,
	2, 2, 138, 140, 3, 2, 2, 2, 139, 80, 3, 2, 2, 2, 139, 88, 3, 2, 2, 2, 139,
	92, 3, 2, 2, 2, 139, 100, 3, 2, 2, 2, 139, 104, 3, 2, 2, 2, 140, 21, 3,
	2, 2, 2, 141, 142, 9, 6, 2, 2, 142, 23, 3, 2, 2, 2, 143, 144, 5, 14, 8,
	2, 144, 145, 7, 17, 2, 2, 145, 146, 5, 22, 12, 2, 146, 25, 3, 2, 2, 2,
	147, 156, 5, 16, 9, 2, 148, 153, 5, 24, 13, 2, 149, 150, 7, 12, 2, 2, 150,
	152, 5, 24, 13, 2, 151, 149, 3, 2, 2, 2, 152, 155, 3, 2, 2, 2, 153, 151,
	3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 157, 3, 2, 2, 2, 155, 153, 3, 2,
	2, 2, 156, 148, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2,
	158, 159, 7, 16, 2, 2, 159, 27, 3, 2, 2, 2, 22, 34, 39, 48, 52, 56, 59,
	63, 67, 74, 82, 94, 108, 111, 117, 126, 129, 134, 139, 153, 156,
}
var literalNames = []string{
	"", "'\\'", "", "", "", "", "", "", "'}'", "", "','", "'\"'", "", "", "']'",
	"'='",
}
var symbolicNames = []string{
	"", "BACKSLASH", "NL", "DOLLARS", "DOLLAR", "GRAVE", "GRAVES", "TEXT",
	"CLOSE_BRACE", "OPEN_BRACE", "COMMA", "QUOTE", "SOPEN_BRACE", "OPEN_BRACKET",
	"CLOSE_BRACKET", "EQUALS", "SPACE", "STR", "NUMBER", "BOOLEAN", "ALPHANUMERIC",
	"MDOLLARS", "MDOLLAR", "MTEXT", "MABRACKET", "MANL", "CNL", "CGRAVE", "CGRAVES",
	"CTEXT", "LANGUAGE", "CABRACKET", "CANL",
}

var ruleNames = []string{
	"newLine", "blankLines", "document", "block", "content", "text", "identifier",
	"openBracket", "openBrace", "command", "val", "attr", "attrs",
}

type NupParser struct {
	*antlr.BaseParser
}

// NewNupParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *NupParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewNupParser(input antlr.TokenStream) *NupParser {
	this := new(NupParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "NupParser.g4"

	return this
}

// NupParser tokens.
const (
	NupParserEOF           = antlr.TokenEOF
	NupParserBACKSLASH     = 1
	NupParserNL            = 2
	NupParserDOLLARS       = 3
	NupParserDOLLAR        = 4
	NupParserGRAVE         = 5
	NupParserGRAVES        = 6
	NupParserTEXT          = 7
	NupParserCLOSE_BRACE   = 8
	NupParserOPEN_BRACE    = 9
	NupParserCOMMA         = 10
	NupParserQUOTE         = 11
	NupParserSOPEN_BRACE   = 12
	NupParserOPEN_BRACKET  = 13
	NupParserCLOSE_BRACKET = 14
	NupParserEQUALS        = 15
	NupParserSPACE         = 16
	NupParserSTR           = 17
	NupParserNUMBER        = 18
	NupParserBOOLEAN       = 19
	NupParserALPHANUMERIC  = 20
	NupParserMDOLLARS      = 21
	NupParserMDOLLAR       = 22
	NupParserMTEXT         = 23
	NupParserMABRACKET     = 24
	NupParserMANL          = 25
	NupParserCNL           = 26
	NupParserCGRAVE        = 27
	NupParserCGRAVES       = 28
	NupParserCTEXT         = 29
	NupParserLANGUAGE      = 30
	NupParserCABRACKET     = 31
	NupParserCANL          = 32
)

// NupParser rules.
const (
	NupParserRULE_newLine     = 0
	NupParserRULE_blankLines  = 1
	NupParserRULE_document    = 2
	NupParserRULE_block       = 3
	NupParserRULE_content     = 4
	NupParserRULE_text        = 5
	NupParserRULE_identifier  = 6
	NupParserRULE_openBracket = 7
	NupParserRULE_openBrace   = 8
	NupParserRULE_command     = 9
	NupParserRULE_val         = 10
	NupParserRULE_attr        = 11
	NupParserRULE_attrs       = 12
)

// INewLineContext is an interface to support dynamic dispatch.
type INewLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNewLineContext differentiates from other interfaces.
	IsNewLineContext()
}

type NewLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNewLineContext() *NewLineContext {
	var p = new(NewLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NupParserRULE_newLine
	return p
}

func (*NewLineContext) IsNewLineContext() {}

func NewNewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NewLineContext {
	var p = new(NewLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NupParserRULE_newLine

	return p
}

func (s *NewLineContext) GetParser() antlr.Parser { return s.parser }

func (s *NewLineContext) NL() antlr.TerminalNode {
	return s.GetToken(NupParserNL, 0)
}

func (s *NewLineContext) MANL() antlr.TerminalNode {
	return s.GetToken(NupParserMANL, 0)
}

func (s *NewLineContext) CANL() antlr.TerminalNode {
	return s.GetToken(NupParserCANL, 0)
}

func (s *NewLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NewLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.EnterNewLine(s)
	}
}

func (s *NewLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.ExitNewLine(s)
	}
}

func (p *NupParser) NewLine() (localctx INewLineContext) {
	localctx = NewNewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NupParserRULE_newLine)
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
		p.SetState(26)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-2)&-(0x1f+1)) == 0 && ((1<<uint((_la-2)))&((1<<(NupParserNL-2))|(1<<(NupParserMANL-2))|(1<<(NupParserCANL-2)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBlankLinesContext is an interface to support dynamic dispatch.
type IBlankLinesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlankLinesContext differentiates from other interfaces.
	IsBlankLinesContext()
}

type BlankLinesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlankLinesContext() *BlankLinesContext {
	var p = new(BlankLinesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NupParserRULE_blankLines
	return p
}

func (*BlankLinesContext) IsBlankLinesContext() {}

func NewBlankLinesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlankLinesContext {
	var p = new(BlankLinesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NupParserRULE_blankLines

	return p
}

func (s *BlankLinesContext) GetParser() antlr.Parser { return s.parser }

func (s *BlankLinesContext) AllNewLine() []INewLineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INewLineContext)(nil)).Elem())
	var tst = make([]INewLineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INewLineContext)
		}
	}

	return tst
}

func (s *BlankLinesContext) NewLine(i int) INewLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewLineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INewLineContext)
}

func (s *BlankLinesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlankLinesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlankLinesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.EnterBlankLines(s)
	}
}

func (s *BlankLinesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.ExitBlankLines(s)
	}
}

func (p *NupParser) BlankLines() (localctx IBlankLinesContext) {
	localctx = NewBlankLinesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NupParserRULE_blankLines)
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
		p.SetState(28)
		p.NewLine()
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-2)&-(0x1f+1)) == 0 && ((1<<uint((_la-2)))&((1<<(NupParserNL-2))|(1<<(NupParserMANL-2))|(1<<(NupParserCANL-2)))) != 0) {
		{
			p.SetState(29)
			p.NewLine()
		}

		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NupParserRULE_document
	return p
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NupParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) AllBlock() []IBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockContext)(nil)).Elem())
	var tst = make([]IBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockContext)
		}
	}

	return tst
}

func (s *DocumentContext) Block(i int) IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *DocumentContext) AllNewLine() []INewLineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INewLineContext)(nil)).Elem())
	var tst = make([]INewLineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INewLineContext)
		}
	}

	return tst
}

func (s *DocumentContext) NewLine(i int) INewLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewLineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INewLineContext)
}

func (s *DocumentContext) AllBlankLines() []IBlankLinesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlankLinesContext)(nil)).Elem())
	var tst = make([]IBlankLinesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlankLinesContext)
		}
	}

	return tst
}

func (s *DocumentContext) BlankLines(i int) IBlankLinesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlankLinesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlankLinesContext)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (p *NupParser) Document() (localctx IDocumentContext) {
	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NupParserRULE_document)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-2)&-(0x1f+1)) == 0 && ((1<<uint((_la-2)))&((1<<(NupParserNL-2))|(1<<(NupParserMANL-2))|(1<<(NupParserCANL-2)))) != 0 {
		{
			p.SetState(34)
			p.NewLine()
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(40)
		p.Block()
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(41)
				p.BlankLines()
			}
			{
				p.SetState(42)
				p.Block()
			}

		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-2)&-(0x1f+1)) == 0 && ((1<<uint((_la-2)))&((1<<(NupParserNL-2))|(1<<(NupParserMANL-2))|(1<<(NupParserCANL-2)))) != 0 {
		{
			p.SetState(49)
			p.BlankLines()
		}

	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NupParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NupParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllContent() []IContentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IContentContext)(nil)).Elem())
	var tst = make([]IContentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IContentContext)
		}
	}

	return tst
}

func (s *BlockContext) Content(i int) IContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IContentContext)
}

func (s *BlockContext) AllNewLine() []INewLineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INewLineContext)(nil)).Elem())
	var tst = make([]INewLineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INewLineContext)
		}
	}

	return tst
}

func (s *BlockContext) NewLine(i int) INewLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewLineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INewLineContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *NupParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, NupParserRULE_block)
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
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NupParserBACKSLASH)|(1<<NupParserDOLLARS)|(1<<NupParserDOLLAR)|(1<<NupParserGRAVE)|(1<<NupParserGRAVES)|(1<<NupParserTEXT)|(1<<NupParserMTEXT)|(1<<NupParserCTEXT))) != 0) {
		{
			p.SetState(52)
			p.Content()
		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(53)
				p.NewLine()
			}

		}
		p.SetState(57)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(56)
				p.Content()
			}

		}

		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IContentContext is an interface to support dynamic dispatch.
type IContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContentContext differentiates from other interfaces.
	IsContentContext()
}

type ContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContentContext() *ContentContext {
	var p = new(ContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NupParserRULE_content
	return p
}

func (*ContentContext) IsContentContext() {}

func NewContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContentContext {
	var p = new(ContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NupParserRULE_content

	return p
}

func (s *ContentContext) GetParser() antlr.Parser { return s.parser }

func (s *ContentContext) Command() ICommandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *ContentContext) Text() ITextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *ContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.EnterContent(s)
	}
}

func (s *ContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.ExitContent(s)
	}
}

func (p *NupParser) Content() (localctx IContentContext) {
	localctx = NewContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, NupParserRULE_content)

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

	p.SetState(65)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NupParserBACKSLASH, NupParserDOLLARS, NupParserDOLLAR, NupParserGRAVE, NupParserGRAVES:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)
			p.Command()
		}

	case NupParserTEXT, NupParserMTEXT, NupParserCTEXT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(64)
			p.Text()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITextContext is an interface to support dynamic dispatch.
type ITextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTextContext differentiates from other interfaces.
	IsTextContext()
}

type TextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextContext() *TextContext {
	var p = new(TextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NupParserRULE_text
	return p
}

func (*TextContext) IsTextContext() {}

func NewTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextContext {
	var p = new(TextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NupParserRULE_text

	return p
}

func (s *TextContext) GetParser() antlr.Parser { return s.parser }

func (s *TextContext) TEXT() antlr.TerminalNode {
	return s.GetToken(NupParserTEXT, 0)
}

func (s *TextContext) MTEXT() antlr.TerminalNode {
	return s.GetToken(NupParserMTEXT, 0)
}

func (s *TextContext) CTEXT() antlr.TerminalNode {
	return s.GetToken(NupParserCTEXT, 0)
}

func (s *TextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.EnterText(s)
	}
}

func (s *TextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.ExitText(s)
	}
}

func (p *NupParser) Text() (localctx ITextContext) {
	localctx = NewTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, NupParserRULE_text)
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
		p.SetState(67)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NupParserTEXT)|(1<<NupParserMTEXT)|(1<<NupParserCTEXT))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NupParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NupParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) AllALPHANUMERIC() []antlr.TerminalNode {
	return s.GetTokens(NupParserALPHANUMERIC)
}

func (s *IdentifierContext) ALPHANUMERIC(i int) antlr.TerminalNode {
	return s.GetToken(NupParserALPHANUMERIC, i)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *NupParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, NupParserRULE_identifier)
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
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == NupParserALPHANUMERIC {
		{
			p.SetState(69)
			p.Match(NupParserALPHANUMERIC)
		}

		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOpenBracketContext is an interface to support dynamic dispatch.
type IOpenBracketContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpenBracketContext differentiates from other interfaces.
	IsOpenBracketContext()
}

type OpenBracketContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpenBracketContext() *OpenBracketContext {
	var p = new(OpenBracketContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NupParserRULE_openBracket
	return p
}

func (*OpenBracketContext) IsOpenBracketContext() {}

func NewOpenBracketContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpenBracketContext {
	var p = new(OpenBracketContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NupParserRULE_openBracket

	return p
}

func (s *OpenBracketContext) GetParser() antlr.Parser { return s.parser }

func (s *OpenBracketContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(NupParserOPEN_BRACKET, 0)
}

func (s *OpenBracketContext) CABRACKET() antlr.TerminalNode {
	return s.GetToken(NupParserCABRACKET, 0)
}

func (s *OpenBracketContext) MABRACKET() antlr.TerminalNode {
	return s.GetToken(NupParserMABRACKET, 0)
}

func (s *OpenBracketContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpenBracketContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpenBracketContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.EnterOpenBracket(s)
	}
}

func (s *OpenBracketContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.ExitOpenBracket(s)
	}
}

func (p *NupParser) OpenBracket() (localctx IOpenBracketContext) {
	localctx = NewOpenBracketContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, NupParserRULE_openBracket)
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
		p.SetState(74)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NupParserOPEN_BRACKET)|(1<<NupParserMABRACKET)|(1<<NupParserCABRACKET))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOpenBraceContext is an interface to support dynamic dispatch.
type IOpenBraceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpenBraceContext differentiates from other interfaces.
	IsOpenBraceContext()
}

type OpenBraceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpenBraceContext() *OpenBraceContext {
	var p = new(OpenBraceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NupParserRULE_openBrace
	return p
}

func (*OpenBraceContext) IsOpenBraceContext() {}

func NewOpenBraceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpenBraceContext {
	var p = new(OpenBraceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NupParserRULE_openBrace

	return p
}

func (s *OpenBraceContext) GetParser() antlr.Parser { return s.parser }

func (s *OpenBraceContext) OPEN_BRACE() antlr.TerminalNode {
	return s.GetToken(NupParserOPEN_BRACE, 0)
}

func (s *OpenBraceContext) SOPEN_BRACE() antlr.TerminalNode {
	return s.GetToken(NupParserSOPEN_BRACE, 0)
}

func (s *OpenBraceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpenBraceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpenBraceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.EnterOpenBrace(s)
	}
}

func (s *OpenBraceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.ExitOpenBrace(s)
	}
}

func (p *NupParser) OpenBrace() (localctx IOpenBraceContext) {
	localctx = NewOpenBraceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, NupParserRULE_openBrace)
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
		p.SetState(76)
		_la = p.GetTokenStream().LA(1)

		if !(_la == NupParserOPEN_BRACE || _la == NupParserSOPEN_BRACE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCmd returns the cmd token.
	GetCmd() antlr.Token

	// SetCmd sets the cmd token.
	SetCmd(antlr.Token)

	// GetAttributes returns the attributes rule contexts.
	GetAttributes() IAttrsContext

	// GetInner returns the inner rule contexts.
	GetInner() IBlockContext

	// SetAttributes sets the attributes rule contexts.
	SetAttributes(IAttrsContext)

	// SetInner sets the inner rule contexts.
	SetInner(IBlockContext)

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	cmd        antlr.Token
	attributes IAttrsContext
	inner      IBlockContext
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NupParserRULE_command
	return p
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NupParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) GetCmd() antlr.Token { return s.cmd }

func (s *CommandContext) SetCmd(v antlr.Token) { s.cmd = v }

func (s *CommandContext) GetAttributes() IAttrsContext { return s.attributes }

func (s *CommandContext) GetInner() IBlockContext { return s.inner }

func (s *CommandContext) SetAttributes(v IAttrsContext) { s.attributes = v }

func (s *CommandContext) SetInner(v IBlockContext) { s.inner = v }

func (s *CommandContext) AllNewLine() []INewLineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INewLineContext)(nil)).Elem())
	var tst = make([]INewLineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INewLineContext)
		}
	}

	return tst
}

func (s *CommandContext) NewLine(i int) INewLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewLineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INewLineContext)
}

func (s *CommandContext) Text() ITextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *CommandContext) MDOLLARS() antlr.TerminalNode {
	return s.GetToken(NupParserMDOLLARS, 0)
}

func (s *CommandContext) DOLLARS() antlr.TerminalNode {
	return s.GetToken(NupParserDOLLARS, 0)
}

func (s *CommandContext) Attrs() IAttrsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrsContext)
}

func (s *CommandContext) MDOLLAR() antlr.TerminalNode {
	return s.GetToken(NupParserMDOLLAR, 0)
}

func (s *CommandContext) DOLLAR() antlr.TerminalNode {
	return s.GetToken(NupParserDOLLAR, 0)
}

func (s *CommandContext) CGRAVES() antlr.TerminalNode {
	return s.GetToken(NupParserCGRAVES, 0)
}

func (s *CommandContext) GRAVES() antlr.TerminalNode {
	return s.GetToken(NupParserGRAVES, 0)
}

func (s *CommandContext) CGRAVE() antlr.TerminalNode {
	return s.GetToken(NupParserCGRAVE, 0)
}

func (s *CommandContext) GRAVE() antlr.TerminalNode {
	return s.GetToken(NupParserGRAVE, 0)
}

func (s *CommandContext) BACKSLASH() antlr.TerminalNode {
	return s.GetToken(NupParserBACKSLASH, 0)
}

func (s *CommandContext) OpenBrace() IOpenBraceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpenBraceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpenBraceContext)
}

func (s *CommandContext) CLOSE_BRACE() antlr.TerminalNode {
	return s.GetToken(NupParserCLOSE_BRACE, 0)
}

func (s *CommandContext) AllALPHANUMERIC() []antlr.TerminalNode {
	return s.GetTokens(NupParserALPHANUMERIC)
}

func (s *CommandContext) ALPHANUMERIC(i int) antlr.TerminalNode {
	return s.GetToken(NupParserALPHANUMERIC, i)
}

func (s *CommandContext) AllBlock() []IBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockContext)(nil)).Elem())
	var tst = make([]IBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockContext)
		}
	}

	return tst
}

func (s *CommandContext) Block(i int) IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *CommandContext) AllBlankLines() []IBlankLinesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlankLinesContext)(nil)).Elem())
	var tst = make([]IBlankLinesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlankLinesContext)
		}
	}

	return tst
}

func (s *CommandContext) BlankLines(i int) IBlankLinesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlankLinesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlankLinesContext)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.EnterCommand(s)
	}
}

func (s *CommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.ExitCommand(s)
	}
}

func (p *NupParser) Command() (localctx ICommandContext) {
	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, NupParserRULE_command)
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

	var _alt int

	p.SetState(137)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NupParserDOLLARS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(78)

			var _m = p.Match(NupParserDOLLARS)

			localctx.(*CommandContext).cmd = _m
		}
		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NupParserOPEN_BRACKET)|(1<<NupParserMABRACKET)|(1<<NupParserCABRACKET))) != 0 {
			{
				p.SetState(79)

				var _x = p.Attrs()

				localctx.(*CommandContext).attributes = _x
			}

		}
		{
			p.SetState(82)
			p.NewLine()
		}
		{
			p.SetState(83)
			p.Text()
		}
		{
			p.SetState(84)
			p.Match(NupParserMDOLLARS)
		}

	case NupParserDOLLAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(86)

			var _m = p.Match(NupParserDOLLAR)

			localctx.(*CommandContext).cmd = _m
		}
		{
			p.SetState(87)
			p.Text()
		}
		{
			p.SetState(88)
			p.Match(NupParserMDOLLAR)
		}

	case NupParserGRAVES:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(90)

			var _m = p.Match(NupParserGRAVES)

			localctx.(*CommandContext).cmd = _m
		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NupParserOPEN_BRACKET)|(1<<NupParserMABRACKET)|(1<<NupParserCABRACKET))) != 0 {
			{
				p.SetState(91)

				var _x = p.Attrs()

				localctx.(*CommandContext).attributes = _x
			}

		}
		{
			p.SetState(94)
			p.NewLine()
		}
		{
			p.SetState(95)
			p.Text()
		}
		{
			p.SetState(96)
			p.Match(NupParserCGRAVES)
		}

	case NupParserGRAVE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(98)

			var _m = p.Match(NupParserGRAVE)

			localctx.(*CommandContext).cmd = _m
		}
		{
			p.SetState(99)
			p.Text()
		}
		{
			p.SetState(100)
			p.Match(NupParserCGRAVE)
		}

	case NupParserBACKSLASH:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(102)
			p.Match(NupParserBACKSLASH)
		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == NupParserALPHANUMERIC {
			{
				p.SetState(103)

				var _m = p.Match(NupParserALPHANUMERIC)

				localctx.(*CommandContext).cmd = _m
			}

			p.SetState(106)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NupParserOPEN_BRACKET)|(1<<NupParserMABRACKET)|(1<<NupParserCABRACKET))) != 0 {
			{
				p.SetState(108)

				var _x = p.Attrs()

				localctx.(*CommandContext).attributes = _x
			}

		}
		{
			p.SetState(111)
			p.OpenBrace()
		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(112)
					p.NewLine()
				}

			}
			p.SetState(117)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NupParserBACKSLASH)|(1<<NupParserDOLLARS)|(1<<NupParserDOLLAR)|(1<<NupParserGRAVE)|(1<<NupParserGRAVES)|(1<<NupParserTEXT)|(1<<NupParserMTEXT)|(1<<NupParserCTEXT))) != 0 {
			{
				p.SetState(118)

				var _x = p.Block()

				localctx.(*CommandContext).inner = _x
			}
			p.SetState(124)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					{
						p.SetState(119)
						p.BlankLines()
					}
					{
						p.SetState(120)

						var _x = p.Block()

						localctx.(*CommandContext).inner = _x
					}

				}
				p.SetState(126)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
			}

		}
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-2)&-(0x1f+1)) == 0 && ((1<<uint((_la-2)))&((1<<(NupParserNL-2))|(1<<(NupParserMANL-2))|(1<<(NupParserCANL-2)))) != 0 {
			{
				p.SetState(129)
				p.NewLine()
			}

			p.SetState(134)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(135)
			p.Match(NupParserCLOSE_BRACE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IValContext is an interface to support dynamic dispatch.
type IValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValContext differentiates from other interfaces.
	IsValContext()
}

type ValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValContext() *ValContext {
	var p = new(ValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NupParserRULE_val
	return p
}

func (*ValContext) IsValContext() {}

func NewValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValContext {
	var p = new(ValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NupParserRULE_val

	return p
}

func (s *ValContext) GetParser() antlr.Parser { return s.parser }

func (s *ValContext) STR() antlr.TerminalNode {
	return s.GetToken(NupParserSTR, 0)
}

func (s *ValContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(NupParserNUMBER, 0)
}

func (s *ValContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(NupParserBOOLEAN, 0)
}

func (s *ValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.EnterVal(s)
	}
}

func (s *ValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.ExitVal(s)
	}
}

func (p *NupParser) Val() (localctx IValContext) {
	localctx = NewValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, NupParserRULE_val)
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
		p.SetState(139)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NupParserSTR)|(1<<NupParserNUMBER)|(1<<NupParserBOOLEAN))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAttrContext is an interface to support dynamic dispatch.
type IAttrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name rule contexts.
	GetName() IIdentifierContext

	// GetValue returns the value rule contexts.
	GetValue() IValContext

	// SetName sets the name rule contexts.
	SetName(IIdentifierContext)

	// SetValue sets the value rule contexts.
	SetValue(IValContext)

	// IsAttrContext differentiates from other interfaces.
	IsAttrContext()
}

type AttrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   IIdentifierContext
	value  IValContext
}

func NewEmptyAttrContext() *AttrContext {
	var p = new(AttrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NupParserRULE_attr
	return p
}

func (*AttrContext) IsAttrContext() {}

func NewAttrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrContext {
	var p = new(AttrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NupParserRULE_attr

	return p
}

func (s *AttrContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrContext) GetName() IIdentifierContext { return s.name }

func (s *AttrContext) GetValue() IValContext { return s.value }

func (s *AttrContext) SetName(v IIdentifierContext) { s.name = v }

func (s *AttrContext) SetValue(v IValContext) { s.value = v }

func (s *AttrContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(NupParserEQUALS, 0)
}

func (s *AttrContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AttrContext) Val() IValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValContext)
}

func (s *AttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.EnterAttr(s)
	}
}

func (s *AttrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.ExitAttr(s)
	}
}

func (p *NupParser) Attr() (localctx IAttrContext) {
	localctx = NewAttrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, NupParserRULE_attr)

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
		p.SetState(141)

		var _x = p.Identifier()

		localctx.(*AttrContext).name = _x
	}
	{
		p.SetState(142)
		p.Match(NupParserEQUALS)
	}
	{
		p.SetState(143)

		var _x = p.Val()

		localctx.(*AttrContext).value = _x
	}

	return localctx
}

// IAttrsContext is an interface to support dynamic dispatch.
type IAttrsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAttribute returns the attribute rule contexts.
	GetAttribute() IAttrContext

	// SetAttribute sets the attribute rule contexts.
	SetAttribute(IAttrContext)

	// IsAttrsContext differentiates from other interfaces.
	IsAttrsContext()
}

type AttrsContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	attribute IAttrContext
}

func NewEmptyAttrsContext() *AttrsContext {
	var p = new(AttrsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NupParserRULE_attrs
	return p
}

func (*AttrsContext) IsAttrsContext() {}

func NewAttrsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrsContext {
	var p = new(AttrsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NupParserRULE_attrs

	return p
}

func (s *AttrsContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrsContext) GetAttribute() IAttrContext { return s.attribute }

func (s *AttrsContext) SetAttribute(v IAttrContext) { s.attribute = v }

func (s *AttrsContext) OpenBracket() IOpenBracketContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpenBracketContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpenBracketContext)
}

func (s *AttrsContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(NupParserCLOSE_BRACKET, 0)
}

func (s *AttrsContext) AllAttr() []IAttrContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttrContext)(nil)).Elem())
	var tst = make([]IAttrContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttrContext)
		}
	}

	return tst
}

func (s *AttrsContext) Attr(i int) IAttrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttrContext)
}

func (s *AttrsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(NupParserCOMMA)
}

func (s *AttrsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(NupParserCOMMA, i)
}

func (s *AttrsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.EnterAttrs(s)
	}
}

func (s *AttrsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.ExitAttrs(s)
	}
}

func (p *NupParser) Attrs() (localctx IAttrsContext) {
	localctx = NewAttrsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, NupParserRULE_attrs)
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
		p.SetState(145)
		p.OpenBracket()
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NupParserALPHANUMERIC {
		{
			p.SetState(146)

			var _x = p.Attr()

			localctx.(*AttrsContext).attribute = _x
		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == NupParserCOMMA {
			{
				p.SetState(147)
				p.Match(NupParserCOMMA)
			}
			{
				p.SetState(148)

				var _x = p.Attr()

				localctx.(*AttrsContext).attribute = _x
			}

			p.SetState(153)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(156)
		p.Match(NupParserCLOSE_BRACKET)
	}

	return localctx
}
