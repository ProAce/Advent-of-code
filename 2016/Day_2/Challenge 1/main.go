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

func parseInput(input string) (output []int) {
	temp := strings.Split(input, "x")
	x, _ := strconv.Atoi(temp[0])
	y, _ := strconv.Atoi(temp[1])
	z, _ := strconv.Atoi(temp[2])
	return []int{x, y, z}
}

func calcArea(input []int) (output int) {
	x := input[0]
	y := input[1]
	z := input[2]
	smallest1 := 100
	smallest2 := 100

	for i := 0; i < len(input); i++ {
		if input[i] < smallest1 {
			smallest2 = smallest1
			smallest1 = input[i]
		} else if input[i] < smallest2 {
			smallest2 = input[i]
		}
	}

	output = (2 * x * y) + (2 * y * z) + (2 * x * z)
	// fmt.Println("Area", output, smallest1, smallest2)
	output += smallest1 * smallest2

	return output
}

func main() {
	start := time.Now()

	data := [][]int{}
	area := 0

	intputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer intputFile.Close()

	scanner := bufio.NewScanner(intputFile)

	for scanner.Scan() {
		input := scanner.Text()
		data = append(data, parseInput(input))
	}

	for i := 0; i < len(data); i++ {
		area += calcArea(data[i])
	}

	fmt.Println(area)
	fmt.Println(time.Since(start))
}
