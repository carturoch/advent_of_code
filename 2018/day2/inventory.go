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

// GetCharsMoreFreqThan returns terms with frequency higher than val
func (table FreqTable) GetCharsMoreFreqThan(val int) (res []string) {
	for char, freq := range table {
		if freq >= val {
			res = append(res, char)
		}
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
	duosCount := 0
	triosCount := 0
	for scanner.Scan() {
		ft := CountChars(scanner.Text())
		trios := ft.GetCharsMoreFreqThan(3)
		duos := ft.GetCharsMoreFreqThan(2)
		if len(trios) > 0 {
			triosCount++
		}
		if len(duos) > 0 {
			duosCount++
		}
	}
	fmt.Println("Trios:", triosCount, "Duos:", duosCount)
}
