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

	var two int16
	var three int16

	intputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer intputFile.Close()

	scanner := bufio.NewScanner(intputFile)

	for scanner.Scan() {
		line := scanner.Text()

		found2 := false
		found3 := false

		for i := 65; i <= 90; i++ {
			check := strings.ToLower(string(i))
			if (strings.Count(line, check) == 2) && (found2 == false) {
				two++
				found2 = true
			} else if (strings.Count(line, check) == 3) && (found3 == false) {
				three++
				found3 = true
			}
		}

	}

	checksum := two * three
	fmt.Println(checksum)
	fmt.Println(time.Since(start))
}
