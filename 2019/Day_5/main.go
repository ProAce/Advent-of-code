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
		opcode := make(map[int]int)
		backup := make(map[int]int)

		for address, codes := range opcodeString {
			i, _ := strconv.Atoi(codes)
			opcode[address] = i
			backup[address] = i
		}

		fmt.Println(runOpcode(opcode, 1))

		opcode = backup

		fmt.Println(runOpcode(opcode, 5))
	}

	fmt.Println(time.Since(start))
}

func parameter(opcode map[int]int, parameterMode, position int) int {
	if parameterMode == 1 {
		return opcode[position]
	}

	return opcode[opcode[position]]
}

func runOpcode(opcode map[int]int, input int) (output []int) {
	i := 0
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
			opcode[opcode[i+1]] = input
			i += 2
			break
		case 4:
			if opcode[opcode[i+1]] != 0 {
				output = append(output, opcode[opcode[i+1]])
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
			return output
		default:
			log.Fatal("Unknown opcode: ", opcode[i], " at address: ", i)
			break
		}
	}

	return output
}
