package bubbleSort

func BubbleSort(input []int) []int {
	sw := 0
	run := 0
	flag := false
	for run < len(input) {
		for i := 0; i < len(input); i++ {
			if i < len(input)-1 && input[i] > input[i+1] {
				sw = input[i+1]
				input[i+1] = input[i]
				input[i] = sw
				flag = true
			}
		}

		if flag == false {
			return input
		}

		flag = false
		run++
	}
	return input
}
