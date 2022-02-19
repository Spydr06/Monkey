package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("Hello %s! this is the Monkey programming language!\n", user.Username)
		fmt.Printf("Starting in REPL mode\n")
		repl.Start(os.Stdin, os.Stdout, true)
	} else if len(args) == 1 {
		buf, err := ioutil.ReadFile(args[0])
		if err != nil {
			panic(err)
		}
		source := string(buf)

		env := object.NewEnvironment()
		macroEnv := object.NewEnvironment()
		l := lexer.New(source)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(os.Stdout, p.Errors())
			return
		}

		evaluator.DefineMacros(program, macroEnv)
		expanded := evaluator.ExpandMacros(program, macroEnv)

		evaluated := evaluator.Eval(expanded, env)
		if evaluated != nil {
			io.WriteString(os.Stdout, evaluated.Inspect())
			io.WriteString(os.Stdout, "\n")
		}
	} else {
		fmt.Errorf("Usage: monkey [<filepath>]\n\nexpected 0, 1 arguments, got=%d", len(args))
	}
}

func printParserErrors(out io.Writer, errors []string) {
	fmt.Fprintf(out, repl.PARSER_ERROR_MESSAGE, len(errors))
	for i, msg := range errors {
		fmt.Fprintf(out, " (%4d) => %s\n", i, msg)
	}
}
