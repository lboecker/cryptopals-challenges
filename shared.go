package main

import (
	"math"
)

// http://www.data-compression.com/english.html
var frequencies = []float64{
	0.0651738, // A
	0.0124248, // B
	0.0217339, // C
	0.0349835, // D
	0.1041442, // E
	0.0197881, // F
	0.0158610, // G
	0.0492888, // H
	0.0558094, // I
	0.0009033, // J
	0.0050529, // K
	0.0331490, // L
	0.0202124, // M
	0.0564513, // N
	0.0596302, // O
	0.0137645, // P
	0.0008606, // Q
	0.0497563, // R
	0.0515760, // S
	0.0729357, // T
	0.0225134, // U
	0.0082903, // V
	0.0171272, // W
	0.0013692, // X
	0.0145984, // Y
	0.0007836, // Z
	0.1918182, // Space
}

func singleByteXor(text []byte, b byte) []byte {
	xored := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		xored[i] = text[i] ^ b
	}
	return xored
}

func breakSingleByteXor(text []byte) (float64, byte) {
	var key byte
	lowest := math.MaxFloat64
	for i := 0; i <= 255; i++ {
		score := scoreText(singleByteXor(text, byte(i)))
		if score < lowest {
			lowest = score
			key = byte(i)
		}
	}
	return lowest, key
}

func repeatingKeyXor(text, key []byte) []byte {
	xored := make([]byte, len(text))
	var keyIndex int
	for i, b := range text {
		xored[i] = b ^ key[keyIndex]
		if keyIndex+1 < len(key) {
			keyIndex++
		} else {
			keyIndex = 0
		}
	}
	return xored
}

// Lower scores are better
func scoreText(text []byte) float64 {
	freqs := make([]float64, 27)
	nonPrintable := 0.0
	for _, b := range text {
		if b >= 0 && b <= 31 || b == 127 {
			nonPrintable++
		} else if b == 32 {
			freqs[26]++
		} else if b >= 65 && b <= 90 {
			freqs[b-65]++
		} else if b >= 97 && b <= 122 {
			freqs[b-97]++
		}
	}
	for i, f := range freqs {
		freqs[i] = f / float64(len(text))
	}
	nonPrintable /= float64(len(text))
	return euclideanDist(frequencies, freqs) + nonPrintable
}

func euclideanDist(a, b []float64) float64 {
	var dist float64
	for i := 0; i < len(a); i++ {
		tmp := a[i] - b[i]
		dist += tmp * tmp
	}
	return math.Sqrt(dist)
}
