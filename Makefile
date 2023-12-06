BIN=cidr
GO=go
GO_BUILD_LD_FLAGS=-ldflags="-s -w -extldflags=-static"
GO_BUILD_DEBUG_FLAGS=-gcflags=all="-N -l"

.PHONY: build build-debug build-race test up

build:
	CGO_ENABLED=0 ${GO} build -${GO_BUILD_LD_FLAGS} ${BIN}.go

build-debug:
	CGO_ENABLED=0 ${GO} build -${GO_BUILD_DEBUG_FLAGS} ${BIN}.go

test:
	${GO} test -covermode=count -coverprofile=coverage.out ./... && \
	${GO} tool cover -html=coverage.out
