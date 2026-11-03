package fizzbuzz_test

import (
	"fmt"
	"testing"

	"github.com/augiraud/go-template/internal/fizzbuzz"
)

func TestOf(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{n: 1, want: "1"},
		{n: 3, want: "Fizz"},
		{n: 5, want: "Buzz"},
		{n: 15, want: "FizzBuzz"},
		{n: 98, want: "98"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := fizzbuzz.Of(tt.n); got != tt.want {
				t.Errorf("Of(%d) = %q, want %q", tt.n, got, tt.want)
			}
		})
	}
}

func ExampleOf() {
	fmt.Println(fizzbuzz.Of(9))
	// Output: Fizz
}
