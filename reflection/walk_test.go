package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		name     string
		input    any
		expected []string
	}{
		{
			name:     "struct with one string field",
			input:    struct{ name string }{name: "Ángel"},
			expected: []string{"Ángel"},
		},
		{
			name:     "struct with two string fields",
			input:    struct{ first, last string }{"First Name", "Last Name"},
			expected: []string{"First Name", "Last Name"},
		},
		{
			name: "struct with a non-string field",
			input: struct {
				name string
				age  int
			}{name: "Ángel", age: 21},
			expected: []string{"Ángel"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			results := []string{}
			Walk(c.input, func(input string) {
				results = append(results, input)
			})
			if !reflect.DeepEqual(results, c.expected) {
				t.Errorf("got %v want %v", results, c.expected)
			}
		})
	}
}
