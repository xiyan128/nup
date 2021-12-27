parser grammar NupParser;

options { tokenVocab=NupLexer; }

newLine: NL;
blankLines: newLine newLine+;


document: newLine* block (blankLines block)* blankLines?;

block: (content newLine? content?)+;

content: command | text;

text: TEXT;
identifier: ALPHANUMERIC+;
command: BACKSLASH cmd=identifier attributes=attrs? OPEN_BRACE inner=block CLOSE_BRACE;
//command: BACKSLASH;
val: (STR | NUMBER | BOOLEAN);
attr :  name=identifier EQUALS value=val;
attrs: OPEN_BRACKET value=val CLOSE_BRACKET | OPEN_BRACKET (attribute=attr (COMMA attribute=attr) *)? CLOSE_BRACKET;