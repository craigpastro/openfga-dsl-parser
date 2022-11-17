grammar OpenFGA;

start: typeDefinition+ EOF;

typeDefinition: 'type' objectType=ID ('relations' relation+)? ;

relation: 'define' name=ID ':' typeRestrictionOrId rewrite? ;

typeRestrictionOrId: typeRestriction # troiTypeRestriction
    | computedUserset=ID             # troiId
    ;

typeRestriction: '[' relationReferences ']' ;

relationReferences: relationReference (',' relationReferences)* ;

relationReference: t=ID  # rrType
    | t=ID '#' r=ID      # rrTypeAndRelation
    | t=ID ':*'          # rrTypeAndWildcard
    ;

rewrite: orTTU | ors | ands | exclusion ;

orTTU: 'or' computedUserset=ID 'from' tupleset=ID ;

ors: or ors* ;
or: 'or' id=ID ;

ands: and ands* ;
and: 'and' id=ID ;

exclusion: 'but not' id=ID ;


ID: [a-zA-Z_][a-zA-Z_0-9]* ;
WS: [ \t\n\r\f]+ -> skip ;
