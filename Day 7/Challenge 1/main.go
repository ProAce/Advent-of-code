package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func parseInput(input string) (before, after string) {
	split := strings.Split(input, " ")
	return split[1], split[7]
}

func main() {
	start := time.Now()

	instructions := []string{}

	intputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(intputFile)

	for scanner.Scan() {
		input := scanner.Text()
		before, after := parseInput(input)
		instructions = append(instructions, before+after)
	}

	for i := 0; i < len(instructions); i++ {
		fmt.Println(instructions[i])
	}

	fmt.Println(time.Since(start))
}
