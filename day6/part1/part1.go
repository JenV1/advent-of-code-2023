package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type race struct {
	time          int
	currentRecord int
}

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var times []int
	var distances []int
	firstLine := true
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		if firstLine {
			times = getNumbersFromString(split)
		} else {
			distances = getNumbersFromString(split)
		}
		firstLine = false
	}
	var races []race

	for index, time := range times {
		races = append(races, race{
			time:          time,
			currentRecord: distances[index],
		})
	}

	total := 1

	for _, race := range races {
		totalWaysToSurpassRecord := 0
		for heldFor := 0; heldFor < race.time; heldFor++ {
			if heldFor*(race.time-heldFor) > race.currentRecord {
				totalWaysToSurpassRecord++
			}
		}
		total *= totalWaysToSurpassRecord
	}
	fmt.Println(total)
}

func getNumbersFromString(entry []string) []int {
	var result []int
	for _, stringValue := range entry {
		if number, err := strconv.Atoi(stringValue); err == nil {
			result = append(result, number)
		}
	}
	return result
}
