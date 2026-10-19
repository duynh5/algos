# Coding Interview Preparation Plan

Interview date: Wednesday, October 21, 2026.

The repository contains all 150 problems in the Top Interview 150 list. This
plan selects a smaller group that covers the main coding-interview methods. The
goal is to recognize and explain those methods reliably, not to rush through
all 150 problems.

Target: complete the 18 priority problems, re-solve representative problems
without notes, and complete two timed mock interviews by Tuesday night. This
schedule assumes four to six focused preparation hours per day.

## Operating rules

- Spend no more than 25 to 30 minutes on an initial independent attempt.
- State the simplest correct approach, even if it tries every possibility,
  before improving it.
- Identify the fact that remains true while the algorithm runs and the data
  structure that stores the required information.
- Explain the approach aloud while coding.
- Test the smallest valid input, values at allowed limits, repeated values, and
  cases designed to expose a faulty assumption.
- State how the work and extra memory grow with the input.
- When stuck, study the solution, close it, and reimplement it without notes.
- Re-solve two problems from the previous day each morning.

## Saturday, October 17

- [ ] Run one 45-minute diagnostic using [49 Group Anagrams](problems/medium_leetcode_0049_group_anagrams.go).
- [ ] Record where time was lost: pattern recognition, data-structure choice,
      implementation, testing, or explanation.
- [ ] Reserve uninterrupted study blocks for Sunday through Tuesday.

## Sunday, October 18

Focus: arrays; hash maps; coordinated moving positions; moving ranges of
neighboring values; left-to-right and right-to-left accumulated results; and
overlapping ranges.

- [ ] [15 3Sum](problems/medium_leetcode_0015_3sum.go): sort, fix one
      value, and move two positions inward to find the other two
- [ ] [49 Group Anagrams](problems/medium_leetcode_0049_group_anagrams.go):
      give words with the same letters the same map key
- [ ] [56 Merge Intervals](problems/medium_leetcode_0056_merge_intervals.go):
      sort ranges and combine them only while they overlap
- [ ] [121 Best Time to Buy and Sell Stock](problems/easy_leetcode_0121_best_time_to_buy_and_sell_stock.go):
      track the cheapest earlier price and best profit in one pass
- [ ] [209 Minimum Size Subarray Sum](problems/medium_leetcode_0209_minimum_size_subarray_sum.go):
      grow and shrink one range of neighboring values
- [ ] [238 Product of Array Except Self](problems/medium_leetcode_0238_product_of_array_except_self.go):
      accumulate products from the left and right

## Monday, October 19

Focus: changing linked-list links, designing stacks, visiting tree nodes, and
checking the lower and upper limits allowed at each binary-search-tree node.

- [ ] Re-solve two Sunday problems without notes.
- [ ] [19 Remove Nth Node From End of List](problems/medium_leetcode_0019_remove_nth_node_from_end_of_list.go):
      keep two moving positions a fixed number of nodes apart
- [ ] [141 Linked List Cycle](problems/easy_leetcode_0141_linked_list_cycle.go):
      detect a loop with fast and slow moving positions
- [ ] [155 Min Stack](problems/medium_leetcode_0155_min_stack.go): store the
      current minimum alongside stack values
- [ ] [104 Maximum Depth of Binary Tree](problems/easy_leetcode_0104_maximum_depth_of_binary_tree.go):
      visit child nodes with nested function calls
- [ ] [102 Binary Tree Level Order Traversal](problems/medium_leetcode_0102_binary_tree_level_order_traversal.go):
      visit every node at one depth before moving to the next depth
- [ ] [98 Validate Binary Search Tree](problems/medium_leetcode_0098_validate_binary_search_tree.go):
      carry each node's allowed lower and upper values to its children

## Tuesday, October 20

Focus: visiting connected items, ordering tasks from prerequisites, structures
that quickly remove the next highest-priority item, building and undoing
choices, and reusing answers to smaller problems.

- [ ] Re-solve two Monday problems without notes.
- [ ] [200 Number of Islands](problems/medium_leetcode_0200_number_of_islands.go):
      visit all land cells connected to one starting cell before counting the
      next island
- [ ] [207 Course Schedule](problems/medium_leetcode_0207_course_schedule.go):
      detect impossible prerequisite loops and build a valid course order
- [ ] [215 Kth Largest Element in an Array](problems/medium_leetcode_0215_kth_largest_element_in_an_array.go):
      keep the next largest value easy to remove or repeatedly divide values
      around a chosen value
- [ ] [39 Combination Sum](problems/medium_leetcode_0039_combination_sum.go):
      build and undo choices, stopping when a partial sum cannot succeed
- [ ] [139 Word Break](problems/medium_leetcode_0139_word_break.go): store
      whether each beginning of the string can be split into dictionary words
- [ ] [322 Coin Change](problems/medium_leetcode_0322_coin_change.go): store the
      best answer for each amount while allowing a coin to be reused
- [ ] Complete two 45-minute mock interviews.

## Wednesday, October 21

- [ ] Warm up with one familiar easy problem.
- [ ] Re-solve one familiar medium problem.
- [ ] Review the clues that suggest each method and how its work and memory grow.
- [ ] Stop learning new algorithms before the interview.

## Readiness gate

Coding-screen readiness is reasonable when all of the following are true:

- [ ] At least 12 of the 18 priority problems can be solved independently.
- [ ] An unseen medium problem can be completed in approximately 30 to 35 minutes.
- [ ] The approach can be explained continuously while coding.
- [ ] Go code compiles without extensive syntax searching.
- [ ] Important edge cases are tested without prompting.
- [ ] The growth of work and extra memory is explained correctly.
- [ ] Two timed mock interviews have been completed.

## Noncoding preparation

The problem archive covers algorithms, not the complete interview.

- [ ] Prepare three concise stories using Situation, Task, Action, and Result,
      covering impact, conflict, and failure or learning.
- [ ] Practice a two-minute explanation of recent work and technical decisions.
- [ ] Research the company, role, product, and interview format.
- [ ] Review how to design larger software systems if the role is senior or the
      interview includes an architecture discussion.

## If time falls behind

Do not compensate by rushing through every problem. Keep coverage across the
main areas by completing at least one representative problem involving arrays
and maps, a moving range, intervals, linked lists, trees, connected networks,
priority queues, building and undoing choices, and reusing smaller answers.
Drop extra repetitions before dropping an entire area.
