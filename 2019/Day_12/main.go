package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type vector struct {
	x, y, z int
}

type moon struct {
	position, velocity vector
	energy             int
}

func (m *moon) applyVelocity() {
	m.position.x += m.velocity.x
	m.position.y += m.velocity.y
	m.position.z += m.velocity.z
}

func (m *moon) applyGravity(m2 moon) {
	if m.position.x < m2.position.x {
		m.velocity.x++
	} else if m.position.x > m2.position.x {
		m.velocity.x--
	}

	if m.position.y < m2.position.y {
		m.velocity.y++
	} else if m.position.y > m2.position.y {
		m.velocity.y--
	}

	if m.position.z < m2.position.z {
		m.velocity.z++
	} else if m.position.z > m2.position.z {
		m.velocity.z--
	}
}

func (m *moon) calculateEnergy() {
	m.energy = m.calculateKineticEnergy() * m.calculatePotentialEnergy()
}

func (m *moon) calculateKineticEnergy() int {
	return abs(m.velocity.x) + abs(m.velocity.y) + abs(m.velocity.z)
}

func (m *moon) calculatePotentialEnergy() int {
	return abs(m.position.x) + abs(m.position.y) + abs(m.position.z)
}

func main() {
	start := time.Now()

	moons := []moon{}

	inputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()

		if err != nil {
			log.Fatal(err)
		}

		line = strings.Replace(line, " ", "", -1) // Remove all spaces
		line = strings.Replace(line, "<", "", 1)  // Remove <
		line = strings.Replace(line, "x=", "", 1) // Remove x=
		line = strings.Replace(line, "y=", "", 1) // Remove y=
		line = strings.Replace(line, "z=", "", 1) // Remove z=
		line = strings.Replace(line, ">", "", 1)  // Remove >

		input := strings.Split(line, ",") // Split line into three "ints"

		x, _ := strconv.Atoi(input[0])
		y, _ := strconv.Atoi(input[1])
		z, _ := strconv.Atoi(input[2])

		moons = append(moons, moon{
			position: vector{x, y, z},
			velocity: vector{0, 0, 0},
		})
	}

	energy := partOne(1000, moons)

	fmt.Println(energy)

	steps := partTwo(moons)

	fmt.Println(steps)

	fmt.Println(time.Since(start))
}

func partOne(count int, m []moon) int {
	for i := 0; i < count; i++ {
		m = applyGravity(m)
		m = applyVelocity(m)
	}

	energy := 0
	for i := range m {
		m[i].calculateEnergy()
		energy += m[i].energy
	}

	return energy
}

func partTwo(m []moon) int {
	x := getPositions(m, 0)
	y := getPositions(m, 1)
	z := getPositions(m, 2)

	var periodX, periodY, periodZ, steps int

	for periodX == 0 || periodY == 0 || periodZ == 0 {
		m = applyGravity(m)
		m = applyVelocity(m)
		steps++

		coordinate := getPositions(m, 0)
		if coordinate == x {
			periodX = steps
		}

		coordinate = getPositions(m, 1)
		if coordinate == y {
			periodY = steps
		}

		coordinate = getPositions(m, 2)
		if coordinate == z {
			periodZ = steps
		}
	}

	lcmXY := (periodX * periodY) / gcd(periodX, periodY)
	lcm := (lcmXY * periodZ) / gcd(lcmXY, periodZ)

	return lcm
}

func getPositions(m []moon, xyz int) (output [8]int) {
	for i, val := range m {
		switch xyz {
		case 0:
			output[i*2] = val.position.x
			output[i*2+1] = val.velocity.x
			break
		case 1:
			output[i*2] = val.position.y
			output[i*2+1] = val.velocity.y
			break
		case 2:
			output[i*2] = val.position.z
			output[i*2+1] = val.velocity.z
			break
		}
	}

	return output
}

func applyGravity(m []moon) []moon {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			if i == j {
				continue
			}
			m[i].applyGravity(m[j])
		}
	}

	return m
}

func applyVelocity(m []moon) []moon {
	for i := 0; i < len(m); i++ {
		m[i].applyVelocity()
	}

	return m
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
