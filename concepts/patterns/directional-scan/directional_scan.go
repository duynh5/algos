package directionalscan

// SumWithSmallerLeftValuesSubtracted applies a local signed-sum rule: a value
// is subtracted when it is smaller than its immediate right neighbor and added
// otherwise.
func SumWithSmallerLeftValuesSubtracted(values []int) int {
	total := 0
	rightValue := 0

	for index := len(values) - 1; index >= 0; index-- {
		value := values[index]
		if value < rightValue {
			total -= value
		} else {
			total += value
		}
		rightValue = value
	}

	return total
}
