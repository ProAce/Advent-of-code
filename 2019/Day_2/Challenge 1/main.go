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
			case 3:
				// opcode[opcode[i+1]] = input
				// i++
				break
			case 4:
				break
			case 99:
				break opcode
			}
		}
		fmt.Println(opcode[0])
	}

	fmt.Println(time.Since(start))
}
