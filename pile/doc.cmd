@Echo Starting Documentation for directory:
@pwd

@go doc    			>	.\.md\doc.md
@go doc -u 			>	.\.md\docu.md
@go doc -u	| grep func	>	.\.md\func.md
@go doc -u	| grep var	>	.\.md\vars.md
@go doc -u	| grep const	>	.\.md\const.md

@Echo Generic types: PileOf ...
@go doc -u . onesofpile		>	.\.md\g.pile.OnesOf.md
@go doc -u . twosofpile		>	.\.md\g.pile.twosOf.md
@go doc -u . lookuppile		>	.\.md\g.pile.LookUp.md
@go doc -u . pileofpile		>	.\.md\g.pile.PileOf.md

@Echo Finished Documentation for directory:
@pwd

@goto :EOF
