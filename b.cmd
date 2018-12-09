go generate ./...
@pause

@cd core\
@Call doc.cmd
@cd ..\

@cd internal\
@Call doc.cmd
@cd ..\

@cd pile\
@Call doc.cmd
@cd ..\

go test ./...
@pause
