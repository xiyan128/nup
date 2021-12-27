package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"reflect"
	"strconv"
	"xiyan.life/nup/parser"
)

type environment struct {
	parent  *environment
	configs map[string]interface{}
}

type NupListener struct {
	*parser.BaseNupParserListener
	env *environment
}

func NewNupListener() *NupListener {
	return &NupListener{
		env: &environment{configs: make(map[string]interface{})},
	}
}

// GetActiveCommands climbs up the environment tree to find the active commands
func (s *NupListener) GetActiveCommands() []string {
	var commands []string
	for env := s.env; env != nil; env = env.parent {
		if env.configs["command"] != nil {
			commands = append(commands, env.configs["command"].(string))
		}
	}
	return commands
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
func (s *NupListener) EnterDocument(ctx *parser.DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *NupListener) ExitDocument(ctx *parser.DocumentContext) {}

// EnterBlock is called when production block is entered.
func (s *NupListener) EnterBlock(ctx *parser.BlockContext) {

}

// ExitBlock is called when production block is exited.
func (s *NupListener) ExitBlock(ctx *parser.BlockContext) {
}

// EnterContent is called when production content is entered.
func (s *NupListener) EnterContent(ctx *parser.ContentContext) {}

// ExitContent is called when production content is exited.
func (s *NupListener) ExitContent(ctx *parser.ContentContext) {}

// EnterText is called when production text is entered.
func (s *NupListener) EnterText(ctx *parser.TextContext) {
	fmt.Printf("(%v)[%s]", s.GetActiveCommands(), ctx.GetText())
}

// ExitText is called when production text is exited.
func (s *NupListener) ExitText(ctx *parser.TextContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *NupListener) EnterIdentifier(ctx *parser.IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *NupListener) ExitIdentifier(ctx *parser.IdentifierContext) {}

// EnterCommand is called when production command is entered.
func (s *NupListener) EnterCommand(ctx *parser.CommandContext) {

	attrsMap := make(map[string]interface{})
	//cmd := ctx.GetCmd()

	cmd := ctx.GetCmd().GetText()
	fmt.Printf("%v", s.GetActiveCommands())
	validAttrs, ok := Attrs(cmd)
	if !ok {
		panic("Wrong command")
	}
	attrsMap["command"] = cmd

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
							panic("Wrong type")
						}
					} else {
						panic("Wrong attribute")
					}
				}
			}
		}
	}

	s.env = &environment{parent: s.env, configs: attrsMap}
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

// EnterAttrs is called when production attrs is entered.
func (s *NupListener) EnterAttrs(ctx *parser.AttrsContext) {}

// ExitAttrs is called when production attrs is exited.
func (s *NupListener) ExitAttrs(ctx *parser.AttrsContext) {}
