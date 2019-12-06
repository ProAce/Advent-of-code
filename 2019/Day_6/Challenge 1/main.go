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
	orbitCount := 0

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

	for _, value := range orbitMap {
		orbitCount++

		i := value
		for {
			if val, exists := orbitMap[i]; exists {
				i = val
				orbitCount++
			} else {
				break
			}
		}
	}

	fmt.Println(orbitCount)
	fmt.Println(time.Since(start))
}
