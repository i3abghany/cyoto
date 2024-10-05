grammar KyotoGrammar;

program
    : (importStatement | cdeclStatement | classDeclaration | implBlock | functionDeclaration)* EOF;

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
    : (methodDeclaration)*
    ;

methodDeclaration
    : 'pub'? 'fn' (IDENTIFIER | '+' | '-' | '*' | '/') '(' ('self' ',')? parameterList? ')' type block
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
    : ((variableDeclaration
    | assignment
    | breakStatement
    | continueStatement
    | functionCall
    | returnStatement) + ';')
    | ifStatement
    | whileStatement
    | forStatement
    | block
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
    : literal                                   #literalExpr
    | IDENTIFIER                                #variableExpr
    | IDENTIFIER '{' initializerList '}'        #classLiteralExpr
    | IDENTIFIER '.' IDENTIFIER                 #memberAccessExpr
    | IDENTIFIER '(' argumentList ')'           #functionCallExpr
    | '(' expression ')'                        #parenthesizedExpr
    | expression '?' expression ':' expression  #ternaryExpr
    | unaryOp expression                        #unaryExpr
    | expression multiplicativeOp expression    #multiplicativeExpr
    | expression additiveOp expression          #additiveExpr
    | expression comparisonOp expression        #comparisonExpr
    ;

unaryOp
    : '+' | '-' | '!'
    ;

additiveOp
    : '+' | '-'
    ;

multiplicativeOp
    : '*' | '/' | '%'
    ;

comparisonOp
    : '==' | '!=' | '<' | '<=' | '>' | '>='
    ;

initializerList
    : IDENTIFIER ':' expression (',' IDENTIFIER ':' expression)*;

literal
    : INTEGER #intLiteral
    | FLOAT   #floatLiteral
    | STRING  #stringLiteral
    | BOOLEAN #booleanLiteral
    ;

COMMENT     : ('//' ~[\r\n]* | '/*' .*? '*/') -> skip;

INTEGER     : [0-9]+;
FLOAT       : [0-9]+'.'[0-9]+;
STRING      : '"' ~'"'* '"';
BOOLEAN     : 'true' | 'false';

IDENTIFIER : [a-zA-Z_][a-zA-Z0-9_]*;
WS          : [ \t\r\n]+ -> skip;
