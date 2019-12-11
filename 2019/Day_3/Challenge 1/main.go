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

type point struct {
	x, y int
}

func main() {
	start := time.Now()

	inputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	grid := make(map[point]int)

	j := 1

	for scanner.Scan() {
		current := point{0, 0}

		line := scanner.Text()

		split := strings.Split(line, ",")

		for _, value := range split {
			count, err := strconv.Atoi(value[1:])

			if err != nil {
				log.Fatal(err)
			}

			for i := 0; i < count; i++ {
				switch value[0] {
				case 'U':
					current.x++
					break
				case 'D':
					current.x--
					break
				case 'R':
					current.y++
					break
				case 'L':
					current.y--
					break
				}
				key := point{current.x, current.y}
				grid[key] += j
			}
		}
		j++
	}

	distance := math.MaxInt64

	for key, i := range grid {
		if i == 3 {
			key.x = abs(key.x)
			key.y = abs(key.y)

			if (key.x + key.y) < distance {
				distance = key.x + key.y
			}
		}
	}

	fmt.Println(distance)

	fmt.Println(time.Since(start))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
