package vm

import (
	"fmt"
	"github.com/IBAX-io/needle/compile"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/shopspring/decimal"
)

func TestVM_Compile(t *testing.T) {
	file, _ := os.ReadFile("../examples/scope.sim")
	tests := []struct {
		name    string
		args    []rune
		wantErr assert.ErrorAssertionFunc
	}{
		{"case1", []rune(string(file)), assert.NoError},
		{"case2", []rune(string(`
contract ABC {
    action{
		$shift = 512
		$shift >> = 8
		Println($shift)
    }
}
`)), assert.NoError},
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
			//func
			//t.Error(vm.Call("adv", nil, extend))
			//Extend func
			//t.Error(vm.Call("str", []any{"ps5"}, map[string]any{}))
			//contract
			t.Log(vm.Call("@1ABC", nil, extend))
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
