package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

type opcode struct {
	operation string
	value     int
}

func main() {
	start := time.Now()

	input := getInput("input.txt")

	part1, index := part1(input)
	part2 := part2(input, part1, index)

	fmt.Printf("Part 1: %d\r\n", part1)
	fmt.Printf("Part 2: %d\r\n", part2)

	fmt.Println(time.Since(start))
}

func part1(input []int) (int, int) {
	preAmble := 25

	// Start looping over the input after the first preamble.
	for i := preAmble; i < len(input); i++ {
		valid := false

		// Loop over the preamble to see if one of those pairs matches up with the input.
		for j := i - preAmble; j < i; j++ {
			for k := j + 1; k < i; k++ {

				// If it does, note it and break out of the loop.
				if input[i] == (input[j] + input[k]) {
					valid = true
					break
				}
			}

			if valid {
				break
			}
		}

		if !valid {
			return input[i], i
		}
	}

	return 0, 0
}

func part2(input []int, solutionPart1 int, index int) int {
	// Up until the index, find the first group of inputs that adds up to the first invalid number.
	for i := 0; i < index; i++ {
		sum := 0

		for j := i; j < index; j++ {
			sum += input[j]

			if sum > solutionPart1 {
				break
			} else if sum == solutionPart1 {
				// If we have found the group of inputs, add up the highest and lowest value in that range.
				return findLowest(input[i:j+1]) + findHighest(input[i:j+1])
			}
		}
	}

	return 0
}

func findLowest(input []int) int {
	lowest := math.MaxInt64

	for _, value := range input {
		if value < lowest {
			lowest = value
		}
	}

	return lowest
}

func findHighest(input []int) int {
	highest := 0

	for _, value := range input {
		if value > highest {
			highest = value
		}
	}

	return highest
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

		value, _ := strconv.Atoi(line)

		output = append(output, value)
	}

	return output
}
