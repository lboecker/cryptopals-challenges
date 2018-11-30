package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math"
	"math/bits"
)

func main() {
	contents, _ := ioutil.ReadFile("06.txt")
	decoded := make([]byte, base64.StdEncoding.DecodedLen(len(contents)))
	base64.StdEncoding.Decode(decoded, contents)
	key := breakRepeatingKeyXor(decoded, 2, 40)
	fmt.Println(string(repeatingKeyXor(decoded, key)))
}

// The result isn't perfect yet. Maybe add some padding.
func breakRepeatingKeyXor(text []byte, minKeyLen, maxKeyLen int) []byte {
	keySize := findKeySize(text, minKeyLen, maxKeyLen)
	blocks := make([][]byte, len(text)/keySize)
	for i := 0; i < len(text)/keySize; i++ {
		blocks[i] = text[i*keySize : (i+1)*keySize]
	}
	transposed := make([][]byte, keySize)
	for i := 0; i < keySize; i++ {
		t := make([]byte, len(blocks))
		for j := 0; j < len(blocks); j++ {
			t[j] = blocks[j][i]
		}
		transposed[i] = t
	}
	key := make([]byte, keySize)
	for i := 0; i < keySize; i++ {
		_, k := breakSingleByteXor(transposed[i])
		key[i] = k
	}
	return key
}

func findKeySize(text []byte, minSize, maxSize int) int {
	var size int
	smallest := math.MaxFloat64
	for i := minSize; i <= maxSize; i++ {
		dist := 0.0
		for j := 0; j < 10; j++ {
			a := text[j*i : (j+1)*i]
			b := text[(j+1)*i : (j+2)*i]
			dist += float64(hammingDist(a, b)) / float64(i)
		}
		dist /= 10
		if dist < smallest {
			smallest = dist
			size = i
		}
	}
	return size
}

func hammingDist(a, b []byte) int {
	dist := 0
	for i := 0; i < len(a); i++ {
		dist += bits.OnesCount8(a[i] ^ b[i])
	}
	return dist
}
