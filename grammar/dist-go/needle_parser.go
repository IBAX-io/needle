// Code generated from NeedleParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package needle // NeedleParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type NeedleParser struct {
	*antlr.BaseParser
}

var NeedleParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func needleparserParserInit() {
	staticData := &NeedleParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'('", "')'", "','", "'.'", "':'", "'='", "'['", "']'", "'{'", "'}'",
		"';'", "'!'", "'*'", "'+'", "'-'", "'/'", "'<'", "'>'", "'!='", "'&&'",
		"'<='", "'=='", "'>='", "'||'", "'&'", "'|'", "'^'", "'%'", "'<<'",
		"'>>'", "'+='", "'-='", "'*='", "'/='", "'%='", "'<<='", "'>>='", "'&='",
		"'|='", "'^='", "'++'", "'--'", "'contract'", "'func'", "'return'",
		"'if'", "'elif'", "'else'", "'while'", "'true'", "'false'", "'var'",
		"'data'", "'settings'", "'break'", "'continue'", "'warning'", "'info'",
		"'nil'", "'action'", "'conditions'", "'...'", "'error'", "'bool'", "'bytes'",
		"'int'", "'address'", "'array'", "'map'", "'money'", "'float'", "'string'",
		"'file'", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "'\"optional\"'",
	}
	staticData.SymbolicNames = []string{
		"", "LPAREN", "RPAREN", "COMMA", "DOT", "COLON", "EQ", "LBRACK", "RBRACK",
		"LBRACE", "RBRACE", "SEMI", "NOT", "MUL", "ADD", "SUB", "QUO", "LESS",
		"GREATER", "NOT_EQ", "AND", "LESS_EQ", "EQ_EQ", "GR_EQ", "OR", "BIT_AND",
		"BIT_OR", "BIT_XOR", "MOD", "LSHIFT", "RSHIFT", "ADD_EQ", "SUB_EQ",
		"MUL_EQ", "DIV_EQ", "MOD_EQ", "LSHIFT_EQ", "RSHIFT_EQ", "BIT_AND_EQ",
		"BIT_OR_EQ", "BIT_XOR_EQ", "INC", "DEC", "CONTRACT", "FUNC", "RETURN",
		"IF", "ELIF", "ELSE", "WHILE", "TRUE", "FALSE", "VAR", "DATA", "SETTINGS",
		"BREAK", "CONTINUE", "ERRWARNING", "ERRINFO", "NIL", "ACTION", "CONDITIONS",
		"TAIL", "ERROR", "BOOL", "BYTES", "INT", "ADDRESS", "ARRAY", "MAP",
		"MONEY", "FLOAT", "STRING", "FILE", "Identifier", "DollarIdentifier",
		"AtIdentifier", "InterpretedStringLiteral", "RawStringLiteral", "DecimalLiteral",
		"FloatLiteral", "HexLiteral", "OctalLiteral", "BinaryLiteral", "RuneLiteral",
		"HexByteValue", "OctalByteValue", "BytesValue", "LittleUValue", "BigUValue",
		"TagOptional", "WS", "COMMENT", "LINE_COMMENT", "TERMINATOR", "WS_NLSEMI",
		"COMMENT_NLSEMI", "LINE_COMMENT_NLSEMI", "EOS", "OTHER",
	}
	staticData.RuleNames = []string{
		"sourceMain", "contractDef", "contractPart", "dataDef", "dataPartList",
		"settingsDef", "funcDef", "defaultFuncDef", "funcDescriptor", "funcSignature",
		"funcTail", "parameterList", "parameter", "returnParameters", "block",
		"statementList", "statement", "simpleStmt", "incDecStmt", "assignment",
		"varDef", "ifStmt", "ifBody", "elseBody", "returnStmt", "continueStmt",
		"breakStmt", "whileStmt", "errorStmt", "arrayExpr", "arrayList", "mapExpr",
		"pairList", "pair", "pairValue", "arguments", "exprList", "expr", "primaryExpr",
		"indexExpr", "sliceExpr", "operand", "literal", "typeName", "incDec_op",
		"mul_op", "unary_op", "add_op", "logical_op", "rel_op", "assign_op",
		"identifierFull", "identifierVar", "identifierList", "stringLiteral",
		"numberLiteral", "booleanLiteral", "eos",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 99, 512, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		1, 0, 1, 0, 3, 0, 119, 8, 0, 1, 0, 1, 0, 5, 0, 123, 8, 0, 10, 0, 12, 0,
		126, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 136, 8,
		1, 10, 1, 12, 1, 139, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 3, 2, 146, 8,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 153, 8, 3, 10, 3, 12, 3, 156, 9,
		3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 3, 4, 163, 8, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 5, 5, 172, 8, 5, 10, 5, 12, 5, 175, 9, 5, 1, 5, 1,
		5, 1, 6, 1, 6, 3, 6, 181, 8, 6, 1, 6, 1, 6, 1, 6, 1, 7, 3, 7, 187, 8, 7,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 3, 9, 195, 8, 9, 1, 9, 5, 9, 198, 8,
		9, 10, 9, 12, 9, 201, 9, 9, 1, 9, 3, 9, 204, 8, 9, 1, 10, 1, 10, 1, 10,
		3, 10, 209, 8, 10, 1, 11, 1, 11, 1, 11, 3, 11, 214, 8, 11, 3, 11, 216,
		8, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 3, 12, 223, 8, 12, 1, 12, 1,
		12, 1, 12, 5, 12, 228, 8, 12, 10, 12, 12, 12, 231, 9, 12, 1, 13, 1, 13,
		3, 13, 235, 8, 13, 1, 13, 5, 13, 238, 8, 13, 10, 13, 12, 13, 241, 9, 13,
		1, 14, 1, 14, 3, 14, 245, 8, 14, 1, 14, 1, 14, 1, 15, 3, 15, 250, 8, 15,
		1, 15, 3, 15, 253, 8, 15, 3, 15, 255, 8, 15, 1, 15, 1, 15, 1, 15, 4, 15,
		260, 8, 15, 11, 15, 12, 15, 261, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 3, 16, 273, 8, 16, 1, 17, 1, 17, 1, 17, 3, 17,
		278, 8, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 294, 8, 21, 10, 21, 12, 21,
		297, 9, 21, 1, 21, 3, 21, 300, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		3, 22, 307, 8, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 3,
		24, 316, 8, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 3, 29, 331, 8, 29, 1, 29, 3, 29, 334,
		8, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 3, 31, 343, 8,
		31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 5, 32, 350, 8, 32, 10, 32, 12, 32,
		353, 9, 32, 1, 32, 3, 32, 356, 8, 32, 1, 32, 1, 32, 1, 33, 1, 33, 3, 33,
		362, 8, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 3, 34, 370, 8, 34,
		1, 34, 1, 34, 1, 34, 3, 34, 375, 8, 34, 1, 35, 1, 35, 3, 35, 379, 8, 35,
		1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 5, 36, 386, 8, 36, 10, 36, 12, 36, 389,
		9, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 3,
		37, 400, 8, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 5, 37, 422, 8, 37, 10, 37, 12, 37, 425, 9, 37, 1, 38, 1, 38,
		1, 38, 1, 38, 1, 38, 1, 38, 3, 38, 433, 8, 38, 1, 38, 5, 38, 436, 8, 38,
		10, 38, 12, 38, 439, 9, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 3,
		40, 447, 8, 40, 1, 40, 1, 40, 3, 40, 451, 8, 40, 1, 40, 1, 40, 1, 41, 1,
		41, 1, 41, 1, 41, 1, 41, 1, 41, 3, 41, 461, 8, 41, 1, 42, 1, 42, 1, 42,
		1, 42, 3, 42, 467, 8, 42, 1, 43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1,
		46, 1, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51,
		1, 51, 1, 52, 1, 52, 1, 53, 1, 53, 3, 53, 491, 8, 53, 1, 53, 5, 53, 494,
		8, 53, 10, 53, 12, 53, 497, 9, 53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 56, 1,
		56, 1, 57, 1, 57, 1, 57, 3, 57, 508, 8, 57, 3, 57, 510, 8, 57, 1, 57, 1,
		295, 2, 74, 76, 58, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
		28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62,
		64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98,
		100, 102, 104, 106, 108, 110, 112, 114, 0, 15, 1, 0, 60, 61, 2, 0, 57,
		58, 63, 63, 1, 0, 64, 73, 1, 0, 41, 42, 4, 0, 13, 13, 16, 16, 25, 25, 28,
		30, 2, 0, 12, 12, 14, 15, 2, 0, 14, 15, 26, 27, 2, 0, 20, 20, 24, 24, 2,
		0, 17, 19, 21, 23, 2, 0, 6, 6, 31, 40, 1, 0, 74, 76, 1, 0, 74, 75, 1, 0,
		77, 78, 1, 0, 79, 83, 1, 0, 50, 51, 528, 0, 124, 1, 0, 0, 0, 2, 129, 1,
		0, 0, 0, 4, 145, 1, 0, 0, 0, 6, 147, 1, 0, 0, 0, 8, 159, 1, 0, 0, 0, 10,
		164, 1, 0, 0, 0, 12, 180, 1, 0, 0, 0, 14, 186, 1, 0, 0, 0, 16, 190, 1,
		0, 0, 0, 18, 194, 1, 0, 0, 0, 20, 205, 1, 0, 0, 0, 22, 210, 1, 0, 0, 0,
		24, 219, 1, 0, 0, 0, 26, 232, 1, 0, 0, 0, 28, 242, 1, 0, 0, 0, 30, 259,
		1, 0, 0, 0, 32, 272, 1, 0, 0, 0, 34, 277, 1, 0, 0, 0, 36, 279, 1, 0, 0,
		0, 38, 282, 1, 0, 0, 0, 40, 286, 1, 0, 0, 0, 42, 289, 1, 0, 0, 0, 44, 306,
		1, 0, 0, 0, 46, 310, 1, 0, 0, 0, 48, 313, 1, 0, 0, 0, 50, 317, 1, 0, 0,
		0, 52, 319, 1, 0, 0, 0, 54, 321, 1, 0, 0, 0, 56, 325, 1, 0, 0, 0, 58, 328,
		1, 0, 0, 0, 60, 337, 1, 0, 0, 0, 62, 340, 1, 0, 0, 0, 64, 346, 1, 0, 0,
		0, 66, 361, 1, 0, 0, 0, 68, 374, 1, 0, 0, 0, 70, 376, 1, 0, 0, 0, 72, 382,
		1, 0, 0, 0, 74, 399, 1, 0, 0, 0, 76, 426, 1, 0, 0, 0, 78, 440, 1, 0, 0,
		0, 80, 444, 1, 0, 0, 0, 82, 460, 1, 0, 0, 0, 84, 466, 1, 0, 0, 0, 86, 468,
		1, 0, 0, 0, 88, 470, 1, 0, 0, 0, 90, 472, 1, 0, 0, 0, 92, 474, 1, 0, 0,
		0, 94, 476, 1, 0, 0, 0, 96, 478, 1, 0, 0, 0, 98, 480, 1, 0, 0, 0, 100,
		482, 1, 0, 0, 0, 102, 484, 1, 0, 0, 0, 104, 486, 1, 0, 0, 0, 106, 488,
		1, 0, 0, 0, 108, 498, 1, 0, 0, 0, 110, 500, 1, 0, 0, 0, 112, 502, 1, 0,
		0, 0, 114, 509, 1, 0, 0, 0, 116, 119, 3, 2, 1, 0, 117, 119, 3, 12, 6, 0,
		118, 116, 1, 0, 0, 0, 118, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120,
		121, 3, 114, 57, 0, 121, 123, 1, 0, 0, 0, 122, 118, 1, 0, 0, 0, 123, 126,
		1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 127, 1, 0,
		0, 0, 126, 124, 1, 0, 0, 0, 127, 128, 5, 0, 0, 1, 128, 1, 1, 0, 0, 0, 129,
		130, 5, 43, 0, 0, 130, 131, 5, 74, 0, 0, 131, 137, 5, 9, 0, 0, 132, 133,
		3, 4, 2, 0, 133, 134, 3, 114, 57, 0, 134, 136, 1, 0, 0, 0, 135, 132, 1,
		0, 0, 0, 136, 139, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0,
		0, 138, 140, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 140, 141, 5, 10, 0, 0, 141,
		3, 1, 0, 0, 0, 142, 146, 3, 6, 3, 0, 143, 146, 3, 10, 5, 0, 144, 146, 3,
		12, 6, 0, 145, 142, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 145, 144, 1, 0, 0,
		0, 146, 5, 1, 0, 0, 0, 147, 148, 5, 53, 0, 0, 148, 154, 5, 9, 0, 0, 149,
		150, 3, 8, 4, 0, 150, 151, 3, 114, 57, 0, 151, 153, 1, 0, 0, 0, 152, 149,
		1, 0, 0, 0, 153, 156, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154, 155, 1, 0,
		0, 0, 155, 157, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 157, 158, 5, 10, 0, 0,
		158, 7, 1, 0, 0, 0, 159, 160, 5, 74, 0, 0, 160, 162, 3, 86, 43, 0, 161,
		163, 3, 108, 54, 0, 162, 161, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 9,
		1, 0, 0, 0, 164, 165, 5, 54, 0, 0, 165, 173, 5, 9, 0, 0, 166, 167, 5, 74,
		0, 0, 167, 168, 5, 6, 0, 0, 168, 169, 3, 84, 42, 0, 169, 170, 3, 114, 57,
		0, 170, 172, 1, 0, 0, 0, 171, 166, 1, 0, 0, 0, 172, 175, 1, 0, 0, 0, 173,
		171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 176, 1, 0, 0, 0, 175, 173,
		1, 0, 0, 0, 176, 177, 5, 10, 0, 0, 177, 11, 1, 0, 0, 0, 178, 181, 3, 16,
		8, 0, 179, 181, 3, 14, 7, 0, 180, 178, 1, 0, 0, 0, 180, 179, 1, 0, 0, 0,
		181, 182, 1, 0, 0, 0, 182, 183, 3, 18, 9, 0, 183, 184, 3, 28, 14, 0, 184,
		13, 1, 0, 0, 0, 185, 187, 5, 44, 0, 0, 186, 185, 1, 0, 0, 0, 186, 187,
		1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 189, 7, 0, 0, 0, 189, 15, 1, 0,
		0, 0, 190, 191, 5, 44, 0, 0, 191, 192, 5, 74, 0, 0, 192, 17, 1, 0, 0, 0,
		193, 195, 3, 22, 11, 0, 194, 193, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195,
		199, 1, 0, 0, 0, 196, 198, 3, 20, 10, 0, 197, 196, 1, 0, 0, 0, 198, 201,
		1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 203, 1, 0,
		0, 0, 201, 199, 1, 0, 0, 0, 202, 204, 3, 26, 13, 0, 203, 202, 1, 0, 0,
		0, 203, 204, 1, 0, 0, 0, 204, 19, 1, 0, 0, 0, 205, 206, 5, 4, 0, 0, 206,
		208, 5, 74, 0, 0, 207, 209, 3, 22, 11, 0, 208, 207, 1, 0, 0, 0, 208, 209,
		1, 0, 0, 0, 209, 21, 1, 0, 0, 0, 210, 215, 5, 1, 0, 0, 211, 213, 3, 24,
		12, 0, 212, 214, 5, 3, 0, 0, 213, 212, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0,
		214, 216, 1, 0, 0, 0, 215, 211, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216,
		217, 1, 0, 0, 0, 217, 218, 5, 2, 0, 0, 218, 23, 1, 0, 0, 0, 219, 220, 3,
		106, 53, 0, 220, 229, 3, 86, 43, 0, 221, 223, 5, 3, 0, 0, 222, 221, 1,
		0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 225, 3, 106,
		53, 0, 225, 226, 3, 86, 43, 0, 226, 228, 1, 0, 0, 0, 227, 222, 1, 0, 0,
		0, 228, 231, 1, 0, 0, 0, 229, 227, 1, 0, 0, 0, 229, 230, 1, 0, 0, 0, 230,
		25, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 232, 239, 3, 86, 43, 0, 233, 235,
		5, 3, 0, 0, 234, 233, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 236, 1, 0,
		0, 0, 236, 238, 3, 86, 43, 0, 237, 234, 1, 0, 0, 0, 238, 241, 1, 0, 0,
		0, 239, 237, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 27, 1, 0, 0, 0, 241,
		239, 1, 0, 0, 0, 242, 244, 5, 9, 0, 0, 243, 245, 3, 30, 15, 0, 244, 243,
		1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 247, 5, 10,
		0, 0, 247, 29, 1, 0, 0, 0, 248, 250, 5, 11, 0, 0, 249, 248, 1, 0, 0, 0,
		249, 250, 1, 0, 0, 0, 250, 255, 1, 0, 0, 0, 251, 253, 5, 98, 0, 0, 252,
		251, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 255, 1, 0, 0, 0, 254, 249,
		1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256, 257, 3, 32,
		16, 0, 257, 258, 3, 114, 57, 0, 258, 260, 1, 0, 0, 0, 259, 254, 1, 0, 0,
		0, 260, 261, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262,
		31, 1, 0, 0, 0, 263, 273, 3, 28, 14, 0, 264, 273, 3, 34, 17, 0, 265, 273,
		3, 40, 20, 0, 266, 273, 3, 42, 21, 0, 267, 273, 3, 54, 27, 0, 268, 273,
		3, 50, 25, 0, 269, 273, 3, 52, 26, 0, 270, 273, 3, 48, 24, 0, 271, 273,
		3, 56, 28, 0, 272, 263, 1, 0, 0, 0, 272, 264, 1, 0, 0, 0, 272, 265, 1,
		0, 0, 0, 272, 266, 1, 0, 0, 0, 272, 267, 1, 0, 0, 0, 272, 268, 1, 0, 0,
		0, 272, 269, 1, 0, 0, 0, 272, 270, 1, 0, 0, 0, 272, 271, 1, 0, 0, 0, 273,
		33, 1, 0, 0, 0, 274, 278, 3, 74, 37, 0, 275, 278, 3, 38, 19, 0, 276, 278,
		3, 36, 18, 0, 277, 274, 1, 0, 0, 0, 277, 275, 1, 0, 0, 0, 277, 276, 1,
		0, 0, 0, 278, 35, 1, 0, 0, 0, 279, 280, 3, 74, 37, 0, 280, 281, 3, 88,
		44, 0, 281, 37, 1, 0, 0, 0, 282, 283, 3, 72, 36, 0, 283, 284, 3, 100, 50,
		0, 284, 285, 3, 72, 36, 0, 285, 39, 1, 0, 0, 0, 286, 287, 5, 52, 0, 0,
		287, 288, 3, 24, 12, 0, 288, 41, 1, 0, 0, 0, 289, 290, 5, 46, 0, 0, 290,
		295, 3, 44, 22, 0, 291, 292, 5, 47, 0, 0, 292, 294, 3, 44, 22, 0, 293,
		291, 1, 0, 0, 0, 294, 297, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 295, 293,
		1, 0, 0, 0, 296, 299, 1, 0, 0, 0, 297, 295, 1, 0, 0, 0, 298, 300, 3, 46,
		23, 0, 299, 298, 1, 0, 0, 0, 299, 300, 1, 0, 0, 0, 300, 43, 1, 0, 0, 0,
		301, 302, 5, 1, 0, 0, 302, 303, 3, 74, 37, 0, 303, 304, 5, 2, 0, 0, 304,
		307, 1, 0, 0, 0, 305, 307, 3, 74, 37, 0, 306, 301, 1, 0, 0, 0, 306, 305,
		1, 0, 0, 0, 307, 308, 1, 0, 0, 0, 308, 309, 3, 28, 14, 0, 309, 45, 1, 0,
		0, 0, 310, 311, 5, 48, 0, 0, 311, 312, 3, 28, 14, 0, 312, 47, 1, 0, 0,
		0, 313, 315, 5, 45, 0, 0, 314, 316, 3, 74, 37, 0, 315, 314, 1, 0, 0, 0,
		315, 316, 1, 0, 0, 0, 316, 49, 1, 0, 0, 0, 317, 318, 5, 56, 0, 0, 318,
		51, 1, 0, 0, 0, 319, 320, 5, 55, 0, 0, 320, 53, 1, 0, 0, 0, 321, 322, 5,
		49, 0, 0, 322, 323, 3, 74, 37, 0, 323, 324, 3, 28, 14, 0, 324, 55, 1, 0,
		0, 0, 325, 326, 7, 1, 0, 0, 326, 327, 3, 74, 37, 0, 327, 57, 1, 0, 0, 0,
		328, 330, 5, 7, 0, 0, 329, 331, 3, 60, 30, 0, 330, 329, 1, 0, 0, 0, 330,
		331, 1, 0, 0, 0, 331, 333, 1, 0, 0, 0, 332, 334, 5, 3, 0, 0, 333, 332,
		1, 0, 0, 0, 333, 334, 1, 0, 0, 0, 334, 335, 1, 0, 0, 0, 335, 336, 5, 8,
		0, 0, 336, 59, 1, 0, 0, 0, 337, 338, 3, 72, 36, 0, 338, 339, 3, 114, 57,
		0, 339, 61, 1, 0, 0, 0, 340, 342, 5, 9, 0, 0, 341, 343, 3, 64, 32, 0, 342,
		341, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 344, 1, 0, 0, 0, 344, 345,
		5, 10, 0, 0, 345, 63, 1, 0, 0, 0, 346, 351, 3, 66, 33, 0, 347, 348, 5,
		3, 0, 0, 348, 350, 3, 66, 33, 0, 349, 347, 1, 0, 0, 0, 350, 353, 1, 0,
		0, 0, 351, 349, 1, 0, 0, 0, 351, 352, 1, 0, 0, 0, 352, 355, 1, 0, 0, 0,
		353, 351, 1, 0, 0, 0, 354, 356, 5, 3, 0, 0, 355, 354, 1, 0, 0, 0, 355,
		356, 1, 0, 0, 0, 356, 357, 1, 0, 0, 0, 357, 358, 3, 114, 57, 0, 358, 65,
		1, 0, 0, 0, 359, 362, 3, 108, 54, 0, 360, 362, 3, 104, 52, 0, 361, 359,
		1, 0, 0, 0, 361, 360, 1, 0, 0, 0, 362, 363, 1, 0, 0, 0, 363, 364, 5, 5,
		0, 0, 364, 365, 3, 68, 34, 0, 365, 67, 1, 0, 0, 0, 366, 369, 3, 104, 52,
		0, 367, 370, 3, 78, 39, 0, 368, 370, 3, 80, 40, 0, 369, 367, 1, 0, 0, 0,
		369, 368, 1, 0, 0, 0, 369, 370, 1, 0, 0, 0, 370, 375, 1, 0, 0, 0, 371,
		375, 3, 84, 42, 0, 372, 375, 3, 62, 31, 0, 373, 375, 3, 58, 29, 0, 374,
		366, 1, 0, 0, 0, 374, 371, 1, 0, 0, 0, 374, 372, 1, 0, 0, 0, 374, 373,
		1, 0, 0, 0, 375, 69, 1, 0, 0, 0, 376, 378, 5, 1, 0, 0, 377, 379, 3, 72,
		36, 0, 378, 377, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 380, 1, 0, 0, 0,
		380, 381, 5, 2, 0, 0, 381, 71, 1, 0, 0, 0, 382, 387, 3, 74, 37, 0, 383,
		384, 5, 3, 0, 0, 384, 386, 3, 74, 37, 0, 385, 383, 1, 0, 0, 0, 386, 389,
		1, 0, 0, 0, 387, 385, 1, 0, 0, 0, 387, 388, 1, 0, 0, 0, 388, 73, 1, 0,
		0, 0, 389, 387, 1, 0, 0, 0, 390, 391, 6, 37, -1, 0, 391, 392, 3, 76, 38,
		0, 392, 393, 3, 114, 57, 0, 393, 400, 1, 0, 0, 0, 394, 400, 3, 62, 31,
		0, 395, 400, 3, 58, 29, 0, 396, 397, 3, 92, 46, 0, 397, 398, 3, 74, 37,
		5, 398, 400, 1, 0, 0, 0, 399, 390, 1, 0, 0, 0, 399, 394, 1, 0, 0, 0, 399,
		395, 1, 0, 0, 0, 399, 396, 1, 0, 0, 0, 400, 423, 1, 0, 0, 0, 401, 402,
		10, 4, 0, 0, 402, 403, 3, 90, 45, 0, 403, 404, 3, 74, 37, 5, 404, 422,
		1, 0, 0, 0, 405, 406, 10, 3, 0, 0, 406, 407, 3, 98, 49, 0, 407, 408, 3,
		74, 37, 4, 408, 422, 1, 0, 0, 0, 409, 410, 10, 2, 0, 0, 410, 411, 3, 96,
		48, 0, 411, 412, 3, 74, 37, 3, 412, 422, 1, 0, 0, 0, 413, 414, 10, 1, 0,
		0, 414, 415, 3, 94, 47, 0, 415, 416, 3, 74, 37, 2, 416, 422, 1, 0, 0, 0,
		417, 418, 10, 9, 0, 0, 418, 422, 3, 78, 39, 0, 419, 420, 10, 8, 0, 0, 420,
		422, 3, 80, 40, 0, 421, 401, 1, 0, 0, 0, 421, 405, 1, 0, 0, 0, 421, 409,
		1, 0, 0, 0, 421, 413, 1, 0, 0, 0, 421, 417, 1, 0, 0, 0, 421, 419, 1, 0,
		0, 0, 422, 425, 1, 0, 0, 0, 423, 421, 1, 0, 0, 0, 423, 424, 1, 0, 0, 0,
		424, 75, 1, 0, 0, 0, 425, 423, 1, 0, 0, 0, 426, 427, 6, 38, -1, 0, 427,
		428, 3, 82, 41, 0, 428, 437, 1, 0, 0, 0, 429, 432, 10, 1, 0, 0, 430, 431,
		5, 4, 0, 0, 431, 433, 5, 74, 0, 0, 432, 430, 1, 0, 0, 0, 432, 433, 1, 0,
		0, 0, 433, 434, 1, 0, 0, 0, 434, 436, 3, 70, 35, 0, 435, 429, 1, 0, 0,
		0, 436, 439, 1, 0, 0, 0, 437, 435, 1, 0, 0, 0, 437, 438, 1, 0, 0, 0, 438,
		77, 1, 0, 0, 0, 439, 437, 1, 0, 0, 0, 440, 441, 5, 7, 0, 0, 441, 442, 3,
		74, 37, 0, 442, 443, 5, 8, 0, 0, 443, 79, 1, 0, 0, 0, 444, 446, 5, 7, 0,
		0, 445, 447, 3, 74, 37, 0, 446, 445, 1, 0, 0, 0, 446, 447, 1, 0, 0, 0,
		447, 448, 1, 0, 0, 0, 448, 450, 5, 5, 0, 0, 449, 451, 3, 74, 37, 0, 450,
		449, 1, 0, 0, 0, 450, 451, 1, 0, 0, 0, 451, 452, 1, 0, 0, 0, 452, 453,
		5, 8, 0, 0, 453, 81, 1, 0, 0, 0, 454, 461, 3, 102, 51, 0, 455, 461, 3,
		84, 42, 0, 456, 457, 5, 1, 0, 0, 457, 458, 3, 74, 37, 0, 458, 459, 5, 2,
		0, 0, 459, 461, 1, 0, 0, 0, 460, 454, 1, 0, 0, 0, 460, 455, 1, 0, 0, 0,
		460, 456, 1, 0, 0, 0, 461, 83, 1, 0, 0, 0, 462, 467, 3, 110, 55, 0, 463,
		467, 3, 108, 54, 0, 464, 467, 3, 112, 56, 0, 465, 467, 5, 59, 0, 0, 466,
		462, 1, 0, 0, 0, 466, 463, 1, 0, 0, 0, 466, 464, 1, 0, 0, 0, 466, 465,
		1, 0, 0, 0, 467, 85, 1, 0, 0, 0, 468, 469, 7, 2, 0, 0, 469, 87, 1, 0, 0,
		0, 470, 471, 7, 3, 0, 0, 471, 89, 1, 0, 0, 0, 472, 473, 7, 4, 0, 0, 473,
		91, 1, 0, 0, 0, 474, 475, 7, 5, 0, 0, 475, 93, 1, 0, 0, 0, 476, 477, 7,
		6, 0, 0, 477, 95, 1, 0, 0, 0, 478, 479, 7, 7, 0, 0, 479, 97, 1, 0, 0, 0,
		480, 481, 7, 8, 0, 0, 481, 99, 1, 0, 0, 0, 482, 483, 7, 9, 0, 0, 483, 101,
		1, 0, 0, 0, 484, 485, 7, 10, 0, 0, 485, 103, 1, 0, 0, 0, 486, 487, 7, 11,
		0, 0, 487, 105, 1, 0, 0, 0, 488, 495, 5, 74, 0, 0, 489, 491, 5, 3, 0, 0,
		490, 489, 1, 0, 0, 0, 490, 491, 1, 0, 0, 0, 491, 492, 1, 0, 0, 0, 492,
		494, 5, 74, 0, 0, 493, 490, 1, 0, 0, 0, 494, 497, 1, 0, 0, 0, 495, 493,
		1, 0, 0, 0, 495, 496, 1, 0, 0, 0, 496, 107, 1, 0, 0, 0, 497, 495, 1, 0,
		0, 0, 498, 499, 7, 12, 0, 0, 499, 109, 1, 0, 0, 0, 500, 501, 7, 13, 0,
		0, 501, 111, 1, 0, 0, 0, 502, 503, 7, 14, 0, 0, 503, 113, 1, 0, 0, 0, 504,
		510, 5, 11, 0, 0, 505, 510, 5, 98, 0, 0, 506, 508, 5, 0, 0, 1, 507, 506,
		1, 0, 0, 0, 507, 508, 1, 0, 0, 0, 508, 510, 1, 0, 0, 0, 509, 504, 1, 0,
		0, 0, 509, 505, 1, 0, 0, 0, 509, 507, 1, 0, 0, 0, 510, 115, 1, 0, 0, 0,
		53, 118, 124, 137, 145, 154, 162, 173, 180, 186, 194, 199, 203, 208, 213,
		215, 222, 229, 234, 239, 244, 249, 252, 254, 261, 272, 277, 295, 299, 306,
		315, 330, 333, 342, 351, 355, 361, 369, 374, 378, 387, 399, 421, 423, 432,
		437, 446, 450, 460, 466, 490, 495, 507, 509,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// NeedleParserInit initializes any static state used to implement NeedleParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewNeedleParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func NeedleParserInit() {
	staticData := &NeedleParserParserStaticData
	staticData.once.Do(needleparserParserInit)
}

