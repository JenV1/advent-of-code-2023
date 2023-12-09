package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type sequence struct {
	previousSequence *sequence
	dataset          []int
	nextSequence     *sequence
}

func (s *sequence) findNextSequence() {
	var dataset []int
	for index := 0; index < len(s.dataset)-1; index++ {
		dataset = append(dataset, s.dataset[index+1]-s.dataset[index])
	}
	s.nextSequence = &sequence{
		previousSequence: s,
		dataset:          dataset,
	}
	for _, val := range s.nextSequence.dataset {
		if val != 0 {
			s.nextSequence.findNextSequence()
			break
		}
	}
}

func (s *sequence) getLowestSequence() *sequence {
	allZero := true
	for _, val := range s.dataset {
		if val != 0 {
			allZero = false
		}
	}
	if allZero {
		return s
	} else {
		return s.nextSequence.getLowestSequence()
	}
}

func (s *sequence) value(val int) int {
	val = s.dataset[0] - val
	if s.previousSequence == nil {
		return val
	} else {
		return s.previousSequence.value(val)
	}
}

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var histories []sequence

	for scanner.Scan() {
		var history []int
		split := strings.Split(scanner.Text(), " ")
		for _, stringNum := range split {
			if num, err := strconv.Atoi(stringNum); err == nil {
				history = append(history, num)
			}
		}
		histories = append(histories, sequence{
			dataset: history,
		})
	}

	total := 0

	for _, history := range histories {
		history.findNextSequence()
		lowestSeq := history.getLowestSequence()
		total += lowestSeq.value(0)
	}
	fmt.Println(total)
}
