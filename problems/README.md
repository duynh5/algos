# Problem Index

Every problem is a Go file directly inside this directory, and each file can be
run on its own. A completed file contains the problem statement, test cases, the
user's original attempt, three agent-authored solutions, an explanation of why
each solution works, and its work and memory costs.

A newly retrieved practice file has the same sections, but its current solution,
list of agent solutions, and analysis begin blank. It still includes the source
statement, official examples, helpers that convert sample text into specific Go
values such as integers or slices and convert results back to text, and
`archiveReady = false`. The current solution may later contain work while the
file remains unfinished.

After implementing a practice file's current attempt, run that file directly to
check it against its sample cases. The shared code in
[`internal/testrunner`](../internal/testrunner/) performs those comparisons. The
bulk test script skips unfinished practice files. A later review adds the three
solutions and analysis before changing `archiveReady` to `true`.

Completed filenames use
`solved_<difficulty>_<source>_<problem_id>_<descriptive_slug>.go`; blank
practice files omit `solved_`. Numeric problem IDs contain leading zeroes until
they are at least four digits long. This keeps ordinary filename sorting in
numeric order inside the same completion status, difficulty, and source group:

- `solved_easy_leetcode_0001_two_sum.go`
- `medium_leetcode_0322_coin_change.go`
- `hard_leetcode_0004_median_of_two_sorted_arrays.go`

LeetCode is the default source. Pass another source explicitly when archiving a
problem from elsewhere. A source URL can remain inside the file for attribution.
Every problem file starts with `//go:build ignore`. This tells package-wide Go
commands not to combine these independent `main` programs into one package.

Archive and run problems with:

```bash
./scripts/archive-workbench.sh easy 1 two-sum
./scripts/run-problem.sh --case 1 solved_easy_leetcode_0001_two_sum.go
./scripts/run-problem.sh solved_easy_leetcode_0001_two_sum.go
./scripts/test-problems.sh
```

## Problems

