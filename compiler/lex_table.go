package compiler

import (
	"sort"
)

const (
	stateError = 0xff
	stateMain  = iota
	stateString
	stateDoubleString
	stateDoubleSlash
	stateDot
	stateDoubleDot
	stateAnd
	stateOr
	stateEq
	stateSolidus
	stateOpNeq
	stateNumber
	stateIdentifier
	stateMustIdent
	stateComment
	stateCommentStop
	stateCommentLine
	statePercent
	stateLess
	stateGreat
	stateShiftEq
	stateDouble
	stateAdd
	stateSub
	stateMul
	stateBitXor
	stateUnderscore
	stateExponentPart
)

// flags of lexical states
const (
	flagEnd  = 0
	flagNext = 1
	flagPush = 2
	flagPop  = 4
	flagSkip = 8
)

type action struct {
	state int
	token Token
	// flag is a set of flags that determine the behavior of the state machine
	flag int
}

var (
	// lexTable is a table of transitions between states of the state machine
	lexTable [][alphaRuleSize]int

	// stateMachine is description of the state machine that passes from one state to another
	// depending on the next received character.
	stateMachine = map[int]map[string]action{
		stateMain: {
			"n;":       {stateMain, NEWLINE, flagNext},
			"()[],{}:": {stateMain, DELIMITER, flagNext},
			"s":        {stateMain, UNKNOWN, flagNext},
			"q":        {stateString, UNKNOWN, flagPush | flagNext},
			"Q":        {stateDoubleString, UNKNOWN, flagPush | flagNext},
			"&":        {stateAnd, UNKNOWN, flagPush | flagNext},
			"|":        {stateOr, UNKNOWN, flagPush | flagNext},
			"=":        {stateEq, UNKNOWN, flagPush | flagNext},
			"/":        {stateSolidus, UNKNOWN, flagPush | flagNext},
			"!":        {stateOpNeq, UNKNOWN, flagPush | flagNext},
			"<":        {stateLess, UNKNOWN, flagPush | flagNext},
			">":        {stateGreat, UNKNOWN, flagPush | flagNext},
			"*":        {stateMul, UNKNOWN, flagPush | flagNext},
			"^":        {stateBitXor, UNKNOWN, flagPush | flagNext},
			"+":        {stateAdd, UNKNOWN, flagPush | flagNext},
			"-":        {stateSub, UNKNOWN, flagPush | flagNext},
			"01":       {stateNumber, UNKNOWN, flagPush | flagNext},
			"a_r":      {stateIdentifier, UNKNOWN, flagPush | flagNext},
			"@$":       {stateMustIdent, UNKNOWN, flagPush | flagNext},
			".":        {stateDot, UNKNOWN, flagPush | flagNext},
			"d":        {stateError, UNKNOWN, flagEnd},
			"%":        {statePercent, UNKNOWN, flagPush | flagNext},
		},
		stateBitXor: {
			"=": {stateMain, OPERATOR, flagPop | flagNext},
			"d": {stateMain, OPERATOR, flagPop},
		},
		stateMul: {
			"=": {stateMain, OPERATOR, flagPop | flagNext},
			"d": {stateMain, OPERATOR, flagPop},
		},
		stateAdd: {
			"=": {stateMain, OPERATOR, flagPop | flagNext},
			"+": {stateDouble, UNKNOWN, flagPop | flagNext},
			"d": {stateMain, OPERATOR, flagPop},
		},
		stateSub: {
			"=": {stateMain, OPERATOR, flagPop | flagNext},
			"-": {stateDouble, UNKNOWN, flagPop | flagNext},
			"d": {stateMain, OPERATOR, flagPop},
		},
		stateDouble: {
			"d": {stateMain, OPERATOR, flagPop},
		},
		statePercent: {
			"=": {stateMain, OPERATOR, flagPop | flagNext},
			"d": {stateMain, OPERATOR, flagPop},
		},
		stateString: {
			"q": {stateMain, LITERAL, flagPop | flagNext},
			"d": {stateString, UNKNOWN, flagNext},
		},
		stateDoubleString: {
			"Q":  {stateMain, LITERAL, flagPop | flagNext},
			`\\`: {stateDoubleSlash, UNKNOWN, flagSkip},
			"d":  {stateDoubleString, UNKNOWN, flagNext},
		},
		stateDoubleSlash: {
			"d": {stateDoubleString, UNKNOWN, flagNext},
		},
		stateDot: {
			".":  {stateDoubleDot, UNKNOWN, flagNext},
			"01": {stateNumber, UNKNOWN, flagNext},
			"d":  {stateMain, DELIMITER, flagPop},
		},
		stateDoubleDot: {
			".": {stateMain, IDENTIFIER, flagPop | flagNext},
			"d": {stateError, UNKNOWN, flagEnd},
		},
		stateAnd: {
			"&=": {stateMain, OPERATOR, flagPop | flagNext},
			"d":  {stateMain, OPERATOR, flagPop},
		},
		stateOr: {
			"|=": {stateMain, OPERATOR, flagPop | flagNext},
			"d":  {stateMain, OPERATOR, flagPop},
		},
		stateEq: {
			"=": {stateMain, OPERATOR, flagPop | flagNext},
			"d": {stateMain, DELIMITER, flagPop},
		},
		stateSolidus: {
			"=": {stateMain, OPERATOR, flagPop | flagNext},
			"/": {stateCommentLine, UNKNOWN, flagPop | flagNext},
			"*": {stateComment, UNKNOWN, flagNext},
			"d": {stateMain, OPERATOR, flagPop},
		},
		stateOpNeq: {
			"=": {stateMain, OPERATOR, flagPop | flagNext},
			"d": {stateMain, OPERATOR, flagPop},
		},
		stateLess: {
			"=": {stateMain, OPERATOR, flagPop | flagNext},
			"<": {stateShiftEq, UNKNOWN, flagPop | flagNext},
			"d": {stateMain, OPERATOR, flagPop},
		},
		stateGreat: {
			"=": {stateMain, OPERATOR, flagPop | flagNext},
			">": {stateShiftEq, UNKNOWN, flagPop | flagNext},
			"d": {stateMain, OPERATOR, flagPop},
		},
		stateShiftEq: {
			"=": {stateMain, OPERATOR, flagPop | flagNext},
			"d": {stateMain, OPERATOR, flagPop},
		},
		stateNumber: {
			"01": {stateNumber, UNKNOWN, flagNext},
			".":  {stateDot, UNKNOWN, flagNext},
			"a":  {stateExponentPart, UNKNOWN, flagNext},
			"_":  {stateUnderscore, UNKNOWN, flagNext},
			"d":  {stateMain, NUMBER, flagPop},
			"r":  {stateError, UNKNOWN, flagEnd},
		},
		stateExponentPart: {
			"+-": {stateExponentPart, UNKNOWN, flagNext},
			"01": {stateNumber, UNKNOWN, flagNext},
			"d":  {stateError, UNKNOWN, flagEnd},
		},
		stateUnderscore: {
			"01": {stateNumber, UNKNOWN, flagNext},
			"d":  {stateError, UNKNOWN, flagEnd},
		},
		stateIdentifier: {
			"01a_r": {stateIdentifier, UNKNOWN, flagNext},
			"d":     {stateMain, IDENTIFIER, flagPop},
		},
		stateMustIdent: {
			"01a_r": {stateIdentifier, UNKNOWN, flagNext},
			"d":     {stateError, UNKNOWN, flagEnd},
		},
		stateComment: {
			"*": {stateCommentStop, UNKNOWN, flagNext},
			"d": {stateComment, UNKNOWN, flagNext},
		},
		stateCommentStop: {
			"/": {stateMain, COMMENT, flagPop | flagNext},
			"d": {stateComment, UNKNOWN, flagNext},
		},
		stateCommentLine: {
			"n": {stateMain, UNKNOWN, flagEnd},
			"d": {stateCommentLine, UNKNOWN, flagNext},
		},
	}
)

