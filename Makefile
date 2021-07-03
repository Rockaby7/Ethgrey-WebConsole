build_agent:
	CGO_ENABLED=0 GOOS="linux" GOARCH="amd64" go build -o cmd/ethgrey_agent cmd/agent/main.go

mv_agent:
	mv cmd/ethgrey_agent deploy/app/
