BIN := gowebdav
SRC := $(wildcard *.go) cmd/gowebdav/main.go

all: test cmd

cmd: ${BIN}

${BIN}: ${SRC}
	go build -o $@ ./cmd/gowebdav

test:
	go test -modfile=go_test.mod -v -short -cover ./...


.PHONY: all cmd clean test api check
