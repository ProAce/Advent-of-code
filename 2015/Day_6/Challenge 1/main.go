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

func parseLine(input string) (output []int) {
	temp := strings.Split(input, " ")
	var onOff, xmin, ymin, xmax, ymax int

	if temp[0] == "turn" {
		coordinates := strings.Split(temp[2], ",")
		coordinates2 := strings.Split(temp[4], ",")
		xmin, _ = strconv.Atoi(coordinates[0])
		ymin, _ = strconv.Atoi(coordinates[1])
		xmax, _ = strconv.Atoi(coordinates2[0])
		ymax, _ = strconv.Atoi(coordinates2[1])
		if temp[1] == "off" {
			onOff = 0
		} else {
			onOff = 1
		}
	} else {
		coordinates := strings.Split(temp[1], ",")
		coordinates2 := strings.Split(temp[3], ",")
		xmin, _ = strconv.Atoi(coordinates[0])
		ymin, _ = strconv.Atoi(coordinates[1])
		xmax, _ = strconv.Atoi(coordinates2[0])
		ymax, _ = strconv.Atoi(coordinates2[1])
		onOff = 3
	}

	output = []int{onOff, xmin, ymin, xmax, ymax}
	return output
}

func litTheLights(input [][]int, grid [][]bool) [][]bool {
	//input of form {onOff, xmin, ymin, xmax, ymax}

	for i := 0; i < len(input); i++ {
		onOff := input[i][0] //0 = off, 1 = on, 3 = togle
		xmin := input[i][1]
		ymin := input[i][2]
		xmax := input[i][3]
		ymax := input[i][4]

		switch onOff {
		case 0:
			for k := xmin; k <= xmax; k++ {
				for l := ymin; l <= ymax; l++ {
					grid[k][l] = false
				}
			}
		case 1:
			for k := xmin; k <= xmax; k++ {
				for l := ymin; l <= ymax; l++ {
					grid[k][l] = true
				}
			}
		case 3:
			for k := xmin; k <= xmax; k++ {
				for l := ymin; l <= ymax; l++ {
					grid[k][l] = !grid[k][l]
				}
			}
		}
	}

	return grid
}

func countTheLitLights(grid [][]bool) (count int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] {
				count++
			}
		}
	}

	return count
}
func main() {
	start := time.Now()

	intputFile, err := os.Open("input.txt")

	coordinates := [][]int{}
	grid := make([][]bool, 1000)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]bool, len(grid))
	}

	if err != nil {
		log.Fatal(err)
	}
	defer intputFile.Close()

	scanner := bufio.NewScanner(intputFile)

	for scanner.Scan() {
		coordinates = append(coordinates, parseLine(scanner.Text()))
	}

	grid = litTheLights(coordinates, grid)
	fmt.Println(countTheLitLights(grid))

	fmt.Println(time.Since(start))
}
