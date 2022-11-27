grammar Match;

prog:	expr EOF ;

// 完整表达式
expr
    : expr andOp expr
    | expr orOp expr
    | notOp leftOp expr rightOp
    | leftOp expr rightOp
    | cmpExpr
    | nearExpr
    | wordExpr
    ;

// 比较表达式
cmpExpr
    : word cmpOp digits
    | leftOp cmpExpr rightOp
    ;

// near表达式
nearExpr
    : (words|word) nearOp '/' digits (words|word)
    | leftOp nearExpr rightOp
    ;

// 单个词表单式
wordExpr
    : word
    | leftOp wordExpr rightOp
    ;

// 词组
words
    : word andOp word
    | word orOp word
    | leftOp words rightOp
    ;

// token
word: WORD;
digits: DIGITS;

// 比较符号
cmpOp
    : '<'
    | '<='
    | '>'
    | '>='
    | '='
    ;
notOp: 'not' ;
andOp: 'and' ;
orOp: 'or' ;
nearOp: 'near' ;
leftOp: '(';
rightOp: ')';

// token
// 预处理：前后加上引号，避免和digits歧义
WORD
    : '"' [a-zA-Z0-9\u4e00-\u9fa5\\.]+? '"'
    | '\'' [a-zA-Z0-9\u4e00-\u9fa5\\.]+? '\''
    ;

DIGITS: [0-9]+ ;

LINE_COMMENT: '//' .*? '\r'? '\n' -> skip ;
COMMENT: '/*' .*? '*/' -> skip ;
WS: [ \t\r\n]+ -> skip ;
