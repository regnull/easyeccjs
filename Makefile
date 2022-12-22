.PHONY: genwasm

ROOT_DIR = $(dir $(realpath $(firstword $(MAKEFILE_LIST))))
OUT_DIR = $(ROOT_DIR)static

genwasm:
	GOOS=js GOARCH=wasm go build -o $(OUT_DIR)/main.wasm $(ROOT_DIR)/cmd/wasm/main.go
