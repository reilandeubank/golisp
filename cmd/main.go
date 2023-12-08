package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/reilandeubank/golisp/pkg/scanner"
	"github.com/reilandeubank/golisp/pkg/interpreter"
	"github.com/reilandeubank/golisp/pkg/parser"
)

var i interpreter.Interpreter = interpreter.NewInterpreter()
func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("Usage: golisp [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}

func runFile(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = run(strings.ToLower(string(bytes)))

	if scanner.HadError() {
		os.Exit(65)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(70)
	}

	return nil
}

func runPrompt() {
	bufscanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">>> ")
		if !bufscanner.Scan() {
			break
		}

		line := bufscanner.Text()
		run(strings.ToLower(line))
		scanner.SetErrorFlag(false)
	}

	if bufscanner.Err() != nil {
		fmt.Println("An error occurred:", bufscanner.Err())
	}
}

func run(source string) error {
	thisScanner := scanner.NewScanner(source)
	tokens := thisScanner.ScanTokens()

	parser := parser.NewParser(tokens)
	expr, err := parser.Parse()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// fmt.Println("Expression: " + expr.String())
	// fmt.Println()

	// fmt.Println("Interpreting...")
	// fmt.Println()

	err = i.Interpret(expr)
	// if output != nil {
	// 	fmt.Println(output)
	// }
	if err != nil {
		fmt.Println(err)
		os.Exit(70)
		return err
	}
	return nil
}
