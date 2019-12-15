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

type baseComponent struct {
	ores   int
	amount int
	needed int
}

type component struct {
	amount   int
	reaction map[string]int
}

type nanoFactory struct {
	reactions     map[string]component
	baseComponent map[string]baseComponent
}

func (n *nanoFactory) parseInput(input string) {
	splitInput := strings.Split(input, "=>")

	componentString := strings.Split(splitInput[1], " ")
	componentKey := componentString[2]
	componentVal, _ := strconv.Atoi(componentString[1])

	inputString := strings.Split(splitInput[0], ",")

	value := make(map[string]int)

	for i := range inputString {
		inputStringSplit := strings.Split(inputString[i], " ")

		inputKey := ""
		inputValue := 0

		if i > 0 {
			inputKey = inputStringSplit[2]
			inputValue, _ = strconv.Atoi(inputStringSplit[1])
		} else {
			inputKey = inputStringSplit[1]
			inputValue, _ = strconv.Atoi(inputStringSplit[0])
		}

		value[inputKey] = inputValue
	}

	if _, ok := value["ORE"]; ok {
		n.baseComponent[componentKey] = baseComponent{
			ores:   value["ORE"],
			amount: componentVal,
		}
	} else {
		n.reactions[componentKey] = component{
			reaction: value,
			amount:   componentVal,
		}
	}
}

func (n *nanoFactory) createFuel() (oreCount int) {
	for i, value := range n.reactions["FUEL"].reaction {
		if _, ok := n.baseComponent[i]; ok {
			n.baseComponent[i].needed += value
		}
	}

	return oreCount
}

func main() {
	start := time.Now()

	factory := nanoFactory{
		reactions:     make(map[string]component),
		baseComponent: make(map[string]baseComponent),
	}

	inputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		factory.parseInput(line)
	}

	fmt.Println(factory)

	fmt.Println(time.Since(start))
}
