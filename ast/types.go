package ast

import "github.com/Lebonesco/go-compiler/token"

// every element is an Attrib
type Attrib interface{}

// root node
type Program struct {
	Statements []Statement `json:"statements"`
	Functions  []Statement `json:"functions"`
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
	Token *token.Token `json:"-"`
	Left  Identifier   `json:"left"`
	Right Expression   `json:"right"`
}

type FunctionStatement struct {
	Token      *token.Token    `json:"-"`
	Name       string          `json:"name"`
	Parameters []FormalArg     `json:"params"`
	Body       *BlockStatement `json:"body"`
	Return     string          `json:"return"`
}

type FormalArg struct {
	Arg  string `json:"arg"`
	Type string `json:"type"`
}

type ForStatement struct {
	Token          *token.Token    `json:"-"`
	Condition      Expression      `json:"condition"`
	BlockStatement *BlockStatement `json:"block"`
}

type ReturnStatement struct {
	Token       *token.Token `json:"-"`
	ReturnValue Expression   `json:"return"`
}

type BlockStatement struct {
	Token      *token.Token `json:"-"`
	Statements []Statement  `json:"statements"`
}

type IfStatement struct {
	Token       *token.Token    `json:"-"`
	Condition   Expression      `json:"condition"`
	Block       *BlockStatement `json:"block"`
	Alternative *BlockStatement `json:"alternative"`
}

type ExpressionStatement struct {
	Token      *token.Token `json:"-"`
	Expression Expression   `json:"statement"`
}

type InitStatement struct {
	Token    *token.Token `json:"-"`
	Expr     Expression   `json:"expression"`
	Location string       `json:"location"`
}

// Expressions
type Identifier struct {
	Token *token.Token `json:"-"`
	Value string       `json:"value"`
}

type Boolean struct {
	Token *token.Token `json:"-"`
	Value bool         `json:"value"`
}

type IntegerLiteral struct {
	Token *token.Token `json:"-"`
	Value string       `json:"value"`
}

type StringLiteral struct {
	Token *token.Token `json:"-"`
	Value string       `json:"value"`
}

type InfixExpression struct {
	Token    *token.Token `json:"-"`
	Type     string       `json:"-"`
	Left     Expression   `json:"left"`
	Right    Expression   `json:"right"`
	Operator string       `json:"operator"`
}

type FunctionCall struct {
	Token *token.Token `json:"-"`
	Name  string       `json:"name"`
	Args  []Expression `json:"args"`
	Type  string       `json:"type"`
}
