package main

import (
	"fmt"
	"learnToLeecode/internal/array/twoSum"
)

func main() {
	var set = []int{5, 3, 6, 7, 20}
	var result = twoSum.TwoSum(set, 13)
	fmt.Println(result)
}
