GOCMD=go
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get 
GENERATOR=grammer.bnf
GENERATE=../../../../bin/gocc
BINARY_NAME=compiler


all: test run

test:
	$(GENERATE) $(GENERATOR)
	$(GOTEST) -v 

clean:
	$(GOCLEAN)
	rm -rf $(BINARY_NAME) util token lexer parser errors
	rm -f LR1_conflicts.txt LR1_sets.txt first.txt lexer_sets.txt terminals.txt
run:
	$(GENERATE) $(GENERATOR) # create lexer and parser
	$(GOCMD) run main.go $(file) # run compiler

deps:
	$(GOGET) github.com/goccmack/gocc