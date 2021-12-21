package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var testInput []int
var input []int

func main() {
	testInput = getInput("day1.test")
	input = getInput("day1.input")

	fmt.Println("Part 1 test:", getPart1Result(testInput))
	fmt.Println("Part 1 result:", getPart1Result(input))
	fmt.Println("Part 2 test:", getPart2Result(testInput))
	fmt.Println("Part 1 test:", getPart2Result(input))
}

func getPart1Result(input []int) int {
	incCount := 0

	for i, v := range input {
		if i > 0 && v > input[i-1] {
			incCount++
		}
	}

	return incCount
}

func getPart2Result(input []int) int {
	incCount := 0

	for i := 1; i < len(input)-2; i++ {
		subs1 := input[i-1 : i+2]
		subs2 := input[i : i+3]

		if sum(subs1) < sum(subs2) {
			incCount++
		}
	}

	return incCount
}

func getInput(fileName string) []int {
	input, _ := os.Open("inputs/" + fileName)
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var result []int
	for scanner.Scan() {
		intRes, _ := strconv.Atoi(scanner.Text())
		result = append(result, intRes)
	}

	return result
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}

	return result
}
