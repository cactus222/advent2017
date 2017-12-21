package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var bytes, _ = ioutil.ReadFile("input.txt")
	var fileString = string(bytes)

	var steps = strings.Split(fileString, "\n")
	var chan0 = make(chan int, 20000)
	var chan1 = make(chan int, 20000)
	go runner(steps, chan0, chan1, false)
	go runner(steps, chan1, chan0, true)
	var quit = make(chan int, 0)
	<-quit
}

func runner(steps []string, ownChan chan int, otherChan chan int, log bool) {
	var registers = make(map[string]int, 0)

	var sent = 0
	if log {
		registers["p"] = 1
	}
	for idx := 0; idx < len(steps); {
		var values = strings.Split(steps[idx], " ")
		var ops = values[0]
		var register1 = values[1]
		var register2Value = 0

		if len(values) > 2 {
			val, err := strconv.Atoi(values[2])
			if err != nil {
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

			val, err := strconv.Atoi(register1)
			if err != nil {
				val = registers[register1]
			}
			if val > 0 {
				idx += register2Value
				continue
			}
		case "add":
			registers[register1] += register2Value
		case "snd":
			sent++
			if log {
				fmt.Printf("sent %d, %d, %d\n", log, registers[register1], sent)
			}

			otherChan <- registers[register1]

			//if log {
			//}

		case "rcv":
			registers[register1] = <-ownChan
		case "mod":
			registers[register1] = registers[register1] % register2Value
		default:
			fmt.Printf("I FORGOT ABOUT %s\n", ops)
		}
		//	fmt.Printf("%v\n", registers)
		idx++

	}
}
