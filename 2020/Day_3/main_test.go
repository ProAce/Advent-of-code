package main

import "testing"

func TestPart1(t *testing.T) {
	sample := getInput("sample.txt")

	outcome := part1(sample)

	if outcome != 7 {
		t.Errorf("Number of trees was incorrect, got: %d, want: %d.", outcome, 7)
	}
}

func TestPart2(t *testing.T) {
	sample := getInput("sample.txt")

	outcome := part2(sample)

	if outcome != 336 {
		t.Errorf("Number of trees was incorrect, got: %d, want: %d.", outcome, 336)
	}
}
