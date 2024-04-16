package ast

type NodeType interface {
	any | DataDef | SettingsDef | FuncDef
}
