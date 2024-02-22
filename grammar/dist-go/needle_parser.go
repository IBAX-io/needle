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
		"settingsDef", "settingsValue", "funcDef", "innerFuncDef", "defaultFuncDef",
		"funcDescriptor", "funcSignature", "parameterList", "parameter", "returnParameters",
		"block", "statList", "stat", "varDef", "ifStat", "returnStat", "continueStat",
		"breakStat", "whileStat", "errorStat", "typeName", "sliceStat", "indexNumber",
		"arrayStat", "arrayList", "arrayValue", "indexStat", "objectStat", "pairList",
		"pair", "pairValue", "arguments", "argumentsList", "simpleStat", "incDecStat",
		"exprStat", "assignMapArrStat", "initMapArrStat", "assignment", "primaryExpr",
		"operand", "exprList", "expr", "incDec_op", "mul_op", "unary_op", "add_op",
		"logical_op", "rel_op", "assign_op", "identifierFull", "identifierVar",
		"identifierList", "stringLiteral", "numberLiteral", "booleanLiteral",
		"eos",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 99, 570, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
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
		2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 1, 0, 1, 0, 3,
		0, 127, 8, 0, 1, 0, 1, 0, 5, 0, 131, 8, 0, 10, 0, 12, 0, 134, 9, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 144, 8, 1, 10, 1, 12,
		1, 147, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 3, 2, 154, 8, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 5, 3, 161, 8, 3, 10, 3, 12, 3, 164, 9, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 4, 3, 4, 171, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 5, 5, 180, 8, 5, 10, 5, 12, 5, 183, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 6, 3, 6, 190, 8, 6, 1, 7, 1, 7, 3, 7, 194, 8, 7, 1, 7, 1, 7, 1, 7, 3,
		7, 199, 8, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 3, 9, 207, 8, 9, 1, 9,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 3, 11, 216, 8, 11, 1, 11, 3, 11,
		219, 8, 11, 1, 12, 1, 12, 1, 12, 3, 12, 224, 8, 12, 1, 12, 5, 12, 227,
		8, 12, 10, 12, 12, 12, 230, 9, 12, 1, 12, 3, 12, 233, 8, 12, 3, 12, 235,
		8, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 3, 13, 242, 8, 13, 1, 13, 1,
		13, 1, 13, 5, 13, 247, 8, 13, 10, 13, 12, 13, 250, 9, 13, 1, 14, 1, 14,
		3, 14, 254, 8, 14, 1, 14, 5, 14, 257, 8, 14, 10, 14, 12, 14, 260, 9, 14,
		1, 15, 1, 15, 3, 15, 264, 8, 15, 1, 15, 1, 15, 1, 16, 3, 16, 269, 8, 16,
		1, 16, 3, 16, 272, 8, 16, 3, 16, 274, 8, 16, 1, 16, 1, 16, 1, 16, 4, 16,
		279, 8, 16, 11, 16, 12, 16, 280, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 3, 17, 292, 8, 17, 1, 18, 1, 18, 1, 18, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 303, 8, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 312, 8, 19, 1, 19, 1, 19, 5, 19,
		316, 8, 19, 10, 19, 12, 19, 319, 9, 19, 1, 19, 1, 19, 3, 19, 323, 8, 19,
		1, 20, 1, 20, 3, 20, 327, 8, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 3, 26,
		344, 8, 26, 1, 26, 1, 26, 3, 26, 348, 8, 26, 1, 26, 1, 26, 1, 27, 1, 27,
		3, 27, 354, 8, 27, 1, 28, 1, 28, 3, 28, 358, 8, 28, 1, 28, 1, 28, 1, 29,
		1, 29, 1, 29, 5, 29, 365, 8, 29, 10, 29, 12, 29, 368, 9, 29, 1, 29, 1,
		29, 1, 30, 1, 30, 1, 30, 3, 30, 375, 8, 30, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 32, 1, 32, 3, 32, 383, 8, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 5,
		33, 390, 8, 33, 10, 33, 12, 33, 393, 9, 33, 1, 33, 3, 33, 396, 8, 33, 1,
		33, 1, 33, 1, 34, 1, 34, 3, 34, 402, 8, 34, 1, 34, 1, 34, 4, 34, 406, 8,
		34, 11, 34, 12, 34, 407, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		3, 35, 417, 8, 35, 1, 36, 1, 36, 3, 36, 421, 8, 36, 1, 36, 1, 36, 1, 37,
		1, 37, 3, 37, 427, 8, 37, 1, 37, 1, 37, 1, 37, 3, 37, 432, 8, 37, 5, 37,
		434, 8, 37, 10, 37, 12, 37, 437, 9, 37, 1, 38, 1, 38, 1, 38, 1, 38, 3,
		38, 443, 8, 38, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41,
		1, 41, 1, 42, 1, 42, 3, 42, 456, 8, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1,
		44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 3, 44, 468, 8, 44, 1, 44, 1, 44,
		1, 44, 3, 44, 473, 8, 44, 5, 44, 475, 8, 44, 10, 44, 12, 44, 478, 9, 44,
		1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 3, 45, 489,
		8, 45, 1, 46, 1, 46, 1, 46, 5, 46, 494, 8, 46, 10, 46, 12, 46, 497, 9,
		46, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 3, 47, 506, 8, 47,
		1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1,
		47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 5, 47, 524, 8, 47, 10, 47, 12, 47,
		527, 9, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1,
		52, 1, 52, 1, 53, 1, 53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 56, 1, 56, 1, 57,
		1, 57, 3, 57, 549, 8, 57, 1, 57, 5, 57, 552, 8, 57, 10, 57, 12, 57, 555,
		9, 57, 1, 58, 1, 58, 1, 59, 1, 59, 1, 60, 1, 60, 1, 61, 1, 61, 1, 61, 3,
		61, 566, 8, 61, 3, 61, 568, 8, 61, 1, 61, 1, 317, 2, 88, 94, 62, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
		42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76,
		78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110,
		112, 114, 116, 118, 120, 122, 0, 15, 1, 0, 60, 61, 2, 0, 57, 58, 63, 63,
		1, 0, 64, 73, 1, 0, 41, 42, 4, 0, 13, 13, 16, 16, 25, 25, 28, 30, 2, 0,
		12, 12, 14, 15, 2, 0, 14, 15, 26, 27, 2, 0, 20, 20, 24, 24, 2, 0, 17, 19,
		21, 23, 2, 0, 6, 6, 31, 40, 1, 0, 74, 76, 1, 0, 74, 75, 1, 0, 77, 78, 1,
		0, 79, 83, 1, 0, 50, 51, 594, 0, 132, 1, 0, 0, 0, 2, 137, 1, 0, 0, 0, 4,
		153, 1, 0, 0, 0, 6, 155, 1, 0, 0, 0, 8, 167, 1, 0, 0, 0, 10, 172, 1, 0,
		0, 0, 12, 189, 1, 0, 0, 0, 14, 193, 1, 0, 0, 0, 16, 202, 1, 0, 0, 0, 18,
		206, 1, 0, 0, 0, 20, 211, 1, 0, 0, 0, 22, 215, 1, 0, 0, 0, 24, 220, 1,
		0, 0, 0, 26, 238, 1, 0, 0, 0, 28, 251, 1, 0, 0, 0, 30, 261, 1, 0, 0, 0,
		32, 278, 1, 0, 0, 0, 34, 291, 1, 0, 0, 0, 36, 293, 1, 0, 0, 0, 38, 296,
		1, 0, 0, 0, 40, 324, 1, 0, 0, 0, 42, 328, 1, 0, 0, 0, 44, 330, 1, 0, 0,
		0, 46, 332, 1, 0, 0, 0, 48, 336, 1, 0, 0, 0, 50, 339, 1, 0, 0, 0, 52, 341,
		1, 0, 0, 0, 54, 353, 1, 0, 0, 0, 56, 355, 1, 0, 0, 0, 58, 361, 1, 0, 0,
		0, 60, 374, 1, 0, 0, 0, 62, 376, 1, 0, 0, 0, 64, 380, 1, 0, 0, 0, 66, 386,
		1, 0, 0, 0, 68, 401, 1, 0, 0, 0, 70, 416, 1, 0, 0, 0, 72, 418, 1, 0, 0,
		0, 74, 426, 1, 0, 0, 0, 76, 442, 1, 0, 0, 0, 78, 444, 1, 0, 0, 0, 80, 447,
		1, 0, 0, 0, 82, 449, 1, 0, 0, 0, 84, 455, 1, 0, 0, 0, 86, 457, 1, 0, 0,
		0, 88, 461, 1, 0, 0, 0, 90, 488, 1, 0, 0, 0, 92, 490, 1, 0, 0, 0, 94, 505,
		1, 0, 0, 0, 96, 528, 1, 0, 0, 0, 98, 530, 1, 0, 0, 0, 100, 532, 1, 0, 0,
		0, 102, 534, 1, 0, 0, 0, 104, 536, 1, 0, 0, 0, 106, 538, 1, 0, 0, 0, 108,
		540, 1, 0, 0, 0, 110, 542, 1, 0, 0, 0, 112, 544, 1, 0, 0, 0, 114, 546,
		1, 0, 0, 0, 116, 556, 1, 0, 0, 0, 118, 558, 1, 0, 0, 0, 120, 560, 1, 0,
		0, 0, 122, 567, 1, 0, 0, 0, 124, 127, 3, 2, 1, 0, 125, 127, 3, 14, 7, 0,
		126, 124, 1, 0, 0, 0, 126, 125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128,
		129, 3, 122, 61, 0, 129, 131, 1, 0, 0, 0, 130, 126, 1, 0, 0, 0, 131, 134,
		1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 135, 1, 0,
		0, 0, 134, 132, 1, 0, 0, 0, 135, 136, 5, 0, 0, 1, 136, 1, 1, 0, 0, 0, 137,
		138, 5, 43, 0, 0, 138, 139, 5, 74, 0, 0, 139, 145, 5, 9, 0, 0, 140, 141,
		3, 4, 2, 0, 141, 142, 3, 122, 61, 0, 142, 144, 1, 0, 0, 0, 143, 140, 1,
		0, 0, 0, 144, 147, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0,
		0, 146, 148, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 148, 149, 5, 10, 0, 0, 149,
		3, 1, 0, 0, 0, 150, 154, 3, 6, 3, 0, 151, 154, 3, 10, 5, 0, 152, 154, 3,
		14, 7, 0, 153, 150, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 153, 152, 1, 0, 0,
		0, 154, 5, 1, 0, 0, 0, 155, 156, 5, 53, 0, 0, 156, 162, 5, 9, 0, 0, 157,
		158, 3, 8, 4, 0, 158, 159, 3, 122, 61, 0, 159, 161, 1, 0, 0, 0, 160, 157,
		1, 0, 0, 0, 161, 164, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0,
		0, 0, 163, 165, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 165, 166, 5, 10, 0, 0,
		166, 7, 1, 0, 0, 0, 167, 168, 5, 74, 0, 0, 168, 170, 3, 50, 25, 0, 169,
		171, 3, 116, 58, 0, 170, 169, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 9,
		1, 0, 0, 0, 172, 173, 5, 54, 0, 0, 173, 181, 5, 9, 0, 0, 174, 175, 5, 74,
		0, 0, 175, 176, 5, 6, 0, 0, 176, 177, 3, 12, 6, 0, 177, 178, 3, 122, 61,
		0, 178, 180, 1, 0, 0, 0, 179, 174, 1, 0, 0, 0, 180, 183, 1, 0, 0, 0, 181,
		179, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 184, 1, 0, 0, 0, 183, 181,
		1, 0, 0, 0, 184, 185, 5, 10, 0, 0, 185, 11, 1, 0, 0, 0, 186, 190, 3, 118,
		59, 0, 187, 190, 3, 120, 60, 0, 188, 190, 3, 116, 58, 0, 189, 186, 1, 0,
		0, 0, 189, 187, 1, 0, 0, 0, 189, 188, 1, 0, 0, 0, 190, 13, 1, 0, 0, 0,
		191, 194, 3, 16, 8, 0, 192, 194, 3, 18, 9, 0, 193, 191, 1, 0, 0, 0, 193,
		192, 1, 0, 0, 0, 194, 198, 1, 0, 0, 0, 195, 196, 5, 4, 0, 0, 196, 197,
		5, 74, 0, 0, 197, 199, 3, 22, 11, 0, 198, 195, 1, 0, 0, 0, 198, 199, 1,
		0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 201, 3, 30, 15, 0, 201, 15, 1, 0, 0,
		0, 202, 203, 3, 20, 10, 0, 203, 204, 3, 22, 11, 0, 204, 17, 1, 0, 0, 0,
		205, 207, 5, 44, 0, 0, 206, 205, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207,
		208, 1, 0, 0, 0, 208, 209, 7, 0, 0, 0, 209, 210, 3, 22, 11, 0, 210, 19,
		1, 0, 0, 0, 211, 212, 5, 44, 0, 0, 212, 213, 5, 74, 0, 0, 213, 21, 1, 0,
		0, 0, 214, 216, 3, 24, 12, 0, 215, 214, 1, 0, 0, 0, 215, 216, 1, 0, 0,
		0, 216, 218, 1, 0, 0, 0, 217, 219, 3, 28, 14, 0, 218, 217, 1, 0, 0, 0,
		218, 219, 1, 0, 0, 0, 219, 23, 1, 0, 0, 0, 220, 234, 5, 1, 0, 0, 221, 228,
		3, 26, 13, 0, 222, 224, 5, 3, 0, 0, 223, 222, 1, 0, 0, 0, 223, 224, 1,
		0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 227, 3, 26, 13, 0, 226, 223, 1, 0,
		0, 0, 227, 230, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0,
		229, 232, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 231, 233, 5, 3, 0, 0, 232,
		231, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 235, 1, 0, 0, 0, 234, 221,
		1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236, 237, 5, 2,
		0, 0, 237, 25, 1, 0, 0, 0, 238, 239, 3, 114, 57, 0, 239, 248, 3, 50, 25,
		0, 240, 242, 5, 3, 0, 0, 241, 240, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242,
		243, 1, 0, 0, 0, 243, 244, 3, 114, 57, 0, 244, 245, 3, 50, 25, 0, 245,
		247, 1, 0, 0, 0, 246, 241, 1, 0, 0, 0, 247, 250, 1, 0, 0, 0, 248, 246,
		1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 27, 1, 0, 0, 0, 250, 248, 1, 0,
		0, 0, 251, 258, 3, 50, 25, 0, 252, 254, 5, 3, 0, 0, 253, 252, 1, 0, 0,
		0, 253, 254, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 257, 3, 50, 25, 0,
		256, 253, 1, 0, 0, 0, 257, 260, 1, 0, 0, 0, 258, 256, 1, 0, 0, 0, 258,
		259, 1, 0, 0, 0, 259, 29, 1, 0, 0, 0, 260, 258, 1, 0, 0, 0, 261, 263, 5,
		9, 0, 0, 262, 264, 3, 32, 16, 0, 263, 262, 1, 0, 0, 0, 263, 264, 1, 0,
		0, 0, 264, 265, 1, 0, 0, 0, 265, 266, 5, 10, 0, 0, 266, 31, 1, 0, 0, 0,
		267, 269, 5, 11, 0, 0, 268, 267, 1, 0, 0, 0, 268, 269, 1, 0, 0, 0, 269,
		274, 1, 0, 0, 0, 270, 272, 5, 98, 0, 0, 271, 270, 1, 0, 0, 0, 271, 272,
		1, 0, 0, 0, 272, 274, 1, 0, 0, 0, 273, 268, 1, 0, 0, 0, 273, 271, 1, 0,
		0, 0, 274, 275, 1, 0, 0, 0, 275, 276, 3, 34, 17, 0, 276, 277, 3, 122, 61,
		0, 277, 279, 1, 0, 0, 0, 278, 273, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280,
		278, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281, 33, 1, 0, 0, 0, 282, 292, 3,
		30, 15, 0, 283, 292, 3, 76, 38, 0, 284, 292, 3, 36, 18, 0, 285, 292, 3,
		38, 19, 0, 286, 292, 3, 46, 23, 0, 287, 292, 3, 42, 21, 0, 288, 292, 3,
		44, 22, 0, 289, 292, 3, 40, 20, 0, 290, 292, 3, 48, 24, 0, 291, 282, 1,
		0, 0, 0, 291, 283, 1, 0, 0, 0, 291, 284, 1, 0, 0, 0, 291, 285, 1, 0, 0,
		0, 291, 286, 1, 0, 0, 0, 291, 287, 1, 0, 0, 0, 291, 288, 1, 0, 0, 0, 291,
		289, 1, 0, 0, 0, 291, 290, 1, 0, 0, 0, 292, 35, 1, 0, 0, 0, 293, 294, 5,
		52, 0, 0, 294, 295, 3, 26, 13, 0, 295, 37, 1, 0, 0, 0, 296, 302, 5, 46,
		0, 0, 297, 298, 5, 1, 0, 0, 298, 299, 3, 94, 47, 0, 299, 300, 5, 2, 0,
		0, 300, 303, 1, 0, 0, 0, 301, 303, 3, 94, 47, 0, 302, 297, 1, 0, 0, 0,
		302, 301, 1, 0, 0, 0, 303, 304, 1, 0, 0, 0, 304, 317, 3, 30, 15, 0, 305,
		311, 5, 47, 0, 0, 306, 307, 5, 1, 0, 0, 307, 308, 3, 94, 47, 0, 308, 309,
		5, 2, 0, 0, 309, 312, 1, 0, 0, 0, 310, 312, 3, 94, 47, 0, 311, 306, 1,
		0, 0, 0, 311, 310, 1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 314, 3, 30, 15,
		0, 314, 316, 1, 0, 0, 0, 315, 305, 1, 0, 0, 0, 316, 319, 1, 0, 0, 0, 317,
		318, 1, 0, 0, 0, 317, 315, 1, 0, 0, 0, 318, 322, 1, 0, 0, 0, 319, 317,
		1, 0, 0, 0, 320, 321, 5, 48, 0, 0, 321, 323, 3, 30, 15, 0, 322, 320, 1,
		0, 0, 0, 322, 323, 1, 0, 0, 0, 323, 39, 1, 0, 0, 0, 324, 326, 5, 45, 0,
		0, 325, 327, 3, 94, 47, 0, 326, 325, 1, 0, 0, 0, 326, 327, 1, 0, 0, 0,
		327, 41, 1, 0, 0, 0, 328, 329, 5, 56, 0, 0, 329, 43, 1, 0, 0, 0, 330, 331,
		5, 55, 0, 0, 331, 45, 1, 0, 0, 0, 332, 333, 5, 49, 0, 0, 333, 334, 3, 94,
		47, 0, 334, 335, 3, 30, 15, 0, 335, 47, 1, 0, 0, 0, 336, 337, 7, 1, 0,
		0, 337, 338, 3, 94, 47, 0, 338, 49, 1, 0, 0, 0, 339, 340, 7, 2, 0, 0, 340,
		51, 1, 0, 0, 0, 341, 343, 5, 7, 0, 0, 342, 344, 3, 54, 27, 0, 343, 342,
		1, 0, 0, 0, 343, 344, 1, 0, 0, 0, 344, 345, 1, 0, 0, 0, 345, 347, 5, 5,
		0, 0, 346, 348, 3, 54, 27, 0, 347, 346, 1, 0, 0, 0, 347, 348, 1, 0, 0,
		0, 348, 349, 1, 0, 0, 0, 349, 350, 5, 8, 0, 0, 350, 53, 1, 0, 0, 0, 351,
		354, 3, 118, 59, 0, 352, 354, 3, 112, 56, 0, 353, 351, 1, 0, 0, 0, 353,
		352, 1, 0, 0, 0, 354, 55, 1, 0, 0, 0, 355, 357, 5, 7, 0, 0, 356, 358, 3,
		58, 29, 0, 357, 356, 1, 0, 0, 0, 357, 358, 1, 0, 0, 0, 358, 359, 1, 0,
		0, 0, 359, 360, 5, 8, 0, 0, 360, 57, 1, 0, 0, 0, 361, 366, 3, 60, 30, 0,
		362, 363, 5, 3, 0, 0, 363, 365, 3, 60, 30, 0, 364, 362, 1, 0, 0, 0, 365,
		368, 1, 0, 0, 0, 366, 364, 1, 0, 0, 0, 366, 367, 1, 0, 0, 0, 367, 369,
		1, 0, 0, 0, 368, 366, 1, 0, 0, 0, 369, 370, 3, 122, 61, 0, 370, 59, 1,
		0, 0, 0, 371, 375, 3, 56, 28, 0, 372, 375, 3, 94, 47, 0, 373, 375, 3, 64,
		32, 0, 374, 371, 1, 0, 0, 0, 374, 372, 1, 0, 0, 0, 374, 373, 1, 0, 0, 0,
		375, 61, 1, 0, 0, 0, 376, 377, 5, 7, 0, 0, 377, 378, 3, 94, 47, 0, 378,
		379, 5, 8, 0, 0, 379, 63, 1, 0, 0, 0, 380, 382, 5, 9, 0, 0, 381, 383, 3,
		66, 33, 0, 382, 381, 1, 0, 0, 0, 382, 383, 1, 0, 0, 0, 383, 384, 1, 0,
		0, 0, 384, 385, 5, 10, 0, 0, 385, 65, 1, 0, 0, 0, 386, 391, 3, 68, 34,
		0, 387, 388, 5, 3, 0, 0, 388, 390, 3, 68, 34, 0, 389, 387, 1, 0, 0, 0,
		390, 393, 1, 0, 0, 0, 391, 389, 1, 0, 0, 0, 391, 392, 1, 0, 0, 0, 392,
		395, 1, 0, 0, 0, 393, 391, 1, 0, 0, 0, 394, 396, 5, 3, 0, 0, 395, 394,
		1, 0, 0, 0, 395, 396, 1, 0, 0, 0, 396, 397, 1, 0, 0, 0, 397, 398, 3, 122,
		61, 0, 398, 67, 1, 0, 0, 0, 399, 402, 3, 116, 58, 0, 400, 402, 3, 112,
		56, 0, 401, 399, 1, 0, 0, 0, 401, 400, 1, 0, 0, 0, 402, 403, 1, 0, 0, 0,
		403, 405, 5, 5, 0, 0, 404, 406, 3, 70, 35, 0, 405, 404, 1, 0, 0, 0, 406,
		407, 1, 0, 0, 0, 407, 405, 1, 0, 0, 0, 407, 408, 1, 0, 0, 0, 408, 69, 1,
		0, 0, 0, 409, 417, 3, 112, 56, 0, 410, 417, 3, 116, 58, 0, 411, 417, 3,
		118, 59, 0, 412, 417, 3, 56, 28, 0, 413, 417, 3, 64, 32, 0, 414, 417, 3,
		52, 26, 0, 415, 417, 3, 62, 31, 0, 416, 409, 1, 0, 0, 0, 416, 410, 1, 0,
		0, 0, 416, 411, 1, 0, 0, 0, 416, 412, 1, 0, 0, 0, 416, 413, 1, 0, 0, 0,
		416, 414, 1, 0, 0, 0, 416, 415, 1, 0, 0, 0, 417, 71, 1, 0, 0, 0, 418, 420,
		5, 1, 0, 0, 419, 421, 3, 74, 37, 0, 420, 419, 1, 0, 0, 0, 420, 421, 1,
		0, 0, 0, 421, 422, 1, 0, 0, 0, 422, 423, 5, 2, 0, 0, 423, 73, 1, 0, 0,
		0, 424, 427, 3, 84, 42, 0, 425, 427, 3, 94, 47, 0, 426, 424, 1, 0, 0, 0,
		426, 425, 1, 0, 0, 0, 427, 435, 1, 0, 0, 0, 428, 431, 5, 3, 0, 0, 429,
		432, 3, 84, 42, 0, 430, 432, 3, 94, 47, 0, 431, 429, 1, 0, 0, 0, 431, 430,
		1, 0, 0, 0, 432, 434, 1, 0, 0, 0, 433, 428, 1, 0, 0, 0, 434, 437, 1, 0,
		0, 0, 435, 433, 1, 0, 0, 0, 435, 436, 1, 0, 0, 0, 436, 75, 1, 0, 0, 0,
		437, 435, 1, 0, 0, 0, 438, 443, 3, 86, 43, 0, 439, 443, 3, 80, 40, 0, 440,
		443, 3, 78, 39, 0, 441, 443, 3, 82, 41, 0, 442, 438, 1, 0, 0, 0, 442, 439,
		1, 0, 0, 0, 442, 440, 1, 0, 0, 0, 442, 441, 1, 0, 0, 0, 443, 77, 1, 0,
		0, 0, 444, 445, 3, 94, 47, 0, 445, 446, 3, 96, 48, 0, 446, 79, 1, 0, 0,
		0, 447, 448, 3, 94, 47, 0, 448, 81, 1, 0, 0, 0, 449, 450, 3, 92, 46, 0,
		450, 451, 5, 6, 0, 0, 451, 452, 3, 84, 42, 0, 452, 83, 1, 0, 0, 0, 453,
		456, 3, 64, 32, 0, 454, 456, 3, 56, 28, 0, 455, 453, 1, 0, 0, 0, 455, 454,
		1, 0, 0, 0, 456, 85, 1, 0, 0, 0, 457, 458, 3, 92, 46, 0, 458, 459, 3, 108,
		54, 0, 459, 460, 3, 92, 46, 0, 460, 87, 1, 0, 0, 0, 461, 462, 6, 44, -1,
		0, 462, 463, 3, 90, 45, 0, 463, 476, 1, 0, 0, 0, 464, 472, 10, 1, 0, 0,
		465, 466, 5, 4, 0, 0, 466, 468, 5, 74, 0, 0, 467, 465, 1, 0, 0, 0, 467,
		468, 1, 0, 0, 0, 468, 469, 1, 0, 0, 0, 469, 473, 3, 72, 36, 0, 470, 473,
		3, 52, 26, 0, 471, 473, 3, 62, 31, 0, 472, 467, 1, 0, 0, 0, 472, 470, 1,
		0, 0, 0, 472, 471, 1, 0, 0, 0, 473, 475, 1, 0, 0, 0, 474, 464, 1, 0, 0,
		0, 475, 478, 1, 0, 0, 0, 476, 474, 1, 0, 0, 0, 476, 477, 1, 0, 0, 0, 477,
		89, 1, 0, 0, 0, 478, 476, 1, 0, 0, 0, 479, 489, 3, 110, 55, 0, 480, 489,
		3, 118, 59, 0, 481, 489, 3, 116, 58, 0, 482, 489, 3, 120, 60, 0, 483, 484,
		5, 1, 0, 0, 484, 485, 3, 94, 47, 0, 485, 486, 5, 2, 0, 0, 486, 489, 1,
		0, 0, 0, 487, 489, 5, 59, 0, 0, 488, 479, 1, 0, 0, 0, 488, 480, 1, 0, 0,
		0, 488, 481, 1, 0, 0, 0, 488, 482, 1, 0, 0, 0, 488, 483, 1, 0, 0, 0, 488,
		487, 1, 0, 0, 0, 489, 91, 1, 0, 0, 0, 490, 495, 3, 94, 47, 0, 491, 492,
		5, 3, 0, 0, 492, 494, 3, 94, 47, 0, 493, 491, 1, 0, 0, 0, 494, 497, 1,
		0, 0, 0, 495, 493, 1, 0, 0, 0, 495, 496, 1, 0, 0, 0, 496, 93, 1, 0, 0,
		0, 497, 495, 1, 0, 0, 0, 498, 499, 6, 47, -1, 0, 499, 500, 3, 88, 44, 0,
		500, 501, 3, 122, 61, 0, 501, 506, 1, 0, 0, 0, 502, 503, 3, 100, 50, 0,
		503, 504, 3, 94, 47, 5, 504, 506, 1, 0, 0, 0, 505, 498, 1, 0, 0, 0, 505,
		502, 1, 0, 0, 0, 506, 525, 1, 0, 0, 0, 507, 508, 10, 4, 0, 0, 508, 509,
		3, 98, 49, 0, 509, 510, 3, 94, 47, 5, 510, 524, 1, 0, 0, 0, 511, 512, 10,
		3, 0, 0, 512, 513, 3, 106, 53, 0, 513, 514, 3, 94, 47, 4, 514, 524, 1,
		0, 0, 0, 515, 516, 10, 2, 0, 0, 516, 517, 3, 104, 52, 0, 517, 518, 3, 94,
		47, 3, 518, 524, 1, 0, 0, 0, 519, 520, 10, 1, 0, 0, 520, 521, 3, 102, 51,
		0, 521, 522, 3, 94, 47, 2, 522, 524, 1, 0, 0, 0, 523, 507, 1, 0, 0, 0,
		523, 511, 1, 0, 0, 0, 523, 515, 1, 0, 0, 0, 523, 519, 1, 0, 0, 0, 524,
		527, 1, 0, 0, 0, 525, 523, 1, 0, 0, 0, 525, 526, 1, 0, 0, 0, 526, 95, 1,
		0, 0, 0, 527, 525, 1, 0, 0, 0, 528, 529, 7, 3, 0, 0, 529, 97, 1, 0, 0,
		0, 530, 531, 7, 4, 0, 0, 531, 99, 1, 0, 0, 0, 532, 533, 7, 5, 0, 0, 533,
		101, 1, 0, 0, 0, 534, 535, 7, 6, 0, 0, 535, 103, 1, 0, 0, 0, 536, 537,
		7, 7, 0, 0, 537, 105, 1, 0, 0, 0, 538, 539, 7, 8, 0, 0, 539, 107, 1, 0,
		0, 0, 540, 541, 7, 9, 0, 0, 541, 109, 1, 0, 0, 0, 542, 543, 7, 10, 0, 0,
		543, 111, 1, 0, 0, 0, 544, 545, 7, 11, 0, 0, 545, 113, 1, 0, 0, 0, 546,
		553, 5, 74, 0, 0, 547, 549, 5, 3, 0, 0, 548, 547, 1, 0, 0, 0, 548, 549,
		1, 0, 0, 0, 549, 550, 1, 0, 0, 0, 550, 552, 5, 74, 0, 0, 551, 548, 1, 0,
		0, 0, 552, 555, 1, 0, 0, 0, 553, 551, 1, 0, 0, 0, 553, 554, 1, 0, 0, 0,
		554, 115, 1, 0, 0, 0, 555, 553, 1, 0, 0, 0, 556, 557, 7, 12, 0, 0, 557,
		117, 1, 0, 0, 0, 558, 559, 7, 13, 0, 0, 559, 119, 1, 0, 0, 0, 560, 561,
		7, 14, 0, 0, 561, 121, 1, 0, 0, 0, 562, 568, 5, 11, 0, 0, 563, 568, 5,
		98, 0, 0, 564, 566, 5, 0, 0, 1, 565, 564, 1, 0, 0, 0, 565, 566, 1, 0, 0,
		0, 566, 568, 1, 0, 0, 0, 567, 562, 1, 0, 0, 0, 567, 563, 1, 0, 0, 0, 567,
		565, 1, 0, 0, 0, 568, 123, 1, 0, 0, 0, 62, 126, 132, 145, 153, 162, 170,
		181, 189, 193, 198, 206, 215, 218, 223, 228, 232, 234, 241, 248, 253, 258,
		263, 268, 271, 273, 280, 291, 302, 311, 317, 322, 326, 343, 347, 353, 357,
		366, 374, 382, 391, 395, 401, 407, 416, 420, 426, 431, 435, 442, 455, 467,
		472, 476, 488, 495, 505, 523, 525, 548, 553, 565, 567,
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
	NeedleParserRULE_settingsValue    = 6
	NeedleParserRULE_funcDef          = 7
	NeedleParserRULE_innerFuncDef     = 8
	NeedleParserRULE_defaultFuncDef   = 9
	NeedleParserRULE_funcDescriptor   = 10
	NeedleParserRULE_funcSignature    = 11
	NeedleParserRULE_parameterList    = 12
	NeedleParserRULE_parameter        = 13
	NeedleParserRULE_returnParameters = 14
	NeedleParserRULE_block            = 15
	NeedleParserRULE_statList         = 16
	NeedleParserRULE_stat             = 17
	NeedleParserRULE_varDef           = 18
	NeedleParserRULE_ifStat           = 19
	NeedleParserRULE_returnStat       = 20
	NeedleParserRULE_continueStat     = 21
	NeedleParserRULE_breakStat        = 22
	NeedleParserRULE_whileStat        = 23
	NeedleParserRULE_errorStat        = 24
	NeedleParserRULE_typeName         = 25
	NeedleParserRULE_sliceStat        = 26
	NeedleParserRULE_indexNumber      = 27
	NeedleParserRULE_arrayStat        = 28
	NeedleParserRULE_arrayList        = 29
	NeedleParserRULE_arrayValue       = 30
	NeedleParserRULE_indexStat        = 31
	NeedleParserRULE_objectStat       = 32
	NeedleParserRULE_pairList         = 33
	NeedleParserRULE_pair             = 34
	NeedleParserRULE_pairValue        = 35
	NeedleParserRULE_arguments        = 36
	NeedleParserRULE_argumentsList    = 37
	NeedleParserRULE_simpleStat       = 38
	NeedleParserRULE_incDecStat       = 39
	NeedleParserRULE_exprStat         = 40
	NeedleParserRULE_assignMapArrStat = 41
	NeedleParserRULE_initMapArrStat   = 42
	NeedleParserRULE_assignment       = 43
	NeedleParserRULE_primaryExpr      = 44
	NeedleParserRULE_operand          = 45
	NeedleParserRULE_exprList         = 46
	NeedleParserRULE_expr             = 47
	NeedleParserRULE_incDec_op        = 48
	NeedleParserRULE_mul_op           = 49
	NeedleParserRULE_unary_op         = 50
	NeedleParserRULE_add_op           = 51
	NeedleParserRULE_logical_op       = 52
	NeedleParserRULE_rel_op           = 53
	NeedleParserRULE_assign_op        = 54
	NeedleParserRULE_identifierFull   = 55
	NeedleParserRULE_identifierVar    = 56
	NeedleParserRULE_identifierList   = 57
	NeedleParserRULE_stringLiteral    = 58
	NeedleParserRULE_numberLiteral    = 59
	NeedleParserRULE_booleanLiteral   = 60
	NeedleParserRULE_eos              = 61
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
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3458790902099607552) != 0 {
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case NeedleParserCONTRACT:
			{
				p.SetState(124)
				p.ContractDef()
			}

		case NeedleParserFUNC, NeedleParserACTION, NeedleParserCONDITIONS:
			{
				p.SetState(125)
				p.FuncDef()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(128)
			p.Eos()
		}

		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(135)
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
		p.SetState(137)
		p.Match(NeedleParserCONTRACT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
		p.Match(NeedleParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.Match(NeedleParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3485803703770808320) != 0 {
		{
			p.SetState(140)
			p.ContractPart()
		}
		{
			p.SetState(141)
			p.Eos()
		}

		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(148)
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
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NeedleParserDATA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(150)
			p.DataDef()
		}

	case NeedleParserSETTINGS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(151)
			p.SettingsDef()
		}

	case NeedleParserFUNC, NeedleParserACTION, NeedleParserCONDITIONS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(152)
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
		p.SetState(155)
		p.Match(NeedleParserDATA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Match(NeedleParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == NeedleParserIdentifier {
		{
			p.SetState(157)
			p.DataPartList()
		}
		{
			p.SetState(158)
			p.Eos()
		}

		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(165)
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
		p.SetState(167)
		p.Match(NeedleParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(168)
		p.TypeName()
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NeedleParserInterpretedStringLiteral || _la == NeedleParserRawStringLiteral {
		{
			p.SetState(169)

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
	AllSettingsValue() []ISettingsValueContext
	SettingsValue(i int) ISettingsValueContext
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

func (s *SettingsDefContext) AllSettingsValue() []ISettingsValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISettingsValueContext); ok {
			len++
		}
	}

	tst := make([]ISettingsValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISettingsValueContext); ok {
			tst[i] = t.(ISettingsValueContext)
			i++
		}
	}

	return tst
}

func (s *SettingsDefContext) SettingsValue(i int) ISettingsValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISettingsValueContext); ok {
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

	return t.(ISettingsValueContext)
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
		p.SetState(172)
		p.Match(NeedleParserSETTINGS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
		p.Match(NeedleParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == NeedleParserIdentifier {
		{
			p.SetState(174)
			p.Match(NeedleParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(175)
			p.Match(NeedleParserEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
			p.SettingsValue()
		}
		{
			p.SetState(177)
			p.Eos()
		}

		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(184)
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

// ISettingsValueContext is an interface to support dynamic dispatch.
type ISettingsValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NumberLiteral() INumberLiteralContext
	BooleanLiteral() IBooleanLiteralContext
	StringLiteral() IStringLiteralContext

	// IsSettingsValueContext differentiates from other interfaces.
	IsSettingsValueContext()
}

type SettingsValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySettingsValueContext() *SettingsValueContext {
	var p = new(SettingsValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_settingsValue
	return p
}

func InitEmptySettingsValueContext(p *SettingsValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_settingsValue
}

func (*SettingsValueContext) IsSettingsValueContext() {}

func NewSettingsValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SettingsValueContext {
	var p = new(SettingsValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_settingsValue

	return p
}

func (s *SettingsValueContext) GetParser() antlr.Parser { return s.parser }

func (s *SettingsValueContext) NumberLiteral() INumberLiteralContext {
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

func (s *SettingsValueContext) BooleanLiteral() IBooleanLiteralContext {
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

func (s *SettingsValueContext) StringLiteral() IStringLiteralContext {
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

func (s *SettingsValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SettingsValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SettingsValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterSettingsValue(s)
	}
}

func (s *SettingsValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitSettingsValue(s)
	}
}

func (p *NeedleParser) SettingsValue() (localctx ISettingsValueContext) {
	localctx = NewSettingsValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, NeedleParserRULE_settingsValue)
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NeedleParserDecimalLiteral, NeedleParserFloatLiteral, NeedleParserHexLiteral, NeedleParserOctalLiteral, NeedleParserBinaryLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(186)
			p.NumberLiteral()
		}

	case NeedleParserTRUE, NeedleParserFALSE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(187)
			p.BooleanLiteral()
		}

	case NeedleParserInterpretedStringLiteral, NeedleParserRawStringLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(188)
			p.StringLiteral()
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

// IFuncDefContext is an interface to support dynamic dispatch.
type IFuncDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	InnerFuncDef() IInnerFuncDefContext
	DefaultFuncDef() IDefaultFuncDefContext
	DOT() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	FuncSignature() IFuncSignatureContext

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

func (s *FuncDefContext) InnerFuncDef() IInnerFuncDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInnerFuncDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInnerFuncDefContext)
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

func (s *FuncDefContext) DOT() antlr.TerminalNode {
	return s.GetToken(NeedleParserDOT, 0)
}

func (s *FuncDefContext) Identifier() antlr.TerminalNode {
	return s.GetToken(NeedleParserIdentifier, 0)
}

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
	p.EnterRule(localctx, 14, NeedleParserRULE_funcDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(191)
			p.InnerFuncDef()
		}

	case 2:
		{
			p.SetState(192)
			p.DefaultFuncDef()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NeedleParserDOT {
		{
			p.SetState(195)
			p.Match(NeedleParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(196)
			p.Match(NeedleParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)
			p.FuncSignature()
		}

	}
	{
		p.SetState(200)
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

// IInnerFuncDefContext is an interface to support dynamic dispatch.
type IInnerFuncDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FuncDescriptor() IFuncDescriptorContext
	FuncSignature() IFuncSignatureContext

	// IsInnerFuncDefContext differentiates from other interfaces.
	IsInnerFuncDefContext()
}

type InnerFuncDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInnerFuncDefContext() *InnerFuncDefContext {
	var p = new(InnerFuncDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_innerFuncDef
	return p
}

func InitEmptyInnerFuncDefContext(p *InnerFuncDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_innerFuncDef
}

func (*InnerFuncDefContext) IsInnerFuncDefContext() {}

func NewInnerFuncDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InnerFuncDefContext {
	var p = new(InnerFuncDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_innerFuncDef

	return p
}

func (s *InnerFuncDefContext) GetParser() antlr.Parser { return s.parser }

func (s *InnerFuncDefContext) FuncDescriptor() IFuncDescriptorContext {
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

func (s *InnerFuncDefContext) FuncSignature() IFuncSignatureContext {
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

func (s *InnerFuncDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerFuncDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InnerFuncDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterInnerFuncDef(s)
	}
}

func (s *InnerFuncDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitInnerFuncDef(s)
	}
}

func (p *NeedleParser) InnerFuncDef() (localctx IInnerFuncDefContext) {
	localctx = NewInnerFuncDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, NeedleParserRULE_innerFuncDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.FuncDescriptor()
	}
	{
		p.SetState(203)
		p.FuncSignature()
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
	FuncSignature() IFuncSignatureContext
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

func (s *DefaultFuncDefContext) FuncSignature() IFuncSignatureContext {
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
	p.EnterRule(localctx, 18, NeedleParserRULE_defaultFuncDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NeedleParserFUNC {
		{
			p.SetState(205)
			p.Match(NeedleParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(208)
		_la = p.GetTokenStream().LA(1)

		if !(_la == NeedleParserACTION || _la == NeedleParserCONDITIONS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(209)
		p.FuncSignature()
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
	p.EnterRule(localctx, 20, NeedleParserRULE_funcDescriptor)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		p.Match(NeedleParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(212)
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
	p.EnterRule(localctx, 22, NeedleParserRULE_funcSignature)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NeedleParserLPAREN {
		{
			p.SetState(214)
			p.ParameterList()
		}

	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&1023) != 0 {
		{
			p.SetState(217)
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

// IParameterListContext is an interface to support dynamic dispatch.
type IParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

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

func (s *ParameterListContext) AllParameter() []IParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterContext); ok {
			len++
		}
	}

	tst := make([]IParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterContext); ok {
			tst[i] = t.(IParameterContext)
			i++
		}
	}

	return tst
}

func (s *ParameterListContext) Parameter(i int) IParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
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

	return t.(IParameterContext)
}

func (s *ParameterListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(NeedleParserCOMMA)
}

func (s *ParameterListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(NeedleParserCOMMA, i)
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
	p.EnterRule(localctx, 24, NeedleParserRULE_parameterList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(220)
		p.Match(NeedleParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NeedleParserIdentifier {
		{
			p.SetState(221)
			p.Parameter()
		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				p.SetState(223)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == NeedleParserCOMMA {
					{
						p.SetState(222)
						p.Match(NeedleParserCOMMA)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(225)
					p.Parameter()
				}

			}
			p.SetState(230)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == NeedleParserCOMMA {
			{
				p.SetState(231)
				p.Match(NeedleParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(236)
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
	p.EnterRule(localctx, 26, NeedleParserRULE_parameter)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.IdentifierList()
	}
	{
		p.SetState(239)
		p.TypeName()
	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(241)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == NeedleParserCOMMA {
				{
					p.SetState(240)
					p.Match(NeedleParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(243)
				p.IdentifierList()
			}
			{
				p.SetState(244)
				p.TypeName()
			}

		}
		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 28, NeedleParserRULE_returnParameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)
		p.TypeName()
	}
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == NeedleParserCOMMA || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&1023) != 0) {
		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == NeedleParserCOMMA {
			{
				p.SetState(252)
				p.Match(NeedleParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(255)
			p.TypeName()
		}

		p.SetState(260)
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
	StatList() IStatListContext

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

func (s *BlockContext) StatList() IStatListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatListContext)
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
	p.EnterRule(localctx, 30, NeedleParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(261)
		p.Match(NeedleParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-8097929526849250814) != 0) || ((int64((_la-74)) & ^0x3f) == 0 && ((int64(1)<<(_la-74))&16778239) != 0) {
		{
			p.SetState(262)
			p.StatList()
		}

	}
	{
		p.SetState(265)
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

// IStatListContext is an interface to support dynamic dispatch.
type IStatListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStat() []IStatContext
	Stat(i int) IStatContext
	AllEos() []IEosContext
	Eos(i int) IEosContext
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode
	AllEOS() []antlr.TerminalNode
	EOS(i int) antlr.TerminalNode

	// IsStatListContext differentiates from other interfaces.
	IsStatListContext()
}

type StatListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatListContext() *StatListContext {
	var p = new(StatListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_statList
	return p
}

func InitEmptyStatListContext(p *StatListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_statList
}

func (*StatListContext) IsStatListContext() {}

func NewStatListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatListContext {
	var p = new(StatListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_statList

	return p
}

func (s *StatListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatListContext) AllStat() []IStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatContext); ok {
			len++
		}
	}

	tst := make([]IStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatContext); ok {
			tst[i] = t.(IStatContext)
			i++
		}
	}

	return tst
}

func (s *StatListContext) Stat(i int) IStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
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

	return t.(IStatContext)
}

func (s *StatListContext) AllEos() []IEosContext {
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

func (s *StatListContext) Eos(i int) IEosContext {
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

func (s *StatListContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(NeedleParserSEMI)
}

func (s *StatListContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(NeedleParserSEMI, i)
}

func (s *StatListContext) AllEOS() []antlr.TerminalNode {
	return s.GetTokens(NeedleParserEOS)
}

func (s *StatListContext) EOS(i int) antlr.TerminalNode {
	return s.GetToken(NeedleParserEOS, i)
}

func (s *StatListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterStatList(s)
	}
}

func (s *StatListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitStatList(s)
	}
}

func (p *NeedleParser) StatList() (localctx IStatListContext) {
	localctx = NewStatListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, NeedleParserRULE_statList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-8097929526849250814) != 0) || ((int64((_la-74)) & ^0x3f) == 0 && ((int64(1)<<(_la-74))&16778239) != 0) {
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
		case 1:
			p.SetState(268)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == NeedleParserSEMI {
				{
					p.SetState(267)
					p.Match(NeedleParserSEMI)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}

		case 2:
			p.SetState(271)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == NeedleParserEOS {
				{
					p.SetState(270)
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
			p.SetState(275)
			p.Stat()
		}
		{
			p.SetState(276)
			p.Eos()
		}

		p.SetState(280)
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

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	SimpleStat() ISimpleStatContext
	VarDef() IVarDefContext
	IfStat() IIfStatContext
	WhileStat() IWhileStatContext
	ContinueStat() IContinueStatContext
	BreakStat() IBreakStatContext
	ReturnStat() IReturnStatContext
	ErrorStat() IErrorStatContext

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_stat
	return p
}

func InitEmptyStatContext(p *StatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_stat
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) Block() IBlockContext {
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

func (s *StatContext) SimpleStat() ISimpleStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatContext)
}

func (s *StatContext) VarDef() IVarDefContext {
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

func (s *StatContext) IfStat() IIfStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatContext)
}

func (s *StatContext) WhileStat() IWhileStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStatContext)
}

func (s *StatContext) ContinueStat() IContinueStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinueStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinueStatContext)
}

func (s *StatContext) BreakStat() IBreakStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakStatContext)
}

func (s *StatContext) ReturnStat() IReturnStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStatContext)
}

func (s *StatContext) ErrorStat() IErrorStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IErrorStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IErrorStatContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterStat(s)
	}
}

func (s *StatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitStat(s)
	}
}

func (p *NeedleParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, NeedleParserRULE_stat)
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NeedleParserLBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(282)
			p.Block()
		}

	case NeedleParserLPAREN, NeedleParserNOT, NeedleParserADD, NeedleParserSUB, NeedleParserTRUE, NeedleParserFALSE, NeedleParserNIL, NeedleParserIdentifier, NeedleParserDollarIdentifier, NeedleParserAtIdentifier, NeedleParserInterpretedStringLiteral, NeedleParserRawStringLiteral, NeedleParserDecimalLiteral, NeedleParserFloatLiteral, NeedleParserHexLiteral, NeedleParserOctalLiteral, NeedleParserBinaryLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(283)
			p.SimpleStat()
		}

	case NeedleParserVAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(284)
			p.VarDef()
		}

	case NeedleParserIF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(285)
			p.IfStat()
		}

	case NeedleParserWHILE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(286)
			p.WhileStat()
		}

	case NeedleParserCONTINUE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(287)
			p.ContinueStat()
		}

	case NeedleParserBREAK:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(288)
			p.BreakStat()
		}

	case NeedleParserRETURN:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(289)
			p.ReturnStat()
		}

	case NeedleParserERRWARNING, NeedleParserERRINFO, NeedleParserERROR:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(290)
			p.ErrorStat()
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
	p.EnterRule(localctx, 36, NeedleParserRULE_varDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.Match(NeedleParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(294)
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

// IIfStatContext is an interface to support dynamic dispatch.
type IIfStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllELIF() []antlr.TerminalNode
	ELIF(i int) antlr.TerminalNode
	ELSE() antlr.TerminalNode
	AllLPAREN() []antlr.TerminalNode
	LPAREN(i int) antlr.TerminalNode
	AllRPAREN() []antlr.TerminalNode
	RPAREN(i int) antlr.TerminalNode

	// IsIfStatContext differentiates from other interfaces.
	IsIfStatContext()
}

type IfStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatContext() *IfStatContext {
	var p = new(IfStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_ifStat
	return p
}

func InitEmptyIfStatContext(p *IfStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_ifStat
}

func (*IfStatContext) IsIfStatContext() {}

func NewIfStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatContext {
	var p = new(IfStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_ifStat

	return p
}

func (s *IfStatContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatContext) IF() antlr.TerminalNode {
	return s.GetToken(NeedleParserIF, 0)
}

func (s *IfStatContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStatContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
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

	return t.(IBlockContext)
}

func (s *IfStatContext) AllExpr() []IExprContext {
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

func (s *IfStatContext) Expr(i int) IExprContext {
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

func (s *IfStatContext) AllELIF() []antlr.TerminalNode {
	return s.GetTokens(NeedleParserELIF)
}

func (s *IfStatContext) ELIF(i int) antlr.TerminalNode {
	return s.GetToken(NeedleParserELIF, i)
}

func (s *IfStatContext) ELSE() antlr.TerminalNode {
	return s.GetToken(NeedleParserELSE, 0)
}

func (s *IfStatContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(NeedleParserLPAREN)
}

func (s *IfStatContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(NeedleParserLPAREN, i)
}

func (s *IfStatContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(NeedleParserRPAREN)
}

func (s *IfStatContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(NeedleParserRPAREN, i)
}

func (s *IfStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterIfStat(s)
	}
}

func (s *IfStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitIfStat(s)
	}
}

func (p *NeedleParser) IfStat() (localctx IIfStatContext) {
	localctx = NewIfStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, NeedleParserRULE_ifStat)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)
		p.Match(NeedleParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(297)
			p.Match(NeedleParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(298)
			p.expr(0)
		}
		{
			p.SetState(299)
			p.Match(NeedleParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(301)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(304)
		p.Block()
	}
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(305)
				p.Match(NeedleParserELIF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(311)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(306)
					p.Match(NeedleParserLPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(307)
					p.expr(0)
				}
				{
					p.SetState(308)
					p.Match(NeedleParserRPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 2:
				{
					p.SetState(310)
					p.expr(0)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}
			{
				p.SetState(313)
				p.Block()
			}

		}
		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(322)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NeedleParserELSE {
		{
			p.SetState(320)
			p.Match(NeedleParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(321)
			p.Block()
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

// IReturnStatContext is an interface to support dynamic dispatch.
type IReturnStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expr() IExprContext

	// IsReturnStatContext differentiates from other interfaces.
	IsReturnStatContext()
}

type ReturnStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatContext() *ReturnStatContext {
	var p = new(ReturnStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_returnStat
	return p
}

func InitEmptyReturnStatContext(p *ReturnStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_returnStat
}

func (*ReturnStatContext) IsReturnStatContext() {}

func NewReturnStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatContext {
	var p = new(ReturnStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_returnStat

	return p
}

func (s *ReturnStatContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatContext) RETURN() antlr.TerminalNode {
	return s.GetToken(NeedleParserRETURN, 0)
}

func (s *ReturnStatContext) Expr() IExprContext {
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

func (s *ReturnStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterReturnStat(s)
	}
}

func (s *ReturnStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitReturnStat(s)
	}
}

func (p *NeedleParser) ReturnStat() (localctx IReturnStatContext) {
	localctx = NewReturnStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, NeedleParserRULE_returnStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(324)
		p.Match(NeedleParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(326)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(325)
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

// IContinueStatContext is an interface to support dynamic dispatch.
type IContinueStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTINUE() antlr.TerminalNode

	// IsContinueStatContext differentiates from other interfaces.
	IsContinueStatContext()
}

type ContinueStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStatContext() *ContinueStatContext {
	var p = new(ContinueStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_continueStat
	return p
}

func InitEmptyContinueStatContext(p *ContinueStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_continueStat
}

func (*ContinueStatContext) IsContinueStatContext() {}

func NewContinueStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStatContext {
	var p = new(ContinueStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_continueStat

	return p
}

func (s *ContinueStatContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinueStatContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(NeedleParserCONTINUE, 0)
}

func (s *ContinueStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterContinueStat(s)
	}
}

func (s *ContinueStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitContinueStat(s)
	}
}

func (p *NeedleParser) ContinueStat() (localctx IContinueStatContext) {
	localctx = NewContinueStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, NeedleParserRULE_continueStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(328)
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

// IBreakStatContext is an interface to support dynamic dispatch.
type IBreakStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BREAK() antlr.TerminalNode

	// IsBreakStatContext differentiates from other interfaces.
	IsBreakStatContext()
}

type BreakStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStatContext() *BreakStatContext {
	var p = new(BreakStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_breakStat
	return p
}

func InitEmptyBreakStatContext(p *BreakStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_breakStat
}

func (*BreakStatContext) IsBreakStatContext() {}

func NewBreakStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStatContext {
	var p = new(BreakStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_breakStat

	return p
}

func (s *BreakStatContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakStatContext) BREAK() antlr.TerminalNode {
	return s.GetToken(NeedleParserBREAK, 0)
}

func (s *BreakStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterBreakStat(s)
	}
}

func (s *BreakStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitBreakStat(s)
	}
}

func (p *NeedleParser) BreakStat() (localctx IBreakStatContext) {
	localctx = NewBreakStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, NeedleParserRULE_breakStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(330)
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

// IWhileStatContext is an interface to support dynamic dispatch.
type IWhileStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expr() IExprContext
	Block() IBlockContext

	// IsWhileStatContext differentiates from other interfaces.
	IsWhileStatContext()
}

type WhileStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStatContext() *WhileStatContext {
	var p = new(WhileStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_whileStat
	return p
}

func InitEmptyWhileStatContext(p *WhileStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_whileStat
}

func (*WhileStatContext) IsWhileStatContext() {}

func NewWhileStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatContext {
	var p = new(WhileStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_whileStat

	return p
}

func (s *WhileStatContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatContext) WHILE() antlr.TerminalNode {
	return s.GetToken(NeedleParserWHILE, 0)
}

func (s *WhileStatContext) Expr() IExprContext {
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

func (s *WhileStatContext) Block() IBlockContext {
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

func (s *WhileStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterWhileStat(s)
	}
}

func (s *WhileStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitWhileStat(s)
	}
}

func (p *NeedleParser) WhileStat() (localctx IWhileStatContext) {
	localctx = NewWhileStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, NeedleParserRULE_whileStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(332)
		p.Match(NeedleParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(333)
		p.expr(0)
	}
	{
		p.SetState(334)
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

// IErrorStatContext is an interface to support dynamic dispatch.
type IErrorStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	ERRWARNING() antlr.TerminalNode
	ERRINFO() antlr.TerminalNode
	ERROR() antlr.TerminalNode

	// IsErrorStatContext differentiates from other interfaces.
	IsErrorStatContext()
}

type ErrorStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyErrorStatContext() *ErrorStatContext {
	var p = new(ErrorStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_errorStat
	return p
}

func InitEmptyErrorStatContext(p *ErrorStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_errorStat
}

func (*ErrorStatContext) IsErrorStatContext() {}

func NewErrorStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ErrorStatContext {
	var p = new(ErrorStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_errorStat

	return p
}

func (s *ErrorStatContext) GetParser() antlr.Parser { return s.parser }

func (s *ErrorStatContext) Expr() IExprContext {
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

func (s *ErrorStatContext) ERRWARNING() antlr.TerminalNode {
	return s.GetToken(NeedleParserERRWARNING, 0)
}

func (s *ErrorStatContext) ERRINFO() antlr.TerminalNode {
	return s.GetToken(NeedleParserERRINFO, 0)
}

func (s *ErrorStatContext) ERROR() antlr.TerminalNode {
	return s.GetToken(NeedleParserERROR, 0)
}

func (s *ErrorStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ErrorStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ErrorStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterErrorStat(s)
	}
}

func (s *ErrorStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitErrorStat(s)
	}
}

func (p *NeedleParser) ErrorStat() (localctx IErrorStatContext) {
	localctx = NewErrorStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, NeedleParserRULE_errorStat)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(336)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-8791026472627208192) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(337)
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
	p.EnterRule(localctx, 50, NeedleParserRULE_typeName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
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

// ISliceStatContext is an interface to support dynamic dispatch.
type ISliceStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	COLON() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	AllIndexNumber() []IIndexNumberContext
	IndexNumber(i int) IIndexNumberContext

	// IsSliceStatContext differentiates from other interfaces.
	IsSliceStatContext()
}

type SliceStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceStatContext() *SliceStatContext {
	var p = new(SliceStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_sliceStat
	return p
}

func InitEmptySliceStatContext(p *SliceStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_sliceStat
}

func (*SliceStatContext) IsSliceStatContext() {}

func NewSliceStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceStatContext {
	var p = new(SliceStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_sliceStat

	return p
}

func (s *SliceStatContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceStatContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(NeedleParserLBRACK, 0)
}

func (s *SliceStatContext) COLON() antlr.TerminalNode {
	return s.GetToken(NeedleParserCOLON, 0)
}

func (s *SliceStatContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(NeedleParserRBRACK, 0)
}

func (s *SliceStatContext) AllIndexNumber() []IIndexNumberContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIndexNumberContext); ok {
			len++
		}
	}

	tst := make([]IIndexNumberContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIndexNumberContext); ok {
			tst[i] = t.(IIndexNumberContext)
			i++
		}
	}

	return tst
}

func (s *SliceStatContext) IndexNumber(i int) IIndexNumberContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexNumberContext); ok {
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

	return t.(IIndexNumberContext)
}

func (s *SliceStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterSliceStat(s)
	}
}

func (s *SliceStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitSliceStat(s)
	}
}

func (p *NeedleParser) SliceStat() (localctx ISliceStatContext) {
	localctx = NewSliceStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, NeedleParserRULE_sliceStat)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(341)
		p.Match(NeedleParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(343)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-74)) & ^0x3f) == 0 && ((int64(1)<<(_la-74))&995) != 0 {
		{
			p.SetState(342)
			p.IndexNumber()
		}

	}
	{
		p.SetState(345)
		p.Match(NeedleParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(347)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-74)) & ^0x3f) == 0 && ((int64(1)<<(_la-74))&995) != 0 {
		{
			p.SetState(346)
			p.IndexNumber()
		}

	}
	{
		p.SetState(349)
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

// IIndexNumberContext is an interface to support dynamic dispatch.
type IIndexNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NumberLiteral() INumberLiteralContext
	IdentifierVar() IIdentifierVarContext

	// IsIndexNumberContext differentiates from other interfaces.
	IsIndexNumberContext()
}

type IndexNumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexNumberContext() *IndexNumberContext {
	var p = new(IndexNumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_indexNumber
	return p
}

func InitEmptyIndexNumberContext(p *IndexNumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_indexNumber
}

func (*IndexNumberContext) IsIndexNumberContext() {}

func NewIndexNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexNumberContext {
	var p = new(IndexNumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_indexNumber

	return p
}

func (s *IndexNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexNumberContext) NumberLiteral() INumberLiteralContext {
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

func (s *IndexNumberContext) IdentifierVar() IIdentifierVarContext {
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

func (s *IndexNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterIndexNumber(s)
	}
}

func (s *IndexNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitIndexNumber(s)
	}
}

func (p *NeedleParser) IndexNumber() (localctx IIndexNumberContext) {
	localctx = NewIndexNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, NeedleParserRULE_indexNumber)
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NeedleParserDecimalLiteral, NeedleParserFloatLiteral, NeedleParserHexLiteral, NeedleParserOctalLiteral, NeedleParserBinaryLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(351)
			p.NumberLiteral()
		}

	case NeedleParserIdentifier, NeedleParserDollarIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(352)
			p.IdentifierVar()
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

// IArrayStatContext is an interface to support dynamic dispatch.
type IArrayStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	ArrayList() IArrayListContext

	// IsArrayStatContext differentiates from other interfaces.
	IsArrayStatContext()
}

type ArrayStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayStatContext() *ArrayStatContext {
	var p = new(ArrayStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_arrayStat
	return p
}

func InitEmptyArrayStatContext(p *ArrayStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_arrayStat
}

func (*ArrayStatContext) IsArrayStatContext() {}

func NewArrayStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayStatContext {
	var p = new(ArrayStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_arrayStat

	return p
}

func (s *ArrayStatContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayStatContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(NeedleParserLBRACK, 0)
}

func (s *ArrayStatContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(NeedleParserRBRACK, 0)
}

func (s *ArrayStatContext) ArrayList() IArrayListContext {
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

func (s *ArrayStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterArrayStat(s)
	}
}

func (s *ArrayStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitArrayStat(s)
	}
}

func (p *NeedleParser) ArrayStat() (localctx IArrayStatContext) {
	localctx = NewArrayStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, NeedleParserRULE_arrayStat)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		p.Match(NeedleParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&579838452024005250) != 0) || ((int64((_la-74)) & ^0x3f) == 0 && ((int64(1)<<(_la-74))&1023) != 0) {
		{
			p.SetState(356)
			p.ArrayList()
		}

	}
	{
		p.SetState(359)
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
	AllArrayValue() []IArrayValueContext
	ArrayValue(i int) IArrayValueContext
	Eos() IEosContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

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

func (s *ArrayListContext) AllArrayValue() []IArrayValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArrayValueContext); ok {
			len++
		}
	}

	tst := make([]IArrayValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArrayValueContext); ok {
			tst[i] = t.(IArrayValueContext)
			i++
		}
	}

	return tst
}

func (s *ArrayListContext) ArrayValue(i int) IArrayValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayValueContext); ok {
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

	return t.(IArrayValueContext)
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

func (s *ArrayListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(NeedleParserCOMMA)
}

func (s *ArrayListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(NeedleParserCOMMA, i)
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
	p.EnterRule(localctx, 58, NeedleParserRULE_arrayList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(361)
		p.ArrayValue()
	}
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == NeedleParserCOMMA {
		{
			p.SetState(362)
			p.Match(NeedleParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(363)
			p.ArrayValue()
		}

		p.SetState(368)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(369)
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

// IArrayValueContext is an interface to support dynamic dispatch.
type IArrayValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ArrayStat() IArrayStatContext
	Expr() IExprContext
	ObjectStat() IObjectStatContext

	// IsArrayValueContext differentiates from other interfaces.
	IsArrayValueContext()
}

type ArrayValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayValueContext() *ArrayValueContext {
	var p = new(ArrayValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_arrayValue
	return p
}

func InitEmptyArrayValueContext(p *ArrayValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_arrayValue
}

func (*ArrayValueContext) IsArrayValueContext() {}

func NewArrayValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayValueContext {
	var p = new(ArrayValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_arrayValue

	return p
}

func (s *ArrayValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayValueContext) ArrayStat() IArrayStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayStatContext)
}

func (s *ArrayValueContext) Expr() IExprContext {
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

func (s *ArrayValueContext) ObjectStat() IObjectStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectStatContext)
}

func (s *ArrayValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterArrayValue(s)
	}
}

func (s *ArrayValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitArrayValue(s)
	}
}

func (p *NeedleParser) ArrayValue() (localctx IArrayValueContext) {
	localctx = NewArrayValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, NeedleParserRULE_arrayValue)
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NeedleParserLBRACK:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(371)
			p.ArrayStat()
		}

	case NeedleParserLPAREN, NeedleParserNOT, NeedleParserADD, NeedleParserSUB, NeedleParserTRUE, NeedleParserFALSE, NeedleParserNIL, NeedleParserIdentifier, NeedleParserDollarIdentifier, NeedleParserAtIdentifier, NeedleParserInterpretedStringLiteral, NeedleParserRawStringLiteral, NeedleParserDecimalLiteral, NeedleParserFloatLiteral, NeedleParserHexLiteral, NeedleParserOctalLiteral, NeedleParserBinaryLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(372)
			p.expr(0)
		}

	case NeedleParserLBRACE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(373)
			p.ObjectStat()
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

