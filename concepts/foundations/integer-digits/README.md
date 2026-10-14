# Integer Digit Processing

Decimal digits can be processed without converting an integer to a string.
For a nonnegative value, `value % 10` extracts the final digit and `value / 10`
removes it. A loop can therefore inspect digits from right to left.

## Reversing digits

To append an extracted digit to another integer:

```go
reversed = reversed*10 + value%10
value /= 10
```

After `k` iterations, `value` is the original number with its last `k` digits
removed, while `reversed` contains those digits in reverse order. When `value`
reaches zero, `reversed` is the complete decimal reversal.

## Integer safety

Check `reversed > (limit-digit)/10` before multiplying and adding when the
reversed value might exceed its integer type. A wider intermediate is enough
when the input contract guarantees that every possible reversal fits it.

Negative values need an explicit contract: their sign is not a decimal digit.
Reversing `120` produces `21`, because integers do not preserve leading zeroes.

## Complexity

Processing every digit takes Theta(d) time for a `d`-digit value and O(1)
auxiliary space. Reversing only half the digits has the same asymptotic bounds
with roughly half as many iterations.

## Related problems

- [LeetCode 9: Palindrome Number](../../../problems/solved_easy_leetcode_0009_palindrome_number.go)
  compares full reversal with the original and also demonstrates half reversal.
