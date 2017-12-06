package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	var bytes, _ = ioutil.ReadFile("input.txt")
	var fileString = string(bytes)
	var bucketStrs = strings.Split(fileString, "\t")

	var maxSize = len(bucketStrs)
	var buckets = make([]int, maxSize)
	for idx, val := range bucketStrs {
		buckets[idx], _ = strconv.Atoi(val)
	}

	var history = make(map[string]int, 0)
	history[genHistory(buckets)] = 0

	for step := 1; ; step++ {
		var max = buckets[0]
		var maxIndex = 0
		for i := 0; i < maxSize; i++ {
			if buckets[i] > max {
				maxIndex = i
				max = buckets[i]
			}
		}
		buckets[maxIndex] = 0
		for i := (maxIndex + 1) % maxSize; max > 0; {
			buckets[i]++
			i = (i + 1) % maxSize
			max--
		}
		historyKey := genHistory(buckets)
		if val, exists := history[historyKey]; exists {

			fmt.Println(step)
			fmt.Println(val)
			fmt.Println(step - val)
			break
		}
		history[historyKey] = step
	}

}

func genHistory(ar []int) string {
	var ans = fmt.Sprintf("%v", ar)
	fmt.Println(ans)
	return ans
}
