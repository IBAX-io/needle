package compiler

import (
	"reflect"
	"strings"
)

// ContractInfo contains the contract information.
type ContractInfo struct {
	Id       uint32
	Name     string
	Owner    *OwnerInfo
	Used     map[string]bool // Called contracts
	Tx       *[]*FieldInfo   // contract fields
	Settings map[string]any
	CanWrite bool // true if the function can update DB
}

// TxMap returns the map of the contract fields.
func (c *ContractInfo) TxMap() map[string]*FieldInfo {
	if c == nil {
		return nil
	}
	m := make(map[string]*FieldInfo)
	if c.Tx == nil {
		return m
	}
	for _, n := range *c.Tx {
		m[n.Name] = n
	}
	return m
}

// FieldInfo describes the field of the data structure.
type FieldInfo struct {
	Name     string
	Type     reflect.Type
	Original Token
	Tags     string
}

// ContainsTag returns whether the tag is contained in this field.
func (fi *FieldInfo) ContainsTag(tag string) bool {
	return strings.Contains(fi.Tags, tag)
}
