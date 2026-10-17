# Coding Interview Preparation Plan

Interview date: Wednesday, October 21, 2026.

The existing Top Interview 150 archive is sufficient for the coding portion of
a general software engineering interview. The goal is focused pattern coverage
and clear interview execution, not completing all 150 problems.

Target: complete the 18 priority problems, re-solve representative problems
without notes, and complete two timed mock interviews by Tuesday night. This
schedule assumes four to six focused preparation hours per day.

## Operating rules

- Spend no more than 25 to 30 minutes on an initial independent attempt.
- State the brute-force approach before optimizing.
- Identify the invariant or data structure that drives the solution.
- Explain the approach aloud while coding.
- Test minimum, boundary, duplicate, and adversarial cases.
- State time and auxiliary-space complexity accurately.
- When stuck, study the solution, close it, and reimplement it without notes.
- Re-solve two problems from the previous day each morning.

## Saturday, October 17

- [ ] Run one 45-minute diagnostic using [49 Group Anagrams](problems/medium_leetcode_0049_group_anagrams.go).
- [ ] Record where time was lost: pattern recognition, data-structure choice,
      implementation, testing, or explanation.
- [ ] Reserve uninterrupted study blocks for Sunday through Tuesday.

## Sunday, October 18

Focus: arrays, hashing, two pointers, sliding windows, prefix and suffix
processing, and intervals.

- [ ] [15 3Sum](problems/medium_leetcode_0015_3sum.go) - sorting and two pointers
- [ ] [49 Group Anagrams](problems/medium_leetcode_0049_group_anagrams.go) - hashing and canonical keys
- [ ] [56 Merge Intervals](problems/medium_leetcode_0056_merge_intervals.go) - sorting and interval invariants
- [ ] [121 Best Time to Buy and Sell Stock](problems/easy_leetcode_0121_best_time_to_buy_and_sell_stock.go) - single-pass greedy state
- [ ] [209 Minimum Size Subarray Sum](problems/medium_leetcode_0209_minimum_size_subarray_sum.go) - sliding window
- [ ] [238 Product of Array Except Self](problems/medium_leetcode_0238_product_of_array_except_self.go) - prefix and suffix products

## Monday, October 19

Focus: pointer manipulation, stack design, tree traversal, and binary search
tree invariants.

- [ ] Re-solve two Sunday problems without notes.
- [ ] [19 Remove Nth Node From End of List](problems/medium_leetcode_0019_remove_nth_node_from_end_of_list.go) - fast and slow pointers
- [ ] [141 Linked List Cycle](problems/easy_leetcode_0141_linked_list_cycle.go) - Floyd cycle detection
- [ ] [155 Min Stack](problems/medium_leetcode_0155_min_stack.go) - augmented stack state
- [ ] [104 Maximum Depth of Binary Tree](problems/easy_leetcode_0104_maximum_depth_of_binary_tree.go) - recursive tree traversal
- [ ] [102 Binary Tree Level Order Traversal](problems/medium_leetcode_0102_binary_tree_level_order_traversal.go) - breadth-first search
- [ ] [98 Validate Binary Search Tree](problems/medium_leetcode_0098_validate_binary_search_tree.go) - propagated value bounds

## Tuesday, October 20

Focus: graph traversal, topological sorting, heaps, backtracking, and dynamic
programming.

- [ ] Re-solve two Monday problems without notes.
- [ ] [200 Number of Islands](problems/medium_leetcode_0200_number_of_islands.go) - grid traversal
- [ ] [207 Course Schedule](problems/medium_leetcode_0207_course_schedule.go) - cycle detection and topological sorting
- [ ] [215 Kth Largest Element in an Array](problems/medium_leetcode_0215_kth_largest_element_in_an_array.go) - heap or selection
- [ ] [39 Combination Sum](problems/medium_leetcode_0039_combination_sum.go) - backtracking and pruning
- [ ] [139 Word Break](problems/medium_leetcode_0139_word_break.go) - one-dimensional dynamic programming
- [ ] [322 Coin Change](problems/medium_leetcode_0322_coin_change.go) - unbounded dynamic programming
- [ ] Complete two 45-minute mock interviews.

## Wednesday, October 21

- [ ] Warm up with one familiar easy problem.
- [ ] Re-solve one familiar medium problem.
- [ ] Review pattern triggers and complexity statements.
- [ ] Stop learning new algorithms before the interview.

## Readiness gate

Coding-screen readiness is reasonable when all of the following are true:

- [ ] At least 12 of the 18 priority problems can be solved independently.
- [ ] An unseen medium problem can be completed in approximately 30 to 35 minutes.
- [ ] The approach can be explained continuously while coding.
- [ ] Go code compiles without extensive syntax searching.
- [ ] Important edge cases are tested without prompting.
- [ ] Time and space complexity are explained correctly.
- [ ] Two timed mock interviews have been completed.

## Noncoding preparation

The problem archive covers algorithms, not the complete interview.

- [ ] Prepare three concise STAR stories covering impact, conflict, and failure
      or learning.
- [ ] Practice a two-minute explanation of recent work and technical decisions.
- [ ] Research the company, role, product, and interview format.
- [ ] Review system-design fundamentals if the role is senior or the interview
      includes design.

## If time falls behind

Do not compensate by rushing through every problem. Preserve breadth by
completing at least one representative problem from each major area: arrays and
hashing, sliding window, intervals, linked lists, trees, graphs, heaps,
backtracking, and dynamic programming. Drop extra repetitions before dropping
an entire pattern.
