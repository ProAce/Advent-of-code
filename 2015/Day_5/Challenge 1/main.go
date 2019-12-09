package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func stringNice(input string) bool {
	if strings.Contains(input, "ab") || strings.Contains(input, "cd") || strings.Contains(input, "pq") || strings.Contains(input, "xy") {
		return false
	}

	vowels := strings.Count(input, "a") + strings.Count(input, "e") + strings.Count(input, "o") + strings.Count(input, "u") + strings.Count(input, "i")
	foundDoubles := false

	for i := 97; i <= 122; i++ {
		doubleString := string(i) + string(i)
		if strings.Contains(input, doubleString) {
			foundDoubles = true
			break
		}
	}

	if vowels >= 3 && foundDoubles {
		return true
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
