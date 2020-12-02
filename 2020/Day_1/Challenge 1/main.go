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

	inputFile, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var expenses []int

	// Get the input values from the text file and store them in a slice.
	for scanner.Scan() {
		line := scanner.Text()

		i, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal(err)
		}

		expenses = append(expenses, i)
	}

	// Loop over the slice to find the values that add up to 2020.
	for a := 0; a < len(expenses); a++ {
		for b := a + 1; b < len(expenses); b++ {
			if (expenses[a] + expenses[b]) == 2020 {
				fmt.Println(expenses[a] * expenses[b])
				break
			}
		}
	}

	fmt.Println(time.Since(start))
}
