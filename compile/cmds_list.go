package compile

import "strconv"

type CmdT uint16

func (tok CmdT) String() string {
	s := ""
	if v, ok := cmdName[tok]; ok {
		s = v
	}
	if s == "" {
		s = "cmd(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}

// here are described the commands of bytecode
const (
	CmdPush         CmdT = iota + 1 // Push value to stack
	CmdVar                          // Push variable to stack
	CmdExtend                       // Push extend variable to stack
	CmdCallExtend                   // Call extend function
	CmdPushStr                      // Push identifier as string
	CmdCall                         // call a function
	CmdCallVariadic                 // call a variadic function
	CmdReturn                       // return from function
	CmdIf                           // run block if Value is true
	CmdElse                         // run block if Value is false
	CmdAssignVar                    // list of assigned var
	CmdAssign                       // assign
	CmdLabel                        // label for continue
	CmdContinue                     // continue from label
	CmdWhile                        // while
	CmdBreak                        // break
	CmdIndex                        // get index []
	CmdSetIndex                     // set index []
	CmdFuncName                     // set func name Func(...).Name(...)
	CmdUnwrapArr                    // unwrap array to stack
	CmdMapInit                      // map initialization
	CmdArrayInit                    // array initialization
	CmdError                        // error command
)

// the commands for operations in expressions are listed below
const (
	CmdNot CmdT = iota | 0x0100
	CmdSign
	CmdInc          // ++
	CmdDec          // --
	CmdAssignAdd    // +=
	CmdAssignSub    // -=
	CmdAssignMul    // *=
	CmdAssignDiv    // /=
	CmdAssignMod    // %=
	CmdAssignAnd    // &=
	CmdAssignOr     // |=
	CmdAssignXor    // ^=
	CmdAssignLShift // <<=
	CmdAssignRShift // >>=
)

const (
	CmdAdd      CmdT = iota | 0x0200 // +
	CmdSub                           // -
	CmdMul                           // *
	CmdDiv                           // /
	CmdMod                           // %
	CmdAnd                           // &&
	CmdBitAnd                        // &
	CmdBitXor                        // ^
	CmdBitOr                         // |
	CmdOr                            // ||
	CmdEqual                         // ==
	CmdNotEq                         // !=
	CmdLess                          // <
	CmdGrEq                          // >=
	CmdGreat                         // >
	CmdLessEq                        // <=
	CmdShiftL                        // <<
	CmdShiftR                        // >>

	CmdSys   CmdT = 0xff
	CmdUnary      = 50
)

var cmdName = map[CmdT]string{
	CmdPush:         `PUSH`,
	CmdVar:          `VAR`,
	CmdExtend:       `EXTEND`,
	CmdCallExtend:   `CALLEXTEND`,
	CmdPushStr:      `PUSHSTR`,
	CmdCall:         `CALL`,
	CmdCallVariadic: `CALLVARIADIC`,
	CmdReturn:       `RETURN`,
	CmdIf:           `IF`,
	CmdElse:         `ELSE`,
	CmdAssignVar:    `ASSIGNVAR`,
	CmdAssign:       `ASSIGN`,
	CmdLabel:        `LABEL`,
	CmdContinue:     `CONTINUE`,
	CmdWhile:        `WHILE`,
	CmdBreak:        `BREAK`,
	CmdIndex:        `INDEX`,
	CmdSetIndex:     `SETINDEX`,
	CmdFuncName:     `FUNCNAME`,
	CmdUnwrapArr:    `UNWRAPARR`,
	CmdMapInit:      `MAPINIT`,
	CmdArrayInit:    `ARRAYINIT`,
	CmdError:        `ERROR`,
	CmdNot:          `!`,
	CmdSign:         `SIGN`,
	CmdAdd:          `+`,
	CmdSub:          `-`,
	CmdMul:          `*`,
	CmdDiv:          `/`,
	CmdAnd:          `&&`,
	CmdOr:           `||`,
	CmdEqual:        `==`,
	CmdNotEq:        `!=`,
	CmdLess:         `<`,
	CmdGrEq:         `>=`,
	CmdGreat:        `>`,
	CmdLessEq:       `<=`,
	CmdAssignMod:    `%=`,
	CmdInc:          `++`,
	CmdDec:          `--`,
	CmdMod:          `%`,
	CmdBitAnd:       `&`,
	CmdShiftL:       `<<`,
	CmdShiftR:       `>>`,
	CmdAssignLShift: `<<=`,
	CmdAssignRShift: `>>=`,
	CmdAssignAdd:    `+=`,
	CmdAssignSub:    `-=`,
	CmdAssignMul:    `*=`,
	CmdAssignDiv:    `/=`,
	CmdAssignAnd:    `&=`,
	CmdAssignOr:     `|=`,
	CmdAssignXor:    `^=`,
	CmdBitXor:       `^`,
	CmdBitOr:        `|`,
}
