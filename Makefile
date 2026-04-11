BINARY := mana
MODULE := github.com/k07g/mana
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")

.PHONY: build run test clean install uninstall help

build: ## バイナリをビルドする
	go build -ldflags "-X main.version=$(VERSION)" -o $(BINARY) .

run: ## アプリを実行する
	go run .

test: fmt vet ## テストを実行する
	go test ./...

fmt: ## コードをフォーマットする
	gofmt -w .

vet: ## 静的解析を実行する
	go vet ./...

install: ## バイナリをインストールする
	go install .

uninstall: ## バイナリをアンインストールする
	rm -f $(shell go env GOPATH)/bin/$(BINARY)

clean: ## ビルド成果物を削除する
	rm -f $(BINARY)

help: ## ヘルプを表示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
