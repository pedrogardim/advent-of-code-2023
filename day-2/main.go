// Advent of Code 2023 - Day 2
// https://adventofcode.com/2023/day/2
// soli deo gloria

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	data, err := os.ReadFile("./input.txt")
	check(err)
	str := string(data)
	arr := strings.Split(str, "\n")

	// res := possibleGamesWithinLimit(arr)
	res := leastCubesPossible(arr)

	fmt.Println(res)
}

func possibleGamesWithinLimit(lines []string) int {

	limits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	possibleGamesIdSum := 0

	for i, gameString := range lines {
		isGamePossible := true
		gameId := i + 1
		gameString = strings.Split(gameString, ":")[1]
		roundStringArr := strings.Split(gameString, ";")
		for _, round := range roundStringArr {
			for _, comb := range strings.Split(round, ",") {
				splitComb := strings.Split(comb, " ")
				amount, err := strconv.Atoi(splitComb[1])
				color := splitComb[2]
				check(err)
				if amount > limits[color] {
					isGamePossible = false
				}
			}
		}
		if isGamePossible {
			possibleGamesIdSum += gameId
		}
	}

	return possibleGamesIdSum
}

func leastCubesPossible(lines []string) int {
	leastCubesPowerSum := 0

	for _, gameString := range lines {
		maxCubeEntry := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		gameString = strings.Split(gameString, ":")[1]
		roundStringArr := strings.Split(gameString, ";")
		for _, round := range roundStringArr {
			for _, comb := range strings.Split(round, ",") {
				splitComb := strings.Split(comb, " ")
				amount, err := strconv.Atoi(splitComb[1])
				color := splitComb[2]
				check(err)
				fmt.Println(amount, color)
				if maxCubeEntry[color] < amount {
					maxCubeEntry[color] = amount
				}
			}
		}

		power := 1

		for _, value := range maxCubeEntry {
			power *= value
		}

		leastCubesPowerSum += power
	}

	return leastCubesPowerSum
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

