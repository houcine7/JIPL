package types

import "fmt"

type TypeObj string

type ObjectJIPL interface {
	GetType() TypeObj
	ToString() string
}

type Integer struct {
	Val int
}

type Boolean struct {
	Val bool
}

type Undefined struct{}

type Return struct {
	Val ObjectJIPL
}

type Context struct {
	Store map[string]ObjectJIPL
	Outer *Context // for nested scopes
}

func (ctx *Context) Get(key string) (ObjectJIPL, bool) {
	val, ok := ctx.Store[key]
	if !ok && ctx.Outer != nil {
		return ctx.Outer.Get(key)
	}
	return val, ok
}

func (ctx *Context) Set(key string, val ObjectJIPL) ObjectJIPL {
	ctx.Store[key] = val
	return val
}


func NewContext() *Context {
	return &Context{Store: make(map[string]ObjectJIPL)}
}

// implementing OBjectJIPL interface by supported types
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


// cte of types
const (
	T_INTEGER  = "INTEGER"
	T_BOOLEAN   = "BOOLEAN"
	T_UNDEFINED = "UNDEFINED"
	T_RETURN   = "RETURN"
)

var (
	TRUE      = &Boolean{Val: true}
	FALSE     = &Boolean{Val: false}
	UNDEFIEND = &Undefined{}
)
