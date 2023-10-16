lexer grammar NeedleLexer;

// delimiters
LPAREN: '(';
RPAREN: ')';
COMMA: ',';
DOT: '.';
COLON: ':';
EQ: '=';
LBRACK: '[';
RBRACK: ']';
LBRACE: '{';
RBRACE: '}';
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
INC: '++';
DEC: '--';

// Keywords
CONTRACT: 'contract';
FUNC: 'func';
RETURN: 'return';
IF: 'if';
ELIF: 'elif';
ELSE: 'else';
WHILE: 'while';
TRUE: 'true';
FALSE: 'false';
VAR: 'var';
DATA: 'data';
SETTINGS: 'settings';
BREAK: 'break';
CONTINUE: 'continue';
ERRWARNING: 'warning';
ERRINFO: 'info';
NIL: 'nil';
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

Identifier: UnicodeLiteral (Literal | UnicodeDigits)*;

DollarIdentifier: '$' Identifier;

AtIdentifier: '@' UnicodeDigits* (UnicodeLiteral Literal*)+;

StringLiteral: InterpretedStringLiteral | RawStringLiteral;

InterpretedStringLiteral: '"' DoubleQuotedStringCharacter* '"';

RawStringLiteral: '`' ~'`'* '`';

DecimalLiteral: DecimalDigits+;

FloatLiteral:
    (DecimalDigits ExponentPart)
    | (DecimalDigits '.' DecimalDigits ExponentPart?)
    | ('.' DecimalDigits ExponentPart?)
    | (DecimalDigits '.' DecimalDigits);

HexLiteral: '0' [xX] ('_'? HexDigits)+;

OctalLiteral: '0' [oO] ('_'? OctalDigits)+;

BinaryLiteral: '0' [bB] ('_'? BinDigits)+;

RuneLiteral: '\'' (UnicodeValue | BytesValue) '\'';

HexByteValue: '\\' 'x' HexDigits HexDigits;

OctalByteValue: '\\' OctalDigits OctalDigits OctalDigits;

BytesValue: HexByteValue | OctalByteValue;

LittleUValue: '\\' 'u' HexDigits HexDigits HexDigits HexDigits;

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
	| LittleUValue
	| BigUValue
	| EscapedChar;

fragment Literal: UnicodeLiteral | '_';

fragment EscapedChar: '\\' [abfnrtv\\'"];

fragment DoubleQuotedStringCharacter: ~["\r\n\\] | ('\\' .);

fragment SingleQuotedStringCharacter: ~['\r\n\\] | ('\\' .);

WS: [ \t]+ -> channel(HIDDEN);
COMMENT: '/*' .*? '*/' -> channel(HIDDEN);
LINE_COMMENT: '//' ~[\r\n]* -> channel(HIDDEN);
TERMINATOR: [\r\n]+ -> channel(HIDDEN);
EOS: ([\r\n]+ | ';' | '/*' .*? '*/' | EOF) -> channel(HIDDEN);
