# Makefile

# The build target executable:
TARGET = main

.PHONY: all build run clean

all: build run

build:
	go build -o $(TARGET) cmd/main.go

run:
	./$(TARGET) test/tester.lsp

clean:
	rm $(TARGET)