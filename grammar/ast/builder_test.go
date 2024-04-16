package ast

import (
	"testing"

	parser "github.com/IBAX-io/needle/grammar"
)

func TestNewBuilder(t *testing.T) {
	input := `
contract test {
	data {
		paraInt int
		paraStr string "sdj"
	}
settings{
	s1 = 1234
	s2 = true
	s3 = "This is a string"
	s4 = 56.7834
}

func testFunc(i i1 int, j j2 bool).tail1() int,string {
	abc*def
a={"a":123}
}
	action{
		a++
	}
}
func main{
var mf int mf2 int
mf = 123
mf2 = mf
}
`
	base, err := parser.NewParserBase([]byte(string(input)), "")
	if err != nil {
		t.Error(err)
	}
	listener := NewBuilder()
	base.RegisterListener("SourceMain", listener)
	// base.RegisterListener("NewExpr", NewExpr(listener))
	base.Parse()
	base.PrintlnError()
}