// IIndexStatContext is an interface to support dynamic dispatch.
type IIndexStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	Expr() IExprContext
	RBRACK() antlr.TerminalNode

	// IsIndexStatContext differentiates from other interfaces.
	IsIndexStatContext()
}

type IndexStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexStatContext() *IndexStatContext {
	var p = new(IndexStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_indexStat
	return p
}

func InitEmptyIndexStatContext(p *IndexStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_indexStat
}

func (*IndexStatContext) IsIndexStatContext() {}

func NewIndexStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexStatContext {
	var p = new(IndexStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_indexStat

	return p
}

func (s *IndexStatContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexStatContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(NeedleParserLBRACK, 0)
}

func (s *IndexStatContext) Expr() IExprContext {
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

func (s *IndexStatContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(NeedleParserRBRACK, 0)
}

func (s *IndexStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterIndexStat(s)
	}
}

func (s *IndexStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitIndexStat(s)
	}
}

func (p *NeedleParser) IndexStat() (localctx IIndexStatContext) {
	localctx = NewIndexStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, NeedleParserRULE_indexStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(376)
		p.Match(NeedleParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(377)
		p.expr(0)
	}
	{
		p.SetState(378)
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

// IObjectStatContext is an interface to support dynamic dispatch.
type IObjectStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	PairList() IPairListContext

	// IsObjectStatContext differentiates from other interfaces.
	IsObjectStatContext()
}

type ObjectStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectStatContext() *ObjectStatContext {
	var p = new(ObjectStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_objectStat
	return p
}

func InitEmptyObjectStatContext(p *ObjectStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_objectStat
}

func (*ObjectStatContext) IsObjectStatContext() {}

func NewObjectStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectStatContext {
	var p = new(ObjectStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_objectStat

	return p
}

func (s *ObjectStatContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectStatContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(NeedleParserLBRACE, 0)
}

func (s *ObjectStatContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(NeedleParserRBRACE, 0)
}

func (s *ObjectStatContext) PairList() IPairListContext {
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

func (s *ObjectStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterObjectStat(s)
	}
}

func (s *ObjectStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitObjectStat(s)
	}
}

func (p *NeedleParser) ObjectStat() (localctx IObjectStatContext) {
	localctx = NewObjectStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, NeedleParserRULE_objectStat)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(380)
		p.Match(NeedleParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-74)) & ^0x3f) == 0 && ((int64(1)<<(_la-74))&27) != 0 {
		{
			p.SetState(381)
			p.PairList()
		}

	}
	{
		p.SetState(384)
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
	p.EnterRule(localctx, 66, NeedleParserRULE_pairList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(386)
		p.Pair()
	}
	p.SetState(391)
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
				p.SetState(387)
				p.Match(NeedleParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(388)
				p.Pair()
			}

		}
		p.SetState(393)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NeedleParserCOMMA {
		{
			p.SetState(394)
			p.Match(NeedleParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(397)
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
	StringLiteral() IStringLiteralContext
	IdentifierVar() IIdentifierVarContext
	AllPairValue() []IPairValueContext
	PairValue(i int) IPairValueContext

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

func (s *PairContext) AllPairValue() []IPairValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPairValueContext); ok {
			len++
		}
	}

	tst := make([]IPairValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPairValueContext); ok {
			tst[i] = t.(IPairValueContext)
			i++
		}
	}

	return tst
}

