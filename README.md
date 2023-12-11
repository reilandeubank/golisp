# golisp: a lisp (yisp) tree-walk interpreter written in Go
 
This is a working interpreter for a subset of the Lisp language (aka Yisp)
as outlined by Don Yessick for CS403 at the University of Alabama

This interpreter is feature complete for the language as defined by Yessick

Notes on choices made:

```true``` for truthy values

```nil``` for falsey/nil values

```=``` used for equality checking

```cond``` used for conditional statements

The language is not case sensitive

# Instructions

## Installation
In order to compile the source code, you will need [Go 1.20+](https://go.dev/dl/), as well as [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git). Go installation
instructions can be found [here](https://go.dev/doc/install). To make sure the compiler has been properly installed, type the following command:
```
$ go version
```

Finally, clone the repository into the current directory with
```
$ git clone https://github.com/reilandeubank/golisp
```
or into ```path/to/directory``` using 
```
$ git clone https://github.com/reilandeubank/golisp path/to/directory
```

## Compiling
First, move into the ```golisp``` directory that was just created

From here, you can use the ```make``` command to compile the interpreter into the executable ```./main``` in the current directory.

For submission purposes, ```make``` also runs ```/test/tester.lsp```, which tests every implemented feature in the language
and outputs the checking to the console

Alternatively, you can compile an executable ```./main``` in the current directory. 
```
$ go build cmd/main.go
```

## Usage
Usage for the interpreter is
```
$ ./main
```
to start a REPL or
```
$ ./main file.lsp
```
to run ```file.lsp```

# About the project
This project was my second ever project in Go after writing my Lox interpreter. I have to say I enjoyed the language just as much as I did the first go around, and I was again glad I had chosen a language that was both so simple to pick up and so powerful. My largest problems that I ran into in this implementation mainly revolved around working through the underlying workings of the Lisp language that I had not considered before. Once I figured out that everything in the language was either a list or an atom/symbol, it became much easier to work through the implementation. I'll admit that I may not have done everything the most optimally (see my giant switch statements in interpreter/visitExpr.go) but I worked through most things multiple times in order to make it work as intended. An example would be the functions, which I initially attempted to detect at runtime, meaning I just parsed the definition and calls as lists, and tried to work those into definition or call statements at runtime. This nearly broke my brain and produced some code very reminiscent of spaghetti, but after trashing all of my changes and starting over, I managed to make definitions and calls into special cases in the parser that far simplified the process, as I could borrow a lot of the interpretation logic from the Lox interpreter. Other than functions, most of the project was fairly smooth sailing and I'm pretty proud of my ability to bang out a working interpreter without having to follow the guidance of a textbook.
