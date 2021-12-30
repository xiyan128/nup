package main

import (
	"fmt"
	"reflect"
	"strings"
)

// define a map of command and attributes

type HtmlTarget struct {
	_htmlOpenTag  string
	_htmlCloseTag string
}

type BaseCommand struct {
	HtmlTarget
	name string
	id   string `html:"id"`
}

type BlockCommand struct {
	blockId string
}

type MathCommand struct {
	mathId string
}

var commands = map[string]interface{}{
	// Strong
	"bf": struct{ BaseCommand }{
		BaseCommand{
			HtmlTarget: HtmlTarget{`<b{{.attrs}}>`, `</b>`},
		},
	},
	// Italic
	"it": struct{ BaseCommand }{
		BaseCommand{
			HtmlTarget: HtmlTarget{`<i{{.attrs}}>`, `</i>`},
		},
	},
	// inline code
	// for now, attributes are not supported
	"`": struct{ BaseCommand }{
		BaseCommand{
			HtmlTarget: HtmlTarget{`<code{{.attrs}}>`, `</code>`},
		},
	},
	// inline math
	// for now, attributes are not supported
	"$": struct {
		BaseCommand
		MathCommand
	}{
		BaseCommand: BaseCommand{
			HtmlTarget: HtmlTarget{
				_htmlOpenTag:  `<span{{.attrs}}>\(`,
				_htmlCloseTag: `\)</span>`,
			}},
	},
	// link
	"href": struct {
		BaseCommand
		to string `html:"href"`
	}{
		BaseCommand: BaseCommand{
			HtmlTarget: HtmlTarget{
				_htmlOpenTag:  `<a{{.attrs}}>`,
				_htmlCloseTag: `</a>`,
			},
		},
	},
	// reference
	"ref": struct {
		BaseCommand
		refNum int
		to     string
	}{
		BaseCommand: BaseCommand{
			HtmlTarget: HtmlTarget{
				_htmlOpenTag:  `<span{{.attrs}}>`,
				_htmlCloseTag: `<a href="#{{.to}}"><sup>{{.refNum}}</sup></a></span>`,
			},
		},
	},
	// dereference
	"deref": struct {
		BaseCommand
		refNum int
	}{
		BaseCommand: BaseCommand{
			HtmlTarget: HtmlTarget{
				_htmlOpenTag:  "<span{{.attrs}}><sup>{{.refNum}}</sup>",
				_htmlCloseTag: "</span>",
			},
		},
	},
	// paragraph
	"para": struct {
		BaseCommand
		BlockCommand
	}{
		BaseCommand: BaseCommand{
			HtmlTarget: HtmlTarget{`<div{{.attrs}}>`, `</div>`},
		},
	},
	// quote
	"quote": struct {
		BaseCommand
		BlockCommand
		cite string `html:"cite"`
	}{
		BaseCommand: BaseCommand{
			HtmlTarget: HtmlTarget{
				_htmlOpenTag:  "<blockquote{{.attrs}}>",
				_htmlCloseTag: "</blockquote>",
			},
		},
	},
	// math block
	// for not, attributes are not supported
	"$$": struct {
		BaseCommand
		MathCommand
	}{
		BaseCommand: BaseCommand{
			HtmlTarget: HtmlTarget{`<div{{.attrs}}>$$`, `$$</div>`},
		},
	},
	// code block
	"```": struct {
		BaseCommand
		lang string `html:"data-lang"`
	}{
		BaseCommand: BaseCommand{
			HtmlTarget: HtmlTarget{
				_htmlOpenTag:  "<pre><code{{.attrs}}>",
				_htmlCloseTag: "</code></pre>",
			},
		},
	},
	"figure": struct {
		BaseCommand
		BlockCommand
		src    string `html:"src"`
		width  int    `html:"width"`
		height int    `html:"height"`
		align  string `html:"data-align"`
	}{
		BaseCommand: BaseCommand{
			HtmlTarget: HtmlTarget{
				_htmlOpenTag:  "<figure><img{{.attrs}}/><figcaption>",
				_htmlCloseTag: "</figcaption></figure>",
			},
		},
	},
	// box
	"box": struct {
		BaseCommand
		BlockCommand
		title string
	}{
		BaseCommand: BaseCommand{
			HtmlTarget: HtmlTarget{
				_htmlOpenTag:  "<div{{.attrs}}>{{if .title}}<div>{{.title}}</div>{{end}}",
				_htmlCloseTag: "</div>",
			},
		},
	},
	"title": struct {
		BaseCommand
	}{
		BaseCommand: BaseCommand{
			HtmlTarget: HtmlTarget{
				_htmlOpenTag:  "<h1{{.attrs}}>",
				_htmlCloseTag: "</h3>",
			},
		},
	},
	"subtitle": struct {
		BaseCommand
	}{
		BaseCommand: BaseCommand{
			HtmlTarget: HtmlTarget{
				_htmlOpenTag:  "<h2{{.attrs}}>",
				_htmlCloseTag: "</h3>",
			},
		},
	},
	"heading": struct {
		BaseCommand
	}{
		BaseCommand: BaseCommand{
			HtmlTarget: HtmlTarget{
				_htmlOpenTag:  "<h3{{.attrs}}>",
				_htmlCloseTag: "</h3>",
			},
		},
	},
}

var attrsTypes map[string]map[string]reflect.Type

// Attrs returns the lookup map of the attributes of the command
func Attrs(command string) (map[string]reflect.Type, bool) {
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

func HasAttr(command string, attr string) (ok bool) {
	attrs, ok := Attrs(command)
	if !ok {
		return false
	}
	_, ok = attrs[attr]
	return
}

func IsBlock(command string) bool {
	return HasAttr(command, "blockId") || command == ""
}

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
			fields = append(fields, extractFields(v.Interface(), res)...)
		default:
			res[ift.Field(i).Name] = v.Type()
			fields = append(fields, v)
		}
	}
	return fields
}

func extractHtmlAttrName(command, name string) (tag string, ok bool) {
	attrs, ok := commands[command]
	if !ok {
		return
	}
	return extractTag(attrs, name)
}

func extractTag(value interface{}, key string) (htmlAttr string, ok bool) {

	ifv, ift := reflect.ValueOf(value), reflect.TypeOf(value)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		if f := ift.Field(i); f.Name == key {
			res := f.Tag.Get("html")
			if res != "" {
				return f.Tag.Get("html"), true
			}
		}
		if v.Kind() == reflect.Struct {
			htmlAttr, ok = extractTag(v.Interface(), key)
			if ok {
				return
			}
		}
	}
	return "", false
}

func GetHtmlTags(command string, vars map[string]interface{}) (string, string) {

	if attrs, ok := commands[command]; ok {
		htmlTarget := reflect.
			ValueOf(attrs).
			FieldByName(reflect.TypeOf(BaseCommand{}).Name()).
			FieldByName(reflect.TypeOf(HtmlTarget{}).Name())
		op, cl := htmlTarget.Field(0).String(), htmlTarget.Field(1).String()
		attrsBuilder := strings.Builder{}
		attrsBuilder.WriteString("")
		for k, v := range vars {
			htmlAttrName, ok := extractTag(attrs, k)
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
