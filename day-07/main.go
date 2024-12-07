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

var numRE = regexp.MustCompile(`\d+`)

type StackNode struct {
	RollingOperand int
	AddNext        bool
	NextNums       []int
}

func runExercise1() (int, error) {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	input := strings.TrimSpace(string(fileBytes))
	summation := 0
	for _, record := range strings.Split(input, "\n") {
		rawNums := numRE.FindAllString(record, -1)
		target, err := strconv.Atoi(rawNums[0])
		if err != nil {
			return 0, err
		}
		nums := make([]int, 0, len(rawNums)-1)
		for i := 1; i < len(rawNums); i++ {
			num, err := strconv.Atoi(rawNums[i])
			if err != nil {
				return 0, err
			}
			nums = append(nums, num)
		}
		isPossible := false
		stack := []StackNode{
			StackNode{RollingOperand: nums[0], AddNext: true, NextNums: nums[1:]},
			StackNode{RollingOperand: nums[0], AddNext: false, NextNums: nums[1:]},
		}
		for len(stack) > 0 {
			if isPossible {
				break
			}
			// Pop off stack
			stackNode := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			if len(stackNode.NextNums) == 0 {
				if stackNode.RollingOperand == target {
					isPossible = true
				}
				continue
			}
			nextRollingOperand := stackNode.RollingOperand
			if stackNode.AddNext {
				nextRollingOperand += stackNode.NextNums[0]
			} else {
				nextRollingOperand *= stackNode.NextNums[0]
			}
			stack = append(stack, StackNode{RollingOperand: nextRollingOperand, AddNext: true, NextNums: stackNode.NextNums[1:]})
			stack = append(stack, StackNode{RollingOperand: nextRollingOperand, AddNext: false, NextNums: stackNode.NextNums[1:]})
		}
		if isPossible {
			summation += target
		}
	}
	return summation, nil
}

type Part2StackNode struct {
	RollingOperand int
	NextOperator   rune
	NextNums       []int
}

func runExercise2() (int, error) {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	input := strings.TrimSpace(string(fileBytes))
	summation := 0
	for _, record := range strings.Split(input, "\n") {
		rawNums := numRE.FindAllString(record, -1)
		target, err := strconv.Atoi(rawNums[0])
		if err != nil {
			return 0, err
		}
		nums := make([]int, 0, len(rawNums)-1)
		for i := 1; i < len(rawNums); i++ {
			num, err := strconv.Atoi(rawNums[i])
			if err != nil {
				return 0, err
			}
			nums = append(nums, num)
		}
		isPossible := false
		stack := []Part2StackNode{
			Part2StackNode{RollingOperand: nums[0], NextOperator: '+', NextNums: nums[1:]},
			Part2StackNode{RollingOperand: nums[0], NextOperator: '*', NextNums: nums[1:]},
			Part2StackNode{RollingOperand: nums[0], NextOperator: '|', NextNums: nums[1:]},
		}
		for len(stack) > 0 {
			if isPossible {
				break
			}
			// Pop off stack
			stackNode := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			if len(stackNode.NextNums) == 0 {
				if stackNode.RollingOperand == target {
					isPossible = true
				}
				continue
			}
			nextRollingOperand := stackNode.RollingOperand
			switch stackNode.NextOperator {
			case '+':
				nextRollingOperand += stackNode.NextNums[0]
			case '*':
				nextRollingOperand *= stackNode.NextNums[0]
			case '|':
				nextRollingOperand, err = strconv.Atoi(fmt.Sprintf("%d%d", nextRollingOperand, stackNode.NextNums[0]))
				if err != nil {
					return 0, err
				}
			}
			stack = append(stack, Part2StackNode{RollingOperand: nextRollingOperand, NextOperator: '+', NextNums: stackNode.NextNums[1:]})
			stack = append(stack, Part2StackNode{RollingOperand: nextRollingOperand, NextOperator: '*', NextNums: stackNode.NextNums[1:]})
			stack = append(stack, Part2StackNode{RollingOperand: nextRollingOperand, NextOperator: '|', NextNums: stackNode.NextNums[1:]})
		}
		if isPossible {
			summation += target
		}
	}
	return summation, nil
}
