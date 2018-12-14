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

var images []*image.Paletted
var delay []int
var palette = []color.Color{
	color.RGBA{0xff, 0xff, 0xff, 0xff},
}

func parseData(input string) (output []int) {
	input = strings.TrimPrefix(input, "position=<")
	input = strings.TrimSuffix(input, ">")
	input = strings.TrimSpace(input)

	temp := strings.Split(input, ",")
	temp[2] = strings.TrimSpace(temp[2])

	temp2 := strings.Split(temp[1], "<")
	temp2[0] = strings.TrimSuffix(temp2[0], "> velocity=")
	temp2[0] = strings.TrimSpace(temp2[0])
	temp2[1] = strings.TrimSpace(temp2[1])

	x, _ := strconv.Atoi(temp[0])
	y, _ := strconv.Atoi(temp2[0])
	xV, _ := strconv.Atoi(temp2[1])
	yV, _ := strconv.Atoi(temp[2])

	output = append(output, x, y, xV, yV)

	return output
}

func updatePos(positions [][]int, iterationsize int) (output [][]int, offset []int) {
	xMin := 100000
	xMax := 0
	yMin := 100000
	yMax := 0

	for i := 0; i < len(positions); i++ {
		positions[i][0] += positions[i][2] * iterationsize
		if positions[i][0] > xMax {
			xMax = positions[i][0]
		}
		if positions[i][0] < xMin {
			xMin = positions[i][0]
		}

		positions[i][1] += positions[i][3] * iterationsize
		if positions[i][1] > yMax {
			yMax = positions[i][1]
		}
		if positions[i][1] < xMin {
			yMin = positions[i][1]
		}

	}

	offset = []int{xMax, xMin, yMax, yMin}

	return positions, offset
}

func drawImageFrame(datapoints [][]int, offset []int, width int, height int) {
	img := image.NewPaletted(image.Rect(0, 0, width, height), palette)
	for i := 0; i < len(datapoints); i++ {
		x := datapoints[i][0]
		y := datapoints[i][1]

		img.Set(x, y, palette[0])
	}

	images = append(images, img)
	delay = append(delay, 0)
}

func createGif() {
	f, _ := os.Create("rgb.gif")
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delay,
	})
}

func main() {
	start := time.Now()

	intputFile, err := os.Open("input.txt")

	positions := [][]int{}
	offset := []int{}

	if err != nil {
		log.Fatal(err)
	}
	defer intputFile.Close()

	scanner := bufio.NewScanner(intputFile)

	for scanner.Scan() {
		positions = append(positions, parseData(scanner.Text()))
	}

	positions, offset = updatePos(positions, 100)

	for offset[0]-offset[1] > 1000 {
		positions, offset = updatePos(positions, 100)
	}

	for offset[0]-offset[1] > 100 {
		positions, offset = updatePos(positions, 10)
	}

	width := offset[0] - offset[1]
	height := offset[2] - offset[3]

	fmt.Println(width, height, offset)

	for offset[0]-offset[1] < 100 {
		positions, _ = updatePos(positions, 1)

		drawImageFrame(positions, offset, width, height)
		if offset[0]-offset[1] < 10 {
			break
		}
	}

	createGif()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(time.Since(start))
}
