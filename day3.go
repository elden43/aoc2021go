package main

import (
	"aoc2021/inputter"
	"fmt"
	"strconv"
)

func main() {

	testInput := inputter.GetLines("inputs/day3.test")
	input := inputter.GetLines("inputs/day3.input")

	fmt.Println("Part 1 test: ", getPart1Result(getRates(testInput)))
	fmt.Println("Part 1 result: ", getPart1Result(getRates(input)))

}

func getPart1Result(gamma string, epsilon string) (int64) {
	var gammaInt, epsilonInt int64
	gammaInt, _ = strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ = strconv.ParseInt(epsilon, 2, 64)
	return gammaInt * epsilonInt
}

func getRates(input []string) (string, string) {
	gamma := ""
	epsilon := ""
	zeroCount := 0

	for i := 0; i < len(input[0]); i++ {
		zeroCount = 0
		for _, v := range input {
			if string(v[i]) == "0" {
				zeroCount++
			}
		}
		if zeroCount > len(input)/2 {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	return gamma, epsilon
}
