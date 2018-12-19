package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func checkValue(frequency int, outputRecord []int) bool {
	for i := 0; i < len(outputRecord); i++ {
		if outputRecord[i] == frequency {
			return true
		}
	}

	return false
}

func main() {
	start := time.Now()

	frequency := 0
	input := []int{}
	outputRecord := []int{}

	found := false

	intputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer intputFile.Close()

	scanner := bufio.NewScanner(intputFile)

	for scanner.Scan() {
		line := scanner.Text()

		v, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal(err)
		}

		input = append(input, v)
	}

	for !found {
		for i := 0; i < len(input); i++ {
			frequency += input[i]
			found = checkValue(frequency, outputRecord)
			if found == true {
				break
			} else {
				outputRecord = append(outputRecord, frequency)
			}
		}
	}

	fmt.Println(frequency)

	fmt.Println(time.Since(start))

}
