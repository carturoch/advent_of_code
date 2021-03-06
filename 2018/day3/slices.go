package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	s "strings"
)

// State represents each one of the possible states of each block
type State int

const (
	Empty State = iota
	Marked
	Conflicted
)

// Layout represents a square matrix
type Layout [][]State

// Claim allocates lines in a structured way
// #123 @ 3,2: 5x4
type Claim struct {
	ID                       string
	Left, Top, Width, Height int
}

// InitLayout initializes a layout with empty state
func InitLayout(size int) Layout {
	layout := make(Layout, size)
	for i := 0; i < size; i++ {
		layout[i] = make([]State, size)
		for j := 0; j < size; j++ {
			layout[i][j] = Empty
		}
	}
	return layout
}

// ApplyClaim applies a given claim to a layout
func ApplyClaim(c Claim, l *Layout) *Layout {
	minTop, minLeft := boundaries(c, l)
	for i := c.Top; i < minTop; i++ {
		for j := c.Left; j < minLeft; j++ {
			state := Marked
			if st := (*l)[i][j]; st == Marked || st == Conflicted {
				state = Conflicted
			}
			(*l)[i][j] = state
		}
	}
	return l
}

// HasConflicts determines whether a claim has conflicts
func HasConflicts(c Claim, l *Layout) bool {
	minTop, minLeft := boundaries(c, l)
	for i := c.Top; i < minTop; i++ {
		for j := c.Left; j < minLeft; j++ {
			if (*l)[i][j] == Conflicted {
				return true
			}
		}
	}
	return false
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

func boundaries(c Claim, l *Layout) (int, int) {
	minTop := int(math.Min(float64(len(*l)), float64(c.Top+c.Height)))
	minLeft := int(math.Min(float64(len(*l)), float64(c.Left+c.Width)))
	return minTop, minLeft
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

func step1(l Layout) int {
	count := 0
	size := len(l)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if l[i][j] == Conflicted {
				count++
			}
		}
	}
	return count
}

func main() {
	filename := "./input.in"
	size := 1000
	layout := InitLayout(size)
	var claims []Claim
	eachLine(filename, func(line string) {
		claim, _ := ParseClaim(line)
		claims = append(claims, claim)
		ApplyClaim(claim, &layout)
	})

	// fmt.Println(step1(layout))

	for _, c := range claims {
		if !HasConflicts(c, &layout) {
			fmt.Println("Found:", c.ID)
		}
	}
	fmt.Println("Done")
}
