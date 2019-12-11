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

	asteroidMap := []point{}

	y := 0

	for scanner.Scan() {
		line := scanner.Text()

		for x, value := range line { // Make a map of all the asteroids
			switch value {
			case '#':
				asteroidMap = append(asteroidMap, point{x, y})
				break
			default:
				break
			}
		}

		y++
	}

	maximumLength := 0
	center := point{}
	angleMap := []angle{}

	for i := range asteroidMap { // Check for every asteroid if it has the longest length
		output := calculateAngles(asteroidMap, asteroidMap[i])
		if len(output) > maximumLength {
			maximumLength = len(output)
			center = asteroidMap[i]
			angleMap = output
		}
	}

	fmt.Println(center, maximumLength)

	coordinate := anihilateAsteroids(angleMap, center, 299) // Check which is the 200th asteroid to be destroyed

	fmt.Println(coordinate)

	fmt.Println(time.Since(start))
}

// calculateAngles takes an map and position and returns a sorted slice of type angle
func calculateAngles(input []point, center point) (output []angle) {
	angleMap := make(map[float64][]point)
	for i := range input {
		inputAngle := getAngle(center, input[i])
		angleMap[inputAngle] = append(angleMap[inputAngle], input[i])
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

		if len(input[i].points) <= 1 {
			copy(input[i:], input[i+1:])
			input = input[:len(input)-1]
		} else {
			input[i].points = input[i].points[1:]
			i++
		}

		if i == len(input) {
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
		x := ((input[i].x - center.x) * (input[i].x - center.x))
		y := ((input[i].y - center.y) * (input[i].y - center.y))
		distanceI := x + y
		x = ((input[j].x - center.x) * (input[j].x - center.x))
		y = ((input[j].y - center.y) * (input[j].y - center.y))
		distanceJ := x + y
		return distanceI < distanceJ
	})

	return input
}

func linear(x, inMin, inMax, outMin, outMax float64) float64 {
	return (x-inMin)*(outMax-outMin)/(inMax-inMin) + outMin
}
