package math

// Abs calculates the absolute value for an int
func Abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}