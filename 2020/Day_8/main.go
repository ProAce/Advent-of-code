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

type opcode struct {
	operation string
	value     int
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

func part1(input []opcode) int {
	accumulator, _ := runOpCode(input)

	return accumulator
}

func part2(input []opcode) int {
	changePointer := -1
	changedToNop := false

	// Find the first operation we can change to test if that fixes the bootloop.
	for true {
		for i := changePointer + 1; i < len(input); i++ {
			if input[i].operation == "nop" {
				changedToNop = false
				changePointer = i
				input[i].operation = "jmp"
				break
			}

			if input[i].operation == "jmp" {
				changedToNop = true
				changePointer = i
				input[i].operation = "nop"
				break
			}
		}
		accumulator, endState := runOpCode(input)

		if endState == 0 {
			return accumulator
		}

		if changedToNop {
			input[changePointer].operation = "jmp"
		} else {
			input[changePointer].operation = "nop"
		}
	}

	return 0
}

func runOpCode(input []opcode) (int, int) {
	accumulator := 0
	endState := 0
	pointer := 0

	executedCommands := make(map[int]int)

	for true {
		if _, ok := executedCommands[pointer]; ok {
			endState = 1 // Infinite loop
			break        // If we have executed this command before than break
		}

		executedCommands[pointer]++

		switch input[pointer].operation {
		case "nop":
			pointer++
		case "acc":
			accumulator += input[pointer].value
			pointer++
		case "jmp":
			pointer += input[pointer].value
		}

		if pointer >= len(input) {
			break // Terminate when the pointer is beyond the scope of the given program.
		}
	}

	return accumulator, endState
}

func getInput(path string) []opcode {
	inputFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var output []opcode

	// Get the input values from the text file and store them in a slice.
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")

		value, _ := strconv.Atoi(parts[1])

		output = append(output, opcode{parts[0], value})
	}

	return output
}
