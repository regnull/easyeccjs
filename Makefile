.PHONY: genwasm

ROOT_DIR = $(dir $(realpath $(firstword $(MAKEFILE_LIST))))
OUT_DIR = $(ROOT_DIR)static
GOROOT := $(shell go env GOROOT)

genwasm:
	GOOS=js GOARCH=wasm go build -o $(OUT_DIR)/easyeccjs.wasm $(ROOT_DIR)/cmd/wasm/main.go
	cp "$(GOROOT)/misc/wasm/wasm_exec.js" $(OUT_DIR)
