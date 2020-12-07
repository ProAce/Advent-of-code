package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type bag struct {
	color    string
	capacity map[string]int
}

type content struct {
	color  string
	amount int
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

func part1(input map[string][]content) int {
	output := make(map[string]int)

	secondaryInput := whichBagCanContain(input, "shiny gold")

	for _, value := range secondaryInput {
		output[value]++
	}

	for true {
		temp := make(map[string]int)
		for _, value := range secondaryInput {
			bags := whichBagCanContain(input, value)

			for _, color := range bags {
				temp[color]++
				output[color]++
			}
		}

		secondaryInput = []string{}
		for key := range temp {
			secondaryInput = append(secondaryInput, key)
		}

		if len(temp) == 0 {
			break
		}
	}

	return len(output)
}

func part2(input map[string][]content) int {
	return countContents(input, "shiny gold")
}

func whichBagCanContain(input map[string][]content, color string) (output []string) {
	for key, value := range input {
		for _, content := range value {
			if color == content.color {
				output = append(output, key)
			}
		}
	}

	return output
}

func countContents(input map[string][]content, color string) int {
	count := 0

	for _, value := range input[color] {
		count += value.amount + (value.amount * countContents(input, value.color))
	}

	return count
}

func getInput(path string) map[string][]content {
	inputFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	output := make(map[string][]content)

	// Get the input values from the text file and store them in a slice.
	for scanner.Scan() {
		line := scanner.Text()

		partRules := strings.Split(line, " contain ")
		color := strings.TrimSuffix(partRules[0], " bags")

		var capacity []content

		if partRules[1] != "no other bags." {
			contents := strings.Split(partRules[1], ", ")

			for _, value := range contents {
				words := strings.SplitN(value, " ", 2)
				numberOfBags, _ := strconv.Atoi(words[0])
				temp := strings.Split(words[1], " ")
				bagColor := temp[0] + " " + temp[1]
				capacity = append(capacity, content{bagColor, numberOfBags})
			}
		}

		output[color] = capacity
	}

	return output
}
