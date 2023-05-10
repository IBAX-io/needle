package vm

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type opFunc func(x, y interface{}) (interface{}, error)

var ops = map[string]opFunc{
	"+":   add,
	"+=":  addAssign,
	"++":  increment,
	"-":   subtract,
	"-=":  subtractAssign,
	"--":  decrement,
	"*":   multiply,
	"*=":  multiplyAssign,
	"/":   divide,
	"/=":  divideAssign,
	"%":   modulo,
	"%=":  moduloAssign,
	"==":  equal,
	"=":   assign,
	"!=":  notEqual,
	"!":   not,
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

func evaluateCmd(x, y interface{}, op string) (interface{}, error) {
	if f, ok := ops[op]; ok {
		return f(x, y)
	} else {
		return nil, fmt.Errorf("unsupported operator: %s", op)
	}
}

func add(x, y interface{}) (interface{}, error) {
	switch x := x.(type) {
	case int64:
		if y, ok := y.(int64); ok {
			return x + y, nil
		}
	case float64:
		if y, ok := y.(float64); ok {
			return decimal.NewFromFloat(x).Add(decimal.NewFromFloat(y)).InexactFloat64(), nil
		}
	case string:
		if y, ok := y.(string); ok {
			return x + y, nil
		}
	case decimal.Decimal:
		if y, ok := y.(decimal.Decimal); ok {
			return x.Add(y), nil
		}
	}

	return nil, fmt.Errorf("invalid types for + operation: %T and %T", x, y)
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
	switch x := x.(type) {
	case int64:
		if y, ok := y.(int64); ok {
			return x - y, nil
		}
	case float64:
		if y, ok := y.(float64); ok {
			return decimal.NewFromFloat(x).Sub(decimal.NewFromFloat(y)).InexactFloat64(), nil
		}
	case decimal.Decimal:
		if y, ok := y.(decimal.Decimal); ok {
			return x.Sub(y), nil
		}
	}

	return nil, fmt.Errorf("invalid types for - operation: %T and %T", x, y)
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
	switch x := x.(type) {
	case int64:
		if y, ok := y.(int64); ok {
			return x * y, nil
		}
	case float64:
		if y, ok := y.(float64); ok {
			return decimal.NewFromFloat(x).Mul(decimal.NewFromFloat(y)).InexactFloat64(), nil
		}
	case decimal.Decimal:
		if y, ok := y.(decimal.Decimal); ok {
			return x.Mul(y), nil
		}
	}

	return nil, fmt.Errorf("invalid types for * operation: %T and %T", x, y)
}

func multiplyAssign(x, y interface{}) (interface{}, error) {
	z, err := multiply(x, y)
	if err != nil {
		return nil, err
	}
	return z, nil
}

func divide(x, y interface{}) (interface{}, error) {
	switch x := x.(type) {
	case int64:
		if y, ok := y.(int64); ok {
			if y == 0 {
				return nil, fmt.Errorf("division by zero for /")
			}
			return x / y, nil
		}
	case float64:
		if y, ok := y.(float64); ok {
			if y == 0 {
				return nil, fmt.Errorf("division by zero for /")
			}
			return decimal.NewFromFloat(x).Div(decimal.NewFromFloat(y)).InexactFloat64(), nil
		}
	case decimal.Decimal:
		if y, ok := y.(decimal.Decimal); ok {
			if y.Equal(decimal.Zero) {
				return nil, fmt.Errorf("division by zero for /")
			}
			return x.Div(y), nil
		}
	}

	return nil, fmt.Errorf("invalid types for / operation: %T and %T", x, y)
}

func divideAssign(x, y interface{}) (interface{}, error) {
	z, err := divide(x, y)
	if err != nil {
		return nil, err
	}
	return z, nil
}

func modulo(x, y interface{}) (interface{}, error) {
	switch x := x.(type) {
	case int64:
		if y, ok := y.(int64); ok {
			if y == 0 {
				return nil, fmt.Errorf("division by zero for %s", "%")
			}
			return x % y, nil
		}
	case float64:
		if y, ok := y.(float64); ok {
			if y == 0 {
				return nil, fmt.Errorf("division by zero for %s", "%")
			}
			return decimal.NewFromFloat(x).Mod(decimal.NewFromFloat(y)).InexactFloat64(), nil
		}
	case decimal.Decimal:
		if y, ok := y.(decimal.Decimal); ok {
			if y.Equal(decimal.Zero) {
				return nil, fmt.Errorf("division by zero for %s", "%")
			}
			return x.Mod(y), nil
		}
	}

	return nil, fmt.Errorf("invalid types for %s operation: %T and %T", "%", x, y)
}

func moduloAssign(x, y interface{}) (interface{}, error) {
	z, err := modulo(x, y)
	if err != nil {
		return nil, err
	}
	return z, nil
}

func equal(x, y interface{}) (interface{}, error) {
	return x == y, nil
}

func assign(x, y interface{}) (interface{}, error) {
	return y, nil
}

func notEqual(x, y interface{}) (interface{}, error) {
	return x != y, nil
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
	switch x := x.(type) {
	case int64:
		if y, ok := y.(int64); ok {
			return x <= y, nil
		}
	case float64:
		if y, ok := y.(float64); ok {
			return x <= y, nil
		}
	case decimal.Decimal:
		if y, ok := y.(decimal.Decimal); ok {
			return x.Cmp(y) <= 0, nil
		}
	case string:
		if y, ok := y.(string); ok {
			return x <= y, nil
		}
	}

	return nil, fmt.Errorf("invalid types for <= operation: %T and %T", x, y)
}

func lessThan(x, y interface{}) (interface{}, error) {
	switch x := x.(type) {
	case int64:
		if y, ok := y.(int64); ok {
			return x < y, nil
		}
	case float64:
		if y, ok := y.(float64); ok {
			return x < y, nil
		}
	case decimal.Decimal:
		if y, ok := y.(decimal.Decimal); ok {
			return x.Cmp(y) < 0, nil
		}
	case string:
		if y, ok := y.(string); ok {
			return x < y, nil
		}
	}

	return nil, fmt.Errorf("invalid types for < operation: %T and %T", x, y)
}

func rightShift(x, y interface{}) (interface{}, error) {
	switch x := x.(type) {
	case int64:
		if y, ok := y.(int64); ok {
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
	switch x := x.(type) {
	case int64:
		if y, ok := y.(int64); ok {
			return x >= y, nil
		}
	case float64:
		if y, ok := y.(float64); ok {
			return x >= y, nil
		}
	case decimal.Decimal:
		if y, ok := y.(decimal.Decimal); ok {
			return x.Cmp(y) >= 0, nil
		}
	case string:
		if y, ok := y.(string); ok {
			return x >= y, nil
		}
	}

	return nil, fmt.Errorf("invalid types for >= operation: %T and %T", x, y)
}

func greaterThan(x, y interface{}) (interface{}, error) {
	switch x := x.(type) {
	case int64:
		if y, ok := y.(int64); ok {
			return x > y, nil
		}
	case float64:
		if y, ok := y.(float64); ok {
			return x > y, nil
		}
	case decimal.Decimal:
		if y, ok := y.(decimal.Decimal); ok {
			return x.Cmp(y) > 0, nil
		}
	case string:
		if y, ok := y.(string); ok {
			return x > y, nil
		}
	}

	return nil, fmt.Errorf("invalid types for > operation: %T and %T", x, y)
}

func logicalOr(x, y interface{}) (interface{}, error) {
	switch x := x.(type) {
	case bool:
		if y, ok := y.(bool); ok {
			return x || y, nil
		}
	}

	return nil, fmt.Errorf("invalid types for || operation: %T and %T", x, y)
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
	switch x := x.(type) {
	case bool:
		if y, ok := y.(bool); ok {
			return x && y, nil
		}
	}

	return nil, fmt.Errorf("invalid types for && operation: %T and %T", x, y)
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
