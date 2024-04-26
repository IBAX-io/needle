parser grammar NeedleParser;

options {
	tokenVocab = NeedleLexer;
}

sourceMain: ((contractDef | funcDef) eos)* EOF;

contractDef:
	CONTRACT Identifier LBRACE (contractPart eos)* RBRACE;

contractPart: dataDef | settingsDef | funcDef;

dataDef: DATA LBRACE (dataPart eos)* RBRACE;

dataPart: Identifier typeName (dataTag = stringLiteral)?;

settingsDef:
	SETTINGS LBRACE (Identifier EQ literal eos)* RBRACE;

funcDef: (funcDescriptor | defaultFuncDef) funcSignature block;

defaultFuncDef: FUNC? (CONDITIONS | ACTION);

funcDescriptor: FUNC Identifier;

funcSignature: parameterList? funcTail* returnParameters?;

funcTail: '.' Identifier parameterList?;

parameterList: LPAREN (parameter COMMA?)? (Identifier TAIL)? RPAREN;

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

simpleStmt: expr | assignment | incDecStmt;

incDecStmt: expr incDec_op;

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

arrayExpr: LBRACK arrayList? COMMA? RBRACK;

arrayList: exprList eos;

mapExpr: LBRACE pairList? RBRACE;

pairList: pair (COMMA pair)* COMMA? eos;

pair: (stringLiteral | identifierVar) COLON expr;

arguments: LPAREN (exprList TAIL?)? RPAREN;

exprList: expr (COMMA expr)*;

expr:
	primaryExpr eos
	| expr indexExpr
	| expr sliceExpr
	| mapExpr
	| arrayExpr
	| contractCall
	| unary_op expr
	| expr mul_op expr
	| expr rel_op expr
	| expr logical_op expr
	| expr add_op expr;

primaryExpr: operand | primaryExpr (DOT Identifier)? arguments;

indexExpr: LBRACK index=expr RBRACK;

sliceExpr: LBRACK low=expr? COLON high=expr? RBRACK;

contractCall: AtIdentifier arguments;

operand: identifierVar | literal | LPAREN expr RPAREN;

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

identifierVar: Identifier | DollarIdentifier;

identifierList: Identifier (COMMA? Identifier)*;

stringLiteral: InterpretedStringLiteral | RawStringLiteral;

// @todo add RuneLiteral.
numberLiteral:
	DecimalLiteral
	| BinaryLiteral
	| OctalLiteral
	| HexLiteral
	| FloatLiteral;

booleanLiteral: TRUE | FALSE;

eos: SEMI | EOS | EOF?;