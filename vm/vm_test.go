package vm

import (
	"fmt"
	"github.com/IBAX-io/needle/compile"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
	"time"
)

func TestVM_Compile(t *testing.T) {
	file, _ := os.ReadFile("../examples/scope.sim")
	expr, _ := os.ReadFile("../examples/expr.sim")
	tests := []struct {
		name    string
		method  string
		args    []rune
		wantErr assert.ErrorAssertionFunc
	}{
		{"case_contract", "@1ABC", []rune(string(file)), assert.NoError},
		{"case_expr", "operand", []rune(string(expr)), assert.NoError},
		{"case_expr_add", "operand_add", []rune(string(expr)), assert.NoError},
	}
	limit := int64(100000)
	extend := map[string]any{
		Extend_txcost:    limit,
		Extend_gen_block: 1,
	}
	start := time.Now()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vm := GetVM()
			vm.SetExtendCost(getcost)
			vm.Objects = compile.NewExtendData(
				obj, map[string]string{}, writeFuncs,
			).MakeObj()
			tt.wantErr(t, vm.Compile(tt.args, &compile.OwnerInfo{StateID: 1, Active: true, TableID: 1}))
			//func|contract|golang func
			t.Error(vm.Call(tt.method, nil, extend))
		})
	}
	fmt.Println("time used", time.Since(start), limit-extend[Extend_txcost].(int64))
}

func getcost(name string) int64 {
	_, ok := obj[name]
	if ok {
		return 10
	}
	return -1
}

var obj = map[string]any{
	"Println":  fmt.Println,
	"Sprintf":  fmt.Sprintf,
	"lenArray": lenArray,
	"str":      str,
	"Money":    Money,
	"Replace":  strings.Replace}

var writeFuncs = map[string]struct{}{"Sprintf": {}}

// Str converts the value to a string
func str(v any) (ret string) {
	return fmt.Sprint(v)
}

func lenArray(par []any) int64 {
	return int64(len(par))
}

func Money(v any) (ret decimal.Decimal) {
	ret, _ = ValueToDecimal(v)
	return ret
}
