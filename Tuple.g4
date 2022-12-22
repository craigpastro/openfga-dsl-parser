grammar Tuple;

tuple: obj = object '#' relation = ID '@' user EOF;

object: namespace = ID ':' objectID = ID;

user:
	obj = object '#' relation = ID	# userUserset
	| obj = object					# userObject
	| ID							# userID;

ID: [a-zA-Z0-9_-]+;
WS: [ \t\n\r\f]+ -> skip;
