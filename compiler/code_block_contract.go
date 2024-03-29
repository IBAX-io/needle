package compiler

import (
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
