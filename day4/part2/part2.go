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
	noOfLines := numberOfLines()

	var scratchCardsByCopies []int

	for cardNumber := 1; cardNumber <= noOfLines; cardNumber++ {
		scratchCardsByCopies = append(scratchCardsByCopies, 1)
	}

	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		removeIdentifier := strings.Split(line, ": ")
		stringCards := strings.Split(removeIdentifier[1], " | ")
		stringOurCards := strings.Split(stringCards[0], " ")
		stringWinningCards := strings.Split(stringCards[1], " ")

		var ourCards []int
		var winningCards []int
		for _, stringCard := range stringOurCards {
			if number, err := strconv.Atoi(stringCard); err == nil {
				ourCards = append(ourCards, number)
			}
		}
		for _, stringCard := range stringWinningCards {
			if number, err := strconv.Atoi(stringCard); err == nil {
				winningCards = append(winningCards, number)
			}
		}

		matchingNumbers := 0
		for _, ourCard := range ourCards {
			for _, winningCard := range winningCards {
				if ourCard == winningCard {
					matchingNumbers++
					break
				}
			}
		}
		for i := index + 1; i <= index+matchingNumbers; i++ {
			scratchCardsByCopies[i] += scratchCardsByCopies[index]
		}
		index++
	}
	total := 0
	for _, number := range scratchCardsByCopies {
		total += number
	}
	fmt.Println(total)
}

func numberOfLines() int {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	lines := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines++
	}
	return lines
}
