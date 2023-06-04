package main

import (
	"fmt"
	// "learnToLeecode/internal/array/twoSum"
	"learnToLeecode/internal/hashTable/twoSum"
	"learnToLeecode/internal/math/palindroMenumber"
	"learnToLeecode/internal/sort/bubbleSort"
	"learnToLeecode/internal/string/bigAdd"
	"learnToLeecode/internal/string/reverseStr"
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

	set4 := []int{6, 5, 4, 1, 2, 3}
	result4 := bubbleSort.BubbleSort(set4)
	fmt.Println(result4)

	result5 := bigAdd.BigAdd("9123467", "33")
	fmt.Println(result5)

	result6 := reverseStr.ReverseStr("123456789")
	fmt.Println(result6)
}
