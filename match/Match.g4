grammar Match;

prog:	expr EOF ;

// 完整表达式
expr
    : '(' expr ')'
    | notOp expr
    | expr andOp expr
    | expr orOp expr
    | baseExpr
    ;

// and/or左边或者右边的表达式
baseExpr
    : wordBaseExpr
    | nearBaseExpr
    | words
    | '(' baseExpr ')'
    ;

// 词表达式的基础形式
wordBaseExpr
    : WORD comparisonOp DIGITS
    | notOp WORD
    | wordBaseExpr andOp wordBaseExpr
    | wordBaseExpr orOp wordBaseExpr
    | '(' wordBaseExpr ')'
    | WORD
    ;

// near表达式的基础形式
nearBaseExpr
    : (words|WORD) nearOp '/' DIGITS (words|WORD)
    | '(' nearBaseExpr ')'
    ;

// 词组
words
    : WORD andOp WORD
    | WORD orOp WORD
    | '(' words ')'
    ;

// 比较符号
comparisonOp
    : '<'
    | '<='
    | '>'
    | '>='
    | '='
    ;

notOp
    : 'not'
    ;

andOp
    : 'and'
    ;

orOp
    : 'or'
    ;

nearOp
    : 'near'
    ;

// WORD与DIGITS定义有冲突
WORD
    : [a-zA-Z0-9\u4e00-\u9fa5\\.]+
    ;

DIGITS
    : [0-9]+
    ;

LINE_COMMENT
    : '//' .*? '\r'? '\n' -> skip
    ;

COMMENT
    : '/*' .*? '*/' -> skip
    ;

WS
    : [ \t\r\n]+ -> skip
    ;
