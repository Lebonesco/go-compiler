package main

import (
	"bytes"
	"github.com/Lebonesco/go-compiler/ast"
	"github.com/Lebonesco/go-compiler/checker"
	"github.com/Lebonesco/go-compiler/gen"
	"github.com/Lebonesco/go-compiler/lexer"
	"github.com/Lebonesco/go-compiler/parser"
	"testing"
)

func TestGen(t *testing.T) {
	tests := []struct {
		src string
		res string
	}{
		{
			src: ``,
			res: ``}}

	for i, test := range tests {
		l := lexer.NewLexer([]byte(input))
		p := parser.NewParser()
		res, err := p.Parse(l)
		if err != nil {
			return t.Log(err)
		}

		program, _ := res.(*ast.Program)
		env, err = checker.Checker(program)
		if err != nil {
			return t.Log(err)
		}

		var b bytes.Buffer
		code, err := gen.Gen(program, &b, env)
		if err != nil {
			t.Log(err)
		}

		codeString := code.String()

		// remove spaces for comparison
		for _, rep := range []string{" ", "\n", "\t"} {
			code = strings.Replace(code, rep, "", -1)
			test.res = strings.Replace(test.res, rep, "", -1)
		}

		if code != test.res {
			t.Fatalf("test [%d] failed at character", i)
		}

	}

}

func TestOutPut(T *testing.T) {

}
