package compile

// opPriority contains command and its priority
type opPriority struct {
	Cmd      uint16 // identifier of the command
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
		Neg:    {Cmd: CmdSign, Priority: CmdUnary},
		Not:    {Cmd: CmdNot, Priority: CmdUnary},
		LPAREN: {Cmd: CmdSys, Priority: 0xff},
		RPAREN: {Cmd: CmdSys, Priority: 0},
	}
)
