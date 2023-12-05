package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type sourceToDestination struct {
	lowestSource  int
	highestSource int
	difference    int
}

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var seeds []int
	var mapsInOrder [][]sourceToDestination

	firstLine := true
	var mapToDestination []sourceToDestination
	for scanner.Scan() {
		if firstLine {
			fullSeedsString := strings.Split(scanner.Text(), ": ")
			seedsStrings := strings.Split(fullSeedsString[1], " ")
			for _, seedString := range seedsStrings {
				if seed, err := strconv.Atoi(seedString); err == nil {
					seeds = append(seeds, seed)
				}
			}
			firstLine = false
			continue
		}
		line := scanner.Text()
		if line == "" {
			if len(mapToDestination) > 0 {
				mapsInOrder = append(mapsInOrder, mapToDestination)
			}
			mapToDestination = []sourceToDestination{}
		}
		stringNumbers := strings.Split(line, " ")

		var numbers []int

		for _, stringNumber := range stringNumbers {
			if number, err := strconv.Atoi(stringNumber); err == nil {
				numbers = append(numbers, number)
			}
		}
		if len(numbers) == 3 {
			mapToDestination = append(mapToDestination, sourceToDestination{
				lowestSource:  numbers[1],
				highestSource: numbers[1] + numbers[2] - 1,
				difference:    numbers[0] - numbers[1],
			})
		}
	}

	var lowestLocation int

	firstValue := true
	for i := 0; i < len(seeds); i += 2 {
		for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
			currentValue := j
			for _, maps := range mapsInOrder {
				for _, mapToDestination := range maps {
					if currentValue >= mapToDestination.lowestSource &&
						currentValue <= mapToDestination.highestSource {
						currentValue += mapToDestination.difference
						break
					}
				}
			}
			if firstValue {
				lowestLocation = currentValue
				firstValue = false
			}
			if lowestLocation > currentValue {
				lowestLocation = currentValue
			}
		}
	}
	fmt.Println(lowestLocation)
}
