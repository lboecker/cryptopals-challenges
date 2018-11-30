package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	a, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	b, _ := hex.DecodeString("686974207468652062756c6c277320657965")
	c := fixedXor(a, b)
	fmt.Println(hex.EncodeToString(c))
}

func fixedXor(a, b []byte) []byte {
	c := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i] ^ b[i]
	}
	return c
}
