package main

import (
	"fmt"
	"strconv"
)

func encode(input string) string {
	var ans string
	stage := make([]int, len(input)+1)
	r := 0
	stage[0] = 0

	for i, s := range input {
		j := i + 1
		switch s {
		case '=':
			stage[j] = stage[j-1]

		case 'R':
			r = j // change to previous "R"
			stage[j] = stage[j-1] + 1

		case 'L':
			stage[j] = 0
			for n := j - 1; n > r; n-- {
				stage[n]++
			}
			if stage[r+1] >= stage[r] {
				stage[r]++
			}
		default:
			continue
		}
		// fmt.Println(string(s), stage)
	}
	ans = sliceOfStringToString(stage)
	fmt.Println("Encoded Answer: ", ans)

	return ans
}

func sliceOfStringToString(stage []int) string {
	var target string

	for _, num := range stage {
		s := strconv.Itoa(num) //Int to string
		target += s
	}
	return target
}
