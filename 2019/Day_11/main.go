package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type opcode struct {
	commands     map[int]int
	relativeBase int
	input        int
	index        int
	running      bool
}

func (o *opcode) runOpcode() (output int) {
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
			o.writeParameter(firstParameterMode, 1, o.input)
			o.index += 2
			break
		case 4:
			output = o.readParameter(firstParameterMode, 1)
			o.index += 2
			return output
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
			return output
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

type point struct {
	x, y int
}

func (p point) smallerThen(in point) (out point) {
	if p.x > in.x && p.y < in.y {
		return p
	}

	return in
}

func (p point) biggerThen(in point) (out point) {
	if p.x < in.x && p.y > in.y {
		return p
	}

	return in
}

type paintRobot struct {
	opcode
	points    map[point]int
	position  point
	direction int
}

// switchDirection: 0 = left, 1 = right
func (p *paintRobot) switchDirection(right int) {
	if right == 1 {
		p.direction++
		if p.direction > 3 {
			p.direction = 0
		}
	} else {
		p.direction--
		if p.direction < 0 {
			p.direction = 3
		}
	}
}

// walkDirection walks one step in the direction the robot is facing
func (p *paintRobot) walkDirection() {
	switch p.direction {
	case 0: // Up
		p.position = point{p.position.x, p.position.y + 1}
		break
	case 1: // Right
		p.position = point{p.position.x + 1, p.position.y}
		break
	case 2: // Down
		p.position = point{p.position.x, p.position.y - 1}
		break
	case 3: // Left
		p.position = point{p.position.x - 1, p.position.y}
		break
	}
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
			running:      true,
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

		for robot.running {
			robot.input = robot.points[robot.position]

			output := []int{}
			for len(output) < 2 {
				output = append(output, robot.runOpcode())
			}

			robot.points[robot.position] = output[0]
			robot.switchDirection(output[1])
			robot.walkDirection()
		}

		fmt.Println(len(robot.points))

		robot = paintRobot{
			opcode:    opcode,
			points:    make(map[point]int),
			position:  point{0, 0},
			direction: 0,
		}

		robot.points[robot.position] = 1

		for robot.running {
			robot.input = robot.points[robot.position]

			output := []int{}
			for len(output) < 2 {
				output = append(output, robot.runOpcode())
			}

			robot.points[robot.position] = output[0]
			robot.switchDirection(output[1])
			robot.walkDirection()
		}

		minimum := point{math.MaxInt64, math.MaxInt64}
		maximum := point{math.MinInt64, math.MinInt64}
		for key := range robot.points {
			minimum = key.smallerThen(minimum)
			maximum = key.biggerThen(maximum)
		}

		upLeft := image.Point{minimum.x, minimum.y}
		lowRight := image.Point{maximum.x, maximum.y}

		img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

		white := color.White
		black := color.Black

		for key, value := range robot.points {
			if value == 1 {
				img.Set(key.x, key.y, white)
			} else {
				img.Set(key.x, key.y, black)
			}
		}

		f, _ := os.Create("SolutionDay11Part2.png")
		png.Encode(f, img)
	}

	fmt.Println(time.Since(start))
}
