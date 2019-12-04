package main

import (
	"fmt"
)

func main(){
	counter := 0
	for i := 372304; i <= 847060; i++{ // Puzzle input
		if criteriaCheck(i) {
			counter++
		}
	}
	fmt.Println(counter)
}

func criteriaCheck(i int) bool{
	digits := []int{}
	for i > 0 {
		digits = append(digits, (i % 10))
		i /= 10
	}

	same := false
	
	for j := 0; j < len(digits) - 1; j++{
		if digits[j] < digits[j + 1]{
			return false
		}
		if digits[j] == digits[j + 1]{ 
			same = true;
		}
	}

	return same;
}