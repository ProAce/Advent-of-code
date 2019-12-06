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

	firstCommon := firstCommonCharacter(santaOrbit, yourOrbit)

	orbitStepsCount := (stepsToCommonCharacter(santaOrbit, firstCommon) + stepsToCommonCharacter(yourOrbit, firstCommon))

	fmt.Println(orbitStepsCount)
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

func firstCommonCharacter(input1, input2 []string) string {
	for _, value1 := range input1 {
		for _, value2 := range input2 {
			if value1 == value2 {
				return value1
			}
		}
	}

	return ""
}

func stepsToCommonCharacter(input []string, character string) (count int) {
	for _, value := range input {
		if value == character {
			break
		} else {
			count++
		}
	}

	return count - 1 // Compensate for the first step
}
