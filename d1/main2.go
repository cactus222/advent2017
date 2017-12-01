package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	var bytes, _ = ioutil.ReadFile("input.txt")

	//bytes = []byte("1122")

	var sum = 0
	var zero byte = '0'
	var halfLen = len(bytes) / 2
	for i := 0; i < halfLen; i++ {
		if bytes[i] == bytes[i+halfLen] {
			sum += int(bytes[i] - zero)
		}
	}

	fmt.Println(sum * 2)

}
