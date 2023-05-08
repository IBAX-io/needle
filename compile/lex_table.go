package compile

import (
	"fmt"
	"os"
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
		0x01, 0x0A, ' ', '`', '"', ';', '(', ')', '[', ']',
		'{', '}', '&', '|', '#', '.', ',', '<', '>', '=',
		'!', '*', '$', '@', ':', '+', '-', '/', '\\',
		'0', '1', 'a', '_',
		'~', '^', '%', '?', 0x27,
		0x80,
	}
	stateRule = map[int]map[string]action{
		stateMain: {
			"n;":        {stateMain, NEWLINE, flagNext},
			"()#[],{}:": {stateMain, SYSTEM, flagNext},
			"s":         {stateMain, UNKNOWN, flagNext},
			"q":         {stateString, UNKNOWN, flagPush | flagNext},
			"Q":         {stateDString, UNKNOWN, flagPush | flagNext},
			"&":         {stateAnd, UNKNOWN, flagPush | flagNext},
			"|":         {stateOr, UNKNOWN, flagPush | flagNext},
			"=":         {stateEq, UNKNOWN, flagPush | flagNext},
			"/":         {stateSolidus, UNKNOWN, flagPush | flagNext},
			"!":         {stateOpNeq, UNKNOWN, flagPush | flagNext},
			"<":         {stateLess, UNKNOWN, flagPush | flagNext},
			">":         {stateGreat, UNKNOWN, flagPush | flagNext},
			"*+-":       {stateMain, OPERATOR, flagNext},
			"01":        {stateNumber, UNKNOWN, flagPush | flagNext},
			"a_r":       {stateIdentifier, UNKNOWN, flagPush | flagNext},
			"@$":        {stateMustIdent, UNKNOWN, flagPush | flagNext},
			".":         {stateDot, UNKNOWN, flagPush | flagNext},
			"d":         {stateError, UNKNOWN, flagEnd},
			"%":         {statePercent, UNKNOWN, flagPush | flagNext},
		},
		statePercent: {
			"=": {stateMain, OPERATOR, flagPop | flagNext},
			"d": {stateMain, OPERATOR, flagPop},
		},
		stateString: {
			"q": {stateMain, LITERAL, flagPop | flagNext},
			"d": {stateString, UNKNOWN, flagNext},
		},
		stateDString: {
			"Q":  {stateMain, LITERAL, flagPop | flagNext},
			`\\`: {stateDoubleSlash, UNKNOWN, flagSkip},
			"d":  {stateDString, UNKNOWN, flagNext},
		},
		stateDoubleSlash: {
			"d": {stateDString, UNKNOWN, flagNext},
		},
		stateDot: {
			".":  {stateDoubleDot, UNKNOWN, flagNext},
			"01": {stateNumber, UNKNOWN, flagNext},
			"d":  {stateMain, SYSTEM, flagPop},
		},
		stateDoubleDot: {
			".": {stateMain, IDENTIFIER, flagPop | flagNext},
			"d": {stateError, UNKNOWN, flagEnd},
		},
		stateAnd: {
			"&": {stateMain, OPERATOR, flagPop | flagNext},
			"d": {stateMain, OPERATOR, flagPop},
		},
		stateOr: {
			"|": {stateMain, OPERATOR, flagPop | flagNext},
			"d": {stateError, UNKNOWN, flagEnd},
		},
		stateEq: {
			"=": {stateMain, OPERATOR, flagPop | flagNext},
			"d": {stateMain, SYSTEM, flagPop},
		},
		stateSolidus: {
			"/": {stateComLine, UNKNOWN, flagPop | flagNext},
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
			"*": {stateComStop, UNKNOWN, flagNext},
			"d": {stateComment, UNKNOWN, flagNext},
		},
		stateComStop: {
			"/": {stateMain, COMMENT, flagPop | flagNext},
			"d": {stateComment, UNKNOWN, flagNext},
		},
		stateComLine: {
			"n": {stateMain, UNKNOWN, flagEnd},
			"d": {stateComLine, UNKNOWN, flagNext},
		},
	}
)

const (
	stateError = iota
	stateMain
	stateString
	stateDString
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
	stateComStop
	stateComLine
	statePercent
	stateLess
	stateGreat
	stateShiftEq
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
	//b2 := alphabets[i]
	//fmt.Printf("%d %c [%d,%d]\n", i, i, b, b2)
	//}
	fmt.Println()
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
			var val = 0xFF << 16
			if v.state != stateError {
				val = (arr[v.state-1] - 1) << 16
			}
			val |= int(v.token << 8)
			val |= v.flag
			//fmt.Println(curState, key, v.State, v.Lexeme, v.Flag, val)
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
	return
	out := "package lex\n\nvar lexTable2 = [][" + fmt.Sprint(alphaRuleSize) + "]uint32{\n"
	for _, line := range lexTable {
		out += "\t{"
		for _, ival := range line {
			out += fmt.Sprintf("0x%x, ", ival)
		}
		out += "},\n"
	}
	out += "}\n"
	err := os.WriteFile("./lex_table.go", []byte(out), 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
}
