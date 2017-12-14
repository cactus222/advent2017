package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	children []Node
	val      int
}

func main() {

	file, _ := os.Open("input.txt")
	var entriesSet = make(map[string][]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var keys = strings.Split(scanner.Text(), " ")
		entriesSet[keys[0]] = keys[2:]
	}
	file.Close()
	fmt.Printf("%v\n", entriesSet)

	var visitedNodes = make(map[string]bool, 0)

	var list = []string{"0"}
	for i := 0; len(list) > 0; i++ {
		var node string
		node, list = list[0], list[1:]
		node = strings.Trim(node, ",")

		if visitedNodes[node] {
			continue
		}
		visitedNodes[node] = true
		list = append(list, entriesSet[node]...)
	}

	fmt.Println(len(visitedNodes))

}
