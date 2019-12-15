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

type opcode struct {
	commands map[int]int

	relativeBase int
	index        int

	running    bool
	firstStart bool

	input  int
	output int
	phase  int
}

func (o *opcode) runOpcode() {
	o.running = true
	for {
		// True = immediate mode, False = position mode
		firstParameterMode := (o.commands[o.index] / 100) % 10
		secondParameterMode := (o.commands[o.index] / 1000) % 10
		thirdParameterMode := (o.commands[o.index] / 10000) % 10

		switch o.commands[o.index] % 100 {
		case 1:
			value := o.readParameter(firstParameterMode, 1) + o.readParameter(secondParameterMode, 2)
			o.writeParameter(thirdParameterMode, 3, value)
			o.index += 4
			break
		case 2:
			value := o.readParameter(firstParameterMode, 1) * o.readParameter(secondParameterMode, 2)
			o.writeParameter(thirdParameterMode, 3, value)
			o.index += 4
			break
		case 3:
			if o.firstStart {
				o.firstStart = false
				o.writeParameter(firstParameterMode, 1, o.phase)
			} else {
				o.writeParameter(firstParameterMode, 1, o.input)
			}
			o.index += 2
			break
		case 4:
			o.output = o.readParameter(firstParameterMode, 1)
			o.index += 2
			return
		case 5:
			if o.readParameter(firstParameterMode, 1) != 0 {
				o.index = o.readParameter(secondParameterMode, 2)
			} else {
				o.index += 3
			}
			break
		case 6:
			if o.readParameter(firstParameterMode, 1) == 0 {
				o.index = o.readParameter(secondParameterMode, 2)
			} else {
				o.index += 3
			}
			break
		case 7:
			if o.readParameter(firstParameterMode, 1) < o.readParameter(secondParameterMode, 2) {
				o.writeParameter(thirdParameterMode, 3, 1)
			} else {
				o.writeParameter(thirdParameterMode, 3, 0)
			}
			o.index += 4
			break
		case 8:
			if o.readParameter(firstParameterMode, 1) == o.readParameter(secondParameterMode, 2) {
				o.writeParameter(thirdParameterMode, 3, 1)
			} else {
				o.writeParameter(thirdParameterMode, 3, 0)
			}
			o.index += 4
			break
		case 9:
			o.relativeBase += o.readParameter(firstParameterMode, 1)
			o.index += 2
			break
		case 99:
			o.running = false
			return
		default:
			log.Fatal("Unknown opcode: ", o.commands[o.index], " at address: ", o.index)
			break
		}
	}
}

func (o *opcode) readParameter(parameterMode, position int) int {
	position += o.index
	if parameterMode == 1 {
		return o.commands[position]
	}
	if parameterMode == 2 {
		return o.commands[o.relativeBase+o.commands[position]]
	}

	return o.commands[o.commands[position]]
}

func (o *opcode) writeParameter(parameterMode, position, value int) {
	position += o.index
	if parameterMode == 1 {
		o.commands[position] = value
		return
	}
	if parameterMode == 2 {
		o.commands[o.relativeBase+o.commands[position]] = value
		return
	}

	o.commands[o.commands[position]] = value
}

type amplifier struct {
	op     opcode
	backup opcode
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
		op := opcode{commands: make(map[int]int)}
		backup := op

		for address, codes := range opcodeString {
			i, _ := strconv.Atoi(codes)
			op.commands[address] = i
		}

		thrust := 0
		// Check every combination of 5 to 9 for the phase
		for j := 56789; j <= 98765; j++ {
			phaseSettings := []int{}
			value := j

			for i := 0; i < 5; i++ { // 5 digits
				phaseSettings = append(phaseSettings, (value % 10))
				value /= 10
			}

			amplifiers := []amplifier{}

			for i := 0; i < 5; i++ { // Create 5 new amplifiers
				amplifiers = append(amplifiers, amplifier{
					op:     op,
					backup: op,
				})
			}

			if uniqueSlice(phaseSettings) {
				for i := range amplifiers { // Set the correct settings
					amplifiers[i].op = backup
					amplifiers[i].op.phase = phaseSettings[i]
					amplifiers[i].op.running = true
					amplifiers[i].op.firstStart = true
				}

				for amplifiers[4].op.running {
					amplifiers[0].op.runOpcode()
					amplifiers[1].op.input = amplifiers[0].op.output
					amplifiers[1].op.runOpcode()
					amplifiers[2].op.input = amplifiers[1].op.output
					amplifiers[2].op.runOpcode()
					amplifiers[3].op.input = amplifiers[2].op.output
					amplifiers[3].op.runOpcode()
					amplifiers[4].op.input = amplifiers[3].op.output
					amplifiers[4].op.runOpcode()
					amplifiers[0].op.input = amplifiers[4].op.output
				}

				if amplifiers[4].op.output > thrust {
					thrust = amplifiers[4].op.output
				}
			}
		}

		fmt.Println(thrust)
	}

	fmt.Println(time.Since(start))
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
