grammar Tuple;

tuple:
	namespace = ID ':' objectID = ID '#' relation = ID '@' user EOF;

user:
	namespace = ID ':' objectID = ID '#' relation = ID	# userUserset
	| namespace = ID ':' objectID = ID					# userObject
	| userID = ID										# userID;

ID: [a-zA-Z0-9_-]+;
WS: [ \t\n\r\f]+ -> skip;
