package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	// Load variabel from file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	data := os.Getenv("CONS_ID")         //Enter consID in .env file
	secretKey := os.Getenv("SECRET_KEY") //Enter secretKey in .env file

	// Computes the timestamp
	tStamp := strconv.FormatInt(time.Now().Unix()-time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), 10)

	// Computes the signature by hashing the salt with the secret key as the key
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data + "&" + tStamp))
	signature := h.Sum(nil)

	// base64 encode and urlencode
	encodedSignature := base64.StdEncoding.EncodeToString(signature)

	fmt.Println("X-cons-id:", data)
	fmt.Println("X-timestamp:", tStamp)
	fmt.Println("X-signature:", encodedSignature)

}
