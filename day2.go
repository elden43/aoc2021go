package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1 test:", getPart1Result("day2.test"))
	fmt.Println("Part 1 result:", getPart1Result("day2.input"))
	fmt.Println("Part 2 test:", getPart2Result("day2.test"))
	fmt.Println("Part 1 test:", getPart2Result("day2.input"))
}

type instruction struct {
	direction string
	amount    int
}

var instructionVersion = map[string]int{
	"day1": 1,
	"day2": 2,
}

func getPart1Result(fileName string) int {
	input := getInput(fileName)
	horPos := 0
	depth := 0
	aim := 0
	for _, v := range input {
		processInstruction(v, &horPos, &depth, &aim, instructionVersion["day1"])
	}

	return horPos * depth
}

func getPart2Result(fileName string) int {
	input := getInput(fileName)
	horPos := 0
	depth := 0
	aim := 0
	for _, v := range input {
		processInstruction(v, &horPos, &depth, &aim, instructionVersion["day2"])
	}

	return horPos * depth
}

func processInstruction(v instruction, horPos *int, depth *int, aim *int, version int) {
	if version == 1 {
		switch v.direction {
		case "forward":
			*horPos += v.amount
		case "down":
			*depth += v.amount
		case "up":
			*depth -= v.amount
		}
	}
	if version == 2 {
		switch v.direction {
		case "forward":
			*horPos += v.amount
			*depth += *aim * v.amount
		case "down":
			*aim += v.amount
		case "up":
			*aim -= v.amount
		}
	}

}

func getInput(fileName string) []instruction {
	fileBytes, err := ioutil.ReadFile("inputs/" + fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var res []instruction
	var tmpRes []string
	strSlice := strings.Split(string(fileBytes), "\r\n")

	for _, v := range strSlice {
		tmpRes = strings.Split(v, " ")
		res = append(res, instruction{direction: tmpRes[0], amount: str2int(tmpRes[1])})
	}

	return res
}

func str2int(str string) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return v
}
