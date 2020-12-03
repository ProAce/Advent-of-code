package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func part1(input [][]bool) int {
	return calculatingHitTrees(input, 3, 1)
}

func part2(input [][]bool) int {
	output := calculatingHitTrees(input, 1, 1)
	output *= calculatingHitTrees(input, 3, 1)
	output *= calculatingHitTrees(input, 5, 1)
	output *= calculatingHitTrees(input, 7, 1)
	output *= calculatingHitTrees(input, 1, 2)

	return output
}

func calculatingHitTrees(input [][]bool, right int, down int) int {
	indexY, indexX, treesHit := 0, 0, 0

	for true {
		if indexX >= len(input[0]) {
			indexX -= len(input[0])
		}

		if indexY >= len(input) {
			break
		}

		if input[indexY][indexX] == true {
			treesHit++
		}

		indexX += right
		indexY += down
	}

	return treesHit
}

func getInput(path string) [][]bool {
	inputFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var input [][]bool

	// Get the input values from the text file and store them in a slice.
	for scanner.Scan() {
		line := scanner.Text()

		var tempLine []bool

		for _, value := range line {
			temp := false
			if value == '#' {
				temp = true
			}
			tempLine = append(tempLine, temp)
		}

		input = append(input, tempLine)
	}

	return input
}
