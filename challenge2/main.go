package main

import (
	"fmt"
)

func main() {
	testCases := []string{"LLRR=", "==RLL", "=LLRR", "RRL=R", "LLRLLR"}

	// max_num := maxStack1(testCases[1])
	// fmt.Println(max_num)

	for _, test := range testCases {
		max_index, max_num := maxStack1(test)
		fmt.Printf("Input: '%s'  max number: %d max index: %d\n", test, max_num, max_index)
	}

}
