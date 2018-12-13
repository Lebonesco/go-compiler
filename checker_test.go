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

func TestOperations(t *testing.T) {
	tests := []Test{
		{`5 + 5;`, true},
		{`5 + "5";`, false},
		{`5 * 5;`, true},
		{`5 * 5 + 9;`, true},
		{`5 < 10;`, true},
		{`true and true;`, true},
		{`4 and 2;`, false},
		{`true or false;`, true}}

	runTests(tests, t)
}

func TestIdents(t *testing.T) {
	tests := []Test{
		{`let x = 5;`, true},
		{`
			let x = 6;
			let x = 8;`, false},
		{
			`let x = 5;
			x = "hello";`, false},
		{
			`let y = "hey";
			y = "cool";`, true},
		{
			`let x = "hello ";
			let y = "world";
			let z = x + y;`, true},
		{
			`x = 5;`, false}}

	runTests(tests, t)
}

func TestFunctions(t *testing.T) {
	tests := []Test{
		{
			`func add(x Int, y Int) Int {
				return x + y;
			}

			let a = add(1, 3);`, true},
		{
			`func add(x Int, y Int) Int {
				return x + y;
			}

			let z = add("test", 3);`, false},
		{
			`func one() Int {
				return "test";
			}`, false},
	}

	runTests(tests, t)
}

func runTests(tests []Test, t *testing.T) {
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

func stringToChecker(input string) error {
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	res, err := p.Parse(l)
	if err != nil {
		return err
	}

	program, _ := res.(*ast.Program)
	err = checker.Checker(program)
	if err != nil {
		return err
	}
	return nil
}
