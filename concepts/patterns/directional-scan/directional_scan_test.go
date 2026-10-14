package directionalscan

import "testing"

func TestSumWithSmallerLeftValuesSubtracted(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   int
	}{
		{name: "empty", want: 0},
		{name: "single value", values: []int{5}, want: 5},
		{name: "descending values are additive", values: []int{1000, 100, 10, 1}, want: 1111},
		{name: "smaller left value is subtractive", values: []int{10, 100}, want: 90},
		{name: "mixed local contributions", values: []int{1000, 100, 1000, 10, 100, 1, 5}, want: 1994},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := SumWithSmallerLeftValuesSubtracted(test.values); got != test.want {
				t.Fatalf("SumWithSmallerLeftValuesSubtracted(%v) = %d, want %d", test.values, got, test.want)
			}
		})
	}
}
