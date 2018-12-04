package ast

import "github.com/Lebonesco/go-compiler/token"

// every element is an Attrib
type Attrib interface{}

// root node
type Program struct {
	Statements []Statement
}

// base interface
type Node interface {
	TokenLiteral() string
}

// all statement nodes
type Statement interface {
	Node
	statementNode()
}

// all expression nodes
type Expression interface {
	Node
	expressionNode()
}

// Statements
type AssignStatement struct {
	Token token.Token
	Left  *Identifier
	Right Expression
}

type FunctionStatement struct {
	Token      token.Token
	Name       string
	Parameters []FormalArg
	Body       *BlockStatement
	Return     string
}

type FormalArg struct {
	Arg  string
	Type string
}

type ForStatement struct {
	Token          token.Token
	Condition      Expression
	BlockStatement *BlockStatement
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

type IfStatement struct {
	Token       token.Token
	Condition   Expression
	Block       *BlockStatement
	Alternative *BlockStatement
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

type InitStatement struct {
	Token    token.Token
	Expr     Expression
	Location string
}

// Expressions
type Identifier struct {
	Token token.Token
	Value string
}

type Boolean struct {
	Token token.Token
	Value bool
}

type IntegerLiteral struct {
	Token token.Token
	Value string
}

type StringLiteral struct {
	Token token.Token
	Value string
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Right    Expression
	Operator string
}

type FunctionCall struct {
	Token token.Token
	Name  string
	Args  []Expression
}
