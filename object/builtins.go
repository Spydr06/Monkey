package object

import (
	"bufio"
	"fmt"
	"os"
)

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

var Builtins = []struct {
	Name    string
	Builtin *Builtin
}{
	{
		"len",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *Array:
				return &Integer{Value: int64(len(arg.Elements))}
			case *String:
				return &Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got=%s", args[0].Type())
			}
		},
		},
	},
	{
		"puts",
		&Builtin{Fn: func(args ...Object) Object {
			for _, arg := range args {
				fmt.Print(arg.Inspect())
			}

			fmt.Println()

			return nil
		},
		},
	},
	{
		"first",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to ´first´ must be ARRAY, got=%s", args[0].Type())
			}
			arr := args[0].(*Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}
			return nil
		},
		},
	},
	{
		"last",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to `last` must be ARRAY, got=%s", args[0].Type())
			}

			arr := args[0].(*Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}

			return nil
		},
		},
	},
	{
		"rest",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to `rest` must be ARRAY, got=%s", args[0].Type())
			}

			arr := args[0].(*Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]Object, length-1)
				copy(newElements, arr.Elements[1:length])
				return &Array{Elements: newElements}
			}

			return nil
		},
		},
	},
	{
		"push",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got=%s", args[0].Type())
			}

			arr := args[0].(*Array)
			length := len(arr.Elements)

			newElements := make([]Object, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &Array{Elements: newElements}
		},
		},
	},
	{
		"gets",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) > 1 {
				return newError("wrong number of arguments. got=%d, want=0, 1", len(args))
			}

			if len(args) > 0 && args[0] != nil {
				fmt.Print(args[0].Inspect())
			}

			scanned := scanner.Scan()
			if !scanned {
				return newError("could not scan properly")
			}

			return &String{Value: scanner.Text()}
		},
		},
	},
	{
		"exit",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) > 2 {
				return newError("wrong number of arguments. got=%d, want=1, 2", len(args))
			}

			switch len(args) {
			case 0:
				os.Exit(0)
			case 1:
				_, ok := args[0].(*Integer)
				if !ok {
					return newError("argument to `exit` not supported, got=%s", args[0].Type())
				}
			case 2:
				_, ok := args[0].(*Integer)
				if !ok {
					return newError("argument to `exit` not supported, got=%s", args[0].Type())
				}
				_, ok = args[1].(*String)
				if !ok {
					return newError("argument to `exit` not supported, got=%s", args[1].Type())
				}
				fmt.Println("Exit: " + args[1].Inspect())
			}

			os.Exit(int(args[0].(*Integer).Value))
			return nil
		},
		},
	},
	{
		"va_args",
		&Builtin{Fn: func(args ...Object) Object {
			return &Array{Elements: args}
		},
		},
	},
}

func GetBuiltinByName(name string) *Builtin {
	for _, def := range Builtins {
		if def.Name == name {
			return def.Builtin
		}
	}
	return nil
}

func newError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}
