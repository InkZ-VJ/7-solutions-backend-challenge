package main

import (
	"fmt"
)

func main() {
	testCases := []string{"LLRR=", "==RLL", "=LLRR", "RRL=R", "LLRLLR"}
	fmt.Println("===========Testing Example===========")

	for _, test := range testCases {
		fmt.Println("Input: ", test)
		_ = encode(test)
	}

	fmt.Println("===========Try It BY Yourself===========")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return
	}
	encode(input)
}
