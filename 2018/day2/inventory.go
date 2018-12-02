package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// CountChars returns a map with every char in the
// given string and how many times it appears
func CountChars(line string) map[string]int {
	res := make(map[string]int)
	for _, char := range line {
		res[string(char)]++
	}
	return res
}

func main() {
	filename := "./sample.in"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
