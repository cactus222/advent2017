package main

import (
	"fmt"
	"os"
)

func main() {

	// var bytes, _ = ioutil.ReadFile("input.txt")

	theArray := make([][]uint64, 200)
	for i := range theArray {
		theArray[i] = make([]uint64, 200)
	}

	var x = 100
	var y = 100
	theArray[x][y] = 1
	var size = 2

	for {
		//right one
		x = x + 1
		theArray[x][y] = findSum(theArray, x, y)

		//up size apparently -1
		for i := 0; i < size-1; i++ {
			y -= 1
			theArray[x][y] = findSum(theArray, x, y)
		}
		//left size
		for i := 0; i < size; i++ {
			x -= 1
			theArray[x][y] = findSum(theArray, x, y)
		}
		//down size
		for i := 0; i < size; i++ {
			y += 1
			theArray[x][y] = findSum(theArray, x, y)
		}
		//right size
		for i := 0; i < size; i++ {
			x += 1
			theArray[x][y] = findSum(theArray, x, y)
		}

		size += 2
		if size == 5 {
			break
		}
	}

}

func findSum(theArray [][]uint64, x int, y int) uint64 {
	var sum uint64 = 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			sum += theArray[x+i][y+j]
		}
	}
	fmt.Println(sum)
	if sum > 312051 {
		fmt.Println(sum)
		os.Exit(0)
	}
	return sum
}
