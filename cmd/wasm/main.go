package main

import (
	_ "crypto/sha512"
	"syscall/js"

	"github.com/regnull/easyecc"
)

func main() {
	done := make(chan struct{}, 0)
	js.Global().Set("wasmHash", js.FuncOf(publicKeyFromPassword))
	<-done
}

func publicKeyFromPassword(this js.Value, args []js.Value) any {
	password := args[0]
	salt := args[1]
	key := easyecc.NewPrivateKeyFromPassword([]byte(password.String()), []byte(salt.String()))
	return key.PublicKey().SerializeCompressed()
}
