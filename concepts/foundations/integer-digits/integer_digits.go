package integerdigits

import "math"

// ReverseNonNegative reverses a nonnegative decimal integer. ok is false when
// value is negative or when its reversal would exceed int64.
func ReverseNonNegative(value int64) (reversed int64, ok bool) {
	if value < 0 {
		return 0, false
	}

	for value > 0 {
		digit := value % 10
		if reversed > (math.MaxInt64-digit)/10 {
			return 0, false
		}
		reversed = reversed*10 + digit
		value /= 10
	}
	return reversed, true
}
