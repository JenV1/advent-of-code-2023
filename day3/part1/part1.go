package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

	total := 0

	currentStringNumber := ""
	for lineIndex, line := range lines {
	OUTER:
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

				total += number

				if startingIndex > 0 {
					if notNumberOrPeriod(line[startingIndex-1]) {
						continue
					}
				}
				if endIndex < len(line)-1 {
					if notNumberOrPeriod(line[endIndex+1]) {
						continue
					}
				}
				if lineIndex > 0 {
					for directlyAbove := startingIndex; directlyAbove <= endIndex; directlyAbove++ {
						if notNumberOrPeriod(lines[lineIndex-1][directlyAbove]) {
							continue OUTER

						}
					}
					if startingIndex > 0 {
						if notNumberOrPeriod(lines[lineIndex-1][startingIndex-1]) {
							continue
						}
					}
					if endIndex < len(line)-1 {
						if notNumberOrPeriod(lines[lineIndex-1][endIndex+1]) {
							continue
						}
					}
				}
				if lineIndex < len(lines)-1 {
					for directlyBelow := startingIndex; directlyBelow <= endIndex; directlyBelow++ {
						if notNumberOrPeriod(lines[lineIndex+1][directlyBelow]) {
							continue OUTER
						}
					}
					if startingIndex > 0 {
						if notNumberOrPeriod(lines[lineIndex+1][startingIndex-1]) {
							continue
						}
					}
					if endIndex < len(line)-1 {
						if notNumberOrPeriod(lines[lineIndex+1][endIndex+1]) {
							continue
						}
					}
				}
				total -= number
			}
		}
	}
	fmt.Println(total)
}

func notNumberOrPeriod(entry string) bool {
	if _, err := strconv.Atoi(entry); err == nil {
		return false
	}
	if entry == "." {
		return false
	}
	return true
}
