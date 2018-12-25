package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func deliverPackages(input string, grid [][]int) (output int) {
	xSanta := 500
	ySanta := 500
	xRobo := 500
	yRobo := 500

	robo := false

	// y
	//x---
	// |
	// |

	for i := 0; i < len(input); i++ {
		if robo {
			switch string(input[i]) {
			case "^":
				ySanta--
			case ">":
				xSanta++
			case "v":
				ySanta++
			case "<":
				xSanta--
			}
			grid[xSanta][ySanta]++
			robo = !robo
		} else {
			switch string(input[i]) {
			case "^":
				yRobo--
			case ">":
				xRobo++
			case "v":
				yRobo++
			case "<":
				xRobo--
			}
			grid[xRobo][yRobo]++
			robo = !robo
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] != 0 {
				output++
			}
		}
	}

	return output
}

func main() {
	start := time.Now()

	intputFile, err := os.Open("input.txt")

	grid := make([][]int, 1000)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, 1000)
	}
	grid[500][500] = 1

	if err != nil {
		log.Fatal(err)
	}
	defer intputFile.Close()

	scanner := bufio.NewScanner(intputFile)

	for scanner.Scan() {
		directions := scanner.Text()
		fmt.Println(deliverPackages(directions, grid))
	}

	fmt.Println(time.Since(start))
}
