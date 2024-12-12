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
	Value rune
	I     int
	J     int
}

func runExercise1() (int, error) {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	rows := strings.Split(strings.TrimSpace(string(fileBytes)), "\n")
	grid := make([][]rune, 0, len(rows))
	for _, row := range rows {
		gridRow := make([]rune, 0, len(row))
		for _, char := range row {
			gridRow = append(gridRow, char)
		}
		grid = append(grid, gridRow)
	}
	seen := make(map[Point]struct{})
	score := 0
	for startI := range grid {
		for startJ := range grid[startI] {
			startingPoint := Point{Value: grid[startI][startJ], I: startI, J: startJ}
			if _, ok := seen[startingPoint]; ok {
				continue
			}
			frontier := []Point{startingPoint}
			frontierSeen := map[Point]struct{}{startingPoint: struct{}{}}
			area, perimiter := 0, 0
			for len(frontier) > 0 {
				point := frontier[0]
				frontier = frontier[1:]
				pointValue := point.Value
				area++
				seen[point] = struct{}{}
				if point.I > 0 && grid[point.I-1][point.J] == pointValue {
					p := Point{Value: grid[point.I-1][point.J], I: point.I - 1, J: point.J}
					if _, ok := frontierSeen[p]; !ok {
						frontier = append(frontier, p)
						frontierSeen[p] = struct{}{}
					}
				} else {
					perimiter++
				}
				if point.J < len(grid[point.I])-1 && grid[point.I][point.J+1] == pointValue {
					p := Point{Value: grid[point.I][point.J+1], I: point.I, J: point.J + 1}
					if _, ok := frontierSeen[p]; !ok {
						frontier = append(frontier, p)
						frontierSeen[p] = struct{}{}
					}
				} else {
					perimiter++
				}
				if point.I < len(grid)-1 && grid[point.I+1][point.J] == pointValue {
					p := Point{Value: grid[point.I+1][point.J], I: point.I + 1, J: point.J}
					if _, ok := frontierSeen[p]; !ok {
						frontier = append(frontier, p)
						frontierSeen[p] = struct{}{}
					}
				} else {
					perimiter++
				}
				if point.J > 0 && grid[point.I][point.J-1] == pointValue {
					p := Point{Value: grid[point.I][point.J-1], I: point.I, J: point.J - 1}
					if _, ok := frontierSeen[p]; !ok {
						frontier = append(frontier, p)
						frontierSeen[p] = struct{}{}
					}
				} else {
					perimiter++
				}
			}
			score += area * perimiter
		}
	}
	return score, nil
}

func runExercise2() (int, error) {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	rows := strings.Split(strings.TrimSpace(string(fileBytes)), "\n")
	grid := make([][]rune, 0, len(rows))
	for _, row := range rows {
		gridRow := make([]rune, 0, len(row))
		for _, char := range row {
			gridRow = append(gridRow, char)
		}
		grid = append(grid, gridRow)
	}
	seen := make(map[Point]struct{})
	score := 0
	for startI := range grid {
		for startJ := range grid[startI] {
			startingPoint := Point{Value: grid[startI][startJ], I: startI, J: startJ}
			if _, ok := seen[startingPoint]; ok {
				continue
			}
			frontier := []Point{startingPoint}
			frontierSeen := map[Point]struct{}{startingPoint: struct{}{}}
			area, corners := 0, 0
			for len(frontier) > 0 {
				point := frontier[0]
				frontier = frontier[1:]
				pointValue := point.Value
				area++
				seen[point] = struct{}{}
				matchUp, matchRight, matchDown, matchLeft := false, false, false, false
				if point.I > 0 && grid[point.I-1][point.J] == pointValue {
					matchUp = true
					p := Point{Value: grid[point.I-1][point.J], I: point.I - 1, J: point.J}
					if _, ok := frontierSeen[p]; !ok {
						frontier = append(frontier, p)
						frontierSeen[p] = struct{}{}
					}
				}
				if point.J < len(grid[point.I])-1 && grid[point.I][point.J+1] == pointValue {
					matchRight = true
					p := Point{Value: grid[point.I][point.J+1], I: point.I, J: point.J + 1}
					if _, ok := frontierSeen[p]; !ok {
						frontier = append(frontier, p)
						frontierSeen[p] = struct{}{}
					}
				}
				if point.I < len(grid)-1 && grid[point.I+1][point.J] == pointValue {
					matchDown = true
					p := Point{Value: grid[point.I+1][point.J], I: point.I + 1, J: point.J}
					if _, ok := frontierSeen[p]; !ok {
						frontier = append(frontier, p)
						frontierSeen[p] = struct{}{}
					}
				}
				if point.J > 0 && grid[point.I][point.J-1] == pointValue {
					matchLeft = true
					p := Point{Value: grid[point.I][point.J-1], I: point.I, J: point.J - 1}
					if _, ok := frontierSeen[p]; !ok {
						frontier = append(frontier, p)
						frontierSeen[p] = struct{}{}
					}
				}
				if !matchUp && !matchRight {
					corners++
				}
				if !matchUp && !matchLeft {
					corners++
				}
				if !matchDown && !matchRight {
					corners++
				}
				if !matchDown && !matchLeft {
					corners++
				}
				if matchUp && matchRight && grid[point.I-1][point.J+1] != pointValue {
					corners++
				}
				if matchUp && matchLeft && grid[point.I-1][point.J-1] != pointValue {
					corners++
				}
				if matchDown && matchRight && grid[point.I+1][point.J+1] != pointValue {
					corners++
				}
				if matchDown && matchLeft && grid[point.I+1][point.J-1] != pointValue {
					corners++
				}
			}
			score += area * corners
		}
	}
	return score, nil
}
