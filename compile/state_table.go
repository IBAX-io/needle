package compile

import (
	"fmt"
	"strconv"
)

type stateType int

// The list of state types
const (
	stateRoot stateType = iota
	stateBody
	stateBlock
	stateContract
	stateFunc
	stateFnParams
	stateFnParam
	stateFnParamType
	stateFnTail
	stateFnResult
	stateFnDot
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
	stateFlagPush     = 1 << (8 + iota) // push a new block and add it to the parent block
	stateFlagPop                        // end the current block and return to the parent block
	stateFlagStay                       // keep the current state unchanged when switching to a new state
	stateFlagToBlock                    // move to stateBlock state
	stateFlagToBody                     // move to stateBody state
	stateFlagFork                       // save the position of the lexical token, when the expression starts with the name of the identifier or $
	stateFlagToFork                     // get the lexical token based on the position of stateFlagFork and pass the lexical token to the process function
	stateFlagLabel                      // used to insert cmdLabel instruction, while structure needs this flag
	stateFlagMustEval                   // check the availability of the condition expression in the if and while structures
)

var stateType2Str = map[stateType]string{
	stateRoot:         "Root",
	stateBody:         "Body",
	stateBlock:        "Block",
	stateContract:     "Contract",
	stateFunc:         "Func",
	stateFnParams:     "FnParams",
	stateFnParam:      "FnParam",
	stateFnParamType:  "FnParamType",
	stateFnTail:       "FnTail",
	stateFnResult:     "FnResult",
	stateFnDot:        "FnDot",
	stateVar:          "Var",
	stateVarType:      "VarType",
	stateAssignEval:   "AssignEval",
	stateAssign:       "Assign",
	stateTX:           "TX",
	stateSettings:     "Settings",
	stateConsts:       "Consts",
	stateConstsAssign: "ConstsAssign",
	stateConstsValue:  "ConstsValue",
	stateFields:       "Fields",
	stateEval:         "Eval",

	stateFlagPush:     "Push",
	stateFlagPop:      "Pop",
	stateFlagStay:     "Stay",
	stateFlagToBlock:  "ToBlock",
	stateFlagToBody:   "ToBody",
	stateFlagFork:     "Fork",
	stateFlagToFork:   "ToFork",
	stateFlagLabel:    "Label",
	stateFlagMustEval: "MustEval",
}

func (tok stateType) String() string {
	var s string
	if 0 <= tok {
		s = stateType2Str[tok]
	}
	if s == "" {
		s = "state(" + strconv.Itoa(int(tok)) + ")"
	}
	return fmt.Sprintf("%-15s", s)
}

const (
	// Errors of compilation
	errUnknownCmd         = iota + 1 // unknown command
	errMustName                      // must be the name
	errMustLBRACE                    // must be '{'
	errMustRBRACE                    // must be '}'
	errParams                        // wrong parameters
	errVars                          // wrong variables
	errVarType                       // must be type
	errAssign                        // must be '='
	errStrNum                        // must be number or string
	errMisplacedDotDotDot            // can only use ... with final parameter in list
)

var errTable = map[int]string{
	errUnknownCmd:         "unknown command",
	errMustName:           "must be the name",
	errMustLBRACE:         "must be '{'",
	errMustRBRACE:         "must be '}'",
	errParams:             "wrong parameters",
	errVars:               "wrong variables",
	errVarType:            "must be type",
	errAssign:             "must be '='",
	errStrNum:             "must be number or string",
	errMisplacedDotDotDot: "can only use ... with final parameter in list",
}

// contains a next state and a handle function
type compileState struct {
	fn   compileFunc // a handle function
	next stateType   // a next state
}

// hasState returns true if the state is set
func (c compileState) hasState(state int) bool {
	return int(c.next)&state > 0
}

var stateTable = make(map[stateType]map[Token]compileState)