// NewNeedleParser produces a new parser instance for the optional input antlr.TokenStream.
func NewNeedleParser(input antlr.TokenStream) *NeedleParser {
	NeedleParserInit()
	this := new(NeedleParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &NeedleParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "NeedleParser.g4"

	return this
}

// NeedleParser tokens.
const (
	NeedleParserEOF                      = antlr.TokenEOF
	NeedleParserLPAREN                   = 1
	NeedleParserRPAREN                   = 2
	NeedleParserCOMMA                    = 3
	NeedleParserDOT                      = 4
	NeedleParserCOLON                    = 5
	NeedleParserEQ                       = 6
	NeedleParserLBRACK                   = 7
	NeedleParserRBRACK                   = 8
	NeedleParserLBRACE                   = 9
	NeedleParserRBRACE                   = 10
	NeedleParserSEMI                     = 11
	NeedleParserNOT                      = 12
	NeedleParserMUL                      = 13
	NeedleParserADD                      = 14
	NeedleParserSUB                      = 15
	NeedleParserQUO                      = 16
	NeedleParserLESS                     = 17
	NeedleParserGREATER                  = 18
	NeedleParserNOT_EQ                   = 19
	NeedleParserAND                      = 20
	NeedleParserLESS_EQ                  = 21
	NeedleParserEQ_EQ                    = 22
	NeedleParserGR_EQ                    = 23
	NeedleParserOR                       = 24
	NeedleParserBIT_AND                  = 25
	NeedleParserBIT_OR                   = 26
	NeedleParserBIT_XOR                  = 27
	NeedleParserMOD                      = 28
	NeedleParserLSHIFT                   = 29
	NeedleParserRSHIFT                   = 30
	NeedleParserADD_EQ                   = 31
	NeedleParserSUB_EQ                   = 32
	NeedleParserMUL_EQ                   = 33
	NeedleParserDIV_EQ                   = 34
	NeedleParserMOD_EQ                   = 35
	NeedleParserLSHIFT_EQ                = 36
	NeedleParserRSHIFT_EQ                = 37
	NeedleParserBIT_AND_EQ               = 38
	NeedleParserBIT_OR_EQ                = 39
	NeedleParserBIT_XOR_EQ               = 40
	NeedleParserINC                      = 41
	NeedleParserDEC                      = 42
	NeedleParserCONTRACT                 = 43
	NeedleParserFUNC                     = 44
	NeedleParserRETURN                   = 45
	NeedleParserIF                       = 46
	NeedleParserELIF                     = 47
	NeedleParserELSE                     = 48
	NeedleParserWHILE                    = 49
	NeedleParserTRUE                     = 50
	NeedleParserFALSE                    = 51
	NeedleParserVAR                      = 52
	NeedleParserDATA                     = 53
	NeedleParserSETTINGS                 = 54
	NeedleParserBREAK                    = 55
	NeedleParserCONTINUE                 = 56
	NeedleParserERRWARNING               = 57
	NeedleParserERRINFO                  = 58
	NeedleParserNIL                      = 59
	NeedleParserACTION                   = 60
	NeedleParserCONDITIONS               = 61
	NeedleParserTAIL                     = 62
	NeedleParserERROR                    = 63
	NeedleParserBOOL                     = 64
	NeedleParserBYTES                    = 65
	NeedleParserINT                      = 66
	NeedleParserADDRESS                  = 67
	NeedleParserARRAY                    = 68
	NeedleParserMAP                      = 69
	NeedleParserMONEY                    = 70
	NeedleParserFLOAT                    = 71
	NeedleParserSTRING                   = 72
	NeedleParserFILE                     = 73
	NeedleParserIdentifier               = 74
	NeedleParserDollarIdentifier         = 75
	NeedleParserAtIdentifier             = 76
	NeedleParserInterpretedStringLiteral = 77
	NeedleParserRawStringLiteral         = 78
	NeedleParserDecimalLiteral           = 79
	NeedleParserFloatLiteral             = 80
	NeedleParserHexLiteral               = 81
	NeedleParserOctalLiteral             = 82
	NeedleParserBinaryLiteral            = 83
	NeedleParserRuneLiteral              = 84
	NeedleParserHexByteValue             = 85
	NeedleParserOctalByteValue           = 86
	NeedleParserBytesValue               = 87
	NeedleParserLittleUValue             = 88
	NeedleParserBigUValue                = 89
	NeedleParserTagOptional              = 90
	NeedleParserWS                       = 91
	NeedleParserCOMMENT                  = 92
	NeedleParserLINE_COMMENT             = 93
	NeedleParserTERMINATOR               = 94
	NeedleParserWS_NLSEMI                = 95
	NeedleParserCOMMENT_NLSEMI           = 96
	NeedleParserLINE_COMMENT_NLSEMI      = 97
	NeedleParserEOS                      = 98
	NeedleParserOTHER                    = 99
)

// NeedleParser rules.
const (
	NeedleParserRULE_sourceMain       = 0
	NeedleParserRULE_contractDef      = 1
	NeedleParserRULE_contractPart     = 2
	NeedleParserRULE_dataDef          = 3
	NeedleParserRULE_dataPartList     = 4
	NeedleParserRULE_settingsDef      = 5
	NeedleParserRULE_funcDef          = 6
	NeedleParserRULE_defaultFuncDef   = 7
	NeedleParserRULE_funcDescriptor   = 8
	NeedleParserRULE_funcSignature    = 9
	NeedleParserRULE_funcTail         = 10
	NeedleParserRULE_parameterList    = 11
	NeedleParserRULE_parameter        = 12
	NeedleParserRULE_returnParameters = 13
	NeedleParserRULE_block            = 14
	NeedleParserRULE_statementList    = 15
	NeedleParserRULE_statement        = 16
	NeedleParserRULE_simpleStmt       = 17
	NeedleParserRULE_incDecStmt       = 18
	NeedleParserRULE_assignment       = 19
	NeedleParserRULE_varDef           = 20
	NeedleParserRULE_ifStmt           = 21
	NeedleParserRULE_ifBody           = 22
	NeedleParserRULE_elseBody         = 23
	NeedleParserRULE_returnStmt       = 24
	NeedleParserRULE_continueStmt     = 25
	NeedleParserRULE_breakStmt        = 26
	NeedleParserRULE_whileStmt        = 27
	NeedleParserRULE_errorStmt        = 28
	NeedleParserRULE_arrayExpr        = 29
	NeedleParserRULE_arrayList        = 30
	NeedleParserRULE_mapExpr          = 31
	NeedleParserRULE_pairList         = 32
	NeedleParserRULE_pair             = 33
	NeedleParserRULE_pairValue        = 34
	NeedleParserRULE_arguments        = 35
	NeedleParserRULE_exprList         = 36
	NeedleParserRULE_expr             = 37
	NeedleParserRULE_primaryExpr      = 38
	NeedleParserRULE_indexExpr        = 39
	NeedleParserRULE_sliceExpr        = 40
	NeedleParserRULE_operand          = 41
	NeedleParserRULE_literal          = 42
	NeedleParserRULE_typeName         = 43
	NeedleParserRULE_incDec_op        = 44
	NeedleParserRULE_mul_op           = 45
	NeedleParserRULE_unary_op         = 46
	NeedleParserRULE_add_op           = 47
	NeedleParserRULE_logical_op       = 48
	NeedleParserRULE_rel_op           = 49
	NeedleParserRULE_assign_op        = 50
	NeedleParserRULE_identifierFull   = 51
	NeedleParserRULE_identifierVar    = 52
	NeedleParserRULE_identifierList   = 53
	NeedleParserRULE_stringLiteral    = 54
	NeedleParserRULE_numberLiteral    = 55
	NeedleParserRULE_booleanLiteral   = 56
	NeedleParserRULE_eos              = 57
)

// ISourceMainContext is an interface to support dynamic dispatch.
type ISourceMainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllEos() []IEosContext
	Eos(i int) IEosContext
	AllContractDef() []IContractDefContext
	ContractDef(i int) IContractDefContext
	AllFuncDef() []IFuncDefContext
	FuncDef(i int) IFuncDefContext

	// IsSourceMainContext differentiates from other interfaces.
	IsSourceMainContext()
}

type SourceMainContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceMainContext() *SourceMainContext {
	var p = new(SourceMainContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_sourceMain
	return p
}

func InitEmptySourceMainContext(p *SourceMainContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_sourceMain
}

func (*SourceMainContext) IsSourceMainContext() {}

func NewSourceMainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceMainContext {
	var p = new(SourceMainContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_sourceMain

	return p
}

func (s *SourceMainContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceMainContext) EOF() antlr.TerminalNode {
	return s.GetToken(NeedleParserEOF, 0)
}

func (s *SourceMainContext) AllEos() []IEosContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEosContext); ok {
			len++
		}
	}

	tst := make([]IEosContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEosContext); ok {
			tst[i] = t.(IEosContext)
			i++
		}
	}

	return tst
}

func (s *SourceMainContext) Eos(i int) IEosContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *SourceMainContext) AllContractDef() []IContractDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IContractDefContext); ok {
			len++
		}
	}

	tst := make([]IContractDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IContractDefContext); ok {
			tst[i] = t.(IContractDefContext)
			i++
		}
	}

	return tst
}

func (s *SourceMainContext) ContractDef(i int) IContractDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContractDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContractDefContext)
}

func (s *SourceMainContext) AllFuncDef() []IFuncDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncDefContext); ok {
			len++
		}
	}

	tst := make([]IFuncDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncDefContext); ok {
			tst[i] = t.(IFuncDefContext)
			i++
		}
	}

	return tst
}

func (s *SourceMainContext) FuncDef(i int) IFuncDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDefContext)
}

func (s *SourceMainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceMainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceMainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterSourceMain(s)
	}
}

func (s *SourceMainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitSourceMain(s)
	}
}

func (p *NeedleParser) SourceMain() (localctx ISourceMainContext) {
	localctx = NewSourceMainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NeedleParserRULE_sourceMain)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3458790902099607552) != 0 {
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case NeedleParserCONTRACT:
			{
				p.SetState(116)
				p.ContractDef()
			}

		case NeedleParserFUNC, NeedleParserACTION, NeedleParserCONDITIONS:
			{
				p.SetState(117)
				p.FuncDef()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(120)
			p.Eos()
		}

		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(127)
		p.Match(NeedleParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContractDefContext is an interface to support dynamic dispatch.
type IContractDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTRACT() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllContractPart() []IContractPartContext
	ContractPart(i int) IContractPartContext
	AllEos() []IEosContext
	Eos(i int) IEosContext

	// IsContractDefContext differentiates from other interfaces.
	IsContractDefContext()
}

type ContractDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContractDefContext() *ContractDefContext {
	var p = new(ContractDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_contractDef
	return p
}

func InitEmptyContractDefContext(p *ContractDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_contractDef
}

func (*ContractDefContext) IsContractDefContext() {}

func NewContractDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContractDefContext {
	var p = new(ContractDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_contractDef

	return p
}

func (s *ContractDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ContractDefContext) CONTRACT() antlr.TerminalNode {
	return s.GetToken(NeedleParserCONTRACT, 0)
}

func (s *ContractDefContext) Identifier() antlr.TerminalNode {
	return s.GetToken(NeedleParserIdentifier, 0)
}

func (s *ContractDefContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(NeedleParserLBRACE, 0)
}

func (s *ContractDefContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(NeedleParserRBRACE, 0)
}

func (s *ContractDefContext) AllContractPart() []IContractPartContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IContractPartContext); ok {
			len++
		}
	}

	tst := make([]IContractPartContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IContractPartContext); ok {
			tst[i] = t.(IContractPartContext)
			i++
		}
	}

	return tst
}

func (s *ContractDefContext) ContractPart(i int) IContractPartContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContractPartContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContractPartContext)
}

func (s *ContractDefContext) AllEos() []IEosContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEosContext); ok {
			len++
		}
	}

	tst := make([]IEosContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEosContext); ok {
			tst[i] = t.(IEosContext)
			i++
		}
	}

	return tst
}

func (s *ContractDefContext) Eos(i int) IEosContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ContractDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContractDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContractDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterContractDef(s)
	}
}

func (s *ContractDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitContractDef(s)
	}
}

func (p *NeedleParser) ContractDef() (localctx IContractDefContext) {
	localctx = NewContractDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NeedleParserRULE_contractDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.Match(NeedleParserCONTRACT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(130)
		p.Match(NeedleParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
		p.Match(NeedleParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3485803703770808320) != 0 {
		{
			p.SetState(132)
			p.ContractPart()
		}
		{
			p.SetState(133)
			p.Eos()
		}

		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(140)
		p.Match(NeedleParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContractPartContext is an interface to support dynamic dispatch.
type IContractPartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DataDef() IDataDefContext
	SettingsDef() ISettingsDefContext
	FuncDef() IFuncDefContext

	// IsContractPartContext differentiates from other interfaces.
	IsContractPartContext()
}

type ContractPartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContractPartContext() *ContractPartContext {
	var p = new(ContractPartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_contractPart
	return p
}

func InitEmptyContractPartContext(p *ContractPartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_contractPart
}

func (*ContractPartContext) IsContractPartContext() {}

func NewContractPartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContractPartContext {
	var p = new(ContractPartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_contractPart

	return p
}

func (s *ContractPartContext) GetParser() antlr.Parser { return s.parser }

func (s *ContractPartContext) DataDef() IDataDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataDefContext)
}

func (s *ContractPartContext) SettingsDef() ISettingsDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISettingsDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISettingsDefContext)
}

func (s *ContractPartContext) FuncDef() IFuncDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDefContext)
}

func (s *ContractPartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContractPartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContractPartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterContractPart(s)
	}
}

func (s *ContractPartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitContractPart(s)
	}
}

func (p *NeedleParser) ContractPart() (localctx IContractPartContext) {
	localctx = NewContractPartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NeedleParserRULE_contractPart)
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NeedleParserDATA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(142)
			p.DataDef()
		}

	case NeedleParserSETTINGS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(143)
			p.SettingsDef()
		}

	case NeedleParserFUNC, NeedleParserACTION, NeedleParserCONDITIONS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(144)
			p.FuncDef()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDataDefContext is an interface to support dynamic dispatch.
type IDataDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DATA() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllDataPartList() []IDataPartListContext
	DataPartList(i int) IDataPartListContext
	AllEos() []IEosContext
	Eos(i int) IEosContext

	// IsDataDefContext differentiates from other interfaces.
	IsDataDefContext()
}

type DataDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataDefContext() *DataDefContext {
	var p = new(DataDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_dataDef
	return p
}

func InitEmptyDataDefContext(p *DataDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_dataDef
}

func (*DataDefContext) IsDataDefContext() {}

func NewDataDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataDefContext {
	var p = new(DataDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_dataDef

	return p
}

func (s *DataDefContext) GetParser() antlr.Parser { return s.parser }

func (s *DataDefContext) DATA() antlr.TerminalNode {
	return s.GetToken(NeedleParserDATA, 0)
}

func (s *DataDefContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(NeedleParserLBRACE, 0)
}

func (s *DataDefContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(NeedleParserRBRACE, 0)
}

func (s *DataDefContext) AllDataPartList() []IDataPartListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDataPartListContext); ok {
			len++
		}
	}

	tst := make([]IDataPartListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDataPartListContext); ok {
			tst[i] = t.(IDataPartListContext)
			i++
		}
	}

	return tst
}

func (s *DataDefContext) DataPartList(i int) IDataPartListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataPartListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataPartListContext)
}

func (s *DataDefContext) AllEos() []IEosContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEosContext); ok {
			len++
		}
	}

	tst := make([]IEosContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEosContext); ok {
			tst[i] = t.(IEosContext)
			i++
		}
	}

	return tst
}

func (s *DataDefContext) Eos(i int) IEosContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *DataDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterDataDef(s)
	}
}

func (s *DataDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitDataDef(s)
	}
}

