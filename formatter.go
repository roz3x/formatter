// this is a Formatter for all kinds of structs
package Formatter

import (
	"fmt"
	"reflect"
	"strings"
)

// Space is number of <spaces> youll get as indentation
var Space int = 3

var (
	indent []int
	sz     int
)

func push(a int) {
	indent = append(indent, a)
	sz += a
}

func pop() {
	a := indent[len(indent)-1]
	sz -= a
	indent = indent[:len(indent)-1]
}

// Formatter is the formatter
func Format(a interface{}) {
	defer (func() {
		if r := recover(); r != nil {
			// either nil or unexported fields
			fmt.Printf("%snull\n", strings.Repeat(" ", sz))
		}
	})()
	if a == nil {
		fmt.Printf("%snull\n", strings.Repeat(" ", sz))
	} else if reflect.ValueOf(a).Kind() == reflect.Ptr {
		b := reflect.ValueOf(a).Elem().Interface()
		fmt.Printf("%s{\n", strings.Repeat(" ", sz))
		push(Space)
		Format(b)
		pop()
		fmt.Printf("%s}\n", strings.Repeat(" ", sz))
	} else if reflect.ValueOf(a).Kind() == reflect.Struct {
		s := reflect.TypeOf(a)
		fields := make([]interface{}, s.NumField())
		for i := 0; i < s.NumField(); i++ {
			fields[i] = reflect.ValueOf(a).Field(i).Interface()
		}
		tname := fmt.Sprintf("%s", reflect.TypeOf(a))
		fmt.Printf("%s\"%v\":{\n", strings.Repeat(" ", sz), tname)
		push(len(tname))
		for t, i := range fields {
			fmt.Printf("%s\"%v\":\n", strings.Repeat(" ", sz), reflect.TypeOf(a).Field(t).Name)
			push(len(reflect.TypeOf(a).Field(t).Name))
			Format(i)
			pop()
		}
		pop()
		fmt.Printf("%s}\n", strings.Repeat(" ", sz))
	} else if reflect.TypeOf(a).Kind() == reflect.String {
        fmt.Printf("%s\"%v\",\n",strings.Repeat(" ",sz),a)
    } else {
		fmt.Printf("%s%v,\n", strings.Repeat(" ", sz), a)
	}
}

