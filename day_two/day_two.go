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

func getSimilarCharactersInWordsWithOneDiff(values []string) string {
	a, b := getOneCharDifferentStrings(values)
	return getSimilarCharacters(a, b)
}

func getSimilarCharacters(a string, b string) (result string) {
	for index := 0; index < len(a); index++ {
		if a[index] != b[index] {
			return a[0:index] + b[index+1:]
		}
	}
	return
}

func getOneCharDifferentStrings(values []string) (resultOne string, resultTwo string) {
	possibilities := make(map[string]string)
	for _, val := range values {
		for index := 0; index < len(val); index++ {
			newString := val[0:index] + "_" + val[index+1:]
			if result, ok := possibilities[newString]; ok {
				return result, val
			} else {
				possibilities[newString] = val
			}
		}
	}
	return
}

func main() {
	values := readFile("input.txt")

	checkSum := calculateCheckSum(values)
	fmt.Printf("\nCheck Sum: %d", checkSum)

	similarCharacters := getSimilarCharactersInWordsWithOneDiff(values)
	fmt.Printf("\nSimilar Characters: %v", similarCharacters)
}
