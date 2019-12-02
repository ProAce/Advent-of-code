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
		opcode := [1000]int{} // Set length to counteract out of bounds issue

		for address, codes := range opcodeString {
			i, _ := strconv.Atoi(codes)
			opcode[address] = i
		}

		backup := opcode

	loop:
		for x := 0; x <= 99; x++ {
			for y := 0; y <= 99; y++ {
				opcode = backup

				opcode[1] = y
				opcode[2] = x

				for i := 0; i < len(opcode); i += 4 {
					if opcode[i] == 1 { // opcode 1 = addition
						opcode[opcode[i+3]] = opcode[opcode[i+1]] + opcode[opcode[i+2]]
					} else if opcode[i] == 2 { // opcode 2 = multiplication
						opcode[opcode[i+3]] = opcode[opcode[i+1]] * opcode[opcode[i+2]]
					} else if opcode[i] == 99 { // opcode 99 = termination
						break
					}
				}

				if opcode[0] == 19690720 {
					fmt.Println(100*y + x)
					break loop
				}
			}
		}
	}

	fmt.Println(time.Since(start))
}
