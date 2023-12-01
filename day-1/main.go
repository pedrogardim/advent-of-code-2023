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

func main() {
	data, err := os.ReadFile("./input.txt")
	check(err)

	partOne(string(data))
}

func partOne(data string) {
	re := regexp.MustCompile("[a-zA-Z]+")
	str := re.ReplaceAllString(data, "")

	arr := strings.Split(str, "\n")

	sum := 0

	for i := range arr {
		if len(arr[i]) == 0 {
			continue
		}
		first := string(arr[i][0])
		last := string(arr[i][len(arr[i])-1])
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
