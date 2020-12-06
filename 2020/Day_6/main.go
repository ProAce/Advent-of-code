package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type seat struct {
	row, column, id int
}

func main() {
	start := time.Now()

	input := getInput("input.txt")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("Part 1: %d\r\n", part1)
	fmt.Printf("Part 2: %d\r\n", part2)

	fmt.Println(time.Since(start))
}

func part1(input [][]string) int {
	count := 0

	for _, group := range input { // Range over all the given inputs.
		mappedAnswers := make(map[rune]int)
		for _, answers := range group { // Range over all the answers in a group.
			for _, char := range answers { // Range over all the characters in an answer.
				if _, ok := mappedAnswers[char]; !ok {
					mappedAnswers[char] = 1 // If the answer hasn't been found yet add it to the map.
					count++                 // And count up.
				}
			}
		}
	}

	return count
}

func part2(input [][]string) int {
	count := 0

	for _, group := range input { // Range over all the given inputs.
		mappedAnswers := make(map[rune]int)
		for _, answers := range group { // Range over all the answers in a group.
			for _, char := range answers { // Range over all the characters in an answer.
				mappedAnswers[char]++ // If the answer hasn't been found yet add it to the map.
			}
		}

		expectedAnswers := len(group)

		for _, amount := range mappedAnswers {
			if amount == expectedAnswers {
				count++
			}
		}
	}

	return count
}

func getInput(path string) [][]string {
	inputFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var output [][]string
	var answers []string

	// Get the input values from the text file and store them in a slice.
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			output = append(output, answers)
			answers = []string{}
			continue
		}

		answers = append(answers, line)
	}

	// Last line is not a new line, don't forget to add the last set
	output = append(output, answers)

	return output
}
