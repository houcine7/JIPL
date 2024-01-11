package runtime

import (
	"fmt"

	"github.com/houcine7/JIPL/internal/debug"
	"github.com/houcine7/JIPL/internal/types"
)

var builtins = map[string]*types.BuiltIn{
	"out": {Fn: func(args ...types.ObjectJIPL) (types.ObjectJIPL, *debug.Error) {
		for _, arg := range args {
			fmt.Println(arg.ToString())
		}
		return nil, debug.NOERROR
	}},
	"length": {Fn: func(args ...types.ObjectJIPL) (types.ObjectJIPL, *debug.Error) {

		if len(args) != 1 {
			return nil, debug.NewError(fmt.Sprintf("the arguments of the length function should be exactly one instead got %d", len(args)))
		}

		switch t := args[0].(type) {
		case *types.String:
			return &types.Integer{Val: len(t.Val)}, debug.NOERROR
		default:
			return nil, debug.NewError(fmt.Sprintf("the argument of type %T doesn't have the length function", t))
		}
	}},
}
