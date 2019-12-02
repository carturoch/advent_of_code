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

// DeepFuelPerModule calculates the total required fuel
// recursively for all the masses of fuel as well
func DeepFuelPerModule(mass int) int {
	fuel := FuelPerModule(mass)
	if fuel > 0 {
		return fuel + DeepFuelPerModule(fuel)
	}
	return fuel
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

func step2(scanner *bufio.Scanner) int {
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		mass, _ := strconv.Atoi(line)
		total = total + DeepFuelPerModule(mass)
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

	// fmt.Println(step1(scanner))
	fmt.Println(step2(scanner))
}
