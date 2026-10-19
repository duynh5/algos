# Integer Digit Processing

## Mental model

An integer's decimal digits can be removed from right to left without first
turning the number into text:

- `value % 10` returns the final digit;
- `value / 10` removes the final digit because integer division drops the
  remainder.

For example, `123 % 10` is `3`, and `123 / 10` is `12`.

## Worked example: reversing digits

To move each extracted digit into a new integer:

```go
reversed = reversed*10 + value%10
value /= 10
```

Starting with `value = 120` and `reversed = 0`:

| Step | Extracted digit | New `reversed` | Remaining `value` |
| --- | --- | --- | --- |
| 1 | `0` | `0` | `12` |
| 2 | `2` | `2` | `1` |
| 3 | `1` | `21` | `0` |

The result is `21`, not `021`, because an integer does not store leading zeroes.

## Invariant (fact that remains true)

After `k` passes, `value` is the original number with its final `k` digits
removed. `reversed` contains those removed digits in the opposite order. When
`value` reaches zero, every digit has moved.

## Preventing integer overflow

Multiplying by `10` may exceed the largest value the integer type can hold.
Before calculating `reversed*10 + digit`, check:

```go
reversed > (limit-digit)/10
```

If the comparison is true, the next result would exceed `limit`. The package
function returns `false` instead of performing that unsafe calculation.

Negative inputs need an explicit rule because the minus sign is not a digit.
`ReverseNonNegative` rejects them and accepts zero.

## Work and memory

For a number with `d` digits, the loop performs `d` passes, so it uses `O(d)`
work. It stores only the current digit and two integers, so it uses `O(1)` extra
memory.

## Related problems

- [LeetCode 9: Palindrome Number](../../../problems/solved_easy_leetcode_0009_palindrome_number.go)
  compares a full reversal with the original number and also shows how reversing
  only half of the digits can prove the same result.
