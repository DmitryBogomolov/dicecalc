vet:
	@go vet ./...

test:
	@go test -v ./...

outdated_mearure:
	@go run ./outdated/tester
