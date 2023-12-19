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
	case *ast.InfixExpression:
		leftOperand := Eval(node.Left)
		rightOperand := Eval(node.Right)
		return evalInfixExpression(node.Operator,leftOperand,rightOperand)
	default:
		return nil
	}
}

func evalInfixExpression(operator string, leftOperand, rightOperand types.ObjectJIPL) types.ObjectJIPL {
	if leftOperand.GetType() != types.T_INTEGER || 
	rightOperand.GetType() !=types.T_INTEGER {
		return nil
	}
	return evalIntInfixExpression(operator,leftOperand,rightOperand)
}
	

func evalIntInfixExpression(operator string, left, right types.ObjectJIPL)  types.ObjectJIPL {
	intObjRight := right.(*types.Integer)
	intObjLeft := left.(*types.Integer)
	switch operator {
	case "+":
		return &types.Integer{Val: intObjLeft.Val + intObjRight.Val}
	case "-":
		return &types.Integer{Val: intObjLeft.Val - intObjRight.Val}
	case "*":
		return &types.Integer{Val: intObjLeft.Val * intObjRight.Val}
	case "/":
		return &types.Integer{Val: intObjLeft.Val / intObjRight.Val}
	case "%":
		return &types.Integer{Val: intObjLeft.Val % intObjRight.Val}
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
	if operand.GetType() != types.T_INTEGER {
		return nil
	}
	intObj := operand.(*types.Integer)
	return &types.Integer{Val: intObj.Val+1}

}

func evalDecrementPostfix(operand types.ObjectJIPL) types.ObjectJIPL{
	if operand.GetType() != types.T_INTEGER{
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
	if operand.GetType() != types.T_INTEGER {
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

