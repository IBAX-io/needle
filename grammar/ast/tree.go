package ast

//go:generate stringer -type=TreeType -trimprefix=TreeType_ -output=tree_string.go tree.go
type TreeType int32

const (
	TreeType_None_Default TreeType = iota
	TreeType_SourceMain
	TreeType_ContractDef
	TreeType_ContractPart
	TreeType_FuncDef
	TreeType_DataDef
	TreeType_DataPart
	TreeType_SettingsDef
	TreeType_Parameter
	TreeType_ParameterList
	TreeType_ReturnParameters
	TreeType_StatementList
	TreeType_Block
	TreeType_Assignment
	TreeType_VarDef
	TreeType_Kind_ControlStmt
	TreeType_IfStmt
	TreeType_ElseStmt
	TreeType_IncDecStmt
	TreeType_WhileStmt
	TreeType_ContinueStmt
	TreeType_BreakStmt
	TreeType_ReturnStmt
	TreeType_FuncSignature
	TreeType_FuncTail
	TreeType_Kind_ErrorStmt
	TreeType_ErrorStmt
	TreeType_ErrWarningStmt
	TreeType_ErrInfoStmt

	TreeType_Kind_Literal
	TreeType_DecimalLiteral
	TreeType_BinaryLiteral
	TreeType_OctalLiteral
	TreeType_HexLiteral
	TreeType_FloatLiteral
	TreeType_InterpretedStringLiteral
	TreeType_RawStringLiteral
	TreeType_BooleanLiteral
	TreeType_NIL
	TreeType_Identifier
	TreeType_DollarIdentifier
	TreeType_AtIdentifier

	TreeType_Kind_Expr
	TreeType_PrimaryExpr
	TreeType_IndexExpr
	TreeType_SliceExpr
	TreeType_MapExpr
	TreeType_ArrayExpr
	TreeType_ContractCall
	TreeType_UnaryOpExpr
	TreeType_MulOpExpr
	TreeType_RelOpExpr
	TreeType_LogicalOpExpr
	TreeType_AddOpExpr
	TreeType_Operand
	TreeType_PairList
	TreeType_Pair
	TreeType_Arguments
	TreeType_ArrayList
)
