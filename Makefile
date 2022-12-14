BINARY_NAME=server

.PHONY: build
build:
	CGO_ENABLED=0 go build -o bin/${BINARY_NAME}

.PHONY: test
test:
	go test ./... -v -cover

.PHONY: dep_verify
dep_verify:
	go mod verify

.PHONE: lint
lint:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck ./...