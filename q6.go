package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main7() {
	file, err := os.ReadFile("input6.txt")
	if err != nil {
		panic(err)
	}
	content := string(file)
	x := strings.Split(content, "\n")

	y := strings.Split(x[0], "  ")
	z := strings.Split(x[1], "  ")

	time := convertToIntegers(y)
	distance := convertToIntegers(z)
	ans := 1
	for i := 0; i < len(time); i++ {
		timeRecord := time[i]
		distanceRecord := distance[i]
		count := 0
		for j := 0; j < int(timeRecord); j++ {
			dist := (timeRecord - j) * j
			if dist >= distanceRecord {
				count++
			}
		}
		fmt.Println(count)
		ans = ans * count
	}
	fmt.Println(ans)
}

func convertToIntegers(strArray []string) []int {
	intArray := make([]int, len(strArray))

	for i, str := range strArray {
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(err) // return early if conversion fails
		}
		intArray[i] = num
	}

	return intArray
}
