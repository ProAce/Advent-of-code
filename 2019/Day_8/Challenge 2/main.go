package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strconv"
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

	width := 25
	height := 6

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	white := color.White
	black := color.Black

	for scanner.Scan() {

		line := scanner.Text()

		imageLayerLength := width * height

		pixels := make(map[int][]int)

		for i := 0; i < len(line); i += imageLayerLength {
			for j := 0; j < imageLayerLength; j++ {
				count, _ := strconv.Atoi(string(line[i+j]))
				pixels[j] = append(pixels[j], count)
			}
		}

		for key, value := range pixels {
			x := key % width
			y := key / width
			for _, color := range value {
				if color == 1 {
					img.Set(x, y, white)
					break
				}
				if color == 0 {
					img.Set(x, y, black)
					break
				}
			}
		}

		f, _ := os.Create("SolutionDay8Part2.png")
		png.Encode(f, img)
	}

	fmt.Println(time.Since(start))
}
