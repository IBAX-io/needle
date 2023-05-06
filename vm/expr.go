package vm

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func evaluateCmd(x, y interface{}, op string) (interface{}, error) {
	switch op {
	case "+", "+=":
		switch v := x.(type) {
		case int64:
			return v + y.(int64), nil
		case float64:
			return v + y.(float64), nil
		case decimal.Decimal:
			return v.Add(y.(decimal.Decimal)), nil
		case string:
			return v + y.(string), nil
		default:
			return nil, fmt.Errorf("invalid type %T for + operation", x)
		}
	case "-", "-=":
		switch v := x.(type) {
		case int64:
			return v - y.(int64), nil
		case float64:
			return v - y.(float64), nil
		case decimal.Decimal:
			return v.Sub(y.(decimal.Decimal)), nil
		default:
			return nil, fmt.Errorf("invalid type %T for - operation", x)
		}
	case "*", "*=":
		switch v := x.(type) {
		case int64:
			return v * y.(int64), nil
		case float64:
			return v * y.(float64), nil
		case decimal.Decimal:
			return v.Mul(y.(decimal.Decimal)), nil
		default:
			return nil, fmt.Errorf("invalid type %T for * operation", x)
		}
	case "/", "/=":
		switch v := x.(type) {
		case int64:
			if y.(int64) == 0 {
				return nil, fmt.Errorf("divided by zero")
			}
			return v / y.(int64), nil
		case float64:
			if y.(float64) == 0 {
				return nil, fmt.Errorf("divided by zero")
			}
			return v / y.(float64), nil
		case decimal.Decimal:
			if y.(decimal.Decimal).IsZero() {
				return nil, fmt.Errorf("divided by zero")
			}
			return v.Div(y.(decimal.Decimal)), nil
		default:
			return nil, fmt.Errorf("invalid type %T for / operation", x)
		}
	case "%", "%=":
		switch v := x.(type) {
		case int64:
			if y.(int64) == 0 {
				return nil, fmt.Errorf("divided by zero")
			}
			return v % y.(int64), nil
		case decimal.Decimal:
			if y.(decimal.Decimal).IsZero() {
				return nil, fmt.Errorf("divided by zero")
			}
			return v.Mod(y.(decimal.Decimal)), nil
		default:
			return nil, fmt.Errorf("invalid type %T for %s operation", x, op)
		}
	case "<<", "<<=":
		switch v := x.(type) {
		case int64:
			if y.(int64) < 0 {
				return nil, fmt.Errorf("negative shift count")
			}
			return v << uint64(y.(int64)), nil
		default:
			return nil, fmt.Errorf("invalid type %T for %s operation", x, op)
		}
	case ">>", ">>=":
		switch v := x.(type) {
		case int64:
			if y.(int64) < 0 {
				return nil, fmt.Errorf("negative shift count")
			}
			return v >> uint64(y.(int64)), nil
		default:
			return nil, fmt.Errorf("invalid type %T for %s operation", x, op)
		}
	case "&", "&=":
		switch v := x.(type) {
		case int64:
			return v & y.(int64), nil
		default:
			return nil, fmt.Errorf("invalid type %T for %s operation", x, op)
		}
	case "|", "|=":
		switch v := x.(type) {
		case int64:
			return v | y.(int64), nil
		default:
			return nil, fmt.Errorf("invalid type %T for %s operation", x, op)
		}
	case "^", "^=":
		switch v := x.(type) {
		case int64:
			return v ^ y.(int64), nil
		default:
			return nil, fmt.Errorf("invalid type %T for %s operation", x, op)
		}
	case "&&":
		switch v := x.(type) {
		case bool:
			return v && y.(bool), nil
		default:
			return nil, fmt.Errorf("invalid type %T for && operation", x)
		}
	case "||":
		switch v := x.(type) {
		case bool:
			return v || y.(bool), nil
		default:
			return nil, fmt.Errorf("invalid type %T for || operation", x)
		}
	case "<":
		switch v := x.(type) {
		case int64:
			return v < y.(int64), nil
		case float64:
			return v < y.(float64), nil
		case decimal.Decimal:
			return v.LessThan(y.(decimal.Decimal)), nil
		case string:
			return v < y.(string), nil
		default:
			return nil, fmt.Errorf("invalid type %T for < operation", x)
		}
	case ">":
		switch v := x.(type) {
		case int64:
			return v > y.(int64), nil
		case float64:
			return v > y.(float64), nil
		case decimal.Decimal:
			return v.GreaterThan(y.(decimal.Decimal)), nil
		case string:
			return v > y.(string), nil
		default:
			return nil, fmt.Errorf("invalid type %T for > operation", x)
		}
	case "<=":
		switch v := x.(type) {
		case int64:
			return v <= y.(int64), nil
		case float64:
			return v <= y.(float64), nil
		case decimal.Decimal:
			return v.LessThanOrEqual(y.(decimal.Decimal)), nil
		case string:
			return v <= y.(string), nil
		default:
			return nil, fmt.Errorf("invalid type %T for <= operation", x)
		}
	case ">=":
		switch v := x.(type) {
		case int64:
			return v >= y.(int64), nil
		case float64:
			return v >= y.(float64), nil
		case decimal.Decimal:
			return v.GreaterThanOrEqual(y.(decimal.Decimal)), nil
		case string:
			return v >= y.(string), nil
		default:
			return nil, fmt.Errorf("invalid type %T for >= operation", x)
		}
	case "==":
		return x == y, nil
	case "!=":
		return x != y, nil
	default:
		return nil, fmt.Errorf("unsupported operator %s", op)
	}
}
