package romanToInteger

func RomanToInt(s string) int {
	result := 0
	len := len(s)
	chars := []byte(s)
	isMap := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	for k, v := range chars {
		getInt := isMap[v]
		if k < len-1 && getInt < isMap[chars[k+1]] {
			result -= getInt
		} else {
			result += getInt
		}
	}
	return result
}