func (s *PairContext) PairValue(i int) IPairValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPairValueContext); ok {
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

	return t.(IPairValueContext)
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
	p.EnterRule(localctx, 68, NeedleParserRULE_pair)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(401)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NeedleParserInterpretedStringLiteral, NeedleParserRawStringLiteral:
		{
			p.SetState(399)
			p.StringLiteral()
		}

	case NeedleParserIdentifier, NeedleParserDollarIdentifier:
		{
			p.SetState(400)
			p.IdentifierVar()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(403)
		p.Match(NeedleParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(405)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == NeedleParserLBRACK || _la == NeedleParserLBRACE || ((int64((_la-74)) & ^0x3f) == 0 && ((int64(1)<<(_la-74))&1019) != 0) {
		{
			p.SetState(404)
			p.PairValue()
		}

		p.SetState(407)
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

// IPairValueContext is an interface to support dynamic dispatch.
type IPairValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentifierVar() IIdentifierVarContext
	StringLiteral() IStringLiteralContext
	NumberLiteral() INumberLiteralContext
	ArrayStat() IArrayStatContext
	ObjectStat() IObjectStatContext
	SliceStat() ISliceStatContext
	IndexStat() IIndexStatContext

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

func (s *PairValueContext) StringLiteral() IStringLiteralContext {
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

func (s *PairValueContext) NumberLiteral() INumberLiteralContext {
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

func (s *PairValueContext) ArrayStat() IArrayStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayStatContext)
}

func (s *PairValueContext) ObjectStat() IObjectStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectStatContext)
}

func (s *PairValueContext) SliceStat() ISliceStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceStatContext)
}

