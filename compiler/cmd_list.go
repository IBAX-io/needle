package compiler

// regenerate if cmd_list.go is changed
//go:generate stringer -type Cmd -output=cmd_list_string.go cmd_list.go

type Cmd uint16

// here are described the commands of bytecode
const (
	CmdPush         Cmd = iota + 1 // Push value to stack
	CmdVar                         // Push variable to stack
	CmdExtend                      // Push extend variable to stack
	CmdCallExtend                  // Call extend function
	CmdPushStr                     // Push identifier as string
	CmdCall                        // call a function
	CmdCallVariadic                // call a variadic function
	CmdReturn                      // return from function
	CmdIf                          // run block if Value is true
	CmdElse                        // run block if Value is false
	CmdAssignVar                   // list of assigned var
	CmdAssign                      // assign
	CmdLabel                       // label for continue
	CmdContinue                    // continue from label
	CmdWhile                       // while
	CmdBreak                       // break
	CmdGetIndex                    // get index []
	CmdSetIndex                    // set index []
	CmdFuncTail                    // set func tail Func(...).tail(...)
	CmdUnwrapArr                   // unwrap array to stack
	CmdMapInit                     // map initialization
	CmdArrayInit                   // array initialization
	CmdError                       // error command
	CmdSliceColon                  // slice [:]
	CmdNot                         // !
	CmdSign                        // unary sign
)

// the commands for operations in expressions are listed below
const (
	CmdInc          Cmd = iota | 0x0100 // ++
	CmdDec                              // --
	CmdAssignAdd                        // +=
	CmdAssignSub                        // -=
	CmdAssignMul                        // *=
	CmdAssignDiv                        // /=
	CmdAssignMod                        // %=
	CmdAssignAnd                        // &=
	CmdAssignOr                         // |=
	CmdAssignXor                        // ^=
	CmdAssignLShift                     // <<=
	CmdAssignRShift                     // >>=
)

const (
	CmdAdd    Cmd = iota | 0x0200 // +
	CmdSub                        // -
	CmdMul                        // *
	CmdDiv                        // /
	CmdMod                        // %
	CmdAnd                        // &&
	CmdBitAnd                     // &
	CmdBitXor                     // ^
	CmdBitOr                      // |
	CmdOr                         // ||
	CmdEqual                      // ==
	CmdNotEq                      // !=
	CmdLess                       // <
	CmdGrEq                       // >=
	CmdGreat                      // >
	CmdLessEq                     // <=
	CmdShiftL                     // <<
	CmdShiftR                     // >>
)

type cmd = Cmd

const (
	// CmdSys is the command for delimiters
	CmdSys   cmd = 0xff
	CmdUnary     = 50
)
