package translator

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

type refCounterEntry struct {
	count int
	match bool
}

type refCounter struct {
	numRefs int
	m       map[string]*refCounterEntry
}

func NewRefCounter() *refCounter {
	return &refCounter{
		numRefs: 0,
		m:       make(map[string]*refCounterEntry),
	}
}

func (r *refCounter) addRef(id string) int {
	if r.m[id] == nil {
		r.numRefs++
		r.m[id] = &refCounterEntry{r.numRefs, false}
	}
	return r.m[id].count
}

func (r *refCounter) addDeref(id string) (int, error) {
	if r.m[id] == nil {
		return -1, fmt.Errorf("reference \"%s\" was never defined", id)
	}
	r.m[id].match = true
	return r.m[id].count, nil
}

func (r *refCounter) allMatched() error {
	for id, e := range r.m {
		if !e.match {
			return fmt.Errorf("unmatched reference \"%s\" ", id)
		}
	}
	return nil
}

type Translator struct {
	*parser.BaseNupParserListener
	env            *environment
	documentWriter DocumentWriter
	writer         io.Writer
	refCounter     *refCounter
	idPool         map[string]bool
}

func NewNupListener(writer io.Writer) *Translator {
	return &Translator{
		documentWriter: NewHtmlWriter(),
		env:            &environment{vars: make(map[string]interface{})},
		writer:         writer,
		refCounter:     NewRefCounter(),
		idPool:         make(map[string]bool),
	}
}

// GetActiveCommands climbs up the environment tree to find the active commands
func (s *Translator) GetActiveCommands() []string {
	var commands []string
	for env := s.env; env != nil; env = env.parent {
		commands = append(commands, env.command)
	}
	return commands
}

func (s *Translator) GetCurrentCommand() string {
	return s.env.command
}

func (s *Translator) getAttribute(name string) (interface{}, bool) {
	val, ok := s.env.vars[name]
	return val, ok
}

func (s *Translator) setAttribute(name string, value interface{}) {
	s.env.vars[name] = value
}

func (s *Translator) checkDeref() {
	if s.GetCurrentCommand() == Reference {
		if id, ok := s.getAttribute("to"); ok {
			count := s.refCounter.addRef(id.(string))
			s.setAttribute("refNum", count)
		} else {
			panic("Field \"to\" is required")
		}
	}
}

func (s *Translator) checkRef() {
	if s.GetCurrentCommand() == Dereference {
		if id, ok := s.getAttribute("id"); ok {
			count, err := s.refCounter.addDeref(id.(string))
			if err != nil {
				panic(err)
			}
			s.setAttribute("refNum", count)
		} else {
			panic("deref must have an id")
		}
	}
}

func (s *Translator) checkDuplicateId() {
	if val, ok := s.getAttribute("id"); ok {
		if s.idPool[val.(string)] {
			panic("duplicate id \"" + val.(string) + "\" not allowed")
		}
		s.idPool[val.(string)] = true
	}
}

// VisitTerminal is called when a terminal node is visited.
func (s *Translator) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *Translator) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *Translator) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *Translator) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNewLine is called when production newLine is entered.
func (s *Translator) EnterNewLine(ctx *parser.NewLineContext) {}

// ExitNewLine is called when production newLine is exited.
func (s *Translator) ExitNewLine(ctx *parser.NewLineContext) {}

// EnterBlankLines is called when production blankLines is entered.
func (s *Translator) EnterBlankLines(ctx *parser.BlankLinesContext) {}

// ExitBlankLines is called when production blankLines is exited.
func (s *Translator) ExitBlankLines(ctx *parser.BlankLinesContext) {}

// EnterDocument is called when production document is entered.
func (s *Translator) EnterDocument(ctx *parser.DocumentContext) {
	_, err := s.documentWriter.WritePreamble()
	if err != nil {
		return
	}
}

// ExitDocument is called when production document is exited.
func (s *Translator) ExitDocument(ctx *parser.DocumentContext) {
	defer func(documentWriter DocumentWriter, writer io.Writer) {
		_, err := documentWriter.Flush(writer)
		if err != nil {
			panic(err)
		}
	}(s.documentWriter, s.writer)

	err := s.refCounter.allMatched()
	if err != nil {
		panic(err)
	}

	_, err = s.documentWriter.WritePostamble()
	if err != nil {
		return
	}
}

