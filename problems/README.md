# Problem Index

This is a flat archive of individually runnable Go programs. Completed problem
files contain their statement, samples, preserved current attempt, three
agent-authored solutions, and correctness and complexity analysis. Unsolved
practice files follow the same section layout, but leave the current-attempt
body, agent solution registry, and analysis content blank for later work. They
retain the source statement, official examples, typed input adapters, and
`archiveReady = false` marker.

After implementing a practice file's current attempt, run that file directly
to execute its sample cases through
[`internal/testrunner`](../internal/testrunner/). The bulk test script skips
unfinished practice archives; a later agent review can add the three solutions
and analysis, then mark the archive ready for repository-wide execution.

Completed filenames use
`solved_<difficulty>_<source>_<problem_id>_<descriptive_slug>.go`; blank
practice files omit `solved_`:

- `solved_easy_leetcode_1_two_sum.go`
- `medium_leetcode_322_coin_change.go`
- `hard_leetcode_72_edit_distance.go`

LeetCode is the default source; pass another source explicitly when archiving
an exception. A source URL can remain inside the file for attribution. Every
archived file starts with `//go:build ignore`, which prevents directory-wide Go
commands from combining independent `main` programs.

Archive and run problems with:

```bash
./scripts/archive-workbench.sh easy 1 two-sum
./scripts/run-problem.sh solved_easy_leetcode_1_two_sum.go
./scripts/test-problems.sh
```

