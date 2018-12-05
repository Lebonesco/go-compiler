package checker

import (
	"errors"
	"github.com/Lebonesco/go-compiler/ast"
)

func Checker(program *ast.Program) (Environment, error) {
	env := NewEnvironment()
	_, err := checker(program, &env)
	return env, err
}

func checker(node ast.Node, env *Environment) (string, error) {
	switch node := node.(type) {
	// Statements
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.BlockStatement:
		return evalBlockStatement(node, env)
	case *ast.ReturnStatement:
		return evalReturnStatement(node, env)
	case *ast.IfStatement:
		return evalIfStatement(node, env)
	case *ast.ExpressionStatement:
		return evalExpressionStatement(node, env)
	case *ast.AssignStatement:
		return evalAssignStatement(node, env)
	case *ast.InitStatement:
		return evalInitStatement(node, env)
	// Expressions
	case *ast.InfixExpression:
		return evalInfixExpression(node, env)
	case *ast.IntegerLiteral:
		return evalInteger(node, env)
	case *ast.StringLiteral:
		return evalString(node, env)
	case *ast.Boolean:
		return evalBoolean(node, env)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.FunctionCall:
		return evalFunctionCall(node, env)
	}
	return "", nil
}

func evalProgram(p *ast.Program, env *Environment) (string, error) {
	for _, statement := range p.Statements {
		_, err := checker(statement, env)
		if err != nil {
			return "", err
		}
	}
	return "", nil
}

// Statements

func evalBlockStatement(node *ast.BlockStatement, env *Environment) (string, error) {
	for _, statement := range node.Statements {
		_, err := checker(statement, env)
		if err != nil {
			return "", err
		}
	}
	return "", nil
}

func evalReturnStatement(node *ast.ReturnStatement, env *Environment) (string, error) {
	res, err := checker(node.ReturnValue, env)
	return res, err
}

func evalIfStatement(node *ast.IfStatement, env *Environment) (string, error) {
	return "", nil
}

func evalExpressionStatement(node *ast.ExpressionStatement, env *Environment) (string, error) {
	_, err := checker(node.Expression, env)
	if err != nil {
		return "", nil
	}

	return "", nil
}

func evalInitStatement(node *ast.InitStatement, env *Environment) (string, error) {
	if env.IdentExist(node.Location) {
		return "", errors.New("ident already exist")
	}

	right, err := checker(node.Expr, env)
	if err != nil {
		return "", nil
	}

	env.Set(node.Location, right) // set ident type
	return "", nil
}

func evalAssignStatement(node *ast.AssignStatement, env *Environment) (string, error) {
	right, err := checker(node.Right, env)
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

func evalFunctionCall(node *ast.FunctionCall, env *Environment) (string, error) {
	var sig Signature
	var ok bool
	if sig, ok = env.GetFunctionSignature(node.Name); !ok {
		return "", errors.New("function not exist")
	}

	if len(node.Args) != len(sig.Params) {
		return "", errors.New("incorrect amount of arguments to function")
	}

	// check params
	for i, arg := range node.Args {
		res, err := checker(arg, env)
		if err != nil {
			return "", nil
		}

		if res != sig.Params[i] {
			return "", errors.New("incorrect argument type")
		}
	}

	return sig.Return, nil
}

func evalIdentifier(node *ast.Identifier, env *Environment) (string, error) {
	kind, _ := env.Get(node.Value)
	return kind, nil
}

func evalBoolean(node *ast.Boolean, env *Environment) (string, error) {
	return BOOL_TYPE, nil
}

func evalInteger(node *ast.IntegerLiteral, env *Environment) (string, error) {
	return INT_TYPE, nil
}

func evalString(node *ast.StringLiteral, env *Environment) (string, error) {
	return STRING_TYPE, nil
}

// Expressions

func evalInfixExpression(node *ast.InfixExpression, env *Environment) (string, error) {
	left, err := checker(node.Left, env)
	if err != nil {
		return left, err
	}

	right, err := checker(node.Right, env)
	if err != nil {
		return right, err
	}

	if left != right {
		return "", errors.New("incorrect types for operation")
	}

	// check if method exists in type
	//methods := map[string]string{"+": PLUS, "-": MINUS, "==": EQUALS, "<": LESS, ">": MORE, ">=": ATLEAST,
	//	"<=": ATMOST, "*": TIMES, "/": DIVIDE, "or": OR, "and": AND}

	// if _, ok := env.GetClassMethod(left, methods[node.Operator]); !ok {
	// 	return environment.Variable{}, createError(METHOD_NOT_EXIST, "method %s not exist in class %s on line %d", methods[node.Operator], obj.Type, node.Token.Pos.Line)
	// }

	// switch node.Operator { // evaluates to a bool
	// case "<", ">", "<=", ">=", "==", "!=", "and", "or":
	// 	node.Type = string(left.Type)
	// 	return environment.Variable{Type: environment.BOOL_CLASS}, nil
	// }

	return "", nil
}