func (p *NeedleParser) DataDef() (localctx IDataDefContext) {
	localctx = NewDataDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, NeedleParserRULE_dataDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Match(NeedleParserDATA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.Match(NeedleParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == NeedleParserIdentifier {
		{
			p.SetState(149)
			p.DataPartList()
		}
		{
			p.SetState(150)
			p.Eos()
		}

		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(157)
		p.Match(NeedleParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDataPartListContext is an interface to support dynamic dispatch.
type IDataPartListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDataTag returns the dataTag rule contexts.
	GetDataTag() IStringLiteralContext

	// SetDataTag sets the dataTag rule contexts.
	SetDataTag(IStringLiteralContext)

	// Getter signatures
	Identifier() antlr.TerminalNode
	TypeName() ITypeNameContext
	StringLiteral() IStringLiteralContext

	// IsDataPartListContext differentiates from other interfaces.
	IsDataPartListContext()
}

type DataPartListContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	dataTag IStringLiteralContext
}

func NewEmptyDataPartListContext() *DataPartListContext {
	var p = new(DataPartListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_dataPartList
	return p
}

func InitEmptyDataPartListContext(p *DataPartListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_dataPartList
}

func (*DataPartListContext) IsDataPartListContext() {}

func NewDataPartListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataPartListContext {
	var p = new(DataPartListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_dataPartList

	return p
}

func (s *DataPartListContext) GetParser() antlr.Parser { return s.parser }

func (s *DataPartListContext) GetDataTag() IStringLiteralContext { return s.dataTag }

func (s *DataPartListContext) SetDataTag(v IStringLiteralContext) { s.dataTag = v }

func (s *DataPartListContext) Identifier() antlr.TerminalNode {
	return s.GetToken(NeedleParserIdentifier, 0)
}

func (s *DataPartListContext) TypeName() ITypeNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *DataPartListContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *DataPartListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataPartListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataPartListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterDataPartList(s)
	}
}

func (s *DataPartListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitDataPartList(s)
	}
}

func (p *NeedleParser) DataPartList() (localctx IDataPartListContext) {
	localctx = NewDataPartListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, NeedleParserRULE_dataPartList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Match(NeedleParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(160)
		p.TypeName()
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NeedleParserInterpretedStringLiteral || _la == NeedleParserRawStringLiteral {
		{
			p.SetState(161)

			var _x = p.StringLiteral()

			localctx.(*DataPartListContext).dataTag = _x
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISettingsDefContext is an interface to support dynamic dispatch.
type ISettingsDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SETTINGS() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode
	AllEQ() []antlr.TerminalNode
	EQ(i int) antlr.TerminalNode
	AllLiteral() []ILiteralContext
	Literal(i int) ILiteralContext
	AllEos() []IEosContext
	Eos(i int) IEosContext

	// IsSettingsDefContext differentiates from other interfaces.
	IsSettingsDefContext()
}

type SettingsDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySettingsDefContext() *SettingsDefContext {
	var p = new(SettingsDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_settingsDef
	return p
}

func InitEmptySettingsDefContext(p *SettingsDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_settingsDef
}

func (*SettingsDefContext) IsSettingsDefContext() {}

func NewSettingsDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SettingsDefContext {
	var p = new(SettingsDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_settingsDef

	return p
}

func (s *SettingsDefContext) GetParser() antlr.Parser { return s.parser }

func (s *SettingsDefContext) SETTINGS() antlr.TerminalNode {
	return s.GetToken(NeedleParserSETTINGS, 0)
}

func (s *SettingsDefContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(NeedleParserLBRACE, 0)
}

func (s *SettingsDefContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(NeedleParserRBRACE, 0)
}

func (s *SettingsDefContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(NeedleParserIdentifier)
}

func (s *SettingsDefContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(NeedleParserIdentifier, i)
}

func (s *SettingsDefContext) AllEQ() []antlr.TerminalNode {
	return s.GetTokens(NeedleParserEQ)
}

func (s *SettingsDefContext) EQ(i int) antlr.TerminalNode {
	return s.GetToken(NeedleParserEQ, i)
}

func (s *SettingsDefContext) AllLiteral() []ILiteralContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILiteralContext); ok {
			len++
		}
	}

	tst := make([]ILiteralContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILiteralContext); ok {
			tst[i] = t.(ILiteralContext)
			i++
		}
	}

	return tst
}

func (s *SettingsDefContext) Literal(i int) ILiteralContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *SettingsDefContext) AllEos() []IEosContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEosContext); ok {
			len++
		}
	}

	tst := make([]IEosContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEosContext); ok {
			tst[i] = t.(IEosContext)
			i++
		}
	}

	return tst
}

func (s *SettingsDefContext) Eos(i int) IEosContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *SettingsDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SettingsDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SettingsDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterSettingsDef(s)
	}
}

func (s *SettingsDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitSettingsDef(s)
	}
}

func (p *NeedleParser) SettingsDef() (localctx ISettingsDefContext) {
	localctx = NewSettingsDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, NeedleParserRULE_settingsDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.Match(NeedleParserSETTINGS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(165)
		p.Match(NeedleParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == NeedleParserIdentifier {
		{
			p.SetState(166)
			p.Match(NeedleParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(167)
			p.Match(NeedleParserEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.Literal()
		}
		{
			p.SetState(169)
			p.Eos()
		}

		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(176)
		p.Match(NeedleParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncDefContext is an interface to support dynamic dispatch.
type IFuncDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FuncSignature() IFuncSignatureContext
	Block() IBlockContext
	FuncDescriptor() IFuncDescriptorContext
	DefaultFuncDef() IDefaultFuncDefContext

	// IsFuncDefContext differentiates from other interfaces.
	IsFuncDefContext()
}

type FuncDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDefContext() *FuncDefContext {
	var p = new(FuncDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_funcDef
	return p
}

func InitEmptyFuncDefContext(p *FuncDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_funcDef
}

func (*FuncDefContext) IsFuncDefContext() {}

func NewFuncDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDefContext {
	var p = new(FuncDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_funcDef

	return p
}

func (s *FuncDefContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDefContext) FuncSignature() IFuncSignatureContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncSignatureContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncSignatureContext)
}

func (s *FuncDefContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncDefContext) FuncDescriptor() IFuncDescriptorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDescriptorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDescriptorContext)
}

func (s *FuncDefContext) DefaultFuncDef() IDefaultFuncDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultFuncDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultFuncDefContext)
}

func (s *FuncDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterFuncDef(s)
	}
}

func (s *FuncDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitFuncDef(s)
	}
}

func (p *NeedleParser) FuncDef() (localctx IFuncDefContext) {
	localctx = NewFuncDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, NeedleParserRULE_funcDef)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(178)
			p.FuncDescriptor()
		}

	case 2:
		{
			p.SetState(179)
			p.DefaultFuncDef()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(182)
		p.FuncSignature()
	}
	{
		p.SetState(183)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefaultFuncDefContext is an interface to support dynamic dispatch.
type IDefaultFuncDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONDITIONS() antlr.TerminalNode
	ACTION() antlr.TerminalNode
	FUNC() antlr.TerminalNode

	// IsDefaultFuncDefContext differentiates from other interfaces.
	IsDefaultFuncDefContext()
}

type DefaultFuncDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultFuncDefContext() *DefaultFuncDefContext {
	var p = new(DefaultFuncDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_defaultFuncDef
	return p
}

func InitEmptyDefaultFuncDefContext(p *DefaultFuncDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_defaultFuncDef
}

func (*DefaultFuncDefContext) IsDefaultFuncDefContext() {}

func NewDefaultFuncDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultFuncDefContext {
	var p = new(DefaultFuncDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_defaultFuncDef

	return p
}

func (s *DefaultFuncDefContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultFuncDefContext) CONDITIONS() antlr.TerminalNode {
	return s.GetToken(NeedleParserCONDITIONS, 0)
}

func (s *DefaultFuncDefContext) ACTION() antlr.TerminalNode {
	return s.GetToken(NeedleParserACTION, 0)
}

func (s *DefaultFuncDefContext) FUNC() antlr.TerminalNode {
	return s.GetToken(NeedleParserFUNC, 0)
}

func (s *DefaultFuncDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultFuncDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultFuncDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterDefaultFuncDef(s)
	}
}

func (s *DefaultFuncDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitDefaultFuncDef(s)
	}
}

func (p *NeedleParser) DefaultFuncDef() (localctx IDefaultFuncDefContext) {
	localctx = NewDefaultFuncDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, NeedleParserRULE_defaultFuncDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NeedleParserFUNC {
		{
			p.SetState(185)
			p.Match(NeedleParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(188)
		_la = p.GetTokenStream().LA(1)

		if !(_la == NeedleParserACTION || _la == NeedleParserCONDITIONS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncDescriptorContext is an interface to support dynamic dispatch.
type IFuncDescriptorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	Identifier() antlr.TerminalNode

	// IsFuncDescriptorContext differentiates from other interfaces.
	IsFuncDescriptorContext()
}

type FuncDescriptorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDescriptorContext() *FuncDescriptorContext {
	var p = new(FuncDescriptorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_funcDescriptor
	return p
}

func InitEmptyFuncDescriptorContext(p *FuncDescriptorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_funcDescriptor
}

func (*FuncDescriptorContext) IsFuncDescriptorContext() {}

func NewFuncDescriptorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDescriptorContext {
	var p = new(FuncDescriptorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_funcDescriptor

	return p
}

func (s *FuncDescriptorContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDescriptorContext) FUNC() antlr.TerminalNode {
	return s.GetToken(NeedleParserFUNC, 0)
}

func (s *FuncDescriptorContext) Identifier() antlr.TerminalNode {
	return s.GetToken(NeedleParserIdentifier, 0)
}

func (s *FuncDescriptorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDescriptorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDescriptorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterFuncDescriptor(s)
	}
}

func (s *FuncDescriptorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitFuncDescriptor(s)
	}
}

func (p *NeedleParser) FuncDescriptor() (localctx IFuncDescriptorContext) {
	localctx = NewFuncDescriptorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, NeedleParserRULE_funcDescriptor)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Match(NeedleParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(191)
		p.Match(NeedleParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncSignatureContext is an interface to support dynamic dispatch.
type IFuncSignatureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ParameterList() IParameterListContext
	AllFuncTail() []IFuncTailContext
	FuncTail(i int) IFuncTailContext
	ReturnParameters() IReturnParametersContext

	// IsFuncSignatureContext differentiates from other interfaces.
	IsFuncSignatureContext()
}

type FuncSignatureContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncSignatureContext() *FuncSignatureContext {
	var p = new(FuncSignatureContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_funcSignature
	return p
}

func InitEmptyFuncSignatureContext(p *FuncSignatureContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_funcSignature
}

func (*FuncSignatureContext) IsFuncSignatureContext() {}

func NewFuncSignatureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncSignatureContext {
	var p = new(FuncSignatureContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_funcSignature

	return p
}

func (s *FuncSignatureContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncSignatureContext) ParameterList() IParameterListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *FuncSignatureContext) AllFuncTail() []IFuncTailContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncTailContext); ok {
			len++
		}
	}

	tst := make([]IFuncTailContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncTailContext); ok {
			tst[i] = t.(IFuncTailContext)
			i++
		}
	}

	return tst
}

func (s *FuncSignatureContext) FuncTail(i int) IFuncTailContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncTailContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncTailContext)
}

func (s *FuncSignatureContext) ReturnParameters() IReturnParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnParametersContext)
}

func (s *FuncSignatureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncSignatureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncSignatureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterFuncSignature(s)
	}
}

func (s *FuncSignatureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitFuncSignature(s)
	}
}

func (p *NeedleParser) FuncSignature() (localctx IFuncSignatureContext) {
	localctx = NewFuncSignatureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, NeedleParserRULE_funcSignature)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NeedleParserLPAREN {
		{
			p.SetState(193)
			p.ParameterList()
		}

	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == NeedleParserDOT {
		{
			p.SetState(196)
			p.FuncTail()
		}

		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&1023) != 0 {
		{
			p.SetState(202)
			p.ReturnParameters()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncTailContext is an interface to support dynamic dispatch.
type IFuncTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOT() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	ParameterList() IParameterListContext

	// IsFuncTailContext differentiates from other interfaces.
	IsFuncTailContext()
}

type FuncTailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncTailContext() *FuncTailContext {
	var p = new(FuncTailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_funcTail
	return p
}

func InitEmptyFuncTailContext(p *FuncTailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_funcTail
}

func (*FuncTailContext) IsFuncTailContext() {}

func NewFuncTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncTailContext {
	var p = new(FuncTailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_funcTail

	return p
}

func (s *FuncTailContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncTailContext) DOT() antlr.TerminalNode {
	return s.GetToken(NeedleParserDOT, 0)
}

func (s *FuncTailContext) Identifier() antlr.TerminalNode {
	return s.GetToken(NeedleParserIdentifier, 0)
}

func (s *FuncTailContext) ParameterList() IParameterListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *FuncTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncTailContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterFuncTail(s)
	}
}

func (s *FuncTailContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitFuncTail(s)
	}
}

func (p *NeedleParser) FuncTail() (localctx IFuncTailContext) {
	localctx = NewFuncTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, NeedleParserRULE_funcTail)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(NeedleParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.Match(NeedleParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NeedleParserLPAREN {
		{
			p.SetState(207)
			p.ParameterList()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterListContext is an interface to support dynamic dispatch.
type IParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Parameter() IParameterContext
	COMMA() antlr.TerminalNode

	// IsParameterListContext differentiates from other interfaces.
	IsParameterListContext()
}

type ParameterListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterListContext() *ParameterListContext {
	var p = new(ParameterListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_parameterList
	return p
}

func InitEmptyParameterListContext(p *ParameterListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_parameterList
}

func (*ParameterListContext) IsParameterListContext() {}

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext {
	var p = new(ParameterListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_parameterList

	return p
}

func (s *ParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterListContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(NeedleParserLPAREN, 0)
}

func (s *ParameterListContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(NeedleParserRPAREN, 0)
}

func (s *ParameterListContext) Parameter() IParameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *ParameterListContext) COMMA() antlr.TerminalNode {
	return s.GetToken(NeedleParserCOMMA, 0)
}

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterParameterList(s)
	}
}

func (s *ParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitParameterList(s)
	}
}

func (p *NeedleParser) ParameterList() (localctx IParameterListContext) {
	localctx = NewParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, NeedleParserRULE_parameterList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		p.Match(NeedleParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NeedleParserIdentifier {
		{
			p.SetState(211)
			p.Parameter()
		}
		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == NeedleParserCOMMA {
			{
				p.SetState(212)
				p.Match(NeedleParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(217)
		p.Match(NeedleParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifierList() []IIdentifierListContext
	IdentifierList(i int) IIdentifierListContext
	AllTypeName() []ITypeNameContext
	TypeName(i int) ITypeNameContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_parameter
	return p
}

func InitEmptyParameterContext(p *ParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_parameter
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) AllIdentifierList() []IIdentifierListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierListContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierListContext); ok {
			tst[i] = t.(IIdentifierListContext)
			i++
		}
	}

	return tst
}

func (s *ParameterContext) IdentifierList(i int) IIdentifierListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *ParameterContext) AllTypeName() []ITypeNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeNameContext); ok {
			len++
		}
	}

	tst := make([]ITypeNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeNameContext); ok {
			tst[i] = t.(ITypeNameContext)
			i++
		}
	}

	return tst
}

func (s *ParameterContext) TypeName(i int) ITypeNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *ParameterContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(NeedleParserCOMMA)
}

func (s *ParameterContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(NeedleParserCOMMA, i)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (p *NeedleParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, NeedleParserRULE_parameter)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(219)
		p.IdentifierList()
	}
	{
		p.SetState(220)
		p.TypeName()
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(222)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == NeedleParserCOMMA {
				{
					p.SetState(221)
					p.Match(NeedleParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(224)
				p.IdentifierList()
			}
			{
				p.SetState(225)
				p.TypeName()
			}

		}
		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnParametersContext is an interface to support dynamic dispatch.
type IReturnParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTypeName() []ITypeNameContext
	TypeName(i int) ITypeNameContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsReturnParametersContext differentiates from other interfaces.
	IsReturnParametersContext()
}

type ReturnParametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnParametersContext() *ReturnParametersContext {
	var p = new(ReturnParametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_returnParameters
	return p
}

func InitEmptyReturnParametersContext(p *ReturnParametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_returnParameters
}

func (*ReturnParametersContext) IsReturnParametersContext() {}

func NewReturnParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnParametersContext {
	var p = new(ReturnParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_returnParameters

	return p
}

func (s *ReturnParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnParametersContext) AllTypeName() []ITypeNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeNameContext); ok {
			len++
		}
	}

	tst := make([]ITypeNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeNameContext); ok {
			tst[i] = t.(ITypeNameContext)
			i++
		}
	}

	return tst
}