## Problems
- [solved_easy_leetcode_1_two_sum.go](solved_easy_leetcode_1_two_sum.go)
- [solved_easy_leetcode_2_add_two_numbers.go](solved_easy_leetcode_2_add_two_numbers.go)
- [solved_medium_leetcode_3_longest_substring_without_repeating_characters.go](solved_medium_leetcode_3_longest_substring_without_repeating_characters.go)
- [easy_leetcode_100_same_tree.go](easy_leetcode_100_same_tree.go)
- [easy_leetcode_101_symmetric_tree.go](easy_leetcode_101_symmetric_tree.go)
- [easy_leetcode_104_maximum_depth_of_binary_tree.go](easy_leetcode_104_maximum_depth_of_binary_tree.go)
- [easy_leetcode_108_convert_sorted_array_to_binary_search_tree.go](easy_leetcode_108_convert_sorted_array_to_binary_search_tree.go)
- [easy_leetcode_112_path_sum.go](easy_leetcode_112_path_sum.go)
- [easy_leetcode_121_best_time_to_buy_and_sell_stock.go](easy_leetcode_121_best_time_to_buy_and_sell_stock.go)
- [easy_leetcode_125_valid_palindrome.go](easy_leetcode_125_valid_palindrome.go)
- [easy_leetcode_136_single_number.go](easy_leetcode_136_single_number.go)
- [easy_leetcode_13_roman_to_integer.go](easy_leetcode_13_roman_to_integer.go)
- [easy_leetcode_141_linked_list_cycle.go](easy_leetcode_141_linked_list_cycle.go)
- [easy_leetcode_14_longest_common_prefix.go](easy_leetcode_14_longest_common_prefix.go)
- [easy_leetcode_169_majority_element.go](easy_leetcode_169_majority_element.go)
- [easy_leetcode_190_reverse_bits.go](easy_leetcode_190_reverse_bits.go)
- [easy_leetcode_191_number_of_1_bits.go](easy_leetcode_191_number_of_1_bits.go)
- [easy_leetcode_1_two_sum_practice.go](easy_leetcode_1_two_sum_practice.go)
- [easy_leetcode_202_happy_number.go](easy_leetcode_202_happy_number.go)
- [easy_leetcode_205_isomorphic_strings.go](easy_leetcode_205_isomorphic_strings.go)
- [easy_leetcode_20_valid_parentheses.go](easy_leetcode_20_valid_parentheses.go)
- [easy_leetcode_219_contains_duplicate_ii.go](easy_leetcode_219_contains_duplicate_ii.go)
- [easy_leetcode_21_merge_two_sorted_lists.go](easy_leetcode_21_merge_two_sorted_lists.go)
- [easy_leetcode_226_invert_binary_tree.go](easy_leetcode_226_invert_binary_tree.go)
- [easy_leetcode_228_summary_ranges.go](easy_leetcode_228_summary_ranges.go)
- [easy_leetcode_242_valid_anagram.go](easy_leetcode_242_valid_anagram.go)
- [easy_leetcode_26_remove_duplicates_from_sorted_array.go](easy_leetcode_26_remove_duplicates_from_sorted_array.go)
- [easy_leetcode_27_remove_element.go](easy_leetcode_27_remove_element.go)
- [easy_leetcode_28_find_the_index_of_the_first_occurrence_in_a_string.go](easy_leetcode_28_find_the_index_of_the_first_occurrence_in_a_string.go)
- [easy_leetcode_290_word_pattern.go](easy_leetcode_290_word_pattern.go)
- [easy_leetcode_35_search_insert_position.go](easy_leetcode_35_search_insert_position.go)
- [easy_leetcode_383_ransom_note.go](easy_leetcode_383_ransom_note.go)
- [easy_leetcode_392_is_subsequence.go](easy_leetcode_392_is_subsequence.go)
- [easy_leetcode_530_minimum_absolute_difference_in_bst.go](easy_leetcode_530_minimum_absolute_difference_in_bst.go)
- [easy_leetcode_58_length_of_last_word.go](easy_leetcode_58_length_of_last_word.go)
- [easy_leetcode_637_average_of_levels_in_binary_tree.go](easy_leetcode_637_average_of_levels_in_binary_tree.go)
- [easy_leetcode_66_plus_one.go](easy_leetcode_66_plus_one.go)
- [easy_leetcode_67_add_binary.go](easy_leetcode_67_add_binary.go)
- [easy_leetcode_69_sqrtx.go](easy_leetcode_69_sqrtx.go)
- [easy_leetcode_70_climbing_stairs.go](easy_leetcode_70_climbing_stairs.go)
- [easy_leetcode_88_merge_sorted_array_practice.go](easy_leetcode_88_merge_sorted_array_practice.go)
- [easy_leetcode_9_palindrome_number.go](easy_leetcode_9_palindrome_number.go)
- [hard_leetcode_123_best_time_to_buy_and_sell_stock_iii.go](hard_leetcode_123_best_time_to_buy_and_sell_stock_iii.go)
- [hard_leetcode_124_binary_tree_maximum_path_sum.go](hard_leetcode_124_binary_tree_maximum_path_sum.go)
- [hard_leetcode_127_word_ladder.go](hard_leetcode_127_word_ladder.go)
- [hard_leetcode_135_candy.go](hard_leetcode_135_candy.go)
- [hard_leetcode_149_max_points_on_a_line.go](hard_leetcode_149_max_points_on_a_line.go)
- [hard_leetcode_188_best_time_to_buy_and_sell_stock_iv.go](hard_leetcode_188_best_time_to_buy_and_sell_stock_iv.go)
- [hard_leetcode_212_word_search_ii.go](hard_leetcode_212_word_search_ii.go)
- [hard_leetcode_224_basic_calculator.go](hard_leetcode_224_basic_calculator.go)
- [hard_leetcode_23_merge_k_sorted_lists.go](hard_leetcode_23_merge_k_sorted_lists.go)
- [hard_leetcode_25_reverse_nodes_in_k_group.go](hard_leetcode_25_reverse_nodes_in_k_group.go)
- [hard_leetcode_295_find_median_from_data_stream.go](hard_leetcode_295_find_median_from_data_stream.go)
- [hard_leetcode_30_substring_with_concatenation_of_all_words.go](hard_leetcode_30_substring_with_concatenation_of_all_words.go)
- [hard_leetcode_42_trapping_rain_water.go](hard_leetcode_42_trapping_rain_water.go)
- [hard_leetcode_4_median_of_two_sorted_arrays.go](hard_leetcode_4_median_of_two_sorted_arrays.go)
- [hard_leetcode_502_ipo.go](hard_leetcode_502_ipo.go)
- [hard_leetcode_52_n_queens_ii.go](hard_leetcode_52_n_queens_ii.go)
- [hard_leetcode_68_text_justification.go](hard_leetcode_68_text_justification.go)
- [hard_leetcode_76_minimum_window_substring.go](hard_leetcode_76_minimum_window_substring.go)
- [medium_leetcode_102_binary_tree_level_order_traversal.go](medium_leetcode_102_binary_tree_level_order_traversal.go)
- [medium_leetcode_103_binary_tree_zigzag_level_order_traversal.go](medium_leetcode_103_binary_tree_zigzag_level_order_traversal.go)
- [medium_leetcode_105_construct_binary_tree_from_preorder_and_inorder_traversal.go](medium_leetcode_105_construct_binary_tree_from_preorder_and_inorder_traversal.go)
- [medium_leetcode_106_construct_binary_tree_from_inorder_and_postorder_traversal.go](medium_leetcode_106_construct_binary_tree_from_inorder_and_postorder_traversal.go)
- [medium_leetcode_114_flatten_binary_tree_to_linked_list.go](medium_leetcode_114_flatten_binary_tree_to_linked_list.go)
- [medium_leetcode_117_populating_next_right_pointers_in_each_node_ii.go](medium_leetcode_117_populating_next_right_pointers_in_each_node_ii.go)
- [medium_leetcode_11_container_with_most_water.go](medium_leetcode_11_container_with_most_water.go)
- [medium_leetcode_120_triangle.go](medium_leetcode_120_triangle.go)
- [medium_leetcode_122_best_time_to_buy_and_sell_stock_ii.go](medium_leetcode_122_best_time_to_buy_and_sell_stock_ii.go)
- [medium_leetcode_128_longest_consecutive_sequence.go](medium_leetcode_128_longest_consecutive_sequence.go)
- [medium_leetcode_129_sum_root_to_leaf_numbers.go](medium_leetcode_129_sum_root_to_leaf_numbers.go)
- [medium_leetcode_12_integer_to_roman.go](medium_leetcode_12_integer_to_roman.go)
- [medium_leetcode_130_surrounded_regions.go](medium_leetcode_130_surrounded_regions.go)
- [medium_leetcode_133_clone_graph.go](medium_leetcode_133_clone_graph.go)
- [medium_leetcode_134_gas_station.go](medium_leetcode_134_gas_station.go)
- [medium_leetcode_137_single_number_ii.go](medium_leetcode_137_single_number_ii.go)
- [medium_leetcode_138_copy_list_with_random_pointer.go](medium_leetcode_138_copy_list_with_random_pointer.go)
- [medium_leetcode_139_word_break.go](medium_leetcode_139_word_break.go)
- [medium_leetcode_146_lru_cache.go](medium_leetcode_146_lru_cache.go)
- [medium_leetcode_148_sort_list.go](medium_leetcode_148_sort_list.go)
- [medium_leetcode_150_evaluate_reverse_polish_notation.go](medium_leetcode_150_evaluate_reverse_polish_notation.go)
- [medium_leetcode_151_reverse_words_in_a_string.go](medium_leetcode_151_reverse_words_in_a_string.go)
- [medium_leetcode_153_find_minimum_in_rotated_sorted_array.go](medium_leetcode_153_find_minimum_in_rotated_sorted_array.go)
- [medium_leetcode_155_min_stack.go](medium_leetcode_155_min_stack.go)
- [medium_leetcode_15_3sum.go](medium_leetcode_15_3sum.go)
- [medium_leetcode_162_find_peak_element.go](medium_leetcode_162_find_peak_element.go)
- [medium_leetcode_167_two_sum_ii_input_array_is_sorted.go](medium_leetcode_167_two_sum_ii_input_array_is_sorted.go)
- [medium_leetcode_172_factorial_trailing_zeroes.go](medium_leetcode_172_factorial_trailing_zeroes.go)
- [medium_leetcode_173_binary_search_tree_iterator.go](medium_leetcode_173_binary_search_tree_iterator.go)
- [medium_leetcode_17_letter_combinations_of_a_phone_number.go](medium_leetcode_17_letter_combinations_of_a_phone_number.go)
- [medium_leetcode_189_rotate_array.go](medium_leetcode_189_rotate_array.go)
- [medium_leetcode_198_house_robber.go](medium_leetcode_198_house_robber.go)
- [medium_leetcode_199_binary_tree_right_side_view.go](medium_leetcode_199_binary_tree_right_side_view.go)
- [medium_leetcode_19_remove_nth_node_from_end_of_list.go](medium_leetcode_19_remove_nth_node_from_end_of_list.go)
- [medium_leetcode_200_number_of_islands.go](medium_leetcode_200_number_of_islands.go)
- [medium_leetcode_201_bitwise_and_of_numbers_range.go](medium_leetcode_201_bitwise_and_of_numbers_range.go)
- [medium_leetcode_207_course_schedule.go](medium_leetcode_207_course_schedule.go)
- [medium_leetcode_208_implement_trie_prefix_tree.go](medium_leetcode_208_implement_trie_prefix_tree.go)
- [medium_leetcode_209_minimum_size_subarray_sum.go](medium_leetcode_209_minimum_size_subarray_sum.go)
- [medium_leetcode_210_course_schedule_ii.go](medium_leetcode_210_course_schedule_ii.go)
- [medium_leetcode_211_design_add_and_search_words_data_structure.go](medium_leetcode_211_design_add_and_search_words_data_structure.go)
- [medium_leetcode_215_kth_largest_element_in_an_array.go](medium_leetcode_215_kth_largest_element_in_an_array.go)
- [medium_leetcode_221_maximal_square.go](medium_leetcode_221_maximal_square.go)
- [medium_leetcode_222_count_complete_tree_nodes.go](medium_leetcode_222_count_complete_tree_nodes.go)
- [medium_leetcode_22_generate_parentheses.go](medium_leetcode_22_generate_parentheses.go)
- [medium_leetcode_230_kth_smallest_element_in_a_bst.go](medium_leetcode_230_kth_smallest_element_in_a_bst.go)
- [medium_leetcode_236_lowest_common_ancestor_of_a_binary_tree.go](medium_leetcode_236_lowest_common_ancestor_of_a_binary_tree.go)
- [medium_leetcode_238_product_of_array_except_self.go](medium_leetcode_238_product_of_array_except_self.go)
- [medium_leetcode_274_h_index.go](medium_leetcode_274_h_index.go)
- [medium_leetcode_289_game_of_life.go](medium_leetcode_289_game_of_life.go)
- [medium_leetcode_2_add_two_numbers_practice.go](medium_leetcode_2_add_two_numbers_practice.go)
- [medium_leetcode_300_longest_increasing_subsequence.go](medium_leetcode_300_longest_increasing_subsequence.go)
- [medium_leetcode_322_coin_change.go](medium_leetcode_322_coin_change.go)
- [medium_leetcode_33_search_in_rotated_sorted_array.go](medium_leetcode_33_search_in_rotated_sorted_array.go)
- [medium_leetcode_34_find_first_and_last_position_of_element_in_sorted_array.go](medium_leetcode_34_find_first_and_last_position_of_element_in_sorted_array.go)
- [medium_leetcode_36_valid_sudoku.go](medium_leetcode_36_valid_sudoku.go)
- [medium_leetcode_373_find_k_pairs_with_smallest_sums.go](medium_leetcode_373_find_k_pairs_with_smallest_sums.go)
- [medium_leetcode_380_insert_delete_getrandom_o1.go](medium_leetcode_380_insert_delete_getrandom_o1.go)
- [medium_leetcode_399_evaluate_division.go](medium_leetcode_399_evaluate_division.go)
- [medium_leetcode_39_combination_sum.go](medium_leetcode_39_combination_sum.go)
- [medium_leetcode_3_longest_substring_without_repeating_characters_practice.go](medium_leetcode_3_longest_substring_without_repeating_characters_practice.go)
- [medium_leetcode_427_construct_quad_tree.go](medium_leetcode_427_construct_quad_tree.go)
- [medium_leetcode_433_minimum_genetic_mutation.go](medium_leetcode_433_minimum_genetic_mutation.go)
- [medium_leetcode_452_minimum_number_of_arrows_to_burst_balloons.go](medium_leetcode_452_minimum_number_of_arrows_to_burst_balloons.go)
- [medium_leetcode_45_jump_game_ii.go](medium_leetcode_45_jump_game_ii.go)
- [medium_leetcode_46_permutations.go](medium_leetcode_46_permutations.go)
- [medium_leetcode_48_rotate_image.go](medium_leetcode_48_rotate_image.go)
- [medium_leetcode_49_group_anagrams.go](medium_leetcode_49_group_anagrams.go)
- [medium_leetcode_50_powx_n.go](medium_leetcode_50_powx_n.go)
- [medium_leetcode_53_maximum_subarray.go](medium_leetcode_53_maximum_subarray.go)
- [medium_leetcode_54_spiral_matrix.go](medium_leetcode_54_spiral_matrix.go)
- [medium_leetcode_55_jump_game.go](medium_leetcode_55_jump_game.go)
- [medium_leetcode_56_merge_intervals.go](medium_leetcode_56_merge_intervals.go)
- [medium_leetcode_57_insert_interval.go](medium_leetcode_57_insert_interval.go)
- [medium_leetcode_5_longest_palindromic_substring.go](medium_leetcode_5_longest_palindromic_substring.go)
- [medium_leetcode_61_rotate_list.go](medium_leetcode_61_rotate_list.go)
- [medium_leetcode_63_unique_paths_ii.go](medium_leetcode_63_unique_paths_ii.go)
- [medium_leetcode_64_minimum_path_sum.go](medium_leetcode_64_minimum_path_sum.go)
- [medium_leetcode_6_zigzag_conversion.go](medium_leetcode_6_zigzag_conversion.go)
- [medium_leetcode_71_simplify_path.go](medium_leetcode_71_simplify_path.go)
- [medium_leetcode_72_edit_distance.go](medium_leetcode_72_edit_distance.go)
- [medium_leetcode_73_set_matrix_zeroes.go](medium_leetcode_73_set_matrix_zeroes.go)
- [medium_leetcode_74_search_a_2d_matrix.go](medium_leetcode_74_search_a_2d_matrix.go)
- [medium_leetcode_77_combinations.go](medium_leetcode_77_combinations.go)
- [medium_leetcode_79_word_search.go](medium_leetcode_79_word_search.go)
- [medium_leetcode_80_remove_duplicates_from_sorted_array_ii.go](medium_leetcode_80_remove_duplicates_from_sorted_array_ii.go)
- [medium_leetcode_82_remove_duplicates_from_sorted_list_ii.go](medium_leetcode_82_remove_duplicates_from_sorted_list_ii.go)
- [medium_leetcode_86_partition_list.go](medium_leetcode_86_partition_list.go)
- [medium_leetcode_909_snakes_and_ladders.go](medium_leetcode_909_snakes_and_ladders.go)
- [medium_leetcode_918_maximum_sum_circular_subarray.go](medium_leetcode_918_maximum_sum_circular_subarray.go)
- [medium_leetcode_92_reverse_linked_list_ii.go](medium_leetcode_92_reverse_linked_list_ii.go)
- [medium_leetcode_97_interleaving_string.go](medium_leetcode_97_interleaving_string.go)
- [medium_leetcode_98_validate_binary_search_tree.go](medium_leetcode_98_validate_binary_search_tree.go)
- [solved_easy_leetcode_88_merge_sorted_array.go](solved_easy_leetcode_88_merge_sorted_array.go)
