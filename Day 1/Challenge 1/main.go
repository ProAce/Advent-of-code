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
	frequency := 0

	intputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer intputFile.Close()

	scanner := bufio.NewScanner(intputFile)

	for scanner.Scan() {
		line := scanner.Text()

		i, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal(err)
		}

		frequency += i
	}

	fmt.Println(frequency)
	fmt.Println(time.Since(start))
}
