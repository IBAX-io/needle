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
		"||":  {Cmd: CmdOr, Priority: 10},
		"&&":  {Cmd: CmdAnd, Priority: 15},
		"==":  {Cmd: CmdEqual, Priority: 20},
		"!=":  {Cmd: CmdNotEq, Priority: 20},
		"<":   {Cmd: CmdLess, Priority: 22},
		">=":  {Cmd: CmdNotLess, Priority: 22},
		">":   {Cmd: CmdGreat, Priority: 22},
		"<=":  {Cmd: CmdNotGreat, Priority: 22},
		"+":   {Cmd: CmdAdd, Priority: 25},
		"-":   {Cmd: CmdSub, Priority: 25},
		"*":   {Cmd: CmdMul, Priority: 30},
		"/":   {Cmd: CmdDiv, Priority: 30},
		"!":   {Cmd: CmdNot, Priority: uint16(CmdUnary)},
		"%=":  {Cmd: CmdAssignMod, Priority: 5},
		"++":  {Cmd: CmdInc, Priority: 5},
		"--":  {Cmd: CmdDec, Priority: 5},
		"=":   {CmdAssign, 1},
		"%":   {Cmd: CmdMod, Priority: 30},
		"&":   {Cmd: CmdBitAnd, Priority: 30},
		"<<":  {Cmd: CmdShiftL, Priority: 30},
		">>":  {Cmd: CmdShiftR, Priority: 30},
		"<<=": {Cmd: CmdAssignLShift, Priority: 5},
		">>=": {Cmd: CmdAssignRShift, Priority: 5},
		"&=":  {Cmd: CmdAssignAnd, Priority: 5},
		"|=":  {Cmd: CmdAssignOr, Priority: 5},
		"^=":  {Cmd: CmdAssignXor, Priority: 5},
		"|":   {Cmd: CmdBitOr, Priority: 30},
		"^":   {Cmd: CmdBitXor, Priority: 30},
		"-=":  {Cmd: CmdAssignSub, Priority: 5},
		"+=":  {Cmd: CmdAssignAdd, Priority: 5},
		"/=":  {Cmd: CmdAssignDiv, Priority: 5},
		"*=":  {Cmd: CmdAssignMul, Priority: 5},
	}
)
