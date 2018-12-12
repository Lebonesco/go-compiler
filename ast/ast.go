package ast

import (
	"fmt"
	"github.com/Lebonesco/go-compiler/token"
)

// interface methods

func (p Program) TokenLiteral() string { return "Program" }

// Statements
func (ls AssignStatement) statementNode()       {}
func (ls AssignStatement) TokenLiteral() string { return "AssignStatement" }

func (rs ReturnStatement) statementNode()       {}
func (rs ReturnStatement) TokenLiteral() string { return "ReturnStatement" }

func (es ExpressionStatement) statementNode()       {}
func (es ExpressionStatement) TokenLiteral() string { return "ExpressionStatement" }

func (is IfStatement) statementNode()       {}
func (is IfStatement) TokenLiteral() string { return "IfStatement" }

func (bs BlockStatement) statementNode()       {}
func (bs BlockStatement) TokenLiteral() string { return "BlockStatement" }

func (is InitStatement) statementNode()       {}
func (is InitStatement) TokenLiteral() string { return "InitStatement" }

func (fs FunctionStatement) statementNode()       {}
func (fs FunctionStatement) TokenLiteral() string { return "FunctionStatement" }

// Expressions
func (i Identifier) expressionNode()      {}
func (i Identifier) TokenLiteral() string { return string(i.Token.Lit) }

func (sl StringLiteral) expressionNode()      {}
func (sl StringLiteral) TokenLiteral() string { return string(sl.Token.Lit) }

func (b Boolean) expressionNode()      {}
func (b Boolean) TokenLiteral() string { return string(b.Token.Lit) }

func (il IntegerLiteral) expressionNode()      {}
func (il IntegerLiteral) TokenLiteral() string { return string(il.Token.Lit) }

func (oe InfixExpression) expressionNode()      {}
func (oe InfixExpression) TokenLiteral() string { return string(oe.Token.Lit) }

func (fc FunctionCall) expressionNode()      {}
func (fc FunctionCall) TokenLiteral() string { return string(fc.Token.Lit) }

func Error(fun, expected, v string, got interface{}) error {
	return fmt.Errorf("AST construction error: In function: %s, expected %s for %s. got=%T", fun, expected, v, got)
}

// AST Constructors

func NewProgram(funcs, stmts Attrib) (*Program, error) {
	s, ok := stmts.([]Statement)
	if !ok {
		return nil, Error("NewProgram", "[]Statement", "stmts", stmts)
	}

	f, ok := funcs.([]Statement)
	if !ok {
		return nil, Error("NewProgram", "[]Statement", "funcs", funcs)
	}

	return &Program{Functions: f, Statements: s}, nil
}

func NewStatementList() ([]Statement, error) {
	return []Statement{}, nil
}

func AppendStatement(stmtList, stmt Attrib) ([]Statement, error) {
	s, ok := stmt.(Statement)
	if !ok {
		return nil, Error("AppendStatement", "Statement", "stmt", stmt)
	}
	return append(stmtList.([]Statement), s), nil
}

func NewAssignStatement(left, right Attrib) (Statement, error) {
	l, ok := left.(*token.Token)
	if !ok {
		return nil, Error("NewAssignStatement", "Identifier", "left", left)
	}

	r, ok := right.(Expression)
	if !ok {
		return nil, Error("NewAssignStatement", "Expression", "right", right)
	}

	return &AssignStatement{Left: Identifier{Value: string(l.Lit)}, Right: r}, nil
}

func NewExpressionStatement(expr Attrib) (Statement, error) {
	e, ok := expr.(Expression)
	if !ok {
		return nil, Error("NewExpressionStatement", "Expression", "expr", expr)
	}
	return &ExpressionStatement{Expression: e}, nil
}

func NewBlockStatement(stmts Attrib) (*BlockStatement, error) {
	s, ok := stmts.([]Statement)
	if !ok {
		return nil, Error("NewBlockStatement", "[]Statement", "stmts", stmts)
	}

	return &BlockStatement{Statements: s}, nil
}

func NewFunctionStatement(name, args, ret, block Attrib) (Statement, error) {
	n, ok := name.(*token.Token)
	if !ok {
		return nil, Error("NewFunctionStatement", "*token.Token", "name", name)
	}

	b, ok := block.(*BlockStatement)
	if !ok {
		return nil, Error("NewFunctionStatement", "BlockStatement", "block", block)
	}

	a := []FormalArg{}
	if args != nil {
		a, ok = args.([]FormalArg)
		if !ok {
			return nil, Error("NewFunctionStatement", "[]FormalArg", "args", args)
		}
	}

	r, ok := ret.(*token.Token)
	if !ok {
		// bad
	}

	return &FunctionStatement{Name: string(n.Lit), Body: b, Parameters: a, Return: string(r.Lit)}, nil
}

