// Code generated from parser/NupParser.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // NupParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NupParserListener is a complete listener for a parse tree produced by NupParser.
type NupParserListener interface {
	antlr.ParseTreeListener

	// EnterNewLine is called when entering the newLine production.
	EnterNewLine(c *NewLineContext)

	// EnterBlankLines is called when entering the blankLines production.
	EnterBlankLines(c *BlankLinesContext)

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterContent is called when entering the content production.
	EnterContent(c *ContentContext)

	// EnterText is called when entering the text production.
	EnterText(c *TextContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterVal is called when entering the val production.
	EnterVal(c *ValContext)

	// EnterAttr is called when entering the attr production.
	EnterAttr(c *AttrContext)

	// EnterAttrs is called when entering the attrs production.
	EnterAttrs(c *AttrsContext)

	// ExitNewLine is called when exiting the newLine production.
	ExitNewLine(c *NewLineContext)

	// ExitBlankLines is called when exiting the blankLines production.
	ExitBlankLines(c *BlankLinesContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitContent is called when exiting the content production.
	ExitContent(c *ContentContext)

	// ExitText is called when exiting the text production.
	ExitText(c *TextContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitVal is called when exiting the val production.
	ExitVal(c *ValContext)

	// ExitAttr is called when exiting the attr production.
	ExitAttr(c *AttrContext)

	// ExitAttrs is called when exiting the attrs production.
	ExitAttrs(c *AttrsContext)
}
