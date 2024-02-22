package compiler

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
	return fmt.Sprintf("%s", s)
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

// compileState contains a next state and a handle function.
type compileState struct {
	handle compileFunc
	next   stateType
}

// hasState returns true if the state is set
func (c compileState) hasState(state int) bool {
	return int(c.next)&state > 0
}

// stateTable describes a finite machine with states on the base of which a bytecode will be generated
var stateTable = make(map[stateType]map[Token]compileState)

func init() {
	stateTable[stateRoot] = map[Token]compileState{
		NEWLINE:  {handleNothing, stateRoot},
		CONTRACT: {handleNothing, stateContract | stateFlagPush},
		FUNC:     {handleNothing, stateFunc | stateFlagPush},
		UNKNOWN:  {handleError, errUnknownCmd},
	}
	stateTable[stateBody] = map[Token]compileState{
		NEWLINE:    {handleNothing, stateBody},
		FUNC:       {handleNothing, stateFunc | stateFlagPush},
		RETURN:     {handleReturn, stateEval},
		CONTINUE:   {handleContinue, stateBody},
		BREAK:      {handleBreak, stateBody},
		IF:         {handleIf, stateEval | stateFlagPush | stateFlagToBlock | stateFlagMustEval},
		WHILE:      {handleWhile, stateEval | stateFlagPush | stateFlagToBlock | stateFlagLabel | stateFlagMustEval},
		ELSE:       {handleElse, stateBlock | stateFlagPush},
		VAR:        {handleNothing, stateVar},
		TX:         {handleTx, stateTX},
		SETTINGS:   {handleSettings, stateSettings},
		ERROR:      {handleCmdError, stateEval},
		ERRWARNING: {handleCmdError, stateEval},
		ERRINFO:    {handleCmdError, stateEval},
		IDENTIFIER: {handleNothing, stateAssignEval | stateFlagFork},
		EXTEND:     {handleNothing, stateAssignEval | stateFlagFork},
		RBRACE:     {handleNothing, stateFlagPop},
		UNKNOWN:    {handleError, errMustRBRACE},
	}
	stateTable[stateBlock] = map[Token]compileState{
		NEWLINE: {handleNothing, stateBlock},
		LBRACE:  {handleNothing, stateBody},
		UNKNOWN: {handleError, errMustLBRACE},
	}
	stateTable[stateContract] = map[Token]compileState{
		NEWLINE:    {handleNothing, stateContract},
		IDENTIFIER: {handleBlockDecl, stateBlock},
		UNKNOWN:    {handleError, errMustName},
	}
	stateTable[stateFunc] = map[Token]compileState{
		NEWLINE:    {handleNothing, stateFunc},
		IDENTIFIER: {handleBlockDecl, stateFnParams},
		UNKNOWN:    {handleError, errMustName},
	}
	stateTable[stateFnParams] = map[Token]compileState{
		NEWLINE: {handleNothing, stateFnParams},
		LPAREN:  {handleNothing, stateFnParam},
		UNKNOWN: {handleNothing, stateFnResult | stateFlagStay},
	}
	stateTable[stateFnParam] = map[Token]compileState{
		NEWLINE:    {handleNothing, stateFnParam},
		IDENTIFIER: {handleParamName, stateFnParamType},
		COMMA:      {handleNothing, stateFnParam},
		RPAREN:     {handleNothing, stateFnResult},
		UNKNOWN:    {handleError, errParams},
	}
	stateTable[stateFnParamType] = map[Token]compileState{
		IDENTIFIER: {handleParamName, stateFnParamType},
		TYPENAME:   {handleParamType, stateFnParam},
		TAIL:       {handleTailParamType, stateFnTail},
		COMMA:      {handleNothing, stateFnParamType},
		UNKNOWN:    {handleError, errVarType},
	}
	stateTable[stateFnTail] = map[Token]compileState{
		NEWLINE: {handleNothing, stateFnTail},
		RPAREN:  {handleNothing, stateFnResult},
		UNKNOWN: {handleError, errMisplacedDotDotDot},
	}
	stateTable[stateFnResult] = map[Token]compileState{
		NEWLINE:  {handleNothing, stateFnResult},
		DOT:      {handleNothing, stateFnDot},
		TYPENAME: {handleFuncResult, stateFnResult},
		COMMA:    {handleNothing, stateFnResult},
		UNKNOWN:  {handleNothing, stateBlock | stateFlagStay},
	}
	stateTable[stateFnDot] = map[Token]compileState{
		NEWLINE:    {handleNothing, stateFnDot},
		IDENTIFIER: {handleDeclTail, stateFnParams},
		UNKNOWN:    {handleError, errMustName},
	}
	stateTable[stateVar] = map[Token]compileState{
		NEWLINE:    {handleNothing, stateBody},
		IDENTIFIER: {handleParamName, stateVarType},
		RBRACE:     {handleNothing, stateBody | stateFlagStay},
		COMMA:      {handleNothing, stateVar},
		UNKNOWN:    {handleError, errVars},
	}
	stateTable[stateVarType] = map[Token]compileState{
		IDENTIFIER: {handleParamName, stateVarType},
		TYPENAME:   {handleParamType, stateVar},
		COMMA:      {handleNothing, stateVarType},
		UNKNOWN:    {handleError, errVarType},
	}
	stateTable[stateAssignEval] = map[Token]compileState{
		LPAREN:  {handleNothing, stateEval | stateFlagToFork | stateFlagToBody},
		LBRACK:  {handleNothing, stateEval | stateFlagToFork | stateFlagToBody},
		UNKNOWN: {handleNothing, stateAssign | stateFlagToFork | stateFlagStay},
	}
	stateTable[stateAssign] = map[Token]compileState{
		COMMA:      {handleNothing, stateAssign},
		IDENTIFIER: {handleAssignVar, stateAssign},
		EXTEND:     {handleAssignVar, stateAssign},
		EQ:         {handleAssign, stateEval | stateFlagToBody},
		OPERATOR:   {handleAssign, stateEval | stateFlagToBody},
		UNKNOWN:    {handleError, errAssign},
	}
	stateTable[stateTX] = map[Token]compileState{
		NEWLINE: {handleNothing, stateTX},
		LBRACE:  {handleNothing, stateFields},
		// IDENTIFIER:   {stateAssign,fTX},
		EXTEND:  {handleTx, stateAssign},
		UNKNOWN: {handleError, errMustLBRACE},
	}
	stateTable[stateSettings] = map[Token]compileState{
		NEWLINE: {handleNothing, stateSettings},
		LBRACE:  {handleNothing, stateConsts},
		UNKNOWN: {handleError, errMustLBRACE},
	}
	stateTable[stateConsts] = map[Token]compileState{
		NEWLINE:    {handleNothing, stateConsts},
		COMMA:      {handleNothing, stateConsts},
		IDENTIFIER: {handleConstName, stateConstsAssign},
		RBRACE:     {handleNothing, stateFlagToBody},
		UNKNOWN:    {handleError, errMustRBRACE},
	}
	stateTable[stateConstsAssign] = map[Token]compileState{
		EQ:      {handleNothing, stateConstsValue},
		UNKNOWN: {handleError, errAssign},
	}
	stateTable[stateConstsValue] = map[Token]compileState{
		LITERAL: {handleConstValue, stateConsts},
		NUMBER:  {handleConstValue, stateConsts},
		UNKNOWN: {handleError, errStrNum},
	}
	stateTable[stateFields] = map[Token]compileState{
		NEWLINE:    {handleFieldLine, stateFields},
		COMMA:      {handleFieldComma, stateFields},
		IDENTIFIER: {handleField, stateFields},
		TYPENAME:   {handleFieldType, stateFields},
		LITERAL:    {handleFieldTag, stateFields},
		RBRACE:     {handleFields, stateFlagToBody},
		UNKNOWN:    {handleError, errMustRBRACE},
	}
}
