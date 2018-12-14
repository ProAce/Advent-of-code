package main

import (
	"fmt"
	"time"
)

var toFind = []int{3, 6, 0, 7, 8, 1}

func newRecipes(elf1, elf2 int, input []int) (bool, int, int, []int) {
	// fmt.Println(elf1, elf2, input)
	temp := input[elf1] + input[elf2]
	twice := false
	if temp/10 == 0 {
		input = append(input, temp)
	} else {
		input = append(input, temp/10, temp%10)
		twice = true
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

	return twice, elf1, elf2, input
}

func checkOne(input []int) bool {
	toCheck := input[len(input)-len(toFind):]

	for i, v := range toCheck {
		if v != toFind[i] {
			return false
		}
	}

	return true
}

func checkTwo(input []int) bool {
	toCheck := input[len(input)-len(toFind)-1 : len(input)-1]

	for i, v := range toCheck {
		if v != toFind[i] {
			return false
		}
	}

	return true
}

func main() {
	start := time.Now()

	found := false
	input := []int{3, 7}
	elf1 := 0
	elf2 := 1
	twice := false

	for !found {
		twice, elf1, elf2, input = newRecipes(elf1, elf2, input)
		if len(input) > len(toFind) {
			found = checkOne(input)
			if twice == true {
				found = checkTwo(input)
			}
		}
	}

	if twice == false {
		for i := len(input) - len(toFind); i < len(input); i++ {
			fmt.Print(input[i])
		}

		fmt.Println("")
		fmt.Println(len(input) - len(toFind))
	} else {
		for i := len(input) - len(toFind) - 1; i < len(input)-1; i++ {
			fmt.Print(input[i])
		}

		fmt.Println("")
		fmt.Println(len(input) - len(toFind) - 1)
	}

	fmt.Println(time.Since(start))
}
