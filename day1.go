package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1 test:", getPart1Result("day1.test"))
	fmt.Println("Part 1 result:", getPart1Result("day1.input"))
	fmt.Println("Part 2 test:", getPart2Result("day1.test"))
	fmt.Println("Part 1 test:", getPart2Result("day1.input"))
}

func getPart1Result(fileName string) int {
	incCount := 0
	input := getInput(fileName)

	for i, v := range input {
		if i > 0 && v > input[i-1] {
			incCount++
		}
	}

	return incCount
}

func getPart2Result(fileName string) int {
	incCount := 0
	input := getInput(fileName)

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
	fileBytes, err := ioutil.ReadFile("inputs/" + fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var res []int
	var result int
	strSlice := strings.Split(string(fileBytes), "\r\n")

	for _, v := range strSlice {
		result, err = strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		res = append(res, result)
	}

	return res
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}

	return result
}
