package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("08.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		encrypted, _ := hex.DecodeString(scanner.Text())
		blocks := make([][]byte, len(encrypted)/16)
		for i := 0; i < len(encrypted)/16; i++ {
			blocks[i] = encrypted[i*16 : (i+1)*16]
		}
		if containsDupes(blocks) {
			fmt.Println(scanner.Text())
		}
	}
}

func containsDupes(blocks [][]byte) bool {
	for i, a := range blocks {
		for j, b := range blocks {
			if i != j && bytes.Equal(a, b) {
				return true
			}
		}
	}
	return false
}
