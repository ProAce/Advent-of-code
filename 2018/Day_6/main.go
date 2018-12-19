package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

func fillGrid(offset []int, grid [][]int, input [][]int) ([][]int, int) {

	for i := 0; i < len(input); i++ {
		x := input[i][1] - offset[0]
		input[i][1] = x
		y := input[i][2] - offset[1]
		input[i][2] = y
		// fmt.Println(x, y)
		if x == 0 || y == 0 || x == len(grid)-1 || y == len(grid[0])-1 {
			grid[x][y] = -1
		} else {
			grid[x][y] = input[i][0]
		}
	}

	tilesLessThen := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			minimumDistance := len(grid) + len(grid[i])
			minimumID := 0
			sumDistance := 0
			for k := 0; k < len(input); k++ {
				distance := int(math.Abs(float64(i-input[k][1])) + math.Abs(float64(j-input[k][2])))
				sumDistance += distance
				if distance == minimumDistance && minimumID != -1 {
					input[minimumID][3]--
					minimumID = -1
				}
				if distance < minimumDistance {
					if minimumID != -1 {
						input[minimumID][3]--
					}
					minimumID = k
					minimumDistance = distance
					input[k][3]++
				}

			}
			if sumDistance < 10000 {
				tilesLessThen++
			}
		}
	}

	return input, tilesLessThen
}

func biggestArea(datapoints [][]int) (count int) {
	for i := 0; i < len(datapoints); i++ {
		if datapoints[i][3] > count && datapoints[i][0] != -1 {
			count = datapoints[i][3]
		}
	}

	return count
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
		datapoints = append(datapoints, []int{id, x, y, 0})
		id++
	}

	// fmt.Println(datapoints)

	grid, offset := buildGrid(datapoints)
	output, tilesLessThen := fillGrid(offset, grid, datapoints)

	fmt.Println(biggestArea(output), tilesLessThen)

	fmt.Println(time.Since(start))
}
