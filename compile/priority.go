package compile

// opPriority contains command and its priority
type opPriority struct {
	Cmd      CmdT   // identifier of the command
	Priority uint16 // priority of the command
}

var (
	// Array of Operator operations and their priority
	operator = map[Token]opPriority{
		Or:     {Cmd: CmdOr, Priority: 10},
		And:    {Cmd: CmdAnd, Priority: 15},
		EqEq:   {Cmd: CmdEqual, Priority: 20},
		NotEq:  {Cmd: CmdNotEq, Priority: 20},
		Less:   {Cmd: CmdLess, Priority: 22},
		GrEq:   {Cmd: CmdNotLess, Priority: 22},
		Great:  {Cmd: CmdGreat, Priority: 22},
		LessEq: {Cmd: CmdNotGreat, Priority: 22},
		Add:    {Cmd: CmdAdd, Priority: 25},
		Sub:    {Cmd: CmdSub, Priority: 25},
		Mul:    {Cmd: CmdMul, Priority: 30},
		Quo:    {Cmd: CmdDiv, Priority: 30},
		Not:    {Cmd: CmdNot, Priority: uint16(CmdUnary)},
		LPAREN: {Cmd: CmdSys, Priority: 0xff},
		RPAREN: {Cmd: CmdSys, Priority: 0},
		ModEq:  {Cmd: CmdAssignMod, Priority: 5},
		Inc:    {Cmd: CmdInc, Priority: 5},
		Dec:    {Cmd: CmdDec, Priority: 5},
		MOD:    {Cmd: CmdMod, Priority: 30},
		BITAND: {Cmd: CmdBitAnd, Priority: 30},
		LSHIFT: {Cmd: CmdShiftL, Priority: 30},
		RSHIFT: {Cmd: CmdShiftR, Priority: 30},
		LshEq:  {Cmd: CmdAssignLShift, Priority: 5},
		RshEq:  {Cmd: CmdAssignRShift, Priority: 5},
	}
)
