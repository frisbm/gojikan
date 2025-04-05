.PHONY: help
# help:
#    Print this help message
help:
	@grep -o '^\#.*' Makefile | cut -d" " -f2-

.PHONY: update
# update:
#    Update spec and generate code
update:
	go run ./cmd/update.go

.PHONY: fmt
# fmt:
#    Format go code
fmt:
	goimports -local github.com/frisbm -w ./

.PHONY: lint
# lint:
#    lint the code
lint:
	golangci-lint run

.PHONY: test
# test:
#    Run the tests
test:
	go test -v ./...


.PHONY: update-deps
# update-deps:
#    Update dependencies
update-deps:
	go get -u -t ./...
	go mod tidy