// Package fizzbuzz is placeholder code that shows where things go. Replace it with yours.
package fizzbuzz

import "strconv"

const (
	fizz = 3
	buzz = 5
)

// Of returns "Fizz", "Buzz", "FizzBuzz" or the number itself as text.
func Of(n int) string {
	switch {
	case n%(fizz*buzz) == 0:
		return "FizzBuzz"
	case n%fizz == 0:
		return "Fizz"
	case n%buzz == 0:
		return "Buzz"
	default:
		return strconv.Itoa(n)
	}
}
