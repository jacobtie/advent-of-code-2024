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

type LinkedNode struct {
	Value string
	Next  *LinkedNode
	Prev  *LinkedNode
}

func (l *LinkedNode) PrintList() {
	var nums []string
	for curr := l; curr != nil; curr = curr.Next {
		nums = append(nums, curr.Value)
	}
	fmt.Println(nums)
}

func (l *LinkedNode) Size() int {
	count := 0
	for curr := l; curr != nil; curr = curr.Next {
		count++
	}
	return count
}

func runExercise1() (int, error) {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, nil
	}
	rawNums := strings.Split(strings.TrimSpace(string(fileBytes)), " ")
	var first *LinkedNode
	for i := len(rawNums) - 1; i >= 0; i-- {
		if first == nil {
			first = &LinkedNode{Value: rawNums[i]}
			continue
		}
		prev := &LinkedNode{Value: rawNums[i], Next: first}
		first.Prev = prev
		first = prev
	}
	for i := 0; i < 25; i++ {
		for curr := first; curr != nil; curr = curr.Next {
			if curr.Value == "0" {
				curr.Value = "1"
				continue
			}
			if len(curr.Value)%2 == 0 {
				midIndex := len(curr.Value) / 2
				firstHalf, secondHalf := curr.Value[:midIndex], curr.Value[midIndex:]
				for strings.HasPrefix(secondHalf, "0") && len(secondHalf) > 1 {
					secondHalf = strings.TrimPrefix(secondHalf, "0")
				}
				curr.Value = firstHalf
				newNode := &LinkedNode{Value: secondHalf, Prev: curr, Next: curr.Next}
				curr.Next = newNode
				curr = curr.Next
				continue
			}
			num, err := strconv.Atoi(curr.Value)
			if err != nil {
				return 0, err
			}
			curr.Value = fmt.Sprint(num * 2024)
		}
	}
	return first.Size(), nil
}

func runExercise2() (int, error) {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, nil
	}
	rawNums := strings.Split(strings.TrimSpace(string(fileBytes)), " ")
	stones := make(map[string]int)
	for i := len(rawNums) - 1; i >= 0; i-- {
		if _, ok := stones[rawNums[i]]; !ok {
			stones[rawNums[i]] = 0
		}
		stones[rawNums[i]]++
	}
	for i := 0; i < 75; i++ {
		stonesCopy := make(map[string]int)
		for num, count := range stones {
			if num == "0" {
				if _, ok := stonesCopy["1"]; !ok {
					stonesCopy["1"] = 0
				}
				stonesCopy["1"] += count
				continue
			}
			if len(num)%2 == 0 {
				midIndex := len(num) / 2
				firstHalf, secondHalf := num[:midIndex], num[midIndex:]
				for strings.HasPrefix(secondHalf, "0") && len(secondHalf) > 1 {
					secondHalf = strings.TrimPrefix(secondHalf, "0")
				}
				if _, ok := stonesCopy[firstHalf]; !ok {
					stonesCopy[firstHalf] = 0
				}
				stonesCopy[firstHalf] += count
				if _, ok := stonesCopy[secondHalf]; !ok {
					stonesCopy[secondHalf] = 0
				}
				stonesCopy[secondHalf] += count
				continue
			}
			realNum, err := strconv.Atoi(num)
			if err != nil {
				return 0, nil
			}
			newNum := fmt.Sprint(realNum * 2024)
			if _, ok := stonesCopy[newNum]; !ok {
				stonesCopy[newNum] = 0
			}
			stonesCopy[newNum] += count
		}
		stones = stonesCopy
	}
	summation := 0
	for _, count := range stones {
		summation += count
	}
	return summation, nil
}
