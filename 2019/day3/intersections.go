package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Step struct {
	dir rune
	len int
}

func ParseWireIntoSteps(line string) (steps []Step) {
	parts := strings.Split(line, " ")
	for _, part := range parts {
		dir := rune(part[0])
		len, _ := strconv.Atoi(part[1:])
		step := Step{dir, len}
		steps = append(steps, step)
	}
	return
}

func main() {
	fmt.Println("Day 3")
}
