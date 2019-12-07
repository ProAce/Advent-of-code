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
		opcode := make(map[int]int) // Set length to counteract out of bounds issue

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

				i := 0
			opcode:
				for i < len(opcode) {
					switch opcode[i] {
					case 1:
						opcode[opcode[i+3]] = opcode[opcode[i+1]] + opcode[opcode[i+2]]
						i += 4
						break
					case 2:
						opcode[opcode[i+3]] = opcode[opcode[i+1]] * opcode[opcode[i+2]]
						i += 4
						break
					case 99:
						break opcode
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
