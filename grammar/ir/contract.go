package ir

import (
	"fmt"

	"github.com/IBAX-io/needle/grammar/ast"
)

type Contract struct {
	Id        int32
	Owner     *Owner
	Name      string
	Fields    []*Field
	Settings  []*Setting
	Functions []*Function
	Body      *Body
}

func (b *Builder) processContract(ast *ast.ContractDef) *Contract {
	c := &Contract{
		Id:   ast.Id,
		Name: ast.Name,
	}
	if b.currentContract != nil && b.currentContract.Name == c.Name {
		fmt.Println("already processed contract", c.Name)
		return b.currentContract
	}
	for _, part := range ast.ContractParts {
		if part.FuncDef != nil {
			c.Functions = append(c.Functions, b.processFunction(part.FuncDef))
		}
		if part.DataDef != nil {
			c.Fields = b.processFields(part.DataDef)
		}
		if part.SettingsDef != nil {
			c.Settings = b.processSetting(part.SettingsDef)
		}
	}
	return c
}

type Field struct {
	Name     string
	Type     string
	Original string
	Tag      string
}

func (b *Builder) processFields(ast *ast.DataDef) []*Field {
	f := make([]*Field, 0)
	for _, part := range ast.DataParts {
		f = append(f, &Field{
			Name:     part.Name,
			Type:     part.TreeType.String(),
			Original: part.Typename,
			Tag:      part.Tag,
		})
	}
	return f
}

type Setting struct {
	Name  string
	Type  string
	Value *Literal
}

func (b *Builder) processSetting(ast *ast.SettingsDef) []*Setting {
	s := make([]*Setting, 0)
	for i := 0; i < len(ast.SettingsName); i++ {
		s = append(s, &Setting{
			Name:  ast.SettingsName[i],
			Type:  ast.SettingValue[i].Value,
			Value: b.processLiteral(ast.SettingValue[i]),
		})
	}
	return s
}
