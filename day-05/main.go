package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
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

var constraintRE = regexp.MustCompile(`(\d+)|(\d+)`)

func runExercise1() (int, error) {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	rows := strings.Split(strings.TrimSpace(string(fileBytes)), "\n")
	constraints := make(map[string]map[string]struct{})
	idx := 0
	for {
		if rows[idx] == "" {
			idx++
			break
		}
		matches := constraintRE.FindAllStringSubmatch(rows[idx], -1)
		num1, num2 := matches[0][0], matches[1][0]
		if _, ok := constraints[num1]; !ok {
			constraints[num1] = make(map[string]struct{})
		}
		constraints[num1][num2] = struct{}{}
		idx++
	}
	summation := 0
	for ; idx < len(rows); idx++ {
		isValid := true
		nums := strings.Split(rows[idx], ",")
		for i := len(nums) - 1; i >= 0; i-- {
			num := nums[i]
			numConstraints, ok := constraints[num]
			if !ok {
				continue
			}
			rest := nums[0:i]
			for _, restNum := range rest {
				if _, ok := numConstraints[restNum]; ok {
					isValid = false
					break
				}
			}
			if !isValid {
				break
			}
		}
		if isValid {
			midNumStr := nums[len(nums)/2]
			midNum, err := strconv.Atoi(midNumStr)
			if err != nil {
				return 0, err
			}
			summation += midNum
		}
	}
	return summation, nil
}

func runExercise2() (int, error) {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	rows := strings.Split(strings.TrimSpace(string(fileBytes)), "\n")
	constraints := make(map[string]map[string]struct{})
	idx := 0
	for {
		if rows[idx] == "" {
			idx++
			break
		}
		matches := constraintRE.FindAllStringSubmatch(rows[idx], -1)
		num1, num2 := matches[0][0], matches[1][0]
		if _, ok := constraints[num1]; !ok {
			constraints[num1] = make(map[string]struct{})
		}
		constraints[num1][num2] = struct{}{}
		idx++
	}
	summation := 0
	for ; idx < len(rows); idx++ {
		isValid := true
		nums := strings.Split(rows[idx], ",")
		for i := len(nums) - 1; i >= 0; i-- {
			num := nums[i]
			numConstraints, ok := constraints[num]
			if !ok {
				continue
			}
			rest := nums[0:i]
			for _, restNum := range rest {
				if _, ok := numConstraints[restNum]; ok {
					isValid = false
					break
				}
			}
			if !isValid {
				break
			}
		}
		if isValid {
			continue
		}
		sort.Slice(nums, func(i, j int) bool {
			num1, num2 := nums[i], nums[j]
			numConstraints, ok := constraints[num1]
			if ok {
				if _, ok = numConstraints[num2]; ok {
					return true
				}
			}
			numConstraints, ok = constraints[num2]
			if ok {
				if _, ok := numConstraints[num1]; ok {
					return false
				}
			}
			return true
		})
		midNumStr := nums[len(nums)/2]
		midNum, err := strconv.Atoi(midNumStr)
		if err != nil {
			return 0, err
		}
		summation += midNum
	}
	return summation, nil
}
