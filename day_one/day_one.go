package main

import (
	"fmt"
	"io"
	"os"
)

func readFile(filePath string) (frequencies []int, totalFrequency int) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	totalFrequency = 0

	var currentFrequency int
	for {
		_, err := fmt.Fscanf(file, "%d\n", &currentFrequency)

		if err != nil {
			if err == io.EOF {
				return
			}
			panic(err)
		}

		totalFrequency += currentFrequency

		frequencies = append(frequencies, currentFrequency)
	}

	return
}

func getFirstFrequencyRetrievedTwice(frequencies []int) int {
	if len(frequencies) == 0 {
		panic("No retrieved frequencies")
	}

	seenFrequencies := make(map[int]bool)
	currentFrequency := 0
	index := 0

	for {
		if seenFrequencies[currentFrequency] {
			break
		}
		seenFrequencies[currentFrequency] = true

		currentFrequency += frequencies[index]
		if index++; index == len(frequencies) {
			index = 0
		}
	}

	return currentFrequency
}

func main() {
	frequencies, totalFrequency := readFile("input.txt")
	fmt.Printf("\nTotal Frequency: %d", totalFrequency)

	currentFrequency := getFirstFrequencyRetrievedTwice(frequencies)
	fmt.Printf("\nFirst frequency reached twice: %d", currentFrequency)
}
