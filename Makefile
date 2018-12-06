GOCMD=go 
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get 
GENERATOR=grammer.bnf
GENERATE=../../../../bin/gocc
BINARY_NAME=compiler
NO_WARNINGS=-w


all: test build run
build: 
	$(GENERATE) $(GENERATOR) 
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GENERATE) $(GENERATOR)
	$(GOTEST) -v 
test-all:
	$(GOTEST) ./... # recursively runs all test files
clean:
	$(GOCLEAN)
	rm -rf $(BINARY_NAME) util token lexer parser errors
	rm -f LR1_conflicts.txt LR1_sets.txt first.txt lexer_sets.txt terminals.txt
run:
	 $(GENERATE) $(GENERATOR) # create lexer and parser
	$(GOBUILD) -o $(BINARY_NAME) -v # build program
	./$(BINARY_NAME) $(file) # run compiler
#	gcc $(NO_WARNINGS) ./build/main.c ./build/Builtins.c ./build/Builtins.h
#	./a.exe

deps:
	$(GOGET) github.com/goccmack/gocc