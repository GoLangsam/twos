@Echo Starting Documentation for directory:
@pwd

@go doc    			>	.\.md\d.doc.md
@go doc		| grep var	>	.\.md\d.vars.md
@go doc		| grep type	>	.\.md\d.type.md
@go doc		| grep func	>	.\.md\d.func.md
@go doc		| grep const	>	.\.md\d.const.md

@go doc -u 			>	.\.md\d.doc.u.md
@go doc -u	| grep var	>	.\.md\d.vars.u.md
@go doc -u	| grep type	>	.\.md\d.type.u.md
@go doc -u	| grep type | grep struct	>	.\.md\d.data.u.md
@go doc -u	| grep type | grep interf	>	.\.md\d.inter.u.md
@go doc -u	| grep func	>	.\.md\d.func.u.md
@go doc -u	| grep const	>	.\.md\d.const.u.md

@Echo priv types: name nest nilPair
@go doc -u . name		>	.\.md\t.name.md
@go doc -u . kind		>	.\.md\t.kind.md
@go doc -u . nest		>	.\.md\t.nest.md
@go doc -u . nilpair		>	.\.md\t.nilpair.md

@Echo Data types: ID Index Cardinality
@go doc -u . ID			>	.\.md\t.id.md
@go doc -u . index		>	.\.md\t.index.md
@go doc -u . cardinality	>	.\.md\t.cardinality.md

@Echo Interface types: Type Kind Pair Iterable
@go doc -u . type		>	.\.md\i.type.md
@go doc -u . pair		>	.\.md\i.pair.md
@go doc -u . iterable		>	.\.md\i.iterable.md

@Echo Function types: Head Tail
@go doc -u . head		>	.\.md\f.head.md
@go doc -u . tail		>	.\.md\f.tail.md

@Echo Finished Documentation for directory:
@pwd

@goto :EOF
