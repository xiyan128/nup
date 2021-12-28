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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 22, 120,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2,
	3, 2, 3, 3, 3, 3, 6, 3, 29, 10, 3, 13, 3, 14, 3, 30, 3, 4, 7, 4, 34, 10,
	4, 12, 4, 14, 4, 37, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 43, 10, 4, 12,
	4, 14, 4, 46, 11, 4, 3, 4, 5, 4, 49, 10, 4, 3, 5, 3, 5, 5, 5, 53, 10, 5,
	3, 5, 5, 5, 56, 10, 5, 6, 5, 58, 10, 5, 13, 5, 14, 5, 59, 3, 6, 3, 6, 5,
	6, 64, 10, 6, 3, 7, 3, 7, 3, 8, 6, 8, 69, 10, 8, 13, 8, 14, 8, 70, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 6, 9, 83, 10, 9,
	13, 9, 14, 9, 84, 3, 9, 5, 9, 88, 10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9,
	94, 10, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 110, 10, 12, 12, 12, 14, 12,
	113, 11, 12, 5, 12, 115, 10, 12, 3, 12, 5, 12, 118, 10, 12, 3, 12, 2, 2,
	13, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 2, 4, 4, 2, 7, 7, 22, 22, 3,
	2, 16, 18, 2, 124, 2, 24, 3, 2, 2, 2, 4, 26, 3, 2, 2, 2, 6, 35, 3, 2, 2,
	2, 8, 57, 3, 2, 2, 2, 10, 63, 3, 2, 2, 2, 12, 65, 3, 2, 2, 2, 14, 68, 3,
	2, 2, 2, 16, 93, 3, 2, 2, 2, 18, 95, 3, 2, 2, 2, 20, 97, 3, 2, 2, 2, 22,
	117, 3, 2, 2, 2, 24, 25, 7, 4, 2, 2, 25, 3, 3, 2, 2, 2, 26, 28, 5, 2, 2,
	2, 27, 29, 5, 2, 2, 2, 28, 27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 28,
	3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 5, 3, 2, 2, 2, 32, 34, 5, 2, 2, 2,
	33, 32, 3, 2, 2, 2, 34, 37, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 35, 36, 3,
	2, 2, 2, 36, 38, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 38, 44, 5, 8, 5, 2, 39,
	40, 5, 4, 3, 2, 40, 41, 5, 8, 5, 2, 41, 43, 3, 2, 2, 2, 42, 39, 3, 2, 2,
	2, 43, 46, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 48,
	3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 47, 49, 5, 4, 3, 2, 48, 47, 3, 2, 2, 2,
	48, 49, 3, 2, 2, 2, 49, 7, 3, 2, 2, 2, 50, 52, 5, 10, 6, 2, 51, 53, 5,
	2, 2, 2, 52, 51, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 55, 3, 2, 2, 2, 54,
	56, 5, 10, 6, 2, 55, 54, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 58, 3, 2,
	2, 2, 57, 50, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60,
	3, 2, 2, 2, 60, 9, 3, 2, 2, 2, 61, 64, 5, 16, 9, 2, 62, 64, 5, 12, 7, 2,
	63, 61, 3, 2, 2, 2, 63, 62, 3, 2, 2, 2, 64, 11, 3, 2, 2, 2, 65, 66, 9,
	2, 2, 2, 66, 13, 3, 2, 2, 2, 67, 69, 7, 19, 2, 2, 68, 67, 3, 2, 2, 2, 69,
	70, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 15, 3, 2, 2,
	2, 72, 73, 7, 6, 2, 2, 73, 74, 5, 12, 7, 2, 74, 75, 7, 21, 2, 2, 75, 94,
	3, 2, 2, 2, 76, 77, 7, 5, 2, 2, 77, 78, 5, 12, 7, 2, 78, 79, 7, 20, 2,
	2, 79, 94, 3, 2, 2, 2, 80, 82, 7, 3, 2, 2, 81, 83, 7, 19, 2, 2, 82, 81,
	3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2,
	85, 87, 3, 2, 2, 2, 86, 88, 5, 22, 12, 2, 87, 86, 3, 2, 2, 2, 87, 88, 3,
	2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 90, 7, 15, 2, 2, 90, 91, 5, 8, 5, 2, 91,
	92, 7, 8, 2, 2, 92, 94, 3, 2, 2, 2, 93, 72, 3, 2, 2, 2, 93, 76, 3, 2, 2,
	2, 93, 80, 3, 2, 2, 2, 94, 17, 3, 2, 2, 2, 95, 96, 9, 3, 2, 2, 96, 19,
	3, 2, 2, 2, 97, 98, 5, 14, 8, 2, 98, 99, 7, 13, 2, 2, 99, 100, 5, 18, 10,
	2, 100, 21, 3, 2, 2, 2, 101, 102, 7, 11, 2, 2, 102, 103, 5, 18, 10, 2,
	103, 104, 7, 12, 2, 2, 104, 118, 3, 2, 2, 2, 105, 114, 7, 11, 2, 2, 106,
	111, 5, 20, 11, 2, 107, 108, 7, 9, 2, 2, 108, 110, 5, 20, 11, 2, 109, 107,
	3, 2, 2, 2, 110, 113, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2, 111, 112, 3, 2,
	2, 2, 112, 115, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 114, 106, 3, 2, 2, 2,
	114, 115, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 118, 7, 12, 2, 2, 117,
	101, 3, 2, 2, 2, 117, 105, 3, 2, 2, 2, 118, 23, 3, 2, 2, 2, 17, 30, 35,
	44, 48, 52, 55, 59, 63, 70, 84, 87, 93, 111, 114, 117,
}
var literalNames = []string{
	"", "'\\'", "", "", "", "", "'}'", "','", "'\"'", "'['", "']'", "'='",
	"", "'{'",
}
var symbolicNames = []string{
	"", "BACKSLASH", "NL", "DOLLAR", "DOLLARS", "TEXT", "CLOSE_BRACE", "COMMA",
	"QUOTE", "OPEN_BRACKET", "CLOSE_BRACKET", "EQUALS", "SPACE", "OPEN_BRACE",
	"STR", "NUMBER", "BOOLEAN", "ALPHANUMERIC", "MDOLLAR", "MDOLLARS", "MTEXT",
}

