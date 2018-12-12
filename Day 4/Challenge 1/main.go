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

func main() {
	start := time.Now()

	data := []string{}
	guards := []int{}

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
	// for i := 0; i < len(data); i++ {
	// 	fmt.Println(data[i])
	// }

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
			b = append(b, guardID)
			for j := 0; j < 60; j++ {
				if (j >= startSleep) && (j < endSleep) {
					guard[][j] = 1
				} else {
					guard[][j] = 0
				}
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(time.Since(start))
}
