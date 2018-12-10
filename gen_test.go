package main

import (
	"bytes"
	"github.com/Lebonesco/go-compiler/ast"
	"github.com/Lebonesco/go-compiler/checker"
	"github.com/Lebonesco/go-compiler/gen"
	"github.com/Lebonesco/go-compiler/lexer"
	"github.com/Lebonesco/go-compiler/parser"
	"strings"
	"testing"
)

func TestGen(t *testing.T) {
	tests := []struct {
		src string
		res string
	}{
		{
			src: ``,
			res: ``},
		{
			src: `5 + 5;`,
			res: `
			Int tmp_1 = Int(5);
			Int tmp_2 = Int(5);
			Int tmp_3 = tmp_1->PLUS(tmp_2);
			tmp_3;`},
		{
			src: `10 < 4;`,
			res: `
			Int tmp_1 = Int(10);
			Int tmp_2 = Int(4);
			Bool tmp_3 = tmp_1->LESS(tmp_2);
			tmp_3;`}}

	for i, test := range tests {
		l := lexer.NewLexer([]byte(test.src))
		p := parser.NewParser()
		res, err := p.Parse(l)
		if err != nil {
			t.Log(err)
		}

		program, _ := res.(*ast.Program)
		_, err = checker.Checker(program)
		if err != nil {
			t.Log(err)
		}

		var b bytes.Buffer
		code := gen.Gen(program, &b)
		if err != nil {
			t.Log(err)
		}

		codeString := code.String()

		// remove spaces for comparison
		for _, rep := range []string{" ", "\n", "\t"} {
			codeString = strings.Replace(codeString, rep, "", -1)
			test.res = strings.Replace(test.res, rep, "", -1)
		}

		if codeString != test.res {
			t.Log(codeString)
			t.Fatalf("test [%d] failed", i)
		}
	}
}

func TestOutPut(T *testing.T) {

}
