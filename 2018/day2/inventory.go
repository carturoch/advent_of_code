package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// FreqTable is map with a char and its frequency in a given string
type FreqTable map[string]int

// CountChars returns a map with every char in the
// given string and how many times it appears
func CountChars(line string) FreqTable {
	res := make(map[string]int)
	for _, char := range line {
		res[string(char)]++
	}
	return res
}

// GetCharsEqFreq returns terms with frequency equal to val
func (table FreqTable) GetCharsEqFreq(val int) (res []string) {
	for char, freq := range table {
		if freq == val {
			res = append(res, char)
		}
	}
	return res
}

func sliceDiff(a, b []string) (diff []string) {
	ref := make(map[string]bool)
	for _, char := range a {
		ref[char] = true
	}
	for _, char := range b {
		if !ref[char] {
			diff = append(diff, char)
		}
	}
	return diff
}

func main() {
	filename := "./input.in"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	duosCount := 0
	triosCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		ft := CountChars(line)
		trios := ft.GetCharsEqFreq(3)
		duos := ft.GetCharsEqFreq(2)
		if len(trios) > 0 {
			triosCount++
		}
		if len(duos) > 0 && len(sliceDiff(trios, duos)) > 0 {
			duosCount++
		}
	}
	fmt.Println("Trios:", triosCount, "Duos:", duosCount, "Result:", triosCount*duosCount)
}
