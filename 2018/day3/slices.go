package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	s "strings"
)

// Claim allocates lines in a structured way
// #123 @ 3,2: 5x4
type Claim struct {
	ID                       string
	Left, Top, Width, Height int
}

// ParseClaim parses a file line into a claim
func ParseClaim(line string) (Claim, error) {
	claim := Claim{}
	parts := s.Split(line, " ")
	var err error
	if len(parts) != 4 {
		return claim, errors.New("Unexpected format")
	}
	claim.ID = parts[0]
	coords := s.Split(s.TrimSuffix(parts[2], ":"), ",")
	claim.Left, err = strconv.Atoi(coords[0])
	claim.Top, err = strconv.Atoi(coords[1])

	dim := s.Split(parts[3], "x")
	claim.Width, err = strconv.Atoi(dim[0])
	claim.Height, err = strconv.Atoi(dim[1])

	return claim, err
}

func scanFile(filename string) *bufio.Scanner {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewScanner(file)
}

func eachLine(filename string, f func(string)) {
	scanner := scanFile(filename)
	for scanner.Scan() {
		newLine := scanner.Text()
		f(newLine)
	}
}

func main() {
	filename := "./input.in"
	eachLine(filename, func(line string) {
		fmt.Println(ParseClaim(line))
	})
}
