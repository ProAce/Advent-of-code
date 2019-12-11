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
	commands     map[int]int
	relativeBase int
}

type point struct {
	x, y int
}

type paintRobot struct {
	opcode    opcode
	points    map[point]int
	position  point
	direction int
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

		opcode := opcode{
			commands:     make(map[int]int),
			relativeBase: 0,
		}

		for address, codes := range opcodeString {
			i, _ := strconv.Atoi(codes)
			opcode.commands[address] = i
		}

		robot := paintRobot{
			opcode:    opcode,
			points:    make(map[point]int),
			position:  point{0, 0},
			direction: 0,
		}

		input := make(chan int)
		output := make(chan int)

		go runOpcode(opcode, input, output)

		for {
			_, ok := <-output
			if !ok {
				break
			}
			input <- robot.points[robot.position]

			out := make([]int, 2)
			for i := range out {
				out[i] = <-output
			}

			robot.points[robot.position] = out[0]

			robot.direction = switchDirection(robot, out[1])
			robot.position = walkDirection(robot)
		}

		fmt.Println(len(robot.points))
	}

	fmt.Println(time.Since(start))
}

// switchDirection: 0 = up, 1 = right, 2 = down, 3 = left
func switchDirection(robot paintRobot, right int) int {

	if right == 1 {
		robot.direction++
		if robot.direction > 3 {
			robot.direction = 0
		}
	} else {
		robot.direction--
		if robot.direction < 0 {
			robot.direction = 3
		}
	}

	return robot.direction
}

func walkDirection(robot paintRobot) point {
	switch robot.direction {
	case 0: // Up
		robot.position = point{robot.position.x, robot.position.y + 1}
		break
	case 1: // Right
		robot.position = point{robot.position.x + 1, robot.position.y}
		break
	case 2: // Down
		robot.position = point{robot.position.x, robot.position.y - 1}
		break
	case 3: // Left
		robot.position = point{robot.position.x - 1, robot.position.y}
		break
	}

	return robot.position
}

func runOpcode(op opcode, input, output chan int) {
	i := 0
	for i < len(op.commands) {
		// True = immediate mode, False = position mode
		firstParameterMode := (op.commands[i] / 100) % 10
		secondParameterMode := (op.commands[i] / 1000) % 10
		thirdParameterMode := (op.commands[i] / 10000) % 10

		switch op.commands[i] % 100 {
		case 1:
			value := readParameter(&op, firstParameterMode, i+1) + readParameter(&op, secondParameterMode, i+2)
			writeParameter(&op, thirdParameterMode, i+3, value)
			i += 4
			break
		case 2:
			value := readParameter(&op, firstParameterMode, i+1) * readParameter(&op, secondParameterMode, i+2)
			writeParameter(&op, thirdParameterMode, i+3, value)
			i += 4
			break
		case 3:
			writeParameter(&op, firstParameterMode, i+1, <-input)
			i += 2
			break
		case 4:
			output <- readParameter(&op, firstParameterMode, i+1)
			i += 2
			break
		case 5:
			if readParameter(&op, firstParameterMode, i+1) != 0 {
				i = readParameter(&op, secondParameterMode, i+2)
			} else {
				i += 3
			}
			break
		case 6:
			if readParameter(&op, firstParameterMode, i+1) == 0 {
				i = readParameter(&op, secondParameterMode, i+2)
			} else {
				i += 3
			}
			break
		case 7:
			if readParameter(&op, firstParameterMode, i+1) < readParameter(&op, secondParameterMode, i+2) {
				writeParameter(&op, thirdParameterMode, i+3, 1)
			} else {
				writeParameter(&op, thirdParameterMode, i+3, 0)
			}
			i += 4
			break
		case 8:
			if readParameter(&op, firstParameterMode, i+1) == readParameter(&op, secondParameterMode, i+2) {
				writeParameter(&op, thirdParameterMode, i+3, 1)
			} else {
				writeParameter(&op, thirdParameterMode, i+3, 0)
			}
			i += 4
			break
		case 9:
			op.relativeBase += readParameter(&op, firstParameterMode, i+1)
			i += 2
			break
		case 99:
			close(input)
			close(output)
			return
		default:
			log.Fatal("Unknown opcode: ", op.commands[i], " at address: ", i)
			break
		}
	}
}

func readParameter(op *opcode, parameterMode, position int) int {

	if parameterMode == 1 {
		return op.commands[position]
	}
	if parameterMode == 2 {
		return op.commands[op.relativeBase+op.commands[position]]
	}

	return op.commands[op.commands[position]]
}

func writeParameter(op *opcode, parameterMode, position, value int) {
	if parameterMode == 1 {
		op.commands[position] = value
		return
	}
	if parameterMode == 2 {
		op.commands[op.relativeBase+op.commands[position]] = value
		return
	}

	op.commands[op.commands[position]] = value
}
