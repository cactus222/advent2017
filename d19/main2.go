package main

import (
	"bufio"
	"fmt"
	"os"
)

var east = 0
var west = 1
var north = 2
var south = 3

func main() {

	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)
	var theMap = []string{}

	for scanner.Scan() {
		theMap = append(theMap, scanner.Text())
	}
	file.Close()
	for i := 0; i < len(theMap); i++ {
		fmt.Printf("%s\n", theMap[i])
	}
	var x = indexOf(theMap[0], '|')
	var y = 0
	var direction = south
	var maxy = len(theMap)
	var maxx = len(theMap[0])
	fmt.Println(maxy)
	fmt.Println(maxx)
	var theString = ""
	var count = 0
	for x >= 0 && y >= 0 && y < maxy && x < maxx {
		count++
		if direction == north {
			y--
		}
		if direction == south {
			y++
		}
		if direction == east {
			x++
		}
		if direction == west {
			x--
		}
		//	fmt.Println(direction)
		//if !isUseless(theMap, x, y) {
		//fmt.Println(string())
		//}
		if isUseless(theMap, x, y) {
			fmt.Println("dead")
			break
		}

		var currentVal = theMap[y][x]
		if currentVal >= 'A' && currentVal <= 'Z' {
			theString += string(currentVal)
		}

		if theMap[y][x] == '+' {
			//All directinos flipped to not return where we came from
			if y+1 < maxy && direction != north {
				if !isUseless(theMap, x, y+1) {
					direction = south
					continue
				}
			}
			if y-1 >= 0 && direction != south {
				if !isUseless(theMap, x, y-1) {
					direction = north
					continue
				}
			}
			if x-1 >= 0 && direction != east {
				if !isUseless(theMap, x-1, y) {
					direction = west
					continue
				}
			}
			if x+1 < maxx && direction != west {
				if !isUseless(theMap, x+1, y) {
					direction = east
					continue
				}
			}
		}
	}
	fmt.Println(count)
	fmt.Println(theString)
	// fmt.Println(x)
	// fmt.Println(string(theMap[y][x]))
	// var direction;

	// var x = 0;
	// var y = 0;

	// direction =  = y = direction;

}
func isUseless(theMap []string, x, y int) bool {
	return theMap[y][x] == ' '
}

func indexOf(base string, char byte) int {
	for i := 0; i < len(base); i++ {
		if base[i] == char {
			return i
		}
	}
	return -1
}