// EnterBlock is called when production block is entered.
func (s *Translator) EnterBlock(ctx *parser.BlockContext) {

	// if the outer command is a block, or it is the outermost command, the inner command is implicitly a paragraph
	cmd := s.GetCurrentCommand()
	if (IsBlock(cmd) && cmd != Paragraph) || (cmd == "" && s.env.parent == nil) {
		s.pushEnv(map[string]interface{}{"_implicit_para": true}, Paragraph)
	}
	if s.env.vars["_implicit_para"] != nil {
		s.writeHTMLOpenTag()
	}
}

// ExitBlock is called when production block is exited.
func (s *Translator) ExitBlock(ctx *parser.BlockContext) {
	if s.env.vars["_implicit_para"] != nil {
		//s.env = s.env.parent
		s.writeHTMLCloseTag()
		s.popEnv()
	}
}

// EnterContent is called when production content is entered.
func (s *Translator) EnterContent(ctx *parser.ContentContext) {}

// ExitContent is called when production content is exited.
func (s *Translator) ExitContent(ctx *parser.ContentContext) {}

// EnterText is called when production text is entered.
func (s *Translator) EnterText(ctx *parser.TextContext) {
	_, err := s.documentWriter.WriteString(ctx.GetText())
	if err != nil {
		return
	}
}

// ExitText is called when production text is exited.
func (s *Translator) ExitText(ctx *parser.TextContext) {
}

// EnterIdentifier is called when production identifier is entered.
func (s *Translator) EnterIdentifier(ctx *parser.IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *Translator) ExitIdentifier(ctx *parser.IdentifierContext) {}

// EnterCommand is called when production command is entered.
func (s *Translator) EnterCommand(ctx *parser.CommandContext) {

	// parse and check the command
	cmd := ctx.GetCmd().GetText()
	validAttrs, ok := GetAttrTypes(cmd)
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
						parsed := parseAttrValue(attr.GetValue().GetText())
						switch reflect.TypeOf(parsed) {
						case attrType:
							attrsMap[attr.GetName().GetText()] = parsed
						default:
							panic("Wrong type: " + ctx.GetText() + " at " + strconv.FormatInt(int64(ctx.GetStart().GetLine()), 10))
						}
					} else {
						// TODO: better error handling should be implemented
						panic("Wrong attribute")
					}
				}
			}
		}
	}
	s.pushEnv(attrsMap, cmd)
	s.checkDuplicateId()
	s.writeHTMLOpenTag()
}

func parseAttrValue(s string) interface{} {
	value := s
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
	return parsed
}

func (s *Translator) pushEnv(attrsMap map[string]interface{}, cmd string) {
	s.env = &environment{parent: s.env, vars: attrsMap, command: cmd}
}

// ExitCommand is called when production command is exited.
func (s *Translator) ExitCommand(ctx *parser.CommandContext) {
	s.writeHTMLCloseTag()
	// pop the environment
	s.popEnv()
}

func (s *Translator) popEnv() {
	s.env = s.env.parent
}

func (s *Translator) writeHTMLTag(tag bool) {

	s.checkRef()
	s.checkDeref()

	cmd := s.GetCurrentCommand()
	openTag, closeTag := GetHtmlTags(cmd, s.env.vars)
	if tag {
		_, _ = s.documentWriter.WriteString(openTag)
	} else {

		_, _ = s.documentWriter.WriteString(closeTag)
	}
}

func (s *Translator) writeHTMLOpenTag() {
	s.writeHTMLTag(true)
}

func (s *Translator) writeHTMLCloseTag() {
	s.writeHTMLTag(false)
}

// EnterVal is called when production val is entered.
func (s *Translator) EnterVal(ctx *parser.ValContext) {}

// ExitVal is called when production val is exited.
func (s *Translator) ExitVal(ctx *parser.ValContext) {}

// EnterAttr is called when production attr is entered.
func (s *Translator) EnterAttr(ctx *parser.AttrContext) {}

// ExitAttr is called when production attr is exited.
func (s *Translator) ExitAttr(ctx *parser.AttrContext) {}

// EnterAttrs is called when production attrsTypes is entered.
func (s *Translator) EnterAttrs(ctx *parser.AttrsContext) {}

// ExitAttrs is called when production attrsTypes is exited.
func (s *Translator) ExitAttrs(ctx *parser.AttrsContext) {}
