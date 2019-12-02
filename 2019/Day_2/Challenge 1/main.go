package main

import (
	"bufio"
	"fmt"
	"log"
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

	for scanner.Scan() {
		line := scanner.Text()

		opcodeString := strings.Split(line, ",")
		opcode := []int{}
		for _, codes := range opcodeString {
			i, _ := strconv.Atoi(codes)
			opcode = append(opcode, i)
		}

		for i := 0; i < len(opcode); i += 4 {

			if opcode[i] == 1 {
				opcode[opcode[i+3]] = opcode[opcode[i+1]] + opcode[opcode[i+2]]
			} else if opcode[i] == 2 {
				opcode[opcode[i+3]] = opcode[opcode[i+1]] * opcode[opcode[i+2]]
			} else if opcode[i] == 99 {
				break
			}
		}
		fmt.Println(opcode[0])
	}

	fmt.Println(time.Since(start))
}
