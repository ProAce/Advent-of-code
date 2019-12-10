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

	intputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer intputFile.Close()

	scanner := bufio.NewScanner(intputFile)

	asteroidMap := make(map[int][]byte)
	index := 0

	for scanner.Scan() {
		line := scanner.Text()

		for _, value := range line {
			switch value {
			case '.':
				asteroidMap[index] = append(asteroidMap[index], 0)
				break
			case '#':
				asteroidMap[index] = append(asteroidMap[index], 1)
				break
			default:
				asteroidMap[index] = append(asteroidMap[index], 8)
				break
			}
		}

		index++
	}

	maximumLength := 0
	maximumX := 0
	maximumY := 0

	for inputy, row := range asteroidMap {
		for inputx, value := range row {
			if value == 1 {
				angle := calculateAngles(asteroidMap, inputx, inputy)
				if len(angle) > maximumLength {
					maximumLength = len(angle)
					maximumX = inputx
					maximumY = inputy
				}
			}
		}
	}

	fmt.Println(maximumX, maximumY, maximumLength)

	fmt.Println(time.Since(start))
}

func calculateAngles(input map[int][]byte, x, y int) (angle map[float64]int) {
	angle = make(map[float64]int)

	for inputy, row := range input {
		for inputx, value := range row {
			if value == 1 {
				inputAngle := getAngle(x, y, inputx, inputy)
				angle[inputAngle] = 1
			}
		}
	}

	return angle
}

func getAngle(centerx, centery, endx, endy int) float64 {
	dy := float64(endy - centery)
	dx := float64(endx - centerx)
	theta := math.Atan2(dy, dx) // range (-PI, PI]
	theta *= 180 / math.Pi      // rads to degs, range (-180, 180]
	return theta
}

func anihilateAsteroids(input map[int][]byte, x, y, stopCounter int) int {
	angleMap := make(map[float64][]string) // key = angle, value = "x,y"

	asteroidsAnihilated := 0

	for inputy, row := range input {
		for inputx, value := range row {
			if value == 1 {
			}
		}
	}

	return x*100 + y
}

func extractXY(input string) (x, y int) {
	value := strings.Split(input, ",")
	x, _ = strconv.Atoi(value[0])
	y, _ = strconv.Atoi(value[1])
	return x, y
}
