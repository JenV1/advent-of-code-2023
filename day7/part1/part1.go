package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type handsByType [7][]hand

type hand struct {
	cards []string
	bid   int
}

var cardValues = map[string]int{
	"A": 12,
	"K": 11,
	"Q": 10,
	"J": 9,
	"T": 8,
	"9": 7,
	"8": 6,
	"7": 5,
	"6": 4,
	"5": 3,
	"4": 2,
	"3": 1,
	"2": 0,
}

type ByCards []hand

func (h ByCards) Len() int {
	return len(h)
}

func (h ByCards) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h ByCards) Less(i, j int) bool {
	for index, card := range h[i].cards {
		if cardValues[card] == cardValues[h[j].cards[index]] {
			continue
		}
		if cardValues[card] < cardValues[h[j].cards[index]] {
			return true
		} else {
			return false
		}
	}
	return false
}

type handType int

const (
	FiveOfAKind handType = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

func (h hand) Type() handType {
	occuredCards := make(map[string]int)
	for _, card := range h.cards {
		occuredCards[card]++
	}
	if len(occuredCards) == 1 {
		return FiveOfAKind
	}
	if len(occuredCards) == 2 {
		for _, v := range occuredCards {
			if v == 4 {
				return FourOfAKind
			}
		}
		return FullHouse
	}
	if len(occuredCards) == 3 {
		for _, v := range occuredCards {
			if v == 3 {
				return ThreeOfAKind
			}
		}
		return TwoPair
	}
	if len(occuredCards) == 4 {
		return OnePair
	}
	if len(occuredCards) == 5 {
		return HighCard
	}
	panic("hand type not found")
}

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var handsByType handsByType

	for scanner.Scan() {
		lineSplit := strings.Split(scanner.Text(), " ")
		cards := strings.Split(lineSplit[0], "")
		bid, err := strconv.Atoi(lineSplit[1])
		if err != nil {
			log.Fatal("bid isn't a number")
		}
		hand := hand{
			cards: cards,
			bid:   bid,
		}
		handsByType[hand.Type()] = append(handsByType[hand.Type()], hand)
	}

	total := 0
	rank := 1
	for index := len(handsByType) - 1; index >= 0; index-- {
		sort.Sort(ByCards(handsByType[index]))
		for _, hand := range handsByType[index] {
			total += rank * hand.bid
			rank++
		}
	}
	fmt.Println(total)
}
