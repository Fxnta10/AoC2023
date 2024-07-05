package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main5() {
	file, err := os.ReadFile("input4.txt")
	content := string(file)
	if err != nil {
		panic(err)
	}
	actualsum := 0
	// lines := strings.Split(content, "\n")
	games := strings.Split(content, "\n")

	for _, x := range games {
		var sum float64 = 0
		count := 0
		y := strings.Split(x, "|")
		winningRounds := strings.Split(y[0], " ")
		drawnRounds := strings.Split(y[1], " ")
		winningRounds = removeEmptyStrings(winningRounds)
		drawnRounds = removeEmptyStrings(drawnRounds)
		elementMap := make(map[string]bool)

		for _, num := range winningRounds {
			elementMap[num] = true
		}

		for _, num := range drawnRounds {
			if elementMap[num] {
				count++
			}

		}
		sum = math.Pow(2, float64(count-1))
		actualsum += int(sum)
	}
	fmt.Println(actualsum)
}
func removeEmptyStrings(arr []string) []string {
	result := make([]string, 0, len(arr))

	for _, str := range arr {
		if str != "" {
			result = append(result, str)
		}
	}

	return result
}
