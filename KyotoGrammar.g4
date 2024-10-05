grammar KyotoGrammar;

program
    : (importStatement | cdeclStatement | classDeclaration | implBlock | functionDeclaration)*;

importStatement
    : 'import' IDENTIFIER ';'
    ;

cdeclStatement
    : 'cdecl' IDENTIFIER '(' parameterList? (',' '...')? ')' type ';'
    ;

classDeclaration
    : 'class' IDENTIFIER '{' classBody '}'
    ;

classBody
    : ('pub'? variableDeclaration ';')*
    ;

implBlock
    : 'impl' IDENTIFIER '{' implBody '}'
    ;

implBody
    : ('pub'? functionDeclaration)*
    ;

functionDeclaration
    : 'fn' IDENTIFIER '(' parameterList? ')' type block
    ;

parameterList
    : parameter (',' parameter)*
    | parameterList ',' '...'
    ;

parameter
    : type IDENTIFIER
    ;

type
    : 'i8'  | 'i16' | 'i32' | 'i64'
    | 'u8'  | 'u16' | 'u32' | 'u64'
    | 'f32' | 'f64'
    | 'str'
    | 'bool'
    | 'void'
    | IDENTIFIER
    ;

statement
    : (variableDeclaration
    | assignment
    | functionCall
    | returnStatement
    | ifStatement
    | whileStatement
    | forStatement
    | breakStatement
    | continueStatement
    | block) ';'
    ;

variableDeclaration
    : type IDENTIFIER ('=' expression)?
    ;

assignment
    : IDENTIFIER '=' expression
    ;

functionCall
    : IDENTIFIER '(' argumentList? ')'
    ;

argumentList
    : expression (',' expression)*
    ;

returnStatement
    : 'return' expression?
    ;

ifStatement
    : 'if' '(' expression ')' statement ('else' statement)?
    ;

whileStatement
    : 'while' '(' expression ')' statement
    ;

forStatement
    : 'for' '(' variableDeclaration? ';' expression? ';' expression? ')' statement
    ;

breakStatement
    : 'break'
    ;

continueStatement
    : 'continue'
    ;

block
    : '{' statement* '}'
    ;

expression
    : literal
    | IDENTIFIER
    | IDENTIFIER '.' IDENTIFIER
    | functionCall
    | expression '?' expression ':' expression
    | '(' expression ')'
    | expression ('+' | '-' | '*' | '/' | '%' | '==' | '!=' | '<' | '<=' | '>' | '>=') expression
    | ('+' | '-' | '!') expression
    ;

literal
    : INTEGER
    | FLOAT
    | STRING
    | BOOLEAN
    ;

INTEGER     : [0-9]+;
FLOAT       : [0-9]+'.'[0-9]+;
STRING      : '"' ~'"'* '"';
BOOLEAN     : 'true' | 'false';

IDENTIFIER : [a-zA-Z_][a-zA-Z0-9_]*;
WS          : [ \t\r\n]+ -> skip;