func (s *PairValueContext) IndexStat() IIndexStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexStatContext)
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
	p.EnterRule(localctx, 70, NeedleParserRULE_pairValue)
	p.SetState(416)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(409)
			p.IdentifierVar()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(410)
			p.StringLiteral()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(411)
			p.NumberLiteral()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(412)
			p.ArrayStat()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(413)
			p.ObjectStat()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(414)
			p.SliceStat()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(415)
			p.IndexStat()
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

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ArgumentsList() IArgumentsListContext

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

func (s *ArgumentsContext) ArgumentsList() IArgumentsListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsListContext)
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
	p.EnterRule(localctx, 72, NeedleParserRULE_arguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(418)
		p.Match(NeedleParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(420)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&579838452024005250) != 0) || ((int64((_la-74)) & ^0x3f) == 0 && ((int64(1)<<(_la-74))&1023) != 0) {
		{
			p.SetState(419)
			p.ArgumentsList()
		}

	}
	{
		p.SetState(422)
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

// IArgumentsListContext is an interface to support dynamic dispatch.
type IArgumentsListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllInitMapArrStat() []IInitMapArrStatContext
	InitMapArrStat(i int) IInitMapArrStatContext
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgumentsListContext differentiates from other interfaces.
	IsArgumentsListContext()
}

type ArgumentsListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsListContext() *ArgumentsListContext {
	var p = new(ArgumentsListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_argumentsList
	return p
}

func InitEmptyArgumentsListContext(p *ArgumentsListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_argumentsList
}

func (*ArgumentsListContext) IsArgumentsListContext() {}

func NewArgumentsListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsListContext {
	var p = new(ArgumentsListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_argumentsList

	return p
}

func (s *ArgumentsListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsListContext) AllInitMapArrStat() []IInitMapArrStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInitMapArrStatContext); ok {
			len++
		}
	}

	tst := make([]IInitMapArrStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInitMapArrStatContext); ok {
			tst[i] = t.(IInitMapArrStatContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentsListContext) InitMapArrStat(i int) IInitMapArrStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInitMapArrStatContext); ok {
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

	return t.(IInitMapArrStatContext)
}

func (s *ArgumentsListContext) AllExpr() []IExprContext {
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

func (s *ArgumentsListContext) Expr(i int) IExprContext {
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

func (s *ArgumentsListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(NeedleParserCOMMA)
}

func (s *ArgumentsListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(NeedleParserCOMMA, i)
}

func (s *ArgumentsListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterArgumentsList(s)
	}
}

func (s *ArgumentsListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitArgumentsList(s)
	}
}

