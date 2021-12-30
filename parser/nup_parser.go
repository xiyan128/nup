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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 30, 157,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 3, 2, 3, 2, 3, 3, 3, 3, 6, 3, 31, 10, 3, 13, 3, 14, 3, 32, 3, 4,
	7, 4, 36, 10, 4, 12, 4, 14, 4, 39, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4,
	45, 10, 4, 12, 4, 14, 4, 48, 11, 4, 3, 4, 5, 4, 51, 10, 4, 3, 5, 3, 5,
	5, 5, 55, 10, 5, 3, 5, 5, 5, 58, 10, 5, 6, 5, 60, 10, 5, 13, 5, 14, 5,
	61, 3, 6, 3, 6, 5, 6, 66, 10, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 6, 9, 73,
	10, 9, 13, 9, 14, 9, 74, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 91, 10, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 6, 10, 99, 10, 10, 13, 10, 14, 10,
	100, 3, 10, 5, 10, 104, 10, 10, 3, 10, 3, 10, 7, 10, 108, 10, 10, 12, 10,
	14, 10, 111, 11, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 117, 10, 10, 12,
	10, 14, 10, 120, 11, 10, 5, 10, 122, 10, 10, 3, 10, 7, 10, 125, 10, 10,
	12, 10, 14, 10, 128, 11, 10, 3, 10, 5, 10, 131, 10, 10, 3, 11, 3, 11, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 7, 13, 147, 10, 13, 12, 13, 14, 13, 150, 11, 13, 5, 13, 152, 10,
	13, 3, 13, 5, 13, 155, 10, 13, 3, 13, 2, 2, 14, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 22, 24, 2, 4, 5, 2, 9, 9, 25, 25, 29, 29, 3, 2, 18, 20, 2,
	167, 2, 26, 3, 2, 2, 2, 4, 28, 3, 2, 2, 2, 6, 37, 3, 2, 2, 2, 8, 59, 3,
	2, 2, 2, 10, 65, 3, 2, 2, 2, 12, 67, 3, 2, 2, 2, 14, 69, 3, 2, 2, 2, 16,
	72, 3, 2, 2, 2, 18, 130, 3, 2, 2, 2, 20, 132, 3, 2, 2, 2, 22, 134, 3, 2,
	2, 2, 24, 154, 3, 2, 2, 2, 26, 27, 7, 4, 2, 2, 27, 3, 3, 2, 2, 2, 28, 30,
	5, 2, 2, 2, 29, 31, 5, 2, 2, 2, 30, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2,
	32, 30, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 5, 3, 2, 2, 2, 34, 36, 5, 2,
	2, 2, 35, 34, 3, 2, 2, 2, 36, 39, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 37, 38,
	3, 2, 2, 2, 38, 40, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 40, 46, 5, 8, 5, 2,
	41, 42, 5, 4, 3, 2, 42, 43, 5, 8, 5, 2, 43, 45, 3, 2, 2, 2, 44, 41, 3,
	2, 2, 2, 45, 48, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47,
	50, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 49, 51, 5, 4, 3, 2, 50, 49, 3, 2, 2,
	2, 50, 51, 3, 2, 2, 2, 51, 7, 3, 2, 2, 2, 52, 54, 5, 10, 6, 2, 53, 55,
	5, 2, 2, 2, 54, 53, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 57, 3, 2, 2, 2,
	56, 58, 5, 10, 6, 2, 57, 56, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 60, 3,
	2, 2, 2, 59, 52, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 61,
	62, 3, 2, 2, 2, 62, 9, 3, 2, 2, 2, 63, 66, 5, 18, 10, 2, 64, 66, 5, 14,
	8, 2, 65, 63, 3, 2, 2, 2, 65, 64, 3, 2, 2, 2, 66, 11, 3, 2, 2, 2, 67, 68,
	7, 30, 2, 2, 68, 13, 3, 2, 2, 2, 69, 70, 9, 2, 2, 2, 70, 15, 3, 2, 2, 2,
	71, 73, 7, 21, 2, 2, 72, 71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 72, 3,
	2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 17, 3, 2, 2, 2, 76, 77, 7, 5, 2, 2, 77,
	78, 5, 14, 8, 2, 78, 79, 7, 23, 2, 2, 79, 131, 3, 2, 2, 2, 80, 81, 7, 6,
	2, 2, 81, 82, 5, 14, 8, 2, 82, 83, 7, 24, 2, 2, 83, 131, 3, 2, 2, 2, 84,
	85, 7, 7, 2, 2, 85, 86, 5, 14, 8, 2, 86, 87, 7, 27, 2, 2, 87, 131, 3, 2,
	2, 2, 88, 90, 7, 8, 2, 2, 89, 91, 5, 24, 13, 2, 90, 89, 3, 2, 2, 2, 90,
	91, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 93, 7, 22, 2, 2, 93, 94, 5, 14,
	8, 2, 94, 95, 7, 28, 2, 2, 95, 131, 3, 2, 2, 2, 96, 98, 7, 3, 2, 2, 97,
	99, 7, 21, 2, 2, 98, 97, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 98, 3, 2,
	2, 2, 100, 101, 3, 2, 2, 2, 101, 103, 3, 2, 2, 2, 102, 104, 5, 24, 13,
	2, 103, 102, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105,
	109, 7, 17, 2, 2, 106, 108, 5, 2, 2, 2, 107, 106, 3, 2, 2, 2, 108, 111,
	3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110, 121, 3, 2,
	2, 2, 111, 109, 3, 2, 2, 2, 112, 118, 5, 8, 5, 2, 113, 114, 5, 4, 3, 2,
	114, 115, 5, 8, 5, 2, 115, 117, 3, 2, 2, 2, 116, 113, 3, 2, 2, 2, 117,
	120, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 122,
	3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 121, 112, 3, 2, 2, 2, 121, 122, 3, 2,
	2, 2, 122, 126, 3, 2, 2, 2, 123, 125, 5, 2, 2, 2, 124, 123, 3, 2, 2, 2,
	125, 128, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127,
	129, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 129, 131, 7, 10, 2, 2, 130, 76,
	3, 2, 2, 2, 130, 80, 3, 2, 2, 2, 130, 84, 3, 2, 2, 2, 130, 88, 3, 2, 2,
	2, 130, 96, 3, 2, 2, 2, 131, 19, 3, 2, 2, 2, 132, 133, 9, 3, 2, 2, 133,
	21, 3, 2, 2, 2, 134, 135, 5, 16, 9, 2, 135, 136, 7, 15, 2, 2, 136, 137,
	5, 20, 11, 2, 137, 23, 3, 2, 2, 2, 138, 139, 7, 13, 2, 2, 139, 140, 5,
	20, 11, 2, 140, 141, 7, 14, 2, 2, 141, 155, 3, 2, 2, 2, 142, 151, 7, 13,
	2, 2, 143, 148, 5, 22, 12, 2, 144, 145, 7, 11, 2, 2, 145, 147, 5, 22, 12,
	2, 146, 144, 3, 2, 2, 2, 147, 150, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 148,
	149, 3, 2, 2, 2, 149, 152, 3, 2, 2, 2, 150, 148, 3, 2, 2, 2, 151, 143,
	3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 155, 7, 14,
	2, 2, 154, 138, 3, 2, 2, 2, 154, 142, 3, 2, 2, 2, 155, 25, 3, 2, 2, 2,
	22, 32, 37, 46, 50, 54, 57, 61, 65, 74, 90, 100, 103, 109, 118, 121, 126,
	130, 148, 151, 154,
}
var literalNames = []string{
	"", "'\\'", "", "", "", "", "", "", "'}'", "','", "'\"'", "'['", "']'",
	"'='", "", "'{'",
}
var symbolicNames = []string{
	"", "BACKSLASH", "NL", "DOLLARS", "DOLLAR", "GRAVE", "GRAVES", "TEXT",
	"CLOSE_BRACE", "COMMA", "QUOTE", "OPEN_BRACKET", "CLOSE_BRACKET", "EQUALS",
	"SPACE", "OPEN_BRACE", "STR", "NUMBER", "BOOLEAN", "ALPHANUMERIC", "SNL",
	"MDOLLARS", "MDOLLAR", "MTEXT", "CNL", "CGRAVE", "CGRAVES", "CTEXT", "LANGUAGE",
}

