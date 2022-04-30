package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Ángel")

	got := buffer.String()
	want := "Hello, Ángel!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
