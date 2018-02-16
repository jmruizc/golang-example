package functionsutil

import "math"

// A function can take zero or more arguments, type comes after the variable name.
// When two or more consecutive named function parameters share a type, you can omit the type from all but the last
func SumIntegers(x, y int) int {
	return x + y
}

// Multiple results
func Swapstrings(x, y string) (string, string) {
	return y, x
}

//Named return values  are treated as variables defined at the top of the function
// These names should be used to document the meaning of the return values.
// A return statement without arguments returns the named return values. This is known as a "naked" return.
// Naked return statements should be used only in short functions
func Sincos(x float64) (sin, cos float64) {
	sin = math.Sin(x)
	cos = math.Cos(x)
	return
}