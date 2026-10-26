//go:build ignore

// Difficulty: Easy
// Source: leetcode
// Problem ID: 27

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/duyserveacc/algos/dsa/internal/testrunner"
)

/*
======================================================================
Remove Element

Source: LeetCode
Problem ID: 27
Source URL: https://leetcode.com/problems/remove-element/description/

Given an integer array nums and an integer val, remove all occurrences
of val in nums in-place. The order of the elements may be changed.
Then return the number of elements in nums which are not equal to val.

Consider the number of elements in nums which are not equal to val be
k, to get accepted, you need to do the following things:

- Change the array nums such that the first k elements of nums contain
  the elements which are not equal to val. The remaining elements of
  nums are not important as well as the size of nums.

- Return k.

Custom Judge:

The judge will test your solution with the following code:

int[] nums = [...]; // Input array
int val = ...; // Value to remove
int[] expectedNums = [...]; // The expected answer with correct
length.
// It is sorted with no values equaling val.

int k = removeElement(nums, val); // Calls your implementation

assert k == expectedNums.length;
sort(nums, 0, k); // Sort the first k elements of nums
for (int i = 0; i < actualLength; i++) {
assert nums[i] == expectedNums[i];
}

If all assertions pass, then your solution will be accepted.

Example 1:

Input: nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2,_,_]
Explanation: Your function should return k = 2, with the first two
elements of nums being 2.
It does not matter what you leave beyond the returned k (hence they
are underscores).

Example 2:

Input: nums = [0,1,2,2,3,0,4,2], val = 2
Output: 5, nums = [0,1,4,0,3,_,_,_]
Explanation: Your function should return k = 5, with the first five
elements of nums containing 0, 0, 1, 3, and 4.
Note that the five elements can be returned in any order.
It does not matter what you leave beyond the returned k (hence they
are underscores).

Constraints:

- 0 <= nums.length <= 100

- 0 <= nums[i] <= 50

- 0 <= val <= 100
======================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here.
// Name adapters after the problem. Put setup used only by the current
// attempt at the top of solveCurrentAttempt below.
//
// This completed problem archive follows the workbench template.
// Concept: ../CONCEPTS.md#two-pointers

func parseRemoveElementInput(input string) ([]int, int) {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse remove-element input: %w", err))
	}
	if len(args) != 2 {
		panic(
			fmt.Sprintf(
				"parse remove-element input: got %d arguments,"+
					" want 2",
				len(args),
			),
		)
	}
	var nums []int
	argumentError0 := json.Unmarshal(args[0], &nums)
	if argumentError0 != nil {
		panic(
			fmt.Errorf(
				"parse remove-element argument nums: %w",
				argumentError0,
			),
		)
	}
	var val int
	argumentError1 := json.Unmarshal(args[1], &val)
	if argumentError1 != nil {
		panic(
			fmt.Errorf(
				"parse remove-element argument val: %w",
				argumentError1,
			),
		)
	}
	return nums, val
}

// formatRemoveElementOutput applies the custom judge's
// order-insensitive comparison by sorting only the retained prefix
// before serializing it.
func formatRemoveElementOutput(k int, nums []int) string {
	retained := append([]int(nil), nums[:k]...)
	sort.Ints(retained)
	values := make([]string, len(nums))
	for index := range values {
		if index < k {
			values[index] = strconv.Itoa(retained[index])
		} else {
			values[index] = "_"
		}
	}
	return fmt.Sprintf(
		"%d, nums = [%s]",
		k,
		strings.Join(values, ","),
	)
}

func maximumRemoveElementCase() testrunner.Case {
	nums := make([]int, 100)
	retained := make([]int, 0, len(nums))
	for index := range nums {
		nums[index] = index % 51
		if nums[index] != 25 {
			retained = append(retained, nums[index])
		}
	}
	input, err := json.Marshal([]any{nums, 25})
	if err != nil {
		panic(
			fmt.Errorf("build maximum remove-element input: %w", err),
		)
	}
	formattedNums := make([]int, len(nums))
	copy(formattedNums, retained)
	return testrunner.Case{
		Name:  "Boundary: Maximum length with separated removals",
		Input: string(input),
		Want: formatRemoveElementOutput(
			len(retained),
			formattedNums,
		),
	}
}

// sampleCases contains the examples that every registered solution
// must pass. Inputs encode the judge arguments as a positional JSON
// array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `[[3,2,2,3],3]`,
		Want:  `2, nums = [2,2,_,_]`,
	},
	{
		Name:  "Official example 2",
		Input: `[[0,1,2,2,3,0,4,2],2]`,
		Want:  `5, nums = [0,0,1,3,4,_,_,_]`,
	},
	{
		Name:  "Additional: Empty array",
		Input: `[[],0]`,
		Want:  `0, nums = []`,
	},
	{
		Name:  "Additional: Every value removed",
		Input: `[[1,1],1]`,
		Want:  `0, nums = [_,_]`,
	},
	{
		Name:  "Additional: No value removed",
		Input: `[[1,2],3]`,
		Want:  `2, nums = [1,2]`,
	},
	{
		Name:  "Additional: Removals at beginning, middle, and end",
		Input: `[[2,1,2,3,2],2]`,
		Want:  `2, nums = [1,3,_,_,_]`,
	},
	{
		Name:  "Additional: Duplicate retained values",
		Input: `[[1,2,1,2,1],2]`,
		Want:  `3, nums = [1,1,1,_,_]`,
	},
	{
		Name:  "Boundary: Value exceeds element range",
		Input: `[[0,25,50],100]`,
		Want:  `3, nums = [0,25,50]`,
	},
	{
		Name:  "Boundary: Remove zero while retaining maximum element",
		Input: `[[0,50,0],0]`,
		Want:  `1, nums = [50,_,_]`,
	},
	maximumRemoveElementCase(),
}

// archiveReady is true after the current attempt, three agent
// solutions, analysis, and regression coverage are complete and
// verified.
const archiveReady = true

// Registered implementations run against the same cases.
var currentAttemptSolution = testrunner.Solution{
	Name: "current attempt",
	Run:  solveCurrentAttempt,
}

var agentSolutions = []testrunner.Solution{
	{
		Name: "direct repeated shifting",
		Run:  solveByRepeatedShifting,
	},
	{
		Name: "improved swap with active end",
		Run:  solveBySwappingWithActiveEnd,
	},
	{
		Name: "preferred stable write pointer",
		Run:  solveByStableWritePointer,
	},
}

var solutions = append(
	[]testrunner.Solution{currentAttemptSolution},
	agentSolutions...,
)

// ===================================================================
// CURRENT ATTEMPT
//
// Implement your solution here. Keep this body intact when requesting
// review.
// ===================================================================

func solveCurrentAttempt(input string) string {
	nums, val := parseRemoveElementInput(input)
	_ = nums
	_ = val
	occ := 0
	last_occ_i := make([]int, 0)

	for i, v := range nums {
		if v == val {
			last_occ_i = append(last_occ_i, i)
			occ += 1
		} else {
			if len(last_occ_i) > 0 {
				nums[last_occ_i[0]] = v
				last_occ_i = last_occ_i[1:]
				last_occ_i = append(last_occ_i, i)
			}
		}
	}

	// Implement your algorithm with the typed values above.
	return formatRemoveElementOutput(len(nums)-occ, nums)
}

// ===================================================================
// AGENT SOLUTIONS
//
// These functions remain separate from the preserved current attempt
// so each approach can be compared against the same cases.
// ===================================================================

func solveByRepeatedShifting(input string) string {
	nums, val := parseRemoveElementInput(input)
	retainedLength := len(nums)

	for index := 0; index < retainedLength; {
		if nums[index] != val {
			index++
			continue
		}

		copy(
			nums[index:retainedLength-1],
			nums[index+1:retainedLength],
		)
		retainedLength--
	}

	return formatRemoveElementOutput(retainedLength, nums)
}

func solveBySwappingWithActiveEnd(input string) string {
	nums, val := parseRemoveElementInput(input)
	activeEnd := len(nums)

	for index := 0; index < activeEnd; {
		if nums[index] != val {
			index++
			continue
		}

		activeEnd--
		nums[index] = nums[activeEnd]
	}

	return formatRemoveElementOutput(activeEnd, nums)
}

func solveByStableWritePointer(input string) string {
	nums, val := parseRemoveElementInput(input)
	write := 0

	for _, value := range nums {
		if value == val {
			continue
		}

		nums[write] = value
		write++
	}

	return formatRemoveElementOutput(write, nums)
}

/*
======================================================================
AGENT REVIEW AND COMPLEXITY ANALYSIS

Current attempt:
- What is correct: It counts every occurrence of val and treats each
  matching
index as a hole. Every later retained value is copied into the oldest
hole,
so the returned prefix contains exactly the non-val values and
preserves
their relative order. It passes empty, all-removed, no-removal,
duplicate,
  boundary-value, and maximum-length cases.
- First failing assumption or bottleneck: There is no correctness
  failure in
the covered input domain. The bottleneck is last_occ_i: retaining
every open
hole can consume O(n) auxiliary space even though only the next write
position is needed. The range loop always examines all n values, so
best,
average, and worst-case time are Theta(n). If r removals are pending,
the
  queue uses O(r) space, hence O(n) in the worst case.
- Relevant concept and invariant: After processing nums[:i+1], the
  prefix
nums[:i+1-occ] contains all retained values from that processed prefix
in
original order. last_occ_i identifies the remaining logical holes.
This is
the same stable-compaction goal and kept-prefix invariant as the
preferred
write-pointer solution, but the attempt represents every hole rather
than
  only the boundary after the kept prefix.
- Concrete example: For nums = [2,1,2,3,2], val = 2, index 0 records
  hole [0]
and occ = 1. Index 1 copies 1 to nums[0] and moves the logical hole to
[1],
giving [1,1,2,3,2]. Index 2 records another hole, [1,2], and occ = 2.
Index 3 copies 3 to nums[1] and moves that hole to the back, [2,3],
giving
[1,3,2,3,2]. Index 4 records the final hole, so occ = 3 and k = 5-3 =
2;
  nums[:2] is [1,3].

Solution 1 - Direct repeated shifting:
- Correctness: When nums[index] equals val, copying the active suffix
  one slot
left removes that occurrence from the first retainedLength positions.
The
same index is checked again because a new value moved into it.
Otherwise the
index advances. Thus, when index reaches retainedLength, every value
in the
active prefix is non-val and every original non-val value appears
once.
- Concrete example: For [3,2,2,3], val = 3, index 0 shifts [2,2,3]
  left,
producing [2,2,3,3] with retainedLength = 3. The two 2s advance index
to 2.
The 3 there reduces retainedLength to 2; no value needs copying. The
returned
  prefix is [2,2].
- Comparison with current attempt: Both preserve retained order.
  Repeated
shifting uses no queue, but it may move the same retained value many
times;
the current attempt moves each retained value at most once after a
hole.
- Time complexity: Theta(n + s), where s is the total number of
  shifted
elements. Best case Theta(n) with no removals or only trailing
removals;
average cost depends on removal positions; worst case Theta(n^2) when
many
  early removals repeatedly shift long suffixes.
- Auxiliary-space complexity: O(1), excluding the parsed input and the
  output
  formatting copy required by this local harness.

Solution 2 - Improved swap with active end:
- Correctness: nums[:activeEnd] is the unfinalized active region. A
  retained
value finalizes its current position. A val is replaced by the last
active
value while activeEnd shrinks; the replacement is checked before index
can
advance. Each iteration classifies exactly one active value, so the
final
prefix contains all and only retained values. Reordering is permitted.
- Concrete example: For [0,1,2,2,3,0,4,2], val = 2, indices 0 and 1
  advance.
At index 2, activeEnd becomes 7 and the trailing 2 is copied there, so
the
same index remains under review. activeEnd becomes 6 next and 4
replaces it.
Index 2 then advances. At index 3, activeEnd becomes 5 and 0 replaces
the 2.
  Scanning the remaining active positions ends with k = 5 and prefix
  [0,1,4,0,3].
- Comparison with current attempt: Both avoid repeated suffix shifts
  and run
in linear time. This approach uses the problem's permission to change
order
and needs no hole queue; the current attempt is stable and uses O(n)
space in
  the worst case.
- Time complexity: Theta(n) in the best, average, and worst cases.
  Every loop
iteration either advances index or decreases activeEnd, classifying
one
  previously active value.
- Auxiliary-space complexity: O(1), excluding harness parsing and
  formatting.

Solution 3 - Preferred stable write pointer:
- Correctness: Before each read, nums[:write] contains exactly the
  non-val
values from the processed input prefix, in their original order. A val
does
not change that prefix; a retained value is written at nums[write] and
then
extends it by one. write never exceeds the read position, so no unread
value
is overwritten. At the end, write is k and nums[:k] is the required
prefix.
- Concrete example: For [2,1,2,3,2], val = 2, write starts at 0. The
  first 2
is skipped. Reading 1 writes nums[0] = 1 and makes write = 1. The next
2 is
skipped. Reading 3 writes nums[1] = 3 and makes write = 2. The last 2
is
  skipped, so k = 2 and the retained prefix is [1,3].
- Comparison with current attempt: This expresses the same stable
  compaction
and kept-prefix invariant directly. The current attempt tracks a FIFO
of all
logical holes and counts removals; this version stores only write, the
first
position after the valid prefix, making the extra collection
unnecessary.
- Time complexity: Theta(n) in the best, average, and worst cases
  because each
  input value is read exactly once.
- Auxiliary-space complexity: O(1), excluding harness parsing and
  formatting.

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
