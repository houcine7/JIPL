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

// implementing OBjectJIPL interface by supporeted types
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

// cte of types
const (
	T_INTEGER  = "INTEGER"
	T_BOOLEAN   = "BOOLEAN"
	T_UNDEFINED = "UNDEFINED"
)

var (
	TRUE      = &Boolean{Val: true}
	FALSE     = &Boolean{Val: false}
	UNDEFIEND = &Undefined{}
)
