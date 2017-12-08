package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, _ := os.Open("input.txt")
	var entriesSet = make(map[string]int, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var lineVals = strings.Split(scanner.Text(), " ")
		var changeRegister = lineVals[0]
		var sign int
		if lineVals[1] == "inc" {
			sign = 1
		} else {
			sign = -1
		}
		var value, _ = strconv.Atoi(lineVals[2])
		//if
		var conditionRegister = lineVals[4]
		var operator = lineVals[5]
		var conditionVal, _ = strconv.Atoi(lineVals[6])

		if operator == ">" {
			if entriesSet[conditionRegister] > conditionVal {
				var original = entriesSet[changeRegister]
				entriesSet[changeRegister] = original + sign*value
			}
		} else if operator == ">=" {
			if entriesSet[conditionRegister] >= conditionVal {
				var original = entriesSet[changeRegister]
				entriesSet[changeRegister] = original + sign*value
			}
		} else if operator == "==" {
			if entriesSet[conditionRegister] == conditionVal {
				var original = entriesSet[changeRegister]
				entriesSet[changeRegister] = original + sign*value
			}
		} else if operator == "!=" {
			if entriesSet[conditionRegister] != conditionVal {
				var original = entriesSet[changeRegister]
				entriesSet[changeRegister] = original + sign*value
			}
		} else if operator == "<" {
			if entriesSet[conditionRegister] < conditionVal {
				var original = entriesSet[changeRegister]
				entriesSet[changeRegister] = original + sign*value
			}
		} else if operator == "<=" {
			if entriesSet[conditionRegister] <= conditionVal {
				var original = entriesSet[changeRegister]
				entriesSet[changeRegister] = original + sign*value
			}
		}

	}
	var max = 0
	for k, v := range entriesSet {
		fmt.Printf("%v %v\n", k, v)
		if v > max {
			max = v
		}
	}
	fmt.Printf("ans %d\n", max)
	file.Close()
}
