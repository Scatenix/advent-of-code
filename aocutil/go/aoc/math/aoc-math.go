package math

import (
	"fmt"
	"strconv"
)

// Abs calculates the absolute value for an int
func Abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func Pow(a, b int) int {
	result := a
	for i := 0; i < b-1; i++ {
		result = result*a
	}
	return result
}

// ConcatInt64 concat two int64s like string concatenation
func ConcatInt64(a, b int64) int64 {
	ret, _ := strconv.ParseInt(strconv.FormatInt(a, 10)+strconv.FormatInt(b, 10), 10, 64)
	return ret
}

// IntToBaseStringWithPadding create a zero padded string for any number base (can be used to simply pad base 10 numbers with zeros)
func IntToBaseStringWithPadding(num, length, base int) string {
	binaryStr := strconv.FormatInt(int64(num), base)
	padding := length - len(binaryStr)
	if padding > 0 {
		binaryStr = fmt.Sprintf("%0*s", length, binaryStr)
	}
	return binaryStr
}