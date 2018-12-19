package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var commands []string

func parseState(input string) string {
	temp := strings.Split(input, " ")
	return temp[2]
}

func parseCommands(input string) {
	temp := strings.Split(input, " ")
	if temp[2] == "#" {
		commands = append(commands, temp[0])
	}
}

func iterateCommands(input string) string {
	output := ".."
	hit := false

	for i := 2; i < len(input)-2; i++ {
		hit = false
		for j := 0; j < len(commands); j++ {
			if commands[j] == input[i-2:i+3] {
				hit = true
			}
		}

		if hit == true {
			output += "#"
		} else {
			output += "."
		}
	}

	output += "..."

	return output
}

func outcome(input string) (count int) {
	indexShift := 50000000000 - 125
	x := []rune("#")
	count = 0

	for i, n := range input {
		if n == x[0] {
			count += (i - 4) + indexShift
		}
	}

	return count
}

func main() {
	start := time.Now()

	intputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer intputFile.Close()

	scanner := bufio.NewScanner(intputFile)

	line := 0
	state := "...."

	for scanner.Scan() {
		input := scanner.Text()
		if line == 0 {
			state += parseState(input)
		} else if line > 1 {
			parseCommands(input)
		}
		line++
	}

	//Visually observed that after about 125 iterations the outcome is stable
	//For compatibility this should be decided by a function #ToDo
	for i := 0; i < 125; i++ {
		state = iterateCommands(state)
	}

	fmt.Println(outcome(state))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(time.Since(start))
}
