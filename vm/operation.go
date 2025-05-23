package vm

import (
	"fmt"

	"github.com/IBAX-io/needle/compiler"

	"github.com/shopspring/decimal"
)

type (
	opFunc    func(x, y any) (any, error)
	operation struct {
		op string
		f  opFunc
	}
)

var cmd2operation = map[compiler.Cmd]operation{
	compiler.CmdAdd:          {"+", add},
	compiler.CmdAssignAdd:    {"+=", addAssign},
	compiler.CmdInc:          {"++", increment},
	compiler.CmdSub:          {"-", subtract},
	compiler.CmdAssignSub:    {"-=", subtractAssign},
	compiler.CmdDec:          {"--", decrement},
	compiler.CmdMul:          {"*", multiply},
	compiler.CmdAssignMul:    {"*=", multiplyAssign},
	compiler.CmdDiv:          {"/", divide},
	compiler.CmdAssignDiv:    {"/=", divideAssign},
	compiler.CmdMod:          {"%", modulo},
	compiler.CmdAssignMod:    {"%=", moduloAssign},
	compiler.CmdEqual:        {"==", equal},
	compiler.CmdNotEq:        {"!=", notEqual},
	compiler.CmdShiftL:       {"<<", leftShift},
	compiler.CmdAssignLShift: {"<<=", leftShiftAssign},
	compiler.CmdLessEq:       {"<=", lessThanOrEqual},
	compiler.CmdLess:         {"<", lessThan},
	compiler.CmdShiftR:       {">>", rightShift},
	compiler.CmdAssignRShift: {">>=", rightShiftAssign},
	compiler.CmdGrEq:         {">=", greaterThanOrEqual},
	compiler.CmdGreat:        {">", greaterThan},
	compiler.CmdOr:           {"||", logicalOr},
	compiler.CmdAssignOr:     {"|=", bitwiseOrAssign},
	compiler.CmdBitOr:        {"|", bitwiseOr},
	compiler.CmdAnd:          {"&&", logicalAnd},
	compiler.CmdAssignAnd:    {"&=", bitwiseAndAssign},
	compiler.CmdBitAnd:       {"&", bitwiseAnd},
	compiler.CmdBitXor:       {"^", bitwiseXor},
	compiler.CmdAssignXor:    {"^=", bitwiseXorAssign},
}

func operationExpr(x, y any, cmd compiler.Cmd) (any, error) {
	opera, ok := cmd2operation[cmd]
	if ok {
		z, err := opera.f(x, y)
		if err != nil {
			return nil, fmt.Errorf("operation error: %v", err)
		}
		return z, nil
	}
	return nil, fmt.Errorf("unsupported operator: %s", opera.op)
}

func reportOpError(op string, x, y any) error {
	return fmt.Errorf("invalid types for %s operation: %T and %T", op, x, y)
}

func reportConvertError(op string, err error) error {
	return fmt.Errorf("convert %s error: %w", op, err)
}

func add(x, y any) (any, error) {
	return binaryArithmeticOperator(x, y, "+")
}

func addAssign(x, y any) (any, error) {
	z, err := add(x, y)
	if err != nil {
		return nil, err
	}
	return z, nil
}

func increment(x, y any) (any, error) {
	switch x := x.(type) {
	case int64:
		return x + 1, nil
	case float64:
		return x + 1, nil
	case decimal.Decimal:
		return x.Add(decimal.NewFromInt(1)), nil
	}

	return nil, fmt.Errorf("invalid types for ++ operation: %T", x)
}

func subtract(x, y any) (any, error) {
	return binaryArithmeticOperator(x, y, "-")
}

func subtractAssign(x, y any) (any, error) {
	z, err := subtract(x, y)
	if err != nil {
		return nil, err
	}
	return z, nil
}

func decrement(x, y any) (any, error) {
	switch x := x.(type) {
	case int64:
		return x - 1, nil
	case float64:
		return x - 1, nil
	case decimal.Decimal:
		return x.Sub(decimal.NewFromInt(1)), nil
	}

	return nil, fmt.Errorf("invalid types for -- operation: %T", x)
}

func multiply(x, y any) (any, error) {
	return binaryArithmeticOperator(x, y, "*")
}

func multiplyAssign(x, y any) (any, error) {
	z, err := multiply(x, y)
	if err != nil {
		return nil, err
	}
	return z, nil
}

func divide(x, y any) (any, error) {
	return binaryArithmeticOperator(x, y, "/")
}

func divideAssign(x, y any) (any, error) {
	z, err := divide(x, y)
	if err != nil {
		return nil, err
	}
	return z, nil
}

