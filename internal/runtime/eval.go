package runtime

import (
	"fmt"

	ast "github.com/houcine7/JIPL/internal/AST"
	"github.com/houcine7/JIPL/internal/debug"
	"github.com/houcine7/JIPL/internal/types"
)

func Eval(node ast.Node, ctx *types.Context) (types.ObjectJIPL, *debug.Error) {
	switch node := node.(type) {
	case *ast.Program:
		return evalAllProgramStatements(node.Statements, ctx)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, ctx)
	case *ast.ReturnStatement:
		if ctx.Outer == nil {
			return nil, debug.NewError("return statements can only be used insed a function") // TODO:  to be tested
		}
		value, err := Eval(node.ReturnValue, ctx)
		if err != debug.NOERROR {
			return nil, err
		}
		return &types.Return{Val: value}, err
	case *ast.DefStatement:
		val, err := Eval(node.Value, ctx)
		if err != debug.NOERROR {
			return nil, err
		}
		ctx.Set(node.Name.Value, val)

		return val, err
	case *ast.Identifier:
		return evalIdentifier(node, ctx)
	case *ast.ForLoopExpression:
		return evalForLoopExpression(node, ctx)
	case *ast.IfExpression:
		return evalIfExpression(node, ctx)
	case *ast.FunctionExp:
		ctx.Set(node.Name.Value, &types.Function{Name: node.Name.Value, Params: node.Parameters,
			Body: node.FnBody, Ctx: ctx})
		return &types.Function{Name: node.Name.Value, Params: node.Parameters,
			Body: node.FnBody, Ctx: ctx}, debug.NOERROR
	case *ast.FunctionCall:
		function, err := Eval(node.Function, ctx)
		if err != debug.NOERROR {
			return nil, err
		}
		args, err := evalExpressions(node.Arguments, ctx)
		if err != debug.NOERROR {
			return nil, err
		}
		return applyFunction(function, args)
	case *ast.BlockStm:
		return evalABlockStatements(node.Statements, ctx)
	case *ast.IntegerLiteral:
		return &types.Integer{Val: node.Value}, debug.NOERROR
	case *ast.StringLiteral:
		return &types.String{Val: node.Value}, debug.NOERROR
	case *ast.BooleanExp:
		return types.BoolToObJIPL(node.Value), debug.NOERROR
	case *ast.PrefixExpression:
		operand, _ := Eval(node.Right, ctx)
		return evalPrefixExpression(node.Operator, operand)
	case *ast.PostfixExpression:
		operand, _ := Eval(node.Left, ctx)
		return evalPostfixExpression(node.Operator, operand)
	case *ast.InfixExpression:
		leftOperand, _ := Eval(node.Left, ctx)
		rightOperand, _ := Eval(node.Right, ctx)
		return evalInfixExpression(node.Operator, leftOperand, rightOperand)
	default:
		return nil, debug.NewError("unknown ast node type")
	}
}

func applyFunction(function types.ObjectJIPL, args []types.ObjectJIPL) (types.ObjectJIPL, *debug.Error) {
	switch fn := function.(type) {
	case *types.Function:

		appendedCtx := appedCtx(fn, args)

		eval, err := Eval(fn.Body, appendedCtx)
		if err != debug.NOERROR {
			return nil, err
		}

		return uwrapReturnValue(eval), debug.NOERROR
		// No return statement for the fn body
	case *types.BuiltIn:
		return fn.Fn(args...), debug.NOERROR
	default:
		return nil, debug.NewError("function not defined")
	}
}

func appedCtx(fn *types.Function, args []types.ObjectJIPL) *types.Context {
	ctx := types.NewContextWithOuter(fn.Ctx)

	for i, param := range fn.Params {
		ctx.Set(param.Value, args[i])
	}

	return ctx
}

func uwrapReturnValue(obj types.ObjectJIPL) types.ObjectJIPL {
	if returnVal, ok := obj.(*types.Return); ok {
		return returnVal.Val
	}
	return types.UNDEFIEND
}

func evalExpressions(node []ast.Expression, ctx *types.Context) ([]types.ObjectJIPL, *debug.Error) {
	var result []types.ObjectJIPL
	for _, exp := range node {
		evaluated, err := Eval(exp, ctx)

		if err != debug.NOERROR {
			return nil, err
		}
		result = append(result, evaluated)
	}
	return result, debug.NOERROR
}

