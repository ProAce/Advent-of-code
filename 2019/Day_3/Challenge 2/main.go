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

func main() {
	start := time.Now()

	inputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	grid := make(map[string]int)

	commands := [2][]string{}

	j := 1

	for scanner.Scan() {
		currentx := 0
		currenty := 0

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
					currentx++
					break
				case 'D':
					currentx--
					break
				case 'R':
					currenty++
					break
				case 'L':
					currenty--
					break
				}
				key := strconv.Itoa(currentx) + "," + strconv.Itoa(currenty)
				grid[key] += j
			}
		}
		j++
	}

	steps := math.MaxInt64

	for key, i := range grid {
		if i == 3 {
			coordinates := strings.Split(key, ",")
			x, _ := strconv.Atoi(coordinates[0])
			y, _ := strconv.Atoi(coordinates[1])

			takenSteps := 0

			for j := range commands {
				_x := 0
				_y := 0

			loop:
				for _, value := range commands[j] {
					count, err := strconv.Atoi(value[1:])

					if err != nil {
						log.Fatal(err)
					}

					for i := 0; i < count; i++ {
						switch value[0] {
						case 'U':
							_x++
							break
						case 'D':
							_x--
							break
						case 'R':
							_y++
							break
						case 'L':
							_y--
							break
						}
						if _x == x && _y == y {
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
