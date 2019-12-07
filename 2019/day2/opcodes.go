package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const path = "./input.txt"
const step = 4

const (
	sum  int = 1
	mul  int = 2
	halt int = 99
)

func readInput() []int {
	data, _ := ioutil.ReadFile(path)
	instructions := strings.Split(string(data), ",")
	program := make([]int, len(instructions))
	for i, instruction := range instructions {
		program[i], _ = strconv.Atoi(instruction)
	}
	return program
}

// Exec process each if the instructions in the program
func Exec(instructions []int) []int {
	size := len(instructions)
	for index := 0; index < size-step; index += step {
		code := instructions[index]
		ref1 := instructions[index+1]
		ref2 := instructions[index+2]
		ref3 := instructions[index+3]
		val1 := instructions[ref1]
		val2 := instructions[ref2]

		switch code {
		case sum:
			instructions[ref3] = val1 + val2
		case mul:
			instructions[ref3] = val1 * val2
		case halt:
			return instructions
		}
	}
	return instructions
}

func main() {
	fmt.Println(Exec(readInput()))
}
