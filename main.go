package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"monkey/compiler"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"monkey/repl"
	"monkey/vm"
	"os"
	"os/user"
)

var engine = flag.String("engine", "vm", "use 'vm' or 'eval'")
var file = flag.String("file", "", "define a file for execution")
var print_result = flag.String("print-result", "false", "use 'true' or 'false'")

func main() {
	flag.Parse()

	if *file == "" {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s! This is the Monkey programming language.\nFeel free to type in commands.\n", user.Username)
		repl.Start(os.Stdin, os.Stdout, *engine == "vm")
	} else {
		buf, err := ioutil.ReadFile(*file)
		if err != nil {
			panic(err)
		}
		source := string(buf)
		result, err := execute(source, *engine == "vm")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return
		}

		if *print_result == "true" {
			fmt.Printf("%s\n", result.Inspect())
		}
	}
}

func execute(source string, useVM bool) (object.Object, error) {
	l := lexer.New(source)
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(os.Stderr, p.Errors())
		return nil, fmt.Errorf("%d parser errors", len(p.Errors()))
	}

	if useVM {
		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			return nil, fmt.Errorf("compiler error: %s", err)
		}

		machine := vm.New(comp.Bytecode())
		err = machine.Run()
		if err != nil {
			return nil, fmt.Errorf("vm error: %s", err)
		}

		return machine.LastPoppedStackElem(), nil
	} else {
		env := object.NewEnvironment()
		macroEnv := object.NewEnvironment()

		evaluator.DefineMacros(program, macroEnv)
		expanded := evaluator.ExpandMacros(program, macroEnv)

		return evaluator.Eval(expanded, env), nil
	}
}

func printParserErrors(out io.Writer, errors []string) {
	fmt.Fprintf(out, repl.PARSER_ERROR_MESSAGE, len(errors))
	for i, msg := range errors {
		fmt.Fprintf(out, " (%4d) => %s\n", i, msg)
	}
}
