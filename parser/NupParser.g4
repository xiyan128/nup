parser grammar NupParser;

options { tokenVocab=NupLexer; }

newLine: NL;
blankLines: newLine newLine+;


document: newLine* block (blankLines block)* blankLines?;

block: (content newLine? content?)+;

content: command | text;

text: TEXT | MTEXT | CTEXT;
identifier: ALPHANUMERIC+;
command: cmd=DOLLARS text MDOLLARS
    | cmd=DOLLAR text MDOLLAR
    | cmd=GRAVE text CGRAVE
    | cmd=GRAVES attributes=attrs? SNL text CGRAVES
    | BACKSLASH cmd=ALPHANUMERIC+ attributes=attrs? OPEN_BRACE newLine* (inner=block (blankLines inner=block)*)? newLine* CLOSE_BRACE;
val: (STR | NUMBER | BOOLEAN);
attr :  name=identifier EQUALS value=val;
attrs: OPEN_BRACKET value=val CLOSE_BRACKET | OPEN_BRACKET (attribute=attr (COMMA attribute=attr) *)? CLOSE_BRACKET;