func (s *ReturnParametersContext) TypeName(i int) ITypeNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *ReturnParametersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(NeedleParserCOMMA)
}

func (s *ReturnParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(NeedleParserCOMMA, i)
}

func (s *ReturnParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterReturnParameters(s)
	}
}

func (s *ReturnParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitReturnParameters(s)
	}
}

func (p *NeedleParser) ReturnParameters() (localctx IReturnParametersContext) {
	localctx = NewReturnParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, NeedleParserRULE_returnParameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		p.TypeName()
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == NeedleParserCOMMA || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&1023) != 0) {
		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == NeedleParserCOMMA {
			{
				p.SetState(233)
				p.Match(NeedleParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(236)
			p.TypeName()
		}

		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	StatementList() IStatementListContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(NeedleParserLBRACE, 0)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(NeedleParserRBRACE, 0)
}

func (s *BlockContext) StatementList() IStatementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *NeedleParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, NeedleParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(NeedleParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-8097929526849250686) != 0) || ((int64((_la-74)) & ^0x3f) == 0 && ((int64(1)<<(_la-74))&16778239) != 0) {
		{
			p.SetState(243)
			p.StatementList()
		}

	}
	{
		p.SetState(246)
		p.Match(NeedleParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	AllEos() []IEosContext
	Eos(i int) IEosContext
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode
	AllEOS() []antlr.TerminalNode
	EOS(i int) antlr.TerminalNode

	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_statementList
	return p
}

func InitEmptyStatementListContext(p *StatementListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_statementList
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementListContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementListContext) AllEos() []IEosContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEosContext); ok {
			len++
		}
	}

	tst := make([]IEosContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEosContext); ok {
			tst[i] = t.(IEosContext)
			i++
		}
	}

	return tst
}

func (s *StatementListContext) Eos(i int) IEosContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *StatementListContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(NeedleParserSEMI)
}

func (s *StatementListContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(NeedleParserSEMI, i)
}

func (s *StatementListContext) AllEOS() []antlr.TerminalNode {
	return s.GetTokens(NeedleParserEOS)
}

func (s *StatementListContext) EOS(i int) antlr.TerminalNode {
	return s.GetToken(NeedleParserEOS, i)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterStatementList(s)
	}
}

func (s *StatementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitStatementList(s)
	}
}

func (p *NeedleParser) StatementList() (localctx IStatementListContext) {
	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, NeedleParserRULE_statementList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-8097929526849250686) != 0) || ((int64((_la-74)) & ^0x3f) == 0 && ((int64(1)<<(_la-74))&16778239) != 0) {
		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
		case 1:
			p.SetState(249)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == NeedleParserSEMI {
				{
					p.SetState(248)
					p.Match(NeedleParserSEMI)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}

		case 2:
			p.SetState(252)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == NeedleParserEOS {
				{
					p.SetState(251)
					p.Match(NeedleParserEOS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		{
			p.SetState(256)
			p.Statement()
		}
		{
			p.SetState(257)
			p.Eos()
		}

		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	SimpleStmt() ISimpleStmtContext
	VarDef() IVarDefContext
	IfStmt() IIfStmtContext
	WhileStmt() IWhileStmtContext
	ContinueStmt() IContinueStmtContext
	BreakStmt() IBreakStmtContext
	ReturnStmt() IReturnStmtContext
	ErrorStmt() IErrorStmtContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) SimpleStmt() ISimpleStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStmtContext)
}

func (s *StatementContext) VarDef() IVarDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDefContext)
}

func (s *StatementContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StatementContext) WhileStmt() IWhileStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStmtContext)
}

func (s *StatementContext) ContinueStmt() IContinueStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinueStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinueStmtContext)
}

func (s *StatementContext) BreakStmt() IBreakStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakStmtContext)
}

func (s *StatementContext) ReturnStmt() IReturnStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStmtContext)
}

func (s *StatementContext) ErrorStmt() IErrorStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IErrorStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IErrorStmtContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *NeedleParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, NeedleParserRULE_statement)
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(263)
			p.Block()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(264)
			p.SimpleStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(265)
			p.VarDef()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(266)
			p.IfStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(267)
			p.WhileStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(268)
			p.ContinueStmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(269)
			p.BreakStmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(270)
			p.ReturnStmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(271)
			p.ErrorStmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimpleStmtContext is an interface to support dynamic dispatch.
type ISimpleStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Assignment() IAssignmentContext
	IncDecStmt() IIncDecStmtContext

	// IsSimpleStmtContext differentiates from other interfaces.
	IsSimpleStmtContext()
}

type SimpleStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleStmtContext() *SimpleStmtContext {
	var p = new(SimpleStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_simpleStmt
	return p
}

func InitEmptySimpleStmtContext(p *SimpleStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_simpleStmt
}

func (*SimpleStmtContext) IsSimpleStmtContext() {}

func NewSimpleStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleStmtContext {
	var p = new(SimpleStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_simpleStmt

	return p
}

func (s *SimpleStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SimpleStmtContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *SimpleStmtContext) IncDecStmt() IIncDecStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncDecStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncDecStmtContext)
}

func (s *SimpleStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterSimpleStmt(s)
	}
}

func (s *SimpleStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitSimpleStmt(s)
	}
}

func (p *NeedleParser) SimpleStmt() (localctx ISimpleStmtContext) {
	localctx = NewSimpleStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, NeedleParserRULE_simpleStmt)
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(274)
			p.expr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(275)
			p.Assignment()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(276)
			p.IncDecStmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIncDecStmtContext is an interface to support dynamic dispatch.
type IIncDecStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	IncDec_op() IIncDec_opContext

	// IsIncDecStmtContext differentiates from other interfaces.
	IsIncDecStmtContext()
}

type IncDecStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncDecStmtContext() *IncDecStmtContext {
	var p = new(IncDecStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_incDecStmt
	return p
}

func InitEmptyIncDecStmtContext(p *IncDecStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_incDecStmt
}

func (*IncDecStmtContext) IsIncDecStmtContext() {}

func NewIncDecStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncDecStmtContext {
	var p = new(IncDecStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_incDecStmt

	return p
}

func (s *IncDecStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IncDecStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IncDecStmtContext) IncDec_op() IIncDec_opContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncDec_opContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncDec_opContext)
}

func (s *IncDecStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncDecStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncDecStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterIncDecStmt(s)
	}
}

func (s *IncDecStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitIncDecStmt(s)
	}
}

func (p *NeedleParser) IncDecStmt() (localctx IIncDecStmtContext) {
	localctx = NewIncDecStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, NeedleParserRULE_incDecStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		p.expr(0)
	}
	{
		p.SetState(280)
		p.IncDec_op()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExprList() []IExprListContext
	ExprList(i int) IExprListContext
	Assign_op() IAssign_opContext

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) AllExprList() []IExprListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprListContext); ok {
			len++
		}
	}

	tst := make([]IExprListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprListContext); ok {
			tst[i] = t.(IExprListContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentContext) ExprList(i int) IExprListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *AssignmentContext) Assign_op() IAssign_opContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_opContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_opContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *NeedleParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, NeedleParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.ExprList()
	}
	{
		p.SetState(283)
		p.Assign_op()
	}
	{
		p.SetState(284)
		p.ExprList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarDefContext is an interface to support dynamic dispatch.
type IVarDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	Parameter() IParameterContext

	// IsVarDefContext differentiates from other interfaces.
	IsVarDefContext()
}

type VarDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDefContext() *VarDefContext {
	var p = new(VarDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_varDef
	return p
}

func InitEmptyVarDefContext(p *VarDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_varDef
}

func (*VarDefContext) IsVarDefContext() {}

func NewVarDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDefContext {
	var p = new(VarDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_varDef

	return p
}

func (s *VarDefContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDefContext) VAR() antlr.TerminalNode {
	return s.GetToken(NeedleParserVAR, 0)
}

func (s *VarDefContext) Parameter() IParameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *VarDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterVarDef(s)
	}
}

func (s *VarDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitVarDef(s)
	}
}

func (p *NeedleParser) VarDef() (localctx IVarDefContext) {
	localctx = NewVarDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, NeedleParserRULE_varDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		p.Match(NeedleParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(287)
		p.Parameter()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	AllIfBody() []IIfBodyContext
	IfBody(i int) IIfBodyContext
	AllELIF() []antlr.TerminalNode
	ELIF(i int) antlr.TerminalNode
	ElseBody() IElseBodyContext

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_ifStmt
	return p
}

func InitEmptyIfStmtContext(p *IfStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_ifStmt
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(NeedleParserIF, 0)
}

func (s *IfStmtContext) AllIfBody() []IIfBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIfBodyContext); ok {
			len++
		}
	}

	tst := make([]IIfBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIfBodyContext); ok {
			tst[i] = t.(IIfBodyContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) IfBody(i int) IIfBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfBodyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfBodyContext)
}

func (s *IfStmtContext) AllELIF() []antlr.TerminalNode {
	return s.GetTokens(NeedleParserELIF)
}

func (s *IfStmtContext) ELIF(i int) antlr.TerminalNode {
	return s.GetToken(NeedleParserELIF, i)
}

