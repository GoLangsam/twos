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
@go doc -u	| grep func	>	.\.md\d.func.u.md
@go doc -u	| grep const	>	.\.md\d.const.u.md

@Echo Generic types: PileOf ...
@go doc -u . onesofanytype	>	.\.md\g.any.OnesOf.md
@go doc -u . twosofanytype	>	.\.md\g.any.TwosOf.md
@go doc -u . lookupanytype	>	.\.md\g.any.LookUp.md
@go doc -u . pileofanytype	>	.\.md\g.any.PileOf.md
@go doc -u . anytypes		>	.\.md\g.any.S.md
@go doc -u . anytypeis		>	.\.md\g.any.Is.md

@Echo Finished Documentation for directory:
@pwd

@goto :EOF
