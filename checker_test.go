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
		Test{`5 + 5;`, true},
		Test{`5 + "5";`, false},
		Test{`5 * 5;`, true},
		Test{`"hello" + " " + "world";`, true},
		Test{`5 <= 10;`, true},
		Test{`true and true;`, true},
		Test{`4 and 2;`, false},
		Test{`true or false;`, true}}

	for _, test := range tests {
		err := stringToChecker(test.src)
		if err != nil && !test.Success {
			t.Fatalf(err.Error())
		}
	}
}

func TestIdents(t *testing.T) {

}

func TestFunctions(t *testing.T) {

}
