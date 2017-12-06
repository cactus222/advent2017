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
		var min, _ = strconv.Atoi(nums[0])
		var max, _ = strconv.Atoi(nums[0])
		fmt.Println(min)
		for _, valStr := range nums {
			var val, _ = strconv.Atoi(valStr)
			if val < min {
				min = val
			}

			if val > max {
				max = val
			}
		}
		sum += max - min
	}

	fmt.Println(sum)

}
