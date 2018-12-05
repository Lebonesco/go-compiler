package main

import (
	"encoding/json"
	"fmt"
	"github.com/Lebonesco/go-compiler/ast"
	"github.com/Lebonesco/go-compiler/lexer"
	"github.com/Lebonesco/go-compiler/parser"
	"reflect"
	"testing"
)

func TestAST(t *testing.T) {
	const input = `
			/* comment should not be scanned */
			let five = "test";
			let ten = 10;

			func add() {
				return x + y;
			};

			let result = 4;  
			5 <= 10;

			if (5 == 10) {
				print();
			} else {
				print();
			}

			10 == 10;
			`

	out := &ast.Program{
		Statements: []ast.Statement{
			ast.InitStatement{Expr: ast.StringLiteral{Value: "\"test\""}, Location: "five"},
			ast.InitStatement{Expr: ast.StringLiteral{Value: "10"}, Location: "ten"},
			ast.FunctionStatement{Name: "add", Parameters: []ast.FormalArg{}, Body: &ast.BlockStatement{
				Statements: []ast.Statement{ast.ReturnStatement{ReturnValue: ast.InfixExpression{Left: ast.StringLiteral{Value: "x"}, Operator: "+", Right: ast.StringLiteral{Value: "y"}}}}}},
			ast.InitStatement{Expr: ast.IntegerLiteral{Value: "4"}, Location: "result"},
			ast.ExpressionStatement{Expression: ast.InfixExpression{Left: ast.IntegerLiteral{Value: "5"}, Operator: "<=", Right: ast.IntegerLiteral{Value: "10"}}},
			ast.IfStatement{Condition: ast.InfixExpression{Left: ast.IntegerLiteral{Value: "5"}, Operator: "==", Right: ast.IntegerLiteral{Value: "10"}},
				Block:       &ast.BlockStatement{Statements: []ast.Statement{ast.ExpressionStatement{Expression: ast.FunctionCall{Name: "print", Args: []ast.Expression{}}}}},
				Alternative: &ast.BlockStatement{Statements: []ast.Statement{ast.ExpressionStatement{Expression: ast.FunctionCall{Name: "print", Args: []ast.Expression{}}}}}},
			ast.ExpressionStatement{Expression: ast.InfixExpression{Left: ast.IntegerLiteral{Value: "10"}, Operator: "==", Right: ast.IntegerLiteral{Value: "10"}}}}}

	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	res, err := p.Parse(l)
	if err != nil {
		t.Fatalf(err.Error())
	}

	program, ok := res.(*ast.Program)
	if !ok {
		t.Fatalf("res not *ast.Program, got=%T", res)
	}
	js, _ := json.MarshalIndent(program, "", "    ")
	jsOut, _ := json.MarshalIndent(out, "", "    ")

	if !reflect.DeepEqual(js, jsOut) {
		fmt.Printf("\n%s\n", js)
		fmt.Println("****************************")
		fmt.Printf("\n%s\n", jsOut)

		str1 := string(js)
		str2 := string(jsOut)

		for i := 1; i <= len(str1); i++ {
			if str1[i] != str2[i] {
				fmt.Println(str1[1:i], str2[1:i], i)
				break
			}
		}

		t.Fatalf("Wrong AST")
	}

}
