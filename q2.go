package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	totalRed := 12
	totalGreen := 13
	totalBlue := 14
	sum := 0
	file, err := os.ReadFile("testinput2.txt")
	content := string(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(content, "\n")
	count := 1 //game ID count
	for _, line := range lines {
		if check(line, totalBlue, totalGreen, totalRed) {
			sum += count
		}
		count++
	}
}

func check(line string, totalRed int, totalGreen int, totalBlue int) bool {
	allRounds := strings.Split(line, ":")
	allRounds = delete(allRounds, 0)
	allRounds = strings.Split(line, ";")
	//now we have an array containing each round of the game
	for _, rounds := range allRounds {
		// rounds ="5 red, 2 blue, 2 green"
		x := strings.Split(rounds, ",")
		//x = ["5 red","2 blue","2 green"]
		for i := 0; i < len(x); i++ {
			if strings.Contains(x[i], "red") {
				redNumber := getNumber(x[i])
				if redNumber < totalRed {
					return false
				}
			}
			if strings.Contains(x[i], "blue") {
				blueNumber := getNumber(x[i])
				if blueNumber < totalBlue {
					return false
				}
			}
			if strings.Contains(x[i], "green") {
				greenNumber := getNumber(x[i])
				if greenNumber < totalGreen {
					return false
				}
			}
		}

	}
	return true
}

func delete(line []string, index int) []string {

	if index < 0 || index >= len(line) {
		// Handle out-of-bounds index (optional)
		return line
	}
	return append(line[:index], line[index+1:]...)

}

func getNumber(line string) int { // takes "5 red" as input and returns 5
	stringnum := ""
	for i := 0; i < len(line); i++ {

		if string(line[i]) == " " {
			number, err := strconv.Atoi(string(stringnum))

			if err != nil {
				panic(err)
			}
			return number

		}
		stringnum += string(line[i])
	}
	return 0
}
