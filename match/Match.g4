grammar Match;

prog:	expr EOF ;

// 完整表达式
expr
    : expr andOp expr
    | expr orOp expr
    | notOp '(' expr ')'
    | '(' expr ')'
    | wordCompExpr
    | nearBaseExpr
    | word
    ;

// 带比较符号的基础形式
wordCompExpr
    : word comparisonOp digits
    | '(' wordCompExpr ')'
    ;

// near表达式的基础形式
nearBaseExpr
    : (words|word) nearOp '/' digits (words|word)
    | '(' nearBaseExpr ')'
    ;

// 词组
words
    : word andOp word
    | word orOp word
    | '(' words ')'
    ;

word: WORD;
digits: DIGITS;

// 比较符号
comparisonOp
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

// 关键词
// 预处理：前后加上引号，避免和digits歧义
WORD
    : '"' [a-zA-Z0-9\u4e00-\u9fa5\\.]+? '"'
    | '\'' [a-zA-Z0-9\u4e00-\u9fa5\\.]+? '\''
    ;

DIGITS: [0-9]+ ;

LINE_COMMENT: '//' .*? '\r'? '\n' -> skip ;
COMMENT: '/*' .*? '*/' -> skip ;
WS: [ \t\r\n]+ -> skip ;