func NewIfStatement(cond, cons, alt Attrib) (Statement, error) {
	c, ok := cond.(Expression)
	if !ok {
		return nil, fmt.Errorf("invalid type of cond. got=%T", cond)
	}

	cs, ok := cons.(*BlockStatement)
	if !ok {
		return nil, fmt.Errorf("invalid type of cons. got=%T", cons)
	}

	a, ok := alt.(*BlockStatement)
	if !ok {
		return nil, fmt.Errorf("invalid type of alt. got=%T", alt)
	}

	return &IfStatement{Condition: c, Block: cs, Alternative: a}, nil
}

func NewInfixExpression(left, right, oper Attrib) (Expression, error) {
	l, ok := left.(Expression)
	if !ok {
		fmt.Println(left)
		return nil, Error("NewInfixExpression", "Expression", "left", left)
	}

	o, ok := oper.(*token.Token)
	if !ok {
		return nil, Error("NewInfixExpression", "*token.Token", "oper", oper)
	}

	r, ok := right.(Expression)
	if !ok {
		return nil, Error("NewInfixExpression", "Expression", "right", right)
	}

	return &InfixExpression{Left: l, Operator: string(o.Lit), Right: r, Token: o}, nil
}

func NewIntegerLiteral(integer Attrib) (Expression, error) {
	intLit, ok := integer.(*token.Token)
	if !ok {
		return nil, Error("NewIntegerLiteral", "*token.Token", "integer", integer)
	}

	return &IntegerLiteral{Token: intLit, Value: string(intLit.Lit)}, nil
}

func NewStringLiteral(str Attrib) (Expression, error) {
	return &StringLiteral{Value: string(str.(*token.Token).Lit), Token: str.(*token.Token)}, nil
}

func NewIdentInit(ident, expr Attrib) (Statement, error) {
	e, ok := expr.(Expression)
	if !ok {
		return nil, Error("NewIdentInit", "Expression", "expr", expr)
	}

	return &InitStatement{Location: string(ident.(*token.Token).Lit), Token: ident.(*token.Token), Expr: e}, nil
}

func NewIdentExpression(ident Attrib) (*Identifier, error) {
	return &Identifier{Value: string(ident.(*token.Token).Lit), Token: ident.(*token.Token)}, nil
}

func NewBoolExpression(val Attrib) (Expression, error) {
	return &Boolean{Value: val.(bool)}, nil
}

func NewReturnStatement(exp Attrib) (Statement, error) {
	e, ok := exp.(Expression)
	if !ok {
		return nil, Error("NewReturnExpression", "Expression", "exp", exp)
	}
	return &ReturnStatement{ReturnValue: e}, nil
}

func NewFunctionCall(name, args Attrib) (Expression, error) {
	n, ok := name.(*token.Token)
	if !ok {
		return nil, fmt.Errorf("invalid type of name. got=%T", name)
	}

	a := []Expression{}
	if args != nil {
		var ok bool
		a, ok = args.([]Expression)
		if !ok {
			return nil, Error("NewFunctionCall", "[]Expression", "args", args)
		}
	}

	return &FunctionCall{Name: string(n.Lit), Args: a, Token: n}, nil
}

func NewFormalArg() ([]FormalArg, error) {
	return []FormalArg{}, nil
}

func AppendFormalArgs(args, arg, kind Attrib) ([]FormalArg, error) {
	as, ok := args.([]FormalArg)
	if !ok {
		return nil, Error("AppendFormalArgs", "[]FormalArg", "args", args)
	}

	a, ok := arg.(*token.Token)
	if !ok {
		return nil, Error("AppendFormalArgs", "*token.Token", "arg", arg)
	}

	k, ok := kind.(*token.Token)
	if !ok {
		return nil, fmt.Errorf("invalid type of kind. got=%T", kind)
	}

	return append(as, FormalArg{string(a.Lit), string(k.Lit)}), nil
}

func NewArg() ([]Expression, error) {
	return []Expression{}, nil
}

func AppendArgs(expr, args Attrib) ([]Expression, error) {
	as, ok := args.([]Expression)
	if !ok {
		return nil, Error("AppendFormalArgs", "[]FormalArgs", "args", args)
	}

	e, ok := expr.(Expression)
	if !ok {
		return nil, Error("AppendFormalArgs", "Expression", "expr", expr)
	}

	return append(as, e), nil
}