func modulo(x, y any) (any, error) {
	return binaryArithmeticOperator(x, y, "%")
}

func moduloAssign(x, y any) (any, error) {
	z, err := modulo(x, y)
	if err != nil {
		return nil, err
	}
	return z, nil
}

func equal(x, y any) (any, error) {
	return comparisonOperator(x, y, "==")
}

func assign(_, y any) (any, error) {
	return y, nil
}

func notEqual(x, y any) (any, error) {
	return comparisonOperator(x, y, "!=")
}

func not(x, _ any) (any, error) {
	switch x := x.(type) {
	case bool:
		return !x, nil
	}

	return nil, fmt.Errorf("invalid types for ! operation: %T", x)
}

func leftShift(x, y any) (any, error) {
	switch x := x.(type) {
	case int64:
		if y, ok := y.(int64); ok {
			if y < 0 {
				return nil, fmt.Errorf("division by zero for <<")
			}
			return x << y, nil
		}
	}

	return nil, reportOpError("<<", x, y)
}

func leftShiftAssign(x, y any) (any, error) {
	z, err := leftShift(x, y)
	if err != nil {
		return nil, err
	}
	return z, nil
}

func lessThanOrEqual(x, y any) (any, error) {
	return comparisonOperator(x, y, "<=")
}

func lessThan(x, y any) (any, error) {
	return comparisonOperator(x, y, "<")
}

func rightShift(x, y any) (any, error) {
	switch x := x.(type) {
	case int64:
		if y, ok := y.(int64); ok {
			if y < 0 {
				return nil, fmt.Errorf("division by zero for >>")
			}
			return x >> y, nil
		}
	}

	return nil, reportOpError(">>", x, y)
}

func rightShiftAssign(x, y any) (any, error) {
	z, err := rightShift(x, y)
	if err != nil {
		return nil, err
	}
	return z, nil
}

func greaterThanOrEqual(x, y any) (any, error) {
	return comparisonOperator(x, y, ">=")
}

func greaterThan(x, y any) (any, error) {
	return comparisonOperator(x, y, ">")
}

func logicalOr(x, y any) (any, error) {
	return valueToBool(x) || valueToBool(y), nil
}

func bitwiseOrAssign(x, y any) (any, error) {
	switch x := x.(type) {
	case int64:
		if y, ok := y.(int64); ok {
			x |= y
			return x, nil
		}
	}

	return nil, reportOpError("|=", x, y)
}

func bitwiseOr(x, y any) (any, error) {
	switch x := x.(type) {
	case int64:
		if y, ok := y.(int64); ok {
			return x | y, nil
		}
	}

	return nil, reportOpError("|", x, y)
}

func logicalAnd(x, y any) (any, error) {
	return valueToBool(x) && valueToBool(y), nil
}

func bitwiseAndAssign(x, y any) (any, error) {
	switch x := x.(type) {
	case int64:
		if y, ok := y.(int64); ok {
			x &= y
			return x, nil
		}
	}

	return nil, reportOpError("&=", x, y)
}

func bitwiseAnd(x, y any) (any, error) {
	switch x := x.(type) {
	case int64:
		if y, ok := y.(int64); ok {
			return x & y, nil
		}
	}

	return nil, reportOpError("&", x, y)
}

func bitwiseXor(x, y any) (any, error) {
	switch x := x.(type) {
	case int64:
		if y, ok := y.(int64); ok {
			return x ^ y, nil
		}
	}

	return nil, reportOpError("^", x, y)
}

func bitwiseXorAssign(x, y any) (any, error) {
	switch x := x.(type) {
	case int64:
		if y, ok := y.(int64); ok {
			x ^= y
			return x, nil
		}
	}

	return nil, reportOpError("^", x, y)
}

func stringstring(x, y string, op string) (any, error) {
	switch op {
	case "+":
		return x + y, nil
	case "==":
		return x == y, nil
	case "!=":
		return x != y, nil
	case ">":
		return x > y, nil
	case ">=":
		return x >= y, nil
	case "<":
		return x < y, nil
	case "<=":
		return x <= y, nil
	}
	return nil, reportOpError(op, x, y)
}

func intint(x, y int64, op string) (any, error) {
	switch op {
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "*":
		return x * y, nil
	case "/":
		if y == 0 {
			return 0, fmt.Errorf("division by zero for %s", op)
		}
		return x / y, nil
	case "%":
		if y == 0 {
			return nil, fmt.Errorf("division by zero for %s", op)
		}
		return x % y, nil
	case "==":
		return x == y, nil
	case "!=":
		return x != y, nil
	case ">":
		return x > y, nil
	case ">=":
		return x >= y, nil
	case "<":
		return x < y, nil
	case "<=":
		return x <= y, nil
	}
	return nil, reportOpError(op, x, y)
}

