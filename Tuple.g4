grammar Tuple;

tuple: object '#' relation '@' user EOF;

object: namespace = ID ':' object_id = ID;
relation: ID;
user:
	object '#' relation	# user_userset
	| object			# user_object
	| ID				# user_id;

ID: [a-zA-Z][a-zA-Z0-9_-]*;
WS: [ \t\n\r\f]+ -> skip;
