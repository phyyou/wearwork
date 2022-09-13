# WEAR Work

This repository is assignment by wear.

evaluate the string expression.

### evaluate function

```golang
// main.go
package main

/// ...

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

/// ...

```

## How to use

### 0. Install golang

Install golang.  
See [Golang](https://golang.org/doc/install).

### 1. Clone this repository

```bash
$ git clone https://github.com/phyyou/wearwork.git
$ cd wearwork
```

### 2. Run

```bash
$ go mod tidy
```

```bash
$ go run main.go
```

Also, you can run it with repl mode.

```bash
$ go run main.go --repl
```

## Author

[phyyou](https://github.com/phyyou)

gydudwls@gmail.com

## Reference
