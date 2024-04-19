package ast

type NodeType interface {
	any
}

type TreeType int32

const (
	TreeType_None_Default TreeType = iota

	TreeType_Block
	TreeType_Kind_ControlStmt
	TreeType_IfStmt
	TreeType_ElseStmt

	TreeType_WhileStmt
	TreeType_ContinueStmt
	TreeType_BreakStmt
	TreeType_ReturnStmt

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
	TreeType_UnaryOpExpr
	TreeType_MulOpExpr
	TreeType_RelOpExpr
	TreeType_LogicalOpExpr
	TreeType_AddOpExpr
)

type OperatorType int32

const (
	OperatorType_None_Default OperatorType = iota
	OperatorType_NOT
	OperatorType_MUL
	OperatorType_ADD
	OperatorType_SUB
	OperatorType_QUO
	OperatorType_LESS
	OperatorType_GREATER
	OperatorType_NOT_EQ
	OperatorType_AND
	OperatorType_LESS_EQ
	OperatorType_EQ_EQ
	OperatorType_GR_EQ
	OperatorType_OR
	OperatorType_BIT_AND
	OperatorType_BIT_OR
	OperatorType_BIT_XOR
	OperatorType_MOD
	OperatorType_LSHIFT
	OperatorType_RSHIFT
	OperatorType_ADD_EQ
	OperatorType_SUB_EQ
	OperatorType_MUL_EQ
	OperatorType_DIV_EQ
	OperatorType_MOD_EQ
	OperatorType_LSHIFT_EQ
	OperatorType_RSHIFT_EQ
	OperatorType_BIT_AND_EQ
	OperatorType_BIT_OR_EQ
	OperatorType_BIT_XOR_EQ
	OperatorType_INC
	OperatorType_DEC
)
