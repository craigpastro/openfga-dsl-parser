grammar DSL;

dsl: typeDefinition+ EOF;

typeDefinition: 'type' objectType = ID ('relations' relation+)?;

relation: 'define' name = ID typeRestriction? 'as' rewrite;

typeRestriction: ':' '[' relationReferences ']';

relationReferences: relationReference (',' relationReferences)*;

relationReference:
	t = ID				# rrType
	| t = ID '#' r = ID	# rrTypeAndRelation
	| t = ID ':*'		# rrTypeAndWildcard;

rewrite:
	'self'										# this
	| computedUserset = ID 'from' tupleset = ID	# tupleToUserset
	| rewrite 'or' rewrite						# union
	| rewrite 'and' rewrite						# intersection
	| rewrite 'but not' rewrite					# exclusion
	| computedUserset = ID						# computedUserset
	| '(' rewrite ')'							# grouping;

ID: [a-zA-Z_][a-zA-Z_0-9]*;
WS: [ \t\n\r\f]+ -> skip;
