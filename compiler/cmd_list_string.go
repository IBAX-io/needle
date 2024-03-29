// Code generated by "stringer -type Cmd -output=cmd_list_string.go cmd_list.go"; DO NOT EDIT.

package compiler

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CmdPush-1]
	_ = x[CmdVar-2]
	_ = x[CmdExtend-3]
	_ = x[CmdCallExtend-4]
	_ = x[CmdPushStr-5]
	_ = x[CmdCall-6]
	_ = x[CmdCallVariadic-7]
	_ = x[CmdReturn-8]
	_ = x[CmdIf-9]
	_ = x[CmdElse-10]
	_ = x[CmdAssignVar-11]
	_ = x[CmdAssign-12]
	_ = x[CmdLabel-13]
	_ = x[CmdContinue-14]
	_ = x[CmdWhile-15]
	_ = x[CmdBreak-16]
	_ = x[CmdGetIndex-17]
	_ = x[CmdSetIndex-18]
	_ = x[CmdFuncTail-19]
	_ = x[CmdUnwrapArr-20]
	_ = x[CmdMapInit-21]
	_ = x[CmdArrayInit-22]
	_ = x[CmdError-23]
	_ = x[CmdSliceColon-24]
	_ = x[CmdNot-25]
	_ = x[CmdSign-26]
	_ = x[CmdInc-256]
	_ = x[CmdDec-257]
	_ = x[CmdAssignAdd-258]
	_ = x[CmdAssignSub-259]
	_ = x[CmdAssignMul-260]
	_ = x[CmdAssignDiv-261]
	_ = x[CmdAssignMod-262]
	_ = x[CmdAssignAnd-263]
	_ = x[CmdAssignOr-264]
	_ = x[CmdAssignXor-265]
	_ = x[CmdAssignLShift-266]
	_ = x[CmdAssignRShift-267]
	_ = x[CmdAdd-512]
	_ = x[CmdSub-513]
	_ = x[CmdMul-514]
	_ = x[CmdDiv-515]
	_ = x[CmdMod-516]
	_ = x[CmdAnd-517]
	_ = x[CmdBitAnd-518]
	_ = x[CmdBitXor-519]
	_ = x[CmdBitOr-520]
	_ = x[CmdOr-521]
	_ = x[CmdEqual-522]
	_ = x[CmdNotEq-523]
	_ = x[CmdLess-524]
	_ = x[CmdGrEq-525]
	_ = x[CmdGreat-526]
	_ = x[CmdLessEq-527]
	_ = x[CmdShiftL-528]
	_ = x[CmdShiftR-529]
}

const (
	_Cmd_name_0 = "CmdPushCmdVarCmdExtendCmdCallExtendCmdPushStrCmdCallCmdCallVariadicCmdReturnCmdIfCmdElseCmdAssignVarCmdAssignCmdLabelCmdContinueCmdWhileCmdBreakCmdGetIndexCmdSetIndexCmdFuncTailCmdUnwrapArrCmdMapInitCmdArrayInitCmdErrorCmdSliceColonCmdNotCmdSign"
	_Cmd_name_1 = "CmdIncCmdDecCmdAssignAddCmdAssignSubCmdAssignMulCmdAssignDivCmdAssignModCmdAssignAndCmdAssignOrCmdAssignXorCmdAssignLShiftCmdAssignRShift"
	_Cmd_name_2 = "CmdAddCmdSubCmdMulCmdDivCmdModCmdAndCmdBitAndCmdBitXorCmdBitOrCmdOrCmdEqualCmdNotEqCmdLessCmdGrEqCmdGreatCmdLessEqCmdShiftLCmdShiftR"
)

var (
	_Cmd_index_0 = [...]uint8{0, 7, 13, 22, 35, 45, 52, 67, 76, 81, 88, 100, 109, 117, 128, 136, 144, 155, 166, 177, 189, 199, 211, 219, 232, 238, 245}
	_Cmd_index_1 = [...]uint8{0, 6, 12, 24, 36, 48, 60, 72, 84, 95, 107, 122, 137}
	_Cmd_index_2 = [...]uint8{0, 6, 12, 18, 24, 30, 36, 45, 54, 62, 67, 75, 83, 90, 97, 105, 114, 123, 132}
)

func (i Cmd) String() string {
	switch {
	case 1 <= i && i <= 26:
		i -= 1
		return _Cmd_name_0[_Cmd_index_0[i]:_Cmd_index_0[i+1]]
	case 256 <= i && i <= 267:
		i -= 256
		return _Cmd_name_1[_Cmd_index_1[i]:_Cmd_index_1[i+1]]
	case 512 <= i && i <= 529:
		i -= 512
		return _Cmd_name_2[_Cmd_index_2[i]:_Cmd_index_2[i+1]]
	default:
		return "Cmd(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
