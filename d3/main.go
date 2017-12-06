package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {

	var bytes, _ = ioutil.ReadFile("input.txt")

	var valStr = strings.TrimSpace(string(bytes))

	var value, _ = strconv.Atoi(valStr)
	fmt.Println(value)
	var i = 1
	for ; i*i < value; i += 2 {

	}
	i -= 2

	var next = i + 1
	var halfnext = next / 2
	fmt.Println(i)
	fmt.Println(i * i)
	fmt.Println(value - i*i - next - next - next)

	var x = math.Abs(float64(halfnext - (value-i*i)%next))
	var y = halfnext
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(int(x) + y)
	// for _, val := range bytes {

	// }

	// fmt.Println(sum)

}
