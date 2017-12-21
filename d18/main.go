package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	var bytes, _ = ioutil.ReadFile("input.txt")
	var fileString = string(bytes)

	var registers = make(map[string]int, 0)
	var steps = strings.Split(fileString, "\n")

	var lastSend = 0
	for idx := 0; idx < len(steps); {
		var values = strings.Split(steps[idx], " ")
		var ops = values[0]
		var register1 = values[1]
		var register2Value = 0

		if len(values) > 2 {
			val, err := strconv.Atoi(values[2])
			if err != nil {
				fmt.Println(values[2])
				register2Value = registers[values[2]]
			} else {
				register2Value = val
			}
		}

		switch ops {
		case "set":
			registers[register1] = register2Value
		case "mul":
			registers[register1] *= register2Value
		case "jgz":
			if registers[register1] > 0 {
				idx += register2Value
				continue
			}
		case "add":
			registers[register1] += register2Value
		case "snd":
			lastSend = registers[register1]
		case "rcv":
			if registers[register1] != 0 {
				fmt.Println(lastSend)
				os.Exit(1)
			}
		case "mod":
			fmt.Println(register2Value)
			fmt.Println(steps[idx])
			registers[register1] %= register2Value
		default:
			fmt.Printf("I FORGOT ABOUT %s\n", ops)
		}
		fmt.Printf("%v\n", registers)
		idx++

	}

}
