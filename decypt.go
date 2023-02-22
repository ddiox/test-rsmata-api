package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	lzstring "github.com/daku10/go-lz-string"
)

const key_decrypt = ""   //Enter consId+secretKey+timeStamp
const response_data = "" //Enter the encrytped reponse data

func main() {
	fmt.Println(lzstring.DecompressFromEncodedURIComponent((stringDecrypt(key_decrypt, response_data))))
}

func stringDecrypt(key string, str string) string {
	// hash
	keyHash := sha256.Sum256([]byte(key))

	// iv - encrypt method AES-256-CBC expects 16 bytes - else you will get a warning
	iv := keyHash[:16]

	// create a new aes cipher using the key and iv
	block, err := aes.NewCipher(keyHash[:32])
	if err != nil {
		panic(err)
	}

	// decode the base64 string to a []byte
	ciphertext, _ := base64.StdEncoding.DecodeString(str)

	// create the decrypter
	decrypter := cipher.NewCBCDecrypter(block, iv)

	// decrypt
	decrypted := make([]byte, len(ciphertext))
	decrypter.CryptBlocks(decrypted, ciphertext)

	// remove padding
	padLen := int(decrypted[len(decrypted)-1])
	decrypted = decrypted[:len(decrypted)-padLen]

	return string(decrypted)
}
