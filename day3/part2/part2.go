package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type potentialGear struct {
	pos             [2]int
	adjacentNumbers []int
}

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var lines [][]string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, strings.Split(line, ""))
	}

	var potentialGears []potentialGear

	for lineIndex, line := range lines {
		for elementIndex, element := range line {
			if element == "*" {
				potentialGears = append(potentialGears, potentialGear{
					pos: [2]int{lineIndex, elementIndex},
				})
			}
		}
	}

	currentStringNumber := ""
	for lineIndex, line := range lines {
		for elementIndex, element := range line {
			if _, err := strconv.Atoi(element); err == nil && elementIndex != len(line)-1 {
				currentStringNumber += element
			} else {
				startingIndex := elementIndex - len(currentStringNumber)
				endIndex := elementIndex - 1

				if _, err := strconv.Atoi(element); err == nil && elementIndex == len(line)-1 {
					currentStringNumber += element
					endIndex = len(line) - 1
					startingIndex = len(line) - len(currentStringNumber)
				}
				if len(currentStringNumber) == 0 {
					continue
				}
				number, err := strconv.Atoi(currentStringNumber)
				if err != nil {
					log.Fatal(err)
				}

				currentStringNumber = ""

				if startingIndex > 0 {
					if isAsterisk(line[startingIndex-1]) {
						addAdjacentNumberToAsterisk(potentialGears, [2]int{lineIndex, startingIndex - 1}, number)
					}
				}
				if endIndex < len(line)-1 {
					if isAsterisk(line[endIndex+1]) {
						addAdjacentNumberToAsterisk(potentialGears, [2]int{lineIndex, endIndex + 1}, number)
					}
				}
				if lineIndex > 0 {
					for directlyAbove := startingIndex; directlyAbove <= endIndex; directlyAbove++ {
						if isAsterisk(lines[lineIndex-1][directlyAbove]) {
							addAdjacentNumberToAsterisk(potentialGears, [2]int{lineIndex - 1, directlyAbove}, number)
						}
					}
					if startingIndex > 0 {
						if isAsterisk(lines[lineIndex-1][startingIndex-1]) {
							addAdjacentNumberToAsterisk(potentialGears, [2]int{lineIndex - 1, startingIndex - 1}, number)
						}
					}
					if endIndex < len(line)-1 {
						if isAsterisk(lines[lineIndex-1][endIndex+1]) {
							addAdjacentNumberToAsterisk(potentialGears, [2]int{lineIndex - 1, endIndex + 1}, number)
						}
					}
				}
				if lineIndex < len(lines)-1 {
					for directlyBelow := startingIndex; directlyBelow <= endIndex; directlyBelow++ {
						if isAsterisk(lines[lineIndex+1][directlyBelow]) {
							addAdjacentNumberToAsterisk(potentialGears, [2]int{lineIndex + 1, directlyBelow}, number)
						}
					}
					if startingIndex > 0 {
						if isAsterisk(lines[lineIndex+1][startingIndex-1]) {
							addAdjacentNumberToAsterisk(potentialGears, [2]int{lineIndex + 1, startingIndex - 1}, number)
						}
					}
					if endIndex < len(line)-1 {
						if isAsterisk(lines[lineIndex+1][endIndex+1]) {
							addAdjacentNumberToAsterisk(potentialGears, [2]int{lineIndex + 1, endIndex + 1}, number)
						}
					}
				}
			}
		}
	}
	total := 0
	for _, potentialGear := range potentialGears {
		if len(potentialGear.adjacentNumbers) == 2 {
			total += potentialGear.adjacentNumbers[0] * potentialGear.adjacentNumbers[1]
		}
	}
	fmt.Println(total)
}

func isAsterisk(entry string) bool {
	return entry == "*"
}

func addAdjacentNumberToAsterisk(potentialGears []potentialGear, pos [2]int, number int) {
	for index, potentialGear := range potentialGears {
		if potentialGear.pos == pos {
			potentialGears[index].adjacentNumbers = append(potentialGears[index].adjacentNumbers, number)
		}
	}
}
