package compile

// here are described the commands of bytecode
const (
	//CmdUnknown = iota // error
	CmdPush         = iota + 1 // Push value to stack
	CmdVar                     // Push variable to stack
	CmdExtend                  // Push extend variable to stack
	CmdCallExtend              // Call extend function
	CmdPushStr                 // Push identifier as string
	CmdCall                    // call a function
	CmdCallVariadic            // call a variadic function
	CmdReturn                  // return from function
	CmdIf                      // run block if Value is true
	CmdElse                    // run block if Value is false
	CmdAssignVar               // list of assigned var
	CmdAssign                  // assign
	CmdLabel                   // label for continue
	CmdContinue                // continue from label
	CmdWhile                   // while
	CmdBreak                   // break
	CmdIndex                   // get index []
	CmdSetIndex                // set index []
	CmdFuncName                // set func name Func(...).Name(...)
	CmdUnwrapArr               // unwrap array to stack
	CmdMapInit                 // map initialization
	CmdArrayInit               // array initialization
	CmdError                   // error command
)

// the commands for operations in expressions are listed below
const (
	CmdNot = iota | 0x0100
	CmdSign
)

const (
	CmdAdd = iota | 0x0200
	CmdSub
	CmdMul
	CmdDiv
	CmdAnd
	CmdOr
	CmdEqual
	CmdNotEq
	CmdLess
	CmdNotLess
	CmdGreat
	CmdNotGreat

	CmdSys          = 0xff
	CmdUnary uint16 = 50
)

var cmdName = map[uint16]string{
	CmdPush:         `cmdPUSH`,
	CmdVar:          `cmdVAR`,
	CmdExtend:       `cmdEXTEND`,
	CmdCallExtend:   `cmdCALLEXTEND`,
	CmdPushStr:      `cmdPUSHSTR`,
	CmdCall:         `cmdCALL`,
	CmdCallVariadic: `cmdCALLVARIADIC`,
	CmdReturn:       `cmdRETURN`,
	CmdIf:           `cmdIF`,
	CmdElse:         `cmdELSE`,
	CmdAssignVar:    `cmdASSIGNVAR`,
	CmdAssign:       `cmdASSIGN`,
	CmdLabel:        `cmdLABEL`,
	CmdContinue:     `cmdCONTINUE`,
	CmdWhile:        `cmdWHILE`,
	CmdBreak:        `cmdBREAK`,
	CmdIndex:        `cmdINDEX`,
	CmdSetIndex:     `cmdSETINDEX`,
	CmdFuncName:     `cmdFUNCNAME`,
	CmdUnwrapArr:    `cmdUNWRAPARR`,
	CmdMapInit:      `cmdMAPINIT`,
	CmdArrayInit:    `cmdARRAYINIT`,
	CmdError:        `cmdERROR`,
	CmdNot:          `cmdNOT`,
	CmdSign:         `cmdSIGN`,
	CmdAdd:          `cmdADD`,
	CmdSub:          `cmdSUB`,
	CmdMul:          `cmdMUL`,
	CmdDiv:          `cmdDIV`,
	CmdAnd:          `cmdAND`,
	CmdOr:           `cmdOR`,
	CmdEqual:        `cmdEQUAL`,
	CmdNotEq:        `cmdNOTEQ`,
	CmdLess:         `cmdLESS`,
	CmdNotLess:      `cmdNOTLESS`,
	CmdGreat:        `cmdGREAT`,
	CmdNotGreat:     `cmdNOTGREAT`,
}
