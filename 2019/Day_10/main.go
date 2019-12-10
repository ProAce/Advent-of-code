package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"time"
)

type point struct {
	x, y int
}

type angle struct {
	points []point
	angle  float64
}

func (p point) distanceTo(centerx, centery int) int {
	return (p.x - centerx) + (p.y - centery)
}

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

		for _, value := range line { // Make a map of all the asteroids
			switch value {
			case '.':
				asteroidMap[index] = append(asteroidMap[index], 0)
				break
			case '#':
				asteroidMap[index] = append(asteroidMap[index], 1)
				break
			default:
				break
			}
		}

		index++
	}

	maximumLength := 0
	center := point{}
	angleMap := []angle{}

	for inputy, row := range asteroidMap { // Check for every asteroid if it has the longest length
		for inputx, value := range row {
			if value == 1 {
				output := calculateAngles(asteroidMap, point{inputx, inputy})
				if len(output) > maximumLength {
					maximumLength = len(output)
					center = point{inputx, inputy}
					angleMap = output
				}
			}
		}
	}

	fmt.Println(center, maximumLength)

	coordinate := anihilateAsteroids(angleMap, center, 200) // Check which is the 200th asteroid to be destroyed

	fmt.Println(coordinate)

	fmt.Println(time.Since(start))
}

// calculateAngles takes an map and position and returns a sorted slice of type angle
func calculateAngles(input map[int][]byte, center point) (output []angle) {
	angleMap := make(map[float64][]point)
	for inputy, row := range input {
		for inputx, value := range row {
			if value == 1 {
				end := point{inputx, inputy}
				inputAngle := getAngle(center, end)

				val := point{inputx, inputy}
				angleMap[inputAngle] = append(angleMap[inputAngle], val)
			}
		}
	}

	for key := range angleMap {
		output = append(output, angle{angleMap[key], key})
	}

	sort.SliceStable(output, func(i, j int) bool {
		return output[i].angle < output[j].angle
	})

	return output
}

// anihilateAsteroids removes asteroids clockwise untl it has reached the stopCounter
func anihilateAsteroids(input []angle, center point, stopCounter int) int {
	asteroidsAnihilated := 0

	for key := range input {
		input[key].points = sortPositions(input[key].points, center)
	}

	i := 0
	for {
		asteroidsAnihilated++

		if asteroidsAnihilated == stopCounter {
			return input[i].points[0].x*100 + input[i].points[0].y
		}

		input[i].points = input[i].points[:1]

		i++
		if i == len(input)-1 {
			i = 0
		}
	}
}

// getAngle returns an angle between the two points and the y axis in degrees
func getAngle(center, end point) float64 {
	dy := float64(end.y - center.y)
	dx := float64(end.x - center.x)
	theta := math.Atan2(dy, dx) // Range (-PI, PI)
	theta *= 180 / math.Pi      // Rads to degs, range (-180, 180)
	if theta < -90 {
		theta = linear(theta, -180, -90, 270, 360)
	} else {
		theta = linear(theta, -90, 180, 0, 270)
	}
	return theta
}

// sortPositions returns a sorted point slice based on its distance to center
func sortPositions(input []point, center point) []point {
	sort.SliceStable(input, func(i, j int) bool {
		distanceI := (input[i].x - center.x) + (input[i].y - center.y)
		distanceJ := (input[j].x - center.x) + (input[j].y - center.y)
		return distanceI < distanceJ
	})

	return input
}

func linear(x, inMin, inMax, outMin, outMax float64) float64 {
	return (x-inMin)*(outMax-outMin)/(inMax-inMin) + outMin
}
