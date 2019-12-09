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
	doublesCheck := false
	repeatCheck := false

	for i := 97; i <= 122; i++ { // Check doubles without overlap
		for j := 97; j <= 122; j++ {
			check := string(i) + string(j)
			if strings.Count(input, check) >= 2 {
				doublesCheck = true
			}

			check2 := string(i) + string(j) + string(i)
			if strings.Contains(input, check2) {
				repeatCheck = true
			}

			if repeatCheck && doublesCheck {
				return true
			}
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
