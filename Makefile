.PHONY: help
# help:
#    Print this help message
help:
	@grep -o '^\#.*' Makefile | cut -d" " -f2-

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


.PHONY: upgrade
# upgrade:
#    Upgrade dependencies
upgrade:
	go get -u -t ./...
	go mod tidy