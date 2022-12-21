grammar Tuple;

tuple:
	namespace = ID ':' objectId = ID '#' relation = ID '@' user EOF;

user:
	namespace = ID ':' objectId = ID '#' relation = ID	# userUserset
	| namespace = ID ':' objectId = ID					# userObject
	| userId = ID										# userId;

ID: [a-zA-Z0-9_-]+;
WS: [ \t\n\r\f]+ -> skip;
