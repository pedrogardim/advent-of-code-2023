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

	res := possibleGamesWithinLimit(arr)

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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
