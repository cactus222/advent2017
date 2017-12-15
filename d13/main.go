package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	children []Node
	val      int
}

func main() {

	file, _ := os.Open("input.txt")
	var ar = make([]int, 100)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		var vals = strings.Split(scanner.Text(), ": ")
		var index, _ = strconv.Atoi(vals[0])
		var size, _ = strconv.Atoi(vals[1])
		ar[index] = size
	}
	file.Close()
	fmt.Printf("%v\n", ar)

	var severity = 0

	for i := 0; i < 100; i++ {
		var size = ar[i]
		if size == 0 {
			continue
		}
		var travelDis = (size - 1) * 2
		if i%travelDis == 0 {
			severity += size * i
			fmt.Printf("Caught %d %d * %d\n", severity, i, size)
		}
	}

	fmt.Println(severity)
	// var visitedNodes = make(map[string]bool, 0)

	// var list = []string{"0"}
	// for i := 0; len(list) > 0; i++ {
	// 	var node string
	// 	node, list = list[0], list[1:]
	// 	node = strings.Trim(node, ",")

	// 	if visitedNodes[node] {
	// 		continue
	// 	}
	// 	visitedNodes[node] = true
	// 	list = append(list, entriesSet[node]...)
	// }

	// fmt.Println(len(visitedNodes))

}
