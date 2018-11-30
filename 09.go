package main

import (
	"fmt"
)

func main() {
	padded := addPadding([]byte("YELLOW SUBMARINE"), 20)
	fmt.Printf("%q\n", padded)
}

// PKCS#7
// https://tools.ietf.org/html/rfc5652#section-6.3
func addPadding(block []byte, blockSize int) []byte {
	if len(block)%blockSize == 0 {
		return block
	} else {
		padded := make([]byte, blockSize)
		val := byte(blockSize - (len(block) % blockSize))
		for i := 0; i < blockSize; i++ {
			if i < len(block) {
				padded[i] = block[i]
			} else {
				padded[i] = val
			}
		}
		return padded
	}
}
