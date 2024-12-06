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

type GridCell struct {
	Char rune
	Seen bool
}

func runExercise1() (int, error) {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	input := strings.TrimSpace(string(fileBytes))
	grid := make([][]GridCell, 0)
	guardI, guardJ := -1, -1
	for i, row := range strings.Split(input, "\n") {
		gridRow := make([]GridCell, 0)
		for j, char := range row {
			if char == '^' {
				gridRow = append(gridRow, GridCell{Char: '.'})
				guardI, guardJ = i, j
				continue
			}
			gridRow = append(gridRow, GridCell{Char: char})
		}
		grid = append(grid, gridRow)
	}
	stepI, stepJ := -1, 0 // Point up to start
	numBlocks := 1        // Starting block counts
	for {
		nextI, nextJ := guardI+stepI, guardJ+stepJ
		// Out of bounds
		if nextI < 0 || nextJ < 0 || nextI >= len(grid) || nextJ >= len(grid[guardI]) {
			break
		}
		// Collision
		if grid[nextI][nextJ].Char == '#' {
			stepI, stepJ = stepJ, stepI*-1 // Rotate 90 degrees clockwise
			continue
		}
		// Move forward
		guardI, guardJ = nextI, nextJ
		if !grid[guardI][guardJ].Seen {
			numBlocks++
		}
		grid[guardI][guardJ].Seen = true
	}
	return numBlocks, nil
}

type CycleGridCell struct {
	Char rune
	Seen map[byte]struct{} // 00-u, 01-r, 02-d, 03-l
}

func runExercise2() (int, error) {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	input := strings.TrimSpace(string(fileBytes))
	grid := make([][]CycleGridCell, 0)
	originalGuardI, originalGuardJ := -1, -1
	for i, row := range strings.Split(input, "\n") {
		gridRow := make([]CycleGridCell, 0)
		for j, char := range row {
			if char == '^' {
				gridRow = append(gridRow, CycleGridCell{Char: '.', Seen: make(map[byte]struct{})})
				originalGuardI, originalGuardJ = i, j
				continue
			}
			gridRow = append(gridRow, CycleGridCell{Char: char, Seen: make(map[byte]struct{})})
		}
		grid = append(grid, gridRow)
	}
	numSolutions := 0
	for obstacleI := 0; obstacleI < len(grid); obstacleI++ {
		for obstacleJ := 0; obstacleJ < len(grid[0]); obstacleJ++ {
			stepI, stepJ := -1, 0 // Point up to start
			for i := range grid {
				for j := range grid[0] {
					grid[i][j].Seen = make(map[byte]struct{})
				}
			}
			guardI, guardJ := originalGuardI, originalGuardJ
			for {
				nextI, nextJ := guardI+stepI, guardJ+stepJ
				// Out of bounds
				if nextI < 0 || nextJ < 0 || nextI >= len(grid) || nextJ >= len(grid[guardI]) {
					break
				}
				// Collision
				if grid[nextI][nextJ].Char == '#' || (nextI == obstacleI && nextJ == obstacleJ) {
					stepI, stepJ = stepJ, stepI*-1 // Rotate 90 degrees clockwise
					continue
				}
				// Move forward
				guardI, guardJ = nextI, nextJ
				if _, ok := grid[guardI][guardJ].Seen[getDirection(stepI, stepJ)]; ok {
					numSolutions++
					break
				}
				grid[guardI][guardJ].Seen[getDirection(stepI, stepJ)] = struct{}{}
			}
		}
	}
	return numSolutions, nil
}

func getDirection(stepI, stepJ int) byte {
	if stepI == -1 && stepJ == 0 {
		return 0 // u
	}
	if stepI == 0 && stepJ == 1 {
		return 1 // r
	}
	if stepI == 1 && stepJ == 0 {
		return 2 // d
	}
	return 3 // l
}