func init() {
	buildLexTable()
}

func buildLexTable() {
	var arr []int
	for i := range stateMachine {
		arr = append(arr, i)
	}
	sort.Ints(arr)
	lexTable = make([][alphaRuleSize]int, len(stateMachine))
	for curState, r := range arr {
		action := stateMachine[r]
		for i := range lexTable[curState] {
			lexTable[curState][i] = 0xFE << 16
		}
		for key, ac := range action {
			val := ac.state << 16
			if ac.state != stateError {
				val = (arr[ac.state-1] - 1) << 16
			}
			val |= int(ac.token << 8)
			val |= ac.flag
			for _, ch := range []byte(key) {
				var ind int
				switch ch {
				case 'd':
					ind = 0
				case 'n':
					ind = 1
				case 's':
					ind = 2
				case 'q':
					ind = 3
				case 'Q':
					ind = 4
				case 'r':
					ind = alphaRuleSize - 1
				default:
					for k, ach := range alphaRule {
						if ach == ch {
							ind = k
							break
						}
					}
				}
				lexTable[curState][ind] = val
				if ind == 0 {
					for i := range lexTable[curState] {
						if lexTable[curState][i] == 0xFE<<16 {
							lexTable[curState][i] = val
						}
					}
				}
			}
		}
	}
}
