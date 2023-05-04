BIN_NAME	:=  tlsg-cli
DIST      :=  target
GO        :=  go
GOLINT    :=  golangci-lint

all: fmt lint test_coverage

build:
	GOARCH=amd64 GOOS=linux go build -o $(DIST)/$(BIN_NAME)-linux main.go

clean:
	$(GO) clean

test:
	$(GO) test ./...

test_coverage:
	$(GO) test ./... -coverprofile=coverage.out

dep:
	$(GO) mod download

lint:
	$(GOLINT) run --enable-all

fmt:
	$(GO) fmt ./...
