package arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertSumIsCorrect := func(t *testing.T, want int, given []int) {
		got := Sum(given)
		if got != want {
			t.Errorf("got %d want %d given: %v", got, want, given)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		assertSumIsCorrect(t, 15, []int{1, 2, 3, 4, 5})
	})
	t.Run("collection of any size", func(t *testing.T) {
		assertSumIsCorrect(t, 6, []int{1, 2, 3})
	})
}

func TestSumAll(t *testing.T) {
	assertSumIsCorrect := func(t *testing.T, want []int, given ...[]int) {
		got := SumAll(given...)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	assertSumIsCorrect(t, []int{7, 1}, []int{3, 4}, []int{-2, 3})
	assertSumIsCorrect(t, []int{0, 0}, []int{9, -9}, []int{1, 4, 6, -11})
}

func TestSumAllTails(t *testing.T) {
	assertSumIsCorrect := func(t *testing.T, want []int, given ...[]int) {
		got := SumAllTails(given...)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	assertSumIsCorrect(t, []int{7, 1, 4}, []int{91, 3, 4}, []int{0, 1, 0}, []int{-4912, 102, -98})
	assertSumIsCorrect(t, []int{0, 0, 0, 0}, []int{}, []int{}, []int{}, []int{})
}
