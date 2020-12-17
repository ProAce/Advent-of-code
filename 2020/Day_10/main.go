package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	input := getInput("input.txt")

	input = append(input, 0) // Add outlet

	// Sort the input once.
	sort.Slice(input, func(a, b int) bool {
		return input[a] < input[b]
	})

	input = append(input, input[len(input)-1]+3) // Add device

	difference := make([]int, 0)

	for i := 0; i < len(input)-1; i++ {
		difference = append(difference, input[i+1]-input[i])
	}

	part1 := part1(difference)
	part2 := part2(difference)

	fmt.Printf("Part 1: %d\r\n", part1)
	fmt.Printf("Part 2: %d\r\n", part2)

	fmt.Println(time.Since(start))
}

func part1(input []int) int {
	diff1, diff3 := 0, 0

	for _, value := range input {
		if value == 1 {
			diff1++
		} else if value == 3 {
			diff3++
		}
	}

	return diff1 * diff3
}

func part2(input []int) int {
	groups := make(map[int]float64)

	for i := 0; i < len(input); i++ {
		count := recursiveSearch(input, i, 0)

		if count != 0 {
			i += count - 1
		}

		groups[count]++
	}

	result := math.Pow(1, groups[1]) * math.Pow(2, groups[2]) * math.Pow(4, groups[3]) * math.Pow(7, groups[4])
	return int(result)
}

func recursiveSearch(input []int, index int, count int) int {
	if input[index] == 1 {
		count = recursiveSearch(input, index+1, count+1)
	}

	return count
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
