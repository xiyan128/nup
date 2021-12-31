// Code generated from parser/NupParser.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // NupParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNupParserListener is a complete listener for a parse tree produced by NupParser.
type BaseNupParserListener struct{}

var _ NupParserListener = &BaseNupParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNupParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNupParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNupParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNupParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNewLine is called when production newLine is entered.
func (s *BaseNupParserListener) EnterNewLine(ctx *NewLineContext) {}

// ExitNewLine is called when production newLine is exited.
func (s *BaseNupParserListener) ExitNewLine(ctx *NewLineContext) {}

// EnterBlankLines is called when production blankLines is entered.
func (s *BaseNupParserListener) EnterBlankLines(ctx *BlankLinesContext) {}

// ExitBlankLines is called when production blankLines is exited.
func (s *BaseNupParserListener) ExitBlankLines(ctx *BlankLinesContext) {}

// EnterDocument is called when production document is entered.
func (s *BaseNupParserListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BaseNupParserListener) ExitDocument(ctx *DocumentContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseNupParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseNupParserListener) ExitBlock(ctx *BlockContext) {}

// EnterContent is called when production content is entered.
func (s *BaseNupParserListener) EnterContent(ctx *ContentContext) {}

// ExitContent is called when production content is exited.
func (s *BaseNupParserListener) ExitContent(ctx *ContentContext) {}

// EnterText is called when production text is entered.
func (s *BaseNupParserListener) EnterText(ctx *TextContext) {}

// ExitText is called when production text is exited.
func (s *BaseNupParserListener) ExitText(ctx *TextContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseNupParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseNupParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterOpenBracket is called when production openBracket is entered.
func (s *BaseNupParserListener) EnterOpenBracket(ctx *OpenBracketContext) {}

// ExitOpenBracket is called when production openBracket is exited.
func (s *BaseNupParserListener) ExitOpenBracket(ctx *OpenBracketContext) {}

// EnterOpenBrace is called when production openBrace is entered.
func (s *BaseNupParserListener) EnterOpenBrace(ctx *OpenBraceContext) {}

// ExitOpenBrace is called when production openBrace is exited.
func (s *BaseNupParserListener) ExitOpenBrace(ctx *OpenBraceContext) {}

// EnterCommand is called when production command is entered.
func (s *BaseNupParserListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BaseNupParserListener) ExitCommand(ctx *CommandContext) {}

// EnterVal is called when production val is entered.
func (s *BaseNupParserListener) EnterVal(ctx *ValContext) {}

// ExitVal is called when production val is exited.
func (s *BaseNupParserListener) ExitVal(ctx *ValContext) {}

// EnterAttr is called when production attr is entered.
func (s *BaseNupParserListener) EnterAttr(ctx *AttrContext) {}

// ExitAttr is called when production attr is exited.
func (s *BaseNupParserListener) ExitAttr(ctx *AttrContext) {}

// EnterAttrs is called when production attrs is entered.
func (s *BaseNupParserListener) EnterAttrs(ctx *AttrsContext) {}

// ExitAttrs is called when production attrs is exited.
func (s *BaseNupParserListener) ExitAttrs(ctx *AttrsContext) {}
