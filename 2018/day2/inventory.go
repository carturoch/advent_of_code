package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

// Common returns the chars that match
// in value and position from the given strings
func Common(a, b string) (res string) {
	max := int(math.Min(float64(len(a)), float64(len(b))))
	for index := 0; index < max; index++ {
		if a[index] == b[index] {
			res += string(a[index])
		}
	}
	return res
}

func step1(scanner *bufio.Scanner) {
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

func step2(scanner *bufio.Scanner) string {
	readLines := []string{}
	for scanner.Scan() {
		newLine := scanner.Text()
		for _, readLine := range readLines {
			if common := Common(newLine, readLine); len(common) == len(newLine)-1 {
				return common
			}
		}
		readLines = append(readLines, newLine)
	}
	return ""
}

func main() {
	filename := "./input.in"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	res := step2(scanner)
	fmt.Println(res)
}
