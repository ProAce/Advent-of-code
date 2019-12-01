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

		i, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal(err)
		}

		mass += (int)(i/3) - 2
	}

	fmt.Println(mass)
	fmt.Println(time.Since(start))
}
