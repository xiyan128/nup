lexer grammar NupLexer;

// tokens
BACKSLASH: '\\' -> pushMode(STRICT);

fragment WS: ' ' | '\t';
NL: WS* ('\r\n' | '\r' | '\n') WS*;
DOLLARS: '$$'-> pushMode(MATH_ATTR);
DOLLAR:'$' -> pushMode(MATH);
GRAVE: '`'-> pushMode(CODE);
GRAVES:'```' -> pushMode(CODE_ATTR);

TEXT: (~[\r\n\\{}$`] | '\\\\' | '\\}' | '\\$' | '\\`')+;

CLOSE_BRACE: '}';
OPEN_BRACE: '{';

mode STRICT;
COMMA: ',';
QUOTE: '"';
SOPEN_BRACE: '{' -> popMode;
OPEN_BRACKET: '[';
CLOSE_BRACKET: ']' -> popMode;
EQUALS: '=';
SPACE: (' ' | '\t' | '\f' | '\r\n' | '\r' | '\n') -> skip;
STR: '"' (~[\\"])* '"';
fragment DIGIT   :   ('0'..'9');
NUMBER: ('0'..'9')+ ('.' ('0'..'9')+)?;
BOOLEAN: 'true' | 'false';
ALPHANUMERIC: [a-zA-Z0-9]+;

mode MATH;
MDOLLARS: '$$'-> popMode;
MDOLLAR:'$' -> popMode;
MTEXT: (~[$] | '\\$')+;

mode MATH_ATTR;
MABRACKET: '['-> pushMode(STRICT);
MANL: ('\r\n' | '\r' | '\n') -> popMode, pushMode(MATH);

mode CODE;
CNL: '\r\n' | '\r' | '\n';
CGRAVE: '`'-> popMode;
CGRAVES:'```' -> popMode;
CTEXT: (~[`] | '\\`')+;
LANGUAGE: [a-zA-Z0-9_]+;

mode CODE_ATTR;
CABRACKET: '['-> pushMode(STRICT);
CANL: ('\r\n' | '\r' | '\n') -> popMode, pushMode(CODE);
