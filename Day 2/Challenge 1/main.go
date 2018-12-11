package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	check := []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
		"j",
		"k",
		"l",
		"m",
		"n",
		"o",
		"p",
		"q",
		"r",
		"s",
		"t",
		"u",
		"v",
		"w",
		"x",
		"y",
		"z",
	}

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

		for i := 0; i < len(check); i++ {
			if (strings.Count(line, check[i]) == 2) && (found2 == false) {
				two++
				found2 = true
			} else if (strings.Count(line, check[i]) == 3) && (found3 == false) {
				three++
				found3 = true
			}
		}
	}

	checksum := two * three
	fmt.Println(checksum)
}
