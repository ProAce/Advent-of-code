package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	inputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()

		imageLayerLength := 25 * 6
		leastZeroesIndex := 0
		leastZeroes := math.MaxInt64

		for i := 0; i < len(line); i += imageLayerLength {
			amount := strings.Count(line[i:i+imageLayerLength], "0")

			if amount < leastZeroes {
				leastZeroes = amount
				leastZeroesIndex = i
			}
		}

		ones := strings.Count(line[leastZeroesIndex:leastZeroesIndex+imageLayerLength], "1")
		twos := strings.Count(line[leastZeroesIndex:leastZeroesIndex+imageLayerLength], "2")
		count := ones * twos

		fmt.Println(count)
	}

	fmt.Println(time.Since(start))
}
