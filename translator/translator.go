package translator

import (
	"crypto/sha1"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"hash"
	"html"
	"io"
	"reflect"
	"strconv"
	"strings"
	"unicode"
	"xiyan.life/nup/errhandler"
	"xiyan.life/nup/parser"
	"xiyan.life/nup/translator/commands"
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

type Translator struct {
	*parser.BaseNupParserListener
	env            *environment
	documentWriter DocumentWriter
	writer         io.Writer
	idPool         map[string]bool
	errorHandler   *errhandler.NupErrorListener
	currCtx        antlr.ParserRuleContext
	refCounter     *refCounter
	sectionCounter *sectionCounter
}

func NewNupListener(writer io.Writer, errorHandler *errhandler.NupErrorListener) *Translator {
	return &Translator{
		documentWriter: NewHtmlWriter(),
		env:            &environment{command: commands.Document, vars: make(map[string]interface{})},
		writer:         writer,
		refCounter:     NewRefCounter(),
		idPool:         make(map[string]bool),
		errorHandler:   errorHandler,
		sectionCounter: NewSectionCounter(),
	}
}

func (t *Translator) reportRuntimeError(msg string) {
	t.errorHandler.RuntimeError(t.currCtx.GetStart().GetLine(), t.currCtx.GetStart().GetColumn(), msg)
}

// GetActiveCommands climbs up the environment tree to find the active commands
func (t *Translator) GetActiveCommands() []string {
	var commands []string
	for env := t.env; env != nil; env = env.parent {
		commands = append(commands, env.command)
	}
	return commands
}

func (t *Translator) GetCurrentCommand() string {
	return t.env.command
}

func (t *Translator) getAttribute(name string) (interface{}, bool) {
	val, ok := t.env.vars[name]
	return val, ok
}

func (t *Translator) setAttribute(name string, value interface{}) {
	t.env.vars[name] = value
}

func (t *Translator) checkRef() {
	if t.GetCurrentCommand() == commands.Reference {
		if id, ok := t.getAttribute("to"); ok {
			count := t.refCounter.addRef(id.(string))
			t.setAttribute("refNum", count)
		} else {
			t.reportRuntimeError("field \"to\" is required")
		}
	}
}

func (t *Translator) checkDeref() {
	if t.GetCurrentCommand() == commands.Dereference {
		if id, ok := t.getAttribute("id"); ok {
			count, err := t.refCounter.addDeref(id.(string))
			if err != nil {
				t.reportRuntimeError(err.Error())
			}
			t.setAttribute("refNum", count)
		} else {
			t.reportRuntimeError("deref must have an id")
		}
	}
}

func (t *Translator) checkDuplicateId() {
	if val, ok := t.getAttribute("id"); ok {
		if t.idPool[val.(string)] {
			t.reportRuntimeError("duplicate id \"" + val.(string) + "\" not allowed")
		}
		t.idPool[val.(string)] = true
	}
}

func (t *Translator) checkBlockCommandConstraint() {
	if cmd := t.GetCurrentCommand(); commands.IsBlock(cmd) {
		if val, ok := t.currCtx.
			GetParent().GetParent().(*parser.BlockContext).
			GetChild(0).(*parser.ContentContext).
			GetChild(0).(*parser.CommandContext); !(ok && val == t.currCtx) {
			t.reportRuntimeError(fmt.Sprintf("block command \"%s\" must be the first element of a block", cmd))
		}
	}
}

func (t *Translator) checkCaptioned() {
	if cmd := t.GetCurrentCommand(); commands.IsCaptioned(cmd) {
		// by default, captioned commands are numbered
		t.setAttribute("_caption_type", commands.GetDisplayName(cmd))
		if numbered, ok := t.getAttribute("numbered"); (ok && numbered.(bool)) || !ok {
			t.setAttribute("numbered", true)
			t.sectionCounter.AddCaptionedCommand(cmd)
			id := t.sectionCounter.GetSectionDisplayId() + "." + strconv.Itoa(t.sectionCounter.GetCaptionNum(cmd))
			t.setAttribute("id", cmd+id)
			t.setAttribute("_caption_num", id)
		}
	}
}

func (t *Translator) checkHeaders() {
	if ctx, ok := t.currCtx.(*parser.CommandContext); ok && ctx.GetInner() != nil {
		inner := ctx.GetInner().GetText()
		switch ctx.GetCmd().GetText() {
		case commands.Title:
			t.sectionCounter.AddTitle(inner)
		case commands.Subtitle:
			t.sectionCounter.AddSubtitle(inner)
		case commands.Heading:
			t.sectionCounter.AddHeading(inner)
		default:
			return
		}
		if _, ok := t.getAttribute("id"); !ok {
			t.setAttribute("id", "h"+t.sectionCounter.GetSectionDisplayId())
		}
		t.setAttribute("sectionNum", t.sectionCounter.GetSectionDisplayId())
	}
}

func (t *Translator) checkImplicitParagraph() {
	// if the command has only one block, the only child is a command block;
	// otherwise, it's not.
	if ctx, ok := t.currCtx.(*parser.CommandContext); ok && ctx.GetInner() != nil && len(ctx.AllBlock()) == 1 {
		t.setAttribute("_command_block", true)
	}
}

func (t *Translator) checkRecursiveCommandConstraint() {
	// check if a block command is nested inside a non-recursive command
	if t.env.parent != nil && !commands.IsRecursive(t.env.parent.command) && commands.IsBlock(t.GetCurrentCommand()) {
		t.reportRuntimeError("cannot have block command inside a non-recursive command")
	}
}

// VisitTerminal is called when a terminal node is visited.
func (t *Translator) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (t *Translator) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (t *Translator) EnterEveryRule(ctx antlr.ParserRuleContext) {
	t.currCtx = ctx
}

// ExitEveryRule is called when any rule is exited.
func (t *Translator) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNewLine is called when production newLine is entered.
func (t *Translator) EnterNewLine(ctx *parser.NewLineContext) {
	// insert spaces between two consecutive lines if there isn't one yet
	if p, ok := ctx.GetParent().(*parser.BlockContext); ok {
		i := 0
		for ; i < p.GetChildCount() && p.GetChild(i) != ctx; i++ {
		}
		if i > 0 {
			if c, ok := p.GetChild(i - 1).(*parser.ContentContext); ok {
				text := c.GetText()
				trimmed := strings.TrimRight(text, " }")
				last := rune(trimmed[len(trimmed)-1])
				// check if the last character is a letter
				if (unicode.IsLetter(last) || unicode.IsPunct(last)) && text[len(text)-1] != ' ' {
					t.documentWriter.WriteString(" ")
				}
			}
		}
	}
}

// ExitNewLine is called when production newLine is exited.
func (t *Translator) ExitNewLine(ctx *parser.NewLineContext) {}

// EnterBlankLines is called when production blankLines is entered.
func (t *Translator) EnterBlankLines(ctx *parser.BlankLinesContext) {}

// ExitBlankLines is called when production blankLines is exited.
func (t *Translator) ExitBlankLines(ctx *parser.BlankLinesContext) {}

// EnterDocument is called when production document is entered.
func (t *Translator) EnterDocument(ctx *parser.DocumentContext) {
	_, err := t.documentWriter.WritePreamble()
	if err != nil {
		return
	}
}

// ExitDocument is called when production document is exited.
func (t *Translator) ExitDocument(ctx *parser.DocumentContext) {
	defer func(documentWriter DocumentWriter, writer io.Writer) {
		_, err := documentWriter.Flush(writer)
		if err != nil {
			t.reportRuntimeError(err.Error())
		}
	}(t.documentWriter, t.writer)

	err := t.refCounter.allMatched()
	if err != nil {
		t.reportRuntimeError(err.Error())
	}

	_, err = t.documentWriter.WritePostamble()
	if err != nil {
		return
	}
}

// EnterBlock is called when production block is entered.
func (t *Translator) EnterBlock(ctx *parser.BlockContext) {
	// if the block is not a direct child of a command, it is implicitly a paragraph
	if val, ok := t.getAttribute("_command_block"); ok && val.(bool) {
		// signal that the next block is no longer the child of a command
		t.setAttribute("_command_block", false)
		// if the child command is already a block, we don't need to wrap it in a paragraph
	} else if c, ok := ctx.GetChild(0).GetChild(0).(*parser.CommandContext); !ok || (ok && !commands.IsBlock(c.GetCmd().GetText())) {
		t.pushEnv(map[string]interface{}{"_implicit_para": true}, commands.Paragraph)
		t.checkRecursiveCommandConstraint()
		t.writeHTMLOpenTag()
	}
}

// ExitBlock is called when production block is exited.
func (t *Translator) ExitBlock(ctx *parser.BlockContext) {
	if t.env.vars["_implicit_para"] != nil {
		t.writeHTMLCloseTag()
		t.popEnv()
	}
}

// EnterContent is called when production content is entered.
func (t *Translator) EnterContent(ctx *parser.ContentContext) {}

// ExitContent is called when production content is exited.
func (t *Translator) ExitContent(ctx *parser.ContentContext) {}

// EnterText is called when production text is entered.
func (t *Translator) EnterText(ctx *parser.TextContext) {
	_, err := t.documentWriter.WriteString(html.EscapeString(ctx.GetText()))
	if err != nil {
		return
	}
}

// ExitText is called when production text is exited.
func (t *Translator) ExitText(ctx *parser.TextContext) {
}

// EnterIdentifier is called when production identifier is entered.
func (t *Translator) EnterIdentifier(ctx *parser.IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (t *Translator) ExitIdentifier(ctx *parser.IdentifierContext) {}

// EnterCommand is called when production command is entered.
func (t *Translator) EnterCommand(ctx *parser.CommandContext) {

	// parse and check the command
	cmd := ctx.GetCmd().GetText()
	validAttrs, ok := commands.GetAttrTypes(cmd)
	if !ok {
		t.reportRuntimeError("command \"" + cmd + "\" is undefined")
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
							t.reportRuntimeError(fmt.Sprintf("wrong attribute type for \"%s\"; expected %s, got %s",
								attr.GetName().GetText(), attrType, reflect.TypeOf(parsed)))
						}
					} else {
						t.reportRuntimeError("attribute \"" + attr.GetName().GetText() + "\" is undefined")
					}
				}
			}
		}
	}
	t.pushEnv(attrsMap, cmd)

	t.checkImplicitParagraph()
	t.checkHeaders()
	t.checkCaptioned()
	t.checkRef()
	t.checkDeref()
	t.checkDuplicateId()
	t.checkRecursiveCommandConstraint()
	t.checkBlockCommandConstraint()
	t.writeHTMLOpenTag()
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

