package evaluator

import "monkey/object"

var builtins = map[string]*object.Builtin{
	"len":     object.GetBuiltinByName("len"),
	"first":   object.GetBuiltinByName("first"),
	"last":    object.GetBuiltinByName("last"),
	"rest":    object.GetBuiltinByName("rest"),
	"push":    object.GetBuiltinByName("push"),
	"puts":    object.GetBuiltinByName("puts"),
	"gets":    object.GetBuiltinByName("gets"),
	"exit":    object.GetBuiltinByName("exit"),
	"va_args": object.GetBuiltinByName("va_args"),
	"mod":     object.GetBuiltinByName("mod"),
}