func (p *NeedleParser) ArgumentsList() (localctx IArgumentsListContext) {
	localctx = NewArgumentsListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, NeedleParserRULE_argumentsList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(426)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NeedleParserLBRACK, NeedleParserLBRACE:
		{
			p.SetState(424)
			p.InitMapArrStat()
		}

	case NeedleParserLPAREN, NeedleParserNOT, NeedleParserADD, NeedleParserSUB, NeedleParserTRUE, NeedleParserFALSE, NeedleParserNIL, NeedleParserIdentifier, NeedleParserDollarIdentifier, NeedleParserAtIdentifier, NeedleParserInterpretedStringLiteral, NeedleParserRawStringLiteral, NeedleParserDecimalLiteral, NeedleParserFloatLiteral, NeedleParserHexLiteral, NeedleParserOctalLiteral, NeedleParserBinaryLiteral:
		{
			p.SetState(425)
			p.expr(0)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(435)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == NeedleParserCOMMA {
		{
			p.SetState(428)
			p.Match(NeedleParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(431)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case NeedleParserLBRACK, NeedleParserLBRACE:
			{
				p.SetState(429)
				p.InitMapArrStat()
			}

		case NeedleParserLPAREN, NeedleParserNOT, NeedleParserADD, NeedleParserSUB, NeedleParserTRUE, NeedleParserFALSE, NeedleParserNIL, NeedleParserIdentifier, NeedleParserDollarIdentifier, NeedleParserAtIdentifier, NeedleParserInterpretedStringLiteral, NeedleParserRawStringLiteral, NeedleParserDecimalLiteral, NeedleParserFloatLiteral, NeedleParserHexLiteral, NeedleParserOctalLiteral, NeedleParserBinaryLiteral:
			{
				p.SetState(430)
				p.expr(0)
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(437)
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

// ISimpleStatContext is an interface to support dynamic dispatch.
type ISimpleStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment() IAssignmentContext
	ExprStat() IExprStatContext
	IncDecStat() IIncDecStatContext
	AssignMapArrStat() IAssignMapArrStatContext

	// IsSimpleStatContext differentiates from other interfaces.
	IsSimpleStatContext()
}

type SimpleStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleStatContext() *SimpleStatContext {
	var p = new(SimpleStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_simpleStat
	return p
}

func InitEmptySimpleStatContext(p *SimpleStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_simpleStat
}

func (*SimpleStatContext) IsSimpleStatContext() {}

func NewSimpleStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleStatContext {
	var p = new(SimpleStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_simpleStat

	return p
}

func (s *SimpleStatContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleStatContext) Assignment() IAssignmentContext {
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

func (s *SimpleStatContext) ExprStat() IExprStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprStatContext)
}

func (s *SimpleStatContext) IncDecStat() IIncDecStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncDecStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncDecStatContext)
}

func (s *SimpleStatContext) AssignMapArrStat() IAssignMapArrStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignMapArrStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignMapArrStatContext)
}

func (s *SimpleStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterSimpleStat(s)
	}
}

func (s *SimpleStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitSimpleStat(s)
	}
}

