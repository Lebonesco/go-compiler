package checker

import (
	"errors"
	"fmt"
	"github.com/Lebonesco/go-compiler/ast"
	"reflect"
)

func Checker(program *ast.Program) error {
	env = NewEnvironment() // reset environment
	_, err := checker(program)
	return err
}

func checker(node ast.Node) (string, error) {
	switch node := node.(type) {
	// Statements
	case *ast.Program:
		return evalProgram(node)
	case *ast.BlockStatement:
		return evalBlockStatement(node)
	case *ast.ReturnStatement:
		return evalReturnStatement(node)
	case *ast.IfStatement:
		return evalIfStatement(node)
	case *ast.ExpressionStatement:
		return evalExpressionStatement(node)
	case *ast.AssignStatement:
		return evalAssignStatement(node)
	case *ast.InitStatement:
		return evalInitStatement(node)
	case *ast.FunctionStatement:
		return evalFunctionStatement(node)
	// Expressions
	case *ast.InfixExpression:
		return evalInfixExpression(node)
	case *ast.IntegerLiteral:
		return evalInteger(node)
	case *ast.StringLiteral:
		return evalString(node)
	case *ast.Boolean:
		return evalBoolean(node)
	case *ast.Identifier:
		return evalIdentifier(node)
	case *ast.FunctionCall:
		return evalFunctionCall(node)
	}
	return "", nil
}

func evalProgram(p *ast.Program) (string, error) {
	for _, function := range p.Functions {
		_, err := checker(function)
		if err != nil {
			return "", err
		}
	}

	for _, statement := range p.Statements {
		_, err := checker(statement)
		if err != nil {
			return "", err
		}
	}
	return "", nil
}

// Statements
func evalBlockStatement(node *ast.BlockStatement) (string, error) {
	for _, statement := range node.Statements {
		result, err := checker(statement)
		if err != nil {
			return "", err
		}
		if reflect.TypeOf(statement) == reflect.TypeOf(&ast.ReturnStatement{}) {
			return result, nil
		}
	}
	return NOTHING_TYPE, nil
}

func evalReturnStatement(node *ast.ReturnStatement) (string, error) {
	res, err := checker(node.ReturnValue)
	return res, err
}

func evalIfStatement(node *ast.IfStatement) (string, error) {
	cond, _ := checker(node.Condition)
	if cond != BOOL_TYPE {
		return "", errors.New("condition not bool type")
	}

	checker(node.Block)
	checker(node.Alternative)
	return "", nil
}

func evalExpressionStatement(node *ast.ExpressionStatement) (string, error) {
	_, err := checker(node.Expression)
	if err != nil {
		return "", err
	}

	return "", nil
}

func evalInitStatement(node *ast.InitStatement) (string, error) {
	if env.IdentExist(node.Location) {
		return "", errors.New("ident already exist")
	}

	right, err := checker(node.Expr)
	if err != nil {
		return "", err
	}

	env.Set(node.Location, right) // set ident type
	return "", nil
}

func evalAssignStatement(node *ast.AssignStatement) (string, error) {
	right, err := checker(node.Right)
	if err != nil {
		return "", nil
	}

	if kind, ok := env.Get(node.Left.Value); ok {
		if kind != right {
			return "", errors.New("invalid type assignment")
		}
	} else {
		return "", errors.New("ident not exist")
	}
	return "", nil
}

func evalFunctionStatement(node *ast.FunctionStatement) (string, error) {
	var params []string
	for _, param := range node.Parameters {
		env.Set(param.Arg, param.Type) // set params into scope
		params = append(params, param.Type)
	}

	res, err := checker(node.Body)
	if err != nil {
		return "", err
	}
	// check if correct return type
	if res != node.Return {
		return "", errors.New("incorrect return type")
	}

	SetFunctionSignature(node.Name, Signature{node.Return, params})
	return "", nil
}

// Expressions

func evalFunctionCall(node *ast.FunctionCall) (string, error) {
	if IsBuiltin(node.Name) {
		res, err := checker(node.Args[0])
		if err != nil {
			return "", err
		}
		node.Type = res
		return NOTHING_TYPE, nil
	}

	var sig Signature
	var ok bool
	if sig, ok = GetFunctionSignature(node.Name); !ok {
		return "", errors.New("function not exist")
	}

	if len(node.Args) != len(sig.Params) {
		return "", errors.New("incorrect amount of arguments to function")
	}

	// check params
	for i, arg := range node.Args {
		res, err := checker(arg)
		if err != nil {
			return "", errors.New(err.Error())
		}

		if res != sig.Params[i] {
			return "", errors.New("incorrect argument type")
		}
	}

	return sig.Return, nil
}

func evalIdentifier(node *ast.Identifier) (string, error) {
	kind, _ := env.Get(node.Value)
	return kind, nil
}

func evalBoolean(node *ast.Boolean) (string, error) {
	return BOOL_TYPE, nil
}

func evalInteger(node *ast.IntegerLiteral) (string, error) {
	return INT_TYPE, nil
}

func evalString(node *ast.StringLiteral) (string, error) {
	return STRING_TYPE, nil
}

func evalInfixExpression(node *ast.InfixExpression) (string, error) {
	left, err := checker(node.Left)
	if err != nil {
		return left, err
	}

	right, err := checker(node.Right)
	if err != nil {
		return right, err
	}

	if left != right {
		return "", errors.New("incorrect types for operation")
	}

	node.Type = left // set type for code generation

	methods := map[string]string{"+": PLUS, "-": MINUS, "==": EQUAL, "<": LT, ">": GT, "*": TIMES, "or": OR, "and": AND}

	if !MethodExist(left, methods[node.Operator]) {
		return NOTHING_TYPE, errors.New(fmt.Sprintf("method %s not exist for type %s", methods[node.Operator], left))
	}

	for _, opr := range []string{"<=", "<", ">=", ">", "or", "and"} {
		if node.Operator == opr {
			return BOOL_TYPE, nil
		}
	}

	return left, nil
}
