package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func reactPolymer(input string) (out string) {
	oldLength := len(input)
	newLength := 0

	for oldLength != newLength {
		oldLength = len(input)
		for j := 65; j <= 90; j++ {
			upperChar := string(j)
			lowerChar := strings.ToLower(upperChar)

			input = strings.Replace(input, upperChar+lowerChar, "", -1)
			input = strings.Replace(input, lowerChar+upperChar, "", -1)
		}
		newLength = len(input)
	}

	out = input

	return out
}

func main() {
	start := time.Now()

	intputFile, err := os.Open("input.txt")

	output := []string{}

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(intputFile)

	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(output); i++ {
		length := len(output[i])
		letter := ""

		for j := 65; j <= 90; j++ {
			removeUpper := string(j)
			removeLower := strings.ToLower(removeUpper)

			temp := strings.Replace(output[i], removeLower, "", -1)
			temp = strings.Replace(temp, removeUpper, "", -1)

			out := reactPolymer(temp)

			if len(out) < length {
				length = len(out)
				letter = removeUpper + removeLower
			}
		}

		fmt.Println(length, letter)
	}

	fmt.Println(time.Since(start))
}