func evalIdentifier(node *ast.Identifier, ctx *types.Context) (types.ObjectJIPL, *debug.Error) {
	val, ok := ctx.Get(node.Value)

	if ok {
		return val, debug.NOERROR
	}
	builtin, ok := builtins[node.Value]
	if ok {
		return builtin, debug.NOERROR
	}
	return nil, debug.NewError(fmt.Sprintf("identifier not found: %s", node.Value))
}

func evalIfExpression(ifExp *ast.IfExpression, ctx *types.Context) (types.ObjectJIPL, *debug.Error) {
	condition, _ := Eval(ifExp.Condition, ctx)
	if condition == types.TRUE {
		return Eval(ifExp.Body, ctx)
	}
	if ifExp.ElseBody != nil {
		return Eval(ifExp.ElseBody, ctx)
	}
	return nil, debug.NOERROR
}

func evalInfixExpression(operator string, leftOperand, rightOperand types.ObjectJIPL) (types.ObjectJIPL, *debug.Error) {

	if leftOperand.GetType() == types.T_INTEGER &&
		rightOperand.GetType() == types.T_INTEGER {
		return evalIntInfixExpression(operator, leftOperand, rightOperand)
	}

	if leftOperand.GetType() == types.T_BOOLEAN &&
		rightOperand.GetType() == types.T_BOOLEAN {
		return evalBoolInfixExpression(operator, leftOperand, rightOperand)
	}

	if leftOperand.GetType() == types.T_STRING &&
		rightOperand.GetType() == types.T_STRING {
		return evlStringInfix(operator, leftOperand, rightOperand)
	}

	return nil, debug.NewError(fmt.Sprintf("type mismatch: %s %s %s", leftOperand.GetType(), operator, rightOperand.GetType()))
}

func evlStringInfix(operator string, left, right types.ObjectJIPL) (types.ObjectJIPL, *debug.Error) {
	stringObjRight := right.(*types.String)
	stringObjLeft := left.(*types.String)
	switch operator {
	case "+":
		return &types.String{Val: stringObjLeft.Val + stringObjRight.Val}, debug.NOERROR
	case "==":
		return types.BoolToObJIPL(stringObjLeft.Val == stringObjRight.Val), debug.NOERROR
	case "!=":
		return types.BoolToObJIPL(stringObjLeft.Val != stringObjRight.Val), debug.NOERROR
	default:
		return nil, debug.NewError("unknown operator")
	}
}
func evalBoolInfixExpression(operator string, left, right types.ObjectJIPL) (types.ObjectJIPL, *debug.Error) {
	boolObjRight := right.(*types.Boolean)
	boolObjLeft := left.(*types.Boolean)
	switch operator {
	case "==":
		return types.BoolToObJIPL(boolObjLeft.Val == boolObjRight.Val), debug.NOERROR
	case "!=":
		return types.BoolToObJIPL(boolObjLeft.Val != boolObjRight.Val), debug.NOERROR
	case "&&":
		return types.BoolToObJIPL(boolObjLeft.Val && boolObjRight.Val), debug.NOERROR
	case "||":
		return types.BoolToObJIPL(boolObjLeft.Val || boolObjRight.Val), debug.NOERROR
	default:
		return nil, debug.NewError("unknown operator")
	}
}

func evalIntInfixExpression(operator string, left, right types.ObjectJIPL) (types.ObjectJIPL, *debug.Error) {
	intObjRight := right.(*types.Integer)
	intObjLeft := left.(*types.Integer)
	switch operator {
	case "+":
		return &types.Integer{Val: intObjLeft.Val + intObjRight.Val}, debug.NOERROR
	case "-":
		return &types.Integer{Val: intObjLeft.Val - intObjRight.Val}, debug.NOERROR
	case "*":
		return &types.Integer{Val: intObjLeft.Val * intObjRight.Val}, debug.NOERROR
	case "/":
		return &types.Integer{Val: intObjLeft.Val / intObjRight.Val}, debug.NOERROR
	case "%":
		return &types.Integer{Val: intObjLeft.Val % intObjRight.Val}, debug.NOERROR
	case "==":
		return types.BoolToObJIPL(intObjLeft.Val == intObjRight.Val), debug.NOERROR
	case "!=":
		return types.BoolToObJIPL(intObjLeft.Val != intObjRight.Val), debug.NOERROR
	case "<":
		return types.BoolToObJIPL(intObjLeft.Val < intObjRight.Val), debug.NOERROR
	case "<=":
		return types.BoolToObJIPL(intObjLeft.Val <= intObjRight.Val), debug.NOERROR
	case ">":
		return types.BoolToObJIPL(intObjLeft.Val > intObjRight.Val), debug.NOERROR
	case ">=":
		return types.BoolToObJIPL(intObjLeft.Val >= intObjRight.Val), debug.NOERROR
	default:
		return nil, debug.NewError("unknown operator")
	}
}

