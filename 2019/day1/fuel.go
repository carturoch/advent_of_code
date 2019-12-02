package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// FuelPerModule calculates the required fuel
func FuelPerModule(mass int) int {
	divided := mass / 3
	if divided > 2 {
		return divided - 2
	}
	return 0
}

func step1(scanner *bufio.Scanner) int {
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		mass, _ := strconv.Atoi(line)
		total = total + FuelPerModule(mass)
	}
	return total
}

func main() {
	filename := "./input.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	fmt.Println(step1(scanner))
}
