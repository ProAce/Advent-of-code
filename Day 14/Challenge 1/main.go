package main

import (
	"fmt"
	"time"
)

func newRecipes(elf1, elf2 int, input []int) (int, int, []int) {
	// fmt.Println(elf1, elf2, input)
	temp := input[elf1] + input[elf2]
	if temp/10 == 0 {
		input = append(input, temp)
	} else {
		input = append(input, temp/10, temp%10)
	}

	temp = elf1 + input[elf1] + 1
	for temp > len(input)-1 {
		temp = temp - len(input)
	}
	elf1 = temp

	temp = elf2 + input[elf2] + 1
	for temp > len(input)-1 {
		temp = temp - len(input)
	}
	elf2 = temp

	return elf1, elf2, input
}

func main() {
	start := time.Now()

	iterations := 360781
	input := []int{3, 7}
	elf1 := 0
	elf2 := 1

	for len(input) < iterations+10 {
		elf1, elf2, input = newRecipes(elf1, elf2, input)
	}

	fmt.Println("")

	for i := len(input) - 10; i < len(input); i++ {
		fmt.Print(input[i])
	}

	fmt.Println("")

	fmt.Println(time.Since(start))
}
