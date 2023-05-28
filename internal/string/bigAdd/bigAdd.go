package bigAdd

import (
	"fmt"
	"strconv"
)

func BigAdd(a string, b string) string {
	lenA := len(a) - 1
	lenB := len(b) - 1
	byteA := []byte(a)
	byteB := []byte(b)
	if lenA == 0 {
		return b
	}

	if lenB == 0 {
		return a
	}

	lenMax := lenA
	if lenA <= lenB {
		lenMax = lenB
	}

	result := make([]byte, lenMax+1)
	carry := 0
	getAdd := 0
	var err error
	for i := lenMax; i >= 0; i-- {
		numB := 0
		numA := 0
		if lenA >= 0 {
			numA, err = strconv.Atoi(string(byteA[lenA]))
			if err != nil {
				fmt.Println("转换失败:", err)
				return ""
			}
		}

		if lenB >= 0 {
			numB, err = strconv.Atoi(string(byteB[lenB]))
			if err != nil {
				fmt.Println("转换失败:", err)
				return ""
			}
		}

		getAdd = numA + numB + carry
		if getAdd > 9 {
			carry = getAdd / 10
			getAdd = getAdd % 10
		} else {
			carry = 0
		}

		result[i] = byte(getAdd + '0')
		lenA--
		lenB--
		fmt.Println(result)
	}
	return string(result)
}
