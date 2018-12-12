package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	check := []string{
		"aA",
		"Aa",
		"bB",
		"Bb",
		"cC",
		"Cc",
		"dD",
		"Dd",
		"eE",
		"Ee",
		"fF",
		"Ff",
		"gG",
		"Gg",
		"hH",
		"Hh",
		"iI",
		"Ii",
		"jJ",
		"Jj",
		"kK",
		"Kk",
		"lL",
		"Ll",
		"mM",
		"Mm",
		"nN",
		"Nn",
		"oO",
		"Oo",
		"pP",
		"Pp",
		"qQ",
		"Qq",
		"rR",
		"Rr",
		"sS",
		"Ss",
		"tT",
		"Tt",
		"uU",
		"Uu",
		"vV",
		"Vv",
		"wW",
		"wW",
		"xX",
		"Xx",
		"yY",
		"Yy",
		"zZ",
		"Zz",
	}

	start := time.Now()

	intputFile, err := os.Open("input.txt")

	// output := ""
	input := ""
	output := ""

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(intputFile)

	for scanner.Scan() {
		input = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(check); j++ {
			if !(input[i:i+1] == check[j]) {
				output += input[i]
			}
		}

	}

	// output = input

	fmt.Println(output)

	fmt.Println(time.Since(start))
}
