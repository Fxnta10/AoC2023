package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main1() {
	sum := 0
	file, err := os.ReadFile("input1.txt")
	content := string(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		num := getdigits(line)
		intnum, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Error:", err)

		}
		sum += intnum
	}
	fmt.Println(sum)
}

func getdigits(line string) string {
	var numarray []int
	for _, char := range line {
		if char >= 48 && char <= 57 {
			digit := int(char - '0') // Convert ASCII to digit value
			numarray = append(numarray, digit)
		}
	}
	if len(numarray) >= 1 {
		return fmt.Sprint(numarray[0]) + fmt.Sprint(numarray[len(numarray)-1])
	} else {
		return "0"
	}
}
