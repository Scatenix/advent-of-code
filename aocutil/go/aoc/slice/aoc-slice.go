package slice

import "strconv"

func ReplaceRange[T any](s []T, i int, v []T) []T{
	for x := 0; x < len(v); x++ {
		s[i+x] = v[x]
	}
	return s
}

func Atoi(input []string) []int {
	ints := make([]int, len(input));
	for i, s := range input {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ints[i] = n
	}
	return ints
}

func Atoi64(input []string) []int64 {
	ints := make([]int64, len(input));
	for i, s := range input {
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		ints[i] = n
	}
	return ints
}

func RemoveZeros(s []int) []int {
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		if s[i] == 0 {
			s = RemoveIndexImpure(s, i)
			sLen--
			i--
		}
	}
	return s
}

func DeepCopy[T any](s []T) []T {
	dc := make([]T, len(s))
	copy(dc, s)
	return dc
}

func DeepCopy2D[T any](s [][]T) [][]T {
	new := make([][]T, len(s))
	for i := range s {
		new[i] = make([]T, len(s[i]))
		copy(new[i], s[i])
	}
	return new
}

func DeepCopy3D[T any](original [][][]T) [][][]T {
	new := make([][][]T, len(original))
	for i := range original {
		new[i] = make([][]T, len(original[i]))
		for j := range original[i] {
			new[i][j] = make([]T, len(original[i][j]))
			copy(new[i][j], original[i][j])
		}
	}
	return new
}

// =============================================================================================================
// A little deep dive into properly removing an item at index from a slice (vector)
// =============================================================================================================

// RemoveIndex is a proper pure removeIndex function which does not modify the input but rather creates a deep copy
// This is what one would expect from this function.
// However, if the input slice is also used for the output slice, a version without deep copy can
// be a better choice performance and space wise
func RemoveIndex(s []int, index int) []int {
	// create a completely new slice with the length of the input slice
	deepCopySlice := make([]int, 0, len(s)-1)
	// append the input slice until excluding index to the new slice
	deepCopySlice = append(deepCopySlice, s[:index]...)
	// append the input slice from index+1 to the end to the new slice
	return append(deepCopySlice, s[index+1:]...)
}

// RemoveIndexImpure shallow copy only!!!
// Assumption: Use only if the input slice will also be used for the returned slice (i.e. mySlice = removeIndex(mySlice, 4)
// ----------------------
// Since the slice (vector) is already implicitly a pointer (because all arrays are just that),
// This function is essentially impure by default.
// the original slice s will be modified as well, leading to drastic side effects,
// when trying to keep the original slice as is and using the output for a new slice.
// A deep copy is required here for a expected function
func RemoveIndexImpure(s []int, index int) []int {
	// needed instead if the slice type is a pointer or a struct with pointer fields to avoid memory leaks:
	//copy(a[i:], a[i+1:])
	//a[len(a)-1] = nil // or the zero value of T
	//a = a[:len(a)-1]

	// Fine for primitives
	return append(s[:index], s[index+1:]...)
}