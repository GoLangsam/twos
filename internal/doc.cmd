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
@go doc -u . onesofanytype	>	.\.md\g.pile.OnesOf.md
@go doc -u . lookupanytype	>	.\.md\g.pile.LookUp.md
@go doc -u . lookerofanytype	>	.\.md\g.pile.LookEr.md
@go doc -u . pileofanytype	>	.\.md\g.pile.PileOf.md
@go doc -u . pilerofanytype	>	.\.md\g.pile.Piler.md
@go doc -u . anytypes		>	.\.md\g.pile.S.md
@go doc -u . anytypeis		>	.\.md\g.pile.Is.md

@Echo Finished Documentation for directory:
@pwd

@goto :EOF
