package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read the file content
	file, err := os.ReadFile("input5.txt")
	if err != nil {
		panic(err)
	}
	content := string(file)

	// Split content by double newline
	x := strings.Split(content, "\n\n")
	if len(x) != 8 {
		panic("Expected 8 sections in input file")
	}

	// Initialize arrays from input sections
	seedTemp := strings.Split(x[0], "\n")
	seedToSoilTemp := strings.Split(x[1], "\n")
	soilToFertTemp := strings.Split(x[2], "\n")
	fertToWaterTemp := strings.Split(x[3], "\n")
	waterToLightTemp := strings.Split(x[4], "\n")
	lightToTempTemp := strings.Split(x[5], "\n")
	tempToHumTemp := strings.Split(x[6], "\n")
	humToLocTemp := strings.Split(x[7], "\n")

	// Convert string arrays to integer arrays
	seed := getArray(seedTemp)
	seedToSoil := getArray(seedToSoilTemp)
	soilToFert := getArray(soilToFertTemp)
	fertToWater := getArray(fertToWaterTemp)
	waterToLight := getArray(waterToLightTemp)
	lightToTemp := getArray(lightToTempTemp)
	tempToHum := getArray(tempToHumTemp)
	humToLoc := getArray(humToLocTemp)

	// Create maps
	seedToSoilMap := getMap(seedToSoil)
	soilToFertMap := getMap(soilToFert)
	fertToWaterMap := getMap(fertToWater)
	waterToLightMap := getMap(waterToLight)
	lightToTempMap := getMap(lightToTemp)
	tempToHumMap := getMap(tempToHum)
	humToLocMap := getMap(humToLoc)

	// Calculate Location Map
	locMap := make(map[int]int)
	for i := 0; i < len(seed); i++ {
		for j := 0; j < len(seed[i]); j++ {
			seed := seed[i][j]
			con1, ok := seedToSoilMap[seed]
			if !ok {
				con1 = seed
			}
			con2, ok := soilToFertMap[con1]
			if !ok {
				con2 = con1
			}
			con3, ok := fertToWaterMap[con2]
			if !ok {
				con3 = con2
			}
			con4, ok := waterToLightMap[con3]
			if !ok {
				con4 = con3
			}
			con5, ok := lightToTempMap[con4]
			if !ok {
				con5 = con4
			}
			con6, ok := tempToHumMap[con5]
			if !ok {
				con6 = con5
			}
			locMap[seed], ok = humToLocMap[con6]
			if !ok {
				locMap[seed] = con6
			}
			// fmt.Println(locMap[seed])
		}
	}

	// Find and print the key with the lowest value in locMap

	fmt.Println(locMap[getLowestLoc(locMap)])
}

func getArray(arr []string) [][]int {
	result := make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		strValues := strings.Split(arr[i], " ")
		result[i] = make([]int, len(strValues))
		for j := 0; j < len(strValues); j++ {
			value, err := strconv.Atoi(strValues[j])
			if err != nil {
				panic(err)
			}
			result[i][j] = value
		}
	}
	return result
}

func getMap(arr [][]int) map[int]int {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		destinationRangeStart := arr[i][0]
		sourceRangeStart := arr[i][1]
		limitRange := arr[i][2]
		destinationArray := getRangeArray(destinationRangeStart, limitRange)
		sourceRangeArray := getRangeArray(sourceRangeStart, limitRange)
		for j := 0; j < len(destinationArray); j++ {
			m[sourceRangeArray[j]] = destinationArray[j]
		}
	}
	return m
}

func getRangeArray(start int, limit int) []int {
	result := make([]int, limit)
	for i := 0; i < limit; i++ {
		result[i] = start + i
	}
	return result
}

func getLowestLoc(loc map[int]int) int {
	minValue := math.MaxInt32
	var minKey int
	for key, value := range loc {
		if value < minValue {
			minValue = value
			minKey = key
		}
	}
	return minKey
}
