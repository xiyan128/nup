package translator

import (
	"fmt"
	"reflect"
	"strings"
)

// BaseCommand is an "interface" common to all commands; it defines the most
// common attributes that all commands must have
type BaseCommand struct {
	_HtmlTarget
	name string
	id   string `html:"id"`
}

// BlockCommand signifies the command can have nested block commands
type BlockCommand struct {
	blockId string
}

type MathCommand struct {
	mathId string
}

// every element in structs that starts with an underscore is not an attribute definition,
// but rather hold concrete information about the command; usually, they are the
// fields with actual values
type _HtmlTarget struct {
	htmlOpenTag  string
	htmlCloseTag string
}

// Enum of all commands
const (
	Strong      = "bf"
	Italic      = "it"
	InlineCode  = "`"
	InlineMath  = "$"
	Link        = "href"
	Reference   = "ref"
	Dereference = "deref"
	Paragraph   = "para"
	Quote       = "quote"
	Math        = "$$"
	Code        = "```"
	Figure      = "figure"
	Box         = "box"
	Title       = "title"
	Subtitle    = "subtitle"
	Heading     = "heading"
)

var commands = map[string]interface{}{
	// Strong
	Strong: struct{ BaseCommand }{
		BaseCommand{
			_HtmlTarget: _HtmlTarget{`<b{{.attrs}}>`, `</b>`},
		},
	},
	// Italic
	Italic: struct{ BaseCommand }{
		BaseCommand{
			_HtmlTarget: _HtmlTarget{`<i{{.attrs}}>`, `</i>`},
		},
	},
	// inline code
	// for now, attributes are not supported
	InlineCode: struct{ BaseCommand }{
		BaseCommand{
			_HtmlTarget: _HtmlTarget{`<code{{.attrs}}>`, `</code>`},
		},
	},
	// inline math
	// for now, attributes are not supported
	InlineMath: struct {
		BaseCommand
		MathCommand
	}{
		BaseCommand: BaseCommand{
			_HtmlTarget: _HtmlTarget{
				htmlOpenTag:  `<span{{.attrs}}>\(`,
				htmlCloseTag: `\)</span>`,
			}},
	},
	// link
	Link: struct {
		BaseCommand
		to string `html:"href"`
	}{
		BaseCommand: BaseCommand{
			_HtmlTarget: _HtmlTarget{
				htmlOpenTag:  `<a{{.attrs}}>`,
				htmlCloseTag: `</a>`,
			},
		},
	},
	// reference
	Reference: struct {
		BaseCommand
		refNum int
		to     string
	}{
		BaseCommand: BaseCommand{
			_HtmlTarget: _HtmlTarget{
				htmlOpenTag:  `<span{{.attrs}}>`,
				htmlCloseTag: `<a href="#{{.to}}"><sup>{{.refNum}}</sup></a></span>`,
			},
		},
	},
	// dereference
	Dereference: struct {
		BaseCommand
		refNum int
	}{
		BaseCommand: BaseCommand{
			_HtmlTarget: _HtmlTarget{
				htmlOpenTag:  "<span{{.attrs}}><sup>{{.refNum}}</sup>",
				htmlCloseTag: "</span>",
			},
		},
	},
	// paragraph
	Paragraph: struct {
		BaseCommand
		BlockCommand
	}{
		BaseCommand: BaseCommand{
			_HtmlTarget: _HtmlTarget{`<div{{.attrs}}>`, `</div>`},
		},
	},
	// quote
	Quote: struct {
		BaseCommand
		BlockCommand
		cite string `html:"cite"`
	}{
		BaseCommand: BaseCommand{
			_HtmlTarget: _HtmlTarget{
				htmlOpenTag:  "<blockquote{{.attrs}}>",
				htmlCloseTag: "</blockquote>",
			},
		},
	},
	// math block
	// for not, attributes are not supported
	Math: struct {
		BaseCommand
		MathCommand
	}{
		BaseCommand: BaseCommand{
			_HtmlTarget: _HtmlTarget{`<div{{.attrs}}>$$`, `$$</div>`},
		},
	},
	// code block
	Code: struct {
		BaseCommand
		lang string `html:"data-lang"`
	}{
		BaseCommand: BaseCommand{
			_HtmlTarget: _HtmlTarget{
				htmlOpenTag:  "<pre><code{{.attrs}}>",
				htmlCloseTag: "</code></pre>",
			},
		},
	},
	// figure
	Figure: struct {
		BaseCommand
		BlockCommand
		src    string `html:"src"`
		width  int    `html:"width"`
		height int    `html:"height"`
		align  string `html:"data-align"`
	}{
		BaseCommand: BaseCommand{
			_HtmlTarget: _HtmlTarget{
				htmlOpenTag:  "<figure><img{{.attrs}}/><figcaption>",
				htmlCloseTag: "</figcaption></figure>",
			},
		},
	},
	// box
	Box: struct {
		BaseCommand
		BlockCommand
		title string
	}{
		BaseCommand: BaseCommand{
			_HtmlTarget: _HtmlTarget{
				htmlOpenTag:  "<div{{.attrs}}>{{if .title}}<div>{{.title}}</div>{{end}}",
				htmlCloseTag: "</div>",
			},
		},
	},
	Title: struct {
		BaseCommand
	}{
		BaseCommand: BaseCommand{
			_HtmlTarget: _HtmlTarget{
				htmlOpenTag:  "<h1{{.attrs}}>",
				htmlCloseTag: "</h1>",
			},
		},
	},
	Subtitle: struct {
		BaseCommand
	}{
		BaseCommand: BaseCommand{
			_HtmlTarget: _HtmlTarget{
				htmlOpenTag:  "<h2{{.attrs}}>",
				htmlCloseTag: "</h2>",
			},
		},
	},
	Heading: struct {
		BaseCommand
	}{
		BaseCommand: BaseCommand{
			_HtmlTarget: _HtmlTarget{
				htmlOpenTag:  "<h3{{.attrs}}>",
				htmlCloseTag: "</h3>",
			},
		},
	},
}