func (s *IfStmtContext) ElseBody() IElseBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseBodyContext)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (p *NeedleParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, NeedleParserRULE_ifStmt)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
		p.Match(NeedleParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(290)
		p.IfBody()
	}
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(291)
				p.Match(NeedleParserELIF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(292)
				p.IfBody()
			}

		}
		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NeedleParserELSE {
		{
			p.SetState(298)
			p.ElseBody()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfBodyContext is an interface to support dynamic dispatch.
type IIfBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	Expr() IExprContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsIfBodyContext differentiates from other interfaces.
	IsIfBodyContext()
}

type IfBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfBodyContext() *IfBodyContext {
	var p = new(IfBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_ifBody
	return p
}

func InitEmptyIfBodyContext(p *IfBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_ifBody
}

func (*IfBodyContext) IsIfBodyContext() {}

func NewIfBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfBodyContext {
	var p = new(IfBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_ifBody

	return p
}

func (s *IfBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *IfBodyContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfBodyContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfBodyContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(NeedleParserLPAREN, 0)
}

func (s *IfBodyContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(NeedleParserRPAREN, 0)
}

func (s *IfBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterIfBody(s)
	}
}

func (s *IfBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitIfBody(s)
	}
}

func (p *NeedleParser) IfBody() (localctx IIfBodyContext) {
	localctx = NewIfBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, NeedleParserRULE_ifBody)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(301)
			p.Match(NeedleParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(302)
			p.expr(0)
		}
		{
			p.SetState(303)
			p.Match(NeedleParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(305)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(308)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseBodyContext is an interface to support dynamic dispatch.
type IElseBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	Block() IBlockContext

	// IsElseBodyContext differentiates from other interfaces.
	IsElseBodyContext()
}

type ElseBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseBodyContext() *ElseBodyContext {
	var p = new(ElseBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_elseBody
	return p
}

func InitEmptyElseBodyContext(p *ElseBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_elseBody
}

func (*ElseBodyContext) IsElseBodyContext() {}

func NewElseBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseBodyContext {
	var p = new(ElseBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_elseBody

	return p
}

func (s *ElseBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseBodyContext) ELSE() antlr.TerminalNode {
	return s.GetToken(NeedleParserELSE, 0)
}

func (s *ElseBodyContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterElseBody(s)
	}
}

func (s *ElseBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitElseBody(s)
	}
}

func (p *NeedleParser) ElseBody() (localctx IElseBodyContext) {
	localctx = NewElseBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, NeedleParserRULE_elseBody)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Match(NeedleParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(311)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnStmtContext is an interface to support dynamic dispatch.
type IReturnStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expr() IExprContext

	// IsReturnStmtContext differentiates from other interfaces.
	IsReturnStmtContext()
}

type ReturnStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStmtContext() *ReturnStmtContext {
	var p = new(ReturnStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_returnStmt
	return p
}

func InitEmptyReturnStmtContext(p *ReturnStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_returnStmt
}

func (*ReturnStmtContext) IsReturnStmtContext() {}

func NewReturnStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_returnStmt

	return p
}

func (s *ReturnStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(NeedleParserRETURN, 0)
}

func (s *ReturnStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (p *NeedleParser) ReturnStmt() (localctx IReturnStmtContext) {
	localctx = NewReturnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, NeedleParserRULE_returnStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(313)
		p.Match(NeedleParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(315)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(314)
			p.expr(0)
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContinueStmtContext is an interface to support dynamic dispatch.
type IContinueStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTINUE() antlr.TerminalNode

	// IsContinueStmtContext differentiates from other interfaces.
	IsContinueStmtContext()
}

type ContinueStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStmtContext() *ContinueStmtContext {
	var p = new(ContinueStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_continueStmt
	return p
}

func InitEmptyContinueStmtContext(p *ContinueStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_continueStmt
}

func (*ContinueStmtContext) IsContinueStmtContext() {}

func NewContinueStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStmtContext {
	var p = new(ContinueStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_continueStmt

	return p
}

func (s *ContinueStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinueStmtContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(NeedleParserCONTINUE, 0)
}

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterContinueStmt(s)
	}
}

func (s *ContinueStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitContinueStmt(s)
	}
}

func (p *NeedleParser) ContinueStmt() (localctx IContinueStmtContext) {
	localctx = NewContinueStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, NeedleParserRULE_continueStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(317)
		p.Match(NeedleParserCONTINUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBreakStmtContext is an interface to support dynamic dispatch.
type IBreakStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BREAK() antlr.TerminalNode

	// IsBreakStmtContext differentiates from other interfaces.
	IsBreakStmtContext()
}

type BreakStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStmtContext() *BreakStmtContext {
	var p = new(BreakStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_breakStmt
	return p
}

func InitEmptyBreakStmtContext(p *BreakStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_breakStmt
}

func (*BreakStmtContext) IsBreakStmtContext() {}

func NewBreakStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStmtContext {
	var p = new(BreakStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_breakStmt

	return p
}

func (s *BreakStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakStmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(NeedleParserBREAK, 0)
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterBreakStmt(s)
	}
}

func (s *BreakStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitBreakStmt(s)
	}
}

func (p *NeedleParser) BreakStmt() (localctx IBreakStmtContext) {
	localctx = NewBreakStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, NeedleParserRULE_breakStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(319)
		p.Match(NeedleParserBREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhileStmtContext is an interface to support dynamic dispatch.
type IWhileStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expr() IExprContext
	Block() IBlockContext

	// IsWhileStmtContext differentiates from other interfaces.
	IsWhileStmtContext()
}

type WhileStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStmtContext() *WhileStmtContext {
	var p = new(WhileStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_whileStmt
	return p
}

func InitEmptyWhileStmtContext(p *WhileStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_whileStmt
}

func (*WhileStmtContext) IsWhileStmtContext() {}

func NewWhileStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStmtContext {
	var p = new(WhileStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_whileStmt

	return p
}

func (s *WhileStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStmtContext) WHILE() antlr.TerminalNode {
	return s.GetToken(NeedleParserWHILE, 0)
}

func (s *WhileStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhileStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterWhileStmt(s)
	}
}

func (s *WhileStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitWhileStmt(s)
	}
}

func (p *NeedleParser) WhileStmt() (localctx IWhileStmtContext) {
	localctx = NewWhileStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, NeedleParserRULE_whileStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(321)
		p.Match(NeedleParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(322)
		p.expr(0)
	}
	{
		p.SetState(323)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IErrorStmtContext is an interface to support dynamic dispatch.
type IErrorStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	ERRWARNING() antlr.TerminalNode
	ERRINFO() antlr.TerminalNode
	ERROR() antlr.TerminalNode

	// IsErrorStmtContext differentiates from other interfaces.
	IsErrorStmtContext()
}

type ErrorStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyErrorStmtContext() *ErrorStmtContext {
	var p = new(ErrorStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_errorStmt
	return p
}

func InitEmptyErrorStmtContext(p *ErrorStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_errorStmt
}

func (*ErrorStmtContext) IsErrorStmtContext() {}

func NewErrorStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ErrorStmtContext {
	var p = new(ErrorStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_errorStmt

	return p
}

func (s *ErrorStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ErrorStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ErrorStmtContext) ERRWARNING() antlr.TerminalNode {
	return s.GetToken(NeedleParserERRWARNING, 0)
}

func (s *ErrorStmtContext) ERRINFO() antlr.TerminalNode {
	return s.GetToken(NeedleParserERRINFO, 0)
}

func (s *ErrorStmtContext) ERROR() antlr.TerminalNode {
	return s.GetToken(NeedleParserERROR, 0)
}

func (s *ErrorStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ErrorStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ErrorStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterErrorStmt(s)
	}
}

func (s *ErrorStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitErrorStmt(s)
	}
}

func (p *NeedleParser) ErrorStmt() (localctx IErrorStmtContext) {
	localctx = NewErrorStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, NeedleParserRULE_errorStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(325)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-8791026472627208192) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(326)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayExprContext is an interface to support dynamic dispatch.
type IArrayExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	ArrayList() IArrayListContext
	COMMA() antlr.TerminalNode

	// IsArrayExprContext differentiates from other interfaces.
	IsArrayExprContext()
}

type ArrayExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayExprContext() *ArrayExprContext {
	var p = new(ArrayExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_arrayExpr
	return p
}

func InitEmptyArrayExprContext(p *ArrayExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_arrayExpr
}

func (*ArrayExprContext) IsArrayExprContext() {}

func NewArrayExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayExprContext {
	var p = new(ArrayExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_arrayExpr

	return p
}

func (s *ArrayExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayExprContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(NeedleParserLBRACK, 0)
}

func (s *ArrayExprContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(NeedleParserRBRACK, 0)
}

func (s *ArrayExprContext) ArrayList() IArrayListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayListContext)
}

func (s *ArrayExprContext) COMMA() antlr.TerminalNode {
	return s.GetToken(NeedleParserCOMMA, 0)
}

func (s *ArrayExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterArrayExpr(s)
	}
}

func (s *ArrayExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitArrayExpr(s)
	}
}

func (p *NeedleParser) ArrayExpr() (localctx IArrayExprContext) {
	localctx = NewArrayExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, NeedleParserRULE_arrayExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(328)
		p.Match(NeedleParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&579838452024005250) != 0) || ((int64((_la-74)) & ^0x3f) == 0 && ((int64(1)<<(_la-74))&1023) != 0) {
		{
			p.SetState(329)
			p.ArrayList()
		}

	}
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NeedleParserCOMMA {
		{
			p.SetState(332)
			p.Match(NeedleParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(335)
		p.Match(NeedleParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayListContext is an interface to support dynamic dispatch.
type IArrayListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExprList() IExprListContext
	Eos() IEosContext

	// IsArrayListContext differentiates from other interfaces.
	IsArrayListContext()
}

type ArrayListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayListContext() *ArrayListContext {
	var p = new(ArrayListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_arrayList
	return p
}

func InitEmptyArrayListContext(p *ArrayListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_arrayList
}

func (*ArrayListContext) IsArrayListContext() {}

func NewArrayListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayListContext {
	var p = new(ArrayListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_arrayList

	return p
}

func (s *ArrayListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayListContext) ExprList() IExprListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *ArrayListContext) Eos() IEosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ArrayListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterArrayList(s)
	}
}

func (s *ArrayListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitArrayList(s)
	}
}

func (p *NeedleParser) ArrayList() (localctx IArrayListContext) {
	localctx = NewArrayListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, NeedleParserRULE_arrayList)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		p.ExprList()
	}
	{
		p.SetState(338)
		p.Eos()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMapExprContext is an interface to support dynamic dispatch.
type IMapExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	PairList() IPairListContext

	// IsMapExprContext differentiates from other interfaces.
	IsMapExprContext()
}

type MapExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapExprContext() *MapExprContext {
	var p = new(MapExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_mapExpr
	return p
}

func InitEmptyMapExprContext(p *MapExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_mapExpr
}

func (*MapExprContext) IsMapExprContext() {}

func NewMapExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapExprContext {
	var p = new(MapExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_mapExpr

	return p
}

func (s *MapExprContext) GetParser() antlr.Parser { return s.parser }

func (s *MapExprContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(NeedleParserLBRACE, 0)
}

func (s *MapExprContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(NeedleParserRBRACE, 0)
}

func (s *MapExprContext) PairList() IPairListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPairListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPairListContext)
}

func (s *MapExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterMapExpr(s)
	}
}

func (s *MapExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitMapExpr(s)
	}
}

func (p *NeedleParser) MapExpr() (localctx IMapExprContext) {
	localctx = NewMapExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, NeedleParserRULE_mapExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(340)
		p.Match(NeedleParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-74)) & ^0x3f) == 0 && ((int64(1)<<(_la-74))&27) != 0 {
		{
			p.SetState(341)
			p.PairList()
		}

	}
	{
		p.SetState(344)
		p.Match(NeedleParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPairListContext is an interface to support dynamic dispatch.
type IPairListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPair() []IPairContext
	Pair(i int) IPairContext
	Eos() IEosContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsPairListContext differentiates from other interfaces.
	IsPairListContext()
}

type PairListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairListContext() *PairListContext {
	var p = new(PairListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_pairList
	return p
}

func InitEmptyPairListContext(p *PairListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_pairList
}

func (*PairListContext) IsPairListContext() {}

func NewPairListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairListContext {
	var p = new(PairListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_pairList

	return p
}

func (s *PairListContext) GetParser() antlr.Parser { return s.parser }

func (s *PairListContext) AllPair() []IPairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPairContext); ok {
			len++
		}
	}

	tst := make([]IPairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPairContext); ok {
			tst[i] = t.(IPairContext)
			i++
		}
	}

	return tst
}

func (s *PairListContext) Pair(i int) IPairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPairContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *PairListContext) Eos() IEosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *PairListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(NeedleParserCOMMA)
}

func (s *PairListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(NeedleParserCOMMA, i)
}

func (s *PairListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterPairList(s)
	}
}

func (s *PairListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitPairList(s)
	}
}

func (p *NeedleParser) PairList() (localctx IPairListContext) {
	localctx = NewPairListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, NeedleParserRULE_pairList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(346)
		p.Pair()
	}
	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(347)
				p.Match(NeedleParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(348)
				p.Pair()
			}

		}
		p.SetState(353)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(355)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NeedleParserCOMMA {
		{
			p.SetState(354)
			p.Match(NeedleParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(357)
		p.Eos()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLON() antlr.TerminalNode
	PairValue() IPairValueContext
	StringLiteral() IStringLiteralContext
	IdentifierVar() IIdentifierVarContext

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_pair
	return p
}

func InitEmptyPairContext(p *PairContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_pair
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) COLON() antlr.TerminalNode {
	return s.GetToken(NeedleParserCOLON, 0)
}

func (s *PairContext) PairValue() IPairValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPairValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPairValueContext)
}

func (s *PairContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *PairContext) IdentifierVar() IIdentifierVarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierVarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierVarContext)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitPair(s)
	}
}

func (p *NeedleParser) Pair() (localctx IPairContext) {
	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, NeedleParserRULE_pair)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NeedleParserInterpretedStringLiteral, NeedleParserRawStringLiteral:
		{
			p.SetState(359)
			p.StringLiteral()
		}

	case NeedleParserIdentifier, NeedleParserDollarIdentifier:
		{
			p.SetState(360)
			p.IdentifierVar()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(363)
		p.Match(NeedleParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(364)
		p.PairValue()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPairValueContext is an interface to support dynamic dispatch.
type IPairValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentifierVar() IIdentifierVarContext
	IndexExpr() IIndexExprContext
	SliceExpr() ISliceExprContext
	Literal() ILiteralContext
	MapExpr() IMapExprContext
	ArrayExpr() IArrayExprContext

	// IsPairValueContext differentiates from other interfaces.
	IsPairValueContext()
}

type PairValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairValueContext() *PairValueContext {
	var p = new(PairValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_pairValue
	return p
}

func InitEmptyPairValueContext(p *PairValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_pairValue
}

func (*PairValueContext) IsPairValueContext() {}

func NewPairValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairValueContext {
	var p = new(PairValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_pairValue

	return p
}

func (s *PairValueContext) GetParser() antlr.Parser { return s.parser }

func (s *PairValueContext) IdentifierVar() IIdentifierVarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierVarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierVarContext)
}

func (s *PairValueContext) IndexExpr() IIndexExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexExprContext)
}

func (s *PairValueContext) SliceExpr() ISliceExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceExprContext)
}

func (s *PairValueContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *PairValueContext) MapExpr() IMapExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapExprContext)
}

func (s *PairValueContext) ArrayExpr() IArrayExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayExprContext)
}

func (s *PairValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterPairValue(s)
	}
}

func (s *PairValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitPairValue(s)
	}
}

func (p *NeedleParser) PairValue() (localctx IPairValueContext) {
	localctx = NewPairValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, NeedleParserRULE_pairValue)
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NeedleParserIdentifier, NeedleParserDollarIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(366)
			p.IdentifierVar()
		}
		p.SetState(369)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(367)
				p.IndexExpr()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) == 2 {
			{
				p.SetState(368)
				p.SliceExpr()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case NeedleParserTRUE, NeedleParserFALSE, NeedleParserNIL, NeedleParserInterpretedStringLiteral, NeedleParserRawStringLiteral, NeedleParserDecimalLiteral, NeedleParserFloatLiteral, NeedleParserHexLiteral, NeedleParserOctalLiteral, NeedleParserBinaryLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(371)
			p.Literal()
		}

	case NeedleParserLBRACE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(372)
			p.MapExpr()
		}

	case NeedleParserLBRACK:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(373)
			p.ArrayExpr()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ExprList() IExprListContext

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_arguments
	return p
}

func InitEmptyArgumentsContext(p *ArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_arguments
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(NeedleParserLPAREN, 0)
}

func (s *ArgumentsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(NeedleParserRPAREN, 0)
}

func (s *ArgumentsContext) ExprList() IExprListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *NeedleParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, NeedleParserRULE_arguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(376)
		p.Match(NeedleParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&579838452024005250) != 0) || ((int64((_la-74)) & ^0x3f) == 0 && ((int64(1)<<(_la-74))&1023) != 0) {
		{
			p.SetState(377)
			p.ExprList()
		}

	}
	{
		p.SetState(380)
		p.Match(NeedleParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprListContext is an interface to support dynamic dispatch.
type IExprListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsExprListContext differentiates from other interfaces.
	IsExprListContext()
}

type ExprListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprListContext() *ExprListContext {
	var p = new(ExprListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_exprList
	return p
}

func InitEmptyExprListContext(p *ExprListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_exprList
}

func (*ExprListContext) IsExprListContext() {}

func NewExprListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprListContext {
	var p = new(ExprListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_exprList

	return p
}

func (s *ExprListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprListContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprListContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(NeedleParserCOMMA)
}

func (s *ExprListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(NeedleParserCOMMA, i)
}

func (s *ExprListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterExprList(s)
	}
}

func (s *ExprListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitExprList(s)
	}
}

func (p *NeedleParser) ExprList() (localctx IExprListContext) {
	localctx = NewExprListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, NeedleParserRULE_exprList)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(382)
		p.expr(0)
	}
	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(383)
				p.Match(NeedleParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(384)
				p.expr(0)
			}

		}
		p.SetState(389)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimaryExpr() IPrimaryExprContext
	Eos() IEosContext
	MapExpr() IMapExprContext
	ArrayExpr() IArrayExprContext
	Unary_op() IUnary_opContext
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	Mul_op() IMul_opContext
	Rel_op() IRel_opContext
	Logical_op() ILogical_opContext
	Add_op() IAdd_opContext
	IndexExpr() IIndexExprContext
	SliceExpr() ISliceExprContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) PrimaryExpr() IPrimaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *ExprContext) Eos() IEosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ExprContext) MapExpr() IMapExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapExprContext)
}

func (s *ExprContext) ArrayExpr() IArrayExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayExprContext)
}

func (s *ExprContext) Unary_op() IUnary_opContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnary_opContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnary_opContext)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) Mul_op() IMul_opContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMul_opContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMul_opContext)
}

func (s *ExprContext) Rel_op() IRel_opContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRel_opContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRel_opContext)
}

func (s *ExprContext) Logical_op() ILogical_opContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogical_opContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogical_opContext)
}

func (s *ExprContext) Add_op() IAdd_opContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdd_opContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdd_opContext)
}

func (s *ExprContext) IndexExpr() IIndexExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexExprContext)
}

func (s *ExprContext) SliceExpr() ISliceExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceExprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *NeedleParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *NeedleParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 74
	p.EnterRecursionRule(localctx, 74, NeedleParserRULE_expr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(399)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NeedleParserLPAREN, NeedleParserTRUE, NeedleParserFALSE, NeedleParserNIL, NeedleParserIdentifier, NeedleParserDollarIdentifier, NeedleParserAtIdentifier, NeedleParserInterpretedStringLiteral, NeedleParserRawStringLiteral, NeedleParserDecimalLiteral, NeedleParserFloatLiteral, NeedleParserHexLiteral, NeedleParserOctalLiteral, NeedleParserBinaryLiteral:
		{
			p.SetState(391)
			p.primaryExpr(0)
		}
		{
			p.SetState(392)
			p.Eos()
		}

	case NeedleParserLBRACE:
		{
			p.SetState(394)
			p.MapExpr()
		}

	case NeedleParserLBRACK:
		{
			p.SetState(395)
			p.ArrayExpr()
		}

	case NeedleParserNOT, NeedleParserADD, NeedleParserSUB:
		{
			p.SetState(396)
			p.Unary_op()
		}
		{
			p.SetState(397)
			p.expr(5)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(423)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(421)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NeedleParserRULE_expr)
				p.SetState(401)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(402)
					p.Mul_op()
				}
				{
					p.SetState(403)
					p.expr(5)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NeedleParserRULE_expr)
				p.SetState(405)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(406)
					p.Rel_op()
				}
				{
					p.SetState(407)
					p.expr(4)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NeedleParserRULE_expr)
				p.SetState(409)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(410)
					p.Logical_op()
				}
				{
					p.SetState(411)
					p.expr(3)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NeedleParserRULE_expr)
				p.SetState(413)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(414)
					p.Add_op()
				}
				{
					p.SetState(415)
					p.expr(2)
				}

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NeedleParserRULE_expr)
				p.SetState(417)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(418)
					p.IndexExpr()
				}

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NeedleParserRULE_expr)
				p.SetState(419)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(420)
					p.SliceExpr()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(425)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryExprContext is an interface to support dynamic dispatch.
type IPrimaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Operand() IOperandContext
	PrimaryExpr() IPrimaryExprContext
	Arguments() IArgumentsContext
	DOT() antlr.TerminalNode
	Identifier() antlr.TerminalNode

	// IsPrimaryExprContext differentiates from other interfaces.
	IsPrimaryExprContext()
}

type PrimaryExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExprContext() *PrimaryExprContext {
	var p = new(PrimaryExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_primaryExpr
	return p
}

func InitEmptyPrimaryExprContext(p *PrimaryExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_primaryExpr
}

func (*PrimaryExprContext) IsPrimaryExprContext() {}

func NewPrimaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_primaryExpr

	return p
}

func (s *PrimaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExprContext) Operand() IOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *PrimaryExprContext) PrimaryExpr() IPrimaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *PrimaryExprContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *PrimaryExprContext) DOT() antlr.TerminalNode {
	return s.GetToken(NeedleParserDOT, 0)
}

func (s *PrimaryExprContext) Identifier() antlr.TerminalNode {
	return s.GetToken(NeedleParserIdentifier, 0)
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitPrimaryExpr(s)
	}
}

func (p *NeedleParser) PrimaryExpr() (localctx IPrimaryExprContext) {
	return p.primaryExpr(0)
}

func (p *NeedleParser) primaryExpr(_p int) (localctx IPrimaryExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewPrimaryExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimaryExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 76
	p.EnterRecursionRule(localctx, 76, NeedleParserRULE_primaryExpr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(427)
		p.Operand()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(437)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPrimaryExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, NeedleParserRULE_primaryExpr)
			p.SetState(429)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			p.SetState(432)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == NeedleParserDOT {
				{
					p.SetState(430)
					p.Match(NeedleParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(431)
					p.Match(NeedleParserIdentifier)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(434)
				p.Arguments()
			}

		}
		p.SetState(439)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIndexExprContext is an interface to support dynamic dispatch.
type IIndexExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	Expr() IExprContext
	RBRACK() antlr.TerminalNode

	// IsIndexExprContext differentiates from other interfaces.
	IsIndexExprContext()
}

type IndexExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexExprContext() *IndexExprContext {
	var p = new(IndexExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_indexExpr
	return p
}

func InitEmptyIndexExprContext(p *IndexExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_indexExpr
}

func (*IndexExprContext) IsIndexExprContext() {}

func NewIndexExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexExprContext {
	var p = new(IndexExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_indexExpr

	return p
}

func (s *IndexExprContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexExprContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(NeedleParserLBRACK, 0)
}

func (s *IndexExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IndexExprContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(NeedleParserRBRACK, 0)
}

func (s *IndexExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterIndexExpr(s)
	}
}

func (s *IndexExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitIndexExpr(s)
	}
}

func (p *NeedleParser) IndexExpr() (localctx IIndexExprContext) {
	localctx = NewIndexExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, NeedleParserRULE_indexExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(440)
		p.Match(NeedleParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(441)
		p.expr(0)
	}
	{
		p.SetState(442)
		p.Match(NeedleParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISliceExprContext is an interface to support dynamic dispatch.
type ISliceExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	COLON() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsSliceExprContext differentiates from other interfaces.
	IsSliceExprContext()
}

type SliceExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceExprContext() *SliceExprContext {
	var p = new(SliceExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_sliceExpr
	return p
}

func InitEmptySliceExprContext(p *SliceExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_sliceExpr
}

func (*SliceExprContext) IsSliceExprContext() {}

func NewSliceExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceExprContext {
	var p = new(SliceExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_sliceExpr

	return p
}

func (s *SliceExprContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceExprContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(NeedleParserLBRACK, 0)
}

func (s *SliceExprContext) COLON() antlr.TerminalNode {
	return s.GetToken(NeedleParserCOLON, 0)
}

func (s *SliceExprContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(NeedleParserRBRACK, 0)
}

func (s *SliceExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *SliceExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SliceExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterSliceExpr(s)
	}
}

func (s *SliceExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitSliceExpr(s)
	}
}

func (p *NeedleParser) SliceExpr() (localctx ISliceExprContext) {
	localctx = NewSliceExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, NeedleParserRULE_sliceExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(444)
		p.Match(NeedleParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(446)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&579838452024005250) != 0) || ((int64((_la-74)) & ^0x3f) == 0 && ((int64(1)<<(_la-74))&1023) != 0) {
		{
			p.SetState(445)
			p.expr(0)
		}

	}
	{
		p.SetState(448)
		p.Match(NeedleParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(450)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&579838452024005250) != 0) || ((int64((_la-74)) & ^0x3f) == 0 && ((int64(1)<<(_la-74))&1023) != 0) {
		{
			p.SetState(449)
			p.expr(0)
		}

	}
	{
		p.SetState(452)
		p.Match(NeedleParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentifierFull() IIdentifierFullContext
	Literal() ILiteralContext
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode

	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_operand
	return p
}

func InitEmptyOperandContext(p *OperandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_operand
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) IdentifierFull() IIdentifierFullContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierFullContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierFullContext)
}

func (s *OperandContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *OperandContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(NeedleParserLPAREN, 0)
}

func (s *OperandContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OperandContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(NeedleParserRPAREN, 0)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterOperand(s)
	}
}

func (s *OperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitOperand(s)
	}
}

func (p *NeedleParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, NeedleParserRULE_operand)
	p.SetState(460)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NeedleParserIdentifier, NeedleParserDollarIdentifier, NeedleParserAtIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(454)
			p.IdentifierFull()
		}

	case NeedleParserTRUE, NeedleParserFALSE, NeedleParserNIL, NeedleParserInterpretedStringLiteral, NeedleParserRawStringLiteral, NeedleParserDecimalLiteral, NeedleParserFloatLiteral, NeedleParserHexLiteral, NeedleParserOctalLiteral, NeedleParserBinaryLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(455)
			p.Literal()
		}

	case NeedleParserLPAREN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(456)
			p.Match(NeedleParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(457)
			p.expr(0)
		}
		{
			p.SetState(458)
			p.Match(NeedleParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NumberLiteral() INumberLiteralContext
	StringLiteral() IStringLiteralContext
	BooleanLiteral() IBooleanLiteralContext
	NIL() antlr.TerminalNode

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) NumberLiteral() INumberLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberLiteralContext)
}

func (s *LiteralContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *LiteralContext) BooleanLiteral() IBooleanLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *LiteralContext) NIL() antlr.TerminalNode {
	return s.GetToken(NeedleParserNIL, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *NeedleParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, NeedleParserRULE_literal)
	p.SetState(466)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NeedleParserDecimalLiteral, NeedleParserFloatLiteral, NeedleParserHexLiteral, NeedleParserOctalLiteral, NeedleParserBinaryLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(462)
			p.NumberLiteral()
		}

	case NeedleParserInterpretedStringLiteral, NeedleParserRawStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(463)
			p.StringLiteral()
		}

	case NeedleParserTRUE, NeedleParserFALSE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(464)
			p.BooleanLiteral()
		}

	case NeedleParserNIL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(465)
			p.Match(NeedleParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeNameContext is an interface to support dynamic dispatch.
type ITypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BOOL() antlr.TerminalNode
	BYTES() antlr.TerminalNode
	INT() antlr.TerminalNode
	ADDRESS() antlr.TerminalNode
	ARRAY() antlr.TerminalNode
	MAP() antlr.TerminalNode
	MONEY() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	STRING() antlr.TerminalNode
	FILE() antlr.TerminalNode

	// IsTypeNameContext differentiates from other interfaces.
	IsTypeNameContext()
}

type TypeNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeNameContext() *TypeNameContext {
	var p = new(TypeNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_typeName
	return p
}

func InitEmptyTypeNameContext(p *TypeNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_typeName
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) BOOL() antlr.TerminalNode {
	return s.GetToken(NeedleParserBOOL, 0)
}

func (s *TypeNameContext) BYTES() antlr.TerminalNode {
	return s.GetToken(NeedleParserBYTES, 0)
}

func (s *TypeNameContext) INT() antlr.TerminalNode {
	return s.GetToken(NeedleParserINT, 0)
}

func (s *TypeNameContext) ADDRESS() antlr.TerminalNode {
	return s.GetToken(NeedleParserADDRESS, 0)
}

func (s *TypeNameContext) ARRAY() antlr.TerminalNode {
	return s.GetToken(NeedleParserARRAY, 0)
}

func (s *TypeNameContext) MAP() antlr.TerminalNode {
	return s.GetToken(NeedleParserMAP, 0)
}

func (s *TypeNameContext) MONEY() antlr.TerminalNode {
	return s.GetToken(NeedleParserMONEY, 0)
}

func (s *TypeNameContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(NeedleParserFLOAT, 0)
}

func (s *TypeNameContext) STRING() antlr.TerminalNode {
	return s.GetToken(NeedleParserSTRING, 0)
}

func (s *TypeNameContext) FILE() antlr.TerminalNode {
	return s.GetToken(NeedleParserFILE, 0)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterTypeName(s)
	}
}

func (s *TypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitTypeName(s)
	}
}

func (p *NeedleParser) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, NeedleParserRULE_typeName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(468)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&1023) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIncDec_opContext is an interface to support dynamic dispatch.
type IIncDec_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INC() antlr.TerminalNode
	DEC() antlr.TerminalNode

	// IsIncDec_opContext differentiates from other interfaces.
	IsIncDec_opContext()
}

type IncDec_opContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncDec_opContext() *IncDec_opContext {
	var p = new(IncDec_opContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_incDec_op
	return p
}

func InitEmptyIncDec_opContext(p *IncDec_opContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_incDec_op
}

func (*IncDec_opContext) IsIncDec_opContext() {}

func NewIncDec_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncDec_opContext {
	var p = new(IncDec_opContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_incDec_op

	return p
}

func (s *IncDec_opContext) GetParser() antlr.Parser { return s.parser }

func (s *IncDec_opContext) INC() antlr.TerminalNode {
	return s.GetToken(NeedleParserINC, 0)
}

func (s *IncDec_opContext) DEC() antlr.TerminalNode {
	return s.GetToken(NeedleParserDEC, 0)
}

func (s *IncDec_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncDec_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncDec_opContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterIncDec_op(s)
	}
}

func (s *IncDec_opContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitIncDec_op(s)
	}
}

func (p *NeedleParser) IncDec_op() (localctx IIncDec_opContext) {
	localctx = NewIncDec_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, NeedleParserRULE_incDec_op)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(470)
		_la = p.GetTokenStream().LA(1)

		if !(_la == NeedleParserINC || _la == NeedleParserDEC) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMul_opContext is an interface to support dynamic dispatch.
type IMul_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MUL() antlr.TerminalNode
	QUO() antlr.TerminalNode
	MOD() antlr.TerminalNode
	LSHIFT() antlr.TerminalNode
	RSHIFT() antlr.TerminalNode
	BIT_AND() antlr.TerminalNode

	// IsMul_opContext differentiates from other interfaces.
	IsMul_opContext()
}

type Mul_opContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMul_opContext() *Mul_opContext {
	var p = new(Mul_opContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_mul_op
	return p
}

func InitEmptyMul_opContext(p *Mul_opContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_mul_op
}

func (*Mul_opContext) IsMul_opContext() {}

func NewMul_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mul_opContext {
	var p = new(Mul_opContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_mul_op

	return p
}

func (s *Mul_opContext) GetParser() antlr.Parser { return s.parser }

func (s *Mul_opContext) MUL() antlr.TerminalNode {
	return s.GetToken(NeedleParserMUL, 0)
}

func (s *Mul_opContext) QUO() antlr.TerminalNode {
	return s.GetToken(NeedleParserQUO, 0)
}

func (s *Mul_opContext) MOD() antlr.TerminalNode {
	return s.GetToken(NeedleParserMOD, 0)
}

func (s *Mul_opContext) LSHIFT() antlr.TerminalNode {
	return s.GetToken(NeedleParserLSHIFT, 0)
}

func (s *Mul_opContext) RSHIFT() antlr.TerminalNode {
	return s.GetToken(NeedleParserRSHIFT, 0)
}

func (s *Mul_opContext) BIT_AND() antlr.TerminalNode {
	return s.GetToken(NeedleParserBIT_AND, 0)
}

func (s *Mul_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mul_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mul_opContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterMul_op(s)
	}
}

func (s *Mul_opContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitMul_op(s)
	}
}

func (p *NeedleParser) Mul_op() (localctx IMul_opContext) {
	localctx = NewMul_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, NeedleParserRULE_mul_op)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(472)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1912676352) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnary_opContext is an interface to support dynamic dispatch.
type IUnary_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ADD() antlr.TerminalNode
	SUB() antlr.TerminalNode
	NOT() antlr.TerminalNode

	// IsUnary_opContext differentiates from other interfaces.
	IsUnary_opContext()
}

type Unary_opContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnary_opContext() *Unary_opContext {
	var p = new(Unary_opContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_unary_op
	return p
}

func InitEmptyUnary_opContext(p *Unary_opContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_unary_op
}

func (*Unary_opContext) IsUnary_opContext() {}

func NewUnary_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unary_opContext {
	var p = new(Unary_opContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_unary_op

	return p
}

func (s *Unary_opContext) GetParser() antlr.Parser { return s.parser }

func (s *Unary_opContext) ADD() antlr.TerminalNode {
	return s.GetToken(NeedleParserADD, 0)
}

func (s *Unary_opContext) SUB() antlr.TerminalNode {
	return s.GetToken(NeedleParserSUB, 0)
}

func (s *Unary_opContext) NOT() antlr.TerminalNode {
	return s.GetToken(NeedleParserNOT, 0)
}

func (s *Unary_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unary_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unary_opContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterUnary_op(s)
	}
}

func (s *Unary_opContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitUnary_op(s)
	}
}

func (p *NeedleParser) Unary_op() (localctx IUnary_opContext) {
	localctx = NewUnary_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, NeedleParserRULE_unary_op)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(474)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&53248) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAdd_opContext is an interface to support dynamic dispatch.
type IAdd_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ADD() antlr.TerminalNode
	SUB() antlr.TerminalNode
	BIT_OR() antlr.TerminalNode
	BIT_XOR() antlr.TerminalNode

	// IsAdd_opContext differentiates from other interfaces.
	IsAdd_opContext()
}

type Add_opContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdd_opContext() *Add_opContext {
	var p = new(Add_opContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_add_op
	return p
}

func InitEmptyAdd_opContext(p *Add_opContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_add_op
}

func (*Add_opContext) IsAdd_opContext() {}

func NewAdd_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Add_opContext {
	var p = new(Add_opContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_add_op

	return p
}

func (s *Add_opContext) GetParser() antlr.Parser { return s.parser }

func (s *Add_opContext) ADD() antlr.TerminalNode {
	return s.GetToken(NeedleParserADD, 0)
}

func (s *Add_opContext) SUB() antlr.TerminalNode {
	return s.GetToken(NeedleParserSUB, 0)
}

func (s *Add_opContext) BIT_OR() antlr.TerminalNode {
	return s.GetToken(NeedleParserBIT_OR, 0)
}

func (s *Add_opContext) BIT_XOR() antlr.TerminalNode {
	return s.GetToken(NeedleParserBIT_XOR, 0)
}

func (s *Add_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Add_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Add_opContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterAdd_op(s)
	}
}

func (s *Add_opContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitAdd_op(s)
	}
}

func (p *NeedleParser) Add_op() (localctx IAdd_opContext) {
	localctx = NewAdd_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, NeedleParserRULE_add_op)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(476)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&201375744) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogical_opContext is an interface to support dynamic dispatch.
type ILogical_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsLogical_opContext differentiates from other interfaces.
	IsLogical_opContext()
}

