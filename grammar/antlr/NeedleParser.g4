parser grammar NeedleParser;

options {
	tokenVocab = NeedleLexer;
}

sourceMain: ((contractDef | funcDef) eos)* EOF;

contractDef:
	CONTRACT Identifier LBRACE (contractPart eos)* RBRACE;

contractPart: dataDef | settingsDef | funcDef;

dataDef: DATA LBRACE (dataPartList eos)* RBRACE;

dataPartList: Identifier typeName (dataTag = stringLiteral)?;

settingsDef:
	SETTINGS LBRACE (Identifier EQ literal eos)* RBRACE;

funcDef: (funcDescriptor | defaultFuncDef) funcSignature block;

defaultFuncDef: FUNC? (CONDITIONS | ACTION);

funcDescriptor: FUNC Identifier;

funcSignature: parameterList? funcTail* returnParameters?;

funcTail: '.' Identifier parameterList?;

parameterList: LPAREN (parameter COMMA?)? RPAREN;

parameter:
	identifierList typeName (COMMA? identifierList typeName)*;

returnParameters: typeName (COMMA? typeName)*;

block: LBRACE statementList? RBRACE;

statementList: ((SEMI? | EOS?) statement eos)+;

statement:
	block
	| simpleStmt
	| varDef
	| ifStmt
	| whileStmt
	| continueStmt
	| breakStmt
	| returnStmt
	| errorStmt;

simpleStmt: expr | assignment | incDecStmt | assignMapArrStmt;

incDecStmt: expr incDec_op;

assignMapArrStmt: identifierVar EQ initMapArrStmt;

initMapArrStmt: mapStmt | arrayStmt;

assignment: exprList assign_op exprList;

varDef: VAR parameter;

ifStmt: IF ifBody (ELIF ifBody)*? elseBody?;

ifBody: ((LPAREN expr RPAREN) | expr) block;

elseBody: ELSE block;

returnStmt: RETURN expr?;

continueStmt: CONTINUE;

breakStmt: BREAK;

whileStmt: WHILE expr block;

errorStmt: (ERRWARNING | ERRINFO | ERROR) expr;

arrayStmt: LBRACK arrayList? RBRACK;

arrayList: arrayValue (COMMA arrayValue)* eos;

arrayValue: expr | initMapArrStmt;

mapStmt: LBRACE pairList? RBRACE;

pairList: pair (COMMA pair)* COMMA? eos;

pair: (stringLiteral | identifierVar) COLON pairValue;

pairValue:
	identifierVar (indexExpr | sliceExpr)?
	| literal
	| initMapArrStmt;

arguments: LPAREN argumentsList? RPAREN;

argumentsList: (initMapArrStmt | expr) (
		COMMA (initMapArrStmt | expr)
	)*;

exprList: expr (COMMA expr)*;

expr:
	primaryExpr eos
	| expr indexExpr
	| expr sliceExpr
	| unary_op expr
	| expr mul_op expr
	| expr rel_op expr
	| expr logical_op expr
	| expr add_op expr;

primaryExpr:
	operand
	| primaryExpr (DOT Identifier)? arguments;

indexExpr:  LBRACK expr RBRACK;

sliceExpr: LBRACK expr? COLON expr? RBRACK;

operand: identifierFull | literal | LPAREN expr RPAREN;

literal: numberLiteral | stringLiteral | booleanLiteral | NIL;

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
	; //	| RuneLiteral

booleanLiteral: TRUE | FALSE;

eos: SEMI | EOS | EOF?;