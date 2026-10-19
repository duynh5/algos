# Common-Prefix Scans

## Mental model

Compare several strings one character position at a time, beginning at the
left. The answer ends as soon as one string ends or one character differs,
because a shared beginning cannot skip that position.

Use this pattern only when the answer must start at position `0` in every
sequence.

## Worked example

Find the shared beginning of `"flower"`, `"flow"`, and `"flight"`:

1. Position `0` is `f` in every string, so keep `"f"`.
2. Position `1` is `l` in every string, so keep `"fl"`.
3. Position `2` is `o`, `o`, and `i`, so stop.

The longest common prefix is `"fl"`.

`LongestCommonPrefix` uses the first string as the proposed answer and checks
the same position in every other string before moving right.

## Invariant (fact that remains true)

Before checking position `c`, every string shares `strs[0][:c]`. If one string
ends at `c` or has a different character there, no longer shared beginning is
possible. If all characters match, the proven prefix grows by one character.

## Work and memory

- Work: `O(S)`, where `S` is the number of strings plus their total bytes. The
  scan stops at the first mismatch and does not inspect positions beyond the
  possible shared prefix.
- Extra memory: `O(1)`, excluding the returned string. The function returns a
  slice of the first string instead of building a new string one character at a
  time.

## Common mistakes

- Check a string's length before reading the current position.
- Decide what empty input means. This implementation returns an empty string.
- The implementation compares bytes, which is correct for lowercase English
  input. General Unicode characters may use several bytes, and one visible
  symbol may contain several encoded parts. Decide what one position should mean
  before reusing this implementation for general text.
- Repeatedly adding one character to a string can copy everything accumulated
  so far. Return the matching portion of the first string instead.

## Related problems

- [LeetCode 14: Longest Common Prefix](../../../problems/solved_easy_leetcode_0014_longest_common_prefix.go)
  compares the same character position across all strings.
