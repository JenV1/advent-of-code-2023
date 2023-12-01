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
		split := strings.Split(line, "")
		for _, element := range split {
			if digit, err := strconv.Atoi(element); err == nil {
				total += digit * 10
				break
			}
		}
		for index := len(split) - 1; index >= 0; index-- {
			if digit, err := strconv.Atoi(split[index]); err == nil {
				total += digit
				break
			}
		}
	}
	fmt.Println(total)
}
