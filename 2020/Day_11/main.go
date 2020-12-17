package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var occupied = '#'
var empty = 'L'
var floor = '.'

func main() {
	start := time.Now()

	input := getInput("input.txt")

	part1 := part1(input)
	// part2 := part2(input)

	fmt.Printf("Part 1: %d\r\n", part1)
	// fmt.Printf("Part 2: %d\r\n", part2)

	fmt.Println(time.Since(start))
}

func part1(input *playingField) int {
	for {
		input = iteratePart1(input)

		if !input.changed {
			return count(input.seats, occupied)
		}
	}
}

func part2(input *playingField) int {
	for {
		input = iteratePart2(input)

		if !input.changed {
			return count(input.seats, occupied)
		}
	}
}

type playingField struct {
	width   int
	seats   []rune
	changed bool
}

func (p *playingField) Print() {
	for i, seat := range p.seats {
		if i%p.width == 0 {
			fmt.Println()
		}
		fmt.Print(string(seat))
	}
	fmt.Println()
}

func iteratePart1(p *playingField) *playingField {
	newPlayingfield := &playingField{
		width:   p.width,
		seats:   make([]rune, len(p.seats)),
		changed: false,
	}

	for i, currentSeat := range p.seats {
		neighbours := findNeighbours(p, i)
		count := count(neighbours, occupied)

		if currentSeat == empty && count == 0 {
			newPlayingfield.changed = true
			newPlayingfield.seats[i] = occupied
		} else if currentSeat == occupied && count >= 4 {
			newPlayingfield.changed = true
			newPlayingfield.seats[i] = empty
		} else {
			newPlayingfield.seats[i] = currentSeat
		}
	}

	return newPlayingfield
}

func iteratePart2(p *playingField) *playingField {
	newPlayingfield := &playingField{
		width:   p.width,
		seats:   make([]rune, len(p.seats)),
		changed: false,
	}

	for i, currentSeat := range p.seats {
		neighbours := findDirectionNeighbours(p, i)
		count := count(neighbours, occupied)

		if currentSeat == empty && count == 0 {
			newPlayingfield.changed = true
			newPlayingfield.seats[i] = occupied
		} else if currentSeat == occupied && count >= 5 {
			newPlayingfield.changed = true
			newPlayingfield.seats[i] = empty
		} else {
			newPlayingfield.seats[i] = currentSeat
		}
	}

	return newPlayingfield
}

func count(seats []rune, check rune) int {
	count := 0

	for _, seat := range seats {
		if seat == check {
			count++
		}
	}

	return count
}

func findNeighbours(p *playingField, index int) []rune {
	neighbours := []rune{}
	indexes := []int{}

	mask := []struct {
		row int
		col int
	}{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	relativeIndex := index % p.width

	for _, m := range mask {
		rowMod := m.row * p.width
		neighbourIndex := index - m.col + rowMod
		relativeNeighbourIndex := relativeIndex - m.col

		if neighbourIndex < 0 || neighbourIndex > len(p.seats)-1 {
			continue
		}

		if relativeNeighbourIndex < 0 || relativeNeighbourIndex >= p.width {
			continue
		}

		indexes = append(indexes, neighbourIndex)
	}

	for _, i := range indexes {
		if i == index {
			continue
		}
		neighbours = append(neighbours, p.seats[i])
	}

	return neighbours
}

func findDirectionNeighbours(p *playingField, index int) []rune {
	neighbours := []rune{}

	return neighbours
}

func getInput(path string) *playingField {
	inputFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	newField := &playingField{
		seats: []rune{},
	}

	count := 0

	// Get the input values from the text file and store them in a slice.
	for scanner.Scan() {
		line := scanner.Text()

		for _, seat := range line {
			newField.seats = append(newField.seats, seat)
		}
		count++
	}

	newField.width = len(newField.seats) / count

	return newField
}
