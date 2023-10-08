package compile

import (
	"sort"
)

const (
	// alphaSize is the length of alphabet
	alphaSize = 129
	// alphaRuleSize is the length of alphabet rule
	alphaRuleSize = 39
)

type action struct {
	state int
	token Token
	flag  int
}

var (
	lexTable   [][alphaRuleSize]int
	alphaTable [alphaSize]byte
	alphaRule  = [alphaRuleSize]byte{
		0x01, 0x0A, ' ', '`', '"',
		';', '(', ')', '[', ']',
		'{', '}', '&', '|', '#', '.', ',', '<', '>', '=',
		'!', '*', '$', '@', ':', '+', '-', '/', '\\',
		'0', '1', 'a', '_',
		'~', '^', '%', '?', 0x27,
		0x80,
	}
	stateRule = map[int]map[string]action{
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
			"01.": {stateNumber, UNKNOWN, flagNext},
			"a_r": {stateError, UNKNOWN, flagEnd},
			"d":   {stateMain, NUMBER, flagPop},
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
)

// flags of lexical states
const (
	flagEnd  = 0
	flagNext = 1
	flagPush = 2
	flagPop  = 4
	flagSkip = 8
)

func init() {
	buildAlphaTable()
	buildLexTable()
}

func buildAlphaTable() {
	for ind, ch := range alphaRule {
		r := byte(ind)
		switch ch {
		case ' ':
			alphaTable[0x09] = r //Horizontal Tab, HT
			alphaTable[0x0d] = r //Carriage Return, CR
			alphaTable[0x20] = r //Space
		case '1':
			for k := '1'; k <= '9'; k++ {
				alphaTable[k] = r
			}
		case 'a':
			for k := 'a'; k <= 'z'; k++ {
				alphaTable[k] = r
				alphaTable[k-32] = r
			}
		case 0x80:
			alphaTable[0x80] = r
		default:
			alphaTable[ch] = r
		}
	}
	//for i, b := range alphaTable {
	//	fmt.Printf("%d %c [%d]\n", i, i, b)
	//}
}

func charToAlpha(ch rune) byte {
	if ch > 127 {
		return alphaTable[len(alphaTable)-1]
	}
	return alphaTable[ch]
}

func buildLexTable() {
	var arr []int
	for i := range stateRule {
		arr = append(arr, i)
	}
	sort.Ints(arr)
	lexTable = make([][alphaRuleSize]int, len(stateRule))
	for curState, r := range arr {
		action := stateRule[r]
		for i := range lexTable[curState] {
			lexTable[curState][i] = 0xFE << 16
		}
		for key, v := range action {
			var val = v.state << 16
			if v.state != stateError {
				val = (arr[v.state-1] - 1) << 16
			}
			val |= int(v.token << 8)
			val |= v.flag
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
