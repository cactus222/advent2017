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
	var instrStrs = strings.Split(fileString, "\n")

	var maxSize = len(instrStrs)
	var instrs = make([]int, maxSize)
	for idx, val := range instrStrs {
		instrs[idx], _ = strconv.Atoi(val)
	}

	var currentStep = 0
	var stepsTaken = 0
	//fmt.Printf("MAXSIZE %d\n", maxSize)
	for currentStep >= 0 && currentStep < maxSize {
		var jump = instrs[currentStep]
		instrs[currentStep]++
		currentStep += jump
		stepsTaken++
		//fmt.Println(currentStep)
	}

	fmt.Println(stepsTaken)
}