func (t *Translator) pushEnv(attrsMap map[string]interface{}, cmd string) {
	t.env = &environment{parent: t.env, vars: attrsMap, command: cmd}
}

// ExitCommand is called when production command is exited.
func (t *Translator) ExitCommand(ctx *parser.CommandContext) {
	t.writeHTMLCloseTag()
	// pop the environment
	t.popEnv()
}

func (t *Translator) popEnv() {
	t.env = t.env.parent
}

func (t *Translator) writeHTMLTag(tag bool) {
	cmd := t.GetCurrentCommand()
	openTag, closeTag := commands.GetHtmlTags(cmd, t.env.vars)
	if tag {
		_, err := t.documentWriter.WriteString(openTag)
		if err != nil {
			t.reportRuntimeError(err.Error())
		}
	} else {

		_, err := t.documentWriter.WriteString(closeTag)
		if err != nil {
			t.reportRuntimeError(err.Error())
		}
	}
}

func (t *Translator) writeHTMLOpenTag() {
	t.writeHTMLTag(true)
}

func (t *Translator) writeHTMLCloseTag() {
	t.writeHTMLTag(false)
}

// EnterVal is called when production val is entered.
func (t *Translator) EnterVal(ctx *parser.ValContext) {}

// ExitVal is called when production val is exited.
func (t *Translator) ExitVal(ctx *parser.ValContext) {}

// EnterAttr is called when production attr is entered.
func (t *Translator) EnterAttr(ctx *parser.AttrContext) {}

// ExitAttr is called when production attr is exited.
func (t *Translator) ExitAttr(ctx *parser.AttrContext) {}

// EnterAttrs is called when production attrsTypes is entered.
func (t *Translator) EnterAttrs(ctx *parser.AttrsContext) {}

// ExitAttrs is called when production attrsTypes is exited.
func (t *Translator) ExitAttrs(ctx *parser.AttrsContext) {}
