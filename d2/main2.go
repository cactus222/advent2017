package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var sum = 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var nums = strings.Split(scanner.Text(), "\t")
		var numbers = make([]int, len(nums))

		for idx, valStr := range nums {
			var val, _ = strconv.Atoi(valStr)
			numbers[idx] = val
			for i := 0; i < idx; i++ {
				var larger = val
				var smaller = val
				if numbers[i] > larger {
					larger = numbers[i]
				} else {
					smaller = numbers[i]
				}

				if larger%smaller == 0 {
					sum += (larger / smaller)
					break
				}
			}
		}
	}

	fmt.Println(sum)

}
