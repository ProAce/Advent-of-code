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

func parseData(input string) (output []int) {
	temp := strings.Split(input, " ")

	for i := 0; i < len(temp); i++ {
		tempInt, _ := strconv.Atoi(temp[i])
		output = append(output, tempInt)
	}

	return output
}

func parsePuzzleinput(input []int) (count int) {

	return count
}

func main() {
	start := time.Now()

	intputFile, err := os.Open("input.txt")

	puzzleInput := []int{}
	count := 0

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(intputFile)

	for scanner.Scan() {
		puzzleInput = parseData(scanner.Text())
		count = parsePuzzleinput(puzzleInput)
		fmt.Println("Sum of metadata", count)
	}

	fmt.Println(count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(time.Since(start))
}
