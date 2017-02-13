package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fileText64 := readFile("6.txt")
	keysize := findKeysize(fileText64)
	fmt.Println(keysize)
	breakToBlocks(fileText64, keysize)
}

func breakToBlocks(fText64 []byte, ks int) {
	byteMap := make(map[int][]byte)
	transposedMap := make(map[int][]byte)

	b64Bytes, err := base64.StdEncoding.DecodeString(string(fText64))
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < (len(b64Bytes) / ks); i++ {
		b1 := make([]byte, ks)
		for j := 0; j < ks; j++ {
			b1[j] = b64Bytes[j+(i*ks)]
		}
		byteMap[i] = b1
	}

	for i := 0; i < ks; i++ {
		b1 := make([]byte, 0)
		for _, byteArr := range byteMap {
			b1 = append(b1, byteArr[i])
		}
		transposedMap[i] = b1
	}
	//var ciphorKey []string
	ciphorKeyArr := make([]string, 0)
	for _, byteArr := range transposedMap {
		ciphorKeyArr = append(ciphorKeyArr, decryptCiphor(byteArr))
	}
	ciphorKey := strings.Join(ciphorKeyArr, "")
	fmt.Println("Ciphor Key: ", ciphorKey)

	decryptedBytes := make([]byte, len(b64Bytes))
	ciphorBytes := []byte(ciphorKey)
	for i, b64Byte := range b64Bytes {
		decryptedBytes[i] = (b64Byte ^ ciphorBytes[(i%len(ciphorBytes))])
	}
	fmt.Println("Message: ", string(decryptedBytes))
}

func decryptCiphor(encrs []byte) string {
	largestWordCount := 0
	wordCount := 0
	var messageBytes []byte
	cipherByte := "None"

	letterMap := make(map[string]int)
	letterMap["E"] = 26
	letterMap["T"] = 25
	letterMap["A"] = 24
	letterMap["O"] = 23
	letterMap["I"] = 22
	letterMap["N"] = 21
	letterMap["S"] = 20
	letterMap["H"] = 19
	letterMap["R"] = 18
	letterMap["D"] = 17
	letterMap["L"] = 16
	letterMap["U"] = 15
	letterMap["C"] = 14
	letterMap["M"] = 13
	letterMap["F"] = 12
	letterMap["P"] = 11
	letterMap["G"] = 10
	letterMap["W"] = 9
	letterMap["Y"] = 8
	letterMap["B"] = 7
	letterMap["V"] = 6
	letterMap["K"] = 5
	letterMap["X"] = 4
	letterMap["J"] = 3
	letterMap["Q"] = 2
	letterMap["Z"] = 1
	letterMap["*"] = -50
	letterMap[" "] = 15

	messageBytes = make([]byte, len(encrs))
	for l := 0; l < 128; l++ {
		for j, xorByte := range encrs {
			messageBytes[j] = (byte(l) ^ xorByte)
		}
		wordCount = 0
		for i := range letterMap {
			upperCaseBytes := strings.ToUpper(string(messageBytes))
			if strings.Contains(upperCaseBytes, i) {
				wordCount = wordCount + letterMap[i]
			}
		}
		if wordCount > largestWordCount {
			largestWordCount = wordCount
			cipherByte = string(l)
		}
	}
	return cipherByte
}

//decode base64 text to bytes and find keysize by finding minimum hamming distance
func findKeysize(fText64 []byte) int {
	foundKeysize := 0
	hamCount1 := 5000
	localHam := 0

	b64Bytes, err := base64.StdEncoding.DecodeString(string(fText64))
	if err != nil {
		fmt.Println(err)
	}

	for i := 2; i < 42; i++ {
		r1 := make([]byte, i)
		r2 := make([]byte, i)
		r3 := make([]byte, i)
		r4 := make([]byte, i)
		r5 := make([]byte, i)
		r6 := make([]byte, i)
		r7 := make([]byte, i)
		r8 := make([]byte, i)
		for j := 0; j < i; j++ {
			r1 = append(r1, b64Bytes[j])
			r2 = append(r2, b64Bytes[j+i])
			r3 = append(r3, b64Bytes[j+(i*2)])
			r4 = append(r4, b64Bytes[j+(i*3)])
			r5 = append(r5, b64Bytes[j+(i*4)])
			r6 = append(r6, b64Bytes[j+(i*5)])
			r7 = append(r7, b64Bytes[j+(i*6)])
			r8 = append(r8, b64Bytes[j+(i*7)])
		}
		localHam = ((hammingdist(r1, r2) / i) + (hammingdist(r2, r3) / i) + (hammingdist(r3, r4) / i) + (hammingdist(r4, r5) / i) + (hammingdist(r5, r6) / i) + (hammingdist(r6, r7) / i) + (hammingdist(r7, r8) / i))
		if localHam < hamCount1 {
			hamCount1 = localHam
			foundKeysize = i
		}
	}
	return foundKeysize
}

func hammingdist(bytes1 []byte, bytes2 []byte) int {
	count := 0
	for i, byte1 := range bytes1 {
		for j := 0; j < 8; j++ {
			mask := byte(1 << uint(j))
			if (byte1 & mask) != (bytes2[i] & mask) {
				count++
			}
		}
	}
	return count
}

func readFile(filename string) []byte {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return file
}
