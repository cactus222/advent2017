package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	var bytes, _ = ioutil.ReadFile("input.txt")

	var isCancelling = false
	var isGarbage = false
	var curDepth = 0
	var score = 0
	for _, char := range bytes {

		if isCancelling {
			isCancelling = false
			continue
		}

		if char == '<' {
			isGarbage = true
		} else if char == '>' {
			isGarbage = false
		} else if char == '!' {
			if isGarbage {
				isCancelling = true
			}
		} else if char == '{' {
			if !isGarbage {
				curDepth++
			}
		} else if char == '}' {
			if !isGarbage {
				score += curDepth
				fmt.Println(curDepth)
				curDepth--
			}
		}
	}
	fmt.Printf("score %d\n", score)
}
