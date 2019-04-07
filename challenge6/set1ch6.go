package main

import "fmt"

//"io/ioutil"

func main() {
	bytesOne := []byte("this is a test")
	bytesTwo := []byte("wokka wokka!!!")
	//textBytes := readFile("5.txt")
	fmt.Println(hammingdist(bytesOne, bytesTwo))
}

func hammingdist(bytes1 []byte, bytes2 []byte) int {
	count := 0
	for i, byte1 := range bytes1 {
		for j := 0; j < 8; j++ {
			byte(j)
			mask := byte(1 << uint(j))
			if (byte1 & mask) != (bytes2[i] & mask) {
				count++
			}
		}
	}
	return count
}
