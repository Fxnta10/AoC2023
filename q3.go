package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main4() {
	sum := 0
	file, err := os.ReadFile("testinput3.txt")
	content := string(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(content, "\n")
	//also split by every charectar to make it easier
	var charArray [][]string
	for _, str := range lines {
		chars := strings.Split(str, "")
		charArray = append(charArray, chars)
	}
	for i := 0; i < len(charArray); i++ {
		for j := 0; j < len(charArray[0]); j++ {
			var stringnumber string
			if checkifnumber(charArray[i][j]) == true {

				for charArray[i][j] != "." {
					stringnumber += charArray[i][j]
					j++
				}
				num, err := strconv.Atoi(stringnumber)
				if err != nil {
					return
				}
				cords := getAdjacentCoordinates(i, j, charArray)
				if checkPartNumber(cords, charArray) == true {
					sum += num
				}
			}
			//check j value before a afte debugging in this
		}
	}
	fmt.Println(sum)
}

// adjacent to a symbol is a part number(return true )
func checkPartNumber(cords [][2]int, lines2d [][]string) bool {
	for i := 0; i < len(lines2d); i++ {
		for j := 0; j < len(lines2d[0]); j++ {
			if existinCords(cords, i, j) == true {
				if lines2d[i][j] != "." {
					return true
				}
			}
		}
	}
	return false
}

func existinCords(cords [][2]int, i int, j int) bool {
	for _, cord := range cords {
		if cord[0] == i && cord[1] == j {
			return true
		}
	}
	return false
}

func checkifnumber(char string) bool {
	_, err := strconv.Atoi(char)
	return err == nil
}

func getAdjacentCoordinates(row, col int, grid [][]string) [][2]int {
	var adjacentCoords [][2]int
	directions := [][]int{
		{row - 1, col - 1},
		{row - 1, col + 1},
		{row + 1, col - 1},
		{row + 1, col + 1},
		{row - 1, col},
		{row + 1, col},
		{row, col - 1},
		{row, col + 1},
	}

	numRows := len(grid)
	numCols := len(grid[0]) //assuming the chararray has same columns

	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]

		// Check if the new coordinates are within bounds
		if newRow >= 0 && newRow < numRows && newCol >= 0 && newCol < numCols {
			adjacentCoords = append(adjacentCoords, [2]int{newRow, newCol})
		}
	}

	return adjacentCoords
}
