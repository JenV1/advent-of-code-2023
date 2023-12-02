package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const redLimit = 12
const greenLimit = 13
const blueLimit = 14

type game struct {
	rounds []gameRound
}

type gameRound struct {
	blue  int
	red   int
	green int
}

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var games []game

	for scanner.Scan() {
		line := scanner.Text()
		rounds := strings.Split(line, ";")

		game := game{}
		for _, round := range rounds {
			blue := getColourNumber("blue", round)
			red := getColourNumber("red", round)
			green := getColourNumber("green", round)
			game.rounds = append(game.rounds, gameRound{
				blue:  blue,
				red:   red,
				green: green,
			})
		}
		games = append(games, game)
	}

	idTotal := 0

	for index, game := range games {
		gamePossible := true
		for _, round := range game.rounds {
			if round.blue > blueLimit ||
				round.green > greenLimit ||
				round.red > redLimit {
				gamePossible = false
			}
		}
		if gamePossible {
			idTotal += index + 1
		}
	}
	fmt.Println(idTotal)
}

func getColourNumber(colour string, round string) int {
	var err error
	noColour := 0

	splitByColour := strings.Split(round, fmt.Sprintf(" %s", colour))
	if len(splitByColour) > 1 {
		split := strings.Split(splitByColour[0], " ")
		noColour, err = strconv.Atoi(split[len(split)-1])
		if err != nil {
			log.Fatal(err)
		}
	}
	return noColour
}
