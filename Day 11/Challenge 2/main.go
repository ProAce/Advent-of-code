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

func createSAT(grid [][]int) [][]int {

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j > 0 {
				grid[i][j] += grid[i][j-1]
			} else if i > 0 && j == 0 {
				grid[i][j] += grid[i-1][j]
			} else if i > 0 && j > 0 {
				grid[i][j] += grid[i-1][j] + grid[i][j-1] - grid[i-1][j-1]
			}
		}
	}

	return grid
}

func countPowerLevel(input [][]int, gridSize int) (coordinates []int) {
	max := 0
	count := 0

	for a := 1; a < gridSize; a++ {
		for i := len(input) - 1; i > 0+a; i-- {
			for j := len(input[i]) - 1; j > 0+a; j-- {
				count = input[i][j] - input[i-a][j] - input[i][j-a] + input[i-a][j-a]
				if count > max {
					max = count
					coordinates = []int{i - a + 1, j - a + 1, a}
				}
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
	// fmt.Println(grid[2])

	grid = createSAT(grid)
	// fmt.Println(grid[2])

	coordinates = countPowerLevel(grid, gridSize)

	fmt.Println(coordinates)

	fmt.Println(time.Since(start))
}
