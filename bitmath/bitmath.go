/*
########################
#                      #
#   WORK IN PROGRESS   #
#                      #
#      DO NOT USE      #
#                      #
########################

it doesn't even run or compile
*/

// Package bitmath provides tools to serialize and deserialize basic mathematical
// functions that are represented by readable strings into/from binary. It also
// provides a tool to interpret these string functions.
package bitmath

/*
	Example bitmath functions:

	a,b,c=a+5-b*c
	g, e, p=(p+e/g)/2
	x = 1/x

	Supported:
	* Multiple arguments
	* Arithmetic
	* Parentheses
	* PEMDAS without the E {lolwut}

	Not supported yet:
	* Multicharacter parameters
	* Float literals
	* Moduli, exponents, trig, etc.
	* Basically anything that isn't deductible from the examples and is not
		in the list of supported things
	* A funny joke to go in the "not supported yet" list
*/

var SyntaxError = errors.New("There is a syntax error in the function")

func Interpret(fn string, params float64...) (float64, error) {
	parammap := map[string]int
	var pastArrow bool

	for _, r := range fn {
		if r == ' '
	}
}
