package vm

import (
	"fmt"
	"github.com/IBAX-io/needle/compile"
	"github.com/shopspring/decimal"
)

type opFunc func(x, y interface{}) (interface{}, error)

var ops = map[string]opFunc{
	"+":  add,
	"+=": addAssign,
	"++": increment,
	"-":  subtract,
	"-=": subtractAssign,
	"--": decrement,
	"*":  multiply,
	"*=": multiplyAssign,
	"/":  divide,
	"/=": divideAssign,
	"%":  modulo,
	"%=": moduloAssign,
	"==": equal,
	//"=":   assign,
	"!=": notEqual,
	//"!":   not,
	"<<":  leftShift,
	"<<=": leftShiftAssign,
	"<=":  lessThanOrEqual,
	"<":   lessThan,
	">>":  rightShift,
	">>=": rightShiftAssign,
	">=":  greaterThanOrEqual,
	">":   greaterThan,
	"||":  logicalOr,
	"|=":  bitwiseOrAssign,
	"|":   bitwiseOr,
	"&&":  logicalAnd,
	"&=":  bitwiseAndAssign,
	"&":   bitwiseAnd,
	"^":   bitwiseXor,
	"^=":  bitwiseXorAssign,
}

func evaluateCmd(x, y interface{}, op compile.CmdT) (interface{}, error) {
	if f, ok := ops[op.String()]; ok {
		z, err := f(x, y)
		if err != nil {
			return nil, fmt.Errorf("error evaluating %s: %v", op, err)
		}
		return z, nil
	}
	return nil, fmt.Errorf("unsupported operator: %s", op)
}

// binary arithmetic operator
func binaryArithmeticOperator(x, y interface{}, op string) (interface{}, error) {
	switch x := x.(type) {
	case string:
		if y, ok := y.(string); ok {
			switch op {
			case "+":
				return x + y, nil
			}
		}
		if y, ok := y.(int64); ok {
			x, err := ValueToInt(x)
			if err != nil {
				return nil, err
			}
			switch op {
			case "+":
				return x + y, nil
			case "-":
				return x - y, nil
			case "*":
				return x * y, nil
			case "/":
				if y == 0 {
					return nil, fmt.Errorf("division by zero for %s", op)
				}
				return x / y, nil
			case "%":
				if y == 0 {
					return nil, fmt.Errorf("division by zero for %s", op)
				}
				return x % y, nil
			}
		}
		if y, ok := y.(float64); ok {
			x, err := decimal.NewFromString(x)
			if err != nil {
				return nil, err
			}
			y := decimal.NewFromFloat(y)

			switch op {
			case "+":
				return x.Add(y).InexactFloat64(), nil
			case "-":
				return x.Sub(y).InexactFloat64(), nil
			case "*":
				return x.Mul(y).InexactFloat64(), nil
			case "/":
				if y.IsZero() {
					return nil, fmt.Errorf("division by zero for %s", op)
				}
				return x.Div(y).InexactFloat64(), nil
			case "%":
				if y.IsZero() {
					return nil, fmt.Errorf("division by zero for %s", op)
				}
				return x.Mod(y).InexactFloat64(), nil
			}
		}
	case int64:
		if y, ok := y.(float64); ok {
			x := decimal.NewFromInt(x)
			y := decimal.NewFromFloat(y)
			switch op {
			case "+":
				return x.Add(y).InexactFloat64(), nil
			case "-":
				return x.Sub(y).InexactFloat64(), nil
			case "*":
				return x.Mul(y).InexactFloat64(), nil
			case "/":
				if y.IsZero() {
					return nil, fmt.Errorf("division by zero for %s", op)
				}
				return x.Div(y).InexactFloat64(), nil
			case "%":
				if y.IsZero() {
					return nil, fmt.Errorf("division by zero for %s", op)
				}
				return x.Mod(y).InexactFloat64(), nil
			}
		}
		switch y.(type) {
		case string, int64:
			yv, err := ValueToInt(y)
			if err != nil {
				return nil, err
			}
			switch op {
			case "+":
				return x + yv, nil
			case "-":
				return x - yv, nil
			case "*":
				return x * yv, nil
			case "/":
				if yv == 0 {
					return nil, fmt.Errorf("division by zero for %s", op)
				}
				return x / yv, nil
			case "%":
				if yv == 0 {
					return nil, fmt.Errorf("division by zero for %s", op)
				}
				return x % yv, nil
			}
		}
	case float64:
		switch y := y.(type) {
		case string, int64, float64:
			x := decimal.NewFromFloat(x)
			yv, err := ValueToDecimal(y)
			if err != nil {
				return nil, err
			}
			switch op {
			case "+":
				return x.Add(yv).InexactFloat64(), nil
			case "-":
				return x.Sub(yv).InexactFloat64(), nil
			case "*":
				return x.Mul(yv).InexactFloat64(), nil
			case "/":
				if yv.IsZero() {
					return nil, fmt.Errorf("division by zero for %s", op)
				}
				return x.Div(yv).InexactFloat64(), nil
			case "%":
				if yv.IsZero() {
					return nil, fmt.Errorf("division by zero for %s", op)
				}
				return x.Mod(yv).InexactFloat64(), nil

			}
		}
	case decimal.Decimal:
		y, err := ValueToDecimal(y)
		if err != nil {
			return nil, err
		}
		switch op {
		case "+":
			return x.Add(y), nil
		case "-":
			return x.Sub(y), nil
		case "*":
			return x.Mul(y), nil
		case "/":
			if y.IsZero() {
				return nil, fmt.Errorf("division by zero for %s", op)
			}
			return x.Div(y), nil
		case "%":
			if y.IsZero() {
				return nil, fmt.Errorf("division by zero for %s", op)
			}
			return x.Mod(y).InexactFloat64(), nil
		}

	}
	return nil, fmt.Errorf("invalid types for %s operation: %T and %T", op, x, y)
}

