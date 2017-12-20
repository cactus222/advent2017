package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
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

	var resultsArray = make([][]int, 128)
	for i := 0; i < 128; i++ {
		resultsArray[i] = make([]int, 128)
	}

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
			var theBytes = strconv.FormatInt(int64(result), 2)
			if len(theBytes) < 8 {
				theBytes = PadLeft(theBytes, "0", 8)
				// fmt.Printf("HEH %s %d %d\n", theBytes, asd, i)
			}
			//break
			for idx := 0; idx < 8; idx++ {
				if theBytes[idx] == '1' {
					//Row, col
					resultsArray[asd][i*8+idx] = -1
				}
			}

			//fmt.Println(popcount(result))
			popCount += int(Count64(result))
		}
	}
	// fmt.Println(popCount)
	// fmt.Println(resultsArray[0][0:5])
	// fmt.Println(resultsArray[1][0:5])
	// os.Exit(1)
	for x := 0; x < 128; x++ {
		for y := 0; y < 128; y++ {
			if resultsArray[x][y] == -1 {
				islands++
				resultsArray[x][y] = islands
				// fmt.Println(resultsArray[0][0:3])
				// fmt.Println(resultsArray[1][0:3])

				flood(resultsArray, x, y)

				// fmt.Println(resultsArray[0][0:3])
				// fmt.Println(resultsArray[1][0:3])
				// os.Exit(1)
			}
		}
	}
	fmt.Println(islands)
	// fmt.Printf("%d\n", ar[0]*ar[1])
}

var islands = 0

func flood(resultsArray [][]int, x int, y int) {
	if x-1 >= 0 {
		if resultsArray[x-1][y] == -1 {
			resultsArray[x-1][y] = islands
			flood(resultsArray, x-1, y)
		}
	}

	if y-1 >= 0 {
		if resultsArray[x][y-1] == -1 {
			resultsArray[x][y-1] = islands
			flood(resultsArray, x, y-1)
		}
	}

	if x+1 <= 127 {
		if resultsArray[x+1][y] == -1 {
			resultsArray[x+1][y] = islands
			flood(resultsArray, x+1, y)
		}
	}

	if y+1 <= 127 {
		if resultsArray[x][y+1] == -1 {
			resultsArray[x][y+1] = islands
			flood(resultsArray, x, y+1)
		}
	}
}

func PadLeft(str, pad string, length int) string {
	for {
		str = pad + str
		if len(str) >= length {
			return str[0:length]
		}
	}
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
