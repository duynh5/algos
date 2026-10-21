//go:build ignore

// Difficulty: Easy
// Source: leetcode
// Problem ID: 26

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/duyserveacc/algos/internal/testrunner"
)

/*
======================================================================
Remove Duplicates from Sorted Array

Source: LeetCode
Problem ID: 26
Source URL:
../CONCEPTS_LINKS.md#leetcode-source-0026

Given an integer array nums sorted in non-decreasing order, remove the
duplicates in-place such that each unique element appears only once.
The relative order of the elements should be kept the same.

Consider the number of unique elements in nums to be k. After removing
duplicates, return the number of unique elements k.

The first k elements of nums should contain the unique numbers in
sorted order. The remaining elements beyond index k - 1 can be
ignored.

Custom Judge:

The judge will test your solution with the following code:

int[] nums = [...]; // Input array
int[] expectedNums = [...]; // The expected answer with correct length

int k = removeDuplicates(nums); // Calls your implementation

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
assert nums[i] == expectedNums[i];
}

If all assertions pass, then your solution will be accepted.

Example 1:

Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
Explanation: Your function should return k = 2, with the first two
elements of nums being 1 and 2 respectively.
It does not matter what you leave beyond the returned k (hence they
are underscores).

Example 2:

Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
Explanation: Your function should return k = 5, with the first five
elements of nums being 0, 1, 2, 3, and 4 respectively.
It does not matter what you leave beyond the returned k (hence they
are underscores).

Constraints:

- 1 <= nums.length <= 3 * 10^4

- -100 <= nums[i] <= 100

- nums is sorted in non-decreasing order.
======================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here.
// Name adapters after the problem. Put setup used only by the current
// attempt at the top of solveCurrentAttempt below.
//
// This completed archive follows the workbench structure.

func parseRemoveDuplicatesFromSortedArrayInput(input string) []int {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse remove-duplicates-from-sorted-array "+
			"input: %w", err))
	}
	if len(args) != 1 {
		panic(
			fmt.Sprintf("parse remove-duplicates-from-sorted-array "+
				"input: got %d arguments, want 1", len(args)),
		)
	}
	var nums []int
	argumentError0 := json.Unmarshal(args[0], &nums)
	if argumentError0 != nil {
		panic(fmt.Errorf("parse remove-duplicates-from-sorted-array "+
			"argument nums: %w", argumentError0))
	}
	return nums
}

func formatRemoveDuplicatesFromSortedArrayOutput(
	k int,
	nums []int,
) string {
	values := make([]string, len(nums))
	for index := range values {
		if index < k {
			values[index] = strconv.Itoa(nums[index])
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

func maximumRemoveDuplicatesFromSortedArrayCase() testrunner.Case {
	nums := make([]int, 30000)
	for index := range nums {
		nums[index] = -100 + index*201/len(nums)
	}
	input, err := json.Marshal([]any{nums})
	if err != nil {
		panic(
			fmt.Errorf(
				"build maximum remove-duplicates input: %w",
				err,
			),
		)
	}

	wantNums := make([]int, len(nums))
	for index := range 201 {
		wantNums[index] = index - 100
	}
	return testrunner.Case{
		Name:  "Boundary: Maximum length spans every allowed value",
		Input: string(input),
		Want: formatRemoveDuplicatesFromSortedArrayOutput(
			201,
			wantNums,
		),
	}
}

// sampleCases contains the examples that every registered solution
// must pass. Inputs encode the judge arguments as a positional JSON
// array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `[[1,1,2]]`,
		Want:  `2, nums = [1,2,_]`,
	},
	{
		Name:  "Official example 2",
		Input: `[[0,0,1,1,1,2,2,3,3,4]]`,
		Want:  `5, nums = [0,1,2,3,4,_,_,_,_,_]`,
	},
	{
		Name:  "Additional: Minimum array",
		Input: `[[1]]`,
		Want:  `1, nums = [1]`,
	},
	{
		Name:  "Additional: All values duplicated",
		Input: `[[2,2,2]]`,
		Want:  `1, nums = [2,_,_]`,
	},
	{
		Name:  "Additional: No duplicates",
		Input: `[[-1,0,1]]`,
		Want:  `3, nums = [-1,0,1]`,
	},
	{
		Name:  "Boundary: Duplicate minimum and maximum values",
		Input: `[[-100,-100,100,100]]`,
		Want:  `2, nums = [-100,100,_,_]`,
	},
	{
		Name:  "Additional: Duplicate runs at both ends",
		Input: `[[1,1,2,3,3]]`,
		Want:  `3, nums = [1,2,3,_,_]`,
	},
	{
		Name:  "Additional: Several runs of different lengths",
		Input: `[[-2,-2,-1,-1,0,1,1,2,2]]`,
		Want:  `5, nums = [-2,-1,0,1,2,_,_,_,_]`,
	},
	{
		Name:  "Adversarial: Long runs surround singleton",
		Input: `[[-1,-1,-1,0,1,1,1]]`,
		Want:  `3, nums = [-1,0,1,_,_,_,_]`,
	},
	maximumRemoveDuplicatesFromSortedArrayCase(),
}

// archiveReady is set to true only after the current attempt, three
// agent solutions, analysis, and regression coverage are complete and
// verified.
const archiveReady = true

// Registered implementations run against the same cases.
var currentAttemptSolution = testrunner.Solution{
	Name: "current attempt",
	Run:  solveCurrentAttempt,
}

// agentSolutions contains direct, improved, and preferred
// implementations.
var agentSolutions = []testrunner.Solution{
	{Name: "shift each duplicate left", Run: solveByRepeatedShifting},
	{Name: "record bounded values", Run: solveWithBoundedValueBits},
	{Name: "compact with write pointer", Run: solveByWritePointer},
}

var solutions = append(
	[]testrunner.Solution{currentAttemptSolution},
	agentSolutions...,
)

// ===================================================================
// CURRENT ATTEMPT
//
// This preserves the user-authored implementation reviewed above.
// ===================================================================

func solveCurrentAttempt(input string) string {
	nums := parseRemoveDuplicatesFromSortedArrayInput(input)
	shift := 0

	for i, val := range nums {
		if i > 0 && val == nums[i-1] {
			shift += 1
		}
		nums[i-shift] = val
	}

	// Return the judge-visible count and compacted prefix.
	k := len(nums) - shift
	return formatRemoveDuplicatesFromSortedArrayOutput(k, nums)
}

// ===================================================================
// AGENT SOLUTIONS
//
// These implementations remain separate from the preserved current
// attempt.
// ===================================================================

// solveByRepeatedShifting removes one duplicate at a time by shifting
// the remaining active suffix left.
func solveByRepeatedShifting(input string) string {
	nums := parseRemoveDuplicatesFromSortedArrayInput(input)
	activeLength := len(nums)

	for index := 1; index < activeLength; {
		if nums[index] != nums[index-1] {
			index++
			continue
		}
		copy(nums[index:activeLength-1], nums[index+1:activeLength])
		activeLength--
	}

	return formatRemoveDuplicatesFromSortedArrayOutput(
		activeLength,
		nums,
	)
}

const (
	minimumAllowedValue = -100
	maximumAllowedValue = 100
)

// allowedValueBits stores the problem's 201 possible values in four
// machine words, so its auxiliary-space use is constant and
// independent of input size.
type allowedValueBits struct {
	word0 uint64
	word1 uint64
	word2 uint64
	word3 uint64
}

func (bits *allowedValueBits) add(value int) {
	offset := uint(value - minimumAllowedValue)
	mask := uint64(1) << (offset % 64)

	switch offset / 64 {
	case 0:
		bits.word0 |= mask
	case 1:
		bits.word1 |= mask
	case 2:
		bits.word2 |= mask
	case 3:
		bits.word3 |= mask
	}
}

func (bits allowedValueBits) contains(value int) bool {
	offset := uint(value - minimumAllowedValue)
	mask := uint64(1) << (offset % 64)

	switch offset / 64 {
	case 0:
		return bits.word0&mask != 0
	case 1:
		return bits.word1&mask != 0
	case 2:
		return bits.word2&mask != 0
	case 3:
		return bits.word3&mask != 0
	default:
		return false
	}
}

// solveWithBoundedValueBits records which allowed values occur, then
// rewrites those values into nums in ascending order.
func solveWithBoundedValueBits(input string) string {
	nums := parseRemoveDuplicatesFromSortedArrayInput(input)
	var present allowedValueBits
	for _, value := range nums {
		present.add(value)
	}

	write := 0
	value := minimumAllowedValue
	for value <= maximumAllowedValue {
		if !present.contains(value) {
			value++
			continue
		}
		nums[write] = value
		write++
		value++
	}

	return formatRemoveDuplicatesFromSortedArrayOutput(write, nums)
}

// solveByWritePointer compacts each newly observed value into the
// next output slot without allocating another value collection.
func solveByWritePointer(input string) string {
	nums := parseRemoveDuplicatesFromSortedArrayInput(input)
	k := compactUniqueSorted(nums)
	return formatRemoveDuplicatesFromSortedArrayOutput(k, nums)
}

// removeDuplicates is the judge-facing wrapper around the preferred
// approach.
func removeDuplicates(nums []int) int {
	return compactUniqueSorted(nums)
}

func compactUniqueSorted(nums []int) int {
	write := 1
	for read := 1; read < len(nums); read++ {
		if nums[read] == nums[write-1] {
			continue
		}
		nums[write] = nums[read]
		write++
	}
	return write
}

/*
======================================================================
AGENT REVIEW AND COMPLEXITY ANALYSIS

Current attempt:
- What is correct: The attempt is correct for every valid input.
  Because nums
is sorted, equal values form one adjacent run. shift equals the
duplicates
seen so far, so i-shift is exactly the output position for val. The
range
  loop reads each original position before any later position can be
  overwritten, and k=len(nums)-shift is the number of unique values.
- First failing assumption or bottleneck: No assumption fails under
  the stated
sorted, non-empty constraints. The main readability cost is expressing
the
write position indirectly as i-shift instead of tracking the
kept-prefix
  length directly.
- Relevant concept and invariant: See ../CONCEPTS.md#two-pointers.
After iteration i, nums[:i+1-shift] contains the unique values from
the
  original prefix nums[:i+1], in sorted order.
- Concrete example: For nums=[0,0,1,1,2], (i,val,shift,write)
  progresses as
  (0,0,0,0), (1,0,1,0), (2,1,1,1), (3,1,2,1), and (4,2,2,2).
  The final k is 5-2=3 and the kept prefix is [0,1,2].
- Time complexity: Best, average, and worst are Theta(n); every input
  position
  is visited once. Parsing and formatting are also Theta(n).
- Auxiliary-space complexity: O(1) for the algorithm. The workbench
  adapters
  allocate Theta(n) space for decoded input and formatted output.
- Relationship to agent solutions: Solution 1 mutates in place too,
  but shifts
a suffix for every duplicate while the attempt writes each visited
value at
most once. Solution 2 does not use adjacent comparisons or a
compacted-prefix
invariant; it records membership across the fixed value domain, then
rebuilds
the prefix in value order. Solution 3 uses the same linear in-place
compaction
algorithm and invariant as the attempt; its explicit write pointer is
a
  material implementation difference, not a distinct algorithm.

Solution 1 - Shift each duplicate left (direct):
- Correctness: Before each comparison, nums[:activeLength] contains
  the same
sorted multiset as the unprocessed logical array. When adjacent values
match,
shifting the suffix left removes exactly one duplicate and keeps index
at the
same location so runs longer than two are fully processed. When they
differ,
  advancing index safely preserves both unique values.
- Concrete example: For nums=[1,1,1,2], index=1 finds a duplicate,
  shifts the
active suffix to obtain [1,1,2,_], and reduces activeLength to 3.
Index 1
still duplicates its predecessor, so another shift gives [1,2,_,_] and
activeLength=2. The comparison at index 1 then differs and the result
is k=2.
- Time complexity: Algorithmic best is Theta(n) when all values are
  unique.
Worst is Theta(n^2) when early duplicates repeatedly shift long
suffixes;
average depends on duplicate positions. Parsing and formatting add
Theta(n).
- Auxiliary-space complexity: O(1) for the algorithm; `copy` reuses
  nums. The
  adapters allocate Theta(n) for decoded input and formatted output.

Solution 2 - Record bounded values (improved):
- Correctness: After scanning nums, the bit for value v is set exactly
  when v
occurs in the input. Enumerating every allowed value from -100 through
100
writes each present value once and in ascending order. The resulting
prefix
therefore contains exactly the input's distinct values in their
required
  order, and write is their count.
- Concrete example: For nums=[-100,-100,0,2,2,100], the scan sets
  offsets 0,
100, 102, and 200. Enumerating the domain writes -100 at index 0, 0 at
index
  1, 2 at index 2, and 100 at index 3, producing k=4 and prefix
  [-100,0,2,100].
- Time complexity: Best, average, and worst are Theta(n+r), where
  r=201 is the
fixed value-domain size; this is Theta(n) under the stated
constraints.
  Parsing and formatting are also Theta(n).
- Auxiliary-space complexity: O(1): four uint64 words cover all 201
  allowed
values, independent of n. The adapters allocate Theta(n) for decoded
input
  and formatted output.

Solution 3 - Compact with a write pointer (preferred):
- Correctness: Before each read, nums[:write] contains every distinct
  value
from the processed prefix exactly once. A value equal to nums[write-1]
belongs to the current run and is skipped. A different value starts
the next
run, so writing it at nums[write] and incrementing write
re-establishes the
invariant. Since write never exceeds read, no unread value is
overwritten.
- Concrete example: For nums=[0,0,1,1,2], write starts at 1. read=1
  skips 0;
read=2 writes 1 at index 1 and makes write=2; read=3 skips 1; read=4
writes 2
  at index 2 and makes write=3. The result is k=3 and prefix [0,1,2].
- Time complexity: Best, average, and worst are Theta(n); the read
  pointer
  visits every element once. Parsing and formatting are also Theta(n).
- Auxiliary-space complexity: O(1) for compaction. The adapters
  allocate
  Theta(n) for decoded input and formatted output.

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
