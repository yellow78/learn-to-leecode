package palindroMenumber

func PalindromeNumber(x int) bool {
	if x < 0 {
		return false
	}

	find := x
	reverse := 0

	for x != 0 {
		reverse = reverse*10 + x%10
		x = x / 10
	}

	if reverse == find {
		return true
	}
	return false
}
