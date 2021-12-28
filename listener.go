package main

import (
	"crypto/sha1"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"hash"
	"io"
	"reflect"
	"strconv"
	"xiyan.life/nup/parser"
)

type IdGenerator struct {
	occupancy map[string]bool
	hasher    hash.Hash
}

func (g *IdGenerator) Issue(s string) (id string, ok bool) {
	// compute the sha-1 hash of the string
	// and use the first 8 bytes as the id
	h := sha1.New()
	h.Write([]byte(s))
	b := h.Sum(nil)
	id = strconv.FormatInt(int64(b[0])<<24|int64(b[1])<<16|int64(b[2])<<8|int64(b[3]), 16)
	if g.occupancy[id] {
		return "", false
	}
	g.occupancy[id] = true
	return id, true
}

type environment struct {
	parent  *environment
	command string
	vars    map[string]interface{}
}

type NupListener struct {
	*parser.BaseNupParserListener
	env            *environment
	documentWriter DocumentWriter
	writer         io.Writer
}

func NewNupListener(writer io.Writer) *NupListener {
	return &NupListener{
		documentWriter: NewHtmlWriter(),
		env:            &environment{vars: make(map[string]interface{})},
		writer:         writer,
	}
}

// GetActiveCommands climbs up the environment tree to find the active commands
func (s *NupListener) GetActiveCommands() []string {
	var commands []string
	for env := s.env; env != nil; env = env.parent {
		commands = append(commands, env.command)
	}
	return commands
}

func (s *NupListener) GetCurrentCommand() string {
	return s.env.command
}

