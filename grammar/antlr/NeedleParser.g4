parser grammar NeedleParser;

options {
	tokenVocab = NeedleLexer;
}

sourceMain: ((contractDef | funcDef) eos )* EOF;

contractDef: CONTRACT Identifier LBRACE (contractPart eos)* RBRACE;

contractPart: dataDef | settingsDef | funcDef;

dataDef: DATA LBRACE (dataPartList eos)* RBRACE;

dataPartList: Identifier typeName (dataTag=stringLiteral)?;

settingsDef:
	SETTINGS LBRACE (Identifier EQ settingsValue eos)* RBRACE;

settingsValue: numberLiteral | booleanLiteral | stringLiteral;

funcDef: (innerFuncDef | defaultFuncDef) (
		DOT Identifier funcSignature
	)? block;

innerFuncDef: funcDescriptor funcSignature;

defaultFuncDef: FUNC? (CONDITIONS | ACTION) funcSignature;

funcDescriptor: FUNC Identifier;

funcSignature: parameterList? returnParameters?;

parameterList:
	LPAREN (parameter (COMMA? parameter)* COMMA?)? RPAREN;

parameter:
	identifierList typeName (COMMA? identifierList typeName)*;

returnParameters: typeName (COMMA? typeName)*;

block:  LBRACE statList? RBRACE;

statList: ((SEMI? | EOS?) stat eos)+;

stat:
	block
	| simpleStat
	| varDef
	| ifStat
	| whileStat
	| continueStat
	| breakStat
	| returnStat
	| errorStat;

varDef: VAR parameter;

ifStat:
	IF ((LPAREN expr RPAREN) | expr) block (
		ELIF ((LPAREN expr RPAREN) | expr) block
	)*? (ELSE block)?;

returnStat: RETURN expr?;

continueStat: CONTINUE;

breakStat: BREAK;

whileStat: WHILE expr block;

errorStat: (ERRWARNING | ERRINFO | ERROR) expr;

typeName:
	BOOL
	| BYTES
	| INT
	| ADDRESS
	| ARRAY
	| MAP
	| MONEY
	| FLOAT
	| STRING
	| FILE;

sliceStat: LBRACK indexNumber? COLON indexNumber? RBRACK;

indexNumber: numberLiteral | identifierVar;

arrayStat: LBRACK arrayList? RBRACK;

arrayList: arrayValue (COMMA arrayValue)* eos;

arrayValue: arrayStat | expr | objectStat;

indexStat: LBRACK expr RBRACK;

objectStat: LBRACE pairList? RBRACE;

pairList: pair (COMMA pair)* COMMA? eos;

pair: (stringLiteral | identifierVar) COLON pairValue+;

pairValue:
	identifierVar
	| stringLiteral
	| numberLiteral
	| arrayStat
	| objectStat
	| sliceStat
	| indexStat;

arguments: LPAREN argumentsList? RPAREN;

argumentsList: (initMapArrStat | expr) (
		COMMA (initMapArrStat | expr)
	)*;

simpleStat:
	assignment
	| exprStat
	| incDecStat
	| assignMapArrStat;

incDecStat: expr incDec_op;

exprStat: expr;

assignMapArrStat: exprList EQ initMapArrStat;

initMapArrStat: objectStat | arrayStat;

assignment: exprList assign_op exprList;

primaryExpr:
	operand
	| primaryExpr (
		(DOT Identifier)? arguments
		| sliceStat
		| indexStat
	);

operand:
	identifierFull
	| numberLiteral
	| stringLiteral
	| booleanLiteral
	| LPAREN expr RPAREN
	| NIL;

exprList: expr (COMMA expr)*;

expr:
	 primaryExpr eos
	| unary_op expr
	| expr mul_op expr
	| expr rel_op expr
	| expr logical_op expr
	| expr add_op expr ;

incDec_op: INC | DEC;

mul_op: MUL | QUO | MOD | LSHIFT | RSHIFT | BIT_AND;

unary_op: ADD | SUB | NOT;

add_op: ADD | SUB | BIT_OR | BIT_XOR;

logical_op: AND | OR;

rel_op: LESS | GREATER | LESS_EQ | GR_EQ | EQ_EQ | NOT_EQ;

assign_op:
	EQ
	| ADD_EQ
	| SUB_EQ
	| MUL_EQ
	| DIV_EQ
	| MOD_EQ
	| LSHIFT_EQ
	| RSHIFT_EQ
	| BIT_AND_EQ
	| BIT_OR_EQ
	| BIT_XOR_EQ;

identifierFull: Identifier | DollarIdentifier | AtIdentifier;

identifierVar: Identifier | DollarIdentifier;

identifierList: Identifier (COMMA? Identifier)*;

stringLiteral: InterpretedStringLiteral | RawStringLiteral;

numberLiteral:
	DecimalLiteral
	| BinaryLiteral
	| OctalLiteral
	| HexLiteral
	| FloatLiteral
//	| RuneLiteral
	;

booleanLiteral: TRUE | FALSE;

eos: SEMI | EOS | EOF?;
