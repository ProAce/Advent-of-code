package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	inputFile, err := os.Open("input.txt")

	orbitMap := make(map[string]string)

	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()

		input := strings.Split(line, ")")
		orbitMap[input[1]] = input[0]
	}

	santaOrbit := buildString("SAN", orbitMap)
	yourOrbit := buildString("YOU", orbitMap)

	x, y := firstCommonCharacterPosition(santaOrbit, yourOrbit)

	fmt.Println(x + y)
	fmt.Println(time.Since(start))
}

func buildString(key string, orbitMap map[string]string) (output []string) {
	i := key
	output = append(output, i)
	for {
		if value, exists := orbitMap[i]; exists {
			output = append(output, value)
			i = value
		} else {
			return output
		}
	}
}

func firstCommonCharacterPosition(input1, input2 []string) (x, y int) {
	for x, value1 := range input1 {
		for y, value2 := range input2 {
			if value1 == value2 {
				return x - 1, y - 1 // Compensate for index offset
			}
		}
	}

	return -1, -1
}
