package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func checkDirection(trackMap []string) (int, int) {
	x := 0
	y := 0

	for i := 0; i < len(trackMap); i++ {
		for j := 0; j < len(trackMap[i]); j++ {

		}
	}

	return x, y
}

func updateCartsPositions(cartPositions [][]int, trackMap []string) [][]int {
	for i := 0; i < len(cartPositions); i++ {
		x := cartPositions[i][0]
		y := cartPositions[i][1]

		switch cartPositions[i][2] {
		//Direction: 0 = up, 1 = right, 2 = down, 3 = left

		case 0:
			switch string(trackMap[x-1][y]) {
			case "\\":
				cartPositions[i][2] = 3
			case "/":
				cartPositions[i][2] = 1
			case "+":
				switch cartPositions[i][3] {
				case 0:
					cartPositions
				case 1:
				case 2:
				}
			}
			cartPositions[i][0] = x + 1
		case 1:
			switch string(trackMap[x][y+1]) {
			case "-":
			case "\\":
			case "/":
			case "+":
			}
		case 2:
			switch string(trackMap[x+1][y]) {
			case "|":
			case "\\":
			case "/":
			case "+":
			}
		case 3:
			switch string(trackMap[x][y] - 1) {
			case "-":
			case "\\":
			case "/":
			case "+":
			}
		}
	}

	return cartPositions
}

func getCartPositions(trackMap []string) [][]int {
	cartPositions := [][]int{}
	for i := 0; i < len(trackMap); i++ {
		for j := 0; j < len(trackMap[i]); j++ {
			x := 0
			y := 0
			//Direction: 0 = up, 1 = right, 2 = down, 3 = left
			direction := 0
			switch string(trackMap[i][j]) {
			case "^":
				x = i
				y = j
				direction = 0
				cartPositions = append(cartPositions, []int{x, y, direction, 0})
				trackMap[i] = strings.Replace(trackMap[i], "^", "|", 1)
			case ">":
				x = i
				y = j
				direction = 1
				cartPositions = append(cartPositions, []int{x, y, direction, 0})
				trackMap[i] = strings.Replace(trackMap[i], ">", "-", 1)
			case "V":
				x = i
				y = j
				direction = 2
				cartPositions = append(cartPositions, []int{x, y, direction, 0})
				trackMap[i] = strings.Replace(trackMap[i], "V", "|", 1)
			case "<":
				x = i
				y = j
				direction = 3
				cartPositions = append(cartPositions, []int{x, y, direction, 0})
				trackMap[i] = strings.Replace(trackMap[i], "<", "-", 1)
			default:

			}
		}
	}
	return cartPositions
}

func colisionDetect(cartPositions [][]int) bool {
	for i := 0; i < len(cartPositions); i++ {
		for j := 0; j < len(cartPositions); j++ {
			if cartPositions[i][0] == cartPositions[j][0] && cartPositions[i][1] == cartPositions[j][1] {
				return true
			}
		}
	}
	return false
}

func main() {
	start := time.Now()

	trackMap := []string{}
	carts := [][]int{}

	crash := false

	intputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(intputFile)

	for scanner.Scan() {
		trackMap = append(trackMap, scanner.Text())
	}

	carts = getCartPositions(trackMap)

	for !crash {
		carts = updateCartsPositions(carts, trackMap)

		crash = colisionDetect(carts)
	}

	fmt.Println(carts)

	// for i := 0; i < len(trackMap); i++ {
	// 	fmt.Println(trackMap[i])
	// }

	fmt.Println(time.Since(start))
}
