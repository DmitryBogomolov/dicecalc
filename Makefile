vet:
	@go vet ./...

test:
	@go test -v ./...

outdated_mearure:
	@go run ./outdated/tester

build_app:
	@mkdir -p dist && cd app && go build && mv app ../dist

build_server:
	@mkdir -p dist && cd server && go build && mv server ../dist
