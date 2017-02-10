package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

func main() {
	//encryptKey := "ICE"
	encryptKey := []byte("ICE")
	textBytes := readFile("5.txt")
	fmt.Println("\nCryptopals Set 1 Challenge 5")
	fmt.Println("\nEncrypted message: ")
	fmt.Println(encryptText(textBytes, encryptKey))
}

func encryptText(tBytes []byte, eKey []byte) string {
	//Iterate over input file
	var encryptedBytes []byte
	encryptedBytes = make([]byte, len(tBytes))
	encStr := string(eKey)
	for i, tByte := range tBytes {
		encryptedBytes[i] = (tByte ^ eKey[(i%len(encStr))])
	}

	return hex.EncodeToString(encryptedBytes)
}
func readFile(filename string) []byte {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return file
}
