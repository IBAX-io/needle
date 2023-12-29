package compile

import (
	"errors"
)

const (
	eSysVar               = `system variable $%s cannot be changed`
	eDataParamVarCollides = `param variable '%s' in the data section of the contract '%s' collides with the 'builtin' variable`
	eWrongParams          = `function %s must have %d parameters`
	eUnknownIdent         = `unknown identifier %s`
	eDataType             = `expecting type of the data field`
	eDataName             = `expecting name of the data field`
	eDataTag              = `unexpected tag`
)

var (
	errRecursion     = errors.New(`the contract can't call itself recursively`)
	errUnclosedArray = errors.New(`unclosed array initialization`)
	errUnclosedMap   = errors.New(`unclosed map initialization`)
	errCondWrite     = errors.New(`'conditions' cannot call contracts or functions which can modify the blockchain database`)
	errMultiIndex    = errors.New(`multi-index is not supported`)
	errEndExp        = errors.New(`unexpected end of the expression`)
	errOper          = errors.New(`unexpected operator; expecting operand`)
)
