package twoSum

func TwoSum(nums []int, target int) []int {
	for key, value := range nums {
		for key2, value2 := range nums {
			if key == key2 {
				continue
			}

			var getsum int
			getsum = value + value2
			if getsum == target {
				return []int{key, key2}
			}
		}
	}

	return []int{}
}
