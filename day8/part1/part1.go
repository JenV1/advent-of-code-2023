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
	var currentLocation string

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
	}
	steps := 0
	currentLocation = "AAA"
	foundZZZ := false
	for !foundZZZ {
		for _, leftOrRight := range leftRightPattern {
			if leftOrRight == "L" {
				currentLocation = leftRightLocations[currentLocation][0]
			} else {
				currentLocation = leftRightLocations[currentLocation][1]
			}
			steps++
			if currentLocation == "ZZZ" {
				foundZZZ = true
			}
		}
	}
	fmt.Println(steps)
}
