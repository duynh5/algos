package integerdigits

import "testing"

func TestReverseNonNegative(t *testing.T) {
	tests := []struct {
		name         string
		value        int64
		wantReversed int64
		wantOK       bool
	}{
		{name: "zero", value: 0, wantReversed: 0, wantOK: true},
		{name: "palindrome", value: 1221, wantReversed: 1221, wantOK: true},
		{name: "trailing zeroes disappear", value: 1200, wantReversed: 21, wantOK: true},
		{name: "negative rejected", value: -121, wantOK: false},
		{name: "overflow rejected", value: 8_085_774_586_302_733_229, wantOK: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reversed, ok := ReverseNonNegative(test.value)
			if reversed != test.wantReversed || ok != test.wantOK {
				t.Fatalf("ReverseNonNegative(%d) = (%d, %t), want (%d, %t)",
					test.value, reversed, ok, test.wantReversed, test.wantOK)
			}
		})
	}
}
