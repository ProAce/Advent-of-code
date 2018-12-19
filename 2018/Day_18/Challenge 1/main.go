package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func parseAcres(lumber []string) (output [][]byte) {
	b := []byte{}

	for i := 0; i < len(lumber[0])+2; i++ {
		b = append(b, 9)
	}

	output = append(output, b)

	for i := 0; i < len(lumber); i++ {
		output = append(output, []byte{})
	}

	for i := 0; i < len(lumber); i++ {
		output[i+1] = append(output[i+1], 9)
		for j := 0; j < len(lumber[i]); j++ {
			switch string(lumber[i][j]) {
			case ".":
				output[i+1] = append(output[i+1], 0)
			case "#":
				output[i+1] = append(output[i+1], 1)
			case "|":
				output[i+1] = append(output[i+1], 2)
			}
		}
		output[i+1] = append(output[i+1], 9)
	}

	output = append(output, b)

	return output
}

func changeAcres(input [][]byte) [][]byte {
	output := input

	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[i])-1; j++ {

			switch input[i][j] {
			case 0: //open
				count := 0

				for k := i - 1; k <= i+1; k++ {
					for l := j - 1; l <= j+1; l++ {
						if input[k][l] == 2 {
							count++
						}
					}
				}

				if count >= 3 {
					output[i][j] = 2
				}

			case 1: //lumberyards
				countLumber := 0
				countTrees := 0

				for k := i - 1; k <= i+1; k++ {
					for l := j - 1; l <= j+1; l++ {

						if k != i && l != j {
							if input[k][l] == 1 {
								countLumber++
							}
							if input[k][l] == 2 {
								countTrees++
							}
						}
					}
				}

				if countLumber < 1 || countTrees < 1 {
					output[i][j] = 0
				}

			case 2: //trees
				count := 0

				for k := i - 1; k <= i+1; k++ {
					for l := j - 1; l <= j+1; l++ {
						if input[k][l] == 1 {
							count++
						}
					}
				}

				if count >= 3 {
					output[i][j] = 1
				}
			}

		}

	}
	return output
}

func printInput(input [][]byte) {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			switch input[i][j] {
			case 0:
				fmt.Print(".")
			case 1:
				fmt.Print("#")
			case 2:
				fmt.Print("|")
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func countAcres(input [][]byte) int {
	yards := 0
	wood := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			switch input[i][j] {
			case 1:
				yards++
			case 2:
				wood++
			}
		}
	}

	return yards * wood
}

func main() {
	start := time.Now()

	lumber := []string{}

	intputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer intputFile.Close()

	scanner := bufio.NewScanner(intputFile)

	for scanner.Scan() {
		lumber = append(lumber, scanner.Text())
	}

	input := parseAcres(lumber)
	printInput(input)

	// for i := 0; i < 2; i++ {
	input = changeAcres(input)
	printInput(input)
	// }

	fmt.Println(countAcres(input))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(time.Since(start))
}
