// The arrays_and_slices package implements common routines for wokring with arrays and slices.
package arrays_and_slices

// The Sum function takes an array of integers and returns the result of adding them together
func Sum(integers []int) (sum int) {
	for _, integer := range integers {
		sum += integer
	}
	return
}

// The SumAll function takes multiple arrays of integers and returns an array with the results of adding the numbers
// in each array
func SumAll(integerArrays ...[]int) (sums []int) {
	for _, array := range integerArrays {
		sums = append(sums, Sum(array))
	}
	return
}

// The SumAllTails function takes multiple arrays of integers and returns an array with the results of adding every
// number but the first in each array
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
