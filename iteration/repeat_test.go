package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeatedString := Repeat("piri", 2)
	fmt.Println("Salsita " + repeatedString)
	// Output: Salsita piripiri
}

func TestRepeat(t *testing.T) {
	assertRepeatIsCorrect := func(t *testing.T, repeated string, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	assertRepeatIsCorrect(t, Repeat("A", 5), "AAAAA")
	assertRepeatIsCorrect(t, Repeat("Hello ", 2), "Hello Hello ")
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("Hello, world!", 10)
	}
}
