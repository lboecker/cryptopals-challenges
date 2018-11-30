package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	encrypted, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	_, key := breakSingleByteXor(encrypted)
	fmt.Println(string(singleByteXor(encrypted, key)))
}
