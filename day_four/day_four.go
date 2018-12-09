package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"sort"
)

func readFile(filePath string) []string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

type DataPoint struct {
	Time 		time.Time
	Action 		string
	GuardNum 	int
}

type DataPoints []DataPoint

func (d DataPoints) Len() int {
	return len(d)
}

func (d DataPoints) Less(i, j int) bool {
	return d[i].Time.Before(d[j].Time)
}

func (d DataPoints) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

type Shift struct {
	MinutesAsleep      [60]int
	TotalMinutesAsleep int
}

func main() {
	values := readFile("input.txt")

	var dataPoints DataPoints

	for _, value := range values {
		layout := "2006-01-02 15:04"
		var dateStr, timeStr, action string
		var guardNum int
		fmt.Sscanf(value, "[%s %s %s #%d", &dateStr, &timeStr, &action, &guardNum)
		timeStr = timeStr[0:len(timeStr)-1] // remove trailing bracket

		t, err := time.Parse(layout, dateStr + " " + timeStr)
		if err != nil {
			panic(err)
		}

		dataPoints = append(dataPoints, DataPoint{t, action, guardNum})
	}

	sort.Sort(dataPoints)

	guardMapping := make(map[int]Shift)

	var currentGuardNum, fellAsleepAt int

	highestSeenTotal := -1
	highestSeenTotalCorrespondingGuardNum := -1

	mostFrequentlyAsleepMinute := -1
	mostFrequentlyAsleepMinuteCount := -1
	mostFrequentlyAsleepMinuteCorrespondingGuardNum := -1

	for _, dp := range dataPoints {
		if dp.GuardNum != 0 {
			currentGuardNum = dp.GuardNum
			if _, ok := guardMapping[currentGuardNum]; !ok {
				guardMapping[currentGuardNum] = Shift{}
			}
		} else if dp.Action == "falls" {
			fellAsleepAt = dp.Time.Minute()
		} else { // wakes up
			currentGuardMapping := guardMapping[currentGuardNum]
			currentGuardMapping.TotalMinutesAsleep += dp.Time.Minute() - fellAsleepAt

			if currentGuardMapping.TotalMinutesAsleep > highestSeenTotal {
				highestSeenTotal = currentGuardMapping.TotalMinutesAsleep
				highestSeenTotalCorrespondingGuardNum = currentGuardNum
			}

			for i := fellAsleepAt; i < dp.Time.Minute(); i++ {
				currentGuardMapping.MinutesAsleep[i]++
				if currentGuardMapping.MinutesAsleep[i] > mostFrequentlyAsleepMinuteCount {
					mostFrequentlyAsleepMinuteCount = currentGuardMapping.MinutesAsleep[i]
					mostFrequentlyAsleepMinute = i
					mostFrequentlyAsleepMinuteCorrespondingGuardNum = currentGuardNum
				}
			}

			guardMapping[currentGuardNum] = currentGuardMapping
		}
	}

	highestSeenMinuteTotal := -1
	highestSeenMinute := -1

	for index, val := range guardMapping[highestSeenTotalCorrespondingGuardNum].MinutesAsleep {
		if val > highestSeenMinuteTotal {
			highestSeenMinuteTotal = val
			highestSeenMinute = index
		}
	}

	fmt.Printf("Part 1 Result: %v\n", (highestSeenTotalCorrespondingGuardNum*highestSeenMinute))
	fmt.Printf("Part 2 Result: %v\n", (mostFrequentlyAsleepMinute*mostFrequentlyAsleepMinuteCorrespondingGuardNum))
}
