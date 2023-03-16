vet:
	@go vet ./...

test:
	@go test -v ./...

mearure:
	@go run ./outdated/tester
