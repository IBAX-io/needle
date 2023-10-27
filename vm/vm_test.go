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
	scope, _ := os.ReadFile("../examples/scope.sim")
	expr, _ := os.ReadFile("../examples/expr.sim")
	stack, _ := os.ReadFile("../examples/stack.sim")
	callfn, _ := os.ReadFile("../examples/callfn.sim")

	tests := []struct {
		name    string
		method  string
		args    []rune
		wantErr assert.ErrorAssertionFunc
	}{
		{"case_contract", "@1ABC.action", []rune(string(scope)), assert.NoError},
		{"case_stack", "@1ABC2.action", []rune(string(stack)), assert.NoError},
		{"case_expr", "operand", []rune(string(expr)), assert.NoError},
		{"case_expr_add", "operand_add", []rune(string(expr)), assert.NoError},
		{"case_expr_inc_dec", "operand_inc_dec", []rune(string(expr)), assert.NoError},
		{"case_expr_logic", "operand_logic", []rune(string(expr)), assert.NoError},
		{"case_expr_bit", "operand_bit", []rune(string(expr)), assert.NoError},
		{"case_expr_comparison", "operand_comparison", []rune(string(expr)), assert.NoError},
		{"case_callfn", "@1ABC.conditions", []rune(string(callfn)), assert.NoError},
		{"case_number_literal", "number_literal", []rune(`
func intLit(){
	var i, i1, i2, i3, i4, i5, i6, i7, i8, i9 int
	i = 42
	i1 = 4_2_2
	i2 = 0600
	i3 = 0_600
	i4 = 0o600
	i5 = -0xBadFace
	i6 = 0x_67_7a_2f_cc_40_c6
	i7 = 0b1101_0011_1101_0000
}
func floatLit(){
	var f, f1, f2, f3, f4, f5, f6, f7 float
	f = .2
	f1 = 72.40
	f2 = 072.40
	f3 = 2.71828
	f4 = 6.67428e-11
	f5 = 1E6
	f6 = .12345E5
	f7 = 0.15e+0_2
}
func number_literal(){
	intLit()
	floatLit()
}`), assert.NoError},
		{"case_newline", "newline", []rune(`
func newline(){
	var a b c int;b=2;c=3;Println(a,b,c)
}`), assert.NoError},
		{"case_strF", "strF", []rune(`
func strF(){
	var a b c string i int m map
	 b = "-2s"
	 c ="3s"
	 m[b] = "a"
 	Println("a,b,c", "sdf" ,{b:m})
	info str("sd1")+str("sd2")
}`), assert.NoError},
		{"case_scope", "scope", []rune(`
func scope(){
	var a int
	a = 3
	if true{
		var a int
		a = 4
		Println(a) // 4
    }
	Println(a) // 3
}`), assert.NoError},
		{"case_slice", "slice", []rune(`
func slice(){
	var a b c d e array f bytes
	a = [1,2,3,4,5,6,7,8,9,10]
	b = a[1:3]
	c = a[1:]
	d = a[:3]
	e = a[:]
	f = Bytes("1234567890")
	Println(b,c,d,e)
}`), assert.NoError},
		{"case_return", "returnF", []rune(`
func returnF(){
	var a int
	if a>10{
		Println("a>10")
		return 
	}
	Println("a<10")
}`), assert.NoError},
		{"case_varF", "varF1", []rune(`
func varF(){
	var a int
	var b b1, b2 string
	var c bool, c1 float c2 map c3 array
	
	b = "string"
	b1, b2 = "string1", "string2"
	c c1 c3 = true 1.1 [1]
	c2 = {"c2": "c2 value"}
	Println(a, b, b1, b2, c, c1, c2,c3)
}

func varF1(){
	var a b int c c1 map d d1 array
	a, b = 1, 2
	// c,d = {"a":a, "b":b}, [1, 2, 3] //error
	c = {"a":a, "b":b}
	d = [1, 2, 3]
	c1,d1 = c, d
	Println(c1, d1) //{"a":1, "b":2} [1, 2, 3]
}`), assert.NoError},
		{"case_settingsF", "@1contractName.action", []rune(`
contract contractName {
    settings{
		a = 123
		b = "string"
      	s = "\'23\'"
	}
	action{
		var a int s string
		a = Settings("@1contractName", "a")
		Println(a)
		s = Settings("@1contractName", "s")
		Println(s)
	}
}`), assert.NoError},
		{"case_loop_actionCall", "@1loop_actionCall.action", []rune(`
		contract loop_actionCall {
 			func aa(){
 			}
			action{
				aa().action()
			}
		}`), assert.NoError},
		{"case_loopCall", "@1loopCall.action", []rune(`
contract loopCall2 {}
contract replaceF {
    action{
        var m map
        CallContract("loopCall2",m)
		//loopCall2()
    }
}`), assert.NoError},
		{"case_replaceF", "@1replaceF.action", []rune(`
//func Println(a bool){
//   Println(a)
//}
contract replaceF {
	func Println(a int){
		Println("23ssd3", a)
	}
    action{
         Println(13)
    }
}`), assert.NoError},
		{"case_dataSection", "@1dataSection.action", []rune(`
contract dataSection {
	data{
		Param1 int
	}
    action{ 
		Println($Param1,$pop)
    }
}`), assert.NoError},
		{"case_execRecursion", "@1Recursion.action", []rune(`
contract Recursion {
    action{ 
 		$e = Println2()
    }
}`), assert.NoError},
		{"case_tail", "@1tail.action", []rune(`
contract tail {
func NameA(i int).TailNameA() int,string{
	return 123 "a"
}
func NameB(s string,b bool) string{
	return "b"
}
func NameC(i int,s string) string{return "c"}

    action{ 
 		var i int s string
  i s = NameA(1)

  s = NameA(2).NameB(true)

  s = NameA(3).NameC()
	}
}`), assert.NoError},
		{"case_control", "@1control.action", []rune(`
contract control{
	func fn1().sm int string{
		return 2 ""
	}
	 
	action {
	var a int;a = 1
    //while a < 5 {
    //   a++
    //   if a == 2 {
    //       continue
    //   }
	//	if a == 3 {
	//		break
	//	}
    //}

$s = fn1()
		if 3 {
			Println("ifs -",$s)

		//}elif $s{
		}elif fn1(){
			Println("elifs --")
		}else{
			//break
			Println("elses ---")
		}
	}
}`), assert.NoError},
		{"case_tail", "@1control.action", []rune(`
contract control{
     action {
        var top array limit  int ignore map
        ignore["0"]=1
        ignore["5555"]=1
        limit = 4  
        top = [{"id":"0"},{"id":"2"},{"id":"3"},{"id":"4"}]
        var i int
        while i < lenArray(top){
            var m map id account string
            m = top[i]
			i = i + 1
            id = m["id"]
            if ignore[id] {
                continue
            }
            Println("id--",id)
        }
}
}`), assert.NoError},
		{"case_map_init", "@1map_init.action", []rune(`
contract map_init {
    action{
		var m map
		 $a = "222"
		m  = {if:"is",">"}
		 m[$a] = 123233 
		Println(m)
    }
}`), assert.NoError},
		{"case_error_test", "@1error_test.action", []rune(`
contract error_test {
    action{
		if " " {
			error  -2 
		}elif "3" {
			warning 3
		}else{
			info 4 
		}
    }
}`), assert.NoError},
		{"case_exec", "@1Name.action", []rune(`
contract contractName {
    data {
		paramName1 int paramName2 string ""
    }
    func action() {
        $result = $paramName1 + Int($paramName2)
    }
}

contract Name {
    func action() {
        var a int s map
        a = ExecContract("@1contractName", "paramName1,paramName2", 123, "456")
Println(.23)
    }
}`), assert.NoError},
		{"case_ifstmt", "@1ifstmt.action", []rune(`
contract ifstmt {
	action{
 		if "" {
			Println("if")
		}
		elif ""{
			Println("elif")
		}
		else {
			Println("else")
		}

		if 0x123 {
			var d bool
			if !d{
				Println("double==")
			}
			if ""{
				Println("if2")
			}else{
				Println("else2")
			}
		}
	}
} 
`), assert.NoError},
	}
	limit := int64(2700)
	extend := map[string]any{
		ExtendTxCost:           limit,
		ExtendGenBlock:         true,
		ExtendTimeLimit:        int64(1000),
		ExtendParentContract:   ``,
		ExtendOriginalContract: ``,
		ExtendThisContract:     ``,
		"extFn": func(a1, a2 int64, a ...int64) int64 {
			fmt.Println("extFn", a1, a2, a)
			return 2323
		},
		"pop":    int64(23),
		"Param1": int64(22223),
		"popFn": func(i int64) map[string]string {
			fmt.Println(i, "23")
			return map[string]string{"a1ss": "b1ss"}
		},
	}
	start := time.Now()
	vm := NewVM()
	vm.SetExtendCost(getcost)
	build := &compile.ExtendData{
		Owner:  &compile.OwnerInfo{StateID: 1},
		PreVar: []string{"key_id"},
		Func:   obj,
		Extern: true,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.wantErr(t, vm.Compile(tt.args, build)) {
				return
			}
			ret, err := vm.Call(tt.method, extend)
			if err != nil {
				t.Error(err)
				return
			}
			t.Log(ret)
		})
	}
	fmt.Println("time used:", time.Since(start), "gas:", limit-extend[ExtendTxCost].(int64))
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
	{Name: "Int", Func: ValueToInt},
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
