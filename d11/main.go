package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

var arraySize = 256

func main() {

	var bytes, _ = ioutil.ReadFile("input.txt")
	var fileString = string(bytes)
	var jumpsStrs = strings.Split(fileString, ",")

	var x = 0
	var y = 0
	var z = 0

	for _, jump := range jumpsStrs {

		if jump == "n" {
			y++
			z--
		} else if jump == "ne" {
			x++
			z--
		} else if jump == "nw" {
			x--
			y++
		} else if jump == "se" {
			x++
			y--
		} else if jump == "s" {
			y--
			z++
		} else if jump == "sw" {
			x--
			z++

		}
	}
	fmt.Println(cubeDistance(x, y, z))

}

func cubeDistance(x, y, z int) float64 {
	return (math.Abs(float64(x)) + math.Abs(float64(y)) + math.Abs(float64(z))) / 2
}
