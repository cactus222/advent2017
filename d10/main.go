package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var arraySize = 256

func main() {

	var bytes, _ = ioutil.ReadFile("input.txt")
	var fileString = string(bytes)
	var jumpsStrs = strings.Split(fileString, ",")

	var maxSize = len(jumpsStrs)
	var jumps = make([]int, maxSize)
	for idx, val := range jumpsStrs {
		jumps[idx], _ = strconv.Atoi(val)
	}

	var curPos = 0
	var ar = make([]int, arraySize)

	for i := 0; i < arraySize; i++ {
		ar[i] = i
	}

	for skip, jump := range jumps {

		reverse(ar, curPos, jump)

		curPos = (curPos + jump + skip) % arraySize
		fmt.Printf("%v\n", ar)
	}

	fmt.Printf("%d\n", ar[0]*ar[1])

}

func reverse(ar []int, start, length int) {

	for i := 0; i < length/2; i++ {
		var startSlot = (start + i) % arraySize
		var endSlot = (start + length - i - 1) % arraySize
		var temp = ar[startSlot]
		ar[startSlot] = ar[endSlot]
		ar[endSlot] = temp
	}
}
