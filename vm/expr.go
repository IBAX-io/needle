package vm

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func evaluateCmd(x, y any, op string) (ret any, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("%s: %w", op, err)
		}
	}()
	switch op {
	case "%", "%=":
		switch x.(type) {
		case int64:
			if y.(int64) == 0 {
				err = errDivZero
				return
			}
			ret = x.(int64) % y.(int64)
		case decimal.Decimal:
			if y.(decimal.Decimal).IsZero() {
				err = errDivZero
				return
			}
			ret = x.(decimal.Decimal).Mod(y.(decimal.Decimal))
		default:
			err = fmt.Errorf(`invalid operation: the operator %s is not defined on %T`, op, x)
			return
		}
	default:
		err = fmt.Errorf(`unsupported operator`)
		return
	}
	return
}
