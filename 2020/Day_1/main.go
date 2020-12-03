package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	input := getInput("input.txt")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("Part 1: %d\r\n", part1)
	fmt.Printf("Part 2: %d\r\n", part2)
	fmt.Println(time.Since(start))
}

func part1(input []int) int {
	// Loop over the slice to find the values that add up to 2020.
	for a := 0; a < len(input); a++ {
		for b := a + 1; b < len(input); b++ {
			if (input[a] + input[b]) == 2020 {
				return input[a] * input[b]
			}
		}
	}

	return 0
}

func part2(input []int) int {
	// Loop over the slice to find the values that add up to 2020.
	for a := 0; a < len(input); a++ {
		for b := a + 1; b < len(input); b++ {
			for c := b + 1; c < len(input); c++ {
				if (input[a] + input[b] + input[c]) == 2020 {
					return input[a] * input[b] * input[c]
				}
			}
		}
	}

	return 0
}

func getInput(path string) []int {
	inputFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var input []int

	// Get the input values from the text file and store them in a slice.
	for scanner.Scan() {
		line := scanner.Text()

		i, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal(err)
		}

		input = append(input, i)
	}

	return input
}
