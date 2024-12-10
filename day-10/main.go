package main

import (
	"fmt"
	"os"
	"strconv"
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
	Value int
	I     int
	J     int
}

func runExercise1() (int, error) {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	grid := make([][]int, 0)
	for _, rawRow := range strings.Split(strings.TrimSpace(string(fileBytes)), "\n") {
		row := make([]int, 0, len(rawRow))
		for _, char := range rawRow {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return 0, err
			}
			row = append(row, num)
		}
		grid = append(grid, row)
	}
	numAnswers := 0
	for startI := range grid {
		for startJ := range grid[startI] {
			if grid[startI][startJ] != 0 {
				continue
			}
			stack := []Point{Point{Value: grid[startI][startJ], I: startI, J: startJ}}
			seen := make(map[Point]struct{})
			for len(stack) > 0 {
				currentPoint := stack[0]
				stack = stack[1:]
				if currentPoint.I > 0 && grid[currentPoint.I-1][currentPoint.J] == currentPoint.Value+1 {
					upPoint := Point{Value: grid[currentPoint.I-1][currentPoint.J], I: currentPoint.I - 1, J: currentPoint.J}
					if _, ok := seen[upPoint]; !ok {
						seen[upPoint] = struct{}{}
						if upPoint.Value == 9 {
							numAnswers++
						} else {
							stack = append(stack, upPoint)
						}
					}
				}
				if currentPoint.J < len(grid[startI])-1 && grid[currentPoint.I][currentPoint.J+1] == currentPoint.Value+1 {
					rightPoint := Point{Value: grid[currentPoint.I][currentPoint.J+1], I: currentPoint.I, J: currentPoint.J + 1}
					if _, ok := seen[rightPoint]; !ok {
						seen[rightPoint] = struct{}{}
						if rightPoint.Value == 9 {
							numAnswers++
						} else {
							stack = append(stack, rightPoint)
						}
					}
				}
				if currentPoint.I < len(grid)-1 && grid[currentPoint.I+1][currentPoint.J] == currentPoint.Value+1 {
					downPoint := Point{Value: grid[currentPoint.I+1][currentPoint.J], I: currentPoint.I + 1, J: currentPoint.J}
					if _, ok := seen[downPoint]; !ok {
						seen[downPoint] = struct{}{}
						if downPoint.Value == 9 {
							numAnswers++
						} else {
							stack = append(stack, downPoint)
						}
					}
				}
				if currentPoint.J > 0 && grid[currentPoint.I][currentPoint.J-1] == currentPoint.Value+1 {
					leftPoint := Point{Value: grid[currentPoint.I][currentPoint.J-1], I: currentPoint.I, J: currentPoint.J - 1}
					if _, ok := seen[leftPoint]; !ok {
						seen[leftPoint] = struct{}{}
						if leftPoint.Value == 9 {
							numAnswers++
						} else {
							stack = append(stack, leftPoint)
						}
					}
				}
			}
		}
	}
	return numAnswers, nil
}

func runExercise2() (int, error) {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	grid := make([][]int, 0)
	for _, rawRow := range strings.Split(strings.TrimSpace(string(fileBytes)), "\n") {
		row := make([]int, 0, len(rawRow))
		for _, char := range rawRow {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return 0, err
			}
			row = append(row, num)
		}
		grid = append(grid, row)
	}
	numAnswers := 0
	for startI := range grid {
		for startJ := range grid[startI] {
			if grid[startI][startJ] != 0 {
				continue
			}
			stack := []Point{Point{Value: grid[startI][startJ], I: startI, J: startJ}}
			seen := make(map[Point]struct{})
			for len(stack) > 0 {
				currentPoint := stack[0]
				stack = stack[1:]
				if currentPoint.I > 0 && grid[currentPoint.I-1][currentPoint.J] == currentPoint.Value+1 {
					upPoint := Point{Value: grid[currentPoint.I-1][currentPoint.J], I: currentPoint.I - 1, J: currentPoint.J}
					if _, ok := seen[upPoint]; !ok {
						seen[upPoint] = struct{}{}
					}
					if upPoint.Value == 9 {
						numAnswers++
					} else {
						stack = append(stack, upPoint)
					}
				}
				if currentPoint.J < len(grid[startI])-1 && grid[currentPoint.I][currentPoint.J+1] == currentPoint.Value+1 {
					rightPoint := Point{Value: grid[currentPoint.I][currentPoint.J+1], I: currentPoint.I, J: currentPoint.J + 1}
					if _, ok := seen[rightPoint]; !ok {
						seen[rightPoint] = struct{}{}
					}
					if rightPoint.Value == 9 {
						numAnswers++
					} else {
						stack = append(stack, rightPoint)
					}
				}
				if currentPoint.I < len(grid)-1 && grid[currentPoint.I+1][currentPoint.J] == currentPoint.Value+1 {
					downPoint := Point{Value: grid[currentPoint.I+1][currentPoint.J], I: currentPoint.I + 1, J: currentPoint.J}
					if _, ok := seen[downPoint]; !ok {
						seen[downPoint] = struct{}{}
					}
					if downPoint.Value == 9 {
						numAnswers++
					} else {
						stack = append(stack, downPoint)
					}
				}
				if currentPoint.J > 0 && grid[currentPoint.I][currentPoint.J-1] == currentPoint.Value+1 {
					leftPoint := Point{Value: grid[currentPoint.I][currentPoint.J-1], I: currentPoint.I, J: currentPoint.J - 1}
					if _, ok := seen[leftPoint]; !ok {
						seen[leftPoint] = struct{}{}
					}
					if leftPoint.Value == 9 {
						numAnswers++
					} else {
						stack = append(stack, leftPoint)
					}
				}
			}
		}
	}
	return numAnswers, nil
}
