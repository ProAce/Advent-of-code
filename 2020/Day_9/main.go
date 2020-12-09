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
	sum := input[0]
	begin, end := 0, 0

	for sum != solutionPart1 {
		if sum < solutionPart1 {
			end++
			sum += input[end]
		} else if sum > solutionPart1 {
			sum -= input[begin]
			begin++
		}
	}

	low, high := findExtremes(input[begin : end+1])
	return low + high
}

func findExtremes(input []int) (int, int) {
	lowest := math.MaxInt64
	highest := 0

	for _, value := range input {
		if value < lowest {
			lowest = value
		}

		if value > highest {
			highest = value
		}
	}

	return lowest, highest
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
