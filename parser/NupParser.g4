parser grammar NupParser;

options { tokenVocab=NupLexer; }

newLine: NL | MANL | CANL;
blankLines: newLine newLine+;


document: newLine* block (blankLines block)* blankLines?;

block: (content newLine? content?)+;

content: command | text;

text: TEXT | MTEXT | CTEXT;
identifier: ALPHANUMERIC+;
openBracket: (OPEN_BRACKET | CABRACKET | MABRACKET);
openBrace: (OPEN_BRACE | SOPEN_BRACE);
command:
      cmd=DOLLARS attributes=attrs? newLine text MDOLLARS
    | cmd=DOLLAR text MDOLLAR
    | cmd=GRAVES attributes=attrs? newLine text CGRAVES
    | cmd=GRAVE text CGRAVE
    | BACKSLASH cmd=ALPHANUMERIC+ attributes=attrs? openBrace newLine* (inner=block (blankLines inner=block)*)? newLine* CLOSE_BRACE;
val: (STR | NUMBER | BOOLEAN);
attr :  name=identifier EQUALS value=val;
attrs: openBracket (attribute=attr (COMMA attribute=attr) *)? CLOSE_BRACKET;