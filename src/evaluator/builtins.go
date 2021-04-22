package evaluator

import (
	"bufio"
	"fmt"
	"monkey/object"
	"os"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got=%s", args[0].Type())
			}
		},
	},
	"puts": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			fmt.Println(args[0].Inspect())

			return NULL
		},
	},
	"gets": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) > 1 {
				return newError("wrong number of arguments. got=%d, wand=0, 1", len(args))
			}

			if len(args) > 0 && args[0] != nil {
				fmt.Println(args[0].Inspect())
			}

			scanner := bufio.NewScanner(os.Stdin)
			scanned := scanner.Scan()
			if !scanned {
				return newError("could not scan properly")
			}

			return &object.String{Value: scanner.Text()}
		},
	},
}
