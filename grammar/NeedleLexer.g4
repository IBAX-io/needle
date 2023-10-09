lexer grammar NeedleLexer;

// delimiters
LPAREN: '(';
RPAREN: ')'; //-> mode(NLSEMI);
COMMA: ',';
DOT: '.';
COLON: ':';
EQ: '=';
LBRACK: '[';
RBRACK: ']'; //-> mode(NLSEMI);
LBRACE: '{';
RBRACE: '}'; //-> mode(NLSEMI);
SEMI: ';';

// Operators
NOT: '!';
MUL: '*';
ADD: '+';
SUB: '-';
QUO: '/';
LESS: '<';
GREATER: '>';
NOT_EQ: '!=';
AND: '&&';
LESS_EQ: '<=';
EQ_EQ: '==';
GR_EQ: '>=';
OR: '||';
BIT_AND: '&';
BIT_OR: '|';
BIT_XOR: '^';
MOD: '%';
LSHIFT: '<<';
RSHIFT: '>>';
ADD_EQ: '+=';
SUB_EQ: '-=';
MUL_EQ: '*=';
DIV_EQ: '/=';
MOD_EQ: '%=';
LSHIFT_EQ: '<<=';
RSHIFT_EQ: '>>=';
BIT_AND_EQ: '&=';
BIT_OR_EQ: '|=';
BIT_XOR_EQ: '^=';
INC: '++'; //-> mode(NLSEMI);
DEC: '--'; //-> mode(NLSEMI);

// Keywords
CONTRACT: 'contract';
FUNC: 'func';
RETURN: 'return'; //-> mode(NLSEMI);
IF: 'if';
ELIF: 'elif';
ELSE: 'else';
WHILE: 'while';
TRUE: 'true';
FALSE: 'false';
VAR: 'var';
DATA: 'data';
SETTINGS: 'settings';
BREAK: 'break'; //-> mode(NLSEMI);
CONTINUE: 'continue'; //-> mode(NLSEMI);
ERRWARNING: 'warning';
ERRINFO: 'info';
NIL: 'nil'; //-> mode(NLSEMI);
ACTION: 'action';
CONDITIONS: 'conditions';
TAIL: '...';
ERROR: 'error';

// type name
BOOL: 'bool';
BYTES: 'bytes';
INT: 'int';
ADDRESS: 'address';
ARRAY: 'array';
MAP: 'map';
MONEY: 'money';
FLOAT: 'float';
STRING: 'string';
FILE: 'file';

Question: '?';

Identifier:
	UnicodeLiteral (Literal | UnicodeDigits)*; //-> mode(NLSEMI);

DollarIdentifier: '$' Identifier;

AtIdentifier: '@' UnicodeDigits* (UnicodeLiteral Literal*)+;

DecimalNumber: (
		DecimalDigits
		| (DecimalDigits? '.' DecimalDigits)
	)
//	ExponentPart?
	;

StringLiteral: InterpretedStringLiteral | RawStringLiteral;

InterpretedStringLiteral:
	'"' DoubleQuotedStringCharacter* '"'; //-> mode(NLSEMI);

RawStringLiteral: '`' ~'`'* '`'; //-> mode(NLSEMI);

RuneLiteral:
	'\'' (UnicodeValue | BytesValue) '\''; //-> mode(NLSEMI);

HexLiteral: '0' [xX] ('_'? HexDigits)+; //-> mode(NLSEMI);

OctalLiteral: '0' [oO] ('_'? OctalDigits)+; //-> mode(NLSEMI);

BinaryLiteral: '0' [bB] ('_'? BinDigits)+; //-> mode(NLSEMI);

DecimalLiteral: DecimalDigits;

HexByteValue: '\\' 'x' HexDigits HexDigits;

OctalByteValue: '\\' OctalDigits OctalDigits OctalDigits;

BytesValue: HexByteValue | OctalByteValue;

LiteralUValue: '\\' 'u' HexDigits HexDigits HexDigits HexDigits;

BigUValue:
	'\\' 'U' HexDigits HexDigits HexDigits HexDigits HexDigits HexDigits HexDigits HexDigits;

EmptyStringLiteral: '""';

TagOptional: '"optional"';

fragment DecimalDigits: [0-9] ( '_'? [0-9])*;

fragment OctalDigits: [0-7];

fragment HexDigits: [0-9a-fA-F];

fragment BinDigits: [01];

fragment ExponentPart: [eE] [+-]? DecimalDigits;

fragment UnicodeLiteral: [\p{L}];

fragment UnicodeDigits: [\p{Nd}];

fragment UnicodeValue:
	~[\r\n']
	| LiteralUValue
	| BigUValue
	| EscapedChar;

fragment Literal: UnicodeLiteral | '_';

fragment EscapedChar:
	LiteralUValue
	| BigUValue
	| OctalByteValue
	| HexByteValue
	| '\\' [abfnrtv\\'"];

fragment DoubleQuotedStringCharacter: ~["\r\n\\] | ('\\' .);

fragment SingleQuotedStringCharacter: ~['\r\n\\] | ('\\' .);

WS: [ \t]+ -> channel(HIDDEN);
COMMENT: '/*' .*? '*/' -> channel(HIDDEN);
LINE_COMMENT: '//' ~[\r\n]* -> channel(HIDDEN);
TERMINATOR: [\r\n]+ -> channel(HIDDEN);
EOS: ([\r\n]+ | ';' | '/*' .*? '*/' | EOF) -> channel(HIDDEN);

//mode NLSEMI; WS_NLSEMI : [ \t]+ -> channel(HIDDEN); COMMENT_NLSEMI : '/*' ~[\r\n]*? '*/' ->
// channel(HIDDEN); LINE_COMMENT_NLSEMI : '//' ~[\r\n]* -> channel(HIDDEN); EOS: ([\r\n]+ | ';' |
// '/*' .*? '*/' | EOF) -> mode(DEFAULT_MODE); OTHER: -> mode(DEFAULT_MODE), channel(HIDDEN);
