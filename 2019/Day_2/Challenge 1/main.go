package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	inputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()

		opcodeString := strings.Split(line, ",")
		opcode := make([]int, 1000) // Set length to counteract out of bounds issue

		for address, codes := range opcodeString {
			i, _ := strconv.Atoi(codes)
			opcode[address] = i
		}

		for i := 0; i < len(opcode); i += 4 {
			if opcode[i] == 1 { // opcode 1 = addition
				opcode[opcode[i+3]] = opcode[opcode[i+1]] + opcode[opcode[i+2]]
			} else if opcode[i] == 2 { // opcode 2 = multiplication
				opcode[opcode[i+3]] = opcode[opcode[i+1]] * opcode[opcode[i+2]]
			} else if opcode[i] == 99 { // opcode 99 = termination
				break
			}
		}
		fmt.Println(opcode[0])
	}

	fmt.Println(time.Since(start))
}
