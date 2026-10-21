//go:build ignore

// Difficulty: Easy
// Source: leetcode
// Problem ID: 88

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duyserveacc/algos/internal/testrunner"
)

/*
======================================================================
PASTE THE COMPLETE PROBLEM BELOW THIS LINE

Source: LeetCode
Problem ID: 88
Source URL:
../CONCEPTS_LINKS.md#leetcode-source-0088

You are given two integer arrays nums1 and nums2, sorted in
non-decreasing order, and two integers m and n, representing the
number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing
order.

The final sorted array should not be returned by the function, but
instead be stored inside the array nums1. To accommodate this, nums1
has a length of m + n, where the first m elements denote the elements
that should be merged, and the last n elements are set to 0 and should
be ignored. nums2 has a length of n.

Example 1:
Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with the underlined elements
coming from nums1.

Example 2:
Input: nums1 = [1], m = 1, nums2 = [], n = 0
Output: [1]
Explanation: The arrays we are merging are [1] and [].
The result of the merge is [1].

Example 3:
Input: nums1 = [0], m = 0, nums2 = [1], n = 1
Output: [1]
Explanation: The arrays we are merging are [] and [1].
The result of the merge is [1].
Note that because m = 0, there are no elements in nums1. The 0 is only
there to ensure the merge result can fit in nums1.


Constraints:
nums1.length == m + n
nums2.length == n
0 <= m, n <= 200
1 <= m + n <= 200
-109 <= nums1[i], nums2[j] <= 109

Can you come up with an algorithm that runs in O(m + n) time?

PASTE THE COMPLETE PROBLEM ABOVE THIS LINE
======================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here.
// Name adapters after the problem, such as parseTwoSumInput. Put
// setup used only by the current attempt at the top of
// solveCurrentAttempt below.

// parseMergeSortedArrayInput decodes the two arrays and their
// populated lengths in the same argument order used by the judge.
func parseMergeSortedArrayInput(
	input string,
) ([]int, int, []int, int) {
	var nums1 []int
	var m int
	var nums2 []int
	var n int
	if err := testrunner.DecodeJSONArgs(
		input,
		&nums1,
		&m,
		&nums2,
		&n,
	); err != nil {
		panic(fmt.Errorf("parse merge sorted array input: %w", err))
	}
	return nums1, m, nums2, n
}

// formatMergeSortedArrayOutput serializes the mutated nums1 array in
// the statement's output form.
func formatMergeSortedArrayOutput(nums1 []int) string {
	encoded, err := json.Marshal(nums1)
	if err != nil {
		panic(fmt.Errorf("format merge sorted array output: %w", err))
	}
	return string(encoded)
}

// repeatedMergeSortedArrayCase keeps the two 200-element boundary
// fixtures compact while still passing their inputs through the
// normal JSON boundary.
func repeatedMergeSortedArrayCase(
	name string,
	m, n, nums1Value, nums2Value int,
) testrunner.Case {
	nums1 := make([]int, m+n)
	for index := 0; index < m; index++ {
		nums1[index] = nums1Value
	}
	nums2 := make([]int, n)
	for index := range nums2 {
		nums2[index] = nums2Value
	}

	want := make([]int, 0, m+n)
	for range m {
		want = append(want, nums1Value)
	}
	for range n {
		want = append(want, nums2Value)
	}

	encodedInput, err := json.Marshal([]any{nums1, m, nums2, n})
	if err != nil {
		panic(
			fmt.Errorf(
				"build merge sorted array case %q: %w",
				name,
				err,
			),
		)
	}
	return testrunner.Case{
		Name:  name,
		Input: string(encodedInput),
		Want:  formatMergeSortedArrayOutput(want),
	}
}

// sampleCases contains the examples that every registered solution
// must pass. Inputs encode nums1, m, nums2, and n as a positional
// JSON array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Example 1",
		Input: `[[1,2,3,0,0,0],3,[2,5,6],3]`,
		Want:  `[1,2,2,3,5,6]`,
	},
	{
		Name:  "Example 2",
		Input: `[[1],1,[],0]`,
		Want:  `[1]`,
	},
	{
		Name:  "Example 3",
		Input: `[[0],0,[1],1]`,
		Want:  `[1]`,
	},
	{
		Name:  "Only nums1 has multiple values",
		Input: `[[-3,0,7],3,[],0]`,
		Want:  `[-3,0,7]`,
	},
	{
		Name:  "Only nums2 has multiple values",
		Input: `[[0,0,0,0],0,[-9,-1,0,8],4]`,
		Want:  `[-9,-1,0,8]`,
	},
	{
		Name:  "All nums1 values precede nums2",
		Input: `[[1,2,0,0],2,[3,4],2]`,
		Want:  `[1,2,3,4]`,
	},
	{
		Name:  "All nums2 values precede nums1",
		Input: `[[3,4,0,0],2,[1,2],2]`,
		Want:  `[1,2,3,4]`,
	},
	{
		Name:  "Arrays alternate",
		Input: `[[1,3,5,0,0,0],3,[2,4,6],3]`,
		Want:  `[1,2,3,4,5,6]`,
	},
	{
		Name:  "Duplicate values span both arrays",
		Input: `[[1,2,2,0,0,0],3,[2,2,3],3]`,
		Want:  `[1,2,2,2,2,3]`,
	},
	{
		Name:  "Zero is data as well as padding",
		Input: `[[-1,0,0,0],2,[0,1],2]`,
		Want:  `[-1,0,0,1]`,
	},
	{
		Name:  "Negative values interleave",
		Input: `[[-5,-2,0,0,0],2,[-6,-2,4],3]`,
		Want:  `[-6,-5,-2,-2,4]`,
	},
	{
		Name:  "Value constraint extremes",
		Input: `[[-1000000000,0,0,0],2,[-1000000000,1000000000],2]`,
		Want:  `[-1000000000,-1000000000,0,1000000000]`,
	},
	repeatedMergeSortedArrayCase(
		"Maximum nums1 length",
		200,
		0,
		-1000000000,
		0,
	),
	repeatedMergeSortedArrayCase(
		"Maximum nums2 length",
		0,
		200,
		0,
		1000000000,
	),
}

// archiveReady is set to true only after the problem, samples,
// preserved current attempt, three agent solutions, analysis, and
// regression coverage are complete and verified.
const archiveReady = true

// Registered implementations run against the same cases.
var currentAttemptSolution = testrunner.Solution{
	Name: "current attempt",
	Run:  solveCurrentAttempt,
}

// agentSolutions contains three implementations authored during
// review: a direct approach, an improved approach, and the preferred
// approach.
var agentSolutions = []testrunner.Solution{
	{Name: "insert nums2 values in place", Run: solveByInsertion},
	{Name: "merge into auxiliary buffer", Run: solveByAuxiliaryMerge},
	{Name: "merge backward in place", Run: solveByBackwardMerge},
}

// solutions keeps the user's attempt separate while running every
// case against it and all three agent-authored approaches.
var solutions = append(
	[]testrunner.Solution{currentAttemptSolution},
	agentSolutions...,
)

// ===================================================================
// CURRENT ATTEMPT
//
// Preserve this function's algorithm body during review. Add
// agent-authored implementations in the separate section below, even
// when one uses the same algorithmic method.
// ===================================================================

// solveCurrentAttempt is the only function you need to change
// initially. It receives the exact sample input string and returns
// the expected output form. Parse with the problem-specific adapter
// above, keep the algorithm strongly typed, and format structural
// output with a problem-specific formatter. Leading and trailing
// whitespace is ignored during comparison; internal whitespace
// remains significant.
func solveCurrentAttempt(input string) string {
	nums1, m, nums2, n := parseMergeSortedArrayInput(input)
	_ = m
	_ = nums2
	_ = n

	i := 0
	j := 0

	arr := make([]int, 0, m+n)

	for i < len(nums1) {
		if i >= m {
			arr = append(arr, nums2[j:]...)
			break
		}
		if j >= n {
			arr = append(arr, nums1[i:m]...)
			break
		}
		nums1_val := nums1[i]
		nums2_val := nums2[j]
		if nums1_val <= nums2_val {
			arr = append(arr, nums1_val)
			i = i + 1
		} else {
			arr = append(arr, nums2_val)
			j = j + 1
		}
	}
	copy(nums1, arr)

	// Implement your algorithm with the typed values above, then
	// return
	// nums1.
	return formatMergeSortedArrayOutput(nums1)
}

// ===================================================================
// AGENT SOLUTIONS
//
// Add the direct, improved, and preferred implementations here during
// review. Keep parsing helpers and judge-owned types in the compact
// support block above.
// ===================================================================

// solveByInsertion inserts each nums2 value into the populated prefix
// of nums1.
func solveByInsertion(input string) string {
	nums1, m, nums2, n := parseMergeSortedArrayInput(input)
	populated := m

	for nums2Index := 0; nums2Index < n; nums2Index++ {
		insertAt := populated
		for insertAt > 0 && nums1[insertAt-1] > nums2[nums2Index] {
			nums1[insertAt] = nums1[insertAt-1]
			insertAt--
		}
		nums1[insertAt] = nums2[nums2Index]
		populated++
	}

	return formatMergeSortedArrayOutput(nums1)
}

// solveByAuxiliaryMerge builds the sorted result from left to right,
// then copies it back into nums1.
func solveByAuxiliaryMerge(input string) string {
	nums1, m, nums2, n := parseMergeSortedArrayInput(input)
	merged := make([]int, 0, m+n)
	nums1Index := 0
	nums2Index := 0

	for nums1Index < m && nums2Index < n {
		if nums1[nums1Index] <= nums2[nums2Index] {
			merged = append(merged, nums1[nums1Index])
			nums1Index++
		} else {
			merged = append(merged, nums2[nums2Index])
			nums2Index++
		}
	}
	merged = append(merged, nums1[nums1Index:m]...)
	merged = append(merged, nums2[nums2Index:n]...)
	copy(nums1, merged)

	return formatMergeSortedArrayOutput(nums1)
}

// solveByBackwardMerge uses nums1's spare suffix as the output area.
func solveByBackwardMerge(input string) string {
	nums1, m, nums2, n := parseMergeSortedArrayInput(input)
	mergeSortedArraysFromBack(nums1, m, nums2, n)
	return formatMergeSortedArrayOutput(nums1)
}

// merge is the judge-facing wrapper around the preferred
// implementation.
func merge(nums1 []int, m int, nums2 []int, n int) {
	mergeSortedArraysFromBack(nums1, m, nums2, n)
}

// mergeSortedArraysFromBack writes the largest remaining value into
// the last unfilled slot, so unread nums1 values are never
// overwritten.
func mergeSortedArraysFromBack(
	nums1 []int,
	m int,
	nums2 []int,
	n int,
) {
	nums1Index := m - 1
	nums2Index := n - 1
	writeIndex := m + n - 1

	for nums2Index >= 0 {
		if nums1Index >= 0 && nums1[nums1Index] > nums2[nums2Index] {
			nums1[writeIndex] = nums1[nums1Index]
			nums1Index--
		} else {
			nums1[writeIndex] = nums2[nums2Index]
			nums2Index--
		}
		writeIndex--
	}
}

/*
======================================================================
AGENT REVIEW AND COMPLEXITY ANALYSIS

Current attempt:
- What is correct: The attempt is correct for every valid input.
  Before each
comparison, arr is sorted and contains exactly the smallest i+j
consumed
values. Choosing the smaller head preserves that invariant; when one
input
  is exhausted, appending the other sorted suffix completes the merge.
- First failing assumption or bottleneck: No correctness assumption
  fails under
the stated constraints. The first optimization limit is the
m+n-element arr
allocation plus the final copy, even though nums1 already owns enough
space.
- Relevant concept and invariant: See ../CONCEPTS.md#two-pointers.
  Each pointer names the smallest unread value in its input.
- Concrete example: For nums1=[1,3,5,0,0,0], m=3, nums2=[2,4,6], n=3,
  (i,j,arr) changes from (0,0,[]) to (1,0,[1]), (1,1,[1,2]),
  (2,1,[1,2,3]), (2,2,[1,2,3,4]), and (3,2,[1,2,3,4,5]).
  Since i==m, [6] is appended and copied into nums1.
- Time complexity: Best, average, and worst are Theta(m+n): every
  value is
  appended and the full result is copied.
- Auxiliary-space complexity: Theta(m+n) for arr. JSON parsing and
  output also
  allocate Theta(m+n) in this workbench wrapper.
- Relationship to agent solutions: Solution 2 uses the same forward
  two-pointer
merge with an auxiliary result buffer. It is an independent,
cleaned-up
implementation of this attempt, not a different algorithmic approach.

Solution 1 - Insert nums2 values in place:
- Correctness: Before each outer iteration, nums1[:populated] is the
  sorted
merge of the original nums1 prefix and processed nums2 values.
Shifting every
larger value right opens the unique sorted insertion point; placing
the next
  nondecreasing nums2 value there re-establishes the invariant.
- Concrete example: With nums1=[3,4,0,0], m=2 and nums2=[1,2],
  inserting 1
shifts 4 to index 2 and 3 to index 1, producing [1,3,4,0]. Inserting 2
  shifts 4 to index 3 and 3 to index 2, producing [1,2,3,4].
- Time complexity: The merge work is best Theta(n) when every
  insertion is at
the end and worst Theta(m*n+n), or O((m+n)^2), when nums2 precedes
nums1.
Average time depends on the value distribution. Parsing and output add
  Theta(m+n), so end-to-end best time is Theta(m+n).
- Auxiliary-space complexity: O(1) for the algorithm. The workbench's
  decoded
  inputs and encoded output require Theta(m+n) memory.

Solution 2 - Merge into an auxiliary buffer (same approach as current
attempt):
- Correctness: merged is always sorted and contains exactly the
  consumed
values. The smaller unread head is the smallest value not yet emitted.
Once
an input ends, its counterpart's suffix is already sorted and no
smaller
  unconsumed value remains, so appending both suffixes is correct.
- Concrete example: For nums1=[1,2,3,0,0,0], m=3 and nums2=[2,5,6],
  pointer
pairs (0,0), (1,0), (2,0), and (2,1) append 1, 2 from nums1, 2 from
nums2,
and 3. nums1 is exhausted, so [5,6] is appended before the copy back.
- Time complexity: Best, average, and worst are Theta(m+n); all values
  are
  emitted once and the result is copied once.
- Auxiliary-space complexity: Theta(m+n) for merged, plus Theta(m+n)
  for the
  workbench adapters.

Solution 3 - Merge backward in place (preferred):
- Correctness: Before each iteration, nums1[writeIndex+1:] contains
  the final
sorted suffix. nums1[nums1Index] and nums2[nums2Index] are the largest
unread
values, so writing the larger one at writeIndex extends that suffix
without
overwriting unread nums1 data. When nums2 is exhausted, any remaining
nums1
  prefix is already in its final position.
- Concrete example: For nums1=[1,2,3,0,0,0], m=3 and nums2=[2,5,6],
(nums1Index,nums2Index,writeIndex)=(2,2,5) writes 6, then (2,1,4)
writes 5,
then (2,0,3) writes 3, and (1,0,2) writes nums2's 2. nums2 is
exhausted;
  nums1's [1,2] prefix stays in place, yielding [1,2,2,3,5,6].
- Time complexity: Best is Theta(n) when nums2 supplies every value
  written
before it is exhausted; worst is Theta(m+n). Average time depends on
the
value distribution. Parsing and output make the wrapper Theta(m+n)
overall.
- Auxiliary-space complexity: O(1) for the merge algorithm. The
  workbench
  adapters allocate Theta(m+n) for decoded inputs and encoded output.

AGENT REVIEW AND COMPLEXITY ANALYSIS END
======================================================================
*/

func main() {
	if err := testrunner.Run(
		os.Stdout,
		sampleCases,
		solutions,
	); err != nil {
		fmt.Fprintln(os.Stderr, "workbench:", err)
		os.Exit(1)
	}
}
