package dynamicprogramming

import "testing"

func TestCountOrderedOneOrTwoStepWays(t *testing.T) {
	tests := []struct {
		name  string
		steps int
		want  int
	}{
		{name: "negative", steps: -1, want: 0},
		{name: "empty sequence", steps: 0, want: 1},
		{name: "one step", steps: 1, want: 1},
		{name: "two steps", steps: 2, want: 2},
		{name: "several transitions", steps: 5, want: 8},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := CountOrderedOneOrTwoStepWays(test.steps); got != test.want {
				t.Fatalf("CountOrderedOneOrTwoStepWays(%d) = %d, want %d", test.steps, got, test.want)
			}
		})
	}
}
