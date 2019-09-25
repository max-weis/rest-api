default: all

all: build test run

build:
    go build main.go

test:
    go test ./...

run:
    go run main