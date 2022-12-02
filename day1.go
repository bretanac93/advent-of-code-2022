package main

import (
	"fmt"
	"sort"
	"strconv"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	data := ReadFile("./input/day1.txt")

	fmt.Println(GreatestCalories(data))
	fmt.Println(ThreeGreatestCalories(data))
}

func GreatestCalories(data []string) int {
	max := 0
	sum := 0

	for _, item := range data {
		if item == "" {
			max = intMax(max, sum)
			sum = 0
		} else {
			curr, _ := strconv.Atoi(item)
			sum += curr
		}
	}

	return max
}

func ThreeGreatestCalories(data []string) int {
	sums := make([]int, 0)
	sum := 0

	for _, item := range data {
		if item != "" {
			curr, _ := strconv.Atoi(item)
			sum += curr
		} else {
			sums = append(sums, sum)
			sum = 0
		}
	}

	sort.Ints(sums)

	l := len(sums) - 1

	return sums[l] + sums[l-1] + sums[l-2]
}

func intMax(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
