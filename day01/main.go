package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func part1() {
	content, _ := ioutil.ReadFile("input.txt")

	var lastWord string
	var increased int32

	for index, word := range strings.Split(string(content), "\n") {
		if index == 0 {
			lastWord = word
			continue
		}

		var lastWordInt, _ = strconv.Atoi(lastWord)
		var wordInt, _ = strconv.Atoi(word)

		if wordInt > lastWordInt {
			increased += 1
		}

		lastWord = word
	}

	fmt.Printf("increased %v\n", increased)
}

func part2() {
	content, _ := ioutil.ReadFile("input.txt")

	var three1 int = 0
	var three2 int = 0
	var three3 int = 0
	var lastThree int = 0
	var increased = 0

	for index, word := range strings.Split(string(content), "\n") {
		wordInt, _ := strconv.Atoi(word)

		if index == 0 {
			three1 += wordInt
			continue
		}
		if index == 1 {
			three1 += wordInt
			three2 += wordInt
			continue
		}

		three1 += wordInt
		three2 += wordInt
		three3 += wordInt

		if index == 2 {
			lastThree = three1
			three1 = 0
			continue
		}

		fmt.Printf("looping %v %v %v with index %v\n", three1, three2, three3, index)
		if index % 3 == 0 {
			if three1 > lastThree {
				increased += 1
			}

			lastThree = three1
			three1 = 0
		}
		if index % 3 == 1 {
			if three2 > lastThree {
				increased += 1
			}

			lastThree = three2
			three2 = 0
		}
		if index % 3 == 2 {
			if three3 > lastThree {
				increased += 1
			}

			lastThree = three3
			three3 = 0
		}
	}

	fmt.Printf("increased %v\n", increased)
}

func main() {
	part2()
}
