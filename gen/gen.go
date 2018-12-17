package gen

import (
	"bytes"
	"fmt"
	"github.com/Lebonesco/go-compiler/ast"
	. "github.com/Lebonesco/go-compiler/checker"
	"strings"
)

var TMP_COUNT int

func write(b *bytes.Buffer, code string, args ...interface{}) {
	b.WriteString(fmt.Sprintf(code, args...))
}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func freshTemp() string {
	TMP_COUNT += 1
	return fmt.Sprintf("tmp_%d", TMP_COUNT)
}

func GenWrapper(p *ast.Program) bytes.Buffer {
	TMP_COUNT = 0
	var b bytes.Buffer
	gen(p, &b)
	return b
}

func gen(node ast.Node, b *bytes.Buffer) string {
	switch node := node.(type) {
	// Statements
	case *ast.Program:
		return genProgram(node, b)
	case *ast.BlockStatement:
		return genBlockStatement(node, b)
	case *ast.ReturnStatement:
		return genReturnStatement(node, b)
	case *ast.FunctionStatement:
		return genFunctionStatement(node, b)
	case *ast.IfStatement:
		return genIfStatement(node, b)
	case *ast.ExpressionStatement:
		return genExpressionStatement(node, b)
	case *ast.AssignStatement:
		return genAssignStatement(node, b)
	case *ast.InitStatement:
		return genInitStatement(node, b)
	// // Expressions
	case *ast.InfixExpression:
		return genInfixExpression(node, b)
	case *ast.IntegerLiteral:
		return genInteger(node, b)
	case *ast.StringLiteral:
		return genString(node, b)
	case *ast.Boolean:
		return genBoolean(node, b)
	case *ast.Identifier:
		return genIdentifier(node, b)
	case *ast.FunctionCall:
		return genFunctionCall(node, b)
	}
	return ""
}

func genProgram(node *ast.Program, b *bytes.Buffer) string {
	write(b, "#include <string>\n#include <iostream>\n#include \"Builtins.cpp\"\n\n")

	for _, funcs := range node.Functions {
		gen(funcs, b)
	}

	write(b, "int main() {\n")
	for _, stmt := range node.Statements {
		gen(stmt, b)
	}
	write(b, "return 0;\n}")
	return ""
}

func genBlockStatement(node *ast.BlockStatement, b *bytes.Buffer) string {
	for _, stmt := range node.Statements {
		gen(stmt, b)
	}
	return ""
}

func genExpressionStatement(node *ast.ExpressionStatement, b *bytes.Buffer) string {
	value := gen(node.Expression, b)
	write(b, "%s;\n", value)
	return ""
}

func genAssignStatement(node *ast.AssignStatement, b *bytes.Buffer) string {
	right := gen(node.Right, b)
	// get left type
	write(b, "%s = %s;\n", node.Left.Value, right)
	return ""
}

func genInitStatement(node *ast.InitStatement, b *bytes.Buffer) string {
	right := gen(node.Expr, b)
	kind, _ := GetIdentType(node.Location)
	write(b, "%s %s = %s;\n", kind, node.Location, right)
	return ""
}

func genReturnStatement(node *ast.ReturnStatement, b *bytes.Buffer) string {
	value := gen(node.ReturnValue, b)
	write(b, "return %s;\n", value)
	return ""
}

func genFunctionStatement(node *ast.FunctionStatement, b *bytes.Buffer) string {
	if IsBuiltin(node.Name) {
		panic("built in function")
	}

	write(b, "%s %s(", node.Return, node.Name)

	for i, arg := range node.Parameters {
		write(b, "%s %s", arg.Type, arg.Arg)
		if i != len(node.Parameters)-1 {
			write(b, ",")
		}
	}
	write(b, ") {\n")
	gen(node.Body, b)
	write(b, "}\n\n")
	return ""
}

func genIfStatement(node *ast.IfStatement, b *bytes.Buffer) string {
	cond := gen(node.Condition, b)
	write(b, "if (\"true\" == %s.val) {\n", cond)
	gen(node.Block, b)
	write(b, "} else {\n")
	gen(node.Alternative, b)
	write(b, "}\n\n")
	return ""
}

func genInteger(node *ast.IntegerLiteral, b *bytes.Buffer) string {
	tmp := freshTemp()
	write(b, "Int %s = Int(%s);\n", tmp, string(node.Token.Lit))
	return tmp
}

func genString(node *ast.StringLiteral, b *bytes.Buffer) string {
	tmp := freshTemp()
	str := string(node.Token.Lit)
	str = strings.Replace(str, `\`, "\\", -1)

	write(b, "String %s = String(%s);\n", tmp, str)
	return tmp
}

func genBoolean(node *ast.Boolean, b *bytes.Buffer) string {
	if node.Value {
		return "Bool(\"true\")"
	} else {
		return "Bool(\"false\")"
	}
	return ""
}

func genIdentifier(node *ast.Identifier, b *bytes.Buffer) string {
	return node.Value
}

func genInfixExpression(node *ast.InfixExpression, b *bytes.Buffer) string {
	left := gen(node.Left, b)
	right := gen(node.Right, b)
	kind := node.Type

	tmp := freshTemp()
	methods := map[string]string{"+": PLUS, "-": MINUS, "==": EQUAL, "<": LT, ">": GT, "*": TIMES, "or": OR, "and": AND}

	method, _ := GetMethod(kind, methods[node.Operator])
	write(b, "%s %s = %s.%s(%s);\n", method.Return, tmp, left, methods[node.Operator], right)
	return tmp
}

func genFunctionCall(node *ast.FunctionCall, b *bytes.Buffer) string {
	var sig Signature
	args := make([]string, len(node.Args))
	// store expression tmp vars
	for i, arg := range node.Args {
		res := gen(arg, b)
		args[i] = res
	}

	tmp := freshTemp()
	if IsBuiltin(node.Name) {
		sig, ok := GetMethod(node.Type, node.Name)
		if !ok {
			panic("no builtin function")
		}

		write(b, "%s %s = %s.%s(", sig.Return, tmp, args[0], node.Name)
	} else {
		sig, _ = GetFunctionSignature(node.Name)
		write(b, "%s %s = %s(", sig.Return, tmp, node.Name)
		for i, arg := range args {
			write(b, arg)
			if i != len(args)-1 {
				write(b, ",")
			}
		}
	}

	write(b, ");")
	return tmp
}
