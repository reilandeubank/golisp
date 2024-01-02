package main

import (
	"bufio"
	"fmt"
	"golisp/pkg/interpreter"
	"golisp/pkg/parser"
	"golisp/pkg/scanner"
	"os"
	"strings"
)

var i interpreter.Interpreter = interpreter.NewInterpreter()

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("Usage: golisp [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		err := runFile(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(70)
		}
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
	theScanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">>> ")
		if !theScanner.Scan() {
			break
		}

		line := theScanner.Text()
		err := run(strings.ToLower(line))
		if err != nil {
			fmt.Println(err)
			os.Exit(70)
		}
		scanner.SetErrorFlag(false)
	}

	if theScanner.Err() != nil {
		fmt.Println("An error occurred:", theScanner.Err())
	}
}

func run(source string) error {
	thisScanner := scanner.NewScanner(source)
	tokens := thisScanner.ScanTokens()

	thisParser := parser.NewParser(tokens)
	expr, err := thisParser.Parse()
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
