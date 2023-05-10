package compile

type stateType int

const (
	// The list of state types
	stateRoot stateType = iota
	stateBody
	stateBlock
	stateContract
	stateFunc
	stateFParams
	stateFParam
	stateFParamTYPE
	stateFTail
	stateFResult
	stateFDot
	stateVar
	stateVarType
	stateAssignEval
	stateAssign
	stateTX
	stateSettings
	stateConsts
	stateConstsAssign
	stateConstsValue
	stateFields
	stateEval
)

const (
	// The list of state flags
	statePush = 1 << (8 + iota)
	statePop
	stateStay
	stateToBlock
	stateToBody
	stateFork
	stateToFork
	stateLabel
	stateMustEval
)

const FlushMark = 1 << 20

const (
	// Errors of compilation
	errUnknownCmd = iota + 1 // unknown command
	errMustName              // must be the name
	errMustLCurly            // must be '{'
	errMustRCurly            // must be '}'
	errParams                // wrong parameters
	errVars                  // wrong variables
	errVarType               // must be type
	errAssign                // must be '='
	errStrNum                // must be number or string
)

var errTable = map[int]string{
	errUnknownCmd: "unknown command",
	errMustName:   "must be the name",
	errMustLCurly: "must be '{'",
	errMustRCurly: "must be '}'",
	errParams:     "wrong parameters",
	errVars:       "wrong variables",
	errVarType:    "must be type",
	errAssign:     "must be '='",
	errStrNum:     "must be number or string",
}

// contains a new state and a handle function
type compileState struct {
	newState   stateType   // a new state
	funcHandle compileFunc // a handle function
}

// hasState returns true if the state is set
func (c compileState) hasState(state int) bool {
	return int(c.newState)&state > 0
}

var stateTable = make(map[stateType]map[Token]compileState)

