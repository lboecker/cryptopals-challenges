package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	text := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	key := []byte("ICE")
	fmt.Println(hex.EncodeToString(repeatingKeyXor(text, key)))
}
