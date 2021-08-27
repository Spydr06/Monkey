package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/compiler"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"monkey/vm"
)

const PROMPT = ">> "
const PARSER_ERROR_MESSAGE = "encountered %d Errors while parsing: \n"

func Start(in io.Reader, out io.Writer, useVM bool) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()

	for {
		fmt.Fprintf(out, "%s", PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		if useVM {
			comp := compiler.New()
			err := comp.Compile(program)
			if err != nil {
				fmt.Fprintf(out, "Compilation failed:\n -> %s\n", err)
				continue
			}

			machine := vm.New(comp.Bytecode())
			err = machine.Run()
			if err != nil {
				fmt.Fprintf(out, "Executing Bytecode failed:\n -> %s\n", err)
				continue
			}

			stackTop := machine.StackTop()
			io.WriteString(out, stackTop.Inspect())
			io.WriteString(out, "\n")
		} else {
			evaluator.DefineMacros(program, macroEnv)
			expanded := evaluator.ExpandMacros(program, macroEnv)

			evaluated := evaluator.Eval(expanded, env)
			if evaluated != nil {
				io.WriteString(out, evaluated.Inspect())
				io.WriteString(out, "\n")
			}
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	fmt.Fprintf(out, PARSER_ERROR_MESSAGE, len(errors))
	for i, msg := range errors {
		fmt.Fprintf(out, " (%4d) => %s\n", i, msg)
	}
}
