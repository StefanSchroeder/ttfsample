PROJECT_NAME := "ttfsample"
PKG := "github.com/StefanSchroeder/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
GOLINT := $(shell go env GOPATH)/bin/golint
 
.PHONY: all dep lint vet test test-coverage build clean
 
all: build

dep: ## Get the dependencies
	@go mod download

vet: ## Run go vet
	@go vet ${PKG_LIST}

lint: ## Lint Golang files
	@${GOLINT} -set_exit_status ${PKG_LIST}

test: ## Run unittests
	@go test -short ${PKG_LIST}

test-coverage: ## Run tests with coverage
	@go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST} 
	@cat cover.out >> coverage.txt

build: dep ## Build the binary file
	@go build -o build/ttfsample $(PKG)
 
clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)/build
 
