GOCMD=go
GOBUILD=$ (GOCMD) build
GOCLEAN=$ (GOCMD) clean
GOTEST=$ (GOCMD) test
GOGET=$ (GOCMD) -u get
GOOS=linux
GOARCH=amd64
BINARY_NAME=belee
LINTER=golangci-lint

all: test build

test:
	$(GOTEST) ./... -v

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

lint:
	$(LINTER) run