.PHONY: build fmt vet test run

# The first release targets X11 through XWayland on Linux.
GO_TAGS ?= x11

build: fmt vet
	go build -tags $(GO_TAGS) -o go_math .

fmt:
	go fmt ./...

vet:
	go vet -tags $(GO_TAGS) ./...

test:
	go test -tags ci ./...

run:
	go run -tags $(GO_TAGS) .
