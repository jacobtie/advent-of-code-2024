package main

import (
	"fmt"
	"math"
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

var recordRE = regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)\nButton B: X\+(\d+), Y\+(\d+)\nPrize: X=(\d+), Y=(\d+)`)

type Node struct {
	Cost int
	NumA int
	NumB int
	X    int
	Y    int
}

func runExercise1() (int, error) {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	records := strings.Split(strings.TrimSpace(string(fileBytes)), "\n\n")
	totalCost := 0
	for _, record := range records {
		matches := recordRE.FindAllStringSubmatch(record, -1)
		buttonAX, err := strconv.Atoi(matches[0][1])
		if err != nil {
			return 0, err
		}
		buttonAY, err := strconv.Atoi(matches[0][2])
		if err != nil {
			return 0, err
		}
		buttonBX, err := strconv.Atoi(matches[0][3])
		if err != nil {
			return 0, err
		}
		buttonBY, err := strconv.Atoi(matches[0][4])
		if err != nil {
			return 0, err
		}
		prizeX, err := strconv.Atoi(matches[0][5])
		if err != nil {
			return 0, err
		}
		prizeY, err := strconv.Atoi(matches[0][6])
		if err != nil {
			return 0, err
		}
		foundGoal, lowestCost := false, math.MaxInt
		frontier := []Node{
			{Cost: 3, NumA: 1, NumB: 0, X: buttonAX, Y: buttonAY},
			{Cost: 1, NumA: 0, NumB: 1, X: buttonBX, Y: buttonBY},
		}
		seen := make(map[Node]struct{})
		for len(frontier) > 0 {
			currentNode := frontier[len(frontier)-1]
			frontier = frontier[:len(frontier)-1]
			if _, ok := seen[currentNode]; ok {
				continue
			}
			seen[currentNode] = struct{}{}
			if currentNode.X == prizeX && currentNode.Y == prizeY {
				foundGoal = true
				if currentNode.Cost < lowestCost {
					lowestCost = currentNode.Cost
				}
				continue
			}
			if currentNode.X >= prizeX || currentNode.Y >= prizeY {
				continue
			}
			if currentNode.NumA < 100 {
				frontier = append(frontier, Node{
					Cost: currentNode.Cost + 3,
					NumA: currentNode.NumA + 1,
					NumB: currentNode.NumB,
					X:    currentNode.X + buttonAX,
					Y:    currentNode.Y + buttonAY,
				})
			}
			if currentNode.NumB < 100 {
				frontier = append(frontier, Node{
					Cost: currentNode.Cost + 1,
					NumA: currentNode.NumA,
					NumB: currentNode.NumB + 1,
					X:    currentNode.X + buttonBX,
					Y:    currentNode.Y + buttonBY,
				})
			}
		}
		if foundGoal {
			totalCost += lowestCost
		}
	}
	return totalCost, nil
}

func runExercise2() (int, error) {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	records := strings.Split(strings.TrimSpace(string(fileBytes)), "\n\n")
	totalCost := 0
	for _, record := range records {
		matches := recordRE.FindAllStringSubmatch(record, -1)
		buttonAX, err := strconv.Atoi(matches[0][1])
		if err != nil {
			return 0, err
		}
		buttonAY, err := strconv.Atoi(matches[0][2])
		if err != nil {
			return 0, err
		}
		buttonBX, err := strconv.Atoi(matches[0][3])
		if err != nil {
			return 0, err
		}
		buttonBY, err := strconv.Atoi(matches[0][4])
		if err != nil {
			return 0, err
		}
		prizeX, err := strconv.Atoi(matches[0][5])
		if err != nil {
			return 0, err
		}
		prizeX += 10000000000000
		prizeY, err := strconv.Atoi(matches[0][6])
		if err != nil {
			return 0, err
		}
		prizeY += 10000000000000
		d := (buttonAX * buttonBY) - (buttonBX * buttonAY)
		a := (buttonBY * prizeX) - (buttonBX * prizeY)
		b := (buttonAX * prizeY) - (buttonAY * prizeX)
		if a%d != 0 || b%d != 0 {
			continue
		}
		totalCost += (3 * (a / d)) + (b / d)
	}
	return totalCost, nil
}
