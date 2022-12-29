package main

import (
	"crypto/sha256"
	_ "crypto/sha512"
	"fmt"
	"strings"
	"syscall/js"

	"github.com/regnull/easyecc"
)

func main() {
	done := make(chan struct{}, 0)
	js.Global().Set("publicKeyFromPassword", js.FuncOf(publicKeyFromPassword))
	<-done
}

// Hash256 does two rounds of SHA256 hashing.
func Hash256(data []byte) []byte {
	h := sha256.Sum256(data)
	h1 := sha256.Sum256(h[:])
	return h1[:]
}

func publicKeyFromPassword(this js.Value, args []js.Value) any {
	password := args[0]
	salt := args[1]
	key := easyecc.NewPrivateKeyFromPassword([]byte(password.String()),
		Hash256([]byte(strings.ToLower(salt.String()))))
	keyHex := fmt.Sprintf("0x%x", key.PublicKey().SerializeCompressed())
	return js.ValueOf(keyHex)
}
