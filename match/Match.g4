grammar Match;

prog
    : expr
    ;

// 完整表达式
expr
    : halfExpr
    | expr logicalOp expr
    | notOp expr
    | '(' expr ')'
    ;

// and/or左边或者右边的表达式
halfExpr
    : wordBaseExpr
    | nearBaseExpr
    | words
    ;

// 词表达式的基础形式
wordBaseExpr
    : word
    | notOp word
    | word comparisonOp digits
    | '(' wordBaseExpr ')'
    ;

// near表达式的基础形式
nearBaseExpr
    : (word|words) nearOp '/' digits (word|words)
    | '(' nearBaseExpr ')'
    ;

// 词组
words
    : word logicalOp word
    ;

// 比较符号
comparisonOp
    : '<'
    | '>'
    | '<='
    | '>='
    | '='
    ;

logicalOp
    : andOp
    | orOp
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

digits
    : DIGIT+
    ;

word
    : CHAR+
    ;

DIGIT
    : [0-9]
    ;

CHAR
    : [a-z\\.]
    ;
