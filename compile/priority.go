package compile

// opPriority contains command and its priority
type opPriority struct {
	Cmd      CmdT   // identifier of the command
	Priority uint16 // priority of the command
	//IsLTR    bool   // true for left-to-right associativity, false for right-to-left
}

var (
	// Array of Operator operations and their priority
	operator = map[Token]opPriority{
		Assign: {CmdAssign, 5},
		SubEq:  {CmdAssignSub, 5},
		AddEq:  {CmdAssignAdd, 5},
		MulEq:  {CmdAssignMul, 5},
		DivEq:  {CmdAssignDiv, 5},
		ModEq:  {CmdAssignMod, 5},
		LshEq:  {CmdAssignLShift, 5},
		RshEq:  {CmdAssignRShift, 5},
		AndEq:  {CmdAssignAnd, 5},
		OrEq:   {CmdAssignOr, 5},
		XorEq:  {CmdAssignXor, 5},
		Or:     {CmdOr, 10},
		And:    {CmdAnd, 11},
		BitOr:  {CmdBitOr, 12},
		BitXor: {CmdBitXor, 13},
		BitAnd: {CmdBitAnd, 14},
		EqEq:   {CmdEqual, 15},
		NotEq:  {CmdNotEq, 15},
		Less:   {CmdLess, 20},
		Great:  {CmdGreat, 20},
		GrEq:   {CmdGrEq, 20},
		LessEq: {CmdLessEq, 20},
		LSHIFT: {CmdShiftL, 21},
		RSHIFT: {CmdShiftR, 21},
		Add:    {CmdAdd, 25},
		Sub:    {CmdSub, 25},
		Mul:    {CmdMul, 30},
		Quo:    {CmdDiv, 30},
		MOD:    {CmdMod, 30},
		Not:    {CmdNot, CmdUnary},
		Inc:    {CmdInc, CmdUnary},
		Dec:    {CmdDec, CmdUnary},
	}
)
