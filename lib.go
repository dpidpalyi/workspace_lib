// Package workspace_lib represents functions to work with numbers
// v1 use generic
//
//[MathIsFun]: https://www.mathisfun.com/numbers/
package workspace_lib

import "golang.org/x/exp/constraints"

// Number represents constraints for numbers
type Number interface {
	constraints.Integer | constraints.Float
}

// AddNums sumarize two numbers
func AddNums[T Number](a, b T) T {
	return a + b
}

// SubNums return difference from two numbers
func SubNums[T Number](a, b T) T {
	return a - b
}
