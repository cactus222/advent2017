package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, _ := os.Open("input.txt")
	var entriesSet = make(map[string]bool, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var keys = strings.Split(scanner.Text(), " ")
		entriesSet[keys[0]] = true
	}
	file.Close()

	file, _ = os.Open("input.txt")
	scanner2 := bufio.NewScanner(file)
	for scanner2.Scan() {
		var keys = strings.Split(scanner2.Text(), " ")

		for i := 3; i < len(keys); i++ {
			var val = strings.Trim(keys[i], ",")
			entriesSet[val] = false
		}
	}

	for k, v := range entriesSet {
		if v {
			fmt.Println(k)
		}
	}
	file.Close()

}