func (p *NeedleParser) SimpleStat() (localctx ISimpleStatContext) {
	localctx = NewSimpleStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, NeedleParserRULE_simpleStat)
	p.SetState(442)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(438)
			p.Assignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(439)
			p.ExprStat()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(440)
			p.IncDecStat()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(441)
			p.AssignMapArrStat()
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

// IIncDecStatContext is an interface to support dynamic dispatch.
type IIncDecStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	IncDec_op() IIncDec_opContext

	// IsIncDecStatContext differentiates from other interfaces.
	IsIncDecStatContext()
}

type IncDecStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncDecStatContext() *IncDecStatContext {
	var p = new(IncDecStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_incDecStat
	return p
}

func InitEmptyIncDecStatContext(p *IncDecStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_incDecStat
}

func (*IncDecStatContext) IsIncDecStatContext() {}

func NewIncDecStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncDecStatContext {
	var p = new(IncDecStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_incDecStat

	return p
}

func (s *IncDecStatContext) GetParser() antlr.Parser { return s.parser }

func (s *IncDecStatContext) Expr() IExprContext {
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

func (s *IncDecStatContext) IncDec_op() IIncDec_opContext {
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

func (s *IncDecStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncDecStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncDecStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterIncDecStat(s)
	}
}

func (s *IncDecStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitIncDecStat(s)
	}
}

func (p *NeedleParser) IncDecStat() (localctx IIncDecStatContext) {
	localctx = NewIncDecStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, NeedleParserRULE_incDecStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(444)
		p.expr(0)
	}
	{
		p.SetState(445)
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

// IExprStatContext is an interface to support dynamic dispatch.
type IExprStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsExprStatContext differentiates from other interfaces.
	IsExprStatContext()
}

type ExprStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprStatContext() *ExprStatContext {
	var p = new(ExprStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_exprStat
	return p
}

func InitEmptyExprStatContext(p *ExprStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_exprStat
}

func (*ExprStatContext) IsExprStatContext() {}

func NewExprStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprStatContext {
	var p = new(ExprStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_exprStat

	return p
}

func (s *ExprStatContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprStatContext) Expr() IExprContext {
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

func (s *ExprStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterExprStat(s)
	}
}

func (s *ExprStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitExprStat(s)
	}
}

func (p *NeedleParser) ExprStat() (localctx IExprStatContext) {
	localctx = NewExprStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, NeedleParserRULE_exprStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(447)
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

// IAssignMapArrStatContext is an interface to support dynamic dispatch.
type IAssignMapArrStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExprList() IExprListContext
	EQ() antlr.TerminalNode
	InitMapArrStat() IInitMapArrStatContext

	// IsAssignMapArrStatContext differentiates from other interfaces.
	IsAssignMapArrStatContext()
}

type AssignMapArrStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignMapArrStatContext() *AssignMapArrStatContext {
	var p = new(AssignMapArrStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_assignMapArrStat
	return p
}

func InitEmptyAssignMapArrStatContext(p *AssignMapArrStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_assignMapArrStat
}

func (*AssignMapArrStatContext) IsAssignMapArrStatContext() {}

func NewAssignMapArrStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignMapArrStatContext {
	var p = new(AssignMapArrStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_assignMapArrStat

	return p
}

func (s *AssignMapArrStatContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignMapArrStatContext) ExprList() IExprListContext {
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

func (s *AssignMapArrStatContext) EQ() antlr.TerminalNode {
	return s.GetToken(NeedleParserEQ, 0)
}

func (s *AssignMapArrStatContext) InitMapArrStat() IInitMapArrStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInitMapArrStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInitMapArrStatContext)
}

func (s *AssignMapArrStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignMapArrStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignMapArrStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterAssignMapArrStat(s)
	}
}

func (s *AssignMapArrStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitAssignMapArrStat(s)
	}
}

func (p *NeedleParser) AssignMapArrStat() (localctx IAssignMapArrStatContext) {
	localctx = NewAssignMapArrStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, NeedleParserRULE_assignMapArrStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(449)
		p.ExprList()
	}
	{
		p.SetState(450)
		p.Match(NeedleParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(451)
		p.InitMapArrStat()
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

// IInitMapArrStatContext is an interface to support dynamic dispatch.
type IInitMapArrStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ObjectStat() IObjectStatContext
	ArrayStat() IArrayStatContext

	// IsInitMapArrStatContext differentiates from other interfaces.
	IsInitMapArrStatContext()
}

type InitMapArrStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitMapArrStatContext() *InitMapArrStatContext {
	var p = new(InitMapArrStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_initMapArrStat
	return p
}

func InitEmptyInitMapArrStatContext(p *InitMapArrStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_initMapArrStat
}

func (*InitMapArrStatContext) IsInitMapArrStatContext() {}

func NewInitMapArrStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitMapArrStatContext {
	var p = new(InitMapArrStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_initMapArrStat

	return p
}

func (s *InitMapArrStatContext) GetParser() antlr.Parser { return s.parser }

func (s *InitMapArrStatContext) ObjectStat() IObjectStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectStatContext)
}

func (s *InitMapArrStatContext) ArrayStat() IArrayStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayStatContext)
}

func (s *InitMapArrStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitMapArrStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitMapArrStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterInitMapArrStat(s)
	}
}

