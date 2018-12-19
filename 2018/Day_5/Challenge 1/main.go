package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

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
		oldLength := len(output[i])
		newLength := 0

		for oldLength != newLength {
			oldLength = len(output[i])
			for j := 65; j <= 90; j++ {
				upperChar := string(j)
				lowerChar := strings.ToLower(upperChar)

				output[i] = strings.Replace(output[i], upperChar+lowerChar, "", -1)
				output[i] = strings.Replace(output[i], lowerChar+upperChar, "", -1)
			}
			newLength = len(output[i])
		}
		fmt.Println(len(output[i]))
	}
	fmt.Println(time.Since(start))
}
