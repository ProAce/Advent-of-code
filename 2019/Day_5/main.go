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
	reader := bufio.NewReader(os.Stdin)

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

		i := 0
	opcode:
		for i < len(opcode) {
			// True = immediate mode, False = position mode
			firstParameterMode := (opcode[i] / 100) % 10
			secondParameterMode := (opcode[i] / 1000) % 10

			switch opcode[i] % 100 {
			case 1:
				opcode[opcode[i+3]] = parameter(opcode, firstParameterMode, i+1) + parameter(opcode, secondParameterMode, i+2)
				i += 4
				break
			case 2:
				opcode[opcode[i+3]] = parameter(opcode, firstParameterMode, i+1) * parameter(opcode, secondParameterMode, i+2)
				i += 4
				break
			case 3:
				fmt.Println("Code 3 input:") // Ask for user input

				inputString, _ := reader.ReadString('\n') // Read user input

				//Check wether it is a Unix or a Windows terminal
				if strings.HasSuffix(inputString, "\r\n") {
					inputString = strings.TrimSuffix(inputString, "\r\n")
				} else {
					inputString = strings.TrimSuffix(inputString, "\n")
				}

				input, err := strconv.Atoi(inputString)

				if err != nil {
					log.Fatal(err)
				}

				opcode[opcode[i+1]] = input
				i += 2
				break
			case 4:
				if opcode[opcode[i+1]] != 0 {
					fmt.Println(opcode[opcode[i+1]])
				}
				i += 2
				break
			case 5:
				if parameter(opcode, firstParameterMode, i+1) != 0 {
					i = parameter(opcode, secondParameterMode, i+2)
				} else {
					i += 3
				}
				break
			case 6:
				if parameter(opcode, firstParameterMode, i+1) == 0 {
					i = parameter(opcode, secondParameterMode, i+2)
				} else {
					i += 3
				}
				break
			case 7:
				if parameter(opcode, firstParameterMode, i+1) < parameter(opcode, secondParameterMode, i+2) {
					opcode[opcode[i+3]] = 1
				} else {
					opcode[opcode[i+3]] = 0
				}
				i += 4
				break
			case 8:
				if parameter(opcode, firstParameterMode, i+1) == parameter(opcode, secondParameterMode, i+2) {
					opcode[opcode[i+3]] = 1
				} else {
					opcode[opcode[i+3]] = 0
				}
				i += 4
				break
			case 99:
				break opcode
			default:
				log.Fatal("Unknown opcode:", opcode[i], "at address:", i)
				break
			}
		}
	}

	fmt.Println(time.Since(start))
}

func parameter(opcode []int, parameterMode, position int) int {
	if parameterMode == 1 {
		return opcode[position]
	}

	return opcode[opcode[position]]
}
