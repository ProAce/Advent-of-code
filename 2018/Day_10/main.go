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
		if positions[i][1] < yMin {
			yMin = positions[i][1]
		}
	}

	offset = []int{xMax, xMin, yMax, yMin}

	return positions, offset
}

func drawInTerminal(datapoints [][]int, offset []int) {

	for i := offset[3]; i <= offset[2]; i++ {
		for j := offset[1]; j <= offset[0]; j++ {
			print := false
			for k := 0; k < len(datapoints); k++ {
				if datapoints[k][0] == j && datapoints[k][1] == i && print == false {
					print = true
					fmt.Print("#")
				}
			}
			if print == false {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
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

	fmt.Println(offset)

	count := 0
	second := 0

	for count < 2 {
		positions, offset = updatePos(positions, 1)
		second++
		if offset[2]-offset[3] < 25 {
			fmt.Println(offset)
			drawInTerminal(positions, offset)
			count++
		}
	}

	fmt.Println(second)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(time.Since(start))
}
