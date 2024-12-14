package main

import (
	"fmt"
	"os"
	"regexp"
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

const IS_EXERCISE_1_EXAMPLE = false

var numRE = regexp.MustCompile(`\-?\d+`)

func runExercise1() (int, error) {
	inputFile := "input.txt"
	boardX, boardY := 101, 103
	if IS_EXERCISE_1_EXAMPLE {
		inputFile = "example.txt"
		boardX, boardY = 11, 7
	}
	fileBytes, err := os.ReadFile(inputFile)
	if err != nil {
		return 0, err
	}
	records := strings.Split(strings.TrimSpace(string(fileBytes)), "\n")
	quads := map[rune]int{'a': 0, 'b': 0, 'c': 0, 'd': 0}
	grid := getGrid(boardX, boardY)
	for _, record := range records {
		matches := numRE.FindAllString(record, -1)
		startX, err := strconv.Atoi(matches[0])
		if err != nil {
			return 0, err
		}
		startY, err := strconv.Atoi(matches[1])
		if err != nil {
			return 0, err
		}
		vectorX, err := strconv.Atoi(matches[2])
		if err != nil {
			return 0, err
		}
		vectorY, err := strconv.Atoi(matches[3])
		if err != nil {
			return 0, err
		}
		endX := (((startX + (vectorX * 100)) % boardX) + boardX) % boardX
		endY := (((startY + (vectorY * 100)) % boardY) + boardY) % boardY
		if endX < int(boardX/2) && endY < int(boardY/2) {
			quads['a']++
		}
		if endX > int(boardX/2) && endY < int(boardY/2) {
			quads['b']++
		}
		if endX > int(boardX/2) && endY > int(boardY/2) {
			quads['c']++
		}
		if endX < int(boardX/2) && endY > int(boardY/2) {
			quads['d']++
		}
		grid[endY][endX]++
	}
	fmt.Println(quads)
	product := 1
	for _, num := range quads {
		product *= num
	}
	return product, nil
}

type Robot struct {
	PosX int
	PosY int
	VecX int
	VecY int
}

func runExercise2() (int, error) {
	inputFile := "input.txt"
	boardX, boardY := 101, 103
	if IS_EXERCISE_1_EXAMPLE {
		inputFile = "example.txt"
		boardX, boardY = 11, 7
	}
	fileBytes, err := os.ReadFile(inputFile)
	if err != nil {
		return 0, err
	}
	records := strings.Split(strings.TrimSpace(string(fileBytes)), "\n")
	robots := make([]*Robot, 0, len(records))
	for _, record := range records {
		matches := numRE.FindAllString(record, -1)
		startX, err := strconv.Atoi(matches[0])
		if err != nil {
			return 0, err
		}
		startY, err := strconv.Atoi(matches[1])
		if err != nil {
			return 0, err
		}
		vectorX, err := strconv.Atoi(matches[2])
		if err != nil {
			return 0, err
		}
		vectorY, err := strconv.Atoi(matches[3])
		if err != nil {
			return 0, err
		}
		robots = append(robots, &Robot{PosX: startX, PosY: startY, VecX: vectorX, VecY: vectorY})
	}
	secondsElapsed := 0
	file, err := os.OpenFile("output.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	for secondsElapsed < 10_000 {
		secondsElapsed++
		fmt.Printf("Second: %d\n", secondsElapsed)
		grid := getGrid(boardX, boardY)
		for _, robot := range robots {
			robot.PosX = (((robot.PosX + robot.VecX) % boardX) + boardX) % boardX
			robot.PosY = (((robot.PosY + robot.VecY) % boardY) + boardY) % boardY
			grid[robot.PosY][robot.PosX]++
		}
		if _, err := file.WriteString(fmt.Sprintf("Seconds Elapsed: %d\n", secondsElapsed)); err != nil {
			return 0, err
		}
		for _, row := range grid {
			for _, num := range row {
				char := "."
				if num > 0 {
					char = "#"
				}
				if _, err := file.WriteString(fmt.Sprintf("%s ", char)); err != nil {
					return 0, err
				}
			}
			if _, err := file.WriteString("\n"); err != nil {
				return 0, err
			}
		}
	}
	return 0, nil
}

func getGrid(boardX, boardY int) [][]int {
	grid := make([][]int, 0, boardY)
	for i := 0; i < boardY; i++ {
		row := make([]int, 0, boardX)
		for j := 0; j < boardX; j++ {
			row = append(row, 0)
		}
		grid = append(grid, row)
	}
	return grid
}
