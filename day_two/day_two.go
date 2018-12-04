package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile(filePath string) []string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

func calculateCheckSum(values []string) int {
	seenTwoCount := 0
	seenThreeCount := 0

	for _, val := range values {
		results := make(map[int32]int)
		for _, character := range val {
			results[character]++
		}

		seenTwo := false
		seenThree := false
		for _, result := range results {
			if result == 2 {
				seenTwo = true
			} else if result == 3 {
				seenThree = true
			}
		}

		if seenTwo {
			seenTwoCount++
		}

		if seenThree {
			seenThreeCount++
		}
	}

	return seenTwoCount * seenThreeCount
}

func main() {
	values := readFile("input.txt")

	checkSum := calculateCheckSum(values)
	fmt.Printf("\nCheck Sum: %d", checkSum)
}
