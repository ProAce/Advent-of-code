package main

import (
	"fmt"
	"time"
)

func drawGrid(gridSize int, gridID int) (grid [][]int) {
	grid = make([][]int, gridSize)

	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, gridSize)
		for j := 0; j < len(grid[i]); j++ {
			rackID := i + 10
			powerLevel := (((rackID) * j) + gridID) * rackID
			powerLevel = (powerLevel / 100) % 10
			powerLevel -= 5

			grid[i][j] = powerLevel
		}
	}

	return grid
}

func countPowerLevel(input [][]int) (coordinates []int) {
	max := 0
	count := 0

	for i := 0; i < len(input)-2; i++ {
		for j := 0; j < len(input[i])-2; j++ {
			count = 0
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					count += input[i+k][j+l]
				}
			}
			if count > max {
				fmt.Println(count)
				max = count
				coordinates = []int{i, j}
			}
		}
	}

	return coordinates
}

func main() {
	start := time.Now()

	gridID := 6303
	gridSize := 300
	coordinates := []int{}

	grid := drawGrid(gridSize, gridID)

	coordinates = countPowerLevel(grid)

	fmt.Println(coordinates)

	fmt.Println(time.Since(start))
}