// binary arithmetic operator
func binaryArithmeticOperator(x, y any, op string) (any, error) {
	switch x := x.(type) {
	case string:
		if y, ok := y.(string); ok {
			return stringstring(x, y, op)
		}
		if y, ok := y.(int64); ok {
			xv, err := ValueToInt(x)
			if err != nil {
				return nil, reportConvertError(op, err)
			}
			return intint(xv, y, op)
		}
		if y, ok := y.(float64); ok {
			return floatfloat(x, y, op)
		}
	case int64:
		if y, ok := y.(float64); ok {
			return floatfloat(x, y, op)
		}
		switch y.(type) {
		case string, int64:
			yv, err := ValueToInt(y)
			if err != nil {
				return nil, reportConvertError(op, err)
			}
			return intint(x, yv, op)
		}
	case float64:
		switch y := y.(type) {
		case string, int64, float64:
			return floatfloat(x, y, op)
		}
	case decimal.Decimal:
		return decimaldecimal(x, y, op)
	}
	return nil, reportOpError(op, x, y)
}

// comparison operator
func comparisonOperator(x, y any, op string) (any, error) {
	switch x := x.(type) {
	case nil:
		switch op {
		case "==":
			return x == y, nil
		case "!=":
			return x != nil, nil
		}
	case bool:
		if y, ok := y.(bool); ok {
			switch op {
			case "==":
				return x == y, nil
			case "!=":
				return x != y, nil
			}
		}
	case string:
		switch y.(type) {
		case string, int64, float64, decimal.Decimal:
			y := fmt.Sprintf("%v", y)
			return stringstring(x, y, op)
		}
	case int64:
		switch y.(type) {
		case string, int64, decimal.Decimal:
			yv, err := ValueToInt(y)
			if err != nil {
				return nil, reportConvertError(op, err)
			}
			return intint(x, yv, op)
		case float64:
			return floatfloat(x, y, op)
		}
	case float64:
		return floatfloat(x, y, op)
	case decimal.Decimal:
		return decimaldecimal(x, y, op)
	}
	return nil, reportOpError(op, x, y)
}

func floatfloat(x, y any, op string) (any, error) {
	xv, err := ValueToDecimal(x)
	if err != nil {
		return nil, reportConvertError(op, err)
	}
	var yv decimal.Decimal
	yv, err = ValueToDecimal(y)
	if err != nil {
		return nil, reportConvertError(op, err)
	}
	switch op {
	case "+":
		return xv.Add(yv).InexactFloat64(), nil
	case "-":
		return xv.Sub(yv).InexactFloat64(), nil
	case "*":
		return xv.Mul(yv).InexactFloat64(), nil
	case "/":
		if yv.IsZero() {
			return nil, fmt.Errorf("division by zero for %s", op)
		}
		return xv.Div(yv).InexactFloat64(), nil
	case "%":
		if yv.IsZero() {
			return nil, fmt.Errorf("division by zero for %s", op)
		}
		return xv.Mod(yv).InexactFloat64(), nil
	case "==":
		return xv.Equal(yv), nil
	case "!=":
		return !xv.Equal(yv), nil
	case ">":
		return xv.GreaterThan(yv), nil
	case ">=":
		return xv.GreaterThanOrEqual(yv), nil
	case "<":
		return xv.LessThan(yv), nil
	case "<=":
		return xv.LessThanOrEqual(yv), nil
	}
	return nil, reportOpError(op, x, y)
}

func decimaldecimal(x decimal.Decimal, y any, op string) (any, error) {
	yv, err := ValueToDecimal(y)
	if err != nil {
		return nil, reportConvertError(op, err)
	}
	switch op {
	case "+":
		return x.Add(yv), nil
	case "-":
		return x.Sub(yv), nil
	case "*":
		return x.Mul(yv), nil
	case "/":
		if yv.IsZero() {
			return nil, fmt.Errorf("division by zero for %s", op)
		}
		return x.Div(yv), nil
	case "%":
		if yv.IsZero() {
			return nil, fmt.Errorf("division by zero for %s", op)
		}
		return x.Mod(yv), nil
	case "==":
		return x.Equal(yv), nil
	case "!=":
		return !x.Equal(yv), nil
	case ">":
		return x.GreaterThan(yv), nil
	case ">=":
		return x.GreaterThanOrEqual(yv), nil
	case "<":
		return x.LessThan(yv), nil
	case "<=":
		return x.LessThanOrEqual(yv), nil
	}
	return nil, reportOpError(op, x, yv)
}
