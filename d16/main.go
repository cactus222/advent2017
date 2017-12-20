package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var arraySize = 16

func main() {

	var bytes, _ = ioutil.ReadFile("input.txt")
	var fileString = string(bytes)
	var steps = strings.Split(fileString, ",")
	var ar = make([]string, arraySize)

	for i := 0; i < arraySize; i++ {
		ar[i] = string('a' + byte(i))
	}

	for _, jump := range steps {
		fmt.Println(jump)
		if jump[0] == 's' {
			fmt.Println("hu")
			jumpCount, err := strconv.Atoi(jump[1:])
			if err != nil {
				panic(err)
			}
			ar = append(ar[arraySize-jumpCount:], ar[:arraySize-jumpCount]...)
			//reverse(ar, arraySize-jumpCount, jumpCount)
		} else if jump[0] == 'x' {
			var exchanges = strings.Split(jump[1:], "/")
			var ex1, _ = strconv.Atoi(exchanges[0])
			var ex2, _ = strconv.Atoi(exchanges[1])
			exchange(ar, ex1, ex2)
		} else if jump[0] == 'p' {
			var partners = strings.Split(jump[1:], "/")
			var p1 = indexOf(ar, partners[0])
			var p2 = indexOf(ar, partners[1])
			exchange(ar, p1, p2)
		}

	}
	fmt.Println(ar)
}
func indexOf(ar []string, t string) int {
	for i, val := range ar {
		if val == t {
			return i
		}
	}
	return -1
}
func exchange(ar []string, s1, s2 int) {

	var temp = ar[s1]
	ar[s1] = ar[s2]
	ar[s2] = temp

}
