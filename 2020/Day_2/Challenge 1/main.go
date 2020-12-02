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

type password struct {
	min, max       int
	rule, password string
}

func main() {
	start := time.Now()

	inputFile, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var input []password

	// Get the input values from the text file and store them in a slice.
	for scanner.Scan() {
		line := scanner.Text()

		temp := strings.Split(line, " ")

		minMax := strings.Split(temp[0], "-")
		min, err := strconv.Atoi(minMax[0])
		max, err := strconv.Atoi(minMax[1])

		if err != nil {
			log.Fatal(err)
		}

		rule := strings.TrimSuffix(temp[1], ":")

		input = append(input, password{min, max, rule, temp[2]})
	}

	var validPasswords int

	for _, value := range input {
		count := strings.Count(value.password, value.rule)
		if (count >= value.min) && (count <= value.max) {
			validPasswords++
		}
	}

	fmt.Println(validPasswords)
	fmt.Println(time.Since(start))
}
