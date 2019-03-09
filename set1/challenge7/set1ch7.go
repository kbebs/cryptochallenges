package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {
	fileText64 := readFile("7.txt")
	ciphorKey := "YELLOW SUBMARINE"
	fmt.Println(string(ecb(fileText64, ciphorKey)))
}

func ecb(fText64 []byte, ck string) []byte {

	cipherText, err := base64.StdEncoding.DecodeString(string(fText64))
	if err != nil {
		fmt.Println(err)
	}

	cipher, _ := aes.NewCipher([]byte(ck))
	plainText := make([]byte, len(cipherText))
	size := 16

	for bs, be := 0, size; bs < len(cipherText); bs, be = bs+size, be+size {
		cipher.Decrypt(plainText[bs:be], cipherText[bs:be])
	}
	return plainText
}

func readFile(filename string) []byte {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return file
}
