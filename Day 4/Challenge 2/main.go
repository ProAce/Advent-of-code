package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var data []string
var guards [][]int
var guardMins [][]int

func parseData(input string) (t time.Time, id int, asleep bool) {
	input = strings.TrimPrefix(input, "[")
	inputTime := strings.Split(input, "]")
	timeString := string(inputTime[0])

	format := "2006-01-02 15:04"

	split := strings.Split(inputTime[1], " ")

	if split[2] == "up" {
		asleep = false
	} else if split[2] == "asleep" {
		asleep = true
	} else {
		split[2] = strings.TrimPrefix(split[2], "#")
		id, _ = strconv.Atoi(split[2])
	}

	t, _ = time.Parse(format, timeString)
	return t, id, asleep
}

func countMinsAsleep(id int) {
	b := make([]int, 61)
	b[0] = id

	for i := 0; i < len(guards); i++ {
		if guards[i][0] == id {
			for j := 0; j < 60; j++ {
				if guards[i][j+1] == 1 {
					b[j+1]++
				}
			}
		}
	}

	guardMins = append(guardMins, b)
}

func main() {
	start := time.Now()

	intputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(intputFile)

	for scanner.Scan() {
		input := scanner.Text()
		data = append(data, input)
	}

	sort.Strings(data)

	guardID := 0
	startSleep := 0

	for i := 0; i < len(data); i++ {
		t, id, asleep := parseData(data[i])

		if id != 0 {
			guardID = id
		}
		if (asleep == true) && (id == 0) {
			startSleep = t.Minute()
		}
		if (asleep == false) && (id == 0) {
			endSleep := t.Minute()
			b := make([]int, 61)
			b[0] = guardID
			for j := 0; j < 60; j++ {
				if (j >= startSleep) && (j < endSleep) {
					b[j+1] = 1
				} else {
					b[j+1] = 0
				}
			}
			guards = append(guards, b)
		}
	}

	//Make list of all the guardID's
	b := []int{}

	for i := 0; i < len(guards); i++ {
		add := true
		for j := 0; j < len(b); j++ {
			if guards[i][0] == b[j] {
				add = false
			}
		}
		if add == true {
			b = append(b, guards[i][0])
		}
	}

	for i := 0; i < len(b); i++ {
		countMinsAsleep(b[i])
	}

	min := 0
	count := 0
	id := 0

	for i := 0; i < len(guardMins); i++ {
		for j := 0; j < 60; j++ {
			if guardMins[i][j+1] > count {
				id = guardMins[i][0]
				min = j
				count = guardMins[i][j+1]
			}
		}
	}

	output := id * min

	fmt.Println(id, count, min, output)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(time.Since(start))
}
