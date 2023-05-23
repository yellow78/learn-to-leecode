package main

import (
	"fmt"
	// "learnToLeecode/internal/array/twoSum"
	"learnToLeecode/internal/hashTable/twoSum"
	"learnToLeecode/internal/math/palindroMenumber"
	"learnToLeecode/internal/string/romanToInteger"
)

func main() {
	var set = []int{3, 3}
	var result = twoSum.TwoSum(set, 6)
	fmt.Println(result)

	set2 := "MCMXCIV"
	result2 := romanToInteger.RomanToInt(set2)
	fmt.Println(result2)

	set3 := 121
	result3 := palindroMenumber.PalindromeNumber(set3)
	fmt.Println(result3)
}