func add(x, y interface{}) (interface{}, error) {
	return binaryArithmeticOperator(x, y, "+")
}

func addAssign(x, y interface{}) (interface{}, error) {
	z, err := add(x, y)
	if err != nil {
		return nil, err
	}
	return z, nil
}

func increment(x, y interface{}) (interface{}, error) {
	switch x := x.(type) {
	case int64:
		return x + 1, nil
	case float64:
		return x + 1, nil
	case decimal.Decimal:
		return x.Add(decimal.NewFromInt(1)), nil
	}

	return nil, fmt.Errorf("invalid type for ++ operation: %T", x)
}

func subtract(x, y interface{}) (interface{}, error) {
	return binaryArithmeticOperator(x, y, "-")
}

func subtractAssign(x, y interface{}) (interface{}, error) {
	z, err := subtract(x, y)
	if err != nil {
		return nil, err
	}
	return z, nil
}

func decrement(x, y interface{}) (interface{}, error) {
	switch x := x.(type) {
	case int64:
		return x - 1, nil
	case float64:
		return x - 1, nil
	case decimal.Decimal:
		return x.Sub(decimal.NewFromInt(1)), nil
	}

	return nil, fmt.Errorf("invalid type for -- operation: %T", x)
}

func multiply(x, y interface{}) (interface{}, error) {
	return binaryArithmeticOperator(x, y, "*")
}

func multiplyAssign(x, y interface{}) (interface{}, error) {
	z, err := multiply(x, y)
	if err != nil {
		return nil, err
	}
	return z, nil
}

func divide(x, y interface{}) (interface{}, error) {
	return binaryArithmeticOperator(x, y, "/")
}

func divideAssign(x, y interface{}) (interface{}, error) {
	z, err := divide(x, y)
	if err != nil {
		return nil, err
	}
	return z, nil
}

func modulo(x, y interface{}) (interface{}, error) {
	return binaryArithmeticOperator(x, y, "%")
}

func moduloAssign(x, y interface{}) (interface{}, error) {
	z, err := modulo(x, y)
	if err != nil {
		return nil, err
	}
	return z, nil
}

// comparison operator
func comparisonOperator(x, y interface{}, op string) (interface{}, error) {
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
			switch op {
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
		}
	case int64:
		switch y.(type) {
		case string, int64, decimal.Decimal:
			y, err := ValueToInt(y)
			if err != nil {
				return nil, err
			}
			switch op {
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
		case float64:
			x := decimal.NewFromInt(x)
			y, err := ValueToDecimal(y)
			if err != nil {
				return nil, err
			}
			switch op {
			case "==":
				return x.Equal(y), nil
			case "!=":
				return !x.Equal(y), nil
			case ">":
				return x.GreaterThan(y), nil
			case ">=":
				return x.GreaterThanOrEqual(y), nil
			case "<":
				return x.LessThan(y), nil
			case "<=":
				return x.LessThanOrEqual(y), nil
			}
		}
	case float64:
		xv := decimal.NewFromFloat(x)
		y, err := ValueToDecimal(y)
		if err != nil {
			return nil, err
		}
		switch op {
		case "==":
			return xv.Equal(y), nil
		case "!=":
			return !xv.Equal(y), nil
		case ">":
			return xv.GreaterThan(y), nil
		case ">=":
			return xv.GreaterThanOrEqual(y), nil
		case "<":
			return xv.LessThan(y), nil
		case "<=":
			return xv.LessThanOrEqual(y), nil
		}
	case decimal.Decimal:
		y, err := ValueToDecimal(y)
		if err != nil {
			return nil, err
		}
		switch op {
		case "==":
			return x.Equal(y), nil
		case "!=":
			return !x.Equal(y), nil
		case ">":
			return x.GreaterThan(y), nil
		case ">=":
			return x.GreaterThanOrEqual(y), nil
		case "<":
			return x.LessThan(y), nil
		case "<=":
			return x.LessThanOrEqual(y), nil
		}
	}
	return nil, fmt.Errorf("invalid types for %s operation: %T and %T", op, x, y)
}

