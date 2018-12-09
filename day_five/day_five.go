package main

import (
	"fmt"
	"io/ioutil"
	"unicode"
)

func removeOneElement(polymer string, r rune) string {
	runeLowerCase := unicode.ToLower(r)
	runeUpperCase := unicode.ToUpper(r)

	for i := 0; i < len(polymer); i++ {
		if rune(polymer[i]) == runeLowerCase || rune(polymer[i]) == runeUpperCase {
			polymer = polymer[:i] + polymer[i+1:]
			i--
		}
	}

	return polymer
}

func fullyReactPolymer(polymer string) int {
	for {
		madeAMove := false

		for i := 0; i < len(polymer) - 1; i++ {
			if unicode.ToLower(rune(polymer[i])) == unicode.ToLower(rune(polymer[i+1])) {
				if polymer[i] != polymer[i+1] {
					polymer = polymer[:i] + polymer[i+2:]
					madeAMove = true
					break
				}
			}
		}

		if !madeAMove {
			break
		}
	}

	return len(polymer)
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	part1Result := fullyReactPolymer(string(content))
	fmt.Printf("Part 1 Result: %v\n", part1Result)

	bestSeenLength := -1
	for l := 'a'; l <= 'z'; l++ {
		result := removeOneElement(string(content), l)
		length := fullyReactPolymer(result)
		if bestSeenLength == -1 || length < bestSeenLength {
			bestSeenLength = length
		}
	}

	fmt.Printf("Part 2 Result: %v\n", bestSeenLength)
}