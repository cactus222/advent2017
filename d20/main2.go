package main

import (
	"bufio"
	"falcon/regex"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Node struct {
	x   int
	y   int
	z   int
	vX  int
	vY  int
	vZ  int
	acX int
	acY int
	acZ int
}

type Position struct {
	x, y, z int
}

func main() {

	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)
	//var theMap = []Node{}
	//p=<1609,-863,-779>, v=<-15,54,-69>, a=<-10,0,14>

	var reg = regex.Regex(`p=<([0-9\-]+),([0-9\-]+),([0-9\-]+)>, v=<([0-9\-]+),([0-9\-]+),([0-9\-]+)>, a=<([0-9\-]+),([0-9\-]+),([0-9\-]+)>`)

	var nodes = []Node{}
	for scanner.Scan() {
		var text = scanner.Text()
		var matches = reg.FindStringSubmatch(text)

		var newNode = Node{}
		newNode.x, _ = strconv.Atoi(matches[1])
		newNode.y, _ = strconv.Atoi(matches[2])
		newNode.z, _ = strconv.Atoi(matches[3])
		newNode.vX, _ = strconv.Atoi(matches[4])
		newNode.vY, _ = strconv.Atoi(matches[5])
		newNode.vZ, _ = strconv.Atoi(matches[6])
		newNode.acX, _ = strconv.Atoi(matches[7])
		newNode.acY, _ = strconv.Atoi(matches[8])
		newNode.acZ, _ = strconv.Atoi(matches[9])
		nodes = append(nodes, newNode)
	}
	file.Close()

	for i := 0; i < 25; i++ {
		// fmt.Println(nodes)
		var graphMap = make(map[Position]*int, 0)

		for idx, node := range nodes {
			node.vX += node.acX
			node.vY += node.acY
			node.vZ += node.acZ
			node.x += node.vX
			node.y += node.vY
			node.z += node.vZ
			nodes[idx] = node
		}

		var deleteSet = make(map[int]struct{}, 0)

		for i := len(nodes) - 1; i >= 0; i-- {
			var ref = i
			var node = nodes[i]
			if (graphMap[Position{node.x, node.y, node.z}] != nil) {
				deleteSet[i] = struct{}{}
				deleteSet[*graphMap[Position{node.x, node.y, node.z}]] = struct{}{}
			} else {
				graphMap[Position{node.x, node.y, node.z}] = &ref
			}
		}

		var keys []int
		for k := range deleteSet {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		for i := len(keys) - 1; i >= 0; i-- {
			var index = keys[i]
			nodes = append(nodes[:index], nodes[index+1:]...)
		}
	}

	fmt.Println(len(nodes))
}
