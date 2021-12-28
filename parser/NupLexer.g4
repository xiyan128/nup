lexer grammar NupLexer;

// tokens
BACKSLASH: '\\' -> pushMode(STRICT);

fragment WS: ' ' | '\t';
NL: WS* ('\r\n' | '\r' | '\n') WS*;
DOLLAR: '$$'-> pushMode(MATH);
DOLLARS:'$' -> pushMode(MATH);

TEXT: (~[\r\n\\}$] | '\\\\' | '\\}' | '\\$')+;

CLOSE_BRACE: '}';

mode STRICT;
COMMA: ',';
QUOTE: '"';
OPEN_BRACKET: '[';
CLOSE_BRACKET: ']';
EQUALS: '=';
SPACE: [\t\r\n\f ]+ -> skip;
OPEN_BRACE: '{' -> popMode;
STR: '"' (~[\\"])* '"';
fragment DIGIT   :   ('0'..'9');
NUMBER: ('0'..'9')+ ('.' ('0'..'9')+)?;
BOOLEAN: 'true' | 'false';
ALPHANUMERIC: [a-zA-Z0-9_]+;

mode MATH;
MDOLLAR: '$$'-> popMode;
MDOLLARS:'$' -> popMode;
MTEXT: (~[$] | '\\$')+;

