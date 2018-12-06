package checker

// operations
const (
	PLUS    = "PLUS"
	EQUALS  = "EQUALS"
	ATMOST  = "ATMOST"
	ATLEAST = "ATLEAST"
	LESS    = "LESS"
	MORE    = "MORE"
	MINUS   = "MINUS"
	DIVIDE  = "DIVIDE"
	TIMES   = "TIMES"
	AND     = "AND"
	OR      = "OR"
)

// error types
const (
	INVALID_OPERATION_TYPE   = "INVALID_OPERATION_TYPE"
	VARIABLE_NOT_INITIALIZED = "VARIABLE_NOT_INITIALIZED"
	ALREADY_INITIALIZED      = "ALREADY_INITIALIZED"
	INCOMPATABLE_TYPES       = "INCOMPATABLE_TYPES"
	BAD_FUNCTION_CALL        = "BAD_FUNCTION_CALL"
	CONDITION_NOT_BOOL       = "CONDITION_NOT_BOOL"
	INVALID_RETURN_TYPE      = "INVALID_RETURN_TYPE"
	INCORRECT_ARGUMENT_COUNT = "INCORRECT_ARGUMENT_COUNT"
)

// variable types
const (
	INT_TYPE     = "INT_TYPE"
	STRING_TYPE  = "STRING_TYPE"
	BOOL_TYPE    = "BOOL_TYPE"
	NOTHING_TYPE = "NOTHING_TYPE"
)

type Signature struct {
	Return string
	Params []string // list of types
}

type Methods map[string]bool

// type methods
var TypeTable = map[string]Methods{
	INT_TYPE:    map[string]bool{"+": true, "-": true, "<=": true, "==": true, ">=": true, "!=": true, "*": true, "/": true},
	STRING_TYPE: map[string]bool{"+": true},
	BOOL_TYPE:   map[string]bool{"and": true, "or": true}}

type Environment struct {
	Vals  map[string]string    // map identifier to type
	Funcs map[string]Signature // map function name to return type
	Types map[string]bool      // track valid types
}

func NewEnvironment() Environment {
	return Environment{Vals: map[string]string{}, Funcs: map[string]Signature{}, Types: map[string]bool{}}
}

func MethodExist(kind, method string) bool {
	methods, ok := TypeTable[kind]
	if !ok {
		return false
	}

	_, ok = methods[method]
	return ok
}

func (e *Environment) Set(name, kind string) {
	e.Vals[name] = kind
}

func (e *Environment) GetFunctionSignature(name string) (Signature, bool) {
	kind, ok := e.Funcs[name]
	return kind, ok
}

func (e *Environment) Get(name string) (string, bool) {
	kind, ok := e.Vals[name]
	return kind, ok
}

func (e *Environment) IdentExist(kind string) bool {
	_, ok := e.Vals[kind]
	return ok
}

func (e *Environment) TypeExist(kind string) bool {
	_, ok := e.Types[kind]
	return ok
}
