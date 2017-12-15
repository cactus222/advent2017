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

	var delay = 0
	var bad = false
	for {
		delay++
		severity = 0
		bad = false
		for i := 0; i < 100; i++ {
			var size = ar[i]
			if size == 0 {
				continue
			}
			var travelDis = (size - 1) * 2
			if (i+delay)%travelDis == 0 {
				bad = true
				//fmt.Printf("Caught %d %d * %d\n", severity, i, size)
				break
			}
		}
		if !bad {
			fmt.Printf("DONE %d\n", delay)
			os.Exit(0)
		}
	}
}
