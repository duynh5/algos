package dynamicprogramming

// CountOrderedOneOrTwoStepWays returns the number of ordered sequences of
// one-step and two-step moves that sum to steps. Negative totals have no valid
// sequence; the empty sequence is the single way to make a total of zero.
func CountOrderedOneOrTwoStepWays(steps int) int {
	if steps < 0 {
		return 0
	}
	if steps < 2 {
		return 1
	}

	twoBack, oneBack := 1, 1
	for current := 2; current <= steps; current++ {
		twoBack, oneBack = oneBack, twoBack+oneBack
	}
	return oneBack
}
