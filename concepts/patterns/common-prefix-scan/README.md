# Common-Prefix Scans

A common-prefix scan finds the longest leading segment shared by several
sequences. Recognize it when every candidate answer must begin at index zero
and the answer stops permanently at the first missing or mismatched position.

`LongestCommonPrefix` demonstrates a vertical scan over string columns. It uses
the first string as the candidate prefix and validates each byte position
against every other string before advancing.

## Invariant

Before checking column `c`, every string shares `strs[0][:c]`. If a string ends
at `c` or contains a different byte there, that proven prefix is maximal. If
every byte matches, the invariant extends through column `c`.

## Complexity

- Time: O(S), where S is the number of strings plus their total bytes. The scan
  stops at the first mismatch and never examines a byte beyond the candidate
  prefix.
- Extra space: O(1). The returned prefix is a substring view of the first
  string; callers that serialize or retain it may allocate separately.

## Pitfalls

- Check a string's length before indexing the current column.
- An empty input collection needs an explicit policy. The example returns an
  empty prefix, while some judges guarantee at least one string.
- Byte indexing is correct for lowercase English input. General Unicode text
  requires deciding whether prefixes are measured in bytes, runes, or grapheme
  clusters.
- Repeatedly appending one character to an immutable string can copy the entire
  accumulated prefix. Return a slice of the first string instead.

## Related problems

- [LeetCode 14: Longest Common Prefix](../../../problems/solved_easy_leetcode_0014_longest_common_prefix.go)
  compares each character column across all strings.
