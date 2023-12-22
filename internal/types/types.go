package types

import (
	"bytes"
	"fmt"

	ast "github.com/houcine7/JIPL/internal/AST"
)

type TypeObj string

type ObjectJIPL interface {
	GetType() TypeObj
	ToString() string
}

type Integer struct {
	Val int
}


type String struct {
	Val string
}
type Boolean struct {
	Val bool
}

type Undefined struct{}

type Return struct {
	Val ObjectJIPL
}

type Function struct {
	Name string
	Params []*ast.Identifier
	Body   *ast.BlockStm
	Ctx    *Context
}


type Context struct {
	Store map[string]ObjectJIPL
	Outer *Context // for nested scopes
}

type BuiltIn struct {
	Fn func(args ...ObjectJIPL) ObjectJIPL
}

func (bi *BuiltIn) GetType() TypeObj {
	return T_BUILTIN
}
func (bi *BuiltIn) ToString() string {
	return "builtin function"
}

func NewContextWithOuter(outer *Context) *Context {
	ctx := NewContext()
	ctx.Outer = outer
	return ctx
}


func (ctx *Context) Get(key string) (ObjectJIPL, bool) {
	val, ok := ctx.Store[key]
	if !ok && ctx.Outer != nil {
		// recursively search for the key 
		// in the outer context
		return ctx.Outer.Get(key)
	}
	return val, ok
}

func (ctx *Context) Set(key string, val ObjectJIPL) ObjectJIPL {
	ctx.Store[key] = val
	return val
}


func NewContext() *Context {
	return &Context{
		Store: make(map[string]ObjectJIPL),
		Outer: nil,
	}
}

// implementing OBjectJIPL interface by supported types
func (fn *Function) GetType() TypeObj {
	return T_FUNCTION
}
func (fn *Function) ToString() string {
	var bf bytes.Buffer
	bf.WriteString("function ")
	bf.WriteString(fn.Name)
	bf.WriteString("(")
	for idx, param := range fn.Params {
		bf.WriteString(param.Value)
		if idx != len(fn.Params)-1 {
			bf.WriteString(",")
		}
	}
	bf.WriteString(fn.Body.ToString())

	return bf.String()
}

func (ret *Return) ToString() string {
	return ret.Val.ToString()
}
func (ret *Return) GetType() TypeObj {
	return T_RETURN
}

func (und *Undefined) ToString() string {
	return "undefined"
}

func (und *Undefined) GetType() TypeObj {
	return T_UNDEFINED
}

func (boolObj *Boolean) ToString() string {
	return fmt.Sprintf("%t", boolObj.Val)
}

func (boolObj *Boolean) GetType() TypeObj {
	return T_BOOLEAN
}

func (intObj *Integer) ToString() string {
	return fmt.Sprintf("%d", intObj.Val)
}

func (intObj *Integer) GetType() TypeObj {
	return T_INTEGER
}



func BoolToObJIPL(bl bool) ObjectJIPL{
	if bl {
		return TRUE 
	}else{
		return FALSE
	}
}

func (str *String) GetType() TypeObj {
	return T_STRING
}
func (str *String) ToString() string {
	return str.Val
}


// cte of types
const (
	T_INTEGER  = "INTEGER"
	T_BOOLEAN   = "BOOLEAN"
	T_UNDEFINED = "UNDEFINED"
	T_RETURN   = "RETURN"
	T_FUNCTION = "FUNCTION"
	T_STRING   = "STRING"
	T_BUILTIN = "BUILTIN"
)

var (
	TRUE      = &Boolean{Val: true}
	FALSE     = &Boolean{Val: false}
	UNDEFIEND = &Undefined{}
)