// singleton
var attrsTypes map[string]map[string]reflect.Type

// GetAttrTypes returns the lookup map of the attribute types of the command
func GetAttrTypes(command string) (map[string]reflect.Type, bool) {
	if attrsTypes == nil {
		attrsTypes = make(map[string]map[string]reflect.Type)
		for k, v := range commands {
			attrsTypes[k] = make(map[string]reflect.Type)
			extractFields(v, attrsTypes[k])
		}
	}
	if attrsTypes[command] == nil {
		return nil, false
	}
	return attrsTypes[command], true
}

// HasAttr returns if a command has an attribute
func HasAttr(command string, attr string) (ok bool) {
	attrs, ok := GetAttrTypes(command)
	if !ok {
		return false
	}
	_, ok = attrs[attr]
	return
}

// IsBlock returns if a command is a block
func IsBlock(command string) bool {
	return HasAttr(command, "blockId") || command == ""
}

// IsMath returns if a command is a math command
func IsMath(command string) bool {
	return HasAttr(command, "mathId")
}

func extractFields(value interface{}, res map[string]reflect.Type) []reflect.Value {
	fields := make([]reflect.Value, 0)
	ifv, ift := reflect.ValueOf(value), reflect.TypeOf(value)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		switch v.Kind() {
		case reflect.Struct:
			if ift.Field(i).IsExported() {
				fields = append(fields, extractFields(v.Interface(), res)...)
			}
		default:
			res[ift.Field(i).Name] = v.Type()
			fields = append(fields, v)
		}
	}
	return fields
}

func GetHtmlAttrName(command, name string) (tag string, ok bool) {
	attrs, ok := commands[command]
	if !ok {
		return
	}
	return extractTagValue(attrs, "html", name)
}

// extractTagValue get the value of a tag
func extractTagValue(value interface{}, tag, key string) (htmlAttr string, ok bool) {

	ifv, ift := reflect.ValueOf(value), reflect.TypeOf(value)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		if f := ift.Field(i); f.Name == key {
			res := f.Tag.Get(tag)
			if res != "" {
				return f.Tag.Get("html"), true
			}
		}
		if v.Kind() == reflect.Struct {
			if ift.Field(i).IsExported() {
				htmlAttr, ok = extractTagValue(v.Interface(), tag, key)
				if ok {
					return
				}
			}
		}
	}
	return "", false
}

// GetHtmlTags returns the open tag and the close tag based on the commands definitions
func GetHtmlTags(command string, vars map[string]interface{}) (string, string) {

	if attrs, ok := commands[command]; ok {
		htmlTarget := reflect.
			ValueOf(attrs).
			FieldByName(reflect.TypeOf(BaseCommand{}).Name()).
			FieldByName(reflect.TypeOf(_HtmlTarget{}).Name())
		op, cl := htmlTarget.Field(0).String(), htmlTarget.Field(1).String()
		attrsBuilder := strings.Builder{}
		attrsBuilder.WriteString("")
		for k, v := range vars {
			htmlAttrName, ok := GetHtmlAttrName(command, k)
			if !ok {
				// just ignore the attribute because it is probably not defined
				// for the HTML target
				continue
			}
			attrsBuilder.WriteRune(' ')
			attrsBuilder.WriteString(htmlAttrName)
			attrsBuilder.WriteString("=\"")
			// for special cases, do not use the attribute building; use separate
			// template instead
			attrsBuilder.WriteString(fmt.Sprintf("%v", v))
			attrsBuilder.WriteRune('"')
		}
		op = strings.Replace(op, "{{.attrs}}", attrsBuilder.String(), -1)
		return ProcessString(op, vars), ProcessString(cl, vars)
	}
	return "", ""
}
