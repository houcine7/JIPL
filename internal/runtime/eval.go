package runtime

import (
	ast "github.com/houcine7/JIPL/internal/AST"
	"github.com/houcine7/JIPL/internal/types"
)

func Eval(node ast.Node) types.ObjectJIPL {
	switch node := node.(type) {
	case *ast.Program:
		return evalAllStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		return &types.Integer{Val: node.Value}
	case *ast.BooleanExp:
		return &types.Boolean{Val: node.Value}
	default:
		return nil
	}
}

func evalAllStatements(stms []ast.Statement) types.ObjectJIPL {
	var resrult types.ObjectJIPL

	for _, stm := range stms {
		resrult = Eval(stm)
	}

	return resrult
}
