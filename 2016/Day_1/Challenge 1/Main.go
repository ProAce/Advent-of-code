package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func whichFloor(input string) (floor int) {
	for i := 0; i < len(input); i++ {
		switch string(input[i]) {
		case "(":
			floor++
		case ")":
			floor--
		}
	}

	return floor
}

func main() {
	start := time.Now()

	intputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer intputFile.Close()

	scanner := bufio.NewScanner(intputFile)

	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println("Floor:", whichFloor(input))
	}

	fmt.Println(time.Since(start))
}
