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

	inputFile, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var input []inputStruct

	// Get the input values from the text file and store them in a slice.
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

	var validPasswords int

	for _, value := range input {
		// Check if which of the indexes contain the set rule.
		boolA := value.password[value.indexA-1] == value.rule
		boolB := value.password[value.indexB-1] == value.rule

		// Only one of them may contain the set rule, xor both outcomes.
		if boolA != boolB {
			validPasswords++
		}
	}

	fmt.Println(validPasswords)
	fmt.Println(time.Since(start))
}
