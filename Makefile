BIN ?= poller
CMD ?= ./cmd/poller

OUT_DIR := ./bin

.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: test
test:
	@go test -v -race -cover ./...

.PHONY: run
run:
	@go run cmd/poller/maing.go

.PHONY: build
build:
	@mkdir -p $(OUT_DIR)
	@go build -o $(OUT_DIR)/$(BIN) $(CMD)


.PHONY: clean
clean:
	@rm -rf $(OUT_DIR)
	@go clean -testcache
