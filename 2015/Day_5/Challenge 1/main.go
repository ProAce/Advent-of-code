package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func stringNice(input string) bool {
	vowels := 0
	foundDoubles := false
	for i := 0; i < len(input)-1; i++ {
		test := string(input[i] + input[i+1])
		if test == "ab" || test == "cd" || test == "pq" || test == "xy" {
			return false
		}

		if !foundDoubles {
			for j := 65; j <= 90; j++ {
				if test == string(j)+string(j) {
					foundDoubles = true
				}
			}
		}

		test = string(input[i])
		if vowels < 3 {
			if test == "a" || test == "e" || test == "o" || test == "i" || test == "u" {
				vowels++
			}
		}

		if foundDoubles && vowels >= 3 {
			return true
		}
	}
	return false
}

func main() {
	start := time.Now()

	intputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer intputFile.Close()

	scanner := bufio.NewScanner(intputFile)

	count := 0
	for scanner.Scan() {
		input := scanner.Text()
		if stringNice(input) {
			count++
		}
	}

	fmt.Println(count)

	fmt.Println(time.Since(start))
}
