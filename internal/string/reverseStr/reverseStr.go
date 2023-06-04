package reverseStr

func ReverseStr(input string) string {
	strlen := len(input)
	var out = make([]byte, strlen)
	run := strlen / 2
	out[run] = input[run]
	for i := 0; i < run; i++ {
		out[i] = input[strlen-i-1]
		out[strlen-i-1] = input[i]
	}

	return string(out)
}
