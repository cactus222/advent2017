package main

import "fmt"

var arraySize = 16

func main() {

	var steps = 356

	var ar = make([]int, 0)
	ar = append(ar, 0)
	var currentPosition = 0
	for i := 1; i < 2018; i++ {

		currentPosition += steps
		currentPosition = currentPosition%i + 1

		ar = append(ar, i)
		copy(ar[currentPosition+1:], ar[currentPosition:])
		ar[currentPosition] = i

		//fmt.Println(ar)
		//fmt.Println(currentPosition)
	}
	fmt.Println(ar[currentPosition+1])
}
