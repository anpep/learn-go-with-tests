package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(9, -3)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {
	assertSumIsCorrect := func(t *testing.T, sum interface{}, expected int) {
		t.Helper()
		if sum != expected {
			t.Errorf("expected %d but got %d", expected, sum)
		}
	}

	t.Run("test 2 + 2 = 4", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertSumIsCorrect(t, sum, expected)
	})
	t.Run("test 7 + (-9) = -2", func(t *testing.T) {
		sum := Add(7, -9)
		expected := -2
		assertSumIsCorrect(t, sum, expected)
	})
}
