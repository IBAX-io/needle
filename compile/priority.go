package compile

// opPriority contains command and its priority
type opPriority struct {
	Cmd      CmdT   // identifier of the command
	Priority uint16 // priority of the command
	//IsLTR    bool   // true for left-to-right associativity, false for right-to-left
}

var (
	// Array of Operator operations and their priority
	operator = map[string]opPriority{
		"=":   {CmdAssign, 5},
		"-=":  {CmdAssignSub, 5},
		"+=":  {CmdAssignAdd, 5},
		"*=":  {CmdAssignMul, 5},
		"/=":  {CmdAssignDiv, 5},
		"%=":  {CmdAssignMod, 5},
		"<<=": {CmdAssignLShift, 5},
		">>=": {CmdAssignRShift, 5},
		"&=":  {CmdAssignAnd, 5},
		"|=":  {CmdAssignOr, 5},
		"^=":  {CmdAssignXor, 5},
		"||":  {CmdOr, 10},
		"&&":  {CmdAnd, 11},
		"|":   {CmdBitOr, 12},
		"^":   {CmdBitXor, 13},
		"&":   {CmdBitAnd, 14},
		"==":  {CmdEqual, 15},
		"!=":  {CmdNotEq, 15},
		"<":   {CmdLess, 20},
		">":   {CmdGreat, 20},
		">=":  {CmdNotLess, 20},
		"<=":  {CmdNotGreat, 20},
		"<<":  {CmdShiftL, 21},
		">>":  {CmdShiftR, 21},
		"+":   {CmdAdd, 25},
		"-":   {CmdSub, 25},
		"*":   {CmdMul, 30},
		"/":   {CmdDiv, 30},
		"%":   {CmdMod, 30},
		"!":   {CmdNot, CmdUnary},
		"++":  {CmdInc, CmdUnary},
		"--":  {CmdDec, CmdUnary},
	}
)
