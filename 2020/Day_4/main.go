package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	input := getInput("input.txt")

	part1, part2 := checkValidity(input)

	fmt.Printf("Part 1: %d\r\n", part1)
	fmt.Printf("Part 2: %d\r\n", part2)
	fmt.Println(time.Since(start))
}

func checkValidity(input []map[string]string) (int, int) {
	requirements := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	invalidPasswordsP1 := 0
	invalidPasswordsP2 := 0

	// Loop over all the passports.
	for _, value := range input {
		validity := true
		// Check all the requirements of the password.
		for _, key := range requirements {
			if _, ok := value[key]; !ok {
				invalidPasswordsP1++
				break
			}

			outcome := true

			switch key {
			case requirements[0]:
				outcome = checkByr(value[key])
				break
			case requirements[1]:
				outcome = checkIyr(value[key])
				break
			case requirements[2]:
				outcome = checkEyr(value[key])
				break
			case requirements[3]:
				outcome = checkHgt(value[key])
				break
			case requirements[4]:
				outcome = checkHcl(value[key])
				break
			case requirements[5]:
				outcome = checkEcl(value[key])
				break
			case requirements[6]:
				outcome = checkPid(value[key])
				break
			}

			if outcome == false {
				validity = false
			}
		}

		if !validity {
			invalidPasswordsP2++
		}
	}

	return (len(input) - invalidPasswordsP1), (len(input) - (invalidPasswordsP1 + invalidPasswordsP2))
}

func checkDigits(input string, min, max int) bool {
	temp, err := strconv.Atoi(input)
	if err != nil {
		return false
	}

	if (temp < min) || (temp > max) {
		return false
	}

	return true
}

func checkByr(input string) bool {
	return checkDigits(input, 1920, 2002)
}

func checkIyr(input string) bool {
	return checkDigits(input, 2010, 2020)
}

func checkEyr(input string) bool {
	return checkDigits(input, 2020, 2030)
}

func checkHgt(input string) bool {
	if strings.HasSuffix(input, "in") {
		return checkDigits(input[:(len(input)-2)], 59, 76)
	} else if strings.HasSuffix(input, "cm") {
		return checkDigits(input[:(len(input)-2)], 150, 193)
	}
	return false
}

func checkHcl(input string) bool {
	valid, _ := regexp.MatchString(`(#[a-f0-9]{6})`, input)
	return valid
}

func checkEcl(input string) bool {
	valid, _ := regexp.MatchString(`(amb|blu|brn|gry|grn|hzl|oth)\b`, input)
	return valid
}

func checkPid(input string) bool {
	return len(input) == 9
}

func getInput(path string) []map[string]string {
	inputFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var output []map[string]string
	credentials := make(map[string]string)

	// Get the input values from the text file and store them in a slice.
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			// On a new line store the map and gather the next credential package.
			output = append(output, credentials)
			credentials = make(map[string]string)
			continue
		}

		splitLine := strings.Split(line, " ")

		for _, value := range splitLine {
			splitValue := strings.Split(value, ":")
			credentials[splitValue[0]] = splitValue[1]
		}
	}

	// When we have scanned all lines store the last credentials, the file doesn't end in a new line.
	output = append(output, credentials)

	return output
}
