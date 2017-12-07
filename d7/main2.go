package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	children []*node
	name     string
	weight   int
}

var entriesSet = make(map[string]*node, 0)

func main() {

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	//var graphBase;
	//
	for scanner.Scan() {
		var keys = strings.Split(scanner.Text(), " ")
		var weight, _ = strconv.Atoi(strings.Trim(keys[1], "()"))
		entriesSet[keys[0]] = &node{name: keys[0], weight: weight}
	}

	file.Close()

	file, _ = os.Open("input.txt")
	scanner2 := bufio.NewScanner(file)
	for scanner2.Scan() {
		var keys = strings.Split(scanner2.Text(), " ")
		var key = keys[0]
		for i := 3; i < len(keys); i++ {
			var val = strings.Trim(keys[i], ",")
			entriesSet[key].children = append(entriesSet[key].children, entriesSet[val])
		}
	}
	// for k, v := range entriesSet {
	// 	validate(k)
	// 	v = v
	// }

	// validate("vvsvez")
	// validate("nbyij")
	// validate("tdxow")
	validate("ghwgd")
	file.Close()

}

func validate(key string) {
	var node = entriesSet[key]
	// var success = true
	var expectedWeight = -1
	// var nodes = []int{}
	for _, child := range node.children {

		var newWeight = getWeight(child.name)
		if expectedWeight == -1 {
			expectedWeight = newWeight
		} else {
			if expectedWeight != newWeight {
				// success = false
				fmt.Printf("incorrect child.name %v %d %d %d\n", child.name, child.weight, expectedWeight, newWeight)
			}
		}
		// weights = append(weights, newWeight)
	}
	// if !success {

	// }
}

func getWeight(key string) int {
	var node = entriesSet[key]
	var weight = node.weight
	for _, child := range node.children {
		weight += getWeight(child.name)
	}
	return weight
}
