package compiler

import (
	"fmt"
	"strings"
)

// ContractInfo contains the contract information.
type ContractInfo struct {
	Id       uint32
	Name     string
	Owner    *OwnerInfo
	Used     map[string]bool // Called contracts
	Field    *[]*FieldInfo   // contract fields
	Settings map[string]any
	CanWrite bool // true if the function can update DB
}

// TxMap returns the map of the contract fields.
func (c *ContractInfo) TxMap() map[string]*FieldInfo {
	if c == nil {
		return nil
	}
	m := make(map[string]*FieldInfo)
	if c.Field == nil {
		return m
	}
	for _, n := range *c.Field {
		m[n.Name] = n
	}
	return m
}

const (
	// TagOptional is the tag of the optional parameter in the contract.
	TagOptional = "optional"

	eUndefinedParam = `%s is not defined`
)

// MakeParams generates the external variables of the contract.
func (c *ContractInfo) MakeParams(fields string, params []any) (map[string]any, error) {
	pars := strings.Split(fields, `,`)
	param := make(map[string]struct{})
	for _, par := range pars {
		if _, ok := param[par]; ok {
			return nil, fmt.Errorf("duplicate parameter '%s'", par)
		}
		param[par] = struct{}{}
	}
	if len(pars) != len(params) {
		return nil, fmt.Errorf("wrong number of parameters, expected %d, got %d", len(pars), len(params))
	}

	extVars := make(map[string]any)
	fieldMap := c.TxMap()

	for i, par := range pars {
		if len(par) == 0 {
			continue
		}
		_, ok := fieldMap[par]
		if !ok {
			continue
		}
		if len(par) > 0 {
			extVars[par] = params[i]
		}
	}
	for _, fie := range fieldMap {
		if _, ok := param[fie.Name]; !ok {
			if !strings.Contains(fie.Tags, TagOptional) {
				return nil, fmt.Errorf(eUndefinedParam, fie.Name)
			}
			extVars[fie.Name] = GetFieldDefaultValue(fie.Type)
		}
	}

	return extVars, nil
}

// FieldInfo describes the field of the data structure.
type FieldInfo struct {
	Name     string
	Type     Token
	Original Token
	Tags     string
}

// ContainsTag returns whether the tag is contained in this field.
func (fi *FieldInfo) ContainsTag(tag string) bool {
	return strings.Contains(fi.Tags, tag)
}
