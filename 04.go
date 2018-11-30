package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"math"
	"os"
)

func main() {
	file, _ := os.Open("04.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lowest := math.MaxFloat64
	var winner []byte
	for scanner.Scan() {
		encrypted, _ := hex.DecodeString(scanner.Text())
		score, key := breakSingleByteXor(encrypted)
		decrypted := singleByteXor(encrypted, key)
		if score < lowest {
			lowest = score
			winner = decrypted
		}
	}
	fmt.Println(string(winner))
}
