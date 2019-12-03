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

	j := 1

	for scanner.Scan() {
		currentx := 0
		currenty := 0

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

	distance := math.MaxInt64

	for key, i := range grid {
		if i == 3 {
			coordinates := strings.Split(key, ",")
			x, _ := strconv.Atoi(coordinates[0])
			y, _ := strconv.Atoi(coordinates[1])

			x = abs(x)
			y = abs(y)

			if (x + y) < distance {
				distance = x + y
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
