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
	case *ast.PrefixExpression:
		operand := Eval(node.Right)
		return evalPrefixExpression(node.Operator, operand)
	case *ast.PostfixExpression:
		operand := Eval(node.Left)
		return evalPostfixExpression(node.Operator, operand)
	default:
		return nil
	}
}
	


func evalPostfixExpression(operator string, operand types.ObjectJIPL) types.ObjectJIPL{
	switch operator {
	case "--":
		return evalDecrementPostfix(operand)
	case "++":
		return evalIncrementPostfix(operand)
	default:
		return nil
	}
}

func evalIncrementPostfix(operand types.ObjectJIPL) types.ObjectJIPL {
	if operand.GetType() != types.T_ITNTEGER {
		return nil
	}
	intObj := operand.(*types.Integer)
	return &types.Integer{Val: intObj.Val+1}

}

func evalDecrementPostfix(operand types.ObjectJIPL) types.ObjectJIPL{
	if operand.GetType() != types.T_ITNTEGER {
		return nil
	}
	intObj := operand.(*types.Integer)
	return &types.Integer{Val: intObj.Val-1}
}
func evalAllStatements(stms []ast.Statement) types.ObjectJIPL {
	var resrult types.ObjectJIPL

	for _, stm := range stms {
			resrult = Eval(stm)
	}

	return resrult
}
	

func evalPrefixExpression(operator string, operand types.ObjectJIPL) types.ObjectJIPL{
	switch operator {
	case "!":
		return evalComplementPrefix(operand)
	case "-":
		return evalMinusPrefix(operand)
	default:
		return nil
	}
}

func evalMinusPrefix(operand types.ObjectJIPL) types.ObjectJIPL {
	if operand.GetType() != types.T_ITNTEGER {
		return nil
	}
	intObj := operand.(*types.Integer)
	return &types.Integer{Val: -intObj.Val}
}

func evalComplementPrefix(operand types.ObjectJIPL) types.ObjectJIPL {
	if operand.GetType() != types.T_BOOLEAN {
		return nil
	}
	boolObj := operand.(*types.Boolean)
	if boolObj.Val {
		return types.FALSE
	}

	return types.TRUE
	
}