func equal(x, y interface{}) (interface{}, error) {
	return comparisonOperator(x, y, "==")
}

func assign(x, y interface{}) (interface{}, error) {
	return y, nil
}

func notEqual(x, y interface{}) (interface{}, error) {
	return comparisonOperator(x, y, "!=")
}

func not(x, y interface{}) (interface{}, error) {
	switch x := x.(type) {
	case bool:
		return !x, nil
	}

	return nil, fmt.Errorf("invalid type for ! operation: %T", x)
}

func leftShift(x, y interface{}) (interface{}, error) {
	switch x := x.(type) {
	case int64:
		if y, ok := y.(int64); ok {
			if y < 0 {
				return nil, fmt.Errorf("division by zero for <<")
			}
			return x << y, nil
		}
	}

	return nil, fmt.Errorf("invalid types for << operation: %T and %T", x, y)
}

func leftShiftAssign(x, y interface{}) (interface{}, error) {
	z, err := leftShift(x, y)
	if err != nil {
		return nil, err
	}
	return z, nil
}

func lessThanOrEqual(x, y interface{}) (interface{}, error) {
	return comparisonOperator(x, y, "<=")
}

func lessThan(x, y interface{}) (interface{}, error) {
	return comparisonOperator(x, y, "<")
}

func rightShift(x, y interface{}) (interface{}, error) {
	switch x := x.(type) {
	case int64:
		if y, ok := y.(int64); ok {
			if y < 0 {
				return nil, fmt.Errorf("division by zero for >>")
			}
			return x >> y, nil
		}
	}

	return nil, fmt.Errorf("invalid types for >> operation: %T and %T", x, y)
}

func rightShiftAssign(x, y interface{}) (interface{}, error) {
	z, err := rightShift(x, y)
	if err != nil {
		return nil, err
	}
	return z, nil
}

func greaterThanOrEqual(x, y interface{}) (interface{}, error) {
	return comparisonOperator(x, y, ">=")
}

func greaterThan(x, y interface{}) (interface{}, error) {
	return comparisonOperator(x, y, ">")
}

func logicalOr(x, y interface{}) (interface{}, error) {
	return valueToBool(x) || valueToBool(y), nil
}

func bitwiseOrAssign(x, y interface{}) (interface{}, error) {
	switch x := x.(type) {
	case int64:
		if y, ok := y.(int64); ok {
			x |= y
			return x, nil
		}
	}

	return nil, fmt.Errorf("invalid types for |= operation: %T and %T", x, y)
}

func bitwiseOr(x, y interface{}) (interface{}, error) {
	switch x := x.(type) {
	case int64:
		if y, ok := y.(int64); ok {
			return x | y, nil
		}
	}

	return nil, fmt.Errorf("invalid types for | operation: %T and %T", x, y)
}

func logicalAnd(x, y interface{}) (interface{}, error) {
	return valueToBool(x) && valueToBool(y), nil
}

func bitwiseAndAssign(x, y interface{}) (interface{}, error) {
	switch x := x.(type) {
	case int64:
		if y, ok := y.(int64); ok {
			x &= y
			return x, nil
		}
	}

	return nil, fmt.Errorf("invalid types for &= operation: %T and %T", x, y)
}

func bitwiseAnd(x, y interface{}) (interface{}, error) {
	switch x := x.(type) {
	case int64:
		if y, ok := y.(int64); ok {
			return x & y, nil
		}
	}

	return nil, fmt.Errorf("invalid types for & operation: %T and %T", x, y)
}

func bitwiseXor(x, y interface{}) (interface{}, error) {
	switch x := x.(type) {
	case int64:
		if y, ok := y.(int64); ok {
			return x ^ y, nil
		}
	}

	return nil, fmt.Errorf("invalid types for ^ operation: %T and %T", x, y)
}

func bitwiseXorAssign(x, y interface{}) (interface{}, error) {
	switch x := x.(type) {
	case int64:
		if y, ok := y.(int64); ok {
			x ^= y
			return x, nil
		}
	}

	return nil, fmt.Errorf("invalid types for ^= operation: %T and %T", x, y)
}
