lexer grammar NeedleLexer;

// delimiters
LPAREN: '(';
RPAREN: ')' -> mode(NLSEMI);
COMMA: ',';
DOT: '.';
COLON: ':';
EQ: '=';
LBRACK: '[';
RBRACK: ']' -> mode(NLSEMI);
LBRACE: '{';
RBRACE: '}' -> mode(NLSEMI);
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
INC: '++' -> mode(NLSEMI);
DEC: '--' -> mode(NLSEMI);

// Keywords
CONTRACT: 'contract';
FUNC: 'func';
RETURN: 'return' -> mode(NLSEMI);
IF: 'if';
ELIF: 'elif';
ELSE: 'else';
WHILE: 'while';
TRUE: 'true' ;
FALSE: 'false' ;
VAR: 'var';
DATA: 'data';
SETTINGS: 'settings';
BREAK: 'break' -> mode(NLSEMI);
CONTINUE: 'continue' -> mode(NLSEMI);
ERRWARNING: 'warning' -> mode(NLSEMI);
ERRINFO: 'info' -> mode(NLSEMI);
NIL: 'nil' -> mode(NLSEMI);
ACTION: 'action';
CONDITIONS: 'conditions';
TAIL: '...';
ERROR: 'error' -> mode(NLSEMI);

// type name
BOOL: 'bool' ;
BYTES: 'bytes' ;
INT: 'int' ;
ADDRESS: 'address' ;
ARRAY: 'array' ;
MAP: 'map' ;
MONEY: 'money' ;
FLOAT: 'float' ;
STRING: 'string' ;
FILE: 'file' ;

Identifier: UnicodeLiteral (Literal | UnicodeDigits)* -> mode(NLSEMI);

DollarIdentifier: '$' Identifier -> mode(NLSEMI);

AtIdentifier: '@' UnicodeDigits* Identifier;

InterpretedStringLiteral: '"' DoubleQuotedStringCharacter* '"' -> mode(NLSEMI);

RawStringLiteral: '`' ~'`'* '`' -> mode(NLSEMI);

DecimalLiteral: DecimalDigits+ -> mode(NLSEMI);

FloatLiteral:
    (DecimalDigits ('.' DecimalDigits)? ExponentPart?
    | '.' DecimalDigits ExponentPart?) -> mode(NLSEMI);

HexLiteral: '0' [xX] ('_'? HexDigits)+ -> mode(NLSEMI);

OctalLiteral: '0' [oO] ('_'? OctalDigits)+ -> mode(NLSEMI);

BinaryLiteral: '0' [bB] ('_'? BinDigits)+ -> mode(NLSEMI);

RuneLiteral: '\'' (UnicodeValue | BytesValue) '\'' -> mode(NLSEMI);

HexByteValue: '\\' 'x' HexDigits HexDigits;

OctalByteValue: '\\' OctalDigits OctalDigits OctalDigits;

BytesValue: HexByteValue | OctalByteValue;

LittleUValue: '\\' 'u' HexDigits HexDigits HexDigits HexDigits -> mode(NLSEMI);

BigUValue:
	'\\' 'U' HexDigits HexDigits HexDigits HexDigits HexDigits HexDigits HexDigits HexDigits;

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

WS: [ \t\r\n]+ -> skip;
COMMENT: '/*' .*? '*/' -> channel(HIDDEN) ;
LINE_COMMENT: '//' ~[\r\n]* -> channel(HIDDEN);
TERMINATOR   : [\r\n]+       -> channel(HIDDEN);

mode NLSEMI;

WS_NLSEMI: [ \t]+ -> skip;
COMMENT_NLSEMI: '/*' ~[\r\n]*? '*/' -> channel(HIDDEN);
LINE_COMMENT_NLSEMI: '//' ~[\r\n]* -> channel(HIDDEN);
EOS: ([\r\n]+ | SEMI | '/*' .*? '*/' | EOF) -> mode(DEFAULT_MODE);
OTHER: -> mode(DEFAULT_MODE), channel(HIDDEN);
