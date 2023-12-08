# golisp: a lisp (yisp) tree-walk interpreter written in Go
 
This is a working interpreter for a subset of the Lisp language (aka Yisp)
as outlined by Don Yessick for CS403 at the University of Alabama

This interpreter is feature complete for the language as defined by Yessick
Notes on choices made:

```true``` is the truth value

```nil``` used for false/nil value

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

