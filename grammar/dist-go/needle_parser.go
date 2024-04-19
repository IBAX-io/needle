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
		"statementList", "statement", "simpleStmt", "incDecStmt", "assignMapArrStmt",
		"initMapArrStmt", "assignment", "varDef", "ifStmt", "ifBody", "elseBody",
		"returnStmt", "continueStmt", "breakStmt", "whileStmt", "errorStmt",
		"arrayStmt", "arrayList", "arrayValue", "mapStmt", "pairList", "pair",
		"pairValue", "arguments", "argumentsList", "exprList", "expr", "primaryExpr",
		"indexExpr", "sliceExpr", "operand", "literal", "typeName", "incDec_op",
		"mul_op", "unary_op", "add_op", "logical_op", "rel_op", "assign_op",
		"identifierFull", "identifierVar", "identifierList", "stringLiteral",
		"numberLiteral", "booleanLiteral", "eos",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 99, 548, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
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
		3, 6, 189, 8, 6, 1, 6, 1, 6, 1, 6, 1, 7, 3, 7, 195, 8, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 8, 1, 9, 3, 9, 203, 8, 9, 1, 9, 5, 9, 206, 8, 9, 10, 9, 12,
		9, 209, 9, 9, 1, 9, 3, 9, 212, 8, 9, 1, 10, 1, 10, 1, 10, 3, 10, 217, 8,
		10, 1, 11, 1, 11, 1, 11, 3, 11, 222, 8, 11, 3, 11, 224, 8, 11, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 12, 3, 12, 231, 8, 12, 1, 12, 1, 12, 1, 12, 5, 12,
		236, 8, 12, 10, 12, 12, 12, 239, 9, 12, 1, 13, 1, 13, 3, 13, 243, 8, 13,
		1, 13, 5, 13, 246, 8, 13, 10, 13, 12, 13, 249, 9, 13, 1, 14, 1, 14, 3,
		14, 253, 8, 14, 1, 14, 1, 14, 1, 15, 3, 15, 258, 8, 15, 1, 15, 3, 15, 261,
		8, 15, 3, 15, 263, 8, 15, 1, 15, 1, 15, 1, 15, 4, 15, 268, 8, 15, 11, 15,
		12, 15, 269, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 3, 16, 281, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 287, 8, 17, 1,
		18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 3, 20, 298,
		8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		23, 1, 23, 5, 23, 311, 8, 23, 10, 23, 12, 23, 314, 9, 23, 1, 23, 3, 23,
		317, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 324, 8, 24, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 3, 26, 333, 8, 26, 1, 27, 1,
		27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 31,
		1, 31, 3, 31, 348, 8, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 5, 32, 355,
		8, 32, 10, 32, 12, 32, 358, 9, 32, 1, 32, 1, 32, 1, 33, 1, 33, 3, 33, 364,
		8, 33, 1, 34, 1, 34, 3, 34, 368, 8, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1,
		35, 5, 35, 375, 8, 35, 10, 35, 12, 35, 378, 9, 35, 1, 35, 3, 35, 381, 8,
		35, 1, 35, 1, 35, 1, 36, 1, 36, 3, 36, 387, 8, 36, 1, 36, 1, 36, 1, 36,
		1, 37, 1, 37, 1, 37, 3, 37, 395, 8, 37, 1, 37, 1, 37, 3, 37, 399, 8, 37,
		1, 38, 1, 38, 3, 38, 403, 8, 38, 1, 38, 1, 38, 1, 39, 1, 39, 3, 39, 409,
		8, 39, 1, 39, 1, 39, 1, 39, 3, 39, 414, 8, 39, 5, 39, 416, 8, 39, 10, 39,
		12, 39, 419, 9, 39, 1, 40, 1, 40, 1, 40, 5, 40, 424, 8, 40, 10, 40, 12,
		40, 427, 9, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 3, 41,
		436, 8, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1,
		41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41,
		1, 41, 5, 41, 458, 8, 41, 10, 41, 12, 41, 461, 9, 41, 1, 42, 1, 42, 1,
		42, 1, 42, 1, 42, 1, 42, 3, 42, 469, 8, 42, 1, 42, 5, 42, 472, 8, 42, 10,
		42, 12, 42, 475, 9, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 3, 44,
		483, 8, 44, 1, 44, 1, 44, 3, 44, 487, 8, 44, 1, 44, 1, 44, 1, 45, 1, 45,
		1, 45, 1, 45, 1, 45, 1, 45, 3, 45, 497, 8, 45, 1, 46, 1, 46, 1, 46, 1,
		46, 3, 46, 503, 8, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 50,
		1, 50, 1, 51, 1, 51, 1, 52, 1, 52, 1, 53, 1, 53, 1, 54, 1, 54, 1, 55, 1,
		55, 1, 56, 1, 56, 1, 57, 1, 57, 3, 57, 527, 8, 57, 1, 57, 5, 57, 530, 8,
		57, 10, 57, 12, 57, 533, 9, 57, 1, 58, 1, 58, 1, 59, 1, 59, 1, 60, 1, 60,
		1, 61, 1, 61, 1, 61, 3, 61, 544, 8, 61, 3, 61, 546, 8, 61, 1, 61, 1, 312,
		2, 82, 84, 62, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66,
		68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102,
		104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 0, 15, 1, 0, 60, 61,
		2, 0, 57, 58, 63, 63, 1, 0, 64, 73, 1, 0, 41, 42, 4, 0, 13, 13, 16, 16,
		25, 25, 28, 30, 2, 0, 12, 12, 14, 15, 2, 0, 14, 15, 26, 27, 2, 0, 20, 20,
		24, 24, 2, 0, 17, 19, 21, 23, 2, 0, 6, 6, 31, 40, 1, 0, 74, 76, 1, 0, 74,
		75, 1, 0, 77, 78, 1, 0, 79, 83, 1, 0, 50, 51, 563, 0, 132, 1, 0, 0, 0,
		2, 137, 1, 0, 0, 0, 4, 153, 1, 0, 0, 0, 6, 155, 1, 0, 0, 0, 8, 167, 1,
		0, 0, 0, 10, 172, 1, 0, 0, 0, 12, 188, 1, 0, 0, 0, 14, 194, 1, 0, 0, 0,
		16, 198, 1, 0, 0, 0, 18, 202, 1, 0, 0, 0, 20, 213, 1, 0, 0, 0, 22, 218,
		1, 0, 0, 0, 24, 227, 1, 0, 0, 0, 26, 240, 1, 0, 0, 0, 28, 250, 1, 0, 0,
		0, 30, 267, 1, 0, 0, 0, 32, 280, 1, 0, 0, 0, 34, 286, 1, 0, 0, 0, 36, 288,
		1, 0, 0, 0, 38, 291, 1, 0, 0, 0, 40, 297, 1, 0, 0, 0, 42, 299, 1, 0, 0,
		0, 44, 303, 1, 0, 0, 0, 46, 306, 1, 0, 0, 0, 48, 323, 1, 0, 0, 0, 50, 327,
		1, 0, 0, 0, 52, 330, 1, 0, 0, 0, 54, 334, 1, 0, 0, 0, 56, 336, 1, 0, 0,
		0, 58, 338, 1, 0, 0, 0, 60, 342, 1, 0, 0, 0, 62, 345, 1, 0, 0, 0, 64, 351,
		1, 0, 0, 0, 66, 363, 1, 0, 0, 0, 68, 365, 1, 0, 0, 0, 70, 371, 1, 0, 0,
		0, 72, 386, 1, 0, 0, 0, 74, 398, 1, 0, 0, 0, 76, 400, 1, 0, 0, 0, 78, 408,
		1, 0, 0, 0, 80, 420, 1, 0, 0, 0, 82, 435, 1, 0, 0, 0, 84, 462, 1, 0, 0,
		0, 86, 476, 1, 0, 0, 0, 88, 480, 1, 0, 0, 0, 90, 496, 1, 0, 0, 0, 92, 502,
		1, 0, 0, 0, 94, 504, 1, 0, 0, 0, 96, 506, 1, 0, 0, 0, 98, 508, 1, 0, 0,
		0, 100, 510, 1, 0, 0, 0, 102, 512, 1, 0, 0, 0, 104, 514, 1, 0, 0, 0, 106,
		516, 1, 0, 0, 0, 108, 518, 1, 0, 0, 0, 110, 520, 1, 0, 0, 0, 112, 522,
		1, 0, 0, 0, 114, 524, 1, 0, 0, 0, 116, 534, 1, 0, 0, 0, 118, 536, 1, 0,
		0, 0, 120, 538, 1, 0, 0, 0, 122, 545, 1, 0, 0, 0, 124, 127, 3, 2, 1, 0,
		125, 127, 3, 12, 6, 0, 126, 124, 1, 0, 0, 0, 126, 125, 1, 0, 0, 0, 127,
		128, 1, 0, 0, 0, 128, 129, 3, 122, 61, 0, 129, 131, 1, 0, 0, 0, 130, 126,
		1, 0, 0, 0, 131, 134, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 132, 133, 1, 0,
		0, 0, 133, 135, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 135, 136, 5, 0, 0, 1,
		136, 1, 1, 0, 0, 0, 137, 138, 5, 43, 0, 0, 138, 139, 5, 74, 0, 0, 139,
		145, 5, 9, 0, 0, 140, 141, 3, 4, 2, 0, 141, 142, 3, 122, 61, 0, 142, 144,
		1, 0, 0, 0, 143, 140, 1, 0, 0, 0, 144, 147, 1, 0, 0, 0, 145, 143, 1, 0,
		0, 0, 145, 146, 1, 0, 0, 0, 146, 148, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0,
		148, 149, 5, 10, 0, 0, 149, 3, 1, 0, 0, 0, 150, 154, 3, 6, 3, 0, 151, 154,
		3, 10, 5, 0, 152, 154, 3, 12, 6, 0, 153, 150, 1, 0, 0, 0, 153, 151, 1,
		0, 0, 0, 153, 152, 1, 0, 0, 0, 154, 5, 1, 0, 0, 0, 155, 156, 5, 53, 0,
		0, 156, 162, 5, 9, 0, 0, 157, 158, 3, 8, 4, 0, 158, 159, 3, 122, 61, 0,
		159, 161, 1, 0, 0, 0, 160, 157, 1, 0, 0, 0, 161, 164, 1, 0, 0, 0, 162,
		160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 165, 1, 0, 0, 0, 164, 162,
		1, 0, 0, 0, 165, 166, 5, 10, 0, 0, 166, 7, 1, 0, 0, 0, 167, 168, 5, 74,
		0, 0, 168, 170, 3, 94, 47, 0, 169, 171, 3, 116, 58, 0, 170, 169, 1, 0,
		0, 0, 170, 171, 1, 0, 0, 0, 171, 9, 1, 0, 0, 0, 172, 173, 5, 54, 0, 0,
		173, 181, 5, 9, 0, 0, 174, 175, 5, 74, 0, 0, 175, 176, 5, 6, 0, 0, 176,
		177, 3, 92, 46, 0, 177, 178, 3, 122, 61, 0, 178, 180, 1, 0, 0, 0, 179,
		174, 1, 0, 0, 0, 180, 183, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 181, 182,
		1, 0, 0, 0, 182, 184, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 184, 185, 5, 10,
		0, 0, 185, 11, 1, 0, 0, 0, 186, 189, 3, 16, 8, 0, 187, 189, 3, 14, 7, 0,
		188, 186, 1, 0, 0, 0, 188, 187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190,
		191, 3, 18, 9, 0, 191, 192, 3, 28, 14, 0, 192, 13, 1, 0, 0, 0, 193, 195,
		5, 44, 0, 0, 194, 193, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 196, 1, 0,
		0, 0, 196, 197, 7, 0, 0, 0, 197, 15, 1, 0, 0, 0, 198, 199, 5, 44, 0, 0,
		199, 200, 5, 74, 0, 0, 200, 17, 1, 0, 0, 0, 201, 203, 3, 22, 11, 0, 202,
		201, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 207, 1, 0, 0, 0, 204, 206,
		3, 20, 10, 0, 205, 204, 1, 0, 0, 0, 206, 209, 1, 0, 0, 0, 207, 205, 1,
		0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 211, 1, 0, 0, 0, 209, 207, 1, 0, 0,
		0, 210, 212, 3, 26, 13, 0, 211, 210, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0,
		212, 19, 1, 0, 0, 0, 213, 214, 5, 4, 0, 0, 214, 216, 5, 74, 0, 0, 215,
		217, 3, 22, 11, 0, 216, 215, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 21,
		1, 0, 0, 0, 218, 223, 5, 1, 0, 0, 219, 221, 3, 24, 12, 0, 220, 222, 5,
		3, 0, 0, 221, 220, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 224, 1, 0, 0,
		0, 223, 219, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225,
		226, 5, 2, 0, 0, 226, 23, 1, 0, 0, 0, 227, 228, 3, 114, 57, 0, 228, 237,
		3, 94, 47, 0, 229, 231, 5, 3, 0, 0, 230, 229, 1, 0, 0, 0, 230, 231, 1,
		0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 233, 3, 114, 57, 0, 233, 234, 3, 94,
		47, 0, 234, 236, 1, 0, 0, 0, 235, 230, 1, 0, 0, 0, 236, 239, 1, 0, 0, 0,
		237, 235, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 25, 1, 0, 0, 0, 239, 237,
		1, 0, 0, 0, 240, 247, 3, 94, 47, 0, 241, 243, 5, 3, 0, 0, 242, 241, 1,
		0, 0, 0, 242, 243, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244, 246, 3, 94, 47,
		0, 245, 242, 1, 0, 0, 0, 246, 249, 1, 0, 0, 0, 247, 245, 1, 0, 0, 0, 247,
		248, 1, 0, 0, 0, 248, 27, 1, 0, 0, 0, 249, 247, 1, 0, 0, 0, 250, 252, 5,
		9, 0, 0, 251, 253, 3, 30, 15, 0, 252, 251, 1, 0, 0, 0, 252, 253, 1, 0,
		0, 0, 253, 254, 1, 0, 0, 0, 254, 255, 5, 10, 0, 0, 255, 29, 1, 0, 0, 0,
		256, 258, 5, 11, 0, 0, 257, 256, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258,
		263, 1, 0, 0, 0, 259, 261, 5, 98, 0, 0, 260, 259, 1, 0, 0, 0, 260, 261,
		1, 0, 0, 0, 261, 263, 1, 0, 0, 0, 262, 257, 1, 0, 0, 0, 262, 260, 1, 0,
		0, 0, 263, 264, 1, 0, 0, 0, 264, 265, 3, 32, 16, 0, 265, 266, 3, 122, 61,
		0, 266, 268, 1, 0, 0, 0, 267, 262, 1, 0, 0, 0, 268, 269, 1, 0, 0, 0, 269,
		267, 1, 0, 0, 0, 269, 270, 1, 0, 0, 0, 270, 31, 1, 0, 0, 0, 271, 281, 3,
		28, 14, 0, 272, 281, 3, 34, 17, 0, 273, 281, 3, 44, 22, 0, 274, 281, 3,
		46, 23, 0, 275, 281, 3, 58, 29, 0, 276, 281, 3, 54, 27, 0, 277, 281, 3,
		56, 28, 0, 278, 281, 3, 52, 26, 0, 279, 281, 3, 60, 30, 0, 280, 271, 1,
		0, 0, 0, 280, 272, 1, 0, 0, 0, 280, 273, 1, 0, 0, 0, 280, 274, 1, 0, 0,
		0, 280, 275, 1, 0, 0, 0, 280, 276, 1, 0, 0, 0, 280, 277, 1, 0, 0, 0, 280,
		278, 1, 0, 0, 0, 280, 279, 1, 0, 0, 0, 281, 33, 1, 0, 0, 0, 282, 287, 3,
		82, 41, 0, 283, 287, 3, 42, 21, 0, 284, 287, 3, 36, 18, 0, 285, 287, 3,
		38, 19, 0, 286, 282, 1, 0, 0, 0, 286, 283, 1, 0, 0, 0, 286, 284, 1, 0,
		0, 0, 286, 285, 1, 0, 0, 0, 287, 35, 1, 0, 0, 0, 288, 289, 3, 82, 41, 0,
		289, 290, 3, 96, 48, 0, 290, 37, 1, 0, 0, 0, 291, 292, 3, 112, 56, 0, 292,
		293, 5, 6, 0, 0, 293, 294, 3, 40, 20, 0, 294, 39, 1, 0, 0, 0, 295, 298,
		3, 68, 34, 0, 296, 298, 3, 62, 31, 0, 297, 295, 1, 0, 0, 0, 297, 296, 1,
		0, 0, 0, 298, 41, 1, 0, 0, 0, 299, 300, 3, 80, 40, 0, 300, 301, 3, 108,
		54, 0, 301, 302, 3, 80, 40, 0, 302, 43, 1, 0, 0, 0, 303, 304, 5, 52, 0,
		0, 304, 305, 3, 24, 12, 0, 305, 45, 1, 0, 0, 0, 306, 307, 5, 46, 0, 0,
		307, 312, 3, 48, 24, 0, 308, 309, 5, 47, 0, 0, 309, 311, 3, 48, 24, 0,
		310, 308, 1, 0, 0, 0, 311, 314, 1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 312,
		310, 1, 0, 0, 0, 313, 316, 1, 0, 0, 0, 314, 312, 1, 0, 0, 0, 315, 317,
		3, 50, 25, 0, 316, 315, 1, 0, 0, 0, 316, 317, 1, 0, 0, 0, 317, 47, 1, 0,
		0, 0, 318, 319, 5, 1, 0, 0, 319, 320, 3, 82, 41, 0, 320, 321, 5, 2, 0,
		0, 321, 324, 1, 0, 0, 0, 322, 324, 3, 82, 41, 0, 323, 318, 1, 0, 0, 0,
		323, 322, 1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325, 326, 3, 28, 14, 0, 326,
		49, 1, 0, 0, 0, 327, 328, 5, 48, 0, 0, 328, 329, 3, 28, 14, 0, 329, 51,
		1, 0, 0, 0, 330, 332, 5, 45, 0, 0, 331, 333, 3, 82, 41, 0, 332, 331, 1,
		0, 0, 0, 332, 333, 1, 0, 0, 0, 333, 53, 1, 0, 0, 0, 334, 335, 5, 56, 0,
		0, 335, 55, 1, 0, 0, 0, 336, 337, 5, 55, 0, 0, 337, 57, 1, 0, 0, 0, 338,
		339, 5, 49, 0, 0, 339, 340, 3, 82, 41, 0, 340, 341, 3, 28, 14, 0, 341,
		59, 1, 0, 0, 0, 342, 343, 7, 1, 0, 0, 343, 344, 3, 82, 41, 0, 344, 61,
		1, 0, 0, 0, 345, 347, 5, 7, 0, 0, 346, 348, 3, 64, 32, 0, 347, 346, 1,
		0, 0, 0, 347, 348, 1, 0, 0, 0, 348, 349, 1, 0, 0, 0, 349, 350, 5, 8, 0,
		0, 350, 63, 1, 0, 0, 0, 351, 356, 3, 66, 33, 0, 352, 353, 5, 3, 0, 0, 353,
		355, 3, 66, 33, 0, 354, 352, 1, 0, 0, 0, 355, 358, 1, 0, 0, 0, 356, 354,
		1, 0, 0, 0, 356, 357, 1, 0, 0, 0, 357, 359, 1, 0, 0, 0, 358, 356, 1, 0,
		0, 0, 359, 360, 3, 122, 61, 0, 360, 65, 1, 0, 0, 0, 361, 364, 3, 82, 41,
		0, 362, 364, 3, 40, 20, 0, 363, 361, 1, 0, 0, 0, 363, 362, 1, 0, 0, 0,
		364, 67, 1, 0, 0, 0, 365, 367, 5, 9, 0, 0, 366, 368, 3, 70, 35, 0, 367,
		366, 1, 0, 0, 0, 367, 368, 1, 0, 0, 0, 368, 369, 1, 0, 0, 0, 369, 370,
		5, 10, 0, 0, 370, 69, 1, 0, 0, 0, 371, 376, 3, 72, 36, 0, 372, 373, 5,
		3, 0, 0, 373, 375, 3, 72, 36, 0, 374, 372, 1, 0, 0, 0, 375, 378, 1, 0,
		0, 0, 376, 374, 1, 0, 0, 0, 376, 377, 1, 0, 0, 0, 377, 380, 1, 0, 0, 0,
		378, 376, 1, 0, 0, 0, 379, 381, 5, 3, 0, 0, 380, 379, 1, 0, 0, 0, 380,
		381, 1, 0, 0, 0, 381, 382, 1, 0, 0, 0, 382, 383, 3, 122, 61, 0, 383, 71,
		1, 0, 0, 0, 384, 387, 3, 116, 58, 0, 385, 387, 3, 112, 56, 0, 386, 384,
		1, 0, 0, 0, 386, 385, 1, 0, 0, 0, 387, 388, 1, 0, 0, 0, 388, 389, 5, 5,
		0, 0, 389, 390, 3, 74, 37, 0, 390, 73, 1, 0, 0, 0, 391, 394, 3, 112, 56,
		0, 392, 395, 3, 86, 43, 0, 393, 395, 3, 88, 44, 0, 394, 392, 1, 0, 0, 0,
		394, 393, 1, 0, 0, 0, 394, 395, 1, 0, 0, 0, 395, 399, 1, 0, 0, 0, 396,
		399, 3, 92, 46, 0, 397, 399, 3, 40, 20, 0, 398, 391, 1, 0, 0, 0, 398, 396,
		1, 0, 0, 0, 398, 397, 1, 0, 0, 0, 399, 75, 1, 0, 0, 0, 400, 402, 5, 1,
		0, 0, 401, 403, 3, 78, 39, 0, 402, 401, 1, 0, 0, 0, 402, 403, 1, 0, 0,
		0, 403, 404, 1, 0, 0, 0, 404, 405, 5, 2, 0, 0, 405, 77, 1, 0, 0, 0, 406,
		409, 3, 40, 20, 0, 407, 409, 3, 82, 41, 0, 408, 406, 1, 0, 0, 0, 408, 407,
		1, 0, 0, 0, 409, 417, 1, 0, 0, 0, 410, 413, 5, 3, 0, 0, 411, 414, 3, 40,
		20, 0, 412, 414, 3, 82, 41, 0, 413, 411, 1, 0, 0, 0, 413, 412, 1, 0, 0,
		0, 414, 416, 1, 0, 0, 0, 415, 410, 1, 0, 0, 0, 416, 419, 1, 0, 0, 0, 417,
		415, 1, 0, 0, 0, 417, 418, 1, 0, 0, 0, 418, 79, 1, 0, 0, 0, 419, 417, 1,
		0, 0, 0, 420, 425, 3, 82, 41, 0, 421, 422, 5, 3, 0, 0, 422, 424, 3, 82,
		41, 0, 423, 421, 1, 0, 0, 0, 424, 427, 1, 0, 0, 0, 425, 423, 1, 0, 0, 0,
		425, 426, 1, 0, 0, 0, 426, 81, 1, 0, 0, 0, 427, 425, 1, 0, 0, 0, 428, 429,
		6, 41, -1, 0, 429, 430, 3, 84, 42, 0, 430, 431, 3, 122, 61, 0, 431, 436,
		1, 0, 0, 0, 432, 433, 3, 100, 50, 0, 433, 434, 3, 82, 41, 5, 434, 436,
		1, 0, 0, 0, 435, 428, 1, 0, 0, 0, 435, 432, 1, 0, 0, 0, 436, 459, 1, 0,
		0, 0, 437, 438, 10, 4, 0, 0, 438, 439, 3, 98, 49, 0, 439, 440, 3, 82, 41,
		5, 440, 458, 1, 0, 0, 0, 441, 442, 10, 3, 0, 0, 442, 443, 3, 106, 53, 0,
		443, 444, 3, 82, 41, 4, 444, 458, 1, 0, 0, 0, 445, 446, 10, 2, 0, 0, 446,
		447, 3, 104, 52, 0, 447, 448, 3, 82, 41, 3, 448, 458, 1, 0, 0, 0, 449,
		450, 10, 1, 0, 0, 450, 451, 3, 102, 51, 0, 451, 452, 3, 82, 41, 2, 452,
		458, 1, 0, 0, 0, 453, 454, 10, 7, 0, 0, 454, 458, 3, 86, 43, 0, 455, 456,
		10, 6, 0, 0, 456, 458, 3, 88, 44, 0, 457, 437, 1, 0, 0, 0, 457, 441, 1,
		0, 0, 0, 457, 445, 1, 0, 0, 0, 457, 449, 1, 0, 0, 0, 457, 453, 1, 0, 0,
		0, 457, 455, 1, 0, 0, 0, 458, 461, 1, 0, 0, 0, 459, 457, 1, 0, 0, 0, 459,
		460, 1, 0, 0, 0, 460, 83, 1, 0, 0, 0, 461, 459, 1, 0, 0, 0, 462, 463, 6,
		42, -1, 0, 463, 464, 3, 90, 45, 0, 464, 473, 1, 0, 0, 0, 465, 468, 10,
		1, 0, 0, 466, 467, 5, 4, 0, 0, 467, 469, 5, 74, 0, 0, 468, 466, 1, 0, 0,
		0, 468, 469, 1, 0, 0, 0, 469, 470, 1, 0, 0, 0, 470, 472, 3, 76, 38, 0,
		471, 465, 1, 0, 0, 0, 472, 475, 1, 0, 0, 0, 473, 471, 1, 0, 0, 0, 473,
		474, 1, 0, 0, 0, 474, 85, 1, 0, 0, 0, 475, 473, 1, 0, 0, 0, 476, 477, 5,
		7, 0, 0, 477, 478, 3, 82, 41, 0, 478, 479, 5, 8, 0, 0, 479, 87, 1, 0, 0,
		0, 480, 482, 5, 7, 0, 0, 481, 483, 3, 82, 41, 0, 482, 481, 1, 0, 0, 0,
		482, 483, 1, 0, 0, 0, 483, 484, 1, 0, 0, 0, 484, 486, 5, 5, 0, 0, 485,
		487, 3, 82, 41, 0, 486, 485, 1, 0, 0, 0, 486, 487, 1, 0, 0, 0, 487, 488,
		1, 0, 0, 0, 488, 489, 5, 8, 0, 0, 489, 89, 1, 0, 0, 0, 490, 497, 3, 110,
		55, 0, 491, 497, 3, 92, 46, 0, 492, 493, 5, 1, 0, 0, 493, 494, 3, 82, 41,
		0, 494, 495, 5, 2, 0, 0, 495, 497, 1, 0, 0, 0, 496, 490, 1, 0, 0, 0, 496,
		491, 1, 0, 0, 0, 496, 492, 1, 0, 0, 0, 497, 91, 1, 0, 0, 0, 498, 503, 3,
		118, 59, 0, 499, 503, 3, 116, 58, 0, 500, 503, 3, 120, 60, 0, 501, 503,
		5, 59, 0, 0, 502, 498, 1, 0, 0, 0, 502, 499, 1, 0, 0, 0, 502, 500, 1, 0,
		0, 0, 502, 501, 1, 0, 0, 0, 503, 93, 1, 0, 0, 0, 504, 505, 7, 2, 0, 0,
		505, 95, 1, 0, 0, 0, 506, 507, 7, 3, 0, 0, 507, 97, 1, 0, 0, 0, 508, 509,
		7, 4, 0, 0, 509, 99, 1, 0, 0, 0, 510, 511, 7, 5, 0, 0, 511, 101, 1, 0,
		0, 0, 512, 513, 7, 6, 0, 0, 513, 103, 1, 0, 0, 0, 514, 515, 7, 7, 0, 0,
		515, 105, 1, 0, 0, 0, 516, 517, 7, 8, 0, 0, 517, 107, 1, 0, 0, 0, 518,
		519, 7, 9, 0, 0, 519, 109, 1, 0, 0, 0, 520, 521, 7, 10, 0, 0, 521, 111,
		1, 0, 0, 0, 522, 523, 7, 11, 0, 0, 523, 113, 1, 0, 0, 0, 524, 531, 5, 74,
		0, 0, 525, 527, 5, 3, 0, 0, 526, 525, 1, 0, 0, 0, 526, 527, 1, 0, 0, 0,
		527, 528, 1, 0, 0, 0, 528, 530, 5, 74, 0, 0, 529, 526, 1, 0, 0, 0, 530,
		533, 1, 0, 0, 0, 531, 529, 1, 0, 0, 0, 531, 532, 1, 0, 0, 0, 532, 115,
		1, 0, 0, 0, 533, 531, 1, 0, 0, 0, 534, 535, 7, 12, 0, 0, 535, 117, 1, 0,
		0, 0, 536, 537, 7, 13, 0, 0, 537, 119, 1, 0, 0, 0, 538, 539, 7, 14, 0,
		0, 539, 121, 1, 0, 0, 0, 540, 546, 5, 11, 0, 0, 541, 546, 5, 98, 0, 0,
		542, 544, 5, 0, 0, 1, 543, 542, 1, 0, 0, 0, 543, 544, 1, 0, 0, 0, 544,
		546, 1, 0, 0, 0, 545, 540, 1, 0, 0, 0, 545, 541, 1, 0, 0, 0, 545, 543,
		1, 0, 0, 0, 546, 123, 1, 0, 0, 0, 58, 126, 132, 145, 153, 162, 170, 181,
		188, 194, 202, 207, 211, 216, 221, 223, 230, 237, 242, 247, 252, 257, 260,
		262, 269, 280, 286, 297, 312, 316, 323, 332, 347, 356, 363, 367, 376, 380,
		386, 394, 398, 402, 408, 413, 417, 425, 435, 457, 459, 468, 473, 482, 486,
		496, 502, 526, 531, 543, 545,
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
	NeedleParserRULE_assignMapArrStmt = 19
	NeedleParserRULE_initMapArrStmt   = 20
	NeedleParserRULE_assignment       = 21
	NeedleParserRULE_varDef           = 22
	NeedleParserRULE_ifStmt           = 23
	NeedleParserRULE_ifBody           = 24
	NeedleParserRULE_elseBody         = 25
	NeedleParserRULE_returnStmt       = 26
	NeedleParserRULE_continueStmt     = 27
	NeedleParserRULE_breakStmt        = 28
	NeedleParserRULE_whileStmt        = 29
	NeedleParserRULE_errorStmt        = 30
	NeedleParserRULE_arrayStmt        = 31
	NeedleParserRULE_arrayList        = 32
	NeedleParserRULE_arrayValue       = 33
	NeedleParserRULE_mapStmt          = 34
	NeedleParserRULE_pairList         = 35
	NeedleParserRULE_pair             = 36
	NeedleParserRULE_pairValue        = 37
	NeedleParserRULE_arguments        = 38
	NeedleParserRULE_argumentsList    = 39
	NeedleParserRULE_exprList         = 40
	NeedleParserRULE_expr             = 41
	NeedleParserRULE_primaryExpr      = 42
	NeedleParserRULE_indexExpr        = 43
	NeedleParserRULE_sliceExpr        = 44
	NeedleParserRULE_operand          = 45
	NeedleParserRULE_literal          = 46
	NeedleParserRULE_typeName         = 47
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
			p.Literal()
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
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(186)
			p.FuncDescriptor()
		}

	case 2:
		{
			p.SetState(187)
			p.DefaultFuncDef()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(190)
		p.FuncSignature()
	}
	{
		p.SetState(191)
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
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NeedleParserFUNC {
		{
			p.SetState(193)
			p.Match(NeedleParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(196)
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
		p.SetState(198)
		p.Match(NeedleParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
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
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NeedleParserLPAREN {
		{
			p.SetState(201)
			p.ParameterList()
		}

	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == NeedleParserDOT {
		{
			p.SetState(204)
			p.FuncTail()
		}

		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&1023) != 0 {
		{
			p.SetState(210)
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
		p.SetState(213)
		p.Match(NeedleParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(214)
		p.Match(NeedleParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NeedleParserLPAREN {
		{
			p.SetState(215)
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
		p.SetState(218)
		p.Match(NeedleParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NeedleParserIdentifier {
		{
			p.SetState(219)
			p.Parameter()
		}
		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == NeedleParserCOMMA {
			{
				p.SetState(220)
				p.Match(NeedleParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(225)
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
		p.SetState(227)
		p.IdentifierList()
	}
	{
		p.SetState(228)
		p.TypeName()
	}
	p.SetState(237)
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
			p.SetState(230)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == NeedleParserCOMMA {
				{
					p.SetState(229)
					p.Match(NeedleParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(232)
				p.IdentifierList()
			}
			{
				p.SetState(233)
				p.TypeName()
			}

		}
		p.SetState(239)
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
		p.SetState(240)
		p.TypeName()
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == NeedleParserCOMMA || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&1023) != 0) {
		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == NeedleParserCOMMA {
			{
				p.SetState(241)
				p.Match(NeedleParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(244)
			p.TypeName()
		}

		p.SetState(249)
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
		p.SetState(250)
		p.Match(NeedleParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-8097929526849250814) != 0) || ((int64((_la-74)) & ^0x3f) == 0 && ((int64(1)<<(_la-74))&16778239) != 0) {
		{
			p.SetState(251)
			p.StatementList()
		}

	}
	{
		p.SetState(254)
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
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-8097929526849250814) != 0) || ((int64((_la-74)) & ^0x3f) == 0 && ((int64(1)<<(_la-74))&16778239) != 0) {
		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
		case 1:
			p.SetState(257)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == NeedleParserSEMI {
				{
					p.SetState(256)
					p.Match(NeedleParserSEMI)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}

		case 2:
			p.SetState(260)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == NeedleParserEOS {
				{
					p.SetState(259)
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
			p.SetState(264)
			p.Statement()
		}
		{
			p.SetState(265)
			p.Eos()
		}

		p.SetState(269)
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
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NeedleParserLBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(271)
			p.Block()
		}

	case NeedleParserLPAREN, NeedleParserNOT, NeedleParserADD, NeedleParserSUB, NeedleParserTRUE, NeedleParserFALSE, NeedleParserNIL, NeedleParserIdentifier, NeedleParserDollarIdentifier, NeedleParserAtIdentifier, NeedleParserInterpretedStringLiteral, NeedleParserRawStringLiteral, NeedleParserDecimalLiteral, NeedleParserFloatLiteral, NeedleParserHexLiteral, NeedleParserOctalLiteral, NeedleParserBinaryLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(272)
			p.SimpleStmt()
		}

	case NeedleParserVAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(273)
			p.VarDef()
		}

	case NeedleParserIF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(274)
			p.IfStmt()
		}

	case NeedleParserWHILE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(275)
			p.WhileStmt()
		}

	case NeedleParserCONTINUE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(276)
			p.ContinueStmt()
		}

	case NeedleParserBREAK:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(277)
			p.BreakStmt()
		}

	case NeedleParserRETURN:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(278)
			p.ReturnStmt()
		}

	case NeedleParserERRWARNING, NeedleParserERRINFO, NeedleParserERROR:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(279)
			p.ErrorStmt()
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

// ISimpleStmtContext is an interface to support dynamic dispatch.
type ISimpleStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Assignment() IAssignmentContext
	IncDecStmt() IIncDecStmtContext
	AssignMapArrStmt() IAssignMapArrStmtContext

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

func (s *SimpleStmtContext) AssignMapArrStmt() IAssignMapArrStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignMapArrStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignMapArrStmtContext)
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
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(282)
			p.expr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(283)
			p.Assignment()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(284)
			p.IncDecStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(285)
			p.AssignMapArrStmt()
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
		p.SetState(288)
		p.expr(0)
	}
	{
		p.SetState(289)
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

// IAssignMapArrStmtContext is an interface to support dynamic dispatch.
type IAssignMapArrStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentifierVar() IIdentifierVarContext
	EQ() antlr.TerminalNode
	InitMapArrStmt() IInitMapArrStmtContext

	// IsAssignMapArrStmtContext differentiates from other interfaces.
	IsAssignMapArrStmtContext()
}

type AssignMapArrStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignMapArrStmtContext() *AssignMapArrStmtContext {
	var p = new(AssignMapArrStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_assignMapArrStmt
	return p
}

func InitEmptyAssignMapArrStmtContext(p *AssignMapArrStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_assignMapArrStmt
}

func (*AssignMapArrStmtContext) IsAssignMapArrStmtContext() {}

func NewAssignMapArrStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignMapArrStmtContext {
	var p = new(AssignMapArrStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_assignMapArrStmt

	return p
}

func (s *AssignMapArrStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignMapArrStmtContext) IdentifierVar() IIdentifierVarContext {
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

func (s *AssignMapArrStmtContext) EQ() antlr.TerminalNode {
	return s.GetToken(NeedleParserEQ, 0)
}

func (s *AssignMapArrStmtContext) InitMapArrStmt() IInitMapArrStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInitMapArrStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInitMapArrStmtContext)
}

func (s *AssignMapArrStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignMapArrStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignMapArrStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterAssignMapArrStmt(s)
	}
}

func (s *AssignMapArrStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitAssignMapArrStmt(s)
	}
}

func (p *NeedleParser) AssignMapArrStmt() (localctx IAssignMapArrStmtContext) {
	localctx = NewAssignMapArrStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, NeedleParserRULE_assignMapArrStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.IdentifierVar()
	}
	{
		p.SetState(292)
		p.Match(NeedleParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(293)
		p.InitMapArrStmt()
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

// IInitMapArrStmtContext is an interface to support dynamic dispatch.
type IInitMapArrStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MapStmt() IMapStmtContext
	ArrayStmt() IArrayStmtContext

	// IsInitMapArrStmtContext differentiates from other interfaces.
	IsInitMapArrStmtContext()
}

type InitMapArrStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitMapArrStmtContext() *InitMapArrStmtContext {
	var p = new(InitMapArrStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_initMapArrStmt
	return p
}

func InitEmptyInitMapArrStmtContext(p *InitMapArrStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_initMapArrStmt
}

func (*InitMapArrStmtContext) IsInitMapArrStmtContext() {}

func NewInitMapArrStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitMapArrStmtContext {
	var p = new(InitMapArrStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_initMapArrStmt

	return p
}

func (s *InitMapArrStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *InitMapArrStmtContext) MapStmt() IMapStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapStmtContext)
}

func (s *InitMapArrStmtContext) ArrayStmt() IArrayStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayStmtContext)
}

func (s *InitMapArrStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitMapArrStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitMapArrStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterInitMapArrStmt(s)
	}
}

func (s *InitMapArrStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitInitMapArrStmt(s)
	}
}

func (p *NeedleParser) InitMapArrStmt() (localctx IInitMapArrStmtContext) {
	localctx = NewInitMapArrStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, NeedleParserRULE_initMapArrStmt)
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NeedleParserLBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(295)
			p.MapStmt()
		}

	case NeedleParserLBRACK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(296)
			p.ArrayStmt()
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
	p.EnterRule(localctx, 42, NeedleParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		p.ExprList()
	}
	{
		p.SetState(300)
		p.Assign_op()
	}
	{
		p.SetState(301)
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
	p.EnterRule(localctx, 44, NeedleParserRULE_varDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(303)
		p.Match(NeedleParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(304)
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
	p.EnterRule(localctx, 46, NeedleParserRULE_ifStmt)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(306)
		p.Match(NeedleParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(307)
		p.IfBody()
	}
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(308)
				p.Match(NeedleParserELIF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(309)
				p.IfBody()
			}

		}
		p.SetState(314)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NeedleParserELSE {
		{
			p.SetState(315)
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
	p.EnterRule(localctx, 48, NeedleParserRULE_ifBody)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(318)
			p.Match(NeedleParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(319)
			p.expr(0)
		}
		{
			p.SetState(320)
			p.Match(NeedleParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(322)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(325)
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
	p.EnterRule(localctx, 50, NeedleParserRULE_elseBody)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(327)
		p.Match(NeedleParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(328)
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
	p.EnterRule(localctx, 52, NeedleParserRULE_returnStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(330)
		p.Match(NeedleParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(332)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(331)
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
	p.EnterRule(localctx, 54, NeedleParserRULE_continueStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(334)
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
	p.EnterRule(localctx, 56, NeedleParserRULE_breakStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(336)
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
	p.EnterRule(localctx, 58, NeedleParserRULE_whileStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(338)
		p.Match(NeedleParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(339)
		p.expr(0)
	}
	{
		p.SetState(340)
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
	p.EnterRule(localctx, 60, NeedleParserRULE_errorStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(342)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-8791026472627208192) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(343)
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

// IArrayStmtContext is an interface to support dynamic dispatch.
type IArrayStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	ArrayList() IArrayListContext

	// IsArrayStmtContext differentiates from other interfaces.
	IsArrayStmtContext()
}

type ArrayStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayStmtContext() *ArrayStmtContext {
	var p = new(ArrayStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_arrayStmt
	return p
}

func InitEmptyArrayStmtContext(p *ArrayStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_arrayStmt
}

func (*ArrayStmtContext) IsArrayStmtContext() {}

func NewArrayStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayStmtContext {
	var p = new(ArrayStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_arrayStmt

	return p
}

func (s *ArrayStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayStmtContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(NeedleParserLBRACK, 0)
}

func (s *ArrayStmtContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(NeedleParserRBRACK, 0)
}

func (s *ArrayStmtContext) ArrayList() IArrayListContext {
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

func (s *ArrayStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterArrayStmt(s)
	}
}

func (s *ArrayStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitArrayStmt(s)
	}
}

func (p *NeedleParser) ArrayStmt() (localctx IArrayStmtContext) {
	localctx = NewArrayStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, NeedleParserRULE_arrayStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(345)
		p.Match(NeedleParserLBRACK)
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

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&579838452024005250) != 0) || ((int64((_la-74)) & ^0x3f) == 0 && ((int64(1)<<(_la-74))&1023) != 0) {
		{
			p.SetState(346)
			p.ArrayList()
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
	p.EnterRule(localctx, 64, NeedleParserRULE_arrayList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(351)
		p.ArrayValue()
	}
	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == NeedleParserCOMMA {
		{
			p.SetState(352)
			p.Match(NeedleParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(353)
			p.ArrayValue()
		}

		p.SetState(358)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(359)
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
	Expr() IExprContext
	InitMapArrStmt() IInitMapArrStmtContext

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

func (s *ArrayValueContext) InitMapArrStmt() IInitMapArrStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInitMapArrStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInitMapArrStmtContext)
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
	p.EnterRule(localctx, 66, NeedleParserRULE_arrayValue)
	p.SetState(363)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NeedleParserLPAREN, NeedleParserNOT, NeedleParserADD, NeedleParserSUB, NeedleParserTRUE, NeedleParserFALSE, NeedleParserNIL, NeedleParserIdentifier, NeedleParserDollarIdentifier, NeedleParserAtIdentifier, NeedleParserInterpretedStringLiteral, NeedleParserRawStringLiteral, NeedleParserDecimalLiteral, NeedleParserFloatLiteral, NeedleParserHexLiteral, NeedleParserOctalLiteral, NeedleParserBinaryLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(361)
			p.expr(0)
		}

	case NeedleParserLBRACK, NeedleParserLBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(362)
			p.InitMapArrStmt()
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

// IMapStmtContext is an interface to support dynamic dispatch.
type IMapStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	PairList() IPairListContext

	// IsMapStmtContext differentiates from other interfaces.
	IsMapStmtContext()
}

type MapStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapStmtContext() *MapStmtContext {
	var p = new(MapStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_mapStmt
	return p
}

func InitEmptyMapStmtContext(p *MapStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NeedleParserRULE_mapStmt
}

func (*MapStmtContext) IsMapStmtContext() {}

func NewMapStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapStmtContext {
	var p = new(MapStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeedleParserRULE_mapStmt

	return p
}

func (s *MapStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *MapStmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(NeedleParserLBRACE, 0)
}

func (s *MapStmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(NeedleParserRBRACE, 0)
}

func (s *MapStmtContext) PairList() IPairListContext {
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

func (s *MapStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.EnterMapStmt(s)
	}
}

func (s *MapStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeedleParserListener); ok {
		listenerT.ExitMapStmt(s)
	}
}

func (p *NeedleParser) MapStmt() (localctx IMapStmtContext) {
	localctx = NewMapStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, NeedleParserRULE_mapStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(365)
		p.Match(NeedleParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-74)) & ^0x3f) == 0 && ((int64(1)<<(_la-74))&27) != 0 {
		{
			p.SetState(366)
			p.PairList()
		}

	}
	{
		p.SetState(369)
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
	p.EnterRule(localctx, 70, NeedleParserRULE_pairList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(371)
		p.Pair()
	}
	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(372)
				p.Match(NeedleParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(373)
				p.Pair()
			}

		}
		p.SetState(378)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(380)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NeedleParserCOMMA {
		{
			p.SetState(379)
			p.Match(NeedleParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(382)
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
	p.EnterRule(localctx, 72, NeedleParserRULE_pair)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NeedleParserInterpretedStringLiteral, NeedleParserRawStringLiteral:
		{
			p.SetState(384)
			p.StringLiteral()
		}

	case NeedleParserIdentifier, NeedleParserDollarIdentifier:
		{
			p.SetState(385)
			p.IdentifierVar()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(388)
		p.Match(NeedleParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(389)
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
	InitMapArrStmt() IInitMapArrStmtContext

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

func (s *PairValueContext) InitMapArrStmt() IInitMapArrStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInitMapArrStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInitMapArrStmtContext)
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
	p.EnterRule(localctx, 74, NeedleParserRULE_pairValue)
	p.SetState(398)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NeedleParserIdentifier, NeedleParserDollarIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(391)
			p.IdentifierVar()
		}
		p.SetState(394)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(392)
				p.IndexExpr()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) == 2 {
			{
				p.SetState(393)
				p.SliceExpr()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case NeedleParserTRUE, NeedleParserFALSE, NeedleParserNIL, NeedleParserInterpretedStringLiteral, NeedleParserRawStringLiteral, NeedleParserDecimalLiteral, NeedleParserFloatLiteral, NeedleParserHexLiteral, NeedleParserOctalLiteral, NeedleParserBinaryLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(396)
			p.Literal()
		}

	case NeedleParserLBRACK, NeedleParserLBRACE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(397)
			p.InitMapArrStmt()
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
	p.EnterRule(localctx, 76, NeedleParserRULE_arguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(400)
		p.Match(NeedleParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(402)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&579838452024005250) != 0) || ((int64((_la-74)) & ^0x3f) == 0 && ((int64(1)<<(_la-74))&1023) != 0) {
		{
			p.SetState(401)
			p.ArgumentsList()
		}

	}
	{
		p.SetState(404)
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
	AllInitMapArrStmt() []IInitMapArrStmtContext
	InitMapArrStmt(i int) IInitMapArrStmtContext
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

func (s *ArgumentsListContext) AllInitMapArrStmt() []IInitMapArrStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInitMapArrStmtContext); ok {
			len++
		}
	}

	tst := make([]IInitMapArrStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInitMapArrStmtContext); ok {
			tst[i] = t.(IInitMapArrStmtContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentsListContext) InitMapArrStmt(i int) IInitMapArrStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInitMapArrStmtContext); ok {
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

	return t.(IInitMapArrStmtContext)
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
	p.EnterRule(localctx, 78, NeedleParserRULE_argumentsList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(408)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NeedleParserLBRACK, NeedleParserLBRACE:
		{
			p.SetState(406)
			p.InitMapArrStmt()
		}

	case NeedleParserLPAREN, NeedleParserNOT, NeedleParserADD, NeedleParserSUB, NeedleParserTRUE, NeedleParserFALSE, NeedleParserNIL, NeedleParserIdentifier, NeedleParserDollarIdentifier, NeedleParserAtIdentifier, NeedleParserInterpretedStringLiteral, NeedleParserRawStringLiteral, NeedleParserDecimalLiteral, NeedleParserFloatLiteral, NeedleParserHexLiteral, NeedleParserOctalLiteral, NeedleParserBinaryLiteral:
		{
			p.SetState(407)
			p.expr(0)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(417)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == NeedleParserCOMMA {
		{
			p.SetState(410)
			p.Match(NeedleParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(413)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case NeedleParserLBRACK, NeedleParserLBRACE:
			{
				p.SetState(411)
				p.InitMapArrStmt()
			}

		case NeedleParserLPAREN, NeedleParserNOT, NeedleParserADD, NeedleParserSUB, NeedleParserTRUE, NeedleParserFALSE, NeedleParserNIL, NeedleParserIdentifier, NeedleParserDollarIdentifier, NeedleParserAtIdentifier, NeedleParserInterpretedStringLiteral, NeedleParserRawStringLiteral, NeedleParserDecimalLiteral, NeedleParserFloatLiteral, NeedleParserHexLiteral, NeedleParserOctalLiteral, NeedleParserBinaryLiteral:
			{
				p.SetState(412)
				p.expr(0)
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(419)
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
	p.EnterRule(localctx, 80, NeedleParserRULE_exprList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(420)
		p.expr(0)
	}
	p.SetState(425)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == NeedleParserCOMMA {
		{
			p.SetState(421)
			p.Match(NeedleParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(422)
			p.expr(0)
		}

		p.SetState(427)
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
	_startState := 82
	p.EnterRecursionRule(localctx, 82, NeedleParserRULE_expr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(435)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NeedleParserLPAREN, NeedleParserTRUE, NeedleParserFALSE, NeedleParserNIL, NeedleParserIdentifier, NeedleParserDollarIdentifier, NeedleParserAtIdentifier, NeedleParserInterpretedStringLiteral, NeedleParserRawStringLiteral, NeedleParserDecimalLiteral, NeedleParserFloatLiteral, NeedleParserHexLiteral, NeedleParserOctalLiteral, NeedleParserBinaryLiteral:
		{
			p.SetState(429)
			p.primaryExpr(0)
		}
		{
			p.SetState(430)
			p.Eos()
		}

	case NeedleParserNOT, NeedleParserADD, NeedleParserSUB:
		{
			p.SetState(432)
			p.Unary_op()
		}
		{
			p.SetState(433)
			p.expr(5)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(459)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(457)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NeedleParserRULE_expr)
				p.SetState(437)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(438)
					p.Mul_op()
				}
				{
					p.SetState(439)
					p.expr(5)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NeedleParserRULE_expr)
				p.SetState(441)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(442)
					p.Rel_op()
				}
				{
					p.SetState(443)
					p.expr(4)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NeedleParserRULE_expr)
				p.SetState(445)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(446)
					p.Logical_op()
				}
				{
					p.SetState(447)
					p.expr(3)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NeedleParserRULE_expr)
				p.SetState(449)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(450)
					p.Add_op()
				}
				{
					p.SetState(451)
					p.expr(2)
				}

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NeedleParserRULE_expr)
				p.SetState(453)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(454)
					p.IndexExpr()
				}

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NeedleParserRULE_expr)
				p.SetState(455)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(456)
					p.SliceExpr()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(461)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext())
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
	_startState := 84
	p.EnterRecursionRule(localctx, 84, NeedleParserRULE_primaryExpr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(463)
		p.Operand()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(473)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext())
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
			p.SetState(465)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			p.SetState(468)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == NeedleParserDOT {
				{
					p.SetState(466)
					p.Match(NeedleParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(467)
					p.Match(NeedleParserIdentifier)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(470)
				p.Arguments()
			}

		}
		p.SetState(475)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 86, NeedleParserRULE_indexExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(476)
		p.Match(NeedleParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(477)
		p.expr(0)
	}
	{
		p.SetState(478)
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
	p.EnterRule(localctx, 88, NeedleParserRULE_sliceExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(480)
		p.Match(NeedleParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(482)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&579838452024004610) != 0) || ((int64((_la-74)) & ^0x3f) == 0 && ((int64(1)<<(_la-74))&1023) != 0) {
		{
			p.SetState(481)
			p.expr(0)
		}

	}
	{
		p.SetState(484)
		p.Match(NeedleParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(486)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&579838452024004610) != 0) || ((int64((_la-74)) & ^0x3f) == 0 && ((int64(1)<<(_la-74))&1023) != 0) {
		{
			p.SetState(485)
			p.expr(0)
		}

	}
	{
		p.SetState(488)
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
	p.EnterRule(localctx, 90, NeedleParserRULE_operand)
	p.SetState(496)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NeedleParserIdentifier, NeedleParserDollarIdentifier, NeedleParserAtIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(490)
			p.IdentifierFull()
		}

	case NeedleParserTRUE, NeedleParserFALSE, NeedleParserNIL, NeedleParserInterpretedStringLiteral, NeedleParserRawStringLiteral, NeedleParserDecimalLiteral, NeedleParserFloatLiteral, NeedleParserHexLiteral, NeedleParserOctalLiteral, NeedleParserBinaryLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(491)
			p.Literal()
		}

	case NeedleParserLPAREN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(492)
			p.Match(NeedleParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(493)
			p.expr(0)
		}
		{
			p.SetState(494)
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
	p.EnterRule(localctx, 92, NeedleParserRULE_literal)
	p.SetState(502)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NeedleParserDecimalLiteral, NeedleParserFloatLiteral, NeedleParserHexLiteral, NeedleParserOctalLiteral, NeedleParserBinaryLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(498)
			p.NumberLiteral()
		}

	case NeedleParserInterpretedStringLiteral, NeedleParserRawStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(499)
			p.StringLiteral()
		}

	case NeedleParserTRUE, NeedleParserFALSE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(500)
			p.BooleanLiteral()
		}

	case NeedleParserNIL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(501)
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
	p.EnterRule(localctx, 94, NeedleParserRULE_typeName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(504)
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
	p.EnterRule(localctx, 96, NeedleParserRULE_incDec_op)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(506)
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
		p.SetState(508)
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
		p.SetState(510)
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
		p.SetState(512)
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
		p.SetState(514)
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
		p.SetState(516)
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
		p.SetState(518)
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
		p.SetState(520)
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
		p.SetState(522)
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
		p.SetState(524)
		p.Match(NeedleParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(531)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == NeedleParserCOMMA || _la == NeedleParserIdentifier {
		p.SetState(526)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == NeedleParserCOMMA {
			{
				p.SetState(525)
				p.Match(NeedleParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(528)
			p.Match(NeedleParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(533)
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
		p.SetState(534)
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
		p.SetState(536)
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
		p.SetState(538)
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
	p.SetState(545)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(540)
			p.Match(NeedleParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(541)
			p.Match(NeedleParserEOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(543)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(542)
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
	case 41:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	case 42:
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
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

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
