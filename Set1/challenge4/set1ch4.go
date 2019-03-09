package main

import (
	"fmt"
	"encoding/hex"
	"bufio"
	"os"
	"strings"
)

func main() {
	encryptStrArr := readFile("4.txt")
	dictionary := readFile("dictionary.txt")
	fmt.Println("\nCryptopals Set 1 Challenge 4")			
	fmt.Println("\nDecrypted message: ")
	fmt.Println(decryptCiphor(encryptStrArr,dictionary))
}

func decryptCiphor(encrs []string, dicts []string) string {
	var largestWordCount int
	largestWordCount = 0
	var wordCount int
	decryptedStr := "No message"
	var messageBytes []byte
	cipherByte := "None"
	//Iterate over input file text strings
	for _, encr := range encrs { 	
		//Decode Hex Strings to Bytes 
		xorBytes, err := hex.DecodeString(encr)
		if err != nil {
    			fmt.Println(err)
		}
		//XOR over ascii range and count common words
		messageBytes = make([]byte, len(xorBytes))	
		for l := 0; l < 128; l++ { 
			for j, xorByte := range xorBytes { 		
				messageBytes[j] = (byte(l) ^ xorByte)
			}
			wordCount = 0			
			for _, dict := range dicts { 		
				if(strings.Contains(string(messageBytes), dict)){
					wordCount++
				}
			}				 		
			if wordCount > largestWordCount {
				largestWordCount = wordCount
				cipherByte = string(l)
				decryptedStr = string(messageBytes)
			}
		}
	}
	return decryptedStr
}

func readFile(filename string) []string {
    var strArr []string
    strArr = make([]string, 0)
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
	strArr = append(strArr,scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }	
    return strArr
	
}
