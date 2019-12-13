package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type point struct {
	x, y int
}

func (p point) findBiggest(p2 point) point {
	if p.x >= p2.x && p.y >= p2.y {
		return p
	}
	return p2
}

type opcode struct {
	commands            map[int]int
	relativeBase        int
	input               int
	index               int
	running             bool
	newInput, newOutput bool
}

func (o *opcode) runOpcode() (output []int) {
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
			if o.newInput {
				o.writeParameter(firstParameterMode, 1, o.input)
				o.index += 2
				o.newInput = false
			} else {
				return output
			}
			break
		case 4:
			output = append(output, o.readParameter(firstParameterMode, 1))
			o.index += 2
			break
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

type view struct {
	screen  map[point]color.Color
	images  []*image.Paletted
	delays  []int
	palette []color.Color
	max     point
}

func (s *view) getExtremes() {
	for key := range s.screen {
		s.max = key.findBiggest(s.max)
	}

	s.max.x++
	s.max.y++
}

type game struct {
	opcode
	view
	paddle        point
	ball, oldball point
	ballDirection int // 0 = topleft, 1 = topright, 2 = bottomright, 3 = bottomleft
	output        []int
	score         int
}

func (g *game) runGame() {
	g.running = true
	first := true

	for g.running {
		g.output = g.runOpcode()
		g.procesOutput()

		if first {
			first = false
			g.getExtremes()
		}

		g.drawScreen()
		g.autoPilot()
	}

	g.generateGIF()
}

func (g *game) drawScreen() {
	img := image.NewPaletted(image.Rectangle{image.Point{0, 0}, image.Point{g.max.x, g.max.y}}, g.palette)

	for key, value := range g.screen {
		img.Set(key.x, key.y, value)
	}

	g.images = append(g.images, img)
	g.delays = append(g.delays, 25)
}

func (g *game) generateGIF() {
	f, _ := os.OpenFile("Part2.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: g.images,
		Delay: g.delays,
	})
}

func (g *game) procesOutput() {
	scorePoint := point{-1, 0}

	for i := 0; i < len(g.output)-3; i += 3 {
		p := point{g.output[i], g.output[i+1]}

		if p == scorePoint {
			g.score = g.output[i+2]
		} else {
			tile := g.palette[0]

			switch g.output[i+2] {
			case 1: // Wall tile
				tile = g.palette[4] // Set to white
				break
			case 2: // Block tile
				tile = g.palette[1] // Set to blue
				break
			case 3: // Paddle tile
				tile = g.palette[2] // Set to green
				g.paddle = p
				break
			case 4: // Ball tile
				tile = g.palette[3] // Set to red
				g.ball = p
				break
			}

			g.screen[p] = tile
		}
	}
}

func (g *game) autoPilot() {
	if g.ball.x > g.paddle.x {
		g.input = 1
	} else if g.ball.x < g.paddle.x {
		g.input = -1
	} else {
		g.input = 0
	}

	g.newInput = true
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

		game := game{
			opcode: opcode{
				commands: make(map[int]int),
			},
			view: view{
				screen: make(map[point]color.Color),
				palette: []color.Color{
					color.RGBA{0x00, 0x00, 0x00, 0xff}, // Black
					color.RGBA{0x00, 0x00, 0xff, 0xff}, // Blue
					color.RGBA{0x00, 0xff, 0x00, 0xff}, // Green
					color.RGBA{0xff, 0x00, 0x00, 0xff}, // Red
					color.RGBA{0xff, 0xff, 0xff, 0xff}, // White
				},
			},
		}

		opcodeString := strings.Split(line, ",")

		for address, codes := range opcodeString {
			i, _ := strconv.Atoi(codes)
			game.commands[address] = i
		}

		game.runGame()

		fmt.Println(game.score)
	}

	fmt.Println(time.Since(start))
}