type Logical_opContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogical_opContext() *Logical_opContext {
	var p = new(Logical_opContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_logical_op
	return p
}

func InitEmptyLogical_opContext(p *Logical_opContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_logical_op
}

func (*Logical_opContext) IsLogical_opContext() {}

func NewLogical_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Logical_opContext {
	var p = new(Logical_opContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_logical_op

	return p
}

func (s *Logical_opContext) GetParser() antlr.Parser { return s.parser }

func (s *Logical_opContext) AND() antlr.TerminalNode {
	return s.GetToken(NeedleParserAND, 0)
}

func (s *Logical_opContext) OR() antlr.TerminalNode {
	return s.GetToken(NeedleParserOR, 0)
}

func (s *Logical_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Logical_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Logical_opContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterLogical_op(s)
	}
}

func (s *Logical_opContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitLogical_op(s)
	}
}

func (p *NeedleParser) Logical_op() (localctx ILogical_opContext) {
	localctx = NewLogical_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, NeedleParserRULE_logical_op)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(478)
		_la = p.GetTokenStream().LA(1)

		if !(_la == NeedleParserAND || _la == NeedleParserOR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRel_opContext is an interface to support dynamic dispatch.
type IRel_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LESS() antlr.TerminalNode
	GREATER() antlr.TerminalNode
	LESS_EQ() antlr.TerminalNode
	GR_EQ() antlr.TerminalNode
	EQ_EQ() antlr.TerminalNode
	NOT_EQ() antlr.TerminalNode

	// IsRel_opContext differentiates from other interfaces.
	IsRel_opContext()
}

type Rel_opContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRel_opContext() *Rel_opContext {
	var p = new(Rel_opContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_rel_op
	return p
}

func InitEmptyRel_opContext(p *Rel_opContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_rel_op
}

func (*Rel_opContext) IsRel_opContext() {}

func NewRel_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rel_opContext {
	var p = new(Rel_opContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_rel_op

	return p
}

func (s *Rel_opContext) GetParser() antlr.Parser { return s.parser }

func (s *Rel_opContext) LESS() antlr.TerminalNode {
	return s.GetToken(NeedleParserLESS, 0)
}

func (s *Rel_opContext) GREATER() antlr.TerminalNode {
	return s.GetToken(NeedleParserGREATER, 0)
}

func (s *Rel_opContext) LESS_EQ() antlr.TerminalNode {
	return s.GetToken(NeedleParserLESS_EQ, 0)
}

func (s *Rel_opContext) GR_EQ() antlr.TerminalNode {
	return s.GetToken(NeedleParserGR_EQ, 0)
}

func (s *Rel_opContext) EQ_EQ() antlr.TerminalNode {
	return s.GetToken(NeedleParserEQ_EQ, 0)
}

func (s *Rel_opContext) NOT_EQ() antlr.TerminalNode {
	return s.GetToken(NeedleParserNOT_EQ, 0)
}

func (s *Rel_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rel_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rel_opContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterRel_op(s)
	}
}

func (s *Rel_opContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitRel_op(s)
	}
}

func (p *NeedleParser) Rel_op() (localctx IRel_opContext) {
	localctx = NewRel_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, NeedleParserRULE_rel_op)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(480)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15597568) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssign_opContext is an interface to support dynamic dispatch.
type IAssign_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQ() antlr.TerminalNode
	ADD_EQ() antlr.TerminalNode
	SUB_EQ() antlr.TerminalNode
	MUL_EQ() antlr.TerminalNode
	DIV_EQ() antlr.TerminalNode
	MOD_EQ() antlr.TerminalNode
	LSHIFT_EQ() antlr.TerminalNode
	RSHIFT_EQ() antlr.TerminalNode
	BIT_AND_EQ() antlr.TerminalNode
	BIT_OR_EQ() antlr.TerminalNode
	BIT_XOR_EQ() antlr.TerminalNode

	// IsAssign_opContext differentiates from other interfaces.
	IsAssign_opContext()
}

type Assign_opContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssign_opContext() *Assign_opContext {
	var p = new(Assign_opContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_assign_op
	return p
}

func InitEmptyAssign_opContext(p *Assign_opContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_assign_op
}

func (*Assign_opContext) IsAssign_opContext() {}

func NewAssign_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_opContext {
	var p = new(Assign_opContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_assign_op

	return p
}

func (s *Assign_opContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_opContext) EQ() antlr.TerminalNode {
	return s.GetToken(NeedleParserEQ, 0)
}

func (s *Assign_opContext) ADD_EQ() antlr.TerminalNode {
	return s.GetToken(NeedleParserADD_EQ, 0)
}

func (s *Assign_opContext) SUB_EQ() antlr.TerminalNode {
	return s.GetToken(NeedleParserSUB_EQ, 0)
}

func (s *Assign_opContext) MUL_EQ() antlr.TerminalNode {
	return s.GetToken(NeedleParserMUL_EQ, 0)
}

func (s *Assign_opContext) DIV_EQ() antlr.TerminalNode {
	return s.GetToken(NeedleParserDIV_EQ, 0)
}

func (s *Assign_opContext) MOD_EQ() antlr.TerminalNode {
	return s.GetToken(NeedleParserMOD_EQ, 0)
}

func (s *Assign_opContext) LSHIFT_EQ() antlr.TerminalNode {
	return s.GetToken(NeedleParserLSHIFT_EQ, 0)
}

func (s *Assign_opContext) RSHIFT_EQ() antlr.TerminalNode {
	return s.GetToken(NeedleParserRSHIFT_EQ, 0)
}

func (s *Assign_opContext) BIT_AND_EQ() antlr.TerminalNode {
	return s.GetToken(NeedleParserBIT_AND_EQ, 0)
}

func (s *Assign_opContext) BIT_OR_EQ() antlr.TerminalNode {
	return s.GetToken(NeedleParserBIT_OR_EQ, 0)
}

func (s *Assign_opContext) BIT_XOR_EQ() antlr.TerminalNode {
	return s.GetToken(NeedleParserBIT_XOR_EQ, 0)
}

func (s *Assign_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assign_opContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterAssign_op(s)
	}
}

func (s *Assign_opContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitAssign_op(s)
	}
}

func (p *NeedleParser) Assign_op() (localctx IAssign_opContext) {
	localctx = NewAssign_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, NeedleParserRULE_assign_op)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(482)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2196875771968) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierFullContext is an interface to support dynamic dispatch.
type IIdentifierFullContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	DollarIdentifier() antlr.TerminalNode
	AtIdentifier() antlr.TerminalNode

	// IsIdentifierFullContext differentiates from other interfaces.
	IsIdentifierFullContext()
}

type IdentifierFullContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierFullContext() *IdentifierFullContext {
	var p = new(IdentifierFullContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_identifierFull
	return p
}

func InitEmptyIdentifierFullContext(p *IdentifierFullContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_identifierFull
}

func (*IdentifierFullContext) IsIdentifierFullContext() {}

func NewIdentifierFullContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierFullContext {
	var p = new(IdentifierFullContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_identifierFull

	return p
}

func (s *IdentifierFullContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierFullContext) Identifier() antlr.TerminalNode {
	return s.GetToken(NeedleParserIdentifier, 0)
}

func (s *IdentifierFullContext) DollarIdentifier() antlr.TerminalNode {
	return s.GetToken(NeedleParserDollarIdentifier, 0)
}

func (s *IdentifierFullContext) AtIdentifier() antlr.TerminalNode {
	return s.GetToken(NeedleParserAtIdentifier, 0)
}

func (s *IdentifierFullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierFullContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierFullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterIdentifierFull(s)
	}
}

func (s *IdentifierFullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitIdentifierFull(s)
	}
}

func (p *NeedleParser) IdentifierFull() (localctx IIdentifierFullContext) {
	localctx = NewIdentifierFullContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, NeedleParserRULE_identifierFull)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(484)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-74)) & ^0x3f) == 0 && ((int64(1)<<(_la-74))&7) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierVarContext is an interface to support dynamic dispatch.
type IIdentifierVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	DollarIdentifier() antlr.TerminalNode

	// IsIdentifierVarContext differentiates from other interfaces.
	IsIdentifierVarContext()
}

type IdentifierVarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierVarContext() *IdentifierVarContext {
	var p = new(IdentifierVarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_identifierVar
	return p
}

func InitEmptyIdentifierVarContext(p *IdentifierVarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_identifierVar
}

func (*IdentifierVarContext) IsIdentifierVarContext() {}

func NewIdentifierVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierVarContext {
	var p = new(IdentifierVarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_identifierVar

	return p
}

func (s *IdentifierVarContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierVarContext) Identifier() antlr.TerminalNode {
	return s.GetToken(NeedleParserIdentifier, 0)
}

func (s *IdentifierVarContext) DollarIdentifier() antlr.TerminalNode {
	return s.GetToken(NeedleParserDollarIdentifier, 0)
}

func (s *IdentifierVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierVarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterIdentifierVar(s)
	}
}

func (s *IdentifierVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitIdentifierVar(s)
	}
}

func (p *NeedleParser) IdentifierVar() (localctx IIdentifierVarContext) {
	localctx = NewIdentifierVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, NeedleParserRULE_identifierVar)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(486)
		_la = p.GetTokenStream().LA(1)

		if !(_la == NeedleParserIdentifier || _la == NeedleParserDollarIdentifier) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierListContext is an interface to support dynamic dispatch.
type IIdentifierListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsIdentifierListContext differentiates from other interfaces.
	IsIdentifierListContext()
}

type IdentifierListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierListContext() *IdentifierListContext {
	var p = new(IdentifierListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_identifierList
	return p
}

func InitEmptyIdentifierListContext(p *IdentifierListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_identifierList
}

func (*IdentifierListContext) IsIdentifierListContext() {}

func NewIdentifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierListContext {
	var p = new(IdentifierListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_identifierList

	return p
}

func (s *IdentifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierListContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(NeedleParserIdentifier)
}

func (s *IdentifierListContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(NeedleParserIdentifier, i)
}

func (s *IdentifierListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(NeedleParserCOMMA)
}

func (s *IdentifierListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(NeedleParserCOMMA, i)
}

func (s *IdentifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterIdentifierList(s)
	}
}

func (s *IdentifierListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitIdentifierList(s)
	}
}

func (p *NeedleParser) IdentifierList() (localctx IIdentifierListContext) {
	localctx = NewIdentifierListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, NeedleParserRULE_identifierList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(488)
		p.Match(NeedleParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(495)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == NeedleParserCOMMA || _la == NeedleParserIdentifier {
		p.SetState(490)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == NeedleParserCOMMA {
			{
				p.SetState(489)
				p.Match(NeedleParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(492)
			p.Match(NeedleParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(497)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringLiteralContext is an interface to support dynamic dispatch.
type IStringLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	InterpretedStringLiteral() antlr.TerminalNode
	RawStringLiteral() antlr.TerminalNode

	// IsStringLiteralContext differentiates from other interfaces.
	IsStringLiteralContext()
}

type StringLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLiteralContext() *StringLiteralContext {
	var p = new(StringLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_stringLiteral
	return p
}

func InitEmptyStringLiteralContext(p *StringLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_stringLiteral
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) InterpretedStringLiteral() antlr.TerminalNode {
	return s.GetToken(NeedleParserInterpretedStringLiteral, 0)
}

func (s *StringLiteralContext) RawStringLiteral() antlr.TerminalNode {
	return s.GetToken(NeedleParserRawStringLiteral, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (p *NeedleParser) StringLiteral() (localctx IStringLiteralContext) {
	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, NeedleParserRULE_stringLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(498)
		_la = p.GetTokenStream().LA(1)

		if !(_la == NeedleParserInterpretedStringLiteral || _la == NeedleParserRawStringLiteral) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumberLiteralContext is an interface to support dynamic dispatch.
type INumberLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DecimalLiteral() antlr.TerminalNode
	BinaryLiteral() antlr.TerminalNode
	OctalLiteral() antlr.TerminalNode
	HexLiteral() antlr.TerminalNode
	FloatLiteral() antlr.TerminalNode

	// IsNumberLiteralContext differentiates from other interfaces.
	IsNumberLiteralContext()
}

type NumberLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberLiteralContext() *NumberLiteralContext {
	var p = new(NumberLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_numberLiteral
	return p
}

func InitEmptyNumberLiteralContext(p *NumberLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_numberLiteral
}

func (*NumberLiteralContext) IsNumberLiteralContext() {}

func NewNumberLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberLiteralContext {
	var p = new(NumberLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_numberLiteral

	return p
}

func (s *NumberLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberLiteralContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(NeedleParserDecimalLiteral, 0)
}

func (s *NumberLiteralContext) BinaryLiteral() antlr.TerminalNode {
	return s.GetToken(NeedleParserBinaryLiteral, 0)
}

func (s *NumberLiteralContext) OctalLiteral() antlr.TerminalNode {
	return s.GetToken(NeedleParserOctalLiteral, 0)
}

func (s *NumberLiteralContext) HexLiteral() antlr.TerminalNode {
	return s.GetToken(NeedleParserHexLiteral, 0)
}

func (s *NumberLiteralContext) FloatLiteral() antlr.TerminalNode {
	return s.GetToken(NeedleParserFloatLiteral, 0)
}

func (s *NumberLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterNumberLiteral(s)
	}
}

func (s *NumberLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitNumberLiteral(s)
	}
}

func (p *NeedleParser) NumberLiteral() (localctx INumberLiteralContext) {
	localctx = NewNumberLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, NeedleParserRULE_numberLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(500)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-79)) & ^0x3f) == 0 && ((int64(1)<<(_la-79))&31) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBooleanLiteralContext is an interface to support dynamic dispatch.
type IBooleanLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode

	// IsBooleanLiteralContext differentiates from other interfaces.
	IsBooleanLiteralContext()
}

type BooleanLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_booleanLiteral
	return p
}

func InitEmptyBooleanLiteralContext(p *BooleanLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_booleanLiteral
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(NeedleParserTRUE, 0)
}

func (s *BooleanLiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(NeedleParserFALSE, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (p *NeedleParser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, NeedleParserRULE_booleanLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(502)
		_la = p.GetTokenStream().LA(1)

		if !(_la == NeedleParserTRUE || _la == NeedleParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEosContext is an interface to support dynamic dispatch.
type IEosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SEMI() antlr.TerminalNode
	EOS() antlr.TerminalNode
	EOF() antlr.TerminalNode

	// IsEosContext differentiates from other interfaces.
	IsEosContext()
}

type EosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEosContext() *EosContext {
	var p = new(EosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_eos
	return p
}

func InitEmptyEosContext(p *EosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_eos
}

func (*EosContext) IsEosContext() {}

func NewEosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EosContext {
	var p = new(EosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_eos

	return p
}

func (s *EosContext) GetParser() antlr.Parser { return s.parser }

func (s *EosContext) SEMI() antlr.TerminalNode {
	return s.GetToken(NeedleParserSEMI, 0)
}

func (s *EosContext) EOS() antlr.TerminalNode {
	return s.GetToken(NeedleParserEOS, 0)
}

func (s *EosContext) EOF() antlr.TerminalNode {
	return s.GetToken(NeedleParserEOF, 0)
}

func (s *EosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterEos(s)
	}
}

func (s *EosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitEos(s)
	}
}

func (p *NeedleParser) Eos() (localctx IEosContext) {
	localctx = NewEosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, NeedleParserRULE_eos)
	p.SetState(509)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 52, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(504)
			p.Match(NeedleParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(505)
			p.Match(NeedleParserEOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(507)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 51, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(506)
				p.Match(NeedleParserEOF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *NeedleParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 37:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	case 38:
		var t *PrimaryExprContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryExprContext)
		}
		return p.PrimaryExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *NeedleParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *NeedleParser) PrimaryExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
