package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	input := getInput("input.txt")

	part1 := part1(input)
	// part2 := part2(input, part1, index)

	fmt.Printf("Part 1: %d\r\n", part1)
	// fmt.Printf("Part 2: %d\r\n", part2)

	fmt.Println(time.Since(start))
}

func part1(input []int) int {
	sort.Slice(input, func(a, b int) bool {
		return input[a] < input[b]
	})

	diff1 := 0
	diff3 := 1 // Your device has a 3 jolt difference to the last adapter

	if input[0] == 1 {
		diff1++
	}

	for i := 0; i < len(input)-1; i++ {
		switch input[i+1] - input[i] {
		case 1:
			diff1++
			break
		case 3:
			diff3++
			break
		}
	}

	return diff1 * diff3
}

func part2(input []int) int {

	return 0
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