func init() {
	stateTable[stateRoot] = map[Token]compileState{
		NEWLINE:  {stateRoot, fnNothing},
		CONTRACT: {stateContract | statePush, fnNothing},
		FUNC:     {stateFunc | statePush, fnNothing},
		UNKNOWN:  {errUnknownCmd, fnError},
	}
	stateTable[stateBody] = map[Token]compileState{
		NEWLINE:    {stateBody, fnNothing},
		FUNC:       {stateFunc | statePush, fnNothing},
		RETURN:     {stateEval, fnReturn},
		CONTINUE:   {stateBody, fnContinue},
		BREAK:      {stateBody, fnBreak},
		IF:         {stateEval | statePush | stateToBlock | stateMustEval, fnIf},
		WHILE:      {stateEval | statePush | stateToBlock | stateLabel | stateMustEval, fnWhile},
		ELSE:       {stateBlock | statePush, fnElse},
		VAR:        {stateVar, fnNothing},
		TX:         {stateTX, fnTx},
		SETTINGS:   {stateSettings, fnSettings},
		ERROR:      {stateEval, fnCmdError},
		ERRWARNING: {stateEval, fnCmdError},
		ERRINFO:    {stateEval, fnCmdError},
		IDENTIFIER: {stateAssignEval | stateFork, fnNothing},
		EXTEND:     {stateAssignEval | stateFork, fnNothing},
		RBRACE:     {statePop, fnNothing},
		UNKNOWN:    {errMustRCurly, fnError},
	}
	stateTable[stateBlock] = map[Token]compileState{
		NEWLINE: {stateBlock, fnNothing},
		LBRACE:  {stateBody, fnNothing},
		UNKNOWN: {errMustLCurly, fnError},
	}
	stateTable[stateContract] = map[Token]compileState{
		NEWLINE:    {stateContract, fnNothing},
		IDENTIFIER: {stateBlock, fnNameBlock},
		UNKNOWN:    {errMustName, fnError},
	}
	stateTable[stateFunc] = map[Token]compileState{
		NEWLINE:    {stateFunc, fnNothing},
		IDENTIFIER: {stateFParams, fnNameBlock},
		UNKNOWN:    {errMustName, fnError},
	}
	stateTable[stateFParams] = map[Token]compileState{
		NEWLINE: {stateFParams, fnNothing},
		LPAREN:  {stateFParam, fnNothing},
		UNKNOWN: {stateFResult | stateStay, fnNothing},
	}
	stateTable[stateFParam] = map[Token]compileState{
		NEWLINE:    {stateFParam, fnNothing},
		IDENTIFIER: {stateFParamTYPE, fnFparam},
		//TYPENAME:  {stateFParam,fFType},
		COMMA:   {stateFParam, fnNothing},
		RPAREN:  {stateFResult, fnNothing},
		UNKNOWN: {errParams, fnError},
	}
	stateTable[stateFParamTYPE] = map[Token]compileState{
		IDENTIFIER: {stateFParamTYPE, fnFparam},
		TYPENAME:   {stateFParam, fnFtype},
		TAIL:       {stateFTail, fnFtail},
		COMMA:      {stateFParamTYPE, fnNothing},
		//			isRPar:  {stateFResult,fnunc: cnil},
		UNKNOWN: {errVarType, fnError},
	}
	stateTable[stateFTail] = map[Token]compileState{
		NEWLINE: {stateFTail, fnNothing},
		RPAREN:  {stateFResult, fnNothing},
		UNKNOWN: {errParams, fnError},
	}
	stateTable[stateFResult] = map[Token]compileState{
		NEWLINE:  {stateFResult, fnNothing},
		DOT:      {stateFDot, fnNothing},
		TYPENAME: {stateFResult, fnFuncResult},
		COMMA:    {stateFResult, fnNothing},
		UNKNOWN:  {stateBlock | stateStay, fnNothing},
	}
	stateTable[stateFDot] = map[Token]compileState{
		NEWLINE:    {stateFDot, fnNothing},
		IDENTIFIER: {stateFParams, fnFNameParam},
		UNKNOWN:    {errMustName, fnError},
	}
	stateTable[stateVar] = map[Token]compileState{
		NEWLINE:    {stateBody, fnNothing},
		IDENTIFIER: {stateVarType, fnFparam},
		RBRACE:     {stateBody | stateStay, fnNothing},
		COMMA:      {stateVar, fnNothing},
		UNKNOWN:    {errVars, fnError},
	}
	stateTable[stateVarType] = map[Token]compileState{
		IDENTIFIER: {stateVarType, fnFparam},
		TYPENAME:   {stateVar, fnFtype},
		COMMA:      {stateVarType, fnNothing},
		UNKNOWN:    {errVarType, fnError},
	}
	stateTable[stateAssignEval] = map[Token]compileState{
		LPAREN:  {stateEval | stateToFork | stateToBody, fnNothing},
		LBRACK:  {stateEval | stateToFork | stateToBody, fnNothing},
		UNKNOWN: {stateAssign | stateToFork | stateStay, fnNothing},
	}
	stateTable[stateAssign] = map[Token]compileState{
		COMMA:      {stateAssign, fnNothing},
		IDENTIFIER: {stateAssign, fnAssignVar},
		EXTEND:     {stateAssign, fnAssignVar},
		EQ:         {stateEval | stateToBody, fnAssign},
		OPERATOR:   {stateEval | stateToBody, fnAssign},
		UNKNOWN:    {errAssign, fnError},
	}
	stateTable[stateTX] = map[Token]compileState{
		NEWLINE: {stateTX, fnNothing},
		LBRACE:  {stateFields, fnNothing},
		//IDENTIFIER:   {stateAssign,fTX},
		EXTEND:  {stateAssign, fnTx},
		UNKNOWN: {errMustLCurly, fnError},
	}
	stateTable[stateSettings] = map[Token]compileState{
		NEWLINE: {stateSettings, fnNothing},
		LBRACE:  {stateConsts, fnNothing},
		UNKNOWN: {errMustLCurly, fnError},
	}
	stateTable[stateConsts] = map[Token]compileState{
		NEWLINE:    {stateConsts, fnNothing},
		COMMA:      {stateConsts, fnNothing},
		IDENTIFIER: {stateConstsAssign, fnConstName},
		RBRACE:     {stateToBody, fnNothing},
		UNKNOWN:    {errMustRCurly, fnError},
	}
	stateTable[stateConstsAssign] = map[Token]compileState{
		EQ:      {stateConstsValue, fnNothing},
		UNKNOWN: {errAssign, fnError},
	}
	stateTable[stateConstsValue] = map[Token]compileState{
		LITERAL: {stateConsts, fnConstValue},
		NUMBER:  {stateConsts, fnConstValue},
		UNKNOWN: {errStrNum, fnError},
	}
	stateTable[stateFields] = map[Token]compileState{
		NEWLINE:    {stateFields, fnFieldLine},
		COMMA:      {stateFields, fnFieldComma},
		IDENTIFIER: {stateFields, fnField},
		TYPENAME:   {stateFields, fnFieldType},
		LITERAL:    {stateFields, fnFieldTag},
		RBRACE:     {stateToBody, fnFields},
		UNKNOWN:    {errMustRCurly, fnError},
	}
}
