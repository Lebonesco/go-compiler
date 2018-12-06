package main

import (
	"github.com/Lebonesco/go-compiler/ast"
	"github.com/Lebonesco/go-compiler/checker"
	"github.com/Lebonesco/go-compiler/lexer"
	"github.com/Lebonesco/go-compiler/parser"
	"testing"
)

type Test struct {
	src     string
	Success bool
}

func stringToChecker(input string) error {
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	res, err := p.Parse(l)
	if err != nil {
		return err
	}

	program, _ := res.(*ast.Program)
	_, err = checker.Checker(program)
	if err != nil {
		return err
	}
	return nil
}

func TestOperations(t *testing.T) {
	tests := []Test{
		{`5 + 5;`, true},
		{`5 + "5";`, false},
		{`5 * 5;`, true},
		{`5 * 5 + 9;`, true},
		{`5 <= 10;`, true},
		{`true and true;`, true},
		{`4 and 2;`, false},
		{`true or false;`, true}}

	for i, test := range tests {
		err := stringToChecker(test.src)

		if err != nil && !test.Success {
			continue
		} else if err != nil {
			t.Fatalf("test %d fail: "+err.Error(), i)
		} else if !test.Success {
			t.Fatalf("test %d should have failed", i)
		}
	}
}

func TestIdents(t *testing.T) {

}

func TestFunctions(t *testing.T) {

}
