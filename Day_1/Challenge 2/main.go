package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	frequency := 0
	var outputRec []int

	for {
		intputFile, err := os.Open("input.txt")

		if err != nil {
			log.Fatal(err)
		}

		defer intputFile.Close()

		scanner := bufio.NewScanner(intputFile)

		for scanner.Scan() {
			line := scanner.Text()

			v, err := strconv.Atoi(line)

			if err != nil {
				log.Fatal(err)
			}

			frequency += v

			for i := 0; i < len(outputRec); i++ {
				if frequency == outputRec[i] {
					fmt.Println(frequency)
					fmt.Println(time.Since(start))
					return
				}
			}

			outputRec = append(outputRec, frequency)
		}
	}
}
