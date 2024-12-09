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

func runExercise1() (int, error) {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	input := strings.TrimSpace(string(fileBytes))
	result := make([]int, 0)
	isBlock := true
	for idx, countRune := range input {
		count, err := strconv.Atoi(string(countRune))
		if err != nil {
			return 0, nil
		}
		for i := 0; i < count; i++ {
			toAppend := idx / 2
			if !isBlock {
				toAppend = -1
			}
			result = append(result, toAppend)
		}
		isBlock = !isBlock
	}
	for i := 0; i < len(result); i++ {
		if result[i] != -1 {
			continue
		}
		for i == len(result) || result[i] == -1 {
			lastElement := result[len(result)-1]
			result = result[0 : len(result)-1]
			if i == len(result) {
				break
			}
			if lastElement == -1 {
				continue
			}
			result[i] = lastElement
		}
	}
	summation := 0
	for idx, num := range result {
		summation += idx * num
	}
	return summation, nil
}

func runExercise2() (int, error) {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	input := strings.TrimSpace(string(fileBytes))
	result := make([]int, 0)
	isBlock := true
	for idx, countRune := range input {
		count, err := strconv.Atoi(string(countRune))
		if err != nil {
			return 0, nil
		}
		for i := 0; i < count; i++ {
			toAppend := idx / 2
			if !isBlock {
				toAppend = -1
			}
			result = append(result, toAppend)
		}
		isBlock = !isBlock
	}
	for rightIDX := len(result) - 1; rightIDX >= 0; {
		if rightIDX == -1 {
			rightIDX--
			continue
		}
		length := 0
		num := result[rightIDX]
		for idx := rightIDX; idx >= 0 && result[idx] == num; idx-- {
			length++
		}
		foundSpace := false
		startFreeSpace := 0
		for leftIDX := 0; leftIDX <= rightIDX; leftIDX++ {
			if result[leftIDX] != -1 {
				startFreeSpace = leftIDX + 1
				continue
			}
			if leftIDX+1-startFreeSpace == length {
				foundSpace = true
				break
			}
		}
		if !foundSpace {
			rightIDX -= length
			continue
		}
		for i := startFreeSpace; i < startFreeSpace+length; i++ {
			result[i] = num
			result[rightIDX] = -1
			rightIDX--
		}
	}
	summation := 0
	for idx, num := range result {
		if num == -1 {
			continue
		}
		summation += idx * num
	}
	return summation, nil
}
