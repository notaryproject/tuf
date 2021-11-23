COMMANDS=tuf-notary
BINARIES=$(addprefix bin/,$(COMMANDS))

.PHONY: all fmt vet test vendor binaries .FORCE

.FORCE:

all: fmt vet test binaries

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./...

vendor:
	go mod vendor

binaries: vendor $(BINARIES)

bin/%: .FORCE
	CGO_ENABLED=0 go build ${GO_BUILD_FLAGS} -o bin/$* ./cmd/$*
