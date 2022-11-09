grammar OpenFGA;

start: typedef+ EOF;

typedef: 'type' objectType=ID ('relations' relation+)? ;

relation: 'define' name=ID ':' typesOrID rewrite? ;

typesOrID: '[' relationReferences ']' # types
    | id=ID                           # computedUserset
    ;

relationReferences: head=relationReference (',' tail=relationReferences)* ;

relationReference: t=ID  # type
    | t=ID '#' r=ID      # typeAndRelation
    ;

rewrite: 'from' id=ID  # tupleToUserset
    | ('or' id=ID)+    # union
    | ('and' id=ID)+   # intersection
    | 'but not' id=ID  # exclusion
    ;

ID: [a-zA-Z_][a-zA-Z_0-9]* ;
WS: [ \t\n\r\f]+ -> skip ;
