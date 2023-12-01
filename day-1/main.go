// Advent of Code 2023 - Day 1
// https://adventofcode.com/2023/day/1
// soli deo gloria

package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var spelledNums = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

// REFACTOR ME
var overlappingCases = map[string]string{
	"oneight":   "18",
	"twone":     "21",
	"threeight": "38",
	"fiveight":  "58",
	"sevenine":  "79",
	"eightwo":   "82",
	"eighthree": "83",
	"nineight":  "98",
}

func main() {
	data, err := os.ReadFile("./input.txt")
	check(err)

	str := string(data)

	for key, value := range overlappingCases {
		str = strings.ReplaceAll(str, key, value)
	}

	for i, numWord := range spelledNums {
		str = strings.ReplaceAll(str, numWord, strconv.Itoa(i+1))
	}

	re := regexp.MustCompile("[a-zA-Z]+")
	str = re.ReplaceAllString(str, "")

	arr := strings.Split(str, "\n")

	sum := 0

	for _, word := range arr {
		if len(word) == 0 {
			break
		}
		first := string(word[0])
		last := string(word[len(word)-1])
		concat := first + last
		parsed, err := strconv.Atoi(concat)
		check(err)
		sum += parsed
	}

	fmt.Println(sum)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
