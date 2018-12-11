package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInput(input string) (id, x, y, sizeX, sizeY int) {
	s := strings.Split(input, " ")

	s[0] = strings.TrimPrefix(s[0], "#")
	s[2] = strings.TrimSuffix(s[2], ":")
	coordinates := strings.Split(s[2], ",")

	size := strings.Split(s[3], "x")

	id, _ = strconv.Atoi(s[0])
	x, _ = strconv.Atoi(coordinates[0])
	y, _ = strconv.Atoi(coordinates[1])
	sizeX, _ = strconv.Atoi(size[0])
	sizeY, _ = strconv.Atoi(size[1])
	return id, x, y, sizeX, sizeY
}

func main() {
	grid := [1000][1000]int{}

	intputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(intputFile)

	//Fill up the grid
	for scanner.Scan() {
		input := scanner.Text()

		_, x, y, sizeX, sizeY := parseInput(input)

		for i := 0; i < sizeX; i++ {
			for j := 0; j < sizeY; j++ {
				if grid[x+i][y+j] == 1 {
					grid[x+i][y+j] = 2
				} else if grid[x+i][y+j] == 2 {
					//Do nothing, this grid has already been counted as double
				} else {
					grid[x+i][y+j] = 1
				}
			}
		}
	}

	intputFile, err = os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer intputFile.Close()

	scanner = bufio.NewScanner(intputFile)

	for scanner.Scan() {
		input := scanner.Text()

		id, x, y, sizeX, sizeY := parseInput(input)

		overlap := false

		for i := 0; i < sizeX; i++ {
			for j := 0; j < sizeY; j++ {
				if grid[x+i][y+j] == 2 {
					overlap = true
				}
			}
		}

		if !overlap {
			fmt.Println(id)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
