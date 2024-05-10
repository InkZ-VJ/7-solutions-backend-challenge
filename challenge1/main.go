package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Challenge 1: Max-length")
	fmt.Println("---------Example---------")
	input_example := [][]int{{59}, {73, 41}, {52, 40, 53}, {26, 53, 6, 34}}
	fmt.Printf("example input = %v\n", input_example)
	fmt.Println("expected output = 237")

	result_example := maxSum(input_example)
	fmt.Println("My answers maximum sum: ", result_example)

	fmt.Println("---------Hard---------")
	fmt.Println("expected output = 7273")

	data, err := os.ReadFile("../files/hard.json")
	if err != nil {
		fmt.Println("Error to read files with error:", err)
		return
	}

	var input_hard [][]int

	err = json.Unmarshal(data, &input_hard)
	if err != nil {
		fmt.Println("Error to decode json", err)
		return
	}
	result_hard := maxSum(input_hard)
	fmt.Println("My answers maximum sum: ", result_hard)
}
