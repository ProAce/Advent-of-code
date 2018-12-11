package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func compareStrings(input string) (int, error) {
	intputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer intputFile.Close()

	scanner := bufio.NewScanner(intputFile)

	for scanner.Scan() {
		compareInput := scanner.Text()

		mismatched := 0
		index := 0

		for i := 0; i < len(input); i++ {
			if string(input[i]) != string(compareInput[i]) {
				mismatched++
				index = i
			}
		}

		if mismatched == 1 {
			return index, nil
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return 0, errors.New("No match found")
}

func main() {
	intputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer intputFile.Close()

	scanner := bufio.NewScanner(intputFile)

	common := ""

	for scanner.Scan() {
		input := scanner.Text()

		index, err := compareStrings(input)

		if err == nil {
			if index == 0 {
				common = input[1:len(input)]
			} else if index == len(input)-1 {
				common = input[0 : len(input)-2]
			} else {
				common = input[0:index] + input[index+1:len(input)]
			}

			break
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(common)
}
