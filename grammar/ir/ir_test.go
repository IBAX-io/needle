package ir

import (
	"fmt"
	"testing"

	"github.com/IBAX-io/needle/compiler"

	parser "github.com/IBAX-io/needle/grammar"
	"github.com/IBAX-io/needle/grammar/ast"
)

func TestIr_Build(t *testing.T) {
	input := `
contract test {
    data {
        paraInt int
        paraStr string "optional"
    }
    settings{
        s1 = 1234
        s2 = true
        s3 = "This is a string"
        s4 = 56.7834
    }

    // this is a line comment
    func testFunc(i1 i2 int, j1 j2 bool,sd ...).tail1() int,string {
	
	}
    /* this is a block comment */
    action{
        var a int
        a++
    }
}
func main(pi int,ps string)bool{
    var mf1 mf2 int mf3 bool
    mf1 = 123
    mf2 = mf1
}
`
	base, err := parser.NewParserBase([]byte(input), "")
	if err != nil {
		t.Error(err)
	}
	listener := ast.NewBuilder(base.NeedleParser, []byte(input))
	base.RegisterListener("SourceMain", listener)
	base.Parse()
	base.PrintlnError()

	ir := &Builder{
		Builder: listener,
		CodeBlock: &compiler.CodeBlock{
			Objects: make(map[string]*compiler.Object),
			Type:    compiler.CodeBlockOwner,
			Info: &compiler.OwnerInfo{
				StateId: 1,
			},
		},
	}
	for _, def := range ir.GetSourceMain().FuncDefs {
		function := ir.processFunction(def)
		fmt.Println("function", function.Name)
	}
	for _, def := range ir.GetSourceMain().ContractDefs {
		contract := ir.processContract(def)
		ir.currentContract = contract
		fmt.Println("contract", contract.Name)
	}
}
