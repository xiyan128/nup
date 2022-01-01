package commands

import (
	"fmt"
	"reflect"
	"strings"
	"xiyan.life/nup/translator/mixins"
)

var f = fmt.Sprintf

// BaseCommand is an "interface" common to all commands; it defines the most
// common attributes that all commands must have
type BaseCommand struct {
	name string
	id   string `html:"id"`
}

// RecursiveCommand signifies the command can have nested block commands
// a recursive command must be a block command
type RecursiveCommand struct {
	BlockCommand
}

// BlockCommand signifies the command must be the first command in a block
type BlockCommand struct{}

type MathCommand struct{}

type CaptionedCommand struct {
	numbered bool
	caption  string
}

// every element in structs that starts with an underscore is not an attribute definition,
// but rather hold concrete information about the command; usually, they are the
// fields with actual values
type _HtmlTarget struct {
	htmlOpenTag  string
	htmlCloseTag string
}

var commands = map[string]interface{}{
	// Document
	Document: struct {
		BaseCommand
		BlockCommand
		RecursiveCommand
	}{},
	// Strong
	Strong: struct {
		BaseCommand
		_HtmlTarget
	}{
		_HtmlTarget: _HtmlTarget{`<b{{.attrs}}>`, `</b>`},
	},
	// Italic
	Italic: struct {
		BaseCommand
		_HtmlTarget
	}{
		_HtmlTarget: _HtmlTarget{`<i{{.attrs}}>`, `</i>`},
	},
	// inline code
	// for now, attributes are not supported
	InlineCode: struct {
		BaseCommand
		_HtmlTarget
	}{
		_HtmlTarget: _HtmlTarget{`<code{{.attrs}}>`, `</code>`},
	},
	// inline math
	// for now, attributes are not supported
	InlineMath: struct {
		BaseCommand
		MathCommand
		_HtmlTarget
	}{
		_HtmlTarget: _HtmlTarget{
			htmlOpenTag:  `<span{{.attrs}}>\(`,
			htmlCloseTag: `\)</span>`,
		}},
	// link
	Link: struct {
		BaseCommand
		_HtmlTarget
		to string `html:"href"`
	}{
		_HtmlTarget: _HtmlTarget{
			htmlOpenTag:  `<a{{.attrs}}>`,
			htmlCloseTag: `</a>`,
		},
	},
	// reference
	Reference: struct {
		BaseCommand
		_HtmlTarget
		refNum int
		to     string
	}{
		_HtmlTarget: _HtmlTarget{
			htmlOpenTag:  `<span{{.attrs}}>`,
			htmlCloseTag: `<a href="#{{.to}}" class="refname">{{.refNum}}</a></span>`,
		},
	},
	// dereference
	Dereference: struct {
		BaseCommand
		_HtmlTarget
		refNum int
	}{
		_HtmlTarget: _HtmlTarget{
			htmlOpenTag:  `<span{{.attrs}}><a class="refname" >{{.refNum}}</a>`,
			htmlCloseTag: "</span>",
		},
	},
	// paragraph
	Paragraph: struct {
		BaseCommand
		RecursiveCommand
		_HtmlTarget
	}{
		_HtmlTarget: _HtmlTarget{`<div{{.attrs}}>`, `</div>`},
	},
	// quote
	Quote: struct {
		BaseCommand
		RecursiveCommand
		_HtmlTarget
		cite string `html:"cite"`
	}{
		_HtmlTarget: _HtmlTarget{
			htmlOpenTag:  "<blockquote{{.attrs}}>",
			htmlCloseTag: "</blockquote>",
		},
	},
	// math block
	Math: struct {
		BaseCommand
		MathCommand
		BlockCommand
		CaptionedCommand
		_HtmlTarget
	}{
		_HtmlTarget: _HtmlTarget{`<div{{.attrs}} class="math-block">$$`, f(`$$%s</div>`, mixins.MathNumber)},
	},
	// code block
	Code: struct {
		BaseCommand
		BlockCommand
		CaptionedCommand
		_HtmlTarget
		lang string
	}{
		_HtmlTarget: _HtmlTarget{
			htmlOpenTag:  `<div{{.attrs}} class="captionable"><pre><code class="language-{{or .lang "plain"}}">`,
			htmlCloseTag: f(`</code></pre>%s</div>`, mixins.Caption),
		},
	},
	// figure
	Figure: struct {
		BaseCommand
		BlockCommand
		CaptionedCommand
		_HtmlTarget
		src    string `html:"src"`
		width  int    `html:"width"`
		height int    `html:"height"`
		align  string `html:"data-align"`
	}{
		_HtmlTarget: _HtmlTarget{
			htmlOpenTag: f(`<figure class="captionable"><img{{.attrs}}/>%s<figure>`, mixins.Caption),
		},
	},
	// box
	Box: struct {
		BaseCommand
		RecursiveCommand
		_HtmlTarget
		title string
	}{
		_HtmlTarget: _HtmlTarget{
			htmlOpenTag:  `<div{{.attrs}} class="box">{{if .title}}<div class="box-title">{{.title}}</div>{{end}}`,
			htmlCloseTag: "</div>",
		},
	},
	Title: struct {
		BaseCommand
		BlockCommand
		_HtmlTarget
	}{
		_HtmlTarget: _HtmlTarget{
			htmlOpenTag:  `<h1{{.attrs}}><a class="section-num">{{.sectionNum}}</a>`,
			htmlCloseTag: "</h1>",
		},
	},
	Subtitle: struct {
		BaseCommand
		BlockCommand
		_HtmlTarget
	}{
		_HtmlTarget: _HtmlTarget{
			htmlOpenTag:  `<h2{{.attrs}}><a class="section-num">{{.sectionNum}}</a>`,
			htmlCloseTag: "</h2>",
		},
	},
	Heading: struct {
		BaseCommand
		BlockCommand
		_HtmlTarget
	}{
		_HtmlTarget: _HtmlTarget{
			htmlOpenTag:  `<h3{{.attrs}}><a class="section-num">{{.sectionNum}}</a>`,
			htmlCloseTag: "</h3>",
		},
	},
	OrderedList: struct {
		BaseCommand
		RecursiveCommand
		_HtmlTarget
	}{
		_HtmlTarget: _HtmlTarget{
			htmlOpenTag:  "<ol{{.attrs}}>",
			htmlCloseTag: "</ol>",
		},
	},
	UnorderedList: struct {
		BaseCommand
		RecursiveCommand
		_HtmlTarget
	}{
		_HtmlTarget: _HtmlTarget{
			htmlOpenTag:  "<ul{{.attrs}}>",
			htmlCloseTag: "</ul>",
		},
	},

	ListItem: struct {
		BaseCommand
		RecursiveCommand
		_HtmlTarget
		value   int    `html:"value"`
		numType string `html:"type"`
	}{
		_HtmlTarget: _HtmlTarget{
			htmlOpenTag:  "<li{{.attrs}}>",
			htmlCloseTag: "</li>",
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
			deepExtractField(v, attrsTypes[k])
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

func IsCommand(command string, baseCommand interface{}) bool {
	if def, ok := commands[command]; ok {
		return embedsBaseModel(def, baseCommand)
	}
	return false
}

func embedsBaseModel(v interface{}, baseType interface{}) bool {
	rt := reflect.TypeOf(v)
	if rt.Kind() != reflect.Struct {
		return false
	}
	for i := 0; i < rt.NumField(); i++ {
		if sf := rt.Field(i); sf.Type == reflect.TypeOf(baseType) && sf.Anonymous {
			return true
		} else if sub := reflect.ValueOf(v).Field(i); sf.IsExported() && sub.Kind() == reflect.Struct {
			if embedsBaseModel(sub.Interface(), baseType) {
				return true
			}
		}
	}
	return false
}

// IsBlock returns if a command is a block
func IsBlock(command string) bool {
	return command == Math || IsCommand(command, BlockCommand{})
}

// IsCaptioned returns if a command is a captioned
func IsCaptioned(command string) bool {
	return IsCommand(command, CaptionedCommand{})
}

func IsRecursive(command string) bool {
	return IsCommand(command, RecursiveCommand{})
}

func deepExtractField(value interface{}, res map[string]reflect.Type) []reflect.Value {
	fields := make([]reflect.Value, 0)
	ifv, ift := reflect.ValueOf(value), reflect.TypeOf(value)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		switch v.Kind() {
		case reflect.Struct:
			if ift.Field(i).IsExported() {
				fields = append(fields, deepExtractField(v.Interface(), res)...)
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
