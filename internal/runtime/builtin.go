package runtime

import (
	"fmt"

	"github.com/houcine7/JIPL/internal/types"
)


var builtins = map[string]*types.BuiltIn{
	"out" : &types.BuiltIn{ Fn: func(args ...types.ObjectJIPL) types.ObjectJIPL {
		for _, arg := range args {
			fmt.Println(arg.ToString())
		}
		return nil
	}},}