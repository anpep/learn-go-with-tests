// The arrays_and_slices package implements common routines for wokring with arrays and slices.
package arrays_and_slices

// The Sum function takes a slice of integers and returns the result of adding them together
func Sum(integers []int) (sum int) {
	for _, integer := range integers {
		sum += integer
	}
	return
}

func SumAll(integerArrays ...[]int) (sums []int) {
	for _, array := range integerArrays {
		sums = append(sums, Sum(array))
	}
	return
}

func SumAllTails(integerArrays ...[]int) (sums []int) {
	for _, array := range integerArrays {
		if len(array) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(array[1:]))
		}
	}
	return
}