// VisitTerminal is called when a terminal node is visited.
func (s *NupListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *NupListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *NupListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *NupListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNewLine is called when production newLine is entered.
func (s *NupListener) EnterNewLine(ctx *parser.NewLineContext) {}

// ExitNewLine is called when production newLine is exited.
func (s *NupListener) ExitNewLine(ctx *parser.NewLineContext) {}

// EnterBlankLines is called when production blankLines is entered.
func (s *NupListener) EnterBlankLines(ctx *parser.BlankLinesContext) {}

// ExitBlankLines is called when production blankLines is exited.
func (s *NupListener) ExitBlankLines(ctx *parser.BlankLinesContext) {}

// EnterDocument is called when production document is entered.
func (s *NupListener) EnterDocument(ctx *parser.DocumentContext) {
	_, err := s.documentWriter.WritePreamble()
	if err != nil {
		return
	}
}

// ExitDocument is called when production document is exited.
func (s *NupListener) ExitDocument(ctx *parser.DocumentContext) {
	defer func(documentWriter DocumentWriter, writer io.Writer) {
		_, err := documentWriter.Flush(writer)
		if err != nil {

		}
	}(s.documentWriter, s.writer)

	_, err := s.documentWriter.WritePostamble()
	if err != nil {
		return
	}
}

// EnterBlock is called when production block is entered.
func (s *NupListener) EnterBlock(ctx *parser.BlockContext) {
	// if the outer command is a block, or it is the outermost command, the inner command is by default a paragraph
	if s.GetCurrentCommand() == "" && ((s.env.parent != nil && IsBlock(s.env.parent.command)) || s.env.parent == nil) {
		s.env.command = "para"
	}

	// if a block is not inside a command, it is a paragraph block
	cmd := s.GetCurrentCommand()
	open, _ := GetHtmlTags(cmd, s.env.vars)
	_, err := s.documentWriter.WriteString(open)
	if err != nil {
		return
	}
}

// ExitBlock is called when production block is exited.
func (s *NupListener) ExitBlock(ctx *parser.BlockContext) {
	cmd := s.GetCurrentCommand()
	_, closeTag := GetHtmlTags(cmd, s.env.vars)
	_, err := s.documentWriter.WriteString(closeTag)
	if err != nil {
		return
	}
}

// EnterContent is called when production content is entered.
func (s *NupListener) EnterContent(ctx *parser.ContentContext) {}

// ExitContent is called when production content is exited.
func (s *NupListener) ExitContent(ctx *parser.ContentContext) {}

// EnterText is called when production text is entered.
func (s *NupListener) EnterText(ctx *parser.TextContext) {
	cmd := s.GetCurrentCommand()
	if IsMath(cmd) {
		openTag, _ := GetHtmlTags(cmd, s.env.vars)
		_, err := s.documentWriter.WriteString(openTag)
		if err != nil {
			return
		}
	}
	_, err := s.documentWriter.WriteString(ctx.GetText())
	if err != nil {
		return
	}
}

// ExitText is called when production text is exited.
func (s *NupListener) ExitText(ctx *parser.TextContext) {
	cmd := s.GetCurrentCommand()
	if IsMath(cmd) {
		_, closeTag := GetHtmlTags(cmd, s.env.vars)
		_, err := s.documentWriter.WriteString(closeTag)
		if err != nil {
			return
		}
	}
}

// EnterIdentifier is called when production identifier is entered.
func (s *NupListener) EnterIdentifier(ctx *parser.IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *NupListener) ExitIdentifier(ctx *parser.IdentifierContext) {}

// EnterCommand is called when production command is entered.
func (s *NupListener) EnterCommand(ctx *parser.CommandContext) {

	// parse and check the command
	cmd := ctx.GetCmd().GetText()
	println(cmd)
	validAttrs, ok := Attrs(cmd)
	if !ok {
		panic("Wrong command")
	}

	// check if a block command is nested inside an inline command
	if s.env.parent != nil && !IsBlock(s.env.parent.command) && IsBlock(cmd) {
		panic("Cannot have block command inside inline command")
	}

	// parse attributes
	attrsMap := make(map[string]interface{})
	if attrs := ctx.GetAttributes(); attrs != nil {
		for i := 0; i < attrs.GetChildCount(); i++ {
			if attr, ok := attrs.GetChild(i).(*parser.AttrContext); ok {
				if attr.GetName() != nil {
					if attrType, ok := validAttrs[attr.GetName().GetText()]; ok {
						value := attr.GetValue().GetText()
						var parsed interface{}
						// parse value into int, float, bool, or string
						if value == "true" || value == "false" {
							parsed = value == "true"
						} else if i, err := strconv.Atoi(value); err == nil {
							parsed = i
						} else if f, err := strconv.ParseFloat(value, 64); err == nil {
							parsed = f
						} else if str, err := strconv.Unquote(value); err == nil {
							parsed = str
						} else {
							parsed = nil
						}
						switch reflect.TypeOf(parsed) {
						case attrType:
							attrsMap[attr.GetName().GetText()] = parsed
						default:
							panic("Wrong type: " + ctx.GetText() + " at " + strconv.FormatInt(int64(ctx.GetStart().GetLine()), 10))
						}
					} else {
						panic("Wrong attribute")
					}
				}
			}
		}
	}

	s.env = &environment{parent: s.env, vars: attrsMap, command: cmd}
	fmt.Printf("%v", s.env)
}

// ExitCommand is called when production command is exited.
func (s *NupListener) ExitCommand(ctx *parser.CommandContext) {
	s.env = s.env.parent
}

// EnterVal is called when production val is entered.
func (s *NupListener) EnterVal(ctx *parser.ValContext) {}

// ExitVal is called when production val is exited.
func (s *NupListener) ExitVal(ctx *parser.ValContext) {}

// EnterAttr is called when production attr is entered.
func (s *NupListener) EnterAttr(ctx *parser.AttrContext) {}

// ExitAttr is called when production attr is exited.
func (s *NupListener) ExitAttr(ctx *parser.AttrContext) {}

// EnterAttrs is called when production attrsTypes is entered.
func (s *NupListener) EnterAttrs(ctx *parser.AttrsContext) {}

// ExitAttrs is called when production attrsTypes is exited.
func (s *NupListener) ExitAttrs(ctx *parser.AttrsContext) {}
