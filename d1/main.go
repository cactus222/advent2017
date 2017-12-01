package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	var bytes, _ = ioutil.ReadFile("input.txt")

	//bytes = []byte("1122")

	var lastValue = bytes[len(bytes)-1]
	var sum = 0
	var zero byte = '0'

	for _, val := range bytes {
		if val == lastValue {
			sum += int(val - zero)
		}
		lastValue = val
	}

	fmt.Println(sum)

}
