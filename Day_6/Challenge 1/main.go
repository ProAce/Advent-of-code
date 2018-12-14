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

func buildGrid(input [][]int) (output [][]int, offset []int) {
	xLow := 10000
	xHigh := 0
	yLow := 10000
	yHigh := 0
	for i := 0; i < len(input); i++ {
		if input[i][1] > xHigh {
			xHigh = input[i][1]
		} else if input[i][1] < xLow {
			xLow = input[i][1]
		}
		if input[i][2] > yHigh {
			yHigh = input[i][2]
		} else if input[i][2] < yLow {
			yLow = input[i][2]
		}
	}

	output = make([][]int, xHigh-xLow+1)
	for i := range output {
		output[i] = make([]int, yHigh-yLow+1)
	}

	// fmt.Println(len(output), len(output[0]))

	offset = append(offset, xLow, yLow, xHigh, yHigh)

	return output, offset
}

func fillGrid(offset []int, grid [][]int, input [][]int) (outputGrid [][]int) {

	for i := 0; i < len(input); i++ {
		x := input[i][1] - offset[0]
		y := input[i][2] - offset[1]
		// fmt.Println(x, y)
		grid[x][y] = input[i][0]

	}

	outputGrid = grid
	return outputGrid
}

func countGrid(grid [][]int) (size int) {

	return size
}

func parseCoordinates(input string) (x, y int) {
	coordinates := strings.Split(input, ",")
	coordinates[1] = strings.TrimSpace(coordinates[1])

	x, _ = strconv.Atoi(coordinates[0])
	y, _ = strconv.Atoi(coordinates[1])

	return x, y
}

func main() {
	start := time.Now()

	datapoints := [][]int{}

	intputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(intputFile)

	id := 1

	for scanner.Scan() {
		input := scanner.Text()

		x, y := parseCoordinates(input)
		datapoints = append(datapoints, []int{id, x, y})
		id++
	}

	// fmt.Println(datapoints)

	grid, offset := buildGrid(datapoints)
	filledGrid := fillGrid(offset, grid, datapoints)

	for i := 0; i < len(filledGrid); i++ {
		fmt.Println(filledGrid[i])
	}

	fmt.Println(time.Since(start))
}
