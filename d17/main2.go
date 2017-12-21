package main

import "fmt"

func main() {

	var steps = 356

	var ar = make([]int, 0)
	ar = append(ar, 0)
	var currentPosition = 0
	var ans = 0
	for i := 1; i < 50000001; i++ {
		currentPosition += steps
		currentPosition = currentPosition%i + 1

		if currentPosition == 1 {
			ans = i
		}
	}
	fmt.Println(ans)
}
