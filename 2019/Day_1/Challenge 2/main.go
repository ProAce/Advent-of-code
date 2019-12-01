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
	mass := 0

	inputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()

		input, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal(err)
		}

		moduleMass := (int)(input/3) - 2
		mass += moduleMass

		i := moduleMass
		for true {
			i = (int)(i/3) - 2

			if i <= 0 {
				break
			}

			mass += i
		}
	}

	fmt.Println(mass)
	fmt.Println(time.Since(start))
}