var ruleNames = []string{
	"newLine", "blankLines", "document", "block", "content", "text", "identifier",
	"command", "val", "attr", "attrs",
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
	NupParserDOLLAR        = 3
	NupParserDOLLARS       = 4
	NupParserTEXT          = 5
	NupParserCLOSE_BRACE   = 6
	NupParserCOMMA         = 7
	NupParserQUOTE         = 8
	NupParserOPEN_BRACKET  = 9
	NupParserCLOSE_BRACKET = 10
	NupParserEQUALS        = 11
	NupParserSPACE         = 12
	NupParserOPEN_BRACE    = 13
	NupParserSTR           = 14
	NupParserNUMBER        = 15
	NupParserBOOLEAN       = 16
	NupParserALPHANUMERIC  = 17
	NupParserMDOLLAR       = 18
	NupParserMDOLLARS      = 19
	NupParserMTEXT         = 20
)

// NupParser rules.
const (
	NupParserRULE_newLine    = 0
	NupParserRULE_blankLines = 1
	NupParserRULE_document   = 2
	NupParserRULE_block      = 3
	NupParserRULE_content    = 4
	NupParserRULE_text       = 5
	NupParserRULE_identifier = 6
	NupParserRULE_command    = 7
	NupParserRULE_val        = 8
	NupParserRULE_attr       = 9
	NupParserRULE_attrs      = 10
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
		p.SetState(22)
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
		p.SetState(24)
		p.NewLine()
	}
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == NupParserNL {
		{
			p.SetState(25)
			p.NewLine()
		}

		p.SetState(28)
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
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NupParserNL {
		{
			p.SetState(30)
			p.NewLine()
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(36)
		p.Block()
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(37)
				p.BlankLines()
			}
			{
				p.SetState(38)
				p.Block()
			}

		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NupParserNL {
		{
			p.SetState(45)
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
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NupParserBACKSLASH)|(1<<NupParserDOLLAR)|(1<<NupParserDOLLARS)|(1<<NupParserTEXT)|(1<<NupParserMTEXT))) != 0) {
		{
			p.SetState(48)
			p.Content()
		}
		p.SetState(50)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(49)
				p.NewLine()
			}

		}
		p.SetState(53)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(52)
				p.Content()
			}

		}

		p.SetState(57)
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

	p.SetState(61)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NupParserBACKSLASH, NupParserDOLLAR, NupParserDOLLARS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.Command()
		}

	case NupParserTEXT, NupParserMTEXT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)
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
		p.SetState(63)
		_la = p.GetTokenStream().LA(1)

		if !(_la == NupParserTEXT || _la == NupParserMTEXT) {
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
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == NupParserALPHANUMERIC {
		{
			p.SetState(65)
			p.Match(NupParserALPHANUMERIC)
		}

		p.SetState(68)
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

func (s *CommandContext) BACKSLASH() antlr.TerminalNode {
	return s.GetToken(NupParserBACKSLASH, 0)
}

func (s *CommandContext) OPEN_BRACE() antlr.TerminalNode {
	return s.GetToken(NupParserOPEN_BRACE, 0)
}

func (s *CommandContext) CLOSE_BRACE() antlr.TerminalNode {
	return s.GetToken(NupParserCLOSE_BRACE, 0)
}

func (s *CommandContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *CommandContext) AllALPHANUMERIC() []antlr.TerminalNode {
	return s.GetTokens(NupParserALPHANUMERIC)
}

func (s *CommandContext) ALPHANUMERIC(i int) antlr.TerminalNode {
	return s.GetToken(NupParserALPHANUMERIC, i)
}

func (s *CommandContext) Attrs() IAttrsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrsContext)
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
	p.EnterRule(localctx, 14, NupParserRULE_command)
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

	p.SetState(91)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NupParserDOLLARS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)

			var _m = p.Match(NupParserDOLLARS)

			localctx.(*CommandContext).cmd = _m
		}
		{
			p.SetState(71)
			p.Text()
		}
		{
			p.SetState(72)
			p.Match(NupParserMDOLLARS)
		}

	case NupParserDOLLAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(74)

			var _m = p.Match(NupParserDOLLAR)

			localctx.(*CommandContext).cmd = _m
		}
		{
			p.SetState(75)
			p.Text()
		}
		{
			p.SetState(76)
			p.Match(NupParserMDOLLAR)
		}

	case NupParserBACKSLASH:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(78)
			p.Match(NupParserBACKSLASH)
		}
		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == NupParserALPHANUMERIC {
			{
				p.SetState(79)

				var _m = p.Match(NupParserALPHANUMERIC)

				localctx.(*CommandContext).cmd = _m
			}

			p.SetState(82)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == NupParserOPEN_BRACKET {
			{
				p.SetState(84)

				var _x = p.Attrs()

				localctx.(*CommandContext).attributes = _x
			}

		}
		{
			p.SetState(87)
			p.Match(NupParserOPEN_BRACE)
		}
		{
			p.SetState(88)

			var _x = p.Block()

			localctx.(*CommandContext).inner = _x
		}
		{
			p.SetState(89)
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
	p.EnterRule(localctx, 16, NupParserRULE_val)
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
		p.SetState(93)
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
	p.EnterRule(localctx, 18, NupParserRULE_attr)

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
		p.SetState(95)

		var _x = p.Identifier()

		localctx.(*AttrContext).name = _x
	}
	{
		p.SetState(96)
		p.Match(NupParserEQUALS)
	}
	{
		p.SetState(97)

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
	p.EnterRule(localctx, 20, NupParserRULE_attrs)
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

	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(99)
			p.Match(NupParserOPEN_BRACKET)
		}
		{
			p.SetState(100)

			var _x = p.Val()

			localctx.(*AttrsContext).value = _x
		}
		{
			p.SetState(101)
			p.Match(NupParserCLOSE_BRACKET)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.Match(NupParserOPEN_BRACKET)
		}
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == NupParserALPHANUMERIC {
			{
				p.SetState(104)

				var _x = p.Attr()

				localctx.(*AttrsContext).attribute = _x
			}
			p.SetState(109)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == NupParserCOMMA {
				{
					p.SetState(105)
					p.Match(NupParserCOMMA)
				}
				{
					p.SetState(106)

					var _x = p.Attr()

					localctx.(*AttrsContext).attribute = _x
				}

				p.SetState(111)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(114)
			p.Match(NupParserCLOSE_BRACKET)
		}

	}

	return localctx
}