var ruleNames = []string{
	"newLine", "blankLines", "document", "block", "content", "language", "text",
	"identifier", "command", "val", "attr", "attrs",
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
	NupParserCOMMA         = 9
	NupParserQUOTE         = 10
	NupParserOPEN_BRACKET  = 11
	NupParserCLOSE_BRACKET = 12
	NupParserEQUALS        = 13
	NupParserSPACE         = 14
	NupParserOPEN_BRACE    = 15
	NupParserSTR           = 16
	NupParserNUMBER        = 17
	NupParserBOOLEAN       = 18
	NupParserALPHANUMERIC  = 19
	NupParserSNL           = 20
	NupParserMDOLLARS      = 21
	NupParserMDOLLAR       = 22
	NupParserMTEXT         = 23
	NupParserCNL           = 24
	NupParserCGRAVE        = 25
	NupParserCGRAVES       = 26
	NupParserCTEXT         = 27
	NupParserLANGUAGE      = 28
)

// NupParser rules.
const (
	NupParserRULE_newLine    = 0
	NupParserRULE_blankLines = 1
	NupParserRULE_document   = 2
	NupParserRULE_block      = 3
	NupParserRULE_content    = 4
	NupParserRULE_language   = 5
	NupParserRULE_text       = 6
	NupParserRULE_identifier = 7
	NupParserRULE_command    = 8
	NupParserRULE_val        = 9
	NupParserRULE_attr       = 10
	NupParserRULE_attrs      = 11
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
		p.SetState(24)
		p.Match(NupParserNL)
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
		p.SetState(26)
		p.NewLine()
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == NupParserNL {
		{
			p.SetState(27)
			p.NewLine()
		}

		p.SetState(30)
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
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NupParserNL {
		{
			p.SetState(32)
			p.NewLine()
		}

		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(38)
		p.Block()
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(39)
				p.BlankLines()
			}
			{
				p.SetState(40)
				p.Block()
			}

		}
		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NupParserNL {
		{
			p.SetState(47)
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
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NupParserBACKSLASH)|(1<<NupParserDOLLARS)|(1<<NupParserDOLLAR)|(1<<NupParserGRAVE)|(1<<NupParserGRAVES)|(1<<NupParserTEXT)|(1<<NupParserMTEXT)|(1<<NupParserCTEXT))) != 0) {
		{
			p.SetState(50)
			p.Content()
		}
		p.SetState(52)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(51)
				p.NewLine()
			}

		}
		p.SetState(55)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(54)
				p.Content()
			}

		}

		p.SetState(59)
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

	p.SetState(63)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NupParserBACKSLASH, NupParserDOLLARS, NupParserDOLLAR, NupParserGRAVE, NupParserGRAVES:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)
			p.Command()
		}

	case NupParserTEXT, NupParserMTEXT, NupParserCTEXT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(62)
			p.Text()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILanguageContext is an interface to support dynamic dispatch.
type ILanguageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLanguageContext differentiates from other interfaces.
	IsLanguageContext()
}

type LanguageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLanguageContext() *LanguageContext {
	var p = new(LanguageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NupParserRULE_language
	return p
}

func (*LanguageContext) IsLanguageContext() {}

func NewLanguageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LanguageContext {
	var p = new(LanguageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NupParserRULE_language

	return p
}

func (s *LanguageContext) GetParser() antlr.Parser { return s.parser }

func (s *LanguageContext) LANGUAGE() antlr.TerminalNode {
	return s.GetToken(NupParserLANGUAGE, 0)
}

func (s *LanguageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LanguageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LanguageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.EnterLanguage(s)
	}
}

func (s *LanguageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NupParserListener); ok {
		listenerT.ExitLanguage(s)
	}
}

func (p *NupParser) Language() (localctx ILanguageContext) {
	localctx = NewLanguageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, NupParserRULE_language)

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
		p.SetState(65)
		p.Match(NupParserLANGUAGE)
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
	p.EnterRule(localctx, 12, NupParserRULE_text)
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
	p.EnterRule(localctx, 14, NupParserRULE_identifier)
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

func (s *CommandContext) MDOLLAR() antlr.TerminalNode {
	return s.GetToken(NupParserMDOLLAR, 0)
}

func (s *CommandContext) DOLLAR() antlr.TerminalNode {
	return s.GetToken(NupParserDOLLAR, 0)
}

