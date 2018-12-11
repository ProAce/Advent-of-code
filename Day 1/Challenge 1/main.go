package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	frequency := 0

	intputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer intputFile.Close()

	scanner := bufio.NewScanner(intputFile)

	for !found {
		for scanner.Scan() {
			line := scanner.Text()

			i, err := strconv.Atoi(line)

			if err != nil {
				log.Fatal(err)
			}

			frequency += i
		}

		fmt.Println(frequency)
	}
}
