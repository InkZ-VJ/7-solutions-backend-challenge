package main

// import (
// 	"fmt"
// 	"strconv"
// )

func maxStack1(input string) (int, int) {
	max_stack, max_index, count := 0, 0, 0
	var top_stack rune

	for i, s := range input {
		if s == '=' {
			continue
		}
		if count == 0 {
			top_stack = s
		}

		// fmt.Println("string: ", strconv.QuoteRune(s))

		switch s == top_stack {
		case true:
			count++
		case false:
			count--
		}

		// get only first max index
		if count > max_stack {
			max_stack = count
			max_index = i
		}

		if count == 0 {
			count++
			top_stack = s
		}
		// fmt.Println("current top stack: ", strconv.QuoteRune(top_stack))
	}
	return max_index, max_stack
}
