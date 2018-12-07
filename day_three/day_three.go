package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func readFile(filePath string) []string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

func main() {
	inputValues := readFile("input.txt")

	takenValues := make(map[string][]int)
	seenClaim := make([]bool, len(inputValues) + 1)

	takenFabricCount := 0

	for _, value := range inputValues {
		var claim, x, y, width, height int
		fmt.Sscanf(value, "#%d @ %d,%d: %dx%d", &claim, &x, &y, &width, &height)
		for i := x; i < x + width; i++ {
			for j := y; j < y + height; j++ {
				key := strconv.Itoa(i) + "," + strconv.Itoa(j)
				if val, ok := takenValues[key]; ok {
					if len(val) == 1 {
						takenFabricCount++
						seenClaim[val[0]] = true
					}
					takenValues[key] = append(takenValues[key], claim)
					seenClaim[claim] = true
				} else {
					takenValues[key] = []int{claim}
				}
			}
		}
	}

	fmt.Printf("Part 1 Result: %d\n", takenFabricCount)

	for i := 1; i <= len(inputValues); i++ {
		if !seenClaim[i] {
			fmt.Printf("Part 2 Result: %v\n", i)
			break
		}
	}
}

