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

	var time int
	var distance int
	firstLine := true
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		if firstLine {
			time = getNumberFromString(split)
		} else {
			distance = getNumberFromString(split)
		}
		firstLine = false
	}
	race := race{
		time:          time,
		currentRecord: distance,
	}

	total := 1

	totalWaysToSurpassRecord := 0
	for heldFor := 0; heldFor < race.time; heldFor++ {
		if heldFor*(race.time-heldFor) > race.currentRecord {
			totalWaysToSurpassRecord++
		}
	}
	total *= totalWaysToSurpassRecord

	fmt.Println(total)
}

func getNumberFromString(entry []string) int {
	stringResult := ""
	for _, stringValue := range entry {
		if _, err := strconv.Atoi(stringValue); err == nil {
			stringResult += stringValue
		}
	}
	result, err := strconv.Atoi(stringResult)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
