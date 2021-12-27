package main

import "reflect"

// define a map of command and attributes

type BaseCommand struct {
	name string
	id   string
}

var commands = map[string]interface{}{
	// Strong
	"bf": struct{ BaseCommand }{},
	// Italic
	"it": struct{ BaseCommand }{},
	// inline code
	"ic": struct{ BaseCommand }{},
	// inline math
	"im": struct{ BaseCommand }{},
	// link
	"href": struct {
		BaseCommand
		to string
	}{},
	// reference
	"ref": struct {
		BaseCommand
		to string
	}{},
	// dereference
	"ref*": struct {
		BaseCommand
		from string
	}{},
	// quote
	"quote": struct{ BaseCommand }{},
	// math block
	"math": struct{ BaseCommand }{},
	// code block
	"code": struct {
		BaseCommand
		lang string
	}{},
	// image
	"image": struct {
		BaseCommand
		width  string
		height string
		align  string
	}{},
	// box
	"box": struct{ BaseCommand, title string }{},
}

var attrs map[string]map[string]reflect.Type

// Attrs returns the lookup map of the attributes of the command
func Attrs(command string) (map[string]reflect.Type, bool) {
	if attrs == nil {
		attrs = make(map[string]map[string]reflect.Type)
		for k, v := range commands {
			attrs[k] = make(map[string]reflect.Type)
			extractFields(v, attrs[k])
		}
	}
	if attrs[command] == nil {
		return nil, false
	}
	return attrs[command], true
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
