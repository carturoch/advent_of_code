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

func outbound(ref, size int) bool {
	return ref < 0 || ref >= size
}

func anyOutbound(references []int, upper int) bool {
	for _, ref := range references {
		if outbound(ref, upper) {
			return true
		}
	}
	return false
}

// Exec process each if the instructions in the program
func Exec(instructions []int) []int {
	size := len(instructions)
	for index := 0; index < size-step; index += step {
		code := instructions[index]
		ref1 := instructions[index+1]
		ref2 := instructions[index+2]
		ref3 := instructions[index+3]

		if anyOutbound([]int{ref1, ref2, ref3}, size) {
			return []int{0, 0, 0, 0}
		}

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

func result(noun, verb int, instructions []int) int {
	instructions[1] = noun
	instructions[2] = verb
	res := Exec(instructions)
	return res[0]
}

func bf(target int, instructions []int) (int, int) {
	max := 99
	var res int
	size := len(instructions)
	for noun := 0; noun <= max; noun++ {
		for verb := 0; verb <= max; verb++ {
			disposableInstructions := make([]int, size)
			copy(disposableInstructions, instructions)
			res = result(noun, verb, disposableInstructions)
			if res == target {
				return noun, verb
			}
		}
	}
	return 0, 0
}

func main() {
	// input := readInput()
	// fmt.Println(result(12, 2, input))

	target := 19690720
	noun, verb := bf(target, readInput())
	fmt.Println(noun, verb)
}