- [solved_easy_leetcode_0009_palindrome_number.go](solved_easy_leetcode_0009_palindrome_number.go)
- [solved_easy_leetcode_0013_roman_to_integer.go](solved_easy_leetcode_0013_roman_to_integer.go)
- [solved_easy_leetcode_0014_longest_common_prefix.go](solved_easy_leetcode_0014_longest_common_prefix.go)
- [solved_easy_leetcode_0020_valid_parentheses.go](solved_easy_leetcode_0020_valid_parentheses.go)
- [solved_easy_leetcode_0021_merge_two_sorted_lists.go](solved_easy_leetcode_0021_merge_two_sorted_lists.go)
- [solved_easy_leetcode_0026_remove_duplicates_from_sorted_array.go](solved_easy_leetcode_0026_remove_duplicates_from_sorted_array.go)
- [solved_easy_leetcode_0027_remove_element.go](solved_easy_leetcode_0027_remove_element.go)
- [easy_leetcode_0028_find_the_index_of_the_first_occurrence_in_a_string.go](easy_leetcode_0028_find_the_index_of_the_first_occurrence_in_a_string.go)
- [solved_easy_leetcode_0035_search_insert_position.go](solved_easy_leetcode_0035_search_insert_position.go)
- [easy_leetcode_0058_length_of_last_word.go](easy_leetcode_0058_length_of_last_word.go)
- [easy_leetcode_0066_plus_one.go](easy_leetcode_0066_plus_one.go)
- [easy_leetcode_0067_add_binary.go](easy_leetcode_0067_add_binary.go)
- [solved_easy_leetcode_0069_sqrtx.go](solved_easy_leetcode_0069_sqrtx.go)
- [solved_easy_leetcode_0070_climbing_stairs.go](solved_easy_leetcode_0070_climbing_stairs.go)
- [easy_leetcode_0100_same_tree.go](easy_leetcode_0100_same_tree.go)
- [easy_leetcode_0101_symmetric_tree.go](easy_leetcode_0101_symmetric_tree.go)
- [easy_leetcode_0104_maximum_depth_of_binary_tree.go](easy_leetcode_0104_maximum_depth_of_binary_tree.go)
- [easy_leetcode_0108_convert_sorted_array_to_binary_search_tree.go](easy_leetcode_0108_convert_sorted_array_to_binary_search_tree.go)
- [easy_leetcode_0112_path_sum.go](easy_leetcode_0112_path_sum.go)
- [easy_leetcode_0121_best_time_to_buy_and_sell_stock.go](easy_leetcode_0121_best_time_to_buy_and_sell_stock.go)
- [easy_leetcode_0125_valid_palindrome.go](easy_leetcode_0125_valid_palindrome.go)
- [easy_leetcode_0136_single_number.go](easy_leetcode_0136_single_number.go)
- [easy_leetcode_0141_linked_list_cycle.go](easy_leetcode_0141_linked_list_cycle.go)
- [easy_leetcode_0169_majority_element.go](easy_leetcode_0169_majority_element.go)
- [easy_leetcode_0190_reverse_bits.go](easy_leetcode_0190_reverse_bits.go)
- [easy_leetcode_0191_number_of_1_bits.go](easy_leetcode_0191_number_of_1_bits.go)
- [easy_leetcode_0202_happy_number.go](easy_leetcode_0202_happy_number.go)
- [easy_leetcode_0205_isomorphic_strings.go](easy_leetcode_0205_isomorphic_strings.go)
- [easy_leetcode_0219_contains_duplicate_ii.go](easy_leetcode_0219_contains_duplicate_ii.go)
- [easy_leetcode_0226_invert_binary_tree.go](easy_leetcode_0226_invert_binary_tree.go)
- [easy_leetcode_0228_summary_ranges.go](easy_leetcode_0228_summary_ranges.go)
- [easy_leetcode_0242_valid_anagram.go](easy_leetcode_0242_valid_anagram.go)
- [easy_leetcode_0290_word_pattern.go](easy_leetcode_0290_word_pattern.go)
- [easy_leetcode_0383_ransom_note.go](easy_leetcode_0383_ransom_note.go)
- [easy_leetcode_0392_is_subsequence.go](easy_leetcode_0392_is_subsequence.go)
- [easy_leetcode_0530_minimum_absolute_difference_in_bst.go](easy_leetcode_0530_minimum_absolute_difference_in_bst.go)
- [easy_leetcode_0637_average_of_levels_in_binary_tree.go](easy_leetcode_0637_average_of_levels_in_binary_tree.go)
- [hard_leetcode_0004_median_of_two_sorted_arrays.go](hard_leetcode_0004_median_of_two_sorted_arrays.go)
- [hard_leetcode_0023_merge_k_sorted_lists.go](hard_leetcode_0023_merge_k_sorted_lists.go)
- [hard_leetcode_0025_reverse_nodes_in_k_group.go](hard_leetcode_0025_reverse_nodes_in_k_group.go)
- [hard_leetcode_0030_substring_with_concatenation_of_all_words.go](hard_leetcode_0030_substring_with_concatenation_of_all_words.go)
- [hard_leetcode_0042_trapping_rain_water.go](hard_leetcode_0042_trapping_rain_water.go)
- [hard_leetcode_0052_n_queens_ii.go](hard_leetcode_0052_n_queens_ii.go)
- [hard_leetcode_0068_text_justification.go](hard_leetcode_0068_text_justification.go)
- [hard_leetcode_0076_minimum_window_substring.go](hard_leetcode_0076_minimum_window_substring.go)
- [hard_leetcode_0123_best_time_to_buy_and_sell_stock_iii.go](hard_leetcode_0123_best_time_to_buy_and_sell_stock_iii.go)
- [hard_leetcode_0124_binary_tree_maximum_path_sum.go](hard_leetcode_0124_binary_tree_maximum_path_sum.go)
- [hard_leetcode_0127_word_ladder.go](hard_leetcode_0127_word_ladder.go)
- [hard_leetcode_0135_candy.go](hard_leetcode_0135_candy.go)
- [hard_leetcode_0149_max_points_on_a_line.go](hard_leetcode_0149_max_points_on_a_line.go)
- [hard_leetcode_0188_best_time_to_buy_and_sell_stock_iv.go](hard_leetcode_0188_best_time_to_buy_and_sell_stock_iv.go)
- [hard_leetcode_0212_word_search_ii.go](hard_leetcode_0212_word_search_ii.go)
- [hard_leetcode_0224_basic_calculator.go](hard_leetcode_0224_basic_calculator.go)
- [hard_leetcode_0295_find_median_from_data_stream.go](hard_leetcode_0295_find_median_from_data_stream.go)
- [hard_leetcode_0502_ipo.go](hard_leetcode_0502_ipo.go)
- [solved_medium_leetcode_0005_longest_palindromic_substring.go](solved_medium_leetcode_0005_longest_palindromic_substring.go)
- [medium_leetcode_0006_zigzag_conversion.go](medium_leetcode_0006_zigzag_conversion.go)
- [medium_leetcode_0011_container_with_most_water.go](medium_leetcode_0011_container_with_most_water.go)
- [medium_leetcode_0012_integer_to_roman.go](medium_leetcode_0012_integer_to_roman.go)
- [medium_leetcode_0015_3sum.go](medium_leetcode_0015_3sum.go)
- [medium_leetcode_0017_letter_combinations_of_a_phone_number.go](medium_leetcode_0017_letter_combinations_of_a_phone_number.go)
- [medium_leetcode_0019_remove_nth_node_from_end_of_list.go](medium_leetcode_0019_remove_nth_node_from_end_of_list.go)
- [medium_leetcode_0022_generate_parentheses.go](medium_leetcode_0022_generate_parentheses.go)
- [solved_medium_leetcode_0033_search_in_rotated_sorted_array.go](solved_medium_leetcode_0033_search_in_rotated_sorted_array.go)
- [solved_medium_leetcode_0034_find_first_and_last_position_of_element_in_sorted_array.go](solved_medium_leetcode_0034_find_first_and_last_position_of_element_in_sorted_array.go)
- [medium_leetcode_0036_valid_sudoku.go](medium_leetcode_0036_valid_sudoku.go)
- [medium_leetcode_0039_combination_sum.go](medium_leetcode_0039_combination_sum.go)
- [medium_leetcode_0045_jump_game_ii.go](medium_leetcode_0045_jump_game_ii.go)
- [medium_leetcode_0046_permutations.go](medium_leetcode_0046_permutations.go)
- [medium_leetcode_0048_rotate_image.go](medium_leetcode_0048_rotate_image.go)
- [medium_leetcode_0049_group_anagrams.go](medium_leetcode_0049_group_anagrams.go)
- [medium_leetcode_0050_powx_n.go](medium_leetcode_0050_powx_n.go)
- [medium_leetcode_0053_maximum_subarray.go](medium_leetcode_0053_maximum_subarray.go)
- [medium_leetcode_0054_spiral_matrix.go](medium_leetcode_0054_spiral_matrix.go)
- [medium_leetcode_0055_jump_game.go](medium_leetcode_0055_jump_game.go)
- [medium_leetcode_0056_merge_intervals.go](medium_leetcode_0056_merge_intervals.go)
- [medium_leetcode_0057_insert_interval.go](medium_leetcode_0057_insert_interval.go)
- [medium_leetcode_0061_rotate_list.go](medium_leetcode_0061_rotate_list.go)
- [medium_leetcode_0063_unique_paths_ii.go](medium_leetcode_0063_unique_paths_ii.go)
- [medium_leetcode_0064_minimum_path_sum.go](medium_leetcode_0064_minimum_path_sum.go)
- [medium_leetcode_0071_simplify_path.go](medium_leetcode_0071_simplify_path.go)
- [medium_leetcode_0072_edit_distance.go](medium_leetcode_0072_edit_distance.go)
- [medium_leetcode_0073_set_matrix_zeroes.go](medium_leetcode_0073_set_matrix_zeroes.go)
- [medium_leetcode_0074_search_a_2d_matrix.go](medium_leetcode_0074_search_a_2d_matrix.go)
- [medium_leetcode_0077_combinations.go](medium_leetcode_0077_combinations.go)
- [medium_leetcode_0079_word_search.go](medium_leetcode_0079_word_search.go)
- [medium_leetcode_0080_remove_duplicates_from_sorted_array_ii.go](medium_leetcode_0080_remove_duplicates_from_sorted_array_ii.go)
- [medium_leetcode_0082_remove_duplicates_from_sorted_list_ii.go](medium_leetcode_0082_remove_duplicates_from_sorted_list_ii.go)
- [medium_leetcode_0086_partition_list.go](medium_leetcode_0086_partition_list.go)
- [medium_leetcode_0092_reverse_linked_list_ii.go](medium_leetcode_0092_reverse_linked_list_ii.go)
- [medium_leetcode_0097_interleaving_string.go](medium_leetcode_0097_interleaving_string.go)
- [medium_leetcode_0098_validate_binary_search_tree.go](medium_leetcode_0098_validate_binary_search_tree.go)
- [medium_leetcode_0102_binary_tree_level_order_traversal.go](medium_leetcode_0102_binary_tree_level_order_traversal.go)
- [medium_leetcode_0103_binary_tree_zigzag_level_order_traversal.go](medium_leetcode_0103_binary_tree_zigzag_level_order_traversal.go)
- [medium_leetcode_0105_construct_binary_tree_from_preorder_and_inorder_traversal.go](medium_leetcode_0105_construct_binary_tree_from_preorder_and_inorder_traversal.go)
- [medium_leetcode_0106_construct_binary_tree_from_inorder_and_postorder_traversal.go](medium_leetcode_0106_construct_binary_tree_from_inorder_and_postorder_traversal.go)
- [medium_leetcode_0114_flatten_binary_tree_to_linked_list.go](medium_leetcode_0114_flatten_binary_tree_to_linked_list.go)
- [medium_leetcode_0117_populating_next_right_pointers_in_each_node_ii.go](medium_leetcode_0117_populating_next_right_pointers_in_each_node_ii.go)
- [medium_leetcode_0120_triangle.go](medium_leetcode_0120_triangle.go)
- [medium_leetcode_0122_best_time_to_buy_and_sell_stock_ii.go](medium_leetcode_0122_best_time_to_buy_and_sell_stock_ii.go)
- [medium_leetcode_0128_longest_consecutive_sequence.go](medium_leetcode_0128_longest_consecutive_sequence.go)
- [medium_leetcode_0129_sum_root_to_leaf_numbers.go](medium_leetcode_0129_sum_root_to_leaf_numbers.go)
- [medium_leetcode_0130_surrounded_regions.go](medium_leetcode_0130_surrounded_regions.go)
- [medium_leetcode_0133_clone_graph.go](medium_leetcode_0133_clone_graph.go)
- [medium_leetcode_0134_gas_station.go](medium_leetcode_0134_gas_station.go)
- [medium_leetcode_0137_single_number_ii.go](medium_leetcode_0137_single_number_ii.go)
- [medium_leetcode_0138_copy_list_with_random_pointer.go](medium_leetcode_0138_copy_list_with_random_pointer.go)
- [medium_leetcode_0139_word_break.go](medium_leetcode_0139_word_break.go)
- [medium_leetcode_0146_lru_cache.go](medium_leetcode_0146_lru_cache.go)
- [medium_leetcode_0148_sort_list.go](medium_leetcode_0148_sort_list.go)
- [medium_leetcode_0150_evaluate_reverse_polish_notation.go](medium_leetcode_0150_evaluate_reverse_polish_notation.go)
- [medium_leetcode_0151_reverse_words_in_a_string.go](medium_leetcode_0151_reverse_words_in_a_string.go)
- [medium_leetcode_0153_find_minimum_in_rotated_sorted_array.go](medium_leetcode_0153_find_minimum_in_rotated_sorted_array.go)
- [medium_leetcode_0155_min_stack.go](medium_leetcode_0155_min_stack.go)
- [medium_leetcode_0162_find_peak_element.go](medium_leetcode_0162_find_peak_element.go)
- [medium_leetcode_0167_two_sum_ii_input_array_is_sorted.go](medium_leetcode_0167_two_sum_ii_input_array_is_sorted.go)
- [medium_leetcode_0172_factorial_trailing_zeroes.go](medium_leetcode_0172_factorial_trailing_zeroes.go)
- [medium_leetcode_0173_binary_search_tree_iterator.go](medium_leetcode_0173_binary_search_tree_iterator.go)
- [medium_leetcode_0189_rotate_array.go](medium_leetcode_0189_rotate_array.go)
- [solved_medium_leetcode_0198_house_robber.go](solved_medium_leetcode_0198_house_robber.go)
- [medium_leetcode_0199_binary_tree_right_side_view.go](medium_leetcode_0199_binary_tree_right_side_view.go)
- [medium_leetcode_0200_number_of_islands.go](medium_leetcode_0200_number_of_islands.go)
- [medium_leetcode_0201_bitwise_and_of_numbers_range.go](medium_leetcode_0201_bitwise_and_of_numbers_range.go)
- [medium_leetcode_0207_course_schedule.go](medium_leetcode_0207_course_schedule.go)
- [medium_leetcode_0208_implement_trie_prefix_tree.go](medium_leetcode_0208_implement_trie_prefix_tree.go)
- [medium_leetcode_0209_minimum_size_subarray_sum.go](medium_leetcode_0209_minimum_size_subarray_sum.go)
- [medium_leetcode_0210_course_schedule_ii.go](medium_leetcode_0210_course_schedule_ii.go)
- [medium_leetcode_0211_design_add_and_search_words_data_structure.go](medium_leetcode_0211_design_add_and_search_words_data_structure.go)
- [medium_leetcode_0215_kth_largest_element_in_an_array.go](medium_leetcode_0215_kth_largest_element_in_an_array.go)
- [medium_leetcode_0221_maximal_square.go](medium_leetcode_0221_maximal_square.go)
- [medium_leetcode_0222_count_complete_tree_nodes.go](medium_leetcode_0222_count_complete_tree_nodes.go)
- [medium_leetcode_0230_kth_smallest_element_in_a_bst.go](medium_leetcode_0230_kth_smallest_element_in_a_bst.go)
- [medium_leetcode_0236_lowest_common_ancestor_of_a_binary_tree.go](medium_leetcode_0236_lowest_common_ancestor_of_a_binary_tree.go)
- [medium_leetcode_0238_product_of_array_except_self.go](medium_leetcode_0238_product_of_array_except_self.go)
- [medium_leetcode_0274_h_index.go](medium_leetcode_0274_h_index.go)
- [medium_leetcode_0289_game_of_life.go](medium_leetcode_0289_game_of_life.go)
- [medium_leetcode_0300_longest_increasing_subsequence.go](medium_leetcode_0300_longest_increasing_subsequence.go)
- [medium_leetcode_0322_coin_change.go](medium_leetcode_0322_coin_change.go)
- [medium_leetcode_0373_find_k_pairs_with_smallest_sums.go](medium_leetcode_0373_find_k_pairs_with_smallest_sums.go)
- [medium_leetcode_0380_insert_delete_getrandom_o1.go](medium_leetcode_0380_insert_delete_getrandom_o1.go)
- [medium_leetcode_0399_evaluate_division.go](medium_leetcode_0399_evaluate_division.go)
- [medium_leetcode_0427_construct_quad_tree.go](medium_leetcode_0427_construct_quad_tree.go)
- [medium_leetcode_0433_minimum_genetic_mutation.go](medium_leetcode_0433_minimum_genetic_mutation.go)
- [medium_leetcode_0452_minimum_number_of_arrows_to_burst_balloons.go](medium_leetcode_0452_minimum_number_of_arrows_to_burst_balloons.go)
- [medium_leetcode_0909_snakes_and_ladders.go](medium_leetcode_0909_snakes_and_ladders.go)
- [medium_leetcode_0918_maximum_sum_circular_subarray.go](medium_leetcode_0918_maximum_sum_circular_subarray.go)
- [solved_easy_leetcode_0001_two_sum.go](solved_easy_leetcode_0001_two_sum.go)
- [solved_easy_leetcode_0002_add_two_numbers.go](solved_easy_leetcode_0002_add_two_numbers.go)
- [solved_easy_leetcode_0088_merge_sorted_array.go](solved_easy_leetcode_0088_merge_sorted_array.go)
- [solved_medium_leetcode_0003_longest_substring_without_repeating_characters.go](solved_medium_leetcode_0003_longest_substring_without_repeating_characters.go)
