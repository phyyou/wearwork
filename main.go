package main

import (
	"fmt"
	"os"
	"wearwork/evaluator"
	"wearwork/lexer"
	"wearwork/object"
	"wearwork/parser"
	"wearwork/repl"
)

// to enable repl mode, run with --repl flag
func main() {
	if len(os.Args) > 1 && os.Args[1] == "--repl" {
		fmt.Printf("WEAR WORK REPL PROMPT\n")
		repl.Start(os.Stdin, os.Stdout)
	} else {
		fmt.Println(evaluate("1 + 2 * 3"))                                                // 7
		fmt.Println(evaluate("(1 + 2) * 3"))                                              // 9
		fmt.Println(evaluate("1 / 32.5 + 167 * (3498 - 1155) * -721 * (4885 - 1) / 0.5")) // -2755685654567.969
		fmt.Println(evaluate("sin(cos(1)) * cos(1)"))                                     // 0.2779289443079115
	}
}

func evaluate(expression string) float64 {
	l := lexer.New(expression)
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(p.Errors())
		return 0
	}
	evaluated := evaluator.Eval(program)
	return evaluated.(*object.Double).Value
}

func printParserErrors(errors []string) {
	for _, msg := range errors {
		fmt.Println("\t" + msg)
	}
}
