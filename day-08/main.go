package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	answer1, err := runExercise1()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Exercise 1: %d\n", answer1)
	answer2, err := runExercise2()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Exercise 2: %d\n", answer2)
}

type Point struct {
	I int
	J int
}

func runExercise1() (int, error) {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	rows := strings.Split(strings.TrimSpace(string(fileBytes)), "\n")
	iMax := len(rows) - 1
	jMax := len(rows[0]) - 1
	antennas := make(map[rune][]Point)
	for i, row := range rows {
		for j, cell := range row {
			if cell == '.' {
				continue
			}
			if _, ok := antennas[cell]; !ok {
				antennas[cell] = make([]Point, 0)
			}
			antennas[cell] = append(antennas[cell], Point{I: i, J: j})
		}
	}
	seen := make(map[int]map[int]struct{})
	count := 0
	for _, points := range antennas {
		for idx, pointA := range points {
			for _, pointB := range points[idx+1:] {
				iDiff, jDiff := pointA.I-pointB.I, pointA.J-pointB.J
				pointAAntinodeI, pointAAntinodeJ := pointA.I+iDiff, pointA.J+jDiff
				if isInBounds(pointAAntinodeI, pointAAntinodeJ, iMax, jMax) && setNewAntinode(seen, pointAAntinodeI, pointAAntinodeJ) {
					count++
				}
				pointBAntinodeI, pointBAntinodeJ := pointB.I-iDiff, pointB.J-jDiff
				if isInBounds(pointBAntinodeI, pointBAntinodeJ, iMax, jMax) && setNewAntinode(seen, pointBAntinodeI, pointBAntinodeJ) {
					count++
				}
			}
		}
	}
	return count, nil
}

func runExercise2() (int, error) {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	rows := strings.Split(strings.TrimSpace(string(fileBytes)), "\n")
	iMax := len(rows) - 1
	jMax := len(rows[0]) - 1
	antennas := make(map[rune][]Point)
	for i, row := range rows {
		for j, cell := range row {
			if cell == '.' {
				continue
			}
			if _, ok := antennas[cell]; !ok {
				antennas[cell] = make([]Point, 0)
			}
			antennas[cell] = append(antennas[cell], Point{I: i, J: j})
		}
	}
	seen := make(map[int]map[int]struct{})
	count := 0
	for _, points := range antennas {
		for idx, pointA := range points {
			for _, pointB := range points[idx+1:] {
				iDiff, jDiff := pointA.I-pointB.I, pointA.J-pointB.J
				nextAI, nextAJ := pointA.I, pointA.J
				for isInBounds(nextAI, nextAJ, iMax, jMax) {
					if setNewAntinode(seen, nextAI, nextAJ) {
						count++
					}
					nextAI, nextAJ = nextAI+iDiff, nextAJ+jDiff
				}
				nextBI, nextBJ := pointB.I, pointB.J
				for isInBounds(nextBI, nextBJ, iMax, jMax) {
					if setNewAntinode(seen, nextBI, nextBJ) {
						count++
					}
					nextBI, nextBJ = nextBI-iDiff, nextBJ-jDiff
				}
			}
		}
	}
	return count, nil
}

func setNewAntinode(seen map[int]map[int]struct{}, i, j int) bool {
	if seenI, ok := seen[i]; ok {
		if _, ok := seenI[j]; ok {
			return false
		}
	} else {
		seen[i] = make(map[int]struct{})
	}
	seen[i][j] = struct{}{}
	return true
}

func isInBounds(i, j, iMax, jMax int) bool {
	return i >= 0 && i <= iMax && j >= 0 && j <= jMax
}
