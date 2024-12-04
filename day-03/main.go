package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
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

var mulRE = regexp.MustCompile(`mul\(\d+,\d+\)`)
var singleMulRE = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func runExercise1() (int, error) {
	summation := 0
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	input := string(fileBytes)
	matches := mulRE.FindAllString(input, -1)
	for _, match := range matches {
		nums := singleMulRE.FindAllStringSubmatch(match, -1)
		firstNum, err := strconv.Atoi(nums[0][1])
		if err != nil {
			return 0, err
		}
		secondNum, err := strconv.Atoi(nums[0][2])
		if err != nil {
			return 0, err
		}
		summation += firstNum * secondNum
	}
	return summation, nil
}

var instructionRE = regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)

func runExercise2() (int, error) {
	summation := 0
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	input := string(fileBytes)
	matches := instructionRE.FindAllString(input, -1)
	shouldMul := true
	for _, match := range matches {
		if match == "don't()" {
			shouldMul = false
			continue
		}
		if match == "do()" {
			shouldMul = true
			continue
		}
		if !shouldMul {
			continue
		}
		nums := singleMulRE.FindAllStringSubmatch(match, -1)
		firstNum, err := strconv.Atoi(nums[0][1])
		if err != nil {
			return 0, err
		}
		secondNum, err := strconv.Atoi(nums[0][2])
		if err != nil {
			return 0, err
		}
		summation += firstNum * secondNum
	}
	return summation, nil
}
