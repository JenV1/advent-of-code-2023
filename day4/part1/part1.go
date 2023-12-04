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

	total := 0

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

		cardTotal := 0
		for _, ourCard := range ourCards {
			for _, winningCard := range winningCards {
				if ourCard == winningCard {
					if cardTotal == 0 {
						cardTotal++
					} else {
						cardTotal *= 2
					}
					break
				}
			}
		}
		total += cardTotal
	}
	fmt.Println(total)
}
