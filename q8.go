package main

import (
	"fmt"
	"os"
	"strings"
)

const actualSequence = "LRRLRRRLRRLRRLRRRLRRLRLLRRRLRRRLRRRLRRRLRRRLRRLLRRRLLRLRRRLRRLRRRLLRRRLRLLRRLRLLRLRLLRRLRRRLRRLLRRRLRRRLRLRRRLRRRLRRRLRRRLRLRRRLLRRRLRRLRRRLRLRRRLRRLRLLLLLRRRLRRRLRRRLRRRLRRLLRLRLRRLRRLLRRRLRRRLRRRLLLRRRLRRRLRRRLRLRRRLLRLRLRRLRRLRRRLRRLRRRLRRRLRRRLRRLLRRRLRRLRRLRLLRRRR"

func main() {
	file, err := os.ReadFile("input8.txt")
	if err != nil {
		panic(err)
	}
	content := string(file)
	x := strings.Split(content, "\n\n")
	sequenceOfInstructions := x[0]
	y := strings.Split(x[1], "\n")
	locationMap := make(map[string][]string)
	//getting the input correctly
	for i := 0; i < len(y); i++ {
		temp := strings.Split(y[i], "=")
		leftSide := temp[0]
		rightside := strings.Split(temp[1], ",")
		locationMap[leftSide] = rightside
	}
	current := "AAA"
	stepCount := findNoOfSteps(locationMap, current, sequenceOfInstructions, 0)
	fmt.Println(stepCount)

}

func findNoOfSteps(locationMap map[string][]string, current string, sequence string, stepCount int) int {
	if current == "ZZZ" {
		return stepCount
	}
	if sequence == "" {
		sequence = actualSequence
	}
	nextInstruction := string(sequence[0])
	var nextElement string
	if nextInstruction == "L" {
		nextElement = locationMap[current][0]
	} else if nextInstruction == "R" {
		nextElement = locationMap[current][1]
	} else {
		panic("error")
	}
	return findNoOfSteps(locationMap, nextElement, sequence[1:], stepCount+1)
}
