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

		thrust := 0
		// Check every combination of 0 to 4 for the phase
		for j := 01234; j <= 43210; j++ {
			phaseSettings := []int{}
			value := j

			for i := 0; i < 5; i++ { // 5 digits
				phaseSettings = append(phaseSettings, (value % 10))
				value /= 10
			}

			if uniqueSlice(phaseSettings) {
				input := 0
				output := 0
				for i := 0; i < 5; i++ { // There are 5 amplifiers (A to E)
					output = runOpcode(opcode, []int{phaseSettings[i], input})
					input = output
					opcode = backup
				}

				if output > thrust {
					thrust = output
				}
			}
		}
		fmt.Println(thrust)
	}

	fmt.Println(time.Since(start))
}

func parameter(opcode map[int]int, parameterMode, position int) int {
	if parameterMode == 1 {
		return opcode[position]
	}

	return opcode[opcode[position]]
}

func runOpcode(opcode map[int]int, input []int) (output int) {
	i := 0
	inputCounter := 0
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
			opcode[opcode[i+1]] = input[inputCounter]
			inputCounter++
			i += 2
			break
		case 4:
			output = opcode[opcode[i+1]]
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

func uniqueSlice(input []int) bool {
	for i, val1 := range input {
		for j, val2 := range input {
			if i == j {
				continue
			}
			if val1 > 4 {
				return false
			}
			if val1 == val2 {
				return false
			}
		}
	}

	return true
}
