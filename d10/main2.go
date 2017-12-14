package main

import (
	"fmt"
	"io/ioutil"
)

var arraySize = 256

func main() {

	var bytes, _ = ioutil.ReadFile("input.txt")

	var jumps = make([]int, len(bytes))

	for idx, char := range bytes {
		jumps[idx] = int(char)
	}
	jumps = append(jumps, 17)
	jumps = append(jumps, 31)
	jumps = append(jumps, 73)
	jumps = append(jumps, 47)
	jumps = append(jumps, 23)

	var curPos = 0
	var ar = make([]int, arraySize)

	for i := 0; i < arraySize; i++ {
		ar[i] = i
	}

	var skip = 0
	for round := 0; round < 64; round++ {
		for _, jump := range jumps {

			reverse(ar, curPos, jump)

			curPos = (curPos + jump + skip) % arraySize
			fmt.Printf("%v\n", ar)
			skip++
		}
	}

	for i := 0; i < 16; i++ {
		var result = ar[i*16]
		for j := 1; j < 16; j++ {
			result ^= ar[i*16+j]
		}
		fmt.Printf("%x", result)
	}
	fmt.Println()
	// fmt.Printf("%d\n", ar[0]*ar[1])
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
