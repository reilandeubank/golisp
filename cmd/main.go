package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/reilandeubank/golisp/pkg/scanner"
	// "github.com/reilandeubank/golisp/pkg/interpreter"
	"github.com/reilandeubank/golisp/pkg/parser"
)

// var i interpreter.Interpreter = interpreter.NewInterpreter()
func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("Usage: golox [script]")
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

	run(strings.ToLower(string(bytes)))

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

func run(source string) {
	thisScanner := scanner.NewScanner(source)
	tokens := thisScanner.ScanTokens()

	// For now, just print the tokens
	// for _, token := range tokens {
	// 	fmt.Println(token.String())
	// }

	if scanner.HadError() {
		os.Exit(65)
		return
	}

	parser := parser.NewParser(tokens)
	expr, err := parser.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(65)
		return
	}

	fmt.Println(expr.String())

	// err = i.Interpret(statements)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(70)
	// 	return
	// }
}
