package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var stringDigits = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

type stringDigitInString struct {
	startIndexInString int
	number             int
}

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
		split := strings.Split(line, "")

		lastStringDigitInString := stringDigitInString{
			startIndexInString: -1,
			number:             0,
		}
		firstStringDigitInString := stringDigitInString{
			startIndexInString: len(split),
			number:             0,
		}

		for index, stringDigit := range stringDigits {
			firstSubStrIndex := strings.Index(line, stringDigit)
			lastSubStrIndex := strings.LastIndex(line, stringDigit)
			if firstSubStrIndex == -1 {
				continue
			}
			if lastSubStrIndex > lastStringDigitInString.startIndexInString {
				lastStringDigitInString = stringDigitInString{
					startIndexInString: lastSubStrIndex,
					number:             index + 1,
				}
			}
			if firstSubStrIndex < firstStringDigitInString.startIndexInString {
				firstStringDigitInString = stringDigitInString{
					startIndexInString: firstSubStrIndex,
					number:             index + 1,
				}
			}
		}

		for index, element := range split {
			if digit, err := strconv.Atoi(element); err == nil {
				if index < firstStringDigitInString.startIndexInString {
					firstStringDigitInString = stringDigitInString{
						startIndexInString: index,
						number:             digit,
					}
				}
				break
			}
		}
		for index := len(split) - 1; index >= 0; index-- {
			if digit, err := strconv.Atoi(split[index]); err == nil {
				if index > lastStringDigitInString.startIndexInString {
					lastStringDigitInString = stringDigitInString{
						startIndexInString: index,
						number:             digit,
					}
				}
				break
			}
		}
		total += firstStringDigitInString.number*10 + lastStringDigitInString.number
	}
	fmt.Println(total)
}
