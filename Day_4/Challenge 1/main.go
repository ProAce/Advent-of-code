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

func countOnes(id int) (mins int) {
	for i := 0; i < len(guards); i++ {
		if guards[i][0] == id {
			for j := 1; j < 61; j++ {
				if guards[i][j] == 1 {
					mins++
				}
			}
		}
	}

	return mins
}

func whichMin(id int) (min int) {
	b := make([]int, 60)

	for i := 0; i < len(guards); i++ {
		if guards[i][0] == id {
			for j := 1; j < 61; j++ {
				if guards[i][j] == 1 {
					b[j-1]++
				}
			}
		}
	}

	count := 0

	for i := 0; i < len(b); i++ {
		if b[i] > count {
			count = b[i]
			min = i
		}
	}

	return min
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

	id := 0
	mostMins := 0

	for i := 0; i < len(b); i++ {
		mins := countOnes(b[i])
		if mins > mostMins {
			mostMins = mins
			id = b[i]
		}
	}

	minute := whichMin(id)

	output := minute * id

	fmt.Println(id, mostMins, minute, output)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(time.Since(start))
}
