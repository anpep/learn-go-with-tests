// Package iteration implements routines using loops.
package iteration

// The Repeat function takes a string and an integer and returns a new string which is the original one repeated as
// many times as specified by the times parameter.
func Repeat(str string, times uint) (repeated string) {
	for n := uint(0); n < times; n++ {
		repeated += str
	}
	return
}
