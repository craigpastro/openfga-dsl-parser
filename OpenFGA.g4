grammar OpenFGA;

prog: typedef* EOF;

typedef: 'type' objectType=ID ('relations' relations+)? ;

relations: 'define' relation=ID 'as' rewrite ;

rewrite: 'self'                              # this
    | computedUserset=ID 'from' tupleset=ID  # tupleToUserset
    | rewrite 'or' rewrite                   # union
    | rewrite 'and' rewrite                  # intersection
    | rewrite 'but not' rewrite              # exclusion
    | computedUserset=ID                     # computedUserset
    | '(' rewrite ')'                        # grouping
    ;
    
ID: [a-zA-Z_][a-zA-Z_0-9]* ;
WS: [ \t\n\r\f]+ -> skip ;