func evalForLoopExpression(forLoop *ast.ForLoopExpression, ctx *types.Context) (types.ObjectJIPL, *debug.Error) {

	_, err := Eval(forLoop.InitStm, ctx)

	if err != debug.NOERROR {
		fmt.Println("error while evaluating the init statement", err)
		return nil, err
	}

	// the condition
	condition, _ := Eval(forLoop.Condition, ctx)

	for condition == types.TRUE {
		iterationEval, _ := Eval(forLoop.Body, ctx)
		returnEval, ok := iterationEval.(*types.Return)
		if ok {
			return returnEval.Val, debug.NOERROR
		}

		postEval, _ := Eval(forLoop.PostIteration, ctx)
		postFix, ok := forLoop.PostIteration.(*ast.PostfixExpression)

		if ok {
			ctx.Set(postFix.Left.(*ast.Identifier).Value,
				postEval)
			condition, _ = Eval(forLoop.Condition, ctx)
		}
	}
	return nil, debug.NOERROR
}

func evalPostfixExpression(operator string, operand types.ObjectJIPL) (types.ObjectJIPL, *debug.Error) {
	switch operator {
	case "--":
		return evalDecrementPostfix(operand)
	case "++":
		return evalIncrementPostfix(operand)
	default:
		return types.UNDEFIEND, debug.NewError("unknown operator")
	}
}

func evalIncrementPostfix(operand types.ObjectJIPL) (types.ObjectJIPL, *debug.Error) {
	if operand.GetType() != types.T_INTEGER {
		return nil, debug.NewError("operand is not an integer")
	}
	intObj := operand.(*types.Integer)
	return &types.Integer{Val: intObj.Val + 1}, debug.NOERROR
}

func evalDecrementPostfix(operand types.ObjectJIPL) (types.ObjectJIPL, *debug.Error) {
	if operand.GetType() != types.T_INTEGER {
		return nil, debug.NewError("operand is not an integer")
	}
	intObj := operand.(*types.Integer)
	return &types.Integer{Val: intObj.Val - 1}, debug.NOERROR
}

func evalAllProgramStatements(stms []ast.Statement, ctx *types.Context) (types.ObjectJIPL, *debug.Error) {
	var result types.ObjectJIPL
	var err *debug.Error = debug.NOERROR

	for _, stm := range stms {
		result, err = Eval(stm, ctx)
		if err != debug.NOERROR {
			return nil, err
		}

		if result != nil && result.GetType() == types.T_RETURN {
			return result.(*types.Return).Val, debug.NOERROR
		}
	}
	return result, debug.NOERROR
}

func evalABlockStatements(stms []ast.Statement, ctx *types.Context) (types.ObjectJIPL, *debug.Error) {
	var result types.ObjectJIPL
	var err = debug.NOERROR

	for _, stm := range stms {
		result, err = Eval(stm, ctx)
		if err != debug.NOERROR {
			return nil, err
		}
		if result != nil && result.GetType() == types.T_RETURN {
			return result, debug.NOERROR
		}
	}
	return result, debug.NOERROR
}

func evalPrefixExpression(operator string, operand types.ObjectJIPL) (types.ObjectJIPL, *debug.Error) {
	switch operator {
	case "!":
		return evalComplementPrefix(operand)
	case "-":
		return evalMinusPrefix(operand)
	default:
		return nil, debug.NewError("unknown operator")
	}
}

func evalMinusPrefix(operand types.ObjectJIPL) (types.ObjectJIPL, *debug.Error) {
	if operand.GetType() != types.T_INTEGER {
		return nil, debug.NewError("operand is not an integer")
	}
	intObj := operand.(*types.Integer)
	return &types.Integer{Val: -intObj.Val}, debug.NOERROR
}

func evalComplementPrefix(operand types.ObjectJIPL) (types.ObjectJIPL, *debug.Error) {
	if operand.GetType() != types.T_BOOLEAN {
		return nil, debug.NewError("operand is not a boolean")
	}
	boolObj := operand.(*types.Boolean)
	if boolObj.Val {
		return types.FALSE, debug.NOERROR
	}

	return types.TRUE, debug.NOERROR
}
