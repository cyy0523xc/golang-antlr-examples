grammar Match;

prog:	expr EOF ;

// 完整表达式
expr
    : leftBracket expr rightBracket
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
    | leftBracket baseExpr rightBracket
    ;

// 词表达式的基础形式
wordBaseExpr
    : word comparisonOp digits
    | notOp word
    | leftBracket wordBaseExpr rightBracket
    | word
    ;

// near表达式的基础形式
nearBaseExpr
    : (words|word) nearOp (words|word)
    | leftBracket nearBaseExpr rightBracket
    ;

// 词组
words
    : word andOp word
    | word orOp word
    | leftBracket words rightBracket
    ;

// 比较符号
comparisonOp
    : '<'
    | '<='
    | '>'
    | '>='
    | '='
    ;

leftBracket
    : '('
    ;

rightBracket
    : ')'
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
    : 'near' '/' digits
    ;

word
    : CHAR+
    ;

digits
    : DIGIT+
    ;

DIGIT
    : [0-9]
    ;

CHAR
    : [a-zA-Z0-9\u4e00-\u9fa5\\.]
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
