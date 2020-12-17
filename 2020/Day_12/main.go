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
	// part2 := part2(input)

	fmt.Printf("Part 1: %d\r\n", part1)
	// fmt.Printf("Part 2: %d\r\n", part2)

	fmt.Println(time.Since(start))
}

func part1(input []instruction) int {
	north, east := 0, 0
	direction := 0

	for _, instruction := range input {
		switch instruction.direction {
		case 'L':
			direction -= instruction.value
			break
		case 'R':
			direction += instruction.value
			break
		case 'N':
			north += instruction.value
			break
		case 'S':
			north -= instruction.value
			break
		case 'E':
			east += instruction.value
			break
		case 'W':
			east -= instruction.value
			break
		case 'F':
			switch direction {
			case 0:
				east += instruction.value
				break
			case 90:
				north -= instruction.value
				break
			case 180:
				east -= instruction.value
				break
			case 270:
				north += instruction.value
				break
			}
		}

		for {
			if direction < 0 {
				direction += 360
			}

			break
		}

		for {
			if direction > 270 {
				direction -= 360
			}

			break
		}
	}

	return abs(north) + abs(east)
}

func part2(input []instruction) int {

	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

type instruction struct {
	direction rune
	value     int
}

func getInput(path string) []instruction {
	inputFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	output := []instruction{}

	// Get the input values from the text file and store them in a slice.
	for scanner.Scan() {
		line := scanner.Text()

		val, err := strconv.Atoi(line[1:])

		if err != nil {
			log.Fatal(err)
		}

		instruction := instruction{
			direction: rune(line[0]),
			value:     val,
		}

		output = append(output, instruction)
	}

	return output
}
