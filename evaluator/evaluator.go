package evaluator

import (
	"math"
	"wearwork/ast"
	"wearwork/object"
)

var (
	NULL = &object.Null{}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.CallExpression:
		arg := Eval(node.Argument)
		switch function := node.Function.String(); function {
		case "sin":
			return &object.Double{Value: math.Sin(arg.(*object.Double).Value)}
		case "cos":
			return &object.Double{Value: math.Cos(arg.(*object.Double).Value)}
		}
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator, left, right)
	case *ast.DoubleLiteral:
		return &object.Double{Value: node.Value}
	}
	return nil
}
func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range stmts {
		result = Eval(statement)
	}
	return result
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return NULL
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.DOUBLE_OBJ {
		return NULL
	}
	value := right.(*object.Double).Value
	return &object.Double{Value: -value}
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.DOUBLE_OBJ && right.Type() == object.DOUBLE_OBJ:
		return evalDoubleInfixExpression(operator, left, right)
	default:
		return NULL
	}
}

func evalDoubleInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Double).Value
	rightVal := right.(*object.Double).Value
	switch operator {
	case "+":
		return &object.Double{Value: leftVal + rightVal}
	case "-":
		return &object.Double{Value: leftVal - rightVal}
	case "*":
		return &object.Double{Value: leftVal * rightVal}
	case "/":
		return &object.Double{Value: leftVal / rightVal}
	default:
		return NULL
	}
}
