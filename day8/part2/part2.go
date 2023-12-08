package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var leftRightPattern []string
	leftRightLocations := make(map[string][2]string)
	var locationsEndingInA []string

	firstLine := true
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		if firstLine {
			leftRightPattern = strings.Split(scanner.Text(), "")
			firstLine = false
			continue
		}
		split := strings.Split(scanner.Text(), " = (")
		locations := strings.Split(split[1], ", ")
		leftRightLocations[split[0]] = [2]string{locations[0], strings.Split(locations[1], ")")[0]}
		if strings.Split(split[0], "")[2] == "A" {
			locationsEndingInA = append(locationsEndingInA, split[0])
		}
	}
	foundAllEndingInZ := false
	currentLocations := locationsEndingInA

	numberOfStepsPerStart := make([]int, len(currentLocations))

	for !foundAllEndingInZ {
	OUTER:
		for _, leftOrRight := range leftRightPattern {
			for index, currentLocation := range currentLocations {
				if strings.Split(currentLocation, "")[2] == "Z" {
					continue
				}
				if leftOrRight == "L" {
					currentLocations[index] = leftRightLocations[currentLocation][0]
				} else {
					currentLocations[index] = leftRightLocations[currentLocation][1]
				}
				numberOfStepsPerStart[index]++
			}
			for _, currentLocation := range currentLocations {
				if strings.Split(currentLocation, "")[2] != "Z" {
					continue OUTER
				}
			}
			foundAllEndingInZ = true
		}
	}
	fmt.Println(numberOfStepsPerStart)
	fmt.Println(LCM(numberOfStepsPerStart[0], numberOfStepsPerStart[1],
		numberOfStepsPerStart[2], numberOfStepsPerStart[3], numberOfStepsPerStart[4],
		numberOfStepsPerStart[5]))
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
