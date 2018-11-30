package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {
	contents, _ := ioutil.ReadFile("07.txt")
	decoded := make([]byte, base64.StdEncoding.DecodedLen(len(contents)))
	base64.StdEncoding.Decode(decoded, contents)
	key := []byte("YELLOW SUBMARINE")
	fmt.Println(string(decryptAES128ECB(decoded, key)))
}

func decryptAES128ECB(encrypted, key []byte) []byte {
	cipher, _ := aes.NewCipher(key)
	decrypted := make([]byte, len(encrypted))
	for i := 0; i < len(encrypted); i += len(key) {
		cipher.Decrypt(decrypted[i:i+len(key)], encrypted[i:i+len(key)])
	}
	return decrypted
}
