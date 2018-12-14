package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func updateCartsPositions(cartPositions [][]int, trackMap []string) [][]int {
	for i := 0; i < len(cartPositions); i++ {
		x := cartPositions[i][0]
		y := cartPositions[i][1]

		switch cartPositions[i][2] {
		//Direction: 0 = up, 1 = right, 2 = down, 3 = left

		case 0:
			switch string(trackMap[x-1][y]) {
			case "\\": //left
				cartPositions[i][2] = 3
			case "/": //right
				cartPositions[i][2] = 1
			case "+":
				switch cartPositions[i][3] {
				case 0: //left
					cartPositions[i][3] = 1
					cartPositions[i][2] = 3
				case 1: //straight
					cartPositions[i][3] = 2
				case 2: //right
					cartPositions[i][3] = 0
					cartPositions[i][2] = 1
				}
			}
			cartPositions[i][0] = x - 1
		case 1:
			switch string(trackMap[x][y+1]) {
			case "\\": //down
				cartPositions[i][2] = 2
			case "/": //up
				cartPositions[i][2] = 0
			case "+":
				switch cartPositions[i][3] {
				case 0: //up
					cartPositions[i][3] = 1
					cartPositions[i][2] = 0
				case 1: //straight
					cartPositions[i][3] = 2
				case 2: //down
					cartPositions[i][3] = 0
					cartPositions[i][2] = 2
				}
			}
			cartPositions[i][1] = y + 1
		case 2:
			switch string(trackMap[x+1][y]) {
			case "\\": //right
				cartPositions[i][2] = 1
			case "/": //left
				cartPositions[i][2] = 3
			case "+":
				switch cartPositions[i][3] {
				case 0: //right
					cartPositions[i][3] = 1
					cartPositions[i][2] = 1
				case 1: //straight
					cartPositions[i][3] = 2
				case 2: //left
					cartPositions[i][3] = 0
					cartPositions[i][2] = 3
				}
			}
			cartPositions[i][0] = x + 1
		case 3:
			switch string(trackMap[x][y-1]) {
			case "\\": //up
				cartPositions[i][2] = 0
			case "/": //down
				cartPositions[i][2] = 2
			case "+":
				switch cartPositions[i][3] {
				case 0: //down
					cartPositions[i][3] = 1
					cartPositions[i][2] = 2
				case 1: //straight
					cartPositions[i][3] = 2
				case 2: //up
					cartPositions[i][3] = 0
					cartPositions[i][2] = 0
				}
			}
			cartPositions[i][1] = y - 1
		}
	}

	return cartPositions
}

func drawTracks(trackMap []string, carts [][]int) {
	for i := 0; i < len(trackMap[0]); i++ {
		fmt.Print(i / 100)
	}
	fmt.Println("")
	for i := 0; i < len(trackMap[0]); i++ {
		fmt.Print(i % 100 / 10)
	}
	fmt.Println("")
	for i := 0; i < len(trackMap[0]); i++ {
		fmt.Print(i % 10)
	}
	fmt.Println("")

	for i := 0; i < len(trackMap); i++ {
		for j := 0; j < len(trackMap[i]); j++ {
			print := true
			for k := 0; k < len(carts); k++ {
				if carts[k][0] == i && carts[k][1] == j {
					fmt.Print("#")
					print = false
				}
			}
			if print {
				fmt.Print(string(trackMap[i][j]))
			}

		}
		fmt.Println("")
	}
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
			case "v":
				x = i
				y = j
				direction = 2
				cartPositions = append(cartPositions, []int{x, y, direction, 0})
				trackMap[i] = strings.Replace(trackMap[i], "v", "|", 1)
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

func colisionDetect(cartPositions [][]int) (bool, []int) {
	for i := 0; i < len(cartPositions); i++ {
		for j := i + 1; j < len(cartPositions); j++ {
			if cartPositions[i][0] == cartPositions[j][0] && cartPositions[i][1] == cartPositions[j][1] {
				coordinates := []int{cartPositions[i][0], cartPositions[i][1]}
				return true, coordinates
			}
		}
	}
	return false, nil
}

func main() {
	start := time.Now()

	trackMap := []string{}
	carts := [][]int{}
	found := false

	var coordinates []int

	intputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(intputFile)

	for scanner.Scan() {
		trackMap = append(trackMap, scanner.Text())
	}

	carts = getCartPositions(trackMap)

	for !found {
		carts = updateCartsPositions(carts, trackMap)
		found, coordinates = colisionDetect(carts)
	}

	fmt.Println(coordinates[1], coordinates[0])

	fmt.Println(time.Since(start))
}
