package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

type seat struct {
	row, column, id int
}

func main() {
	start := time.Now()

	input := getInput("input.txt")

	part1, part2 := getMySeat(input)

	fmt.Printf("Part 1: %d\r\n", part1)
	fmt.Printf("Part 2: %d\r\n", part2)

	fmt.Println(time.Since(start))
}

func getMySeat(input []int) (int, int) {
	maxID, myID := 0, 0

	// Sort the input from big to small
	sort.Slice(input, func(i, j int) bool {
		return input[i] > input[j]
	})

	maxID = input[0]

	for i := 1; i < len(input); i++ {
		if input[i] != (input[i-1] - 1) {
			myID = input[i-1] - 1
		}
	}

	return maxID, myID
}

func getInput(path string) []int {
	inputFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var output []int

	// Get the input values from the text file and store them in a slice.
	for scanner.Scan() {
		line := scanner.Text()

		output = append(output, parseSeat(line))
	}

	return output
}

func parseSeat(input string) int {
	minRow, maxRow := 0, 127
	minColumn, maxColumn := 0, 7
	id := 0

	// Parse the row values
	for i := 0; i < 7; i++ {
		if input[i] == 'F' {
			maxRow -= (64 >> i)
		} else if input[i] == 'B' {
			minRow += (64 >> i)
		} else {
			log.Fatalf("Unexpected input: &s, at position: %d, in string: %s", input[i], i, input)
		}
	}

	for i := 0; i < 3; i++ {
		if input[i+7] == 'L' {
			maxColumn -= (4 >> i)
		} else if input[i+7] == 'R' {
			minColumn += (4 >> i)
		} else {
			log.Fatalf("Unexpected input: &s, at position: %d, in string: %s", input[i], i, input)
		}
	}

	id = minRow*8 + minColumn

	return id
}
