package main

import (
	"fmt"
	"os"
	"strings"
)

func main2() {
	totalRed := 12
	totalGreen := 13
	totalBlue := 14
	sum := 0
	file, err := os.ReadFile("input2.txt")
	content := string(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(content, "\n") // creates an array of each line
	count := 1                            //game ID count
	for _, line := range lines {
		if check(line, totalBlue, totalGreen, totalRed) {
			sum += count
		}
		count++
	}

	fmt.Println(sum)
}

func check(line string, totalRed int, totalGreen int, totalBlue int) bool {
	allRounds := strings.Split(line, ":")
	allRounds = allRounds[1:]
	// now allRounds= [“2 green, 8 red, 3 blue; 2 green, 4 blue, 2 red; 2 green, 5 blue, 2 red”]
	ArrayAllRounds := strings.Split(allRounds[0], ";")
	//ArrayallRounds= [“2 green, 8 red, 3 blue” ,” 2 green, 4 blue, 2 red”,” 2 green, 5 blue, 2 red”]

	for _, rounds := range ArrayAllRounds {
		// rounds ="5 red, 2 blue, 2 green"
		x := strings.Split(rounds, ",")
		//x = ["5 red","2 blue","2 green"]
		for _, item := range x {
			switch {
			case strings.Contains(item, "red"):
				redNumber := getNumber(item)
				if redNumber > totalRed {
					return false
				}
			case strings.Contains(item, "blue"):
				blueNumber := getNumber(item)
				if blueNumber > totalBlue {
					return false
				}
			case strings.Contains(item, "green"):
				greenNumber := getNumber(item)
				if greenNumber > totalGreen {
					return false
				}
			}
		}
	}
	return true
}

func getNumber(line string) int { // takes "5 red" as input and returns 5
	var num string
	if len(line) == 0 {
		return 0
	}

	// var number string
	for _, char := range line {
		if char >= 48 && char <= 57 {
			digit := int(char - '0')
			num += fmt.Sprint(digit)
		}
	}

	return strToInt(num)
}

func strToInt(s string) int {
	var result int

	// Iterate through each character of the string
	for _, digit := range s {
		// Convert the character to its integer representation
		digitValue := int(digit - '0')

		// Shift the existing result to the left by one position and add the new digit
		result = result*10 + digitValue
	}

	return result
}
