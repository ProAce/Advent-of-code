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

	commands := [2][]string{}

	j := 1

	for scanner.Scan() {
		current := point{0, 0}

		line := scanner.Text()

		split := strings.Split(line, ",")
		commands[j-1] = split

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

	steps := math.MaxInt64

	for key, i := range grid {
		if i == 3 {
			takenSteps := 0

			for j := range commands {

			loop:
				for _, value := range commands[j] {
					count, err := strconv.Atoi(value[1:])

					current := point{0, 0}

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
						if current.x == key.x && current.y == key.y {
							takenSteps++
							break loop
						} else {
							takenSteps++
						}
					}
				}
			}

			if takenSteps < steps {
				steps = takenSteps
			}
		}
	}

	fmt.Println(steps)

	fmt.Println(time.Since(start))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