func (s *CommandContext) CGRAVE() antlr.TerminalNode {
	return s.GetToken(NupParserCGRAVE, 0)
}

func (s *CommandContext) GRAVE() antlr.TerminalNode {
	return s.GetToken(NupParserGRAVE, 0)
}

func (s *CommandContext) SNL() antlr.TerminalNode {
	return s.GetToken(NupParserSNL, 0)
}

func (s *CommandContext) CGRAVES() antlr.TerminalNode {
	return s.GetToken(NupParserCGRAVES, 0)
}

func (s *CommandContext) GRAVES() antlr.TerminalNode {
	return s.GetToken(NupParserGRAVES, 0)
}

func (s *CommandContext) Attrs() IAttrsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrsContext)
}

func (s *CommandContext) BACKSLASH() antlr.TerminalNode {
	return s.GetToken(NupParserBACKSLASH, 0)
}

func (s *CommandContext) OPEN_BRACE() antlr.TerminalNode {
	return s.GetToken(NupParserOPEN_BRACE, 0)
}

func (s *CommandContext) CLOSE_BRACE() antlr.TerminalNode {
	return s.GetToken(NupParserCLOSE_BRACE, 0)
}

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
	p.EnterRule(localctx, 16, NupParserRULE_command)
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

	p.SetState(128)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NupParserDOLLARS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(74)

			var _m = p.Match(NupParserDOLLARS)

			localctx.(*CommandContext).cmd = _m
		}
		{
			p.SetState(75)
			p.Text()
		}
		{
			p.SetState(76)
			p.Match(NupParserMDOLLARS)
		}

	case NupParserDOLLAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(78)

			var _m = p.Match(NupParserDOLLAR)

			localctx.(*CommandContext).cmd = _m
		}
		{
			p.SetState(79)
			p.Text()
		}
		{
			p.SetState(80)
			p.Match(NupParserMDOLLAR)
		}

	case NupParserGRAVE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(82)

			var _m = p.Match(NupParserGRAVE)

			localctx.(*CommandContext).cmd = _m
		}
		{
			p.SetState(83)
			p.Text()
		}
		{
			p.SetState(84)
			p.Match(NupParserCGRAVE)
		}

	case NupParserGRAVES:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(86)

			var _m = p.Match(NupParserGRAVES)

			localctx.(*CommandContext).cmd = _m
		}
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == NupParserOPEN_BRACKET {
			{
				p.SetState(87)

				var _x = p.Attrs()

				localctx.(*CommandContext).attributes = _x
			}

		}
		{
			p.SetState(90)
			p.Match(NupParserSNL)
		}
		{
			p.SetState(91)
			p.Text()
		}
		{
			p.SetState(92)
			p.Match(NupParserCGRAVES)
		}

	case NupParserBACKSLASH:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(94)
			p.Match(NupParserBACKSLASH)
		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == NupParserALPHANUMERIC {
			{
				p.SetState(95)

				var _m = p.Match(NupParserALPHANUMERIC)

				localctx.(*CommandContext).cmd = _m
			}

			p.SetState(98)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == NupParserOPEN_BRACKET {
			{
				p.SetState(100)

				var _x = p.Attrs()

				localctx.(*CommandContext).attributes = _x
			}

		}
		{
			p.SetState(103)
			p.Match(NupParserOPEN_BRACE)
		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(104)
					p.NewLine()
				}

			}
			p.SetState(109)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NupParserBACKSLASH)|(1<<NupParserDOLLARS)|(1<<NupParserDOLLAR)|(1<<NupParserGRAVE)|(1<<NupParserGRAVES)|(1<<NupParserTEXT)|(1<<NupParserMTEXT)|(1<<NupParserCTEXT))) != 0 {
			{
				p.SetState(110)

				var _x = p.Block()

				localctx.(*CommandContext).inner = _x
			}
			p.SetState(116)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					{
						p.SetState(111)
						p.BlankLines()
					}
					{
						p.SetState(112)

						var _x = p.Block()

						localctx.(*CommandContext).inner = _x
					}

				}
				p.SetState(118)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
			}

		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == NupParserNL {
			{
				p.SetState(121)
				p.NewLine()
			}

			p.SetState(126)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(127)
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
	p.EnterRule(localctx, 18, NupParserRULE_val)
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
		p.SetState(130)
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
	p.EnterRule(localctx, 20, NupParserRULE_attr)

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
		p.SetState(132)

		var _x = p.Identifier()

		localctx.(*AttrContext).name = _x
	}
	{
		p.SetState(133)
		p.Match(NupParserEQUALS)
	}
	{
		p.SetState(134)

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

	// GetValue returns the value rule contexts.
	GetValue() IValContext

	// GetAttribute returns the attribute rule contexts.
	GetAttribute() IAttrContext

	// SetValue sets the value rule contexts.
	SetValue(IValContext)

	// SetAttribute sets the attribute rule contexts.
	SetAttribute(IAttrContext)

	// IsAttrsContext differentiates from other interfaces.
	IsAttrsContext()
}

type AttrsContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	value     IValContext
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

func (s *AttrsContext) GetValue() IValContext { return s.value }

func (s *AttrsContext) GetAttribute() IAttrContext { return s.attribute }

func (s *AttrsContext) SetValue(v IValContext) { s.value = v }

func (s *AttrsContext) SetAttribute(v IAttrContext) { s.attribute = v }

func (s *AttrsContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(NupParserOPEN_BRACKET, 0)
}

func (s *AttrsContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(NupParserCLOSE_BRACKET, 0)
}

func (s *AttrsContext) Val() IValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValContext)
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
	p.EnterRule(localctx, 22, NupParserRULE_attrs)
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

	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(136)
			p.Match(NupParserOPEN_BRACKET)
		}
		{
			p.SetState(137)

			var _x = p.Val()

			localctx.(*AttrsContext).value = _x
		}
		{
			p.SetState(138)
			p.Match(NupParserCLOSE_BRACKET)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(140)
			p.Match(NupParserOPEN_BRACKET)
		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == NupParserALPHANUMERIC {
			{
				p.SetState(141)

				var _x = p.Attr()

				localctx.(*AttrsContext).attribute = _x
			}
			p.SetState(146)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == NupParserCOMMA {
				{
					p.SetState(142)
					p.Match(NupParserCOMMA)
				}
				{
					p.SetState(143)

					var _x = p.Attr()

					localctx.(*AttrsContext).attribute = _x
				}

				p.SetState(148)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(151)
			p.Match(NupParserCLOSE_BRACKET)
		}

	}

	return localctx
}
