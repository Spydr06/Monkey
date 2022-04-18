# Monkey
The Monkey Programming language, written in Go, following Thorsten Ball's Books "Writing an Interpreter in Go" and "Writing a Compiler in Go"

Monkey is a small, interpreted or compiled toy programming language with high-level simplicity and easy to write C-like syntax.
<br/>
*The language features:*

* Two backends: an evaluator or a stack-based virtual-machine with bytecode compiler
* Turing-completeness
* expressions like `+` `-` `*` `/` `-`(prefix) `!`(prefix) `==` `!=` `>` `<`
* local and global variables
* macros and quotes (only with the evaluator backend)
* functions/closures as first-class objects
* everything is an expression
* arrays `[1, 2, 3]`
* hash tables `{a: 1, b: 2, c: 3}`
* strings `"Hello"`
* full unit test coverage
* written in 100% pure Go without external dependencies
* Examples in the `examples` directory

## Installation

- Clone this repository or download a precompiled package

```bash
git clone github.com/spydr06/monkey.git
cd monkey
```

Build the executable:

```bash
go build .
```

## Usage

By default, monkey starts it's interactive REPL, just by typing

```bash
./monkey
```

To run a specific file, use the `-file` flag

```bash
./monkey -file examples/helloworld.monkey
```

To use the evaluator backend, use `-engine eval`

```bash
./monkey -engine eval
```

To get help, use `-help`

```bash
./monkey -help
```