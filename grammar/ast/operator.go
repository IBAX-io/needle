package ast

//go:generate stringer -type=OperatorType -trimprefix=OperatorType_ -output=operator_string.go operator.go
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
