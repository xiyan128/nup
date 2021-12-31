lexer grammar NupLexer;

// tokens
BACKSLASH: '\\' -> pushMode(STRICT);

fragment WS: ' ' | '\t';
NL: WS* ('\r\n' | '\r' | '\n') WS*;
DOLLARS: '$$'-> pushMode(MATH);
DOLLAR:'$' -> pushMode(MATH);
GRAVE: '`'-> pushMode(CODE);
GRAVES:'```' -> pushMode(STRICT);

TEXT: (~[\r\n\\}$`] | '\\\\' | '\\}' | '\\$' | '\\`')+;

CLOSE_BRACE: '}';

mode STRICT;
COMMA: ',';
QUOTE: '"';
OPEN_BRACKET: '[';
CLOSE_BRACKET: ']';
EQUALS: '=';
SPACE: (' ' | '\t' | '\f') -> skip;
OPEN_BRACE: '{' -> popMode;
STR: '"' (~[\\"])* '"';
fragment DIGIT   :   ('0'..'9');
NUMBER: ('0'..'9')+ ('.' ('0'..'9')+)?;
BOOLEAN: 'true' | 'false';
ALPHANUMERIC: [a-zA-Z0-9]+;
SNL: ('\r\n' | '\r' | '\n')  -> popMode, pushMode(CODE);

mode MATH;
MDOLLARS: '$$'-> popMode;
MDOLLAR:'$' -> popMode;
MTEXT: (~[$] | '\\$')+;

mode CODE;
CNL: '\r\n' | '\r' | '\n';
CGRAVE: '`'-> popMode;
CGRAVES:'```' -> popMode;
CTEXT: (~[`] | '\\`')+;
LANGUAGE: [a-zA-Z0-9_]+;
