parser grammar NeedleParser;

options {
	tokenVocab = NeedleLexer;
	superClass = NeedleParserBase;
}

sourceMain: (contractDef | funcDef)* EOF;

contractDef: CONTRACT Identifier LBRACE contractPart* RBRACE;

contractPart: (dataDef | settingsDef | funcDef) eos;

dataDef: DATA eos LBRACE (dataPartList eos)* RBRACE;

dataPartList: Identifier typeName dataTag?;

dataTag: StringLiteral;

settingsDef:
	SETTINGS LBRACE (Identifier EQ settingsValue eos)* RBRACE;

settingsValue: (numberLiteral | booleanLiteral | StringLiteral);

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

returnParameters: typeName (COMMA? typeName)* COMMA? ;

block: LBRACE statementList? RBRACE;

statementList: (
        (SEMI? | EOS? | {p.closingBracket()}?) statement eos
	)+;

statement:
	block
	| simpleStatement
	| varDef
	| ifStatement
	| whileStatement
	| continueStatement
	| breakStatement
	| returnStatement
	| errorStatement;

varDef: VAR parameter;

ifStatement:
	IF ((LPAREN expression RPAREN) | expression) block (
		ELIF ((LPAREN expression RPAREN) | expression) block
	)*? (ELSE block)?;

returnStatement: RETURN expression?;

continueStatement: CONTINUE;

breakStatement: BREAK;

whileStatement: WHILE expression statement;

errorStatement: (ERRWARNING | ERRINFO | ERROR) expression;

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

sliceStatement: LBRACK expression? COLON expression? RBRACK;

arrayStatement: LBRACK arrayList? RBRACK;

arrayList: arrayValue (COMMA arrayValue)*;

arrayValue: arrayStatement | expression | objectStatement;

indexStatement: LBRACK expression RBRACK;

objectStatement: LBRACE pairList? RBRACE;

pairList: pair (COMMA pair)* COMMA?;

pair: (StringLiteral | identifierVar) COLON pairValue+;

pairValue:
	identifierVar
	| StringLiteral
	| numberLiteral
	| arrayStatement
	| objectStatement
	| sliceStatement
	| indexStatement;

arguments: LPAREN argumentsList? RPAREN;

argumentsList: (initMapArrStatement | expression) (
		COMMA (initMapArrStatement | expression)
	)*;

simpleStatement:
	assignment
	| expressionStatement
	| incDecStatement
	| assignMapArrStatement;

incDecStatement: expression incDec_op;

expressionStatement: expression;

assignMapArrStatement: expressionList EQ initMapArrStatement;

initMapArrStatement: objectStatement | arrayStatement;

assignment: expressionList assign_op expressionList;

primaryExpression:
	operand
	| primaryExpression (
		(DOT Identifier)? arguments
		| sliceStatement
		| indexStatement
	);

operand:
	identifierFull
	| numberLiteral
	| StringLiteral
	| booleanLiteral
	| LPAREN expression RPAREN
	| NIL;

expressionList: expression (COMMA expression)*;

expression:
	primaryExpression
	| unary_op expression
	| expression mul_op expression
	| expression rel_op expression
	| expression logical_op expression
	| expression add_op expression;

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

numberLiteral:
	DecimalLiteral
	| BinaryLiteral
	| OctalLiteral
	| HexLiteral
	| FloatLiteral
//	| RuneLiteral
	;

booleanLiteral: TRUE | FALSE;

eos: SEMI | EOF | EOS | {p.closingBracket()}?;
