package vm

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/IBAX-io/needle/compile"
	"github.com/stretchr/testify/assert"
)

func TestVM_Compile(t *testing.T) {
	file, _ := os.ReadFile("../examples/scope.sim")
	expr, _ := os.ReadFile("../examples/expr.sim")
	stack, _ := os.ReadFile("../examples/stack.sim")
	tests := []struct {
		name    string
		method  string
		args    []rune
		wantErr assert.ErrorAssertionFunc
	}{
		{"case_contract", "@1ABC.action", []rune(string(file)), assert.NoError},
		{"case_stack", "@1ABC2.conditions", []rune(string(stack)), assert.NoError},
		{"case_expr", "operand", []rune(string(expr)), assert.NoError},
		{"case_expr_add", "operand_add", []rune(string(expr)), assert.NoError},
		{"case_expr_inc_dec", "operand_inc_dec", []rune(string(expr)), assert.NoError},
		{"case_expr_logic", "operand_logic", []rune(string(expr)), assert.NoError},
		{"case_expr_bit", "operand_bit", []rune(string(expr)), assert.NoError},
		{"case_expr_comparison", "operand_comparison", []rune(string(expr)), assert.NoError},
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
			tt.wantErr(t, vm.Compile(tt.args, compile.NewExtendData(&compile.OwnerInfo{StateID: 1, Active: true, TableID: 1}, obj, []string{"key_id"})))
			//func|contract|golang func
			t.Error(vm.Call(tt.method, nil, extend))
		})
	}
	fmt.Println("time used", time.Since(start), limit-extend[Extend_txcost].(int64))
}

func getcost(name string) int64 {
	for _, f := range obj {
		if f.Name == name {
			return 10
		}
	}
	return -1
}

var obj = []compile.ExtendFunc{
	{Name: "lenArray", Func: lenArray},
	{Name: "str", Func: str},
	{Name: "Money", Func: ValueToDecimal},
	{Name: "Replace", Func: strings.Replace},
}

// Str converts the value to a string
func str(v any) (ret string) {
	return fmt.Sprint(v)
}

func lenArray(par []any) int64 {
	return int64(len(par))
}
