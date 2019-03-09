package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	fileHex := readFile("8.txt")
	fmt.Println(fileHex[0])
	hexBytes, err := hex.DecodeString(fileHex[0])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(hexBytes))
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
		strArr = append(strArr, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return strArr
}
