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

type amplifier struct {
	opcode        map[int]int
	index         int
	input, output int
	stopped       bool
}

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
		amplifiers := []amplifier{}
		opcode := make(map[int]int)

		for address, codes := range opcodeString {
			i, _ := strconv.Atoi(codes)
			opcode[address] = i
		}

		for i := 0; i < 5; i++ {
			amplifiers = append(amplifiers, amplifier{opcode, 0, 0, 0, false})
		}

		thrust := 0
		// Check every combination of 0 to 4 for the phase
		for j := 56789; j <= 98765; j++ {
			phaseSettings := []int{}
			value := j

			for i := 0; i < 5; i++ { // 5 digits
				phaseSettings = append(phaseSettings, (value % 10))
				value /= 10
			}

			for i := range phaseSettings {
				amplifiers[i].input = phaseSettings[i]
			}

			if uniqueSlice(phaseSettings) {
				for !amplifiers[4].stopped {
					amplifiers[0] = runOpcode(amplifiers[0])
					amplifiers[1].input = amplifiers[0].output

					amplifiers[1] = runOpcode(amplifiers[1])
					amplifiers[2].input = amplifiers[1].output

					amplifiers[2] = runOpcode(amplifiers[2])
					amplifiers[3].input = amplifiers[2].output

					amplifiers[3] = runOpcode(amplifiers[3])
					amplifiers[4].input = amplifiers[3].output

					amplifiers[4] = runOpcode(amplifiers[4])
					amplifiers[0].input = amplifiers[4].output
				}
				if amplifiers[4].output > thrust {
					thrust = amplifiers[4].output
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

func runOpcode(in amplifier) (out amplifier) {
	i := in.index
	for i < len(in.opcode) {
		// True = immediate mode, False = position mode
		firstParameterMode := (in.opcode[i] / 100) % 10
		secondParameterMode := (in.opcode[i] / 1000) % 10

		switch in.opcode[i] % 100 {
		case 1:
			in.opcode[in.opcode[i+3]] = parameter(in.opcode, firstParameterMode, i+1) + parameter(in.opcode, secondParameterMode, i+2)
			i += 4
			break
		case 2:
			in.opcode[in.opcode[i+3]] = parameter(in.opcode, firstParameterMode, i+1) * parameter(in.opcode, secondParameterMode, i+2)
			i += 4
			break
		case 3:
			in.opcode[in.opcode[i+1]] = in.input
			i += 2
			break
		case 4:
			output := in.opcode[in.opcode[i+1]]
			i += 2
			return amplifier{
				opcode: in.opcode,
				output: output,
				input:  in.input,
				index:  i,
			}
		case 5:
			if parameter(in.opcode, firstParameterMode, i+1) != 0 {
				i = parameter(in.opcode, secondParameterMode, i+2)
			} else {
				i += 3
			}
			break
		case 6:
			if parameter(in.opcode, firstParameterMode, i+1) == 0 {
				i = parameter(in.opcode, secondParameterMode, i+2)
			} else {
				i += 3
			}
			break
		case 7:
			if parameter(in.opcode, firstParameterMode, i+1) < parameter(in.opcode, secondParameterMode, i+2) {
				in.opcode[in.opcode[i+3]] = 1
			} else {
				in.opcode[in.opcode[i+3]] = 0
			}
			i += 4
			break
		case 8:
			if parameter(in.opcode, firstParameterMode, i+1) == parameter(in.opcode, secondParameterMode, i+2) {
				in.opcode[in.opcode[i+3]] = 1
			} else {
				in.opcode[in.opcode[i+3]] = 0
			}
			i += 4
			break
		case 99:
			return amplifier{
				output:  in.output,
				stopped: true,
			}
		default:
			log.Fatal("Unknown opcode: ", in.opcode[i], " at address: ", i)
			break
		}
	}

	return amplifier{}
}

func uniqueSlice(input []int) bool {
	for i, val1 := range input {
		for j, val2 := range input {
			if i == j {
				continue
			}
			if val1 < 5 {
				return false
			}
			if val1 == val2 {
				return false
			}
		}
	}

	return true
}
