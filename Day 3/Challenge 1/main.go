package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInput(input string) (x, y, sizeX, sizeY int) {
	s := strings.Split(input, " ")

	s[2] = strings.TrimSuffix(s[2], ":")
	coordinates := strings.Split(s[2], ",")

	size := strings.Split(s[3], "x")

	x, _ = strconv.Atoi(coordinates[0])
	y, _ = strconv.Atoi(coordinates[1])
	sizeX, _ = strconv.Atoi(size[0])
	sizeY, _ = strconv.Atoi(size[1])
	return x, y, sizeX, sizeY
}

func main() {
	grid := [1000][1000]int{}
	doubles := 0

	intputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer intputFile.Close()

	scanner := bufio.NewScanner(intputFile)

	for scanner.Scan() {
		input := scanner.Text()

		x, y, sizeX, sizeY := parseInput(input)

		for i := 0; i < sizeX; i++ {
			for j := 0; j < sizeY; j++ {
				if grid[x+i][y+j] == 1 {
					doubles++
					grid[x+i][y+j] = 2
				} else if grid[x+i][y+j] == 2 {
					//Do nothing, this grid has already been counted as double
				} else {
					grid[x+i][y+j] = 1
				}
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(doubles)
}
