package math

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