func init() {
	stateTable[stateRoot] = map[Token]compileState{
		NEWLINE:  {fnNothing, stateRoot},
		CONTRACT: {fnNothing, stateContract | stateFlagPush},
		FUNC:     {fnNothing, stateFunc | stateFlagPush},
		UNKNOWN:  {fnError, errUnknownCmd},
	}
	stateTable[stateBody] = map[Token]compileState{
		NEWLINE:    {fnNothing, stateBody},
		FUNC:       {fnNothing, stateFunc | stateFlagPush},
		RETURN:     {fnReturn, stateEval},
		CONTINUE:   {fnContinue, stateBody},
		BREAK:      {fnBreak, stateBody},
		IF:         {fnIf, stateEval | stateFlagPush | stateFlagToBlock | stateFlagMustEval},
		WHILE:      {fnWhile, stateEval | stateFlagPush | stateFlagToBlock | stateFlagLabel | stateFlagMustEval},
		ELSE:       {fnElse, stateBlock | stateFlagPush},
		VAR:        {fnNothing, stateVar},
		TX:         {fnTx, stateTX},
		SETTINGS:   {fnSettings, stateSettings},
		ERROR:      {fnCmdError, stateEval},
		ERRWARNING: {fnCmdError, stateEval},
		ERRINFO:    {fnCmdError, stateEval},
		IDENTIFIER: {fnNothing, stateAssignEval | stateFlagFork},
		EXTEND:     {fnNothing, stateAssignEval | stateFlagFork},
		RBRACE:     {fnNothing, stateFlagPop},
		UNKNOWN:    {fnError, errMustRBRACE},
	}
	stateTable[stateBlock] = map[Token]compileState{
		NEWLINE: {fnNothing, stateBlock},
		LBRACE:  {fnNothing, stateBody},
		UNKNOWN: {fnError, errMustLBRACE},
	}
	stateTable[stateContract] = map[Token]compileState{
		NEWLINE:    {fnNothing, stateContract},
		IDENTIFIER: {fnBlockDecl, stateBlock},
		UNKNOWN:    {fnError, errMustName},
	}
	stateTable[stateFunc] = map[Token]compileState{
		NEWLINE:    {fnNothing, stateFunc},
		IDENTIFIER: {fnBlockDecl, stateFnParams},
		UNKNOWN:    {fnError, errMustName},
	}
	stateTable[stateFnParams] = map[Token]compileState{
		NEWLINE: {fnNothing, stateFnParams},
		LPAREN:  {fnNothing, stateFnParam},
		UNKNOWN: {fnNothing, stateFnResult | stateFlagStay},
	}
	stateTable[stateFnParam] = map[Token]compileState{
		NEWLINE:    {fnNothing, stateFnParam},
		IDENTIFIER: {fnParamName, stateFnParamType},
		COMMA:      {fnNothing, stateFnParam},
		RPAREN:     {fnNothing, stateFnResult},
		UNKNOWN:    {fnError, errParams},
	}
	stateTable[stateFnParamType] = map[Token]compileState{
		IDENTIFIER: {fnParamName, stateFnParamType},
		TYPENAME:   {fnParamType, stateFnParam},
		TAIL:       {fnTailParamType, stateFnTail},
		COMMA:      {fnNothing, stateFnParamType},
		UNKNOWN:    {fnError, errVarType},
	}
	stateTable[stateFnTail] = map[Token]compileState{
		NEWLINE: {fnNothing, stateFnTail},
		RPAREN:  {fnNothing, stateFnResult},
		UNKNOWN: {fnError, errMisplacedDotDotDot},
	}
	stateTable[stateFnResult] = map[Token]compileState{
		NEWLINE:  {fnNothing, stateFnResult},
		DOT:      {fnNothing, stateFnDot},
		TYPENAME: {fnFuncResult, stateFnResult},
		COMMA:    {fnNothing, stateFnResult},
		UNKNOWN:  {fnNothing, stateBlock | stateFlagStay},
	}
	stateTable[stateFnDot] = map[Token]compileState{
		NEWLINE:    {fnNothing, stateFnDot},
		IDENTIFIER: {fnDeclTail, stateFnParams},
		UNKNOWN:    {fnError, errMustName},
	}
	stateTable[stateVar] = map[Token]compileState{
		NEWLINE:    {fnNothing, stateBody},
		IDENTIFIER: {fnParamName, stateVarType},
		RBRACE:     {fnNothing, stateBody | stateFlagStay},
		COMMA:      {fnNothing, stateVar},
		UNKNOWN:    {fnError, errVars},
	}
	stateTable[stateVarType] = map[Token]compileState{
		IDENTIFIER: {fnParamName, stateVarType},
		TYPENAME:   {fnParamType, stateVar},
		COMMA:      {fnNothing, stateVarType},
		UNKNOWN:    {fnError, errVarType},
	}
	stateTable[stateAssignEval] = map[Token]compileState{
		LPAREN:  {fnNothing, stateEval | stateFlagToFork | stateFlagToBody},
		LBRACK:  {fnNothing, stateEval | stateFlagToFork | stateFlagToBody},
		UNKNOWN: {fnNothing, stateAssign | stateFlagToFork | stateFlagStay},
	}
	stateTable[stateAssign] = map[Token]compileState{
		COMMA:      {fnNothing, stateAssign},
		IDENTIFIER: {fnAssignVar, stateAssign},
		EXTEND:     {fnAssignVar, stateAssign},
		EQ:         {fnAssign, stateEval | stateFlagToBody},
		OPERATOR:   {fnAssign, stateEval | stateFlagToBody},
		UNKNOWN:    {fnError, errAssign},
	}
	stateTable[stateTX] = map[Token]compileState{
		NEWLINE: {fnNothing, stateTX},
		LBRACE:  {fnNothing, stateFields},
		//IDENTIFIER:   {stateAssign,fTX},
		EXTEND:  {fnTx, stateAssign},
		UNKNOWN: {fnError, errMustLBRACE},
	}
	stateTable[stateSettings] = map[Token]compileState{
		NEWLINE: {fnNothing, stateSettings},
		LBRACE:  {fnNothing, stateConsts},
		UNKNOWN: {fnError, errMustLBRACE},
	}
	stateTable[stateConsts] = map[Token]compileState{
		NEWLINE:    {fnNothing, stateConsts},
		COMMA:      {fnNothing, stateConsts},
		IDENTIFIER: {fnConstName, stateConstsAssign},
		RBRACE:     {fnNothing, stateFlagToBody},
		UNKNOWN:    {fnError, errMustRBRACE},
	}
	stateTable[stateConstsAssign] = map[Token]compileState{
		EQ:      {fnNothing, stateConstsValue},
		UNKNOWN: {fnError, errAssign},
	}
	stateTable[stateConstsValue] = map[Token]compileState{
		LITERAL: {fnConstValue, stateConsts},
		NUMBER:  {fnConstValue, stateConsts},
		UNKNOWN: {fnError, errStrNum},
	}
	stateTable[stateFields] = map[Token]compileState{
		NEWLINE:    {fnFieldLine, stateFields},
		COMMA:      {fnFieldComma, stateFields},
		IDENTIFIER: {fnField, stateFields},
		TYPENAME:   {fnFieldType, stateFields},
		LITERAL:    {fnFieldTag, stateFields},
		RBRACE:     {fnFields, stateFlagToBody},
		UNKNOWN:    {fnError, errMustRBRACE},
	}
}
