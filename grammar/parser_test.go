package parser

import (
	"fmt"
	"testing"
)

func TestNewParser(t *testing.T) {
	var files []string
	files = append(files, `
contract Hello {
    func hello() string {
        return "Hello, Needle!"
    }
}
`)
	for i, file := range files {
		if i == 1 {
			break
		}
		p, err := NewParserBase([]byte(file), "")
		if err != nil {
			t.Log(err)
			return
		}
		p.Parse()
		p.PrintlnError()
	}
	fmt.Println("All files parse completed!")
}
