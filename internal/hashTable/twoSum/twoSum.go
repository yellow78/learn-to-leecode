package twoSum

func TwoSum(nums []int, target int) []int {
	save := make(map[int]int)
	for key, value := range nums {
		var find = target - value
		findKey, ok := save[find]
		if ok == true {
			return []int{findKey, key}
		}
		save[value] = key
	}

	return []int{}
}