func (s *InitMapArrStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitInitMapArrStat(s)
	}
}

func (p *NeedleParser) InitMapArrStat() (localctx IInitMapArrStatContext) {
	localctx = NewInitMapArrStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, NeedleParserRULE_initMapArrStat)
	p.SetState(455)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NeedleParserLBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(453)
			p.ObjectStat()
		}

	case NeedleParserLBRACK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(454)
			p.ArrayStat()
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
	p.EnterRule(localctx, 86, NeedleParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(457)
		p.ExprList()
	}
	{
		p.SetState(458)
		p.Assign_op()
	}
	{
		p.SetState(459)
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

// IPrimaryExprContext is an interface to support dynamic dispatch.
type IPrimaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Operand() IOperandContext
	PrimaryExpr() IPrimaryExprContext
	Arguments() IArgumentsContext
	SliceStat() ISliceStatContext
	IndexStat() IIndexStatContext
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

func (s *PrimaryExprContext) SliceStat() ISliceStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceStatContext)
}

func (s *PrimaryExprContext) IndexStat() IIndexStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexStatContext)
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
	_startState := 88
	p.EnterRecursionRule(localctx, 88, NeedleParserRULE_primaryExpr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(462)
		p.Operand()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(476)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 52, p.GetParserRuleContext())
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
			p.SetState(464)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			p.SetState(472)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 51, p.GetParserRuleContext()) {
			case 1:
				p.SetState(467)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == NeedleParserDOT {
					{
						p.SetState(465)
						p.Match(NeedleParserDOT)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					{
						p.SetState(466)
						p.Match(NeedleParserIdentifier)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(469)
					p.Arguments()
				}

			case 2:
				{
					p.SetState(470)
					p.SliceStat()
				}

			case 3:
				{
					p.SetState(471)
					p.IndexStat()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(478)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 52, p.GetParserRuleContext())
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

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentifierFull() IIdentifierFullContext
	NumberLiteral() INumberLiteralContext
	StringLiteral() IStringLiteralContext
	BooleanLiteral() IBooleanLiteralContext
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode
	NIL() antlr.TerminalNode

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

func (s *OperandContext) NumberLiteral() INumberLiteralContext {
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

func (s *OperandContext) StringLiteral() IStringLiteralContext {
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

func (s *OperandContext) BooleanLiteral() IBooleanLiteralContext {
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

func (s *OperandContext) NIL() antlr.TerminalNode {
	return s.GetToken(NeedleParserNIL, 0)
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
	p.EnterRule(localctx, 90, NeedleParserRULE_operand)
	p.SetState(488)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NeedleParserIdentifier, NeedleParserDollarIdentifier, NeedleParserAtIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(479)
			p.IdentifierFull()
		}

	case NeedleParserDecimalLiteral, NeedleParserFloatLiteral, NeedleParserHexLiteral, NeedleParserOctalLiteral, NeedleParserBinaryLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(480)
			p.NumberLiteral()
		}

	case NeedleParserInterpretedStringLiteral, NeedleParserRawStringLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(481)
			p.StringLiteral()
		}

	case NeedleParserTRUE, NeedleParserFALSE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(482)
			p.BooleanLiteral()
		}

	case NeedleParserLPAREN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(483)
			p.Match(NeedleParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(484)
			p.expr(0)
		}
		{
			p.SetState(485)
			p.Match(NeedleParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case NeedleParserNIL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(487)
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
	p.EnterRule(localctx, 92, NeedleParserRULE_exprList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(490)
		p.expr(0)
	}
	p.SetState(495)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == NeedleParserCOMMA {
		{
			p.SetState(491)
			p.Match(NeedleParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(492)
			p.expr(0)
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimaryExpr() IPrimaryExprContext
	Eos() IEosContext
	Unary_op() IUnary_opContext
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	Mul_op() IMul_opContext
	Rel_op() IRel_opContext
	Logical_op() ILogical_opContext
	Add_op() IAdd_opContext

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
	_startState := 94
	p.EnterRecursionRule(localctx, 94, NeedleParserRULE_expr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(505)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NeedleParserLPAREN, NeedleParserTRUE, NeedleParserFALSE, NeedleParserNIL, NeedleParserIdentifier, NeedleParserDollarIdentifier, NeedleParserAtIdentifier, NeedleParserInterpretedStringLiteral, NeedleParserRawStringLiteral, NeedleParserDecimalLiteral, NeedleParserFloatLiteral, NeedleParserHexLiteral, NeedleParserOctalLiteral, NeedleParserBinaryLiteral:
		{
			p.SetState(499)
			p.primaryExpr(0)
		}
		{
			p.SetState(500)
			p.Eos()
		}

	case NeedleParserNOT, NeedleParserADD, NeedleParserSUB:
		{
			p.SetState(502)
			p.Unary_op()
		}
		{
			p.SetState(503)
			p.expr(5)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(525)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(523)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NeedleParserRULE_expr)
				p.SetState(507)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(508)
					p.Mul_op()
				}
				{
					p.SetState(509)
					p.expr(5)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NeedleParserRULE_expr)
				p.SetState(511)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(512)
					p.Rel_op()
				}
				{
					p.SetState(513)
					p.expr(4)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NeedleParserRULE_expr)
				p.SetState(515)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(516)
					p.Logical_op()
				}
				{
					p.SetState(517)
					p.expr(3)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NeedleParserRULE_expr)
				p.SetState(519)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(520)
					p.Add_op()
				}
				{
					p.SetState(521)
					p.expr(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(527)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 96, NeedleParserRULE_incDec_op)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(528)
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
	p.EnterRule(localctx, 98, NeedleParserRULE_mul_op)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(530)
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
	p.EnterRule(localctx, 100, NeedleParserRULE_unary_op)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(532)
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
	p.EnterRule(localctx, 102, NeedleParserRULE_add_op)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(534)
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
	p.EnterRule(localctx, 104, NeedleParserRULE_logical_op)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(536)
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
	p.EnterRule(localctx, 106, NeedleParserRULE_rel_op)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(538)
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
	p.EnterRule(localctx, 108, NeedleParserRULE_assign_op)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(540)
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
	p.EnterRule(localctx, 110, NeedleParserRULE_identifierFull)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(542)
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
	p.EnterRule(localctx, 112, NeedleParserRULE_identifierVar)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(544)
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
	p.EnterRule(localctx, 114, NeedleParserRULE_identifierList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(546)
		p.Match(NeedleParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(553)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == NeedleParserCOMMA || _la == NeedleParserIdentifier {
		p.SetState(548)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == NeedleParserCOMMA {
			{
				p.SetState(547)
				p.Match(NeedleParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(550)
			p.Match(NeedleParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(555)
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
	p.EnterRule(localctx, 116, NeedleParserRULE_stringLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(556)
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
	p.EnterRule(localctx, 118, NeedleParserRULE_numberLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(558)
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
	p.EnterRule(localctx, 120, NeedleParserRULE_booleanLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(560)
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
	p.EnterRule(localctx, 122, NeedleParserRULE_eos)
	p.SetState(567)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 61, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(562)
			p.Match(NeedleParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(563)
			p.Match(NeedleParserEOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(565)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 60, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(564)
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
	case 44:
		var t *PrimaryExprContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryExprContext)
		}
		return p.PrimaryExpr_Sempred(t, predIndex)

	case 47:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *NeedleParser) PrimaryExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *NeedleParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
