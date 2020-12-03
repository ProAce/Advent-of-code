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

type inputStruct struct {
	indexA, indexB int
	rule           rune
	password       []rune
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

func part1(input []inputStruct) int {
	validPasswords := 0

	for _, value := range input {
		count := 0

		// Count the amount of characters of type rule in the password.
		for _, char := range value.password {
			if char == value.rule {
				count++
			}
		}

		// Check if this amount is allowed given the min and max indexes.
		if (count >= value.indexA) && (count <= value.indexB) {
			validPasswords++
		}
	}

	return validPasswords
}

func part2(input []inputStruct) int {
	validPasswords := 0

	for _, value := range input {
		// Check if which of the indexes contain the set rule.
		boolA := value.password[value.indexA-1] == value.rule
		boolB := value.password[value.indexB-1] == value.rule

		// Only one of them may contain the set rule, xor both outcomes.
		if boolA != boolB {
			validPasswords++
		}
	}

	return validPasswords
}

func getInput(path string) []inputStruct {
	inputFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var input []inputStruct

	for scanner.Scan() {
		line := scanner.Text()

		temp := strings.Split(line, " ")

		// Parse the given input so we can do some processing on it later.
		indexes := strings.Split(temp[0], "-")
		indexA, err := strconv.Atoi(indexes[0])
		indexB, err := strconv.Atoi(indexes[1])

		if err != nil {
			log.Fatal(err)
		}

		rule := []rune(strings.TrimSuffix(temp[1], ":"))
		password := []rune(temp[2])

		input = append(input, inputStruct{indexA, indexB, rule[0], password})
	}

	return input
}
