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

func runExercise1() (int, error) {
	summation := 0
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	input := strings.TrimSpace(string(fileBytes))
	rows := strings.Split(input, "\n")
	for i := range rows {
		for j := range rows[i] {
			if rows[i][j] != 'X' {
				continue
			}
			if isXMASAbove(rows, i, j) {
				summation++
			}
			if isXMASRight(rows, i, j) {
				summation++
			}
			if isXMASDown(rows, i, j) {
				summation++
			}
			if isXMASLeft(rows, i, j) {
				summation++
			}
			if isXMASAboveLeft(rows, i, j) {
				summation++
			}
			if isXMASAboveRight(rows, i, j) {
				summation++
			}
			if isXMASDownRight(rows, i, j) {
				summation++
			}
			if isXMASDownLeft(rows, i, j) {
				summation++
			}
		}
	}
	return summation, nil
}

func isXMASAbove(rows []string, i, j int) bool {
	if i < 3 {
		return false
	}
	return rows[i-1][j] == 'M' && rows[i-2][j] == 'A' && rows[i-3][j] == 'S'
}

func isXMASRight(rows []string, i, j int) bool {
	if j > len(rows[i])-4 {
		return false
	}
	return rows[i][j+1] == 'M' && rows[i][j+2] == 'A' && rows[i][j+3] == 'S'
}

func isXMASDown(rows []string, i, j int) bool {
	if i > len(rows)-4 {
		return false
	}
	return rows[i+1][j] == 'M' && rows[i+2][j] == 'A' && rows[i+3][j] == 'S'
}

func isXMASLeft(rows []string, i, j int) bool {
	if j < 3 {
		return false
	}
	return rows[i][j-1] == 'M' && rows[i][j-2] == 'A' && rows[i][j-3] == 'S'
}

func isXMASAboveRight(rows []string, i, j int) bool {
	if i < 3 || j > len(rows[i])-4 {
		return false
	}
	return rows[i-1][j+1] == 'M' && rows[i-2][j+2] == 'A' && rows[i-3][j+3] == 'S'
}

func isXMASDownRight(rows []string, i, j int) bool {
	if i > len(rows)-4 || j > len(rows[i])-4 {
		return false
	}
	return rows[i+1][j+1] == 'M' && rows[i+2][j+2] == 'A' && rows[i+3][j+3] == 'S'
}

func isXMASDownLeft(rows []string, i, j int) bool {
	if i > len(rows)-4 || j < 3 {
		return false
	}
	return rows[i+1][j-1] == 'M' && rows[i+2][j-2] == 'A' && rows[i+3][j-3] == 'S'
}

func isXMASAboveLeft(rows []string, i, j int) bool {
	if i < 3 || j < 3 {
		return false
	}
	return rows[i-1][j-1] == 'M' && rows[i-2][j-2] == 'A' && rows[i-3][j-3] == 'S'
}

func runExercise2() (int, error) {
	summation := 0
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	input := strings.TrimSpace(string(fileBytes))
	rows := strings.Split(input, "\n")
	for i := range rows {
		for j := range rows[i] {
			if rows[i][j] != 'A' {
				continue
			}
			if isMASCross(rows, i, j) {
				summation++
			}
		}
	}
	return summation, nil
}

func isMASCross(rows []string, i, j int) bool {
	if i == 0 || j == 0 || i == len(rows)-1 || j == len(rows[i])-1 {
		return false
	}
	upperLeft := rows[i-1][j-1]
	upperRight := rows[i-1][j+1]
	lowerRight := rows[i+1][j+1]
	lowerLeft := rows[i+1][j-1]
	if !((upperLeft == 'M' && lowerRight == 'S') || (upperLeft == 'S' && lowerRight == 'M')) {
		return false
	}
	if !((upperRight == 'M' && lowerLeft == 'S') || (upperRight == 'S' && lowerLeft == 'M')) {
		return false
	}
	return true
}
