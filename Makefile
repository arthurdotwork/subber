test:
	go test ./... -v

lint:
	@golangci-lint run --out-format=github-actions

lint-fix:
	@golangci-lint run --fix