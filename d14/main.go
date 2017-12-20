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
	jumps = append(jumps, int('-'))

	// var resultsArray = make([]int, 128)

	var popCount int = 0
	for asd := 0; asd < 128; asd++ {
		var newJumps = make([]int, len(jumps))
		for idx := 0; idx < len(jumps); idx++ {
			newJumps[idx] = jumps[idx]
		}

		var stringIndex = fmt.Sprintf("%d", asd)
		var bytesString = []byte(stringIndex)
		for _, char := range bytesString {
			newJumps = append(newJumps, int(char))
		}
		newJumps = append(newJumps, 17)
		newJumps = append(newJumps, 31)
		newJumps = append(newJumps, 73)
		newJumps = append(newJumps, 47)
		newJumps = append(newJumps, 23)
		//fmt.Printf("%x\n", newJumps)
		//os.Exit(1)

		// fmt.Println(string(bytesString))

		var curPos = 0
		var ar = make([]uint64, arraySize)

		for i := 0; i < arraySize; i++ {
			ar[i] = uint64(i)
		}

		var skip = 0
		for round := 0; round < 64; round++ {
			for _, jump := range newJumps {

				reverse(ar, curPos, jump)

				curPos = (curPos + jump + skip) % arraySize
				//fmt.Printf("%v\n", ar)
				skip++
			}
		}

		for i := 0; i < 16; i++ {
			var result = ar[i*16]
			for j := 1; j < 16; j++ {
				result ^= ar[i*16+j]
			}
			//fmt.Printf("%x", result)
			//os.Exit(1)
			//fmt.Println(popcount(result))
			popCount += int(Count64(result))

		}
	}

	fmt.Println(popCount)
	// fmt.Printf("%d\n", ar[0]*ar[1])
}

func reverse(ar []uint64, start, length int) {

	for i := 0; i < length/2; i++ {
		var startSlot = (start + i) % arraySize
		var endSlot = (start + length - i - 1) % arraySize
		var temp = ar[startSlot]
		ar[startSlot] = ar[endSlot]
		ar[endSlot] = temp
	}
}

func Count64(x uint64) uint64 {
	x = (x & 0x5555555555555555) + ((x & 0xAAAAAAAAAAAAAAAA) >> 1)
	x = (x & 0x3333333333333333) + ((x & 0xCCCCCCCCCCCCCCCC) >> 2)
	x = (x & 0x0F0F0F0F0F0F0F0F) + ((x & 0xF0F0F0F0F0F0F0F0) >> 4)
	x *= 0x0101010101010101
	return ((x >> 56) & 0xFF)
}
