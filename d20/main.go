package main

import (
	"bufio"
	"falcon/regex"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Node struct {
	x int
	y int
	z int
}

func main() {

	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)
	//var theMap = []Node{}
	//p=<1609,-863,-779>, v=<-15,54,-69>, a=<-10,0,14>
	var i = 0

	var reg = regex.Regex(`p=<([0-9\-]+),([0-9\-]+),([0-9\-]+)>, v=<([0-9\-]+),([0-9\-]+),([0-9\-]+)>, a=<([0-9\-]+),([0-9\-]+),([0-9\-]+)>`)

	var min float64 = 99999
	var maxIndex = 0

	for scanner.Scan() {
		var text = scanner.Text()
		var matches = reg.FindStringSubmatch(text)
		var acX, acY, acZ int
		var err error

		acX, err = strconv.Atoi(matches[7])
		if err != nil {
			panic(err)
		}
		acY, err = strconv.Atoi(matches[8])
		if err != nil {
			panic(err)
		}
		acZ, err = strconv.Atoi(matches[9])
		if err != nil {
			panic(err)
		}
		var accTotal = math.Abs(float64(acX)) + math.Abs(float64(acY)) + math.Abs(float64(acZ))
		//fmt.Println(accTotal)
		if accTotal < min {
			min = accTotal
			maxIndex = i
			fmt.Println("BLARGH")
			fmt.Println(i)
			fmt.Println(accTotal)
			fmt.Printf("%d,%d,%d\n", acX, acY, acZ)
		}
		// for _, val := range matches {
		// 	fmt.Println(val)
		// }
		// fmt.Println(text)
		i++
		// break
	}
	fmt.Println()
	fmt.Println(min)
	fmt.Println(maxIndex)
	file.Close()
}
