# Algorithms and Data Structures in Go

## Contents

- [A. Data structures](#a-data-structures)
  - [A.1 Arrays and slices](
    #a1-arrays-and-slices)
  - [A.2 Strings](
    #a2-strings)
  - [A.3 Linked lists](
    #a3-linked-lists)
  - [A.4 Stacks](
    #a4-stacks)
  - [A.5 Queues and deques](
    #a5-queues-and-deques)
  - [A.6 Hash maps and hash sets](
    #a6-hash-maps-and-hash-sets)
  - [A.7 Trees](
    #a7-trees)
  - [A.8 Heaps and priority queues](
    #a8-heaps-and-priority-queues)
  - [A.9 Tries](
    #a9-tries)
  - [A.10 Graphs](
    #a10-graphs)
  - [A.11 Disjoint sets](
    #a11-disjoint-sets)
  - [A.12 Range-query structures](
    #a12-range-query-structures)
  - [A.13 Bit sets and bitmasks](
    #a13-bit-sets-and-bitmasks)
  - [A.14 Composite structures](
    #a14-composite-structures)
  - [A.15 Ordered maps, sets, and skip lists](
    #a15-ordered-maps-sets-and-skip-lists)
  - [A.16 String indexes (advanced)](
    #a16-string-indexes-advanced)
  - [A.17 Spatial structures (advanced)](
    #a17-spatial-structures-advanced)
  - [A.18 Probabilistic structures (advanced)](
    #a18-probabilistic-structures-advanced)
  - [A.19 Persistent and immutable structures (advanced)](
    #a19-persistent-and-immutable-structures-advanced)
  - [A.20 Succinct and specialized structures (advanced)](
    #a20-succinct-and-specialized-structures-advanced)
- [B. Interview question families](#b-interview-question-families)
  - [B.1 Membership, duplicates, and frequencies](
    #b1-membership-duplicates-and-frequencies)
  - [B.2 Pairs, triplets, and target sums](
    #b2-pairs-triplets-and-target-sums)
  - [B.3 Ordering, ranking, and selection](
    #b3-ordering-ranking-and-selection)
  - [B.4 Searching sorted data and answer boundaries](
    #b4-searching-sorted-data-and-answer-boundaries)
  - [B.5 Contiguous subarrays and substrings](
    #b5-contiguous-subarrays-and-substrings)
  - [B.6 Ranges, updates, and aggregate queries](
    #b6-ranges-updates-and-aggregate-queries)
  - [B.7 String matching and word lookup](
    #b7-string-matching-and-word-lookup)
  - [B.8 Palindromes and string transformations](
    #b8-palindromes-and-string-transformations)
  - [B.9 Parsing, expressions, and simulation](
    #b9-parsing-expressions-and-simulation)
  - [B.10 Linked-list transformations and cycles](
    #b10-linked-list-transformations-and-cycles)
  - [B.11 Tree traversal, paths, and reconstruction](
    #b11-tree-traversal-paths-and-reconstruction)
  - [B.12 Reachability, components, and graph cycles](
    #b12-reachability-components-and-graph-cycles)
  - [B.13 Shortest paths and minimum travel costs](
    #b13-shortest-paths-and-minimum-travel-costs)
  - [B.14 Dependencies and valid task ordering](
    #b14-dependencies-and-valid-task-ordering)
  - [B.15 Connecting a network at minimum cost](
    #b15-connecting-a-network-at-minimum-cost)
  - [B.16 Grid and matrix questions](
    #b16-grid-and-matrix-questions)
  - [B.17 Intervals, scheduling, and event processing](
    #b17-intervals-scheduling-and-event-processing)
  - [B.18 Nearest greater or smaller values](
    #b18-nearest-greater-or-smaller-values)
  - [B.19 Enumerating combinations and constrained choices](
    #b19-enumerating-combinations-and-constrained-choices)
  - [B.20 Counting and optimizing over repeated subproblems](
    #b20-counting-and-optimizing-over-repeated-subproblems)
  - [B.21 Numbers, digits, and exact arithmetic](
    #b21-numbers-digits-and-exact-arithmetic)
  - [B.22 Bits and compact state](
    #b22-bits-and-compact-state)
  - [B.23 Streaming and data-structure design](
    #b23-streaming-and-data-structure-design)
  - [B.24 Array rearrangement and in-place transformations](
    #b24-array-rearrangement-and-in-place-transformations)
  - [B.25 Subsequences and sequence alignment](
    #b25-subsequences-and-sequence-alignment)
  - [B.26 Greedy optimization and resource assignment](
    #b26-greedy-optimization-and-resource-assignment)
  - [B.27 Graph routes that use edges or vertices](
    #b27-graph-routes-that-use-edges-or-vertices)
  - [B.28 Graph matching, flow, and cuts (advanced)](
    #b28-graph-matching-flow-and-cuts-advanced)
  - [B.29 Dynamic graph and tree queries (advanced)](
    #b29-dynamic-graph-and-tree-queries-advanced)
  - [B.30 Geometry and spatial queries](
    #b30-geometry-and-spatial-queries)
  - [B.31 Randomization, sampling, and probability](
    #b31-randomization-sampling-and-probability)
  - [B.32 Games and adversarial choices](
    #b32-games-and-adversarial-choices)
  - [B.33 Puzzles, state search, and transformations](
    #b33-puzzles-state-search-and-transformations)
  - [B.34 Counting with combinatorial identities](
    #b34-counting-with-combinatorial-identities)
  - [B.35 Advanced arithmetic and algebra](
    #b35-advanced-arithmetic-and-algebra)
  - [B.36 Offline processing and query reordering](
    #b36-offline-processing-and-query-reordering)
  - [B.37 External-memory and large-data questions](
    #b37-external-memory-and-large-data-questions)
  - [B.38 Encoding, compression, and serialization](
    #b38-encoding-compression-and-serialization)
  - [B.39 Advanced optimization of dynamic programming](
    #b39-advanced-optimization-of-dynamic-programming)
- [C. Choosing algorithms](#c-choosing-algorithms)
  - [C.1 Clarify the question and required output](
    #c1-clarify-the-question-and-required-output)
  - [C.2 Translate constraints into a work budget](
    #c2-translate-constraints-into-a-work-budget)
  - [C.3 Inspect input guarantees](
    #c3-inspect-input-guarantees)
  - [C.4 Start with a correct baseline](
    #c4-start-with-a-correct-baseline)
  - [C.5 Recognize conditions that enable an algorithm](
    #c5-recognize-conditions-that-enable-an-algorithm)
  - [C.6 Choose the next item to process](
    #c6-choose-the-next-item-to-process)
  - [C.7 Compare algorithms for the same question](
    #c7-compare-algorithms-for-the-same-question)
  - [C.8 Match data structures to required operations](
    #c8-match-data-structures-to-required-operations)
  - [C.9 Combine algorithms and data structures](
    #c9-combine-algorithms-and-data-structures)
  - [C.10 Prove the choice and test its assumptions](
    #c10-prove-the-choice-and-test-its-assumptions)
  - [C.11 Account for Go implementation costs](
    #c11-account-for-go-implementation-costs)
  - [C.12 Explain and validate an interview solution](
    #c12-explain-and-validate-an-interview-solution)
  - [C.13 Model the question before choosing a technique](
    #c13-model-the-question-before-choosing-a-technique)
  - [C.14 Recognize limits and avoid false pattern matches](
    #c14-recognize-limits-and-avoid-false-pattern-matches)
  - [C.15 Select advanced techniques only when justified](
    #c15-select-advanced-techniques-only-when-justified)
  - [C.16 Prioritize study and transfer knowledge](
    #c16-prioritize-study-and-transfer-knowledge)

## A. Data structures

### A.1 Arrays and slices

#### A.1.1 Fixed arrays and dynamic slices

#### A.1.2 Multidimensional arrays and matrices

#### A.1.3 Sparse matrices

#### A.1.4 Flat storage and multidimensional indexing

### A.2 Strings

#### A.2.1 Bytes, runes, and character sequences

#### A.2.2 Mutable character buffers

#### A.2.3 String builders

### A.3 Linked lists

#### A.3.1 Singly linked lists

#### A.3.2 Doubly linked lists

#### A.3.3 Circular linked lists

### A.4 Stacks

#### A.4.1 Array-backed and linked stacks

#### A.4.2 Monotonic stacks

#### A.4.3 Stacks with cached aggregates

### A.5 Queues and deques

#### A.5.1 Queues and circular buffers

#### A.5.2 Double-ended queues

#### A.5.3 Array-backed and linked queues

#### A.5.4 Monotonic deques

#### A.5.5 Double-ended priority queues

### A.6 Hash maps and hash sets

#### A.6.1 Direct-address tables

#### A.6.2 Separate chaining and open addressing

#### A.6.3 Multisets and frequency maps

#### A.6.4 Composite keys and custom hashes

### A.7 Trees

#### A.7.1 Binary trees

#### A.7.2 Binary search trees

#### A.7.3 Balanced search trees

#### A.7.4 General rooted trees

#### A.7.5 AVL and red-black trees

#### A.7.6 Treaps (advanced)

#### A.7.7 Order-statistic trees

#### A.7.8 B-trees and B+ trees (advanced)

#### A.7.9 Threaded trees and parent links

### A.8 Heaps and priority queues

#### A.8.1 Min-heaps and max-heaps

#### A.8.2 Binary and d-ary heaps

#### A.8.3 Indexed heaps

#### A.8.4 Meldable heaps (advanced)

### A.9 Tries

#### A.9.1 Prefix trees

#### A.9.2 Compressed radix trees

#### A.9.3 Binary tries for integer keys

#### A.9.4 Ternary search trees (advanced)

### A.10 Graphs

#### A.10.1 Directed and undirected graphs

#### A.10.2 Weighted and unweighted graphs

#### A.10.3 Adjacency lists, matrices, and edge lists

#### A.10.4 Bipartite graphs and directed acyclic graphs

#### A.10.5 Multigraphs and self-loops

#### A.10.6 Implicit graphs and state spaces

#### A.10.7 Reverse adjacency and residual graphs

### A.11 Disjoint sets

#### A.11.1 Parent and size arrays

#### A.11.2 Weighted and parity disjoint sets

#### A.11.3 Rollback disjoint sets (advanced)

### A.12 Range-query structures

#### A.12.1 Prefix-sum arrays

#### A.12.2 Difference arrays

#### A.12.3 Fenwick trees

#### A.12.4 Segment trees

#### A.12.5 Sparse tables

#### A.12.6 Lazy segment trees

#### A.12.7 Dynamic and persistent segment trees (advanced)

#### A.12.8 Two-dimensional range structures (advanced)

#### A.12.9 Square-root decomposition

### A.13 Bit sets and bitmasks

#### A.13.1 Packed bit vectors

#### A.13.2 Multiword bit sets

### A.14 Composite structures

#### A.14.1 Hash map with doubly linked list

#### A.14.2 Heap with hash map

#### A.14.3 Two heaps

#### A.14.4 Two stacks

#### A.14.5 Frequency buckets with linked lists

#### A.14.6 Indexed arrays with hash maps

### A.15 Ordered maps, sets, and skip lists

#### A.15.1 Ordered maps and sets

#### A.15.2 Ordered multisets

#### A.15.3 Skip lists (advanced)

### A.16 String indexes (advanced)

#### A.16.1 Suffix arrays and longest-common-prefix arrays

#### A.16.2 Suffix trees

#### A.16.3 Suffix automata

#### A.16.4 Aho-Corasick automata

#### A.16.5 Palindromic trees

### A.17 Spatial structures (advanced)

#### A.17.1 Quadtrees and octrees

#### A.17.2 k-d trees

#### A.17.3 Interval trees

#### A.17.4 R-trees

### A.18 Probabilistic structures (advanced)

#### A.18.1 Bloom filters

#### A.18.2 Count-Min sketches

#### A.18.3 HyperLogLog sketches

### A.19 Persistent and immutable structures (advanced)

#### A.19.1 Path-copying trees

#### A.19.2 Persistent lists and stacks

#### A.19.3 Versioned maps

### A.20 Succinct and specialized structures (advanced)

#### A.20.1 Rank and select bit vectors

#### A.20.2 Wavelet trees

#### A.20.3 Li Chao trees

#### A.20.4 Dancing links

## B. Interview question families

### B.1 Membership, duplicates, and frequencies

#### B.1.1 Frequency counting (hash map)

#### B.1.2 Membership tracking (hash set)

#### B.1.3 Sorting and scanning (slice)

#### B.1.4 Bounded-value counting (array)

#### B.1.5 Boyer-Moore majority vote (counters)

#### B.1.6 Missing values by placement (slice)

### B.2 Pairs, triplets, and target sums

#### B.2.1 Complement lookup (hash map)

#### B.2.2 Two pointers after sorting (slice)

#### B.2.3 Meet in the middle (slices and hash maps)

#### B.2.4 Generalized k-sum (sorted slice and recursion)

#### B.2.5 Counting pairs (frequency map)

### B.3 Ordering, ranking, and selection

#### B.3.1 Comparison sorting (slice)

#### B.3.2 Counting and radix sorting (arrays)

#### B.3.3 Quickselect (slice)

#### B.3.4 Top-k selection (heap)

#### B.3.5 Merging sorted sources (heap or pointers)

#### B.3.6 Merge sort (slice and buffer)

#### B.3.7 Quicksort (slice and partition bounds)

#### B.3.8 Heapsort (array-backed heap)

#### B.3.9 Insertion and selection sort (slice)

#### B.3.10 Bucket sort (bucket arrays)

#### B.3.11 Inversion counting (merge buffer or Fenwick tree)

### B.4 Searching sorted data and answer boundaries

#### B.4.1 Binary search and bounds (slice)

#### B.4.2 Rotated-array search (slice)

#### B.4.3 Binary search on the answer (feasibility state)

#### B.4.4 Exponential search (indexed data)

#### B.4.5 Peak finding (slice or matrix)

#### B.4.6 Selection across sorted arrays (partition bounds)

### B.5 Contiguous subarrays and substrings

#### B.5.1 Fixed sliding window (slice or counters)

#### B.5.2 Variable sliding window (map or counters)

#### B.5.3 Prefix sums with lookup (array and hash map)

#### B.5.4 Kadane's algorithm (running totals)

#### B.5.5 Window extrema (monotonic deque)

#### B.5.6 Product extrema (minimum and maximum state)

#### B.5.7 Shortest signed-sum range (prefix sums and deque)

#### B.5.8 Exactly-k counting (two windows and frequency map)

#### B.5.9 Circular subarrays (slice and running totals)

### B.6 Ranges, updates, and aggregate queries

#### B.6.1 Static range sums (prefix-sum array)

#### B.6.2 Batch range updates (difference array)

#### B.6.3 Point updates and sums (Fenwick tree)

#### B.6.4 Range queries and updates (segment tree)

#### B.6.5 Static range minima (sparse table)

#### B.6.6 Coordinate compression (sorted slice and map)

#### B.6.7 Offline queries with Mo's algorithm (blocks)

#### B.6.8 Range order statistics (wavelet or persistent tree)

### B.7 String matching and word lookup

#### B.7.1 Common-prefix scanning (strings)

#### B.7.2 KMP matching (prefix-function array)

#### B.7.3 Rabin-Karp matching (rolling hash)

#### B.7.4 Prefix and dictionary lookup (trie)

#### B.7.5 Anagram grouping (hash map and counters)

#### B.7.6 Z-algorithm matching (Z-array)

#### B.7.7 Multi-pattern matching (Aho-Corasick automaton)

#### B.7.8 Repeated substrings (suffix array and LCP array)

#### B.7.9 Wildcard and regex matching (DP table or automaton)

### B.8 Palindromes and string transformations

#### B.8.1 Two pointers and center expansion (string)

#### B.8.2 Palindrome dynamic programming (table)

#### B.8.3 Edit distance and subsequences (DP table)

#### B.8.4 Manacher's algorithm (radius arrays)

#### B.8.5 Palindrome partitioning (backtracking or DP table)

#### B.8.6 Character reorganization (counts and heap)

### B.9 Parsing, expressions, and simulation

#### B.9.1 State-machine parsing (state and counters)

#### B.9.2 Bracket matching (stack)

#### B.9.3 Expression evaluation (stacks)

#### B.9.4 Rule-based simulation (arrays, maps, or queues)

#### B.9.5 Tokenization and recursive descent (tokens and stack)

#### B.9.6 Shunting-yard conversion (operator stack)

#### B.9.7 Nested decoding (stack and string buffer)

### B.10 Linked-list transformations and cycles

#### B.10.1 Reversal, splicing, and merging (node pointers)

#### B.10.2 Fast and slow pointers (linked list)

#### B.10.3 Cycle detection with memory (hash set)

#### B.10.4 Intersection by pointer switching (node pointers)

#### B.10.5 Copying random pointers (map or interleaved nodes)

#### B.10.6 Merge sorting lists (node pointers)

#### B.10.7 Reordering and group reversal (node pointers)

### B.11 Tree traversal, paths, and reconstruction

#### B.11.1 Depth-first traversal (stack or recursion)

#### B.11.2 Breadth-first traversal (queue)

#### B.11.3 Subtree aggregation (recursive return values)

#### B.11.4 Search-tree validation (bounds or stack)

#### B.11.5 Tree reconstruction (index map and recursion)

#### B.11.6 Lowest common ancestor (tree or ancestor table)

#### B.11.7 Serialization and deserialization (queue or stack)

#### B.11.8 Tree diameter and maximum path (subtree state)

#### B.11.9 Tree views and vertical order (queue and map)

#### B.11.10 Morris traversal (temporary tree links)

#### B.11.11 Search-tree rank queries (subtree counts)

#### B.11.12 Rerooting DP (per-node arrays)

### B.12 Reachability, components, and graph cycles

#### B.12.1 Depth-first search (adjacency list and stack)

#### B.12.2 Breadth-first search (adjacency list and queue)

#### B.12.3 Component merging (disjoint sets)

#### B.12.4 Directed-cycle detection (color array)

#### B.12.5 Bipartite checking (color array and queue)

#### B.12.6 Strongly connected components (stacks and arrays)

#### B.12.7 Bridges and articulation points (DFS low-link arrays)

#### B.12.8 Functional-graph cycles (pointers and visit arrays)

#### B.12.9 Transitive closure (bit sets or boolean matrix)

### B.13 Shortest paths and minimum travel costs

#### B.13.1 Unweighted BFS (queue and distance array)

#### B.13.2 0-1 BFS (deque and distance array)

#### B.13.3 Dijkstra (heap or scan, distance array)

#### B.13.4 Bellman-Ford (edge list and distance array)

#### B.13.5 DAG shortest paths (topological order and array)

#### B.13.6 Floyd-Warshall (distance matrix)

#### B.13.7 Multi-source BFS (queue and distance array)

#### B.13.8 Bidirectional BFS (two frontiers and maps)

#### B.13.9 A* search (heap and cost map)

#### B.13.10 Bottleneck paths (heap or disjoint sets)

#### B.13.11 State-augmented paths (distance table and queue)

### B.14 Dependencies and valid task ordering

#### B.14.1 Kahn's topological sort (queue and indegrees)

#### B.14.2 DFS topological sort (stack and visit states)

#### B.14.3 Dependency layers (queue and level array)

#### B.14.4 Lexicographic topological order (heap)

#### B.14.5 Prerequisite queries (bit sets or reachability table)

### B.15 Connecting a network at minimum cost

#### B.15.1 Kruskal's algorithm (sorted edges and disjoint sets)

#### B.15.2 Prim's algorithm (adjacency list and heap)

#### B.15.3 Implicit-edge MST (heap or dense arrays)

### B.16 Grid and matrix questions

#### B.16.1 Directional and spiral scans (matrix and bounds)

#### B.16.2 In-place transformations (matrix)

#### B.16.3 Flood fill and islands (stack, queue, or disjoint sets)

#### B.16.4 Shortest grid paths (queue or heap)

#### B.16.5 Grid dynamic programming (table or rolling row)

#### B.16.6 Simultaneous updates (copied or encoded matrix)

#### B.16.7 Sorted-matrix search (row and column pointers)

#### B.16.8 Two-dimensional sums (prefix matrix)

#### B.16.9 Largest square (DP table or rolling row)

#### B.16.10 Largest rectangle (heights and monotonic stack)

### B.17 Intervals, scheduling, and event processing

#### B.17.1 Interval merging (sorted slice)

#### B.17.2 Interval selection (sorted slice)

#### B.17.3 Resource allocation (min-heap)

#### B.17.4 Sweep line (sorted events and active state)

#### B.17.5 Weighted scheduling (sorted slice and DP array)

#### B.17.6 Cooldown scheduling (heap and queue)

#### B.17.7 Deadline scheduling (sorted jobs and heap)

#### B.17.8 Interval stabbing and coverage (sorted endpoints)

### B.18 Nearest greater or smaller values

#### B.18.1 Monotonic stack (index stack)

#### B.18.2 Histogram boundaries (monotonic stack)

#### B.18.3 Rainwater trapping (two pointers or stack)

#### B.18.4 Circular nearest values (index stack)

### B.19 Enumerating combinations and constrained choices

#### B.19.1 Subsets and permutations (backtracking and path)

#### B.19.2 Constraint search (backtracking and sets)

#### B.19.3 Pruned search (bounds and partial state)

#### B.19.4 Subset enumeration (bitmasks)

#### B.19.5 Duplicate-aware generation (counts or sorted slice)

#### B.19.6 Board word search (trie and visited state)

#### B.19.7 Exact cover (dancing links, advanced)

### B.20 Counting and optimizing over repeated subproblems

#### B.20.1 Memoization (map and recursion)

#### B.20.2 Tabulation and rolling state (arrays)

#### B.20.3 Knapsack and subset sum (DP table or bit set)

#### B.20.4 Sequence DP (arrays or binary-search tails)

#### B.20.5 Interval DP (two-dimensional table)

#### B.20.6 Tree and DAG DP (per-node state)

#### B.20.7 Bitmask DP (state table)

#### B.20.8 Digit DP (position and bound state)

#### B.20.9 Finite-state DP (state array)

#### B.20.10 Probability DP (numeric table)

#### B.20.11 Profile DP (frontier bitmasks, advanced)

#### B.20.12 DP path reconstruction (predecessor array)

### B.21 Numbers, digits, and exact arithmetic

#### B.21.1 Digit extraction and carry handling (integers or lists)

#### B.21.2 Euclid's algorithm (integer remainders)

#### B.21.3 Factorization and sieves (integers and arrays)

#### B.21.4 Fast exponentiation (integers)

#### B.21.5 Modular arithmetic (integers)

#### B.21.6 Normalized ratios and slopes (hash-map keys)

#### B.21.7 Large-number processing (strings or big integers)

#### B.21.8 Combinatorial counting (integers or DP arrays)

#### B.21.9 Overflow-safe reversal (integer bounds)

#### B.21.10 Factorial trailing zeroes (integer quotients)

#### B.21.11 Integer square roots (binary search bounds)

#### B.21.12 Rational arithmetic (normalized integer pairs)

#### B.21.13 Extended Euclid and modular inverses (integers)

#### B.21.14 Matrix exponentiation (matrices)

#### B.21.15 Chinese remainder theorem (residue arrays, advanced)

### B.22 Bits and compact state

#### B.22.1 Bit tests and population counts (integer masks)

#### B.22.2 XOR cancellation (integer accumulator)

#### B.22.3 Subset transitions (bitmasks)

#### B.22.4 Maximum XOR queries (binary trie)

#### B.22.5 Bitwise range aggregates (bit counts or segment tree)

#### B.22.6 XOR linear basis (basis array, advanced)

### B.23 Streaming and data-structure design

#### B.23.1 Running statistics (counters)

#### B.23.2 Streaming median (two heaps)

#### B.23.3 Streaming top-k (heap)

#### B.23.4 Reservoir sampling (sample array)

#### B.23.5 LRU cache (hash map and doubly linked list)

#### B.23.6 Min stack (paired values or two stacks)

#### B.23.7 Queue from stacks (two stacks)

#### B.23.8 Randomized set (slice and hash map)

#### B.23.9 LFU cache (hash maps and frequency buckets)

#### B.23.10 Time-based key store (map and sorted histories)

#### B.23.11 Sliding event counts (queue or time buckets)

#### B.23.12 Stack from queues (queues)

#### B.23.13 Iterators and flattening (stack or cursor state)

### B.24 Array rearrangement and in-place transformations

#### B.24.1 Stable compaction (read and write indexes)

#### B.24.2 Dutch national flag (three partition boundaries)

#### B.24.3 Cyclic placement (slice)

#### B.24.4 Rotation by reversal (slice)

#### B.24.5 Next permutation (slice)

#### B.24.6 Prefix and suffix products (arrays)

#### B.24.7 Shuffle and partition (slice)

### B.25 Subsequences and sequence alignment

#### B.25.1 Subsequence checking (two pointers)

#### B.25.2 Repeated subsequence queries (position lists)

#### B.25.3 Longest common subsequence (DP table)

#### B.25.4 Longest increasing subsequence (tails array)

#### B.25.5 Counting subsequences (DP array or Fenwick tree)

#### B.25.6 Patience sorting (pile tops)

### B.26 Greedy optimization and resource assignment

#### B.26.1 Activity selection (sorted intervals)

#### B.26.2 Jump reachability (farthest-reachable index)

#### B.26.3 Gas-station circuit (running balances)

#### B.26.4 Job assignment (sorting and heap)

#### B.26.5 Huffman coding (min-heap and binary tree)

#### B.26.6 Exchange-based selection (sorted candidates)

### B.27 Graph routes that use edges or vertices

#### B.27.1 Euler trails with Hierholzer (adjacency and stack)

#### B.27.2 Hamiltonian paths (backtracking or bitmask DP)

#### B.27.3 Itinerary reconstruction (ordered adjacency)

#### B.27.4 De Bruijn sequences (implicit graph, advanced)

### B.28 Graph matching, flow, and cuts (advanced)

#### B.28.1 Augmenting-path matching (adjacency and match arrays)

#### B.28.2 Hopcroft-Karp (layers and matching arrays)

#### B.28.3 Edmonds-Karp max flow (residual graph and queue)

#### B.28.4 Dinic max flow (residual graph and level arrays)

#### B.28.5 Minimum cut (residual reachability)

#### B.28.6 Min-cost flow (residual graph and potentials)

#### B.28.7 Assignment with Hungarian algorithm (cost matrix)

### B.29 Dynamic graph and tree queries (advanced)

#### B.29.1 Offline connectivity (rollback disjoint sets)

#### B.29.2 Binary lifting (ancestor table)

#### B.29.3 Euler-tour flattening (array and range structure)

#### B.29.4 Heavy-light decomposition (segment tree)

#### B.29.5 Centroid decomposition (tree and distance state)

### B.30 Geometry and spatial queries

#### B.30.1 Orientation and intersection (point coordinates)

#### B.30.2 Collinearity (normalized slope map)

#### B.30.3 Convex hull (sorted points and stack)

#### B.30.4 Closest pair (sorted points and merge buffer)

#### B.30.5 Rectangle union (sweep events and segment tree)

#### B.30.6 Point-in-polygon (edges and crossing count)

#### B.30.7 Nearest-neighbor queries (k-d tree, advanced)

### B.31 Randomization, sampling, and probability

#### B.31.1 Fisher-Yates shuffle (slice)

#### B.31.2 Weighted sampling (prefix sums and binary search)

#### B.31.3 Alias sampling (probability and alias arrays)

#### B.31.4 Rejection sampling (random draws)

#### B.31.5 Randomized quickselect (slice)

#### B.31.6 Monte Carlo estimation (sample counters)

### B.32 Games and adversarial choices

#### B.32.1 Winning-state DP (state table)

#### B.32.2 Minimax (game tree and recursion)

#### B.32.3 Alpha-beta pruning (game tree and bounds)

#### B.32.4 Sprague-Grundy values (arrays and sets, advanced)

### B.33 Puzzles, state search, and transformations

#### B.33.1 Shortest transformation (BFS and state map)

#### B.33.2 Search with keys or masks (queue and visited table)

#### B.33.3 Iterative deepening (stack and depth bound)

#### B.33.4 Branch and bound (heap or stack and bounds)

#### B.33.5 Symmetry reduction (canonical state keys)

### B.34 Counting with combinatorial identities

#### B.34.1 Binomial coefficients (integer recurrence or DP)

#### B.34.2 Permutation rank and unrank (Fenwick tree)

#### B.34.3 Catalan counting (DP array)

#### B.34.4 Inclusion-exclusion (subset masks)

#### B.34.5 Pigeonhole arguments (counts and buckets)

### B.35 Advanced arithmetic and algebra

#### B.35.1 Segmented sieve (block arrays)

#### B.35.2 Primality testing (modular integers)

#### B.35.3 Gaussian elimination (matrix)

#### B.35.4 Polynomial multiplication (FFT or NTT arrays)

#### B.35.5 Linear recurrences (matrix or coefficient arrays)

### B.36 Offline processing and query reordering

#### B.36.1 Sort queries by threshold (disjoint sets or Fenwick)

#### B.36.2 Reverse processing (disjoint sets or hash maps)

#### B.36.3 Parallel binary search (query buckets, advanced)

#### B.36.4 Divide-and-conquer queries (buffers, advanced)

### B.37 External-memory and large-data questions

#### B.37.1 External merge sort (sorted runs and heap)

#### B.37.2 Partitioned aggregation (hash buckets)

#### B.37.3 Approximate membership (Bloom filter)

#### B.37.4 Approximate frequencies (Count-Min sketch)

#### B.37.5 Approximate distinct counts (HyperLogLog)

### B.38 Encoding, compression, and serialization

#### B.38.1 Run-length encoding (scan and output buffer)

#### B.38.2 Prefix coding (Huffman tree)

#### B.38.3 Length-prefixed serialization (byte buffer)

#### B.38.4 Variable-length integer coding (byte buffer)

#### B.38.5 Dictionary compression (map or trie, advanced)

### B.39 Advanced optimization of dynamic programming

#### B.39.1 Monotonic queue optimization (deque)

#### B.39.2 Convex hull trick (line hull or Li Chao tree)

#### B.39.3 Divide-and-conquer DP optimization (arrays)

#### B.39.4 Knuth optimization (cost and split tables)

#### B.39.5 Subset transforms (bitmask arrays)

## C. Choosing algorithms

### C.1 Clarify the question and required output

#### C.1.1 Existence, count, optimum, one answer, or all answers

#### C.1.2 Subarray, subsequence, subset, or permutation

### C.2 Translate constraints into a work budget

#### C.2.1 Input size and expected time complexity

#### C.2.2 Output size and unavoidable work

#### C.2.3 Memory limits and recursion depth

#### C.2.4 Worst-case, expected, and amortized work

#### C.2.5 Exponential search and practical state counts

### C.3 Inspect input guarantees

#### C.3.1 Sorted order, duplicates, and bounded values

#### C.3.2 Negative values, overflow, and exactness

#### C.3.3 Directed edges, weights, and cycles

#### C.3.4 Mutability, precision, and character semantics

#### C.3.5 Online arrivals and changing data

#### C.3.6 Queries versus updates and repeated workloads

### C.4 Start with a correct baseline

#### C.4.1 Find repeated work and expensive operations

#### C.4.2 Choose what to remember or preprocess

#### C.4.3 Time-space tradeoffs and caching

#### C.4.4 Information needed to reconstruct the answer

### C.5 Recognize conditions that enable an algorithm

#### C.5.1 Monotonicity for binary search

#### C.5.2 Window validity for sliding windows

#### C.5.3 Repeated subproblems for dynamic programming

#### C.5.4 Safe permanent choices for greedy algorithms

#### C.5.5 Associativity and invertibility for range queries

#### C.5.6 Partial orders and acyclic dependencies

#### C.5.7 Exchange arguments and optimal substructure

### C.6 Choose the next item to process

#### C.6.1 Depth-first, breadth-first, cheapest, or dependency-ready

#### C.6.2 Stacks, queues, heaps, and linear scans

### C.7 Compare algorithms for the same question

#### C.7.1 Hashing versus sorting and two pointers

#### C.7.2 Prefix sums versus sliding windows

#### C.7.3 Backtracking versus dynamic programming

#### C.7.4 Greedy versus dynamic programming

#### C.7.5 BFS versus Dijkstra versus Bellman-Ford

#### C.7.6 Heap versus sorting versus quickselect

#### C.7.7 Fenwick tree versus segment tree versus sparse table

#### C.7.8 DFS versus disjoint sets for connectivity

#### C.7.9 Trie versus hashing versus string matching

#### C.7.10 Exact versus randomized or approximate algorithms

### C.8 Match data structures to required operations

#### C.8.1 Lookup, ordered access, insertion, and deletion

#### C.8.2 Static, dynamic, offline, and streaming workloads

#### C.8.3 In-place updates and ownership of input

### C.9 Combine algorithms and data structures

#### C.9.1 Sorting followed by scanning or greedy selection

#### C.9.2 Binary search with a feasibility algorithm

#### C.9.3 Graph traversal with dynamic programming

### C.10 Prove the choice and test its assumptions

#### C.10.1 State meanings, invariants, and termination

#### C.10.2 Counterexamples and adversarial inputs

#### C.10.3 Empty inputs, boundaries, duplicates, and ties

#### C.10.4 Cross-check optimized code against brute force

#### C.10.5 Property-based and metamorphic checks

### C.11 Account for Go implementation costs

#### C.11.1 Allocation, copying, and retained slice storage

#### C.11.2 Map presence, iteration order, and key types

#### C.11.3 Bytes, runes, and integer limits

### C.12 Explain and validate an interview solution

#### C.12.1 Describe the approach before coding

#### C.12.2 Trace concrete input through the state

#### C.12.3 Compare time and memory tradeoffs

#### C.12.4 Test the implementation against the contract

### C.13 Model the question before choosing a technique

#### C.13.1 Turn transformations into graph edges

#### C.13.2 Turn histories into compact states

#### C.13.3 Turn ranges into endpoint events

#### C.13.4 Turn equivalence into canonical keys

#### C.13.5 Separate decision problems from optimization

### C.14 Recognize limits and avoid false pattern matches

#### C.14.1 Why negative values can break a sliding window

#### C.14.2 Why negative edges can break Dijkstra

#### C.14.3 Why local improvement does not prove a greedy choice

#### C.14.4 Why similar wording can require different algorithms

#### C.14.5 Recognize exponential or intractable general cases

### C.15 Select advanced techniques only when justified

#### C.15.1 Identify the simpler method's actual bottleneck

#### C.15.2 Check the specialized algorithm's prerequisites

#### C.15.3 Balance implementation risk against resource savings

### C.16 Prioritize study and transfer knowledge

#### C.16.1 Core structures and reusable algorithm families

#### C.16.2 Advanced and specialized interview topics

#### C.16.3 Practice variations that change one assumption

#### C.16.4 Track mistakes by invariant and state definition
