package checker

// operations
const (
	PLUS   = "PLUS"
	EQUAL  = "EQUAL"
	LT     = "LESS"
	GT     = "MORE"
	MINUS  = "MINUS"
	DIVIDE = "DIVIDE"
	TIMES  = "TIMES"
	AND    = "AND"
	OR     = "OR"
	PRINT  = "PRINT"
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
	INT_TYPE     = "Int"
	STRING_TYPE  = "String"
	BOOL_TYPE    = "Bool"
	NOTHING_TYPE = "Nothing"
)

type Signature struct {
	Return string
	Params []string // list of types
}

type Methods map[string]Signature

// type methods
var TypeTable = map[string]Methods{
	INT_TYPE: {
		PLUS:  {INT_TYPE, []string{INT_TYPE}},
		MINUS: {INT_TYPE, []string{INT_TYPE}},
		TIMES: {INT_TYPE, []string{INT_TYPE}},
		LT:    {BOOL_TYPE, []string{INT_TYPE}},
		GT:    {INT_TYPE, []string{INT_TYPE}},
		EQUAL: {INT_TYPE, []string{INT_TYPE}},
		PRINT: {NOTHING_TYPE, []string{}}},
	STRING_TYPE: {
		PLUS:  {STRING_TYPE, []string{STRING_TYPE}},
		PRINT: {NOTHING_TYPE, []string{}}},
	BOOL_TYPE: {
		AND:   {BOOL_TYPE, []string{BOOL_TYPE}},
		OR:    {BOOL_TYPE, []string{BOOL_TYPE}},
		PRINT: {NOTHING_TYPE, []string{}}}}

type Environment struct {
	Vals  map[string]string    // map identifier to type
	Funcs map[string]Signature // map function name to return type
	Types map[string]bool      // track valid types
}

var env Environment // set global

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

func GetMethod(kind, method string) (Signature, bool) {
	methods, ok := TypeTable[kind]
	if !ok {
		return Signature{}, false
	}

	sig, ok := methods[method]
	return sig, ok
}

func (e *Environment) Set(name, kind string) {
	e.Vals[name] = kind
}

func SetFunctionSignature(name string, sig Signature) {
	env.Funcs[name] = sig
}

func GetFunctionSignature(name string) (Signature, bool) {
	kind, ok := env.Funcs[name]
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

func GetIdentType(name string) (string, bool) {
	kind, ok := env.Vals[name]
	return kind, ok
}

func (e *Environment) TypeExist(kind string) bool {
	_, ok := e.Types[kind]
	return ok
}
