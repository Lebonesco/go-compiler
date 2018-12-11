package main

import (
	"github.com/Lebonesco/go-compiler/lexer"
	"github.com/Lebonesco/go-compiler/parser"
	"testing"
)

func TestParser(t *testing.T) {
	const input = `
			/* comment should not be scanned */
			let five = "test";
			let ten = 10;

			func add() Int {
				return x + y;
			}

			let result = 4;  
			5 < 10;

			if (5 < 10) {
				print(10);
			} else {
				print(5);
			}

			10 == 10;
			`

	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	_, err := p.Parse(l)
	if err != nil {
		t.Fatalf(err.Error())
	}
}
