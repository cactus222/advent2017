package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var sum = 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var passes = strings.Split(scanner.Text(), " ")
		var passSet = make(map[string]bool, 0)
		var success = true
		for _, pass := range passes {
			pass = SortString(pass)

			if passSet[pass] {
				fmt.Println("INVALID " + scanner.Text())
				success = false
				break
			} else {
				passSet[pass] = true
			}
		}
		if success {
			sum++
		}

	}

	fmt.Println(sum)

}
