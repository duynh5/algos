# Algorithms and Data Structures in Go

This is the repository's guide to understanding algorithms and data structures
in Go. It explains what each idea is, when it helps, why it works, and what it
costs to run. The problem archives contain complete judge solutions; this guide
teaches the ideas that those solutions reuse.

Use the book in two directions:

- When starting a problem, begin with [method selection](#problem-solving-method-selection-guide).
- When studying, read foundations first, then data structures, patterns, and
  the preserved Go reference implementations.

Technical names such as **binary search** and **dynamic programming** are kept
because they are the names used in problems, documentation, and interviews.
When a technical term first appears, the nearby text explains it in ordinary
language. You should not need to know the term before reading its explanation.

Every concept follows the same teaching order whenever the topic needs it:

1. **Plain-language picture:** what to imagine before learning the formal name.
2. **Concrete example:** real values traced one step at a time.
3. **Formal rule:** the invariant, data structure, or algorithmic idea.
4. **Cost and limits:** how much work and memory it uses, and when it is the
   wrong tool.

An **ELI5** note appears when the formal explanation can hide the simple idea.
ELI5 means "explain it like I am five": it is an intuition aid, not a separate
or less accurate rule. Read it first, then use the precise explanation to code.

The coverage is cross-checked against the repository's
[Top Interview 150 manifest](../scripts/top-interview-150.tsv), LeetCode's
[Top Interview 150](https://leetcode.com/studyplan/top-interview-150/), the
[Grind 75 study guidance](https://www.techinterviewhandbook.org/grind75/faq),
and the [NeetCode roadmap](https://neetcode.io/roadmap). These lists determine
breadth, but the material below is organized by reusable invariants rather than
by problem title.

## Contemporary interview evidence

Current official preparation material from
[Amazon](https://www.amazon.jobs/content/en/how-we-hire/interview-prep/software-development-topics),
[Meta](https://d3no4ktch0fdq4.cloudfront.net/public/course/files/Meta_SWE_full_loop_guide.pdf),
and [Microsoft](https://careers.microsoft.com/v2/global/en/hiring-tips/technical-interviewing.html)
emphasizes applying fundamentals, comparing approaches, explaining complexity,
writing clear code, and testing it. That is stronger evidence about the skills
being assessed than any remembered list of exact questions.

Recent public candidate reports are also sampled for changes in question style,
including reports from [Amazon in 2026](https://leetcode.com/discuss/post/8420230/amazon-sde-2-interview-experience-2026-e-t1g8/),
[Google in 2025](https://leetcode.com/discuss/post/6522571/),
[Meta in 2025](https://leetcode.com/discuss/post/6460774/Meta-or-E4-SWE-or-USA-or-Feb-2025-or-Passed-Hiring-Committee/),
[Microsoft in 2025](https://leetcode.com/discuss/post/7227003/), and
[Uber in 2025](https://leetcode.com/discuss/post/7327020/). Candidate reports
are anecdotal and cannot prove that one exact question is common. This guide
therefore incorporates only repeated, transferable signals:

- Expect a familiar structure with a changed constraint: duplicates, a limited
  number of edits, streaming input, a value-distance restriction, or a demand
  for lower memory.
- Be able to move from a direct solution to a better one and explain what
  repeated work, one-direction ordering, or data-structure operation made the
  improvement possible.
- Practice constructing a graph from relationships, converting a rooted tree
  into an undirected graph when movement may go upward, and recognizing when a
  range query accelerates a dynamic-programming transition.
- Know online combinations such as counts plus a queue, two heaps, a heap over
  sorted streams, and a map plus linked list. The invariant between the two
  structures matters more than memorizing a class name.
- Treat manual verification as part of the solution: trace one ordinary case,
  one boundary case, and one adversarial case before declaring the code done.

Exact company question banks change and are often confidential. Use company
tags to choose practice inputs, not as a substitute for the concepts below.

## Contents

- [Contemporary interview evidence](#contemporary-interview-evidence)
- [Foundations](#foundations)
- [Data structures](#data-structures)
- [Problem-solving patterns](#problem-solving-patterns)
- [Go reference implementations](#go-reference-implementations)
- [Go techniques](#go-techniques)
- [Further-study checklist](#further-study-checklist)

## Foundations

### Complexity and resource accounting

Let `n` mean the amount of input, such as the number of values in a slice.
**Time complexity** describes how quickly the work grows when `n` grows.
**Extra-space complexity** describes memory created by the algorithm, not the
input it receives or the answer it must return. Count setup work, unfinished
function calls, Go allocations, and the work needed to build the answer. If an
algorithm behaves differently on easy and difficult inputs, describe its best,
usual, and worst cases separately.

The notation below compares broad growth rates rather than exact running times:

| Notation | Plain-language meaning |
| --- | --- |
| `O(1)` | Roughly the same amount of work regardless of input size. |
| `O(log n)` | Repeatedly removes a large part of the remaining search area. |
| `O(n)` | Usually one full pass over the input. |
| `O(n log n)` | Common for efficient sorting and work built around sorting. |
| `O(n^2)` | Often examines every pair of input items. |
| `O(2^n)` | Often tries every subset. It becomes impractical quickly. |
| `O(n!)` | Often tries every ordering. It becomes impractical even faster. |

The hidden constant still matters in real programs, but the maximum input size
usually tells you which growth rates are impossible before small code-level
speed differences matter.

**ELI5:** Imagine looking for a name in a phone book with 1,000 pages.
Checking every page is `O(n)`. Opening near the middle and repeatedly throwing
away the half that cannot contain the name is `O(log n)`. Comparing every page
with every other page is `O(n^2)`. If the phone book doubles in size, the first
method does about twice the work, the second needs only about one extra check,
and the third does about four times the work.

**Concrete comparison:** For 1,024 items, a full scan can inspect 1,024 items.
A halving search needs at most about 10 decisions because `2^10 = 1,024`.
Examining every ordered pair can approach 1,048,576 comparisons. Big-O ignores
small implementation details so these growth differences remain visible.

### Correctness and invariants

An **invariant** is simply a fact the algorithm keeps true while it runs. For
example, in a left-to-right scan, "everything before the current position has
already been handled correctly" may be the invariant. To explain why an
algorithm is correct, answer these questions:

1. **At the start:** Why is the fact true before any work happens?
2. **After one step:** Why does the next step keep the fact true?
3. **Moving forward:** Why does every step bring the algorithm closer to
   finishing?
4. **Nothing missed:** Why can skipped or discarded choices not contain a
   missing answer?
5. **At the end:** Why does the maintained fact make the returned answer
   correct?

**ELI5:** Think of tidying toys into a box from left to right. Your rule is:
"everything left of my hand is already in the right place." After moving one
toy, that rule must still be true. When your hand reaches the end, the rule
covers every toy. An algorithm invariant works the same way.

**Worked example:** To move all non-zero values forward in `[0, 5, 0, 2]`, let
`read` inspect each position and let `write` point to the next output position.
Before every step, `values[:write]` contains exactly the non-zero values already
seen: first `[]`, then `[5]`, and finally `[5, 2]`. The invariant explains both
why no non-zero value is lost and why their original order is preserved.

### Recursion and saved work

**Recursion** means that a function calls itself to solve a smaller part of the
same problem. Every unfinished call must remember its arguments, local values,
and where execution should continue afterward. If there can be `d` unfinished
calls, they use `O(d)` memory even when the code does not create a slice or map.
Use a stack data structure instead when the nesting may become too deep, or
when the saved work must be inspected and controlled directly.

**ELI5:** Recursion is like opening a smaller box inside a box. You cannot
finish closing the outer box until the inner one is finished, so you leave a
note on each open box saying what to do next. The computer keeps the same kind
of note for every unfinished function call.

For `factorial(4)`, the calls wait as `4 * factorial(3)`, then
`3 * factorial(2)`, then `2 * factorial(1)`. After `factorial(1)` returns `1`,
the waiting calls finish in reverse order: `2`, then `6`, then `24`. There are
four nested calls, so the call stack uses `O(n)` extra memory even though the
function creates no explicit collection.

### Integer limits and bits

An integer type can store only a limited range of values. **Overflow** happens
when arithmetic tries to go outside that range. Check before performing a
risky operation: a comparison such as `value > (limit-digit)/10` can show that
the next multiplication and addition would be too large. When working with
binary digits, first decide whether negative values are allowed and how many
bits the type stores; shifting and masking use exactly that representation.

### Bit manipulation

Bits are useful when each position is an independent yes/no flag, the input is
defined as a fixed-width binary value, or an arithmetic identity lets many
values cancel. Do not reach for bit operations merely because the inputs are
integers; ordinary arithmetic or a set is often clearer.

**ELI5:** A bitmask is a row of light switches. Switch `k` answers one yes/no
question: is item `k` selected? Turning switches on, off, or checking them lets
one integer represent many independent choices.

For three permissions (read, write, and execute), use bits `0`, `1`, and `2`.
The value `5` is binary `101`, so read and execute are on while write is off.
`mask & (1 << 1)` checks write; `mask | (1 << 1)` turns it on and changes the
mask from `101` to `111`.

Useful identities for an integer `x` include:

| Expression | Meaning |
| --- | --- |
| `x & (1 << k)` | Test bit `k`. |
| `x | (1 << k)` | Set bit `k`. |
| `x &^ (1 << k)` | Clear bit `k` in Go. |
| `x ^ (1 << k)` | Toggle bit `k`. |
| `x & (x - 1)` | Clear the lowest set bit. |
| `x & -x` | Isolate the lowest set bit in Go's signed-integer representation. |
| `a ^ a` | Produce `0`; equal values cancel under XOR. |
| `a ^ 0` | Preserve `a`. |

For example, if every value occurs twice except `7`, XORing the entire input
removes every pair because order does not matter and `a ^ a = 0`. The remaining
value is `7`. This works only because the occurrence rule matches the identity;
it does not solve a general frequency-counting problem.

A bitmask can also represent a subset: bit `k` is `1` exactly when item `k` is
included. Enumerating all masks from `0` through `(1 << n) - 1` visits all
`2^n` subsets, so the representation is compact but does not remove the
exponential work. For DP over subsets, a state such as `dp[mask]` stores the
answer for exactly the chosen items in `mask`.

In Go, use an unsigned type when a right shift must add zero bits on the left
or when an exact bit width matters. Right-shifting a signed negative value adds
one bits on the left instead. Converting between `int`, `uint32`, and `uint64`
also changes the available width. Parenthesize masks, check that the shift count
is valid, and remember that `^x` flips every bit while `x ^ y` compares two
values with XOR.

Counting set bits with `x &= x - 1` takes one pass per set bit. Scanning a fixed
32- or 64-bit word takes constant time with respect to ordinary input length,
but state the width rather than calling the work universally free.

### Keeping arithmetic values exact and consistent

Two arithmetic mistakes appear often in interview solutions:

1. Two equal values are stored differently. For example, `1/2`, `2/4`, and
   `-1/-2` all mean the same number, but their numerator and denominator pairs
   are different.
2. A calculation loses information. For example, a decimal approximation of
   a fraction may be slightly above or below its exact value.

The general fix is to choose one exact way to store each value. Converting all
equal values to that agreed form is called **normalization**. It makes exact
comparison possible without relying on rounded decimal values.

**ELI5:** Imagine that one person writes a half as `1/2` and another writes it
as `2/4`. Before filing the values, rewrite both labels as `1/2`. Then equal
amounts always go into the same drawer.

#### Shared divisors and multiples (GCD and LCM)

The **greatest common divisor (GCD)** is the largest positive integer that
divides two values. Euclid's algorithm finds it by repeatedly replacing
`(a, b)` with `(b, a % b)` until `b` is zero. For `gcd(18, 12)`, the pairs are
`(18, 12)`, `(12, 6)`, and `(6, 0)`, so the answer is `6`.

The **least common multiple (LCM)** is the smallest positive integer divisible
by both values. For non-negative `a` and `b`, calculate it with:

```text
lcm(a, b) = a / gcd(a, b) * b
```

For `18` and `12`, this gives `18 / 6 * 12 = 36`. Dividing before multiplying
reduces the chance of overflow, but the final result can still be too large for
the integer type. If either value is zero, return an LCM of zero before using
the formula. For negative inputs, use their non-negative magnitudes.

#### Fractions and ratios

For a fraction stored as `(numerator, denominator)`:

1. Find the GCD of the numerator and denominator.
2. Divide both values by the GCD.
3. Keep any negative sign on the numerator, so the denominator is positive.

For example, `(6, -8)` becomes `(-3, 4)`: the GCD of `6` and `8` is `2`, and
the sign moves to the numerator. The equal fractions `2/4`, `3/6`, and
`-1/-2` all become `(1, 2)`. These pairs can safely be used as map keys because
equal fractions now produce equal keys.

A zero denominator is invalid and must not reach the division step. A zero
numerator with a non-zero denominator should become `(0, 1)`.

#### Grouping points on the same line without rounding

Pick one point as the starting point. For every other point, calculate the
horizontal and vertical movement in the usual `(x, y)` order:

```text
changeX = otherX - startX
changeY = otherY - startY
```

The pair `(changeX, changeY)` describes the movement from the starting point.
The slope would be `changeY / changeX`, but division is unnecessary. Movement
pairs `(1, 2)`, `(2, 4)`, and `(-1, -2)` all follow the same straight line.
Reduce them to one shared pair so a map can count them together.

For example, start at `(1, 1)`:

| Other point | Movement from `(1, 1)` | Reduced pair |
| --- | --- | --- |
| `(2, 3)` | `(1, 2)` | `(1, 2)` |
| `(3, 5)` | `(2, 4)` | `(1, 2)` |
| `(0, -1)` | `(-1, -2)` | `(1, 2)` |

Every row has the same reduced pair, so all four points are on one line.

Normalize each movement pair in this order:

1. Calculate both changes in an integer type large enough to hold the
   subtraction.
2. If both changes are zero, the points are identical. Count the duplicate
   separately because it does not define a direction.
3. If `changeX` is zero, use one fixed pair such as `(0, 1)` for every vertical
   line.
4. If `changeY` is zero, use one fixed pair such as `(1, 0)` for every
   horizontal line.
5. Otherwise, ignore the signs while finding the GCD, then divide both changes
   by it. If `changeX` is negative, reverse both signs so the first value is
   always positive.

Using reduced integer pairs avoids decimal rounding. More importantly, it
guarantees that every movement along the same line produces exactly the same
map key.

#### Remainders and division with negative values

In Go, `%` keeps the sign of the number on its left. Therefore `-1 % 5` is
`-1`, not `4`. When a problem requires a result from `0` through `m-1` for a
positive `m`, adjust a negative remainder once:

```go
remainder := x % m
if remainder < 0 {
	remainder += m
}
```

The addition moves a negative result into the desired range. A non-negative
result already lies in that range and needs no change.

Go integer division drops the fractional part toward zero. Thus `-7 / 3` is
`-2`. Mathematical floor division would return `-3` because it rounds down.
Read the problem's rounding rule before using integer division with negative
values.

#### Counting trailing zeroes without calculating a factorial

`n!`, read as "n factorial," means multiplying every whole number from `1` to
`n`. For example:

```text
5! = 1 * 2 * 3 * 4 * 5 = 120
```

Calculating a larger factorial creates an enormous number. If the question only
asks how many zeroes appear at the end, the full number is not needed.

Multiplying by `10` adds a zero to the end of a whole number:

```text
12 * 10 = 120
```

A `10` can be made by multiplying `2 * 5`. Therefore, each pair containing one
`2` and one `5` creates one trailing zero. A factorial has plenty of `2`s
because every second number is even. The `5`s are less common, so the number of
`5`s determines how many pairs can be made.

Consider `25!`. These are the numbers from `1` through `25` that contain a `5`
when broken into multiplication:

```text
5  = 1 * 5       one 5
10 = 2 * 5       one 5
15 = 3 * 5       one 5
20 = 4 * 5       one 5
25 = 5 * 5       two 5s
```

There are six `5`s altogether. Each can pair with a `2`, so `25!` ends in six
zeroes.

Listing the numbers works for `25!`, but not for a very large `n`. Division
counts them more quickly. Go integer division discards the remainder.

| Calculation | What it counts | Total |
| --- | --- | ---: |
| `25 / 5 = 5` | One `5` from each of `5`, `10`, `15`, `20`, and `25` | `5` |
| `25 / 25 = 1` | The second `5` inside `25` | `6` |
| `25 / 125 = 0` | Nothing else, so stop | `6` |

The same steps become this loop:

```go
zeroes := 0
remaining := n

for remaining >= 5 {
	remaining /= 5
	zeroes += remaining
}
```

The first division counts numbers containing at least one `5`. Dividing by `5`
again counts the extra `5` inside each multiple of `25`. Further divisions count
the extra `5`s inside multiples of `125`, `625`, and so on. The loop stops when
there are no more to count.

The loop uses `O(1)` auxiliary space. It takes `O(log n)` time because every
iteration makes `remaining` five times smaller.

#### Finding factors and primes

When testing whether one number is prime, trial division only needs candidates
up to its square root. Any factor larger than the square root must be paired
with a smaller factor that has already been checked.

A **sieve** finds every prime number up to a chosen limit by removing numbers
known not to be prime. For example, to find primes up to `10`, begin with the
numbers `2` through `10`. Keep `2` and cross out its multiples: `4`, `6`, `8`,
and `10`. Then keep `3` and cross out its remaining multiple, `9`. The numbers
left are the primes `2`, `3`, `5`, and `7`. This method is called a sieve
because it filters out non-prime numbers.

When many questions ask which values up to the same limit are prime, a sieve is
usually faster than repeating trial division for every value.

#### Processing a number that does not fit in an integer type

Keep a huge integer as text and process one digit at a time. To find its
remainder after division by a positive integer `m`, start with remainder zero
and update:

```go
remainder = (remainder*10 + digit) % m
```

For the text `"1234"` with `m = 7`, the remainder after each digit is `1`,
`5`, `4`, and `2`. The answer is therefore `2`, and the program never needs to
store the integer `1234` as one value. The same method works even when the text
contains thousands of digits. Validate the optional sign and every digit
separately before doing the arithmetic. The working integer type must be able
to hold `10*(m-1)+9`; otherwise, use a wider type or an overflow-safe modular
calculation.

The shared rule is simple: decide what values count as equal, choose one exact
representation for them, and check that each intermediate calculation fits in
its integer type.

### Combinatorics and output size

**Combinatorics** is the study of counting choices. Before writing an algorithm
that generates choices, first estimate how many answers it may have to return.
Even a perfect algorithm cannot return one million answers without creating
and processing those one million answers.

For the three distinct values `[A, B, C]`:

- A **subset** chooses any number of values, and order does not matter. The
  eight subsets are `[]`, `[A]`, `[B]`, `[C]`, `[A,B]`, `[A,C]`, `[B,C]`, and
  `[A,B,C]`. Each of the `n` values has two choices--included or excluded--so
  `n` distinct values have `2^n` subsets.
- A **permutation** arranges all the values, so order matters. The six
  permutations are `ABC`, `ACB`, `BAC`, `BCA`, `CAB`, and `CBA`. There are `n!`
  permutations, where `n!` means `n * (n-1) * ... * 1`.
- A **size-`k` combination** chooses exactly `k` values, and order does not
  matter. Choosing two values gives `[A,B]`, `[A,C]`, and `[B,C]`; `[B,A]` is
  not another answer. The number of answers is written `C(n, k)`, read as
  "`n` choose `k`."

**ELI5:** If the question asks for every possible group, the number of groups
is part of the cost. There is no shortcut that lets a program print a group
without spending time on it.

For example, generating all subsets takes at least `O(2^n)` time because there
are `2^n` answers. If every answer is copied into its own slice, also count the
values inside those slices. Each input value appears in exactly half of the
subsets, so for `n > 0`, all subsets together contain `n * 2^(n-1)` values.
Similarly, there are `C(n, k)` size-`k` combinations and each contains `k`
values, so storing all of them takes `O(k * C(n, k))` output space. "Extra
space" may still be smaller because it normally excludes the required returned
output.

#### Counting combinations without large factorials

The mathematical formula is:

```text
C(n, k) = n! / (k! * (n-k)!)
```

For example, choosing two values from five gives:

```text
C(5, 2) = 5! / (2! * 3!) = 120 / (2 * 6) = 10
```

This formula is a poor direct implementation for large inputs. The factorials
can overflow the integer type even when the final answer would fit. For valid
inputs with `0 <= k <= n`, build the answer a step at a time instead:

```go
k = min(k, n-k)
result := 1
for i := 1; i <= k; i++ {
	result = result * (n - k + i) / i
}
```

For `C(5, 2)`, the loop does:

```text
i = 1: result = 1 * 4 / 1 = 4
i = 2: result = 4 * 5 / 2 = 10
```

The division is exact after each complete step. The assignment
`k = min(k, n-k)` uses the fact that choosing the two values to keep is the
same counting problem as choosing the three values to leave out:
`C(5, 2) = C(5, 3)`. Using the smaller side means fewer loop iterations.

Another option is **Pascal's recurrence**:

```text
C(n, k) = C(n-1, k-1) + C(n-1, k)
```

Pick one particular value. Every valid group either includes it or does not:

- If the group includes it, choose the remaining `k-1` values from the other
  `n-1` values: `C(n-1, k-1)` possibilities.
- If the group excludes it, choose all `k` values from the other `n-1` values:
  `C(n-1, k)` possibilities.

For example, `C(5, 2) = C(4, 1) + C(4, 2) = 4 + 6 = 10`. This rule is the
source of Pascal's triangle and is useful when dynamic programming needs many
combination counts, not just one. The [dynamic programming](#dynamic-programming)
card explains that method.

Both methods can still overflow. In the loop, even `result * nextFactor` can
overflow before the following division reduces it. Check multiplication and
addition against the integer type's limit. If the final answer must exceed the
ordinary integer range, Go's `big.Int` type from the `math/big` package can
store integers that grow as needed. Another implementation can cancel common
factors in the numerator and denominator before multiplying, but that is an
additional optimization rather than part of the basic counting idea.

#### Repeated input values

When values repeat, clarify what counts as a different answer. For the labeled
positions `[1a, 1b, 2]`, choosing two positions gives three choices:
`[1a,1b]`, `[1a,2]`, and `[1b,2]`. If answers contain only values, however,
the last two both become `[1,2]`, so the distinct value combinations are only
`[1,1]` and `[1,2]`.

To generate each value combination once, sort the input and build one answer a
choice at a time. When choosing a value for the same output position, skip an
equal value already tried there. This is a common
[backtracking](#backtracking) technique. Do not skip the equal value at the
next output position: after choosing the first `1`, the second `1` must remain
available so `[1,1]` can be formed.

### Integer Digit Processing

#### Mental model

An integer's decimal digits can be removed from right to left without first
turning the number into text:

- `value % 10` returns the final digit;
- `value / 10` removes the final digit because integer division drops the
  remainder.

For example, `123 % 10` is `3`, and `123 / 10` is `12`.

#### Worked example: reversing digits

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

#### Invariant (fact that remains true)

After `k` passes, `value` is the original number with its final `k` digits
removed. `reversed` contains those removed digits in the opposite order. When
`value` reaches zero, every digit has moved.

#### Preventing integer overflow

Multiplying by `10` may exceed the largest value the integer type can hold.
Before calculating `reversed*10 + digit`, check:

```go
reversed > (limit-digit)/10
```

This comparison comes directly from the safe condition
`reversed*10 + digit <= limit`. Subtract `digit` from both sides and divide by
positive `10`; the next operation is safe only when
`reversed <= (limit-digit)/10`. The code tests the opposite condition with
`>`; when it is true, the reference implementation returns `false` instead of
performing the unsafe calculation.

Negative inputs need an explicit rule because the minus sign is not a digit.
`ReverseNonNegative` rejects them and accepts zero.

#### Work and memory

For a number with `d` digits, the loop performs `d` passes, so it uses `O(d)`
work. It stores only the current digit and two integers, so it uses `O(1)` extra
memory.

#### Related problems

- [LeetCode 9: Palindrome Number](../problems/solved_easy_leetcode_0009_palindrome_number.go)
  compares a full reversal with the original number and also shows how reversing
  only half of the digits can prove the same result.

## Data structures

### Arrays, slices, and strings

An array stores a fixed number of values. A slice gives access to some or all
of an array.

**Plain-language picture:** An array is a row of numbered lockers. A slice is a
view through which some consecutive lockers are visible. Two slices can look
like separate lists while still viewing the same lockers underneath.

A slice's **length** is the number of values currently visible. Its
**capacity** is the number of positions available from its starting point
before Go must allocate another array to make it grow. Appending may therefore
reuse the existing array or create a new one.

**Worked example:** If `base := []int{10, 20, 30, 40}` and
`middle := base[1:3]`, then `middle` sees `[20, 30]`. Setting `middle[0] = 99`
also changes `base` to `[10, 99, 30, 40]` because both slices refer to the same
underlying array. An append that exceeds `middle`'s capacity allocates a new
underlying array; changes through the returned slice then no longer affect
`base`.

A small slice can keep its much larger underlying array in memory. Copy the
small result when it will live for a long time and the larger array is no
longer needed.

A Go string is a read-only sequence of bytes. Reading `text[i]` returns one
byte, which is not always a complete Unicode character. The
[bytes, runes, and visible characters](#bytes-runes-and-visible-characters)
section explains how to choose the intended text unit.

Arrays also provide an addressable position for every value. When values are
guaranteed to lie in a small range, the position can represent the value and
the stored number can represent its count. This is called a **frequency
slice**. For `[2, 1, 2]`, a three-position frequency slice becomes `[0, 1, 2]`:
value `0` appears zero times, value `1` once, and value `2` twice. This is often
simpler than a map, but it is safe only when every value is a valid slice
position. Negative or unexpectedly large values would cause an invalid access.

#### Moving values in place

An algorithm works **in place** when it rearranges the input's existing
storage instead of creating another collection of similar size. Common forms
use one position to read values and another to write retained values, divide
the input into regions, or reverse a range.

For example, rotate `[1, 2, 3, 4, 5]` two positions to the right:

```text
reverse everything:       [5, 4, 3, 2, 1]
reverse the first 2:      [4, 5, 3, 2, 1]
reverse the remaining 3:  [4, 5, 1, 2, 3]
```

When `k` may exceed the length, `k %= n` replaces it with an equivalent
smaller rotation: rotating five items by seven positions is the same as
rotating by two. Handle an empty slice before taking `k % n`, because division
by zero is invalid.

A **three-way partition** chooses a comparison value called a **pivot** and
maintains three regions: values smaller than the pivot, values equal to it, and
values greater than it. A fourth region contains values not yet classified.
After every swap, one classified region must grow and the unclassified region
must shrink; otherwise an item may be skipped or processed forever.

### Linked Lists

#### Mental model

A singly linked list is a chain of nodes. Each node stores one value and a link
to the next node. The first node is the **head**; a `nil` next link marks the
end.

Unlike an array, a linked list cannot jump directly to position `i`. Reaching a
position requires following links from the head. In exchange, code can reconnect
known nodes without shifting every later value.

#### Core operations

- `FromSlice` builds a list in the slice's order.
- `ToSlice` follows links from the head and collects the values.
- `Reverse` changes every next link so the chain points in the opposite
  direction.

#### Worked example: reversing a list

Reverse `1 -> 2 -> 3 -> nil` while keeping three positions:

1. `previous` is `nil` and `current` is node `1`. Save node `2`, point `1` to
   `nil`, then advance.
2. `previous` is node `1` and `current` is node `2`. Save node `3`, point `2` to
   `1`, then advance.
3. `previous` is node `2` and `current` is node `3`. Save `nil`, point `3` to
   `2`, then advance.
4. `current` is `nil`, so `previous` is the new head of
   `3 -> 2 -> 1 -> nil`.

Saving the next node before changing the link is essential; otherwise the
unvisited remainder of the list is lost.

#### Invariants (facts that remain true)

- Following `Next` from the head visits nodes in list order and eventually
  reaches `nil` when the list contains no loop.
- While building a list with a `tail` position, `tail` always names the final
  node. A new node can therefore be attached without walking from the head.
- During reversal, `previous` is the already reversed portion and `current` is
  the first node not yet reversed.

#### Common techniques

- A **dummy head** is a temporary node placed before the real first node. It
  avoids separate code for inserting or replacing the head.
- Two positions can advance independently when input lists have different
  lengths.
- [Fast and slow positions](#fast-and-slow-pointers) can detect a loop or
  locate the middle by moving one position one link per step and the other two.
- For numbers stored with the lowest digit first, add matching nodes and carry
  `sum / 10` to the next position; the current output digit is `sum % 10`.
  For example, adding digit nodes `8` and `7` produces output digit `5` and
  carry `1` because `15 % 10 = 5` and `15 / 10 = 1`.

#### Work and memory

| Operation | Work | Extra memory |
| --- | --- | --- |
| Visit or search the entire list | `O(n)` | `O(1)` |
| Insert after a node already in hand | `O(1)` | `O(1)` |
| Reverse with a loop | `O(n)` | `O(1)` |
| Copy all values to or from a slice | `O(n)` | `O(n)` |

These memory amounts do not count nodes required for a returned result. A
version that uses nested function calls can also keep as many as `n` unfinished
calls in memory.

#### Common mistakes

- Save `node.Next` before replacing it during reversal.
- Make loop conditions handle an empty list and a one-node list.
- Keep the original head when it is still needed; advancing the only reference
  can make earlier nodes unreachable.
- Following links one by one never ends if the list links back to an earlier
  node.
- Prefer a loop when a list may be too long for deeply nested function calls.

#### Related problems

- [LeetCode 2: Add Two Numbers](../problems/solved_easy_leetcode_0002_add_two_numbers.go)
- [LeetCode 21: Merge Two Sorted Lists](../problems/solved_easy_leetcode_0021_merge_two_sorted_lists.go)

### Queues and deques

A **queue** processes the oldest waiting item first: add at the back and remove
from the front. This first-in, first-out order is what makes breadth-first
search process one distance layer before the next. A **deque** (double-ended
queue) supports adding and removing at both ends. It is useful when both the
oldest candidate and the newest candidate may become irrelevant.

**ELI5:** A queue is a line at a ticket counter: the person who arrived first
is served first. A deque is a line with a door at both ends, so people can enter
or leave from either side.

**Worked example:** Starting with queue `[A, B]`, adding `C` produces
`[A, B, C]`. Removing one item returns `A`, leaving `[B, C]`. This exact order
is why a maze search visits every cell one step away before any cell two steps
away.

With a Go slice, keep a `head` index instead of repeatedly shifting every item:

```go
queue := []int{start}
for head := 0; head < len(queue); head++ {
	current := queue[head]
	// Append newly discovered items to queue.
}
```

Across the whole traversal, this gives constant average work per addition and
removal even though an occasional append may need to copy the slice. This way
of averaging the cost over a sequence is called **amortized analysis**. A
long-lived queue should occasionally compact its unread portion or use a
**circular buffer**, which reuses positions by wrapping from the end of a fixed
slice back to its beginning. Retaining `queue[head:]` can otherwise keep the
entire old backing array alive. Clear removed positions when they contain
pointers so Go's automatic memory manager, called the **garbage collector**,
can reclaim objects no longer in use.

A queue built from two stacks illustrates the same idea. New items go onto an
input stack. Removal uses an output stack; only when that stack is empty are all
input items moved to it. One removal can cost `O(n)`, but each item moves
between stacks at most once, so a sequence of operations costs `O(1)` amortized
time per operation.

### Trees and binary-search trees

A tree is a collection of nodes joined by links with no loop and exactly one
route between any two nodes. A **rooted** tree chooses one node as its starting
point and describes other nodes as its descendants. In a rooted binary tree,
every node has at most a left and right child. Be clear about whether `nil`
children, parent links, duplicate values, or node identity matter; equal values
do not make two nodes the same node.

**ELI5:** Picture a family tree turned upside down. Start at one root, follow
branches to children, and never find a branch that loops back to where it
started. A binary tree merely limits each node to at most two children.

A **binary-search tree (BST)** adds an ordering rule. Under the usual rule for
distinct values, everything in a node's left subtree is smaller than the node,
and everything in its right subtree is larger. A **subtree** is simply one node
together with every node below it.

**Worked search example:** In a BST with root `8`, left child `3`, and right
child `10`, searching for `10` compares with `8`, goes right because `10 > 8`,
and finds it. Searching for `4` goes left at `8`, then right at `3`. If that
link is `nil`, the value is absent. Each comparison rules out one whole subtree
only because the ordering rule applies to every descendant, not just the
direct children.

The common depth-first traversal orders differ only in when the current node is
handled:

| Order | Visit current node | Common use |
| --- | --- | --- |
| Preorder | Before both children | Copy the tree's shape or pass information from a parent to its children. |
| Inorder | Between left and right | Read a BST in sorted order. |
| Postorder | After both children | Calculate a parent result from results already returned by its children. |
| Level order | By increasing depth | Visit every node one edge away before nodes two edges away. |

For this tree:

```text
    2
   / \
  1   3
```

preorder produces `[2, 1, 3]`, inorder produces `[1, 2, 3]`, postorder produces
`[1, 3, 2]`, and level order produces `[2, 1, 3]`. The names describe when the
current node is visited relative to its children.

For recursive tree problems, define exactly what one call returns. A height
function, for example, can return `1 + max(leftHeight, rightHeight)` to its
parent. The return value follows only one side because a path continuing upward
cannot split into both children. A separate best answer stored for the whole
tree may combine both sides at the current node. State these two meanings
separately instead of using one variable for both.

Validating only a BST node against its direct children is insufficient. A
smaller value hidden anywhere in a right subtree still breaks the rule. Pass
down the allowed lower and upper bounds from every ancestor, or verify that an
inorder traversal is strictly increasing. If duplicate values are allowed, the
problem must state whether equality belongs on the left, the right, or neither.

Search, insertion, and deletion take `O(h)` time for tree height `h`. A balanced
BST keeps its left and right sides from becoming extremely uneven, giving
`h = O(log n)`. An unrestricted BST can become a one-sided chain with
`h = O(n)`. Go's standard library does not provide an **ordered map**, which is
a key-value collection that also keeps keys in sorted order. Interview
solutions therefore usually use sorting, a heap, binary search over a slice,
or a problem-specific tree instead of implementing a balanced BST from
scratch.

When reconstructing a tree from traversals, one traversal usually identifies
the next root and another identifies the boundary between its left and right
subtrees. For example, preorder `[2, 1, 3]` says the root is `2`; finding `2`
in inorder `[1, 2, 3]` shows that `[1]` belongs to the left subtree and `[3]`
to the right. Store each value's inorder position to avoid scanning that
sequence again for every node. Reconstruction from values alone requires a
uniqueness rule because repeated values can describe more than one tree.

A rooted tree may need to be traversed in directions the node type does not
store. For a question such as all nodes at distance `k` from a target, first
walk the tree once to map every child to its parent. Then run
[breadth-first search (BFS)](#breadth-first-search) from the target using left
child, right child, and parent as two-way neighbors. BFS visits nodes one link
away, then two links away, and so on. Mark node identities reached when adding
them to the queue so the search cannot bounce between a parent and child. The
conversion and search both take `O(n)` time and space.

### Graphs and grids

A graph stores vertices and connections.

**ELI5:** A graph is a map of things and the allowed moves between them. The
things are called vertices or nodes; the connections are called edges. Cities
and roads, people and friendships, courses and prerequisites, and grid cells
and legal moves are all the same basic shape.

In complexity descriptions below, `V` means the number of vertices and `E`
means the number of edges.

**Worked example:** An **adjacency list** stores each vertex together with its
neighbors. For roads `A-B`, `A-C`, and `C-D`, it can be
`A: [B, C]`, `B: [A]`, `C: [A, D]`, `D: [C]`. A search starting at `A` marks
`A`, schedules `B` and `C`, then reaches `D` through `C`. Marking a town when
it is scheduled prevents the reverse road from scheduling it again.

Before choosing an algorithm, decide these properties:

- **Directed or undirected:** an undirected edge must normally appear in both
  adjacency lists; a directed edge has one direction and contributes to the
  destination's **indegree**, meaning its number of incoming edges.
- **Weighted or unweighted:** [breadth-first search (BFS)](#breadth-first-search)
  minimizes the number of edges used, not an arbitrary total of differing edge
  costs.
- **Cyclic or acyclic:** a general graph needs reached-state tracking; a rooted
  tree can usually avoid revisiting only by not returning to its parent.
- **Connected or disconnected:** one traversal reaches only one component, so
  an outer loop is needed when every component matters. A **component** is one
  group of vertices that can reach one another.

The adjacency list uses `O(V + E)` space and is the usual representation for a
**sparse graph**, where relatively few of the possible vertex pairs have an
edge. An **adjacency matrix** instead uses a `V`-by-`V` table: cell
`matrix[a][b]` records whether the edge `a -> b` exists. It uses `O(V^2)` space
but checks one specific edge in constant time. An **edge list** stores entries
such as `(from, to, cost)` and is convenient when an algorithm repeatedly scans
every edge, as [Bellman-Ford](#bellman-ford-shortest-paths) does.

Mark a vertex reached when it is scheduled, not later when it is removed from a
stack or queue. Otherwise several neighbors may schedule the same vertex and
duplicate work. When cloning a graph, map each original node identity to its
clone; mapping only by value fails when different nodes contain equal values.

A matrix often represents an **implicit graph**: each cell is a vertex, and its
neighbors are calculated from direction offsets such as up, down, left, and
right. Check row and column bounds independently. Decide whether the input may
be mutated to record visited cells; if not, allocate separate state. For
[backtracking](#backtracking) on a word grid, mark a cell only for the current
path and restore it afterward. For ordinary reachability, the visited mark is
permanent.

Sometimes the input does not call itself a graph at all. Create one vertex per
object and an edge whenever the stated relationship permits a transition, such
as one device reaching another or one symbol preceding another. State whether
building those edges costs `O(n^2)` pair checks, whether direction matters, and
whether duplicate edges would corrupt indegrees or other counts.

### Matrix traversal and in-place state

**Plain-language picture:** A matrix is a grid addressed by two positions:
`matrix[row][column]`. Moving vertically changes the row; moving horizontally
changes the column. Check both positions before reading a neighbor.

Matrix problems often test boundary discipline more than a new data structure.
Treat `rows` and `cols` separately; a matrix need not be square, and Go permits
**jagged slices**, whose rows have different lengths, unless the problem
promises a rectangle.

**Worked boundary example:** A `2`-by-`3` matrix has rows `0..1` and columns
`0..2`. From cell `(0, 2)`, moving down to `(1, 2)` is valid, but moving right
to `(0, 3)` is not. Checking both coordinates against `len(matrix)` would be
wrong because the number of rows and columns differ.

- **Boundary traversal:** For spiral order, keep inclusive `top`, `bottom`,
  `left`, and `right` boundaries. After traversing one side, move its boundary
  inward and recheck that the opposite boundaries have not crossed before
  traversing the bottom or left side. For the two rows `[1, 2, 3]` and
  `[4, 5, 6]`, the order is `1, 2, 3, 6, 5, 4`; after reading the top and right
  sides, only the bottom side remains.
- **Layer transformations:** A square rotation can be expressed as a transpose
  followed by reversing each row, or as four-way swaps around each layer.
  A transpose swaps rows with columns, so position `(row, col)` moves to
  `(col, row)`. For `[[1, 2], [3, 4]]`, transposing gives
  `[[1, 3], [2, 4]]`; reversing each row gives the clockwise rotation
  `[[3, 1], [4, 2]]`.
- **First-row or first-column markers:** When rows and columns must be marked
  without extra arrays, reuse the first row and column as marker storage and
  keep separate flags for whether those marker areas originally required a
  change. If a zero at `(1, 2)` means row `1` and column `2` must become zero,
  record those two facts first and apply them only after every original cell
  has been inspected.
- **Current and next state together:** For simultaneous updates, such as a grid
  in which every cell changes once per round, encode both the original and
  future state in each cell or record changes separately. Neighbor decisions
  must read only the original state; immediately overwriting it makes later
  cells observe a different round.

The shared invariant is that every read during a logical phase sees the same
version of the matrix. If an encoding is used, document how to recover both
versions and ensure input values leave enough unused states for the encoding.

For example, suppose `0` means a dead cell and `1` means a live cell. Writing a
future `0` or `1` immediately would erase the old state needed by neighboring
cells. One safe encoding could temporarily use `2` for "was dead, becomes
live" and `3` for "was live, becomes dead." Neighbor checks read whether the
cell was originally live; a final pass converts the temporary states to `0` or
`1`.

### Composing data structures

Design questions often require several operations to be fast at the same time.
No single structure may provide all of them, so combine structures whose
invariants repair one another's weakness:

**ELI5:** Think of an address book paired with a waiting line. The address book
finds a person quickly but does not say who arrived first; the line preserves
arrival order but is slow to search. Keeping them synchronized provides both
abilities.

In the table, **expected `O(1)`** means constant time under the usual behavior
of hashing or random choices; it does not promise that every individual
operation always takes exactly that long.

| Required operations | Composition and invariant |
| --- | --- |
| Stack minimum in `O(1)` | Store each value with the minimum at or below it, or keep a parallel stack of minima. |
| Insert, delete, random item in expected `O(1)` | A slice gives random indexing; a map gives each value's slice index. Delete by swapping in the last value and repairing its map entry. |
| Least-recently-used (LRU) cache `get` and `put` in `O(1)` | A map finds an entry; a doubly linked list, whose nodes link both forward and backward, orders entries from most to least recently used. Every access moves exactly one node. |
| Streaming median | A [max-heap and min-heap](#heap-or-priority-queue) keep the lower and upper halves. Their sizes differ by at most one, and every lower value is no greater than every upper value. |
| Timestamped lookup | Store a sorted history per key and [binary-search](#binary-search) the last timestamp not greater than the query. |

**Worked deletion example:** Suppose a set that also supports random retrieval
stores values `[10, 20, 30]` and the map stores `10 -> 0`, `20 -> 1`, and
`30 -> 2`. To delete `20` without shifting later values, copy the final value
`30` into position `1`, update its map entry to `30 -> 1`, shorten the slice to
`[10, 30]`, and remove `20` from the map. The map and slice agree again after
every step.

Write the cross-structure invariant before implementing operations. For a
randomized set, every slice position and map index must agree after a swap. For
an LRU cache, every map entry must name exactly one linked-list node, and the
list order must match recency. Most bugs in these designs come from updating
only one representation or mishandling the first or last node. Permanent dummy
nodes called **sentinels** can stand just before and after the real entries, so
insertion and removal do not need separate empty-list or end-of-list cases.

### Streaming and online processing

An **online** algorithm must update its answer as values arrive, without
assuming the complete input is available. Define what the stored state means
after every **prefix**, meaning all items received so far, and whether old data
may be discarded.

**Worked example:** Report the first non-repeating letter after each arrival in
the stream `a, b, a`. After the first `a`, its count is one and the answer is
`a`. After `b`, the candidate queue is `[a, b]`, so the answer remains `a`.
After the second `a`, its count becomes two; remove it from the queue front and
report `b`. The algorithm never needs to scan the entire received stream again.

- To report the first non-repeating item after each arrival, keep a frequency
  map and a queue of candidates. Increment the frequency, enqueue a first
  occurrence, and remove repeated items from the queue front. The front is the
  earliest item whose current frequency is one.
- For a fixed recent window, pair counts or another summary with a queue or
  circular buffer, which wraps around and reuses fixed storage, so the item
  leaving the window can undo its contribution.
- For a running median, maintain the two-[heap](#heap-or-priority-queue)
  invariant described above. For streaming top `k`, keep only `k` candidates
  in a heap.
- For sorted streams, keep only the next unread item from each stream in a
  heap; consuming one candidate reveals its successor.

Ask whether arrivals can be late or reordered, whether memory must be bounded,
and whether an answer is required after every item. A **batch** solution waits
for all input before starting, so one that sorts the entire input is not online
even if it produces the same final answer. Amortized analysis is often the
right measure because one arrival may remove several candidates that are no
longer valid while every candidate is still added and removed only once.

### Choosing a structure

- Use a [hash map or set](#hash-map-or-set) for exact-key lookup.
- Use a [stack](#stack) for most-recent-first processing and a
  [queue](#queues-and-deques) for oldest-first processing.
- Use a [deque](#queues-and-deques) when candidates can expire from the front
  while a better and newer candidate makes items at the back permanently
  unnecessary.
- Use a [binary-search tree](#trees-and-binary-search-trees) when keys must stay
  ordered under updates and the environment already supplies a balanced
  implementation.
- Use a [heap](#heap-or-priority-queue) when only the next smallest, largest,
  or highest-priority item must be exposed.
- Use a [trie](#trie) for repeated prefix lookup.
- Use [disjoint-set union](#disjoint-set-union) for connectivity under
  incremental unions.
- Use linked nodes when reconnecting known positions matters more than random
  access.

## Problem-Solving Patterns

Pattern cards explain how to recognize a useful method and why it does not miss
an answer. Each card names the fact that must remain true while the algorithm
runs. Computer science calls that fact an **invariant**.

### Problem-solving method selection guide

Treat the signs listed here as clues, not keyword matches. An array does not
automatically require two pointers, and the word "minimum" does not
automatically require a greedy method. A method fits only when you can explain
why each step cannot discard a needed answer.

Use this guide as a sequence of progressively stronger hints. Start with the
decision tree. Open the linked pattern card only when the tree leaves you
unsure how to begin, then try the card's first steps before reading a solution
explanation. This file owns the complete pattern guide; no separate concept
page is required.

#### Quick decision tree

Before choosing a branch, answer three basic questions:

1. **What must you return?** A value, position, path, ordering, one valid
   answer, or every valid answer?
2. **May you rearrange the input?** Check whether items must stay consecutive
   or retain their original order or indices.
3. **How large can the input be?** Eliminate methods that would use too much
   time or memory at the maximum size.

Now choose the branch that best describes the hardest part of the problem. The
last choice in a branch suggests a method. Read its linked card and confirm why
its steps are safe before using it. More than one branch may apply. For example,
Dijkstra's algorithm usually uses a heap.

- **Are you working with connections between items?**
  * Need to arrange tasks so every prerequisite comes first? Use
    **[topological sort](#topological-sort)**.
    **Example:** If course `C` requires both `A` and `B`, return an order with
    `A` and `B` before `C`.
  * Need the shortest or cheapest route from one starting point?
    **Example:** Find the cheapest route from town `A` to every other town on a
    road map.
    + Every move costs the same? Use
      **[breadth-first search](#breadth-first-search)**.
      **Example:** In a maze, each step to an adjacent open cell costs one move.
    + Costs vary, but none are negative? Use
      **[Dijkstra's algorithm](#dijkstra-shortest-paths)**.
      **Example:** Roads have travel times such as `2`, `5`, and `9` minutes.
    + Some costs are negative? Use
      **[Bellman-Ford](#bellman-ford-shortest-paths)**.
      **Example:** A route can include a `-3` rebate, so a later path may make
      an earlier distance cheaper.
  * Are connections added one at a time while you repeatedly ask whether two
    items are connected? Use **[disjoint-set union](#disjoint-set-union)**.
    **Example:** Add roads one by one and, after each road, ask whether towns
    `1` and `5` are now connected.
  * Need to connect every item with the least total connection cost? Use a
    **[minimum spanning tree](#minimum-spanning-tree)**.
    **Example:** Connect every office while minimizing total cable length.
  * Need to completely explore one path, region, or subtree before the next?
    Use **[depth-first search](#depth-first-search)**.
    **Example:** Calculate a directory's total size by finishing every child
    directory before returning its size to the parent.

- **Are you building an answer by choosing among alternatives?**
  * Can you prove that one choice now will not worsen the best final result?
    Use **[greedy](#greedy)**.
    **Example:** To attend the most meetings, always choose the compatible
    meeting that finishes earliest.
  * Do many choice sequences reach the same remaining situation, and you need
    only its best value, number of solutions, or yes/no result? Use
    **[dynamic programming](#dynamic-programming)**.
    **Example:** Different coin choices can leave the same remaining amount;
    calculate the minimum coins for that amount once.
  * Must you try a choice, explore it, undo it, and try another? Use
    **[backtracking](#backtracking)**.
    **Example:** Place a queen in one column, remove it if a later row cannot be
    completed, and try the next column.

- **Do you need fast repeated lookups?**
  * Looking up an exact value, frequency, matching partner, or group? Use a
    **[hash map or set](#hash-map-or-set)**.
    **Example:** While looking for a pair totaling `9`, ask whether `9 - x` has
    already appeared for each number `x`.
  * Looking up words by a shared beginning? Use a **[trie](#trie)**.
    **Example:** For autocomplete after `ca`, keep `car`, `cart`, and `cat`
    available without scanning every word.
  * Need several different operations to be fast together? Consider
    **[composing data structures](#composing-data-structures)**.
    **Example:** Combine a map and doubly linked list for an LRU cache.

- **Does the problem revolve around ranges?**
  * Need to merge overlaps or find gaps? Use **[intervals](#intervals)**.
    **Example:** Merge reservations `[1, 4]` and `[3, 6]` into `[1, 6]`.
  * Need to know how many ranges are active as positions or times change? Use a
    **[sweep line](#sweep-line)**.
    **Example:** Process booking start and end times to find the maximum number
    of rooms occupied together.
  * Need to answer many sum or count queries over ranges? Use
    **[prefix sums](#prefix-sums)**.
    **Example:** Precompute totals once, then answer many questions such as the
    sum from index `20` through index `80`.
  * Need to apply many range additions and calculate the final values only
    once? Use a **[difference array](#difference-arrays)**.
    **Example:** Apply hundreds of passenger-count changes across stop ranges,
    then reconstruct the count at every stop once.
  * Need interleaved updates and range queries? Use a
    **[Fenwick or segment tree](#fenwick-and-segment-trees)**.
    **Example:** Update one score and repeatedly query the sum or maximum over
    a range without rebuilding every prefix.

- **Does the problem revolve around positions in a sequence?**
  * Would rearranging values make duplicates, intervals, or compatible choices
    adjacent? Use **[sorting and selection](#sorting-and-selection)**.
    **Example:** Sort intervals by start so every overlap appears beside the
    interval it may extend.
  * Comparing several sequences to find the beginning shared by all of them?
    Use a **[common-prefix scan](#common-prefix-scan)**.
    **Example:** Compare `"flower"`, `"flow"`, and `"flight"` one column at a
    time until the first mismatch.
  * Does an item's meaning depend on an adjacent item? Use a
    **[directional scan](#directional-scan)** that processes the needed neighbor
    first.
    **Example:** Scan a Roman numeral from right to left so the symbol to the
    right is already known.
  * Does the statement prescribe a sequence of state changes or formatting
    rules? Use **[simulation](#simulation)**.
    **Example:** Traverse matrix boundaries one side at a time, shrinking the
    remaining rectangle after each side.
  * Need the best consecutive section, and can update the tracked information
    when either end moves? Use a **[sliding window](#sliding-window)**.
    **Example:** Find the longest consecutive part of `"abca"` without a
    repeated character by growing and shrinking one window.
  * Searching sorted values, or the first point where a yes/no test changes?
    Use **[binary search](#binary-search)**.
    **Example:** In `[1, 3, 3, 5]`, find the first position whose value is at
    least `3`.
  * Can two positions move through the sequence while each move safely rules
    out possibilities? Use **[two pointers](#two-pointers)**.
    **Example:** In a sorted array, move the left or right position according
    to whether their sum is too small or too large.
  * Working with a linked structure and need a cycle, midpoint, or fixed gap?
    Use **[fast and slow pointers](#fast-and-slow-pointers)**.
    **Example:** Move one linked-list pointer one step and another two steps;
    they meet if the list contains a cycle.
  * Need the next or nearest greater or smaller value? Use a
    **[monotonic stack](#monotonic-stack)**.
    **Example:** For each daily temperature, find the next later day with a
    higher temperature.
  * Need the best sum of one consecutive section? Use
    **[Kadane's algorithm](#kadanes-algorithm-best-subarray-scan)**.
    **Example:** At each value, decide whether the best section ending here
    should extend the previous section or begin here.
  * Searching for the same exact pattern inside a larger string? Use
    **[string matching](#string-matching)** when restarting comparisons would
    repeat substantial work.
    **Example:** Reuse the known border of a partially matched pattern after a
    mismatch instead of restarting from its first character.
  * Does each character's meaning depend on what has already been parsed? Use
    **[parsing and state machines](#parsing-and-state-machines)**.
    **Example:** A sign is allowed before digits but not after the numeric part
    has begun.

- **Does the order of waiting items determine what happens next?**
  * Must the most recently opened or unfinished item be handled first? Use a
    **[stack](#stack)**.
    **Example:** When checking `([])`, the closing bracket must match the most
    recent unmatched opening bracket.
  * Must you repeatedly remove the smallest or largest item while the stored
    items keep changing? Use a
    **[heap or priority queue](#heap-or-priority-queue)**.
    **Example:** As jobs arrive, repeatedly run the waiting job with the highest
    priority.
  * Need the minimum or maximum of every moving window? Use a
    **[monotonic deque](#monotonic-deque)**.
    **Example:** Remove expired positions from the front and values dominated
    by a new arrival from the back.

- **Does the problem naturally split into independent smaller parts?**
  * If each part can be solved separately and their results combined without
    representing the same state repeatedly, use
    **[divide and conquer](#divide-and-conquer)**.
    **Example:** Sort each half, then merge the two sorted results.

If nothing fits, return to the simplest correct solution and ask what work it
repeats. The problem may need a method outside this guide; do not force a
match.

#### Derive a direction before naming a pattern

Do not begin by guessing an algorithm name. First work through a tiny example
by hand and describe the simplest correct method, even if it tries every
possibility. Then ask what that method repeatedly recalculates or revisits.

For example, "for every starting position, I add the same later values again"
suggests storing totals that can be updated or reused. That observation may
lead to a sliding window or prefix sums. It is more useful than matching one
word in the prompt to a memorized technique.

Use this hint ladder during practice:

1. **First 5 minutes:** write or explain the simplest correct method, even if it
   tries every possibility.
2. **At 5 minutes:** use the decision tree to identify at most two possible
   methods.
3. **At 10 to 15 minutes:** read the relevant pattern card and check
   whether its reason for discarding work applies to the problem.
4. **At 20 minutes:** trace the method on one small input before writing code.
5. **At 25 to 30 minutes:** read a full explanation if still blocked, close it,
   and implement the idea without copying its code.

#### Pattern cards

Each card is the complete conceptual reference for its method: what to picture,
when it fits, how to begin, why it is safe, and when to reject it. The cards are
deliberately method-level rather than problem-specific, so they help produce a
first attempt without revealing a judge solution.

Jump directly to a card:

- Sequence and range methods: [common-prefix scan](#common-prefix-scan),
  [directional scan](#directional-scan), [simulation](#simulation),
  [two pointers](#two-pointers),
  [sorting and selection](#sorting-and-selection),
  [binary search](#binary-search), [string matching](#string-matching),
  [parsing and state machines](#parsing-and-state-machines),
  [sliding window](#sliding-window),
  [prefix sums](#prefix-sums), [difference arrays](#difference-arrays),
  [Fenwick and segment trees](#fenwick-and-segment-trees),
  [fast and slow pointers](#fast-and-slow-pointers),
  [Kadane's algorithm](#kadanes-algorithm-best-subarray-scan),
  [monotonic stack](#monotonic-stack), [monotonic deque](#monotonic-deque),
  [intervals](#intervals), and [sweep line](#sweep-line).
- Lookup and processing-order methods: [hash map or set](#hash-map-or-set),
  [trie](#trie), [stack](#stack), and
  [heap or priority queue](#heap-or-priority-queue).
- Choice methods: [dynamic programming](#dynamic-programming),
  [greedy](#greedy), [backtracking](#backtracking), and
  [divide and conquer](#divide-and-conquer).
- Relationship methods: [depth-first search](#depth-first-search),
  [breadth-first search](#breadth-first-search),
  [topological sort](#topological-sort),
  [Dijkstra](#dijkstra-shortest-paths),
  [Bellman-Ford](#bellman-ford-shortest-paths),
  [disjoint-set union](#disjoint-set-union), and
  [minimum spanning tree](#minimum-spanning-tree).

##### Common-prefix scan

**Mental model:** Compare several strings one character position at a time,
beginning at the left. The answer ends as soon as one string ends or one
character differs, because a shared beginning cannot skip that position.

**Recognize it:** The answer must start at position `0` in every sequence. If
the answer may begin later or skip positions, this is not a common-prefix scan.

**Start here:** Use the first string as the proposed prefix. For each position,
check that every other string is long enough and contains the same character.
Return the part before the first failed check.

**Worked example:** For `"flower"`, `"flow"`, and `"flight"`, position `0` is
`f` everywhere and position `1` is `l` everywhere. Position `2` contains `o`,
`o`, and `i`, so the longest common prefix is `"fl"`.

**Why it stays correct:** Before checking position `c`, every string shares the
prefix `strs[0][:c]`. A mismatch or ended string proves that no longer shared
prefix is possible; otherwise the proven prefix grows by one character.

**Work and memory:** The scan uses `O(S)` time in the worst case, where `S` is
the total input size, and `O(1)` extra memory when it returns a slice of the
first string.

**Common traps:** Check a string's length before reading its current position.
Define the result for empty input. Also decide whether a position means a byte,
a Unicode code point, or a user-visible character before applying this method
to general text.

**Reject it when:** The shared result may start after position `0`, skip
positions, or compare reordered items.

**Representative problem:** [Longest Common Prefix](../problems/solved_easy_leetcode_0014_longest_common_prefix.go).

##### Directional scan

**Mental model:** When an item's meaning depends on a neighboring item, scan in
the direction that makes that neighbor already known instead of repeatedly
looking ahead.

**Recognize it:** A local rule refers consistently to the item immediately to
the right or left. A right-neighbor rule suggests scanning right to left; a
left-neighbor rule usually suggests scanning left to right.

**Start here:** State which neighbor determines the current item's treatment,
choose the direction that processes that neighbor first, and remember exactly
the neighboring information needed for the next step.

**Worked example:** Roman numeral `IV` represents `4`. Scanning numeric values
`[1, 5]` from right to left first adds `5`, then sees that `1` is smaller than
the remembered right value and subtracts it, producing `4`.

**Why it stays correct:** Before processing a position, the already-scanned
side has been handled correctly and the retained value represents the adjacent
item needed by the rule. Applying the rule once extends the correct result by
one position.

**Work and memory:** A single scan takes `O(n)` time and `O(1)` extra memory
when only one neighboring value is retained.

**Common traps:** The neighboring-value rule must come from the problem; an
adjacent comparison is not automatically correct. Handle the first processed
item explicitly when a sentinel would be invalid, and update the remembered
neighbor after every item, including one that contributes negatively.

**Reject it when:** The current result depends on information from farther away
that cannot be made available by choosing one scan direction.

**Representative problem:** [Roman to Integer](../problems/solved_easy_leetcode_0013_roman_to_integer.go).

##### Simulation

**Mental model:** Translate the problem's rules into an explicit state and a
transition that advances that state. Simulation is a legitimate algorithm when
the rules themselves are the shortest clear route to the result.

**Recognize it:** The statement describes movement, formatting, rounds,
simultaneous updates, or a process whose next state follows directly from the
current state. The challenge is usually boundary handling and preserving the
correct version of state, not discovering a hidden optimization.

**Start here:** List every state variable and its meaning, define the order of
one step, and identify a progress measure that proves termination. Separate
phases when later actions must see a snapshot of earlier state. Extract a helper
only when it represents one rule clearly, such as flushing a completed output
line or checking one matrix neighbor.

**Worked example:** A robot starts at `(0, 0)`, facing north, with commands
`move`, `right`, `move`. Store only its position and direction. The first move
changes the position to `(0, 1)`. Turning right changes the direction to east
without moving. The final move changes the position to `(1, 1)`. Each command
is one state transition, and the number of unread commands proves progress.

**Why it stays correct:** Before each step, the stored state represents exactly
the process after the already completed steps. Applying the written rule once
produces the specified next state and advances the progress measure.

**Common traps:** Updating shared state too early, processing one boundary
twice after dimensions cross, forgetting a final partial group, confusing an
index with a count, or adding a fallback that hides an impossible state. Build
tests around empty or smallest inputs, transitions between phases, exact fits,
and one item beyond an exact fit.

**Work and memory:** If the process has `e` events, simulation usually takes
`O(e)` time plus the time needed to produce its output. Extra memory is the
state being tracked, which may be `O(1)` for a few values or larger when the
process must retain a board, history, or pending events.

**Reject it when:** The required event count is too large for the constraints
and the rules allow the state to jump directly to a later time or value.

##### Hash map or set

**Mental model:** Store a fact under a key that can be calculated again. A map
returns associated information; a set records presence only.

**Recognize it:** You repeatedly ask whether a value appeared, how often it
appeared, which earlier value completes a pair, or which group an item belongs
to. Use a set when no value beyond presence is needed.

**Start here:** Define the exact key and the meaning of its stored value. Decide
whether the current item must query the map before or after it is inserted, and
whether repeated keys keep the first value, latest value, count, or collection.

**Worked example:** For `[2, 7, 11, 15]` and target `9`, the scan first asks for
`7`, finds nothing, and stores `2 -> 0`. At value `7`, it asks for `9 - 7 = 2`,
finds position `0`, and returns positions `0` and `1`. Querying before insertion
prevents an item from pairing with itself while still allowing `[3, 3]`.

**Why it stays correct:** Before processing position `i`, the map contains
exactly the facts retained from positions before `i`, under their defined keys.
It contains nothing from the current or later positions.

**Work and memory:** A scan over `n` items usually takes expected `O(n)` time
because map operations are usually close to constant time, and it uses up to
`O(n)` extra memory. A single map operation can occasionally take longer.

**Common traps:** A map does not preserve sorted or insertion order by itself.
In Go, reading a missing key returns the value type's **zero value**, meaning
its default such as `0`, `false`, or `nil`. Use the
`value, found := entries[key]` form when that default could also be stored.

For grouping problems, design a **canonical key**: equal logical objects must
produce the same key and unequal objects must not collide under the chosen
representation. **Anagrams**, which contain the same character counts in a
different order, can use a fixed-size count array as a key when the alphabet is
fixed and the array is comparable. Arbitrary text may need a sorted
representation or encoded counts.

An **isomorphic mapping** consistently replaces characters while preserving
which positions are equal. For example, `egg` can map to `add` with `e -> a`
and `g -> d`, but `ab` cannot map to `cc` because two different source
characters would collapse into one destination. Checking maps in both
directions enforces both requirements.

**Reject it when:** The essential operation is ordered traversal, nearest-key
search, or another relationship a hash key does not preserve.

**Representative problem:** [Two Sum](../problems/solved_easy_leetcode_0001_two_sum.go).

##### Sorting and selection

**Mental model:** Pay `O(n log n)` once to place related values next to one
another or to make later decisions move in only one direction. If only the item
of rank `k` matters, selection may avoid producing the complete order.

**Recognize it:** The input order is not meaningful, and sorted order would
make duplicates adjacent, expose interval overlaps, enable two pointers, or
let a greedy rule process candidates consistently. Also consider sorting when
the desired output itself must be ordered.

**Start here:** Decide what the comparison means and whether the original
indices or input order must be retained. For records, define tie-breaking
explicitly. This comparison rule, called a **comparator**, must be consistent:
if `a` comes before `b` and `b` comes before `c`, then `a` must come before `c`.
If equal items must keep their previous order, use a stable sort or include the
original position as a tie-breaker.

**Worked example:** To detect duplicates in `[4, 1, 4, 2]`, sort it into
`[1, 2, 4, 4]`, then compare neighbors. The two `4`s become adjacent, so one
linear scan finds them. Sorting costs more than a set for this one question,
but the same ordering may also support later range or pair queries.

**Work and memory:** Standard comparison sorts are usually budgeted as
`O(n log n)` time. Quicksort divides values around a pivot and sorts the two
sides; consistently poor pivots can make it degrade to `O(n^2)`. Extra memory
depends on the chosen sorting algorithm; copying the input to preserve it uses
`O(n)` additional memory. Go's `sort.Ints` and `sort.Slice` modify the slice
and are not stable; `sort.SliceStable` preserves equal elements' order at an
additional practical cost.

When integer keys occupy a small known range of size `k`, counting occurrences
can sort or group them in `O(n + k)` time and `O(k)` extra memory. This is useful
only when `k` is reasonably small; allocating a table spanning a huge key range
is worse than comparison sorting.

**Selection:** To find the `k`th smallest value, quickselect partitions values
around a pivot. After partitioning, the pivot is in its final rank, every value
on one side is no greater, and every value on the other side is no smaller.
Continue only on the side containing rank `k`. Randomized or careful pivots give
expected `O(n)` time, but consistently poor pivots give `O(n^2)` worst-case
time. A heap gives a simpler `O(n log k)` worst-case bound for top-`k`
streaming data and preserves the selected items without fully sorting all data.

**Why sorting helps:** The sort itself does not solve the problem. It creates a
monotonic relationship: once a scan proves a value, interval, or candidate too
small, too early, or disjoint, later items cannot reverse that fact.

**Reject it when:** Original order or indices are essential and cannot be
carried along, linear time is required, or a hash table, counting table, heap,
or direct selection answers the question without a complete order.

##### Binary search

**Mental model:** Test the middle possible answer and discard one half. The
method is safe only when the possibilities form two consecutive groups, such
as all failures followed by all successes.

**Recognize it:** Data is sorted, the problem asks for the first or last
position satisfying a condition, or an ordered numeric answer space has a
single boundary where a yes/no test changes. A numeric answer alone is not
enough.

**Start here:** Define precisely what `left` and `right` include. A range
written as `[left, right)` includes `left` but excludes `right`. For that kind
of range, test `middle = left + (right-left)/2` and keep
the half that can still contain the boundary. Decide which side retains an
equal value based on whether the first or last valid position is required.

**Worked example:** To find the first value at least `3` in `[1, 3, 3, 5]`,
start with `[0, 4)`. Testing value `3` at position `2` keeps `[0, 2)`; testing
value `3` at position `1` keeps `[0, 1)`; testing value `1` discards position
`0`. The empty range meets at position `1`, the first valid position.

**Why it stays correct:** Before every test, the first acceptable position lies
in `[left, right)`. Each update discards only positions proved too early or too
late; when the range is empty, `left` is the boundary.

**Wrapped sorted data:** In data such as `[4, 5, 6, 1, 2, 3]`, the whole range
is not sorted, but at least one side of the middle is. Identify that side and
use its outer endpoint to determine whether the target belongs there.

Duplicates can make both sides appear equally sorted and remove the evidence
needed to discard half. When equality at both boundaries is allowed, shrinking
a boundary may be necessary and the worst case can fall to `O(n)`. Likewise,
binary search for a peak or local minimum needs the exact strictness rule;
plateaus invalidate proofs that rely on one adjacent comparison selecting a
unique downhill direction.

**Work and memory:** Each test removes about half of the remaining positions,
so a loop-based search takes `O(log n)` time and `O(1)` extra memory.

**Common traps:** Keeping the tested middle can prevent progress. A valid
insertion position may equal `len(values)`. Ordinary bounds require fully
sorted input; wrapped input needs the separate sorted-side rule.

**Reject it when:** Testing a later candidate can change from yes back to no,
or the middle test cannot prove that one entire half is impossible.

**Representative problems:** [Search Insert Position](../problems/solved_easy_leetcode_0035_search_insert_position.go)
and [Search in Rotated Sorted Array](../problems/solved_medium_leetcode_0033_search_in_rotated_sorted_array.go).

##### String matching

**Mental model:** Search a text for an exact pattern while remembering how much
of the pattern is still useful after a mismatch.

**Recognize it:** The same pattern is searched inside a larger string, and a
straightforward restart would compare the same text characters many times. For
a short pattern or one search, that straightforward method may still be best.

**Start with the direct method:** Try the pattern at every possible starting
position. At each start, compare from left to right until every character
matches or one differs.

**Worked example:** Search for `abab` in `ababab`. Trying position `0` matches
`a`, `b`, `a`, and `b`, so the answer is `0`. If the pattern is `abac`, the
comparison at position `0` matches `aba` and then fails on the fourth
character. A direct search next tries position `1` from the beginning.

**Remembering useful matched text:** The Knuth-Morris-Pratt algorithm, usually
called **KMP**, avoids restarting the pattern from zero after every mismatch.
A **prefix** begins at the pattern's first character; a **suffix** ends at its
last character. A **proper prefix** is shorter than the complete string. A
piece that is both a prefix and suffix is sometimes called a **border**.

For pattern `ababac` and text beginning `ababab`, the first five characters
`ababa` match before pattern character `c` disagrees with text character `b`.
The matched text ends with `aba`, which is also the pattern's first three
characters. KMP keeps those three useful matches and continues instead of
restarting at zero. A table prepared from the pattern records how many
characters remain useful after each mismatch.

KMP's invariant is that the first `matched` pattern bytes equal the ending of
the text processed so far. On mismatch, the table shortens `matched` to the
next possible border without moving the text position backward. Each text
position advances once, so a text of length `n` and pattern of length `m` takes
`O(n + m)` time and the table uses `O(m)` memory.

**Rolling hash:** A hash is a compact numeric fingerprint. Maintain the
fingerprint for the current length-`m` text window by removing the outgoing
character's contribution and adding the incoming character's contribution.
Two different windows can occasionally have the same fingerprint; this is a
**collision**. Equal hashes therefore identify candidates, not proof. Compare
the actual bytes before reporting a match unless the problem explicitly
accepts a collision risk.

**Work and memory:** The direct method can take `O(nm)` time and `O(1)` extra
memory, such as when many repeated characters match before the final character
fails. KMP takes `O(n + m)` time and `O(m)` memory. A rolling hash can scan in
`O(n)` expected time after `O(m)` setup, but correctness still requires the
collision check described above.

**Common traps:** Define the empty-pattern result. Byte-based matching is
correct for inputs defined as bytes or a one-byte alphabet, but it can split a
Unicode character encoded with several bytes. KMP is valuable when its
worst-case guarantee matters; it is not a requirement to replace a clear
linear-library search in every interview.

**Reject it when:** One short or one-off search is clearer with a direct scan
or a standard library operation, or when the match permits edits, reordered
characters, or another relationship that exact string matching does not model.

##### Parsing and state machines

**Mental model:** The parser is always in one of a small number of modes, and
the current input symbol determines the next mode and any action. The mode
summarizes exactly what prefixes are legal so far.

Here, a **grammar** is the set of arrangements the input permits. A **token**
is one meaningful piece, such as a number, name, or operator. **Tokenization**
splits characters into those pieces; **evaluation** applies their meaning.

**Recognize it:** Input has a grammar or phased format: optional whitespace,
an optional sign, digits, perhaps a decimal or exponent, nested delimiters, or
operators with precedence. A collection of unrelated boolean flags is often an
implicit state machine that should be made explicit.

**Start here:** Separate tokenization from evaluation when that makes the rules
clearer. List the states, the symbols accepted in each state, the transition,
and which ending states are valid. Reject an invalid transition immediately
rather than letting a later fallback reinterpret malformed input.

For integer parsing, a useful sequence is `leading space -> optional sign ->
digits -> stop`. Before `value = value*10 + digit`, check the bound using
`value > (limit-digit)/10`. Apply the sign according to the required range;
the negative limit may have a larger magnitude than the positive limit.

**Worked example:** Parsing `"  -42x"` begins in `leading space`, moves to
`optional sign` at `-`, enters `digits` at `4`, builds `42`, and stops at `x`.
The sign then produces `-42`. By contrast, `"4-2"` is rejected or stopped at
`-` according to the stated grammar; the sign transition is no longer legal
after digits have begun.

For arithmetic expressions, a stack handles nested parentheses, while operator
precedence can be handled by operator/value stacks or by recursive-descent
functions: one function handles additions, another multiplications, and
another individual values or parenthesized expressions. This organization is
called **recursive descent** because the functions follow the grammar's nested
levels. Decide whether **unary** `+` and `-`, which apply to one following
value, are operators or part of numeric tokenization. Do not rely on splitting
by spaces when the grammar permits no spaces or arbitrary spacing.

**Why it stays correct:** After consuming a prefix, the state describes every
legal way that prefix may continue and the accumulated values represent exactly
the completed tokens. Each transition consumes input or terminates, so the
parser cannot silently loop on an invalid character.

**Work and memory:** Reading `n` input characters once takes `O(n)` time. A
simple number parser can use `O(1)` extra memory. Nested expressions may need a
stack whose size is proportional to the nesting depth, and retaining every
token can require `O(n)` memory.

**Reject it when:** The input has no meaningful grammar or phases, or a
well-tested standard parser already implements exactly the required language.

##### Two pointers

**Mental model:** Track two meaningful positions and move them in a coordinated
way so the scan never needs to restart. Common forms move inward from both
ends, merge two sorted inputs, use separate read and write positions, or expand
outward from a possible center.

**Recognize it:** Look for sorted pair selection, comparisons between two ends,
merging sorted inputs, modifying a sequence in place, or testing symmetric
positions. For every movement, you must explain why the position left behind
cannot be needed again.

**Sorted pair example:** Find a pair totaling `9` in `[1, 2, 4, 7, 11]`.
`1 + 11` is too large, so move the right position left; moving the left one
would only increase the sum. Then `1 + 7` is too small, so move the left
position right and find `2 + 7 = 9`.

**Why sorted pair search stays correct:** No pair using a position already
passed by `left` or `right` can meet the target. Sorted order makes each discard
safe.

**Merging sorted inputs:** Track the smallest unread value in each input, copy
the smaller one, and advance only its position. When merging into unused space
at the end of the first input, work backward from the largest values so a write
cannot overwrite unread input.

**Read and write positions:** Let `read` inspect every value and `write` mark
where the next retained value belongs. The invariant is that the part before
`write` contains exactly the retained values seen so far, in their required
order. Because `write` never passes `read`, writing cannot destroy unread data.

**Expanding around centers:** Every palindrome is centered on either one
character or one gap. Test all character and gap centers, then move left and
right while the characters match. The already-accepted interior must be a
palindrome before a matching outer pair is added.

When at most one deletion is allowed, scan inward until the first mismatch.
Only two possibilities remain: skip the left value or skip the right value,
then verify the rest of each range with an ordinary palindrome scan. This is a
bounded branch, not general backtracking, so the total work remains `O(n)`.

**Work and memory:** Pair search, merging, and in-place compaction take linear
time. They use `O(1)` extra memory unless a merged output is created.
Expanding around all palindrome centers takes between `O(n)` and `O(n^2)` time
and `O(1)` extra memory.

**Common traps:** Sorting may destroy required original positions. The inward
pair-sum form requires sorted values, though negative values are fine. Use the
provided populated lengths when an array contains spare capacity, advance a
write position only after retaining a value, and test both character and gap
centers for palindromes.

**Reject it when:** No pointer movement can be proved to preserve every
remaining answer, or the method would require sorting data whose original
positions or order must remain intact.

**Representative problems:** [Palindrome Number](../problems/solved_easy_leetcode_0009_palindrome_number.go),
[Longest Palindromic Substring](../problems/solved_medium_leetcode_0005_longest_palindromic_substring.go),
[Remove Duplicates from Sorted Array](../problems/solved_easy_leetcode_0026_remove_duplicates_from_sorted_array.go),
and [Remove Element](../problems/solved_easy_leetcode_0027_remove_element.go).

##### Dynamic programming

**Mental model:** Dynamic programming (**DP**) starts with a search that tries
choices. If different paths keep reaching the same smaller question, answer
that question once and reuse the result.

**Recognize it:** All three conditions should hold:

1. The full answer can be built from answers to smaller versions of the
   problem.
2. A direct search answers the same smaller question more than once.
3. A small set of facts completely describes each smaller question.

That small set of identifying facts is called the **state**. For climbing
stairs, the current stair number is enough state because the choices available
next do not depend on how you reached that stair. Describe the state in one
sentence before choosing an array or map to store answers.

**Start here:** First write a rule that builds one answer from smaller answers;
this rule is called a **recurrence**. Then write the smallest answers known
without using that rule; these are the **base cases**. Finally, choose an order
that calculates every smaller answer before something needs it. You can start
from the requested answer and save results during recursive calls, or start
from the smallest answers and build upward with a loop.

**Worked example:** If each move climbs one or two stairs, every route to stair
`4` ends with one step from stair `3` or two steps from stair `2`. Therefore
`ways[4] = ways[3] + ways[2]`. Starting with `ways[0] = 1` and `ways[1] = 1`
produces `1, 1, 2, 3, 5` without recalculating earlier stairs.

**Why this counts every route exactly once:** Every non-empty route has one
final move. Removing that move puts the route into exactly one group: routes
whose final move was one step, or routes whose final move was two steps. Those
groups cover every allowed final move and do not overlap.

**Work and memory:** Without caching, the stair search grows roughly like
`2^n`. Saving the answer for each stair reduces it to `O(n)` time and `O(n)`
memory. When the next answer needs only a fixed number of earlier answers, keep
just those recent answers and reduce extra memory to `O(1)`.

Common state families recur across interview problems:

| Family | Typical state | Ordering rule |
| --- | --- | --- |
| Prefix or position | `dp[i]` answers for the first `i` items. | Process smaller prefixes before larger ones. |
| Take or skip | `dp[i]` is the best result ending at or through item `i`. | Retain exactly the earlier states each choice can follow. |
| Capacity or amount | `dp[c]` answers for capacity or total `c`. | When each item may be used once, scan capacity backward so it is not reused. Scan forward only when reuse is allowed. |
| Grid | `dp[row][col]` answers for a cell or rectangle boundary. | Visit every cell that can lead into this cell first. |
| Two sequences | `dp[i][j]` answers for prefixes of lengths `i` and `j`. | Fill by increasing prefix length or save each pair during recursive calls. |
| Interval | `dp[left][right]` answers inside one range. | Process shorter ranges before longer ranges. |
| Finite state machine | `dp[i][state]` records a small mode such as holding stock or transaction count. | Every transition must represent one legal action. |

**Reducing table memory:** State compression means retaining only the table
entries future calculations still need. Suppose a grid calculation uses the
cell above and the cell to the left. In one row-sized slice scanned left to
right, `dp[col]` still holds the result from above before it is replaced, while
`dp[col-1]` already holds the new result from the left. Scanning in the opposite
direction would change those meanings. Explain this dependency before reducing
`O(rows * cols)` storage to `O(cols)`; otherwise overwritten values may be
reused or lost incorrectly.

**Changing what a state stores:** For longest increasing subsequence, let
`tails[L]` store the smallest ending value found for any increasing subsequence
of length `L`, where `L` is that length. For `[3, 5, 2, 4, 1]`, the stored
endings change as follows:

```text
3 -> [3]
5 -> [3, 5]
2 -> [2, 5]
4 -> [2, 4]
1 -> [1, 4]
```

The two stored entries prove that a length-two subsequence exists, but `[1, 4]`
is not itself a subsequence of this input because `1` arrived after `4`.
Reconstructing the actual values requires each chosen value to link to the
earlier chosen value before it. For each new value,
[binary search](#binary-search) finds the first stored ending greater than or
equal to it; replacing that ending produces an `O(n log n)` length calculation
for strictly increasing subsequences. A non-decreasing variant must use a
different equality rule.

###### Coordinate compression

Coordinate compression replaces each distinct value with its position in
sorted order. For `[100, -5, 100, 20]`, the sorted distinct values are
`[-5, 20, 100]`, so the ranks are `[2, 0, 2, 1]`. The ranks are small, but
comparisons keep the same order as the original values.

This is useful when a transition asks for the best earlier state in a value
range such as `value-k <= earlierValue < value`. Store the best DP result at
each rank, query the permitted rank interval with a
[segment tree](#fenwick-and-segment-trees), and update the current rank. This
can reduce an `O(n^2)` scan of earlier states to `O(n log n)` when only a range
minimum or maximum is needed. Compression preserves value order, not original
input positions, so still process the values in their required sequence.

Derive complexity from the number of reachable states multiplied by the work
per transition. A two-dimensional table is not automatically `O(n^2)` if each
cell loops over many choices. **Memoization**, meaning saving the result of a
recursive state for reuse, does not help states that are never repeated.

**Common traps:** Writing the rule alone does not prevent repeated work; the
answers must also be saved. Do not store facts that cannot affect future
choices. Make the smallest known answers agree with the rule, update recent
values without destroying one that is still needed, count memory used by
unfinished recursive calls, and check whether the count can exceed the chosen
integer type.

**Reject it when:** The smaller questions do not repeat, the saved facts leave
out something that affects future choices, or there are still too many
different smaller questions for the input limits.

**Representative problem:** [Climbing Stairs](../problems/solved_easy_leetcode_0070_climbing_stairs.go).

##### Divide and conquer

**Mental model:** Split one problem into smaller independent parts, solve each
part by the same rule, and combine their answers.

**Recognize it:** A balanced split makes the input substantially smaller, the
parts can be solved without sharing evolving state, and their results have a
clear merge step. Merge sort, tree reconstruction, fast exponentiation, and
some geometric algorithms use this shape.

**Start here:** Define the smallest input that can be answered directly, how
the input is divided, and exactly what information the combine step needs. Then
write a recurrence that includes both recursive work and the combine cost.

**Worked example:** Merge-sort `[4, 1, 3, 2]` by splitting it into `[4, 1]`
and `[3, 2]`. Those split into one-item lists, which are already sorted. Merge
them into `[1, 4]` and `[2, 3]`, then compare the two front values repeatedly to
produce `[1, 2, 3, 4]`. No subproblem needs information from its sibling until
the final merge.

For merge sort, each call sorts two halves, then merges them by repeatedly
taking the smaller unread value. The invariant during merging is that the
output prefix is sorted and contains exactly the smallest values consumed from
both halves. The recurrence `T(n) = 2T(n/2) + O(n)` gives `O(n log n)` time.
The reason is concrete: repeatedly halving `n` creates about `log n` levels,
and all merges on one level process `n` values in total. Multiplying the levels
by the work per level gives `O(n log n)`. Ordinary slice-based merging needs
`O(n)` extra storage in addition to `O(log n)` recursive stack depth for
balanced splits.

Exponentiation by squaring uses the same idea without two recursive calls. If
`n` is even, `x^n = (x^(n/2))^2`; if odd, multiply by one extra `x`. Repeatedly
halving the exponent takes `O(log |n|)` multiplications. For example,
`x^10 = (x^5)^2`, and `x^5 = x * (x^2)^2`, so the calculation reuses squared
results instead of multiplying `x` ten times. For negative exponents, use
`x^(-n) = 1/x^n`. Take care when converting `-n` to a positive magnitude:
negating the most negative signed integer exceeds that type's positive limit.

**Contrast with dynamic programming:** Divide and conquer is most natural when
subproblems do not overlap. Dynamic programming is needed when different
branches repeatedly ask for the same state. Memoizing a divide-and-conquer
recursion because its parts overlap has effectively turned it into DP.

**Work and memory:** The cost depends on the number and sizes of the smaller
parts plus the work needed to combine them. Write that relationship before
claiming a bound: merge sort is `O(n log n)` time with `O(n)` merge storage,
while exponentiation by squaring is `O(log |n|)` time and can use `O(1)` extra
memory as a loop.

**Reject it when:** Splits can remain badly unbalanced, the parts repeatedly
solve the same smaller questions, or combining their answers costs as much as
solving the original problem directly.

##### Kadane's algorithm (best subarray scan)

**Mental model:** At each position, keep the best consecutive section that must
end here. It either extends the best section ending one position earlier or
starts fresh at the current value.

For value `x`, update:

```text
endingHere = max(x, endingHere + x)
bestOverall = max(bestOverall, endingHere)
```

**Worked example:** For `[-2, 3, -1, 4, -5]`, `endingHere` becomes `-2`, `3`,
`2`, `6`, and `1`. At `3`, starting fresh beats extending `-2`. At `-1` and
`4`, extending remains better. `bestOverall` reaches `6`, representing the
consecutive section `[3, -1, 4]`.

**Why it stays correct:** Every non-empty subarray ending at the current
position is either the current value alone or a subarray ending at the previous
position followed by the current value. Among extensions, only the best
previous ending can matter. `bestOverall` records the best ending among every
position processed so far.

Initialize from the first value when the answer must be non-empty. Initializing
both values to zero incorrectly returns zero for an all-negative input. To
return indices, remember the candidate start whenever beginning at the current
value is better, and save it when the global best improves.

For a circular array, a wrapping maximum equals the total sum minus the minimum
non-wrapping subarray because the wrapping answer keeps everything outside one
consecutive middle section. For `[5, -3, 5]`, the total is `7`; removing the
minimum section `[-3]` gives wrapping sum `7 - (-3) = 10`, representing the two
end values joined around the circle. Compare that result with ordinary Kadane.
When all values are negative, removing the minimum would remove the whole array
and produce an invalid empty answer, so use the ordinary non-empty result.

**Work and memory:** The scan takes `O(n)` time and `O(1)` extra memory.

**Reject it when:** The answer may skip values, is not additive, or deciding
whether to extend needs more information than the best section ending at the
previous position.

##### Stack

**Mental model:** Add and remove only at the top, so the most recently added
item is the first removed. This is last in, first out.

**Recognize it:** The most recently opened or unfinished item must be handled
first, as with nested brackets, undo history, expression evaluation, or a
depth-first search implemented with a loop.

**Start here:** Define what each stored item represents and what makes it
finished. Push newly opened work, inspect or pop only the top, and decide how an
empty stack is reported.

In the Go reference implementation, `Push` adds a value, `Pop` removes the
top, `Peek` inspects it, `Len` reports the count, and `IsEmpty` reports whether
the stack contains no values. `Pop` and `Peek` return a zero value plus `false`
when the stack is empty; the boolean distinguishes that default value from a
successful removal.

**Worked example:** To validate `([])`, push `(` and then `[`. When `]` arrives,
pop `[` and verify the match. When `)` arrives, pop `(`. The input is valid
because every closer matched and the stack is empty at the end.

**Why it stays correct:** The stack contains exactly the opened items not yet
completed. The top is the most recent unfinished item and must be completed
first.

A stack can also compress adjacent runs. Store `(value, count)` for each run;
merge a new value with the top when equal, and pop when a rule eliminates the
completed run. After a pop, recheck the new top because two previously separated
runs may now touch and combine. For example, if three equal letters disappear,
processing `aabbbaa` first removes the run `bbb`; the two stored `aa` runs then
touch and combine into `aaaa`. Rechecking the top handles that new adjacency.
This avoids repeated full rescans when each item is pushed and popped only a
constant number of times.

**Work and memory:** Push, pop, and top inspection take `O(1)` time, aside from
occasional slice growth when Go uses a slice as the stack. The stack may retain
`O(n)` items. Clear a removed slice position when it may hold pointers so Go's
garbage collector can reclaim objects no longer in use.

**Common traps:** Use a queue when the oldest waiting item must be handled
first. Removing from the front of a Go slice does not model a stack and may
retain a large underlying array. Recursive calls already use memory like a
stack behind the scenes; use your own stack when nesting may be unsafe or
direct control is needed.

**Reject it when:** The oldest item must be processed first or the algorithm
needs efficient access to arbitrary stored items.

**Representative problem:** [Valid Parentheses](../problems/solved_easy_leetcode_0020_valid_parentheses.go).

##### Sliding window

**Mental model:** Put a movable fence around one consecutive portion of the
input. The fence always contains neighboring items with no gaps.
Extend the right side to consider new values and move the left side when earlier
values must leave.

**Recognize it:** The prompt asks for a longest, shortest, largest, or smallest
consecutive portion of an array or string. You can update the tracked sum,
character counts, or other information when one item enters or leaves.

**Start here:** Define exactly what condition the current range must meet. Move
`right` once per step and add that item to the tracked information. Move `left`
and remove items until the range meets the condition again. Record the answer
at the moment required by the problem.

A common form gives the window a budget of at most `k` violations, such as
zeros that may be flipped or characters that may be replaced. Track the number
of violations or the amount required to repair the window. Shrink while the
budget is exceeded; afterward the window is the longest valid one ending at
the current right position under that invariant.

**Worked example:** Find the shortest positive-number range with sum at least
`7` in `[2, 3, 1, 2, 4, 3]`. Grow the window until `[2, 3, 1, 2]` sums to `8`,
then remove values from the left while the sum stays large enough. Later the
window `[4, 3]` reaches `7` with length `2`, which is the best answer. Positive
values make shrinking safe because adding future values can never reduce the
sum.

**Why moving the window is safe:** Before advancing `left`, explain why no
future value could make the current starting position part of a better answer.
For example, when all values are positive and the sum is already too large,
adding more values can only make it larger.

**Work and memory:** Each end moves forward at most `n` times, so the usual
window scan takes `O(n)` time. A running sum uses `O(1)` extra memory; character
or value counts use memory proportional to the number of distinct values kept.

**Reject it when:** The answer may skip items, or adding an item can make the
condition switch back and forth unpredictably. Mixed positive and negative
values often cause that problem for sum-based windows.

##### Prefix sums

**Mental model:** Store a summary of everything before each position, and when
useful a second summary of everything after it, so later answers do not repeat
the same work.

**Recognize it:** Many questions ask for sums or counts over overlapping ranges.
The simplest solution repeatedly adds values it has already added for an
earlier question.

**Start here:** Let `prefix[i]` mean "the total of the first `i` values." Include
`prefix[0] = 0`. The sum from `left` through `right` is then
`prefix[right+1] - prefix[left]`.

**Example:** For `[2, 4, 1, 3]`, store `[0, 2, 6, 7, 10]`. The sum from index
`1` through index `3` is `prefix[4] - prefix[1] = 10 - 2 = 8`.

**Why subtraction works:** Both stored totals begin at index `0`. Subtracting
the total before `left` removes exactly the values that are outside the
requested range.

The same preprocessing idea works beyond sums when the combination rule
supports the required query. For `Product of Array Except Self`, a left-to-right
pass stores the product before each position, then a right-to-left running
product multiplies in everything after it. Zeros need no division special case,
and the output slice can hold the prefix values while only one suffix product
uses extra memory.

Prefix frequencies can answer how many times a character appears in a range.
Prefix XOR works for range XOR because equal prefixes cancel. A prefix minimum
can answer "smallest value seen so far" but cannot generally be subtracted to
recover an arbitrary range minimum; use a monotonic deque for moving ranges or
a range-query structure for arbitrary updates and queries.

**Work and memory:** Building summaries for `n` values takes `O(n)` time and
usually `O(n)` memory. Each sum, count, or XOR range query then takes `O(1)`
time when its answer can be recovered from two stored prefixes.

**Reject it when:** The requested property cannot be recovered by combining two
running summaries. For example, you cannot find a range's smallest value by
subtracting two previously stored smallest values.

##### Difference arrays

**Mental model:** When many updates affect whole ranges, record where each
change starts and stops instead of modifying every covered position
immediately.

**Recognize it:** The problem applies many additions or subtractions to ranges
and asks for the final array after all updates.

**Start here:** Create a bookkeeping array with one more position than the
result array. For an update that adds `value` from `left` through `right`, add
`value` at `left` and subtract `value` at `right + 1`. After recording every
update, scan the positions of the result array from left to right while keeping
a running total of the changes.

**Example:** Begin with `[0, 0, 0, 0]` and add `3` to indices `1` through `3`.
In the five-position bookkeeping array, record `+3` at position `1` and `-3`
at position `4`. Running totals across the original four positions become
`[0, 3, 3, 3]`; position `4` merely marks where the update stops.

**Why it works:** The running change increases when an update begins and
decreases immediately after that update ends. Every position between those
points receives the update exactly once.

**Work and memory:** Recording `u` updates and rebuilding `n` final values
takes `O(u + n)` time and `O(n)` extra memory.

**Reject it when:** You must answer detailed questions after every individual
update rather than after the updates have been collected.

##### Fenwick and segment trees

**Mental model:** Store summaries for a hierarchy of ranges so an update or
query touches only `O(log n)` summaries instead of rescanning the entire input.

**ELI5:** Imagine asking for the total number of marbles in boxes `1` through
`12`. Instead of recounting every box, keep trustworthy totals for useful
groups of boxes. A query combines a few groups, and changing one box repairs
only the groups that contain it.

**Worked example:** A segment tree over `[2, 1, 5, 3]` stores leaf sums and
combines them upward: `[2, 1]` has sum `3`, `[5, 3]` has sum `8`, and the whole
range has sum `11`. Changing `5` to `6` updates its leaf, the right-half sum to
`9`, and the root to `12`; the unrelated left half stays untouched.

A **Fenwick tree** (binary indexed tree) is compact and especially useful for
point updates plus prefix sums. With one-based positions, `i & -i` gives the
size of the range summarized at position `i`. For example, position `6` is
binary `110`; its lowest set bit has value `2`, so the entry at position `6`
summarizes two positions, `5..6`. A prefix query at `6` adds that summary, then
removes the low bit to continue at `4`, whose summary covers `1..4`. An update
does the reverse: it adds the low bit to move to larger ranges containing the
changed position. Both operations take `O(log n)` time and the structure uses
`O(n)` memory. A range sum is the difference of two prefix sums.

A **segment tree** stores one summary for each range in a binary partition of
the array. It builds in `O(n)` and handles point updates in `O(log n)`. It also
answers range queries such as sum, minimum, or maximum in `O(log n)` when the
combining rule is **associative**, meaning that grouping the same values in a
different way does not change the answer.

Lazy propagation records a pending range update at an internal node and pushes
it to children only when necessary. For example, adding `3` to an entire half
of the array can update that half's stored sum and attach a pending `+3` marker
without immediately visiting every leaf. A later query that enters the half
passes the marker to the relevant children before reading them. This supports
many range updates and range queries without touching every element each time.

The invariant is that every node's stored value correctly summarizes its range
after accounting for any pending update. Define interval boundaries and the
combine **identity**, or neutral starting value, explicitly: zero is an
identity for sum because adding it changes nothing, but zero is not a safe
starting value for a minimum over arbitrary positive and negative inputs.

Prefer a prefix array for immutable data, a difference array when all updates
can be collected before one final reconstruction, and these trees only when
updates and queries must be interleaved. They are valuable but lower priority
than hash tables, ordinary traversal, sorting, and binary search for most
general interviews.

[Coordinate compression](#coordinate-compression) lets a Fenwick or segment
tree answer value-range queries even when the original values are negative or
enormous. Include every endpoint that a query may use, and distinguish a
value's sorted rank from its original input position.

**Work and memory:** Both trees use `O(n)` memory. A Fenwick tree supports
prefix queries and point updates in `O(log n)` time. A segment tree builds in
`O(n)` time and usually supports each compatible query or update in
`O(log n)` time.

**Reject it when:** The data never changes, all updates can be collected before
one final pass, or the number of operations is small enough that a direct scan
is clearer and fast enough.

##### Greedy

**Mental model:** Make one choice now and never undo it because you can prove
that it does not reduce the best result still possible.

**Recognize it:** Look for a problem that asks for a minimum or maximum and
builds the answer through repeated choices, such as selecting items, scheduling
jobs, or connecting nodes. Greedy is a candidate when one available choice may
be taken permanently without lowering the best result achievable afterward.

The deciding question is: **after making this choice permanently, can you prove
that the best result still possible has not become worse?** If the answer is
yes at every step, use greedy. If a discarded alternative could produce a
better result because of later choices, keep the alternatives with dynamic
programming or search instead.

Before committing to a greedy approach, answer all three questions:

1. **What are you optimizing?** State exactly what must be as small or large as
   possible and what facts affect the choices that remain.
2. **Why is the next choice safe?** Prove that taking it cannot make the best
   final result worse, even though you discard other choices.
3. **Can you repeat the argument?** After making the choice, the remaining work
   must be a smaller version of the same problem with the same rule.

**Safe to discard** has a precise meaning. Suppose the greedy rule chooses `A`
and rejects the alternatives. If the best achievable result before the choice
is `X`, it must still be possible to achieve `X` after choosing `A`. You do not
need to know the value of `X`; the proof establishes that it is unchanged.

One way to prove safety is to show that choice `A` is never worse than choice
`B`: every valid final answer beginning with `B` can be changed to begin with
`A` without becoming worse. This is sometimes called **dominance**. The
comparison must include both what the choice gains now and what it leaves
possible later.

Three common proof styles express that same basic idea:

- **Replace one choice (exchange):** Start with any best answer using `B`,
  replace `B` with `A`, and prove that the answer remains valid and no worse.
- **Remain at least as well placed (stays ahead):** Identify the facts that
  control future choices and prove that the greedy method keeps those facts at
  least as favorable as any competing method after every step.
- **Choose safely across a boundary (cut property):** Divide the work into a
  finished part and an unfinished part, then prove that `A` is a safe choice
  among everything connecting those two parts.

For interval scheduling, all first choices contribute one meeting. If `A` ends
no later than `B`, every meeting that could follow `B` can also follow `A`.
Therefore replacing `B` with `A` preserves every later choice, so `B` is safe
to discard for that position.

For coin values `[1, 3, 4]` and target `6`, choosing the largest coin does not
dominate choosing `3`. Choosing `4` leaves target `2`, which needs two more
coins, while choosing `3` leaves target `3`, which needs only one. The future
after `4` is worse, so the alternative `3` cannot be discarded.

**Cancellation as a greedy invariant:** Boyer-Moore majority vote keeps one
candidate and a balance. A matching value increases the balance; a different
value cancels one occurrence. If a value occupies more than half the input, it
cannot be completely canceled by all other values and must be the final
candidate. When the problem does not guarantee that a majority exists, make a
second pass to verify the candidate; the first pass alone returns a candidate,
not proof that its frequency is large enough.

Choose another method when any claim is missing:

- Use **dynamic programming** when different choices create genuinely
  different remaining situations, several may still lead to the best answer,
  and the same situations appear repeatedly.
- Use **backtracking or exhaustive search** when alternatives cannot be safely
  discarded and you must try them, especially when there is no small repeated
  situation whose answer can be saved.

**Start here:** State the objective and the exact local rule; "take the best"
is not precise enough. Now imagine a best possible answer that starts
differently. Replace its first choice with yours and check two facts: the
replacement remains valid, and its final value is no worse. Then verify that
the unchosen candidates still form the same kind of problem so the argument can
be repeated.

There is no universal greedy rule such as "always take the smallest value."
Each problem needs its own rule and proof. A practical process is:

1. Define what you are optimizing, which choices are currently allowed, and
   what facts describe the work remaining after one choice.
2. Propose a precise rule that leaves the remaining possibilities no worse
   than a competing choice would. Sorting, a priority queue, or a one-pass
   comparison often reveals that choice.
3. Prove the rule by replacing a choice, showing that it stays ahead, or
   showing that it is safe across the boundary between finished and unfinished
   work. If the proof fails, do not commit greedily.
4. Repeatedly take the proven-safe choice, discard choices it makes impossible,
   and solve the remaining smaller problem with the same rule.

**Execution on one input:** Suppose one meeting may start when another ends. To
attend the maximum number of meetings from `(1, 9)`, `(2, 4)`, `(5, 7)`, and
`(8, 10)`, the best possible result before making a choice is three: the last
three meetings are compatible, while `(1, 9)` overlaps all of them.

1. Sort the meetings by end time, producing `(2, 4)`, `(5, 7)`, `(1, 9)`,
   `(8, 10)`.
2. Set `lastEnd` to negative infinity and start with no selected meetings.
3. Scan the sorted meetings. Select a meeting exactly when its start is at
   least `lastEnd`; after selecting it, set `lastEnd` to its end.

| Meeting | Test | Action | `lastEnd` afterward |
| --- | --- | --- | --- |
| `(2, 4)` | `2 >= -infinity` | Select | `4` |
| `(5, 7)` | `5 >= 4` | Select | `7` |
| `(1, 9)` | `1 < 7` | Skip | `7` |
| `(8, 10)` | `8 >= 7` | Select | `10` |

Choosing `(2, 4)` first leaves `(5, 7)` and `(8, 10)` available, so the best
result still possible remains three. The choice has not made the answer worse.
Choosing `(1, 9)` first would leave no compatible meeting, reducing the best
result still possible from three to one. That choice is therefore unsafe.
This comparison explains this input, but it does not prove the rule for every
input.

**Proof for every input:** At any step, let `R` contain every meeting compatible
with the meetings already selected, and let `E` be the meeting in `R` with the
earliest end. Consider an optimal schedule from `R`, and call its first meeting
`F`.

Because `E` has the earliest end, `E.end <= F.end`. Replace `F` with `E`. If a
later meeting `L` followed `F`, then `L.start >= F.end`. Combining the two
inequalities gives `L.start >= F.end >= E.end`, so `L` can also follow `E`.
The replacement is therefore valid and preserves the schedule's number of
meetings. This proves that at least one optimal schedule from `R` begins with
the greedy choice `E`.

After committing to `E`, remove the meetings that overlap it. The remaining
meetings form another scheduling problem of the same kind, so the same argument
applies again. Repeating this replacement proof shows that every choice made by
the algorithm is safe, not merely that it works for the example above.

**Work and memory:** Sorting costs `O(n log n)` time, and the scan costs `O(n)`
time. The scan uses `O(1)` extra memory when the meetings are sorted in place,
excluding the space used to return the selected meetings.

**Why it stays correct:** After every choice, the best final value is still
achievable. The method removes only alternatives that are unnecessary for
reaching that value.

**Reject it when:** A choice that looks best now can block a better combination
later, and you cannot prove that the replacement described above is harmless.
Keep multiple possibilities with dynamic programming or search instead.

##### Backtracking

**Mental model:** Walk through a tree of decisions while carrying one partial
answer. At each branch, make a choice, explore what it can lead to, and undo the
choice before trying its sibling branch.

**Recognize it:** The answer is assembled one decision at a time, and each
partial answer has a clear set of legal next choices. You may need to find one
valid construction, generate every construction, or compare a manageable
number of constructions. Input rules often let you stop exploring a branch as
soon as its partial answer cannot become valid or cannot beat the best answer
already found.

Backtracking is a candidate when all of these are true:

1. A complete answer can be represented as a sequence of choices.
2. From a partial answer, you can enumerate the choices allowed next.
3. No single next choice can be proved safe enough to discard all alternatives.
4. There are few enough possibilities, or invalid branches can be stopped
   early enough, to make exploration practical.

**Procedure:** Define the partial answer, the facts needed to check whether the
next choice is allowed, and the condition that marks a complete answer. Then:

1. If the partial answer is complete, record a copy or report success.
2. Otherwise, list every legal next choice.
3. Apply one choice to the partial answer and its supporting facts.
4. Recursively explore what can follow that choice.
5. Undo every change made in step 3.
6. Repeat steps 3 through 5 for the other choices.

If only one answer is needed, return as soon as a recursive call succeeds. If
every answer is needed, finish every branch. If a partial answer is already
invalid or cannot improve the current best result, stop that branch immediately
instead of exploring the choices below it. This early stopping is called
**pruning**.

**Worked example:** Generate every subset of `[a, b]`. At each position, branch
into including the current value or excluding it:

1. Start with `[]`; include `a`, producing `[a]`.
2. Include `b`, producing the complete subset `[a, b]`; save a copy.
3. Undo the inclusion of `b`, returning to `[a]`; exclude `b` and save `[a]`.
4. Undo the inclusion of `a`, returning to `[]`; exclude `a`.
5. Include `b` and save `[b]`; undo it, exclude it, and save `[]`.

The crucial detail is that undoing `b` restores exactly `[a]`, and undoing `a`
restores exactly `[]`. Each sibling branch therefore begins with only the
choices made on its own path.

**Work and memory:** Generating subsets visits `2^n` complete branches. Saving
a copy of every subset takes `O(n * 2^n)` total time and `O(n * 2^n)` output
space. The current path and unfinished function calls use `O(n)` extra memory,
excluding the output.

**Why it stays correct:** At the start of each recursive call, the partial
answer and its supporting facts describe exactly the choices on the path to
that call. After the call returns, undoing its choice restores exactly what the
caller had before trying that branch.

**Reject it when:**

- Use **greedy** when one next choice can be proved not to reduce the best
  result still possible, so the other branches may be discarded.
- Use **dynamic programming** when many choice sequences reach the same
  remaining situation and only its best value, count, or yes/no result matters.
  Save that result instead of rebuilding the same search tree.
- Use ordinary **[depth-first search (DFS)](#depth-first-search)** when
  traversing an existing graph or tree without constructing and undoing a
  candidate answer.
- Avoid exhaustive backtracking when there are too many possibilities and no
  useful early-stopping rule or repeated smaller question exists.

##### Depth-first search

**Mental model:** Depth-first search (**DFS**) follows one route as far as
possible. When that route ends, return to the most recent place that had
another route available.

**Recognize it:** You need to visit an entire connected region, inspect every
complete path, or calculate a tree result from the results below each node. You
do not need the first discovered path to use the fewest moves.

**Start here:** Define one position and how to list every position reachable in
one move. Mark a position reached before exploring from it so a loop cannot send
the search around forever. Use nested function calls or your own stack to
remember where to return.

**Example:** To calculate each directory's total size, descend through its files
and subdirectories first. Once a subdirectory has been fully explored, return
its total and add it to its parent directory's size.

**Why it stays correct:** Every position marked reached has either been fully
explored or is waiting on the current return path. No position needs to be
scheduled twice.

**Common trap:** Nested function calls consume memory. A path with too many
steps may exceed the program's call limit, so use your own stack collection for
very deep inputs.

**Work and memory:** With an adjacency list, a complete traversal takes
`O(V + E)` time because it visits each vertex and examines each edge a constant
number of times. Reached state and the explicit or recursive stack use `O(V)`
extra memory in the worst case.

For cycle detection, one permanent `visited` flag is not always enough. In a
directed graph, use three states: unseen, currently on the active DFS path, and
fully finished. An edge to an active vertex is a directed cycle; an edge to a
finished vertex is not. In an undirected graph, ignore the edge back to the
current vertex's parent and treat another visited neighbor as a cycle.

**Reject it when:** You need a route with the fewest equal-cost moves; use BFS
instead. Also avoid recursive DFS when a valid path can be too deep for the
call stack unless the environment guarantees enough stack space.

##### Breadth-first search

**Mental model:** Breadth-first search (**BFS**) explores outward in layers:
first everything one move away, then everything two moves away, and so on.

**Recognize it:** You need the fewest moves through a network where every move
costs the same, the nearest reachable position, or tree nodes grouped by depth.

**Start here:** Use a first-in, first-out queue. Add the starting position and
mark it reached immediately. Repeatedly remove the oldest position, add each
unseen neighbor, and record its distance as one more than the current distance.

**Example:** In a maze, process the start at distance `0`, its open neighbors at
distance `1`, and their unseen neighbors at distance `2`. The first time the
exit is added, you have found a route with the fewest moves.

**Why it stays correct:** The queue never processes a position with a greater
distance before a position with a smaller distance. Each position is added at
most once.

**Multi-source BFS:** When the question asks for the nearest source for every
position, add all sources to the queue at distance zero before the search
starts. The layers then grow from every source simultaneously, and the first
time a position is reached is its minimum distance to any source. Running a
separate BFS from every source repeats work and is usually much slower.

**Bidirectional BFS:** When both one start and one target are known and moves
can be reversed, grow a search from both ends until they meet. A **frontier** is
the outermost layer reached so far. Expanding the smaller frontier can reduce
the number of states explored dramatically, though the worst-case growth-rate
bound may be unchanged. Keep the two visited-distance maps separate so the
meeting distance can be combined correctly.

**Work and memory:** With an adjacency list, ordinary BFS takes `O(V + E)` time
and `O(V)` extra memory for its queue, distances, and reached state.

**Reject it when:** Moves have different costs. The route with the fewest moves
may then cost more than a longer route; use Dijkstra's algorithm when all costs
are zero or greater.

##### Topological sort

**Mental model:** Repeatedly perform tasks whose prerequisites are already
satisfied.

**Recognize it:** The prompt describes courses, builds, jobs, or other tasks
with rules such as "A must finish before B."

**Start here:** For every prerequisite rule, record which later task it unlocks.
Count each task's unfinished prerequisites. Add every task with a count of zero
to a first-in, first-out queue. When one task finishes, subtract one from the
count of every task that depended on it.

**Worked example:** Suppose `Shop` must happen before `Cook`, and `Cook` before
`Eat`. Initially only `Shop` has zero unfinished prerequisites, so output it.
That reduces `Cook`'s count to zero; output `Cook`, which then unlocks `Eat`.
The result is `[Shop, Cook, Eat]`. Adding `Eat -> Shop` would leave every task
waiting forever and correctly reveal a cycle.

**Why it stays correct:** The queue contains exactly the tasks whose prerequisites are all
finished. A task is added to the output only after its remaining prerequisite
count reaches zero.

**Failure signal:** If the queue becomes empty before every task is processed,
the graph of unprocessed tasks contains a dependency cycle, such as `A`
needing `B` while `B` also needs `A`. Some other unprocessed tasks may not be
part of that cycle; they can simply be waiting on a task that is. No valid
order exists.

An alternative uses DFS with the three states described above. After all
outgoing dependencies of a vertex are finished, append that vertex to a
postorder list; reverse the list for a topological order. The queue-based method
above is often simpler when the problem also needs layer-by-layer processing or
an explicit count of unfinished prerequisites.

**Work and memory:** Both forms take `O(V + E)` time and `O(V)` extra memory
with adjacency lists, not counting the stored graph itself.

**Reject it when:** The relationships do not represent prerequisites, the
graph may be undirected, or the problem asks for shortest routes rather than a
valid dependency order.

##### Dijkstra shortest paths

**Mental model:** On a map with travel costs, always process the unvisited
location with the cheapest known route from the starting location.

**Recognize it:** You need the cheapest routes through a network, and every
connection has a cost of zero or greater.

**Start here:** Record cost `0` for the starting location and "unknown" for all
others. Use a priority queue that removes the route with the smallest cost
first. When traveling through the current location gives a neighbor a cheaper
route, store that new cost and add the route to the queue.

**Worked example:** Let `A -> B` cost `4`, `A -> C` cost `1`, and `C -> B`
cost `2`. Starting at `A` discovers `B = 4` and `C = 1`. The queue processes
`C` first because `1` is cheaper, then improves `B` to `1 + 2 = 3`. Processing
`B` finalizes cost `3`; the older queued entry with cost `4` is **stale**,
meaning it no longer equals the stored best cost, so it is ignored.

**Why a processed cost is final:** Suppose the queue removes a route costing
`5`, and `5` is still the cheapest known cost for that location. Every untried
route already costs at least `5` before taking another connection. Because a
connection cannot subtract cost, none can produce a route below `5`.

**Work and memory:** With adjacency lists and a **binary heap**, a priority
queue organized as a tree with at most two children per position, the usual
implementation takes `O((V + E) log V)` time and `O(V + E)` storage including
the graph. Some priority queues support **decrease-key**, which edits the
priority of an entry already stored. When that operation is unavailable, push
a new improved entry and later discard any popped entry whose cost no longer
equals the stored best cost.

**Reject it when:** Any connection may have a negative cost. If every connection
has the same cost, breadth-first search is simpler.

##### Bellman-Ford shortest paths

**Mental model:** Let cheaper route costs ripple through every connection,
repeating until an improvement can cross even the longest useful route.

**Recognize it:** You need shortest routes from one starting location, some
connections may have negative costs, and you may need to detect a reachable
negative-cost cycle.

**Start here:** Record cost `0` for the starting location and "unknown" for all
others. Repeat a full pass over every directed connection up to `V - 1` times,
where `V` is the number of locations. For an edge `from -> to` with cost `w`,
replace `distance[to]` with `distance[from] + w` when `from` is reachable and
that total is smaller. This attempt to lower a stored cost is often called
**relaxing an edge**. Stop early if a full pass changes nothing.

**Example:** Suppose `A -> B` costs `4`, `A -> C` costs `7`, and `C -> B` costs
`-5`. Starting from `A`, checking the first two connections sets `B = 4` and
`C = 7`. Checking `C -> B` then improves `B` to `7 + (-5) = 2`. Dijkstra
cannot safely finalize `B = 4` before that later improvement, but Bellman-Ford
keeps checking every connection.

**Why `V - 1` passes are enough:** A finite shortest route never needs to
repeat a location; removing any repeated non-negative cycle cannot make it more
expensive, while a repeated negative cycle would mean there is no finite
shortest route. A route without repeated locations uses at most `V - 1` edges.
After pass `k`, every shortest route using at most `k` edges has had a chance to
propagate its cost. Therefore `V - 1` passes cover every finite shortest route.

**Checking for a negative-cost loop:** Make one additional pass. If a reachable
connection can still lower a cost, the route can loop and become cheaper every
time around. There is then no single lowest cost for locations on that loop or
locations reachable from it. This is called a **negative cycle**.

**Work and memory:** If `V` is the number of locations and `E` is the number of
connections, up to `V - 1` full passes take `O(VE)` time. The stored distances
use `O(V)` extra memory, not counting the input map or optional links used to
rebuild the chosen routes.

**Reject it when:** Every cost is non-negative and the graph is large enough
that Dijkstra's better running time matters. If a reachable negative cycle may
exist, do not report ordinary finite shortest costs for locations affected by
that cycle.

##### Disjoint-set union

**Mental model:** Imagine that every connected group has one shared group
leader. To check whether two items are connected, find each item's leader and
compare them. To connect two different groups, make one leader report to the
other.

**Recognize it:** Connections are added over time, and queries ask whether two
items are already connected, whether a new connection creates a loop, or how
many separate groups remain.

**Start here:** Store a `parent` for every item. Initially, each item is its own
parent because every item is in a separate group. `find(item)` follows parent
links until it reaches the group leader. `union(a, b)` finds both leaders and,
if they differ, makes one leader the parent of the other.

**Walk-through:** Start with `parent[1] = 1`, `parent[2] = 2`, and
`parent[3] = 3`. After connecting `1` and `2`, make `2` point to `1`. After
connecting `2` and `3`, follow `2 -> 1` and make `3` point to `1`. Now both
`find(1)` and `find(3)` reach leader `1`, so the items are connected.

**Why it stays correct:** Two items are connected exactly when following their parent
links reaches the same group leader.

**Speed improvements:** While finding a leader, make each visited item point
closer to that leader so later searches take fewer steps. This is called
**path compression**. When joining groups, attach the smaller group below the
larger group so the chains stay short; this is called **union by size**. Both
changes improve speed without changing which items belong together.

**Work and memory:** With both improvements, a long sequence of operations is
almost constant time per operation in practice. Spread across the whole
sequence, which is what **amortized** means here, the formal bound is
`O(alpha(n))`, where `alpha` is a mathematical function called the inverse
Ackermann function. You do not need to calculate it for implementation work:
it grows so slowly that it remains below a tiny constant for any realistic
input. The parent and size arrays use `O(n)` memory.

**Reject it when:** You need the actual route, the order in which items were
visited, or the shortest path. Knowing only each item's group does not preserve
that information.

##### Minimum spanning tree

**Mental model:** Choose enough weighted edges to connect every vertex without
forming a cycle, while minimizing their total cost. The result spans all
vertices and has exactly `V - 1` edges when the graph is connected.

Do not confuse a minimum spanning tree with shortest paths. A spanning tree
minimizes the total cost of the selected network; paths from a particular start
inside that tree need not be the cheapest possible paths in the original graph.

**Worked example:** Three offices have cable costs `A-B = 1`, `B-C = 2`, and
`A-C = 5`. Kruskal accepts `A-B`, then `B-C`; all offices are connected for
total cost `3`. It rejects `A-C` because that expensive edge would create a
cycle. The result uses exactly two edges, which is `V - 1` for three offices.

**Kruskal's algorithm:** Sort edges by weight. Scan from cheapest to most
expensive and accept an edge exactly when its endpoints are currently in
different disjoint-set components. The invariant is that the accepted edges
form a **forest**, meaning a collection of separate trees with no cycles, that
can still be extended to some minimum spanning tree. The **cut property** says
that if the vertices are split into two groups, a cheapest edge crossing
between those groups can be included in some minimum spanning tree. That is
why Kruskal can safely accept the next cheapest edge that joins two current
components. Sorting dominates the work at `O(E log E)`; disjoint-set operations
are nearly constant amortized time with path compression and union by size.

**Prim's algorithm:** Begin from any vertex and keep the cheapest edge leaving
the connected portion in a min-heap. Repeatedly take the cheapest edge reaching
an unseen vertex. With adjacency lists and a binary heap, this takes
`O(E log V)` time. Ignore stale heap entries whose destination has already been
added.

**Work and memory:** Kruskal takes `O(E log E)` time and `O(V + E)` storage
including its edge list and disjoint-set data. Prim with adjacency lists and a
binary heap takes `O(E log V)` time and `O(V + E)` storage including the graph.

**Failure signal:** If fewer than `V - 1` edges can connect all vertices, the
graph is disconnected and no single spanning tree exists. Kruskal naturally
produces a **minimum spanning forest**: one minimum-cost tree for each separate
connected group. Prim as described above covers only the group containing its
starting vertex; restart it from every still-unseen vertex to produce the same
kind of forest for the whole graph.

**Reject it when:** The goal is the cheapest route from one starting vertex,
not the cheapest total network, or when the directed edges do not represent an
undirected connection problem.

##### Heap or priority queue

**Mental model:** Keep the next item you need (the smallest, largest, or most
urgent) easy to remove without fully sorting everything after each change.

**Recognize it:** You repeatedly need the smallest or largest item, a requested
number of the best items, or the next event in priority order.

**Start here:** Decide which item must come out next. A min-heap is a priority
queue that exposes the smallest stored item. To retain the three largest values
seen so far, keep three values in a min-heap; when a larger value arrives,
remove the smallest retained value and insert the new one.

**Worked example:** Process job priorities `5`, `2`, and `8` with a min-heap.
Although `2` arrived second, it is removed first because it is smallest. The
next removal returns `5`, then `8`. The heap promises only which item is at the
top; its internal slice does not need to be globally sorted.

**Why it stays correct:** The heap contains exactly the items still under consideration.
The item exposed at the top is always the next smallest or largest item the
algorithm needs.

**Work and memory:** Building a heap from `n` existing items can be done in
`O(n)` time and uses `O(n)` storage. Adding or removing the top takes
`O(log n)`, and inspecting the top takes `O(1)`. A heap does not provide
constant-time search or removal for an arbitrary value unless an additional
index map is maintained and repaired after every swap.

For `k` sorted inputs containing `N` total items, put the first item from each
non-empty input into a min-heap. Whenever one is removed, add the next item from
that same input. The heap contains at most `k` candidates and its top is the
smallest unread item overall, giving `O(N log k)` time and `O(k)` extra
space. If output must be unique, compare against the last emitted value; do not
drop duplicates before they have advanced their source input.

**Reject it when:** You need the complete final order of an unchanging set of
items; sorting is usually simpler. A heap also does not make it fast to find any
item other than the smallest, largest, or highest-priority one it was designed
to expose.

##### Monotonic stack

**Mental model:** Keep earlier values whose answer is still unknown, ordered in
only one direction: always increasing or always decreasing. That one-direction
order is what "monotonic" means here.

**Recognize it:** The prompt asks for the next greater or smaller value, the
nearest earlier or later value that becomes larger or smaller, or how long a
value remains the minimum or maximum.

**Start here:** Decide whether to store positions or values and whether the
stored values increase or decrease from bottom to top. While the new value
answers the question for the top item, remove that item and record its answer.
Then add the new item, whose own answer is still unknown.

**Worked example:** Find the next greater value in `[2, 1, 4]`. Store position
`0` for value `2`, then position `1` for value `1`; both still await a larger
value. When `4` arrives, it resolves `1 -> 4` and then `2 -> 4`, so both
positions are popped. Store `4`; it has no next greater value by the end.

**Why it stays correct:** Every stored item is still waiting for an answer, and
the stored values remain sorted in one direction. An item is removed only when
the current value is the first later value that answers its question.

**Work and memory:** Each position is added once and removed at most once, so
the whole scan takes `O(n)` time even though one step may remove several items.
The stack can hold `O(n)` positions.

**Reject it when:** The answer depends on any earlier value, not specifically
the nearest one that meets the condition, or when removing an item would lose
information that a later value still needs.

##### Monotonic deque

**Mental model:** Keep only candidates that could still become the minimum or
maximum of a current or future window. The front is the answer for the current
window; the back is where dominated candidates are removed.

For a moving-window maximum, store positions whose values decrease from front
to back:

1. Remove front positions that have left the window.
2. While the new value is at least the value at the back, remove the back; the
   new value is both newer and no smaller, so the old one can never win again.
3. Add the new position. The front now identifies the maximum.

**Worked example:** For `[1, 3, 2, 5]` with window size `3`, reading `3`
removes `1` from the back because `1` is older and smaller. After reading `2`,
the deque represents values `[3, 2]`, so the first window's maximum is `3`.
Reading `5` removes both `2` and `3`; the next window's maximum is `5`.

**Why it stays correct:** Positions remain in increasing time order, values
remain monotonic, and every removed position is either expired or dominated by
a later candidate that will remain available for at least as long.

**Work and memory:** Each position enters and leaves the deque at most once, so
all windows are processed in `O(n)` time with `O(window size)` memory. Store
positions rather than values when expiration depends on location or duplicate
values exist.

Use a heap when priorities change independently or arbitrary items remain
eligible; use a monotonic deque when eligibility expires in sequence order and
new values permanently dominate older ones.

**Reject it when:** Candidates do not expire in input order, or a newer value
cannot be proved to make an older value permanently irrelevant.

##### Intervals

**Mental model:** Each interval covers everything between a start and an end.
Sort intervals by start so overlaps appear next to one another.

**Recognize it:** The input contains ranges, and the output must combine
overlapping ranges or report the gaps between covered regions.

**Start here:** Sort by start and keep one current interval. Compare the next
start with the current end using the problem's rule for touching endpoints. If
they overlap, extend the current end when necessary. If a gap separates them,
save the current interval and begin a new one.

**Worked example:** Sort `[5, 7]`, `[1, 3]`, and `[2, 6]` into start order.
Begin with `[1, 3]`; `[2, 6]` overlaps, so extend the current interval to
`[1, 6]`. `[5, 7]` also overlaps, producing the single merged interval
`[1, 7]`.

**Why saving an interval is safe:** Once a gap appears, every remaining
interval starts even later because the input is sorted. None can overlap the
interval you just saved.

**Common trap:** The problem must define whether `[1, 2]` and `[2, 3]` overlap.
If both ending points are included, the ranges share `2`. If an end marks the
first excluded value, the first range stops immediately before `2`, so they do
not overlap.

Different interval objectives require different rules:

- To select the largest number of non-overlapping intervals, sort by end and
  greedily accept the next compatible interval. Finishing earliest leaves every
  later option available.
- To maximize total interval weight or profit, earliest finish is not enough.
  Sort by end, binary-search the last compatible interval for each choice, and
  use DP to compare taking it plus its compatible prefix against skipping it.
- To count simultaneous intervals or required rooms, sort start/end events or
  keep current ending times in a min-heap. The endpoint convention decides
  whether an end at time `t` frees a resource before a start at `t`.

**Work and memory:** Sorting `n` intervals takes `O(n log n)` time, and the
following scan takes `O(n)`. Sorting in place uses no separate interval copy,
but the merged or selected result can contain `O(n)` intervals.

**Reject it when:** The ranges change online and each change must be queried
immediately, or when the problem is really about individual points rather than
relationships between ranges.

##### Sweep line

**Mental model:** Move from left to right across a timeline or number line.
Record only the positions where something starts or ends, because nothing
changes between those positions.

**Recognize it:** You need the maximum number of simultaneous bookings, the
number of active ranges at each position, or another result that changes only
when a range starts or ends.

**Start here:** Turn each range into two events. A start event adds one to the
active count, and an end event subtracts one. Sort the events by position and
update the count in that order.

**Example:** One booking starts at `1` and ends just before `4`; another starts
at `3` and ends just before `5`. Process `start at 1`, `start at 3`, `end at 4`,
and `end at 5`. The active counts are `1`, `2`, `1`, and `0`, so the maximum is
`2`.

**Why it stays correct:** After processing all events at a position, the running count
equals the number of ranges active immediately after that position.

**Common trap:** Events at the same position need an explicit rule. If a booking
ending at `4` frees its room before another starts at `4`, process the ending
before the new start.

**Work and memory:** Creating and sorting two events per range takes
`O(n log n)` time and `O(n)` event storage. The sweep itself is linear after
sorting.

**Reject it when:** Values change continuously between event positions, or
events arrive online and must be answered before they can be collected and
sorted.

##### Fast and slow pointers

**Mental model:** Move two positions through the same linked list. Their speeds
or the distance between them reveals a loop, the middle, or a node near the end.

**Recognize it:** A linked list may loop back to an earlier node, you need its
middle without first counting every node, or one position must remain a fixed
number of nodes behind another.

The structure need not be stored as nodes. Repeatedly applying a deterministic
function `next(state)` creates a **functional graph** in which every state has
one outgoing edge. Fast and slow states can therefore detect repetition in
digit transforms or other generated sequences without storing every seen state.

**Start here:** Decide how fast each position moves or how far apart they must
remain. Before following the link to the next node, check that the current node
exists. For a fixed distance, move the leading position ahead first and then
move both one node at a time.

**Worked example:** In `1 -> 2 -> 3 -> 4 -> nil`, move `slow` one link and
`fast` two links. After the first move they are at `2` and `3`; after the next,
`slow` is at `3` while `fast` has reached the end. The slow position is the
second middle under this stopping rule. State the desired middle for even
lengths because a different loop condition can return node `2` instead.

**Why it stays correct:** After setup, the pointers keep the same speed
relationship or the same fixed gap after every move.

**Work and memory:** Each pointer makes at most a linear number of moves, so the
method takes `O(n)` time and `O(1)` extra memory.

**Reject it when:** You need to jump directly to any chosen position, remember
the full history, or return the actual visited path. Two moving positions do not
retain those details.

##### Trie

**Mental model:** Store words one character at a time. Words with the same
beginning reuse the same initial steps.

**Recognize it:** You repeatedly ask for words beginning with the same letters,
check dictionary words one character at a time, or want to stop as soon as a
beginning does not exist.

**Start here:** Each stored step records possible next characters and whether a
complete word ends there. To insert a word, follow or create one step for each
character. To look up a word, follow the same steps and check the end marker
after the last character.

**Worked example:** Insert `car` and `cat`. Both words share the steps `c -> a`
and then branch to `r` or `t`. A prefix query for `ca` succeeds at the shared
step. A full-word query for `ca` fails unless that step has its own end marker;
being a valid prefix is not the same as being a stored word.

Use a fixed child array for a small fixed alphabet and a map for a sparse or
large alphabet. A wildcard search branches over every existing child at that
position, so its worst case can be exponential in the number of wildcards even
though ordinary exact lookup is `O(word length)`. During a board search, carry
the trie node reached by the current path so a prefix missing from the trie can
stop that branch immediately.

**Work and memory:** Inserting or looking up a word of length `L` takes `O(L)`
time. An insertion may create `O(L)` nodes; across all words, storage is
proportional to the number of distinct prefix nodes. Wildcards can require
exploring many branches and may take exponential time in the number of
wildcards.

**Why it stays correct:** The characters followed from the starting point always spell the
beginning represented by the current step. A word exists only if the last step
is marked as the end of a complete word.

**Reject it when:** You only need to know whether a complete word exists. A hash
set is usually simpler and may use less memory.

#### Choosing among overlapping methods

- Sliding window versus prefix sums: use a window when moving its left or right
  boundary changes the tracked condition predictably, so an earlier starting
  position is never needed again. Use prefix sums when a range total can be
  calculated from two stored totals but boundary movement is not predictable.
- Prefix sums versus Fenwick or segment trees: prefix sums are simpler and give
  constant-time queries when the input does not change. Use a range-query tree
  when updates and queries are interleaved and logarithmic time for both is
  required.
- Two pointers versus binary search: use two pointers when both positions can
  move through the data together and each move permanently rules out some
  possibilities. Use binary search when testing the middle independently rules
  out half of the remaining possibilities.
- Greedy versus dynamic programming: use greedy when one choice can be proved
  not to reduce the best result still possible, and the same proof repeats
  after committing. Use dynamic programming when competing choices lead to
  different future states that must remain under consideration and those
  states recur.
- Dynamic programming versus backtracking: use dynamic programming when many
  choice sequences lead to the same remaining problem and you only need a
  total, minimum, maximum, or yes/no answer. Use backtracking when you must
  produce each individual result or when two paths with similar progress can
  still have different legal next choices.
- Divide and conquer versus dynamic programming: both define smaller problems.
  Use divide and conquer when the smaller parts are independent and solved once;
  use dynamic programming when branches overlap and repeatedly request the same
  state.
- Backtracking versus DFS: DFS describes the order of exploration: follow one
  route fully before returning. Backtracking uses that order to build an answer
  through choices, then removes each choice before trying the next one. A
  backtracking solution therefore often performs a DFS, but not every DFS
  builds and undoes choices.
- DFS versus BFS: use DFS when you must finish one path, tree branch, or
  connected region before returning to alternatives. Use BFS when you need the
  nearest result or fewest moves and every move has the same cost.
- BFS versus Dijkstra: both explore outward from a starting location. BFS
  handles every waiting location in arrival order and gives cheapest routes
  only when every move costs the same. Dijkstra instead handles the location
  with the cheapest known route next, so it supports differing costs as long as
  none are negative. Dijkstra is therefore related to BFS, but it is not
  ordinary BFS.
- Dijkstra versus Bellman-Ford: both find cheapest routes from one starting
  location. Prefer Dijkstra when every connection cost is non-negative because
  it avoids repeatedly scanning every connection. Use Bellman-Ford when costs
  may be negative or when you must detect a reachable negative-cost loop.
- Two pointers versus fast and slow pointers: both coordinate two moving
  positions. The fast/slow version is a specialized two-pointer method for
  linked lists and similar structures, where different speeds or a fixed gap
  reveal a loop, a middle, or a position near the end.
- Heap versus sorting: sort when you need every item in order or all data is
  available at once. Use a heap when you repeatedly need only the next smallest
  or largest item, or when you must retain a fixed number of best items as new
  values arrive.
- Heap versus quickselect: use quickselect for one rank in mutable batch data
  when expected linear time is worthwhile. Use a heap for streaming input,
  repeated extraction, or a simpler `O(n log k)` top-`k` bound.
- Monotonic stack versus monotonic deque: a stack resolves nearest earlier or
  later relationships as new values arrive. A deque maintains an extreme over
  a moving window, removing expired items from the front and dominated items
  from the back.

#### Let constraints eliminate bad choices

Let `n` mean the number of input items. The notation describes how quickly the
work grows: `O(n)` is roughly one pass, `O(n^2)` examines pairs, and `O(n^3)`
examines triples. These are rough interview-scale checks, not guarantees;
memory, programming language, and time limits still matter.

| Input scale | Methods worth considering |
| --- | --- |
| `n <= 10` | Trying every ordering (`n!` possibilities) may be practical if impossible partial answers stop early. |
| `n <= 20` | Trying every subset (`2^n` possibilities) may be practical. |
| `n <= 100` | Three nested passes (`O(n^3)`) may be possible; examining every pair (`O(n^2)`) is commonly reasonable. |
| `n <= 1,000` | Examining every pair (`O(n^2)`) may be possible but needs scrutiny. |
| `n <= 100,000` | Usually target sorting or a similar amount of work (`O(n log n)`), or a small number of full passes (`O(n)`). |
| `n >= 1,000,000` | Usually target a few full passes and avoid creating large extra arrays, maps, or sets. |

Also count the work required to produce the output. If the problem asks for all
`2^n` subsets or all `n!` orderings, merely writing those results already takes
that much work; generating them in a different order cannot remove it.

#### Before committing to a method

Write one sentence for each item:

- **Stored meaning:** What facts does each variable, table entry, queue item, or
  stack entry represent?
- **Invariant:** What fact remains true before and after every step?
- **Progress:** Why does each step reduce the work left to do or move closer to
  finishing?
- **Coverage:** Why can none of the choices you skipped or discarded contain an
  answer?
- **Work and memory:** How many smaller problems, positions, or elements are
  processed, how often, and what extra information is retained?

Then rehearse the follow-ups that recent interviews frequently attach to an
otherwise familiar problem:

- What changes if equal values are allowed, input is empty, or values may be
  negative?
- What changes if the input cannot be mutated, arrives as a stream, or is too
  large to retain?
- Can the method return an actual path, ordering, or chosen indices rather than
  only the optimal value?
- Can preprocessing make repeated queries faster, and what update pattern would
  invalidate that preprocessing?
- Which assumption makes the optimization correct, and what is the fallback
  when that assumption is removed?

If one sentence cannot be made precise, the method is still only a guess.

## Go reference implementations

These snippets show the ideas above as small Go programs. They are examples to
read and adapt, not packages built by this repository. The note after each
snippet lists important boundary cases to test.

Read a snippet in three passes: first read the function comment to learn its
contract, then trace one ordinary input through the changing variables, and
finally check how it handles the listed boundary cases. The matching pattern
card above explains why the code is correct; the snippet shows how that idea is
expressed in Go.

Some snippets use a type parameter written as `T`. Treat `T` as a placeholder
for the value type chosen by the caller, such as `int` or `string`; the
[type-parameter note](#type-parameters) explains why this is useful.

### Binary search bounds

```go

package binarysearch

// LowerBound returns the first index whose value is greater than or equal to
// target. It returns len(values) when no such index exists.
func LowerBound(values []int, target int) int {
	left, right := 0, len(values)
	for left < right {
		middle := left + (right-left)/2
		if values[middle] < target {
			left = middle + 1
		} else {
			right = middle
		}
	}
	return left
}

// UpperBound returns the first index whose value is greater than target. It
// returns len(values) when no such index exists.
func UpperBound(values []int, target int) int {
	left, right := 0, len(values)
	for left < right {
		middle := left + (right-left)/2
		if values[middle] <= target {
			left = middle + 1
		} else {
			right = middle
		}
	}
	return left
}

// Search returns the index of target in a sorted slice, or -1 when target is
// absent. When duplicates exist, it returns the first matching index.
func Search(values []int, target int) int {
	index := LowerBound(values, target)
	if index == len(values) || values[index] != target {
		return -1
	}
	return index
}

```

**Cases to test:** Empty input, targets before and after the data,
duplicate ranges, values between entries, first-match lookup, and absent
targets.

### Rolling dynamic programming

```go

package dynamicprogramming

// CountOrderedOneOrTwoStepWays returns the number of ordered sequences of
// one-step and two-step moves that sum to steps. Negative totals have no valid
// sequence; the empty sequence is the single way to make a total of zero. The
// result must fit in an int.
func CountOrderedOneOrTwoStepWays(steps int) int {
	if steps < 0 {
		return 0
	}
	if steps < 2 {
		return 1
	}

	twoBack, oneBack := 1, 1
	for current := 2; current <= steps; current++ {
		twoBack, oneBack = oneBack, twoBack+oneBack
	}
	return oneBack
}

```

**Cases to test:** Negative totals, the empty sequence, one and two
steps, and several points where a new answer is built from earlier answers.

### Hash-map complement lookup

```go

// Package hashmap demonstrates lookup-based uses of Go maps.
package hashmap

// PairSumIndices returns indices of two distinct values whose sum is target.
// The boolean is false when no pair exists. Subtracting each value from target
// must fit in an int.
func PairSumIndices(values []int, target int) (int, int, bool) {
	// Map every previously visited value to one of its earlier indices.
	indexByValue := make(map[int]int)
	// Visit each value once from left to right.
	for index, value := range values {
		// Calculate the value needed to complete the target.
		complement := target - value
		// Look up the complement and whether it exists as two explicit results.
		earlierIndex, found := indexByValue[complement]
		// Continue into this branch only when an earlier complement was found.
		if found {
			// Return both distinct indices and report that a pair was found.
			return earlierIndex, index, true
		}
		// Insert after lookup so the current element cannot match itself.
		indexByValue[value] = index
	}

	// Return the zero indices with false when the input contains no valid pair.
	return 0, 0, false
}

```

**Cases to test:** A previously seen matching value, equal values at
distinct indices, and inputs with no pair.

### Generic singly linked list

```go

package linkedlist

// Node stores one value and a link to the next node.
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// FromSlice returns a list containing values in their slice order.
func FromSlice[T any](values []T) *Node[T] {
	dummy := &Node[T]{}
	tail := dummy
	for _, value := range values {
		tail.Next = &Node[T]{Value: value}
		tail = tail.Next
	}
	return dummy.Next
}

// ToSlice returns the values encountered from head to nil.
func ToSlice[T any](head *Node[T]) []T {
	values := make([]T, 0)
	for node := head; node != nil; node = node.Next {
		values = append(values, node.Value)
	}
	return values
}

// Reverse reverses a list in place and returns its new head.
func Reverse[T any](head *Node[T]) *Node[T] {
	var previous *Node[T]
	for current := head; current != nil; {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}
	return previous
}

```

**Cases to test:** Slice round trips, reversal, and empty-list
construction, conversion, and reversal.

### Generic stack

```go

package stack

// Stack stores values in last-in, first-out order.
// The zero value is ready to use.
type Stack[T any] struct {
	items []T
}

// Push places value on top of the stack.
func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

// Pop removes and returns the top value. The boolean is false when the stack
// is empty.
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	last := len(s.items) - 1
	value := s.items[last]
	var zero T
	s.items[last] = zero
	s.items = s.items[:last]
	return value, true
}

// Peek returns the top value without removing it. The boolean is false when
// the stack is empty.
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	return s.items[len(s.items)-1], true
}

// Len returns the number of values in the stack.
func (s *Stack[T]) Len() int {
	return len(s.items)
}

// IsEmpty reports whether the stack contains no values.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

```

**Cases to test:** Last-in-first-out ordering, top inspection, length
tracking, complete removal, and empty operations.

### Overflow-safe integer reversal

```go

package integerdigits

import "math"

// ReverseNonNegative reverses a nonnegative decimal integer. ok is false when
// value is negative or when its reversal would exceed int64.
func ReverseNonNegative(value int64) (reversed int64, ok bool) {
	if value < 0 {
		return 0, false
	}

	for value > 0 {
		digit := value % 10
		if reversed > (math.MaxInt64-digit)/10 {
			return 0, false
		}
		reversed = reversed*10 + digit
		value /= 10
	}
	return reversed, true
}

```

**Cases to test:** Zero, palindromes, trailing zeroes, negative
rejection, and a reversal that would overflow `int64`.

### Common-prefix scan

```go

package commonprefixscan

// LongestCommonPrefix returns the longest byte prefix shared by all strings.
// It returns an empty string when strs is empty.
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for column := 0; column < len(strs[0]); column++ {
		expected := strs[0][column]
		for row := 1; row < len(strs); row++ {
			if column >= len(strs[row]) || strs[row][column] != expected {
				return strs[0][:column]
			}
		}
	}
	return strs[0]
}

```

**Cases to test:** Shared and missing prefixes, an early shortest
string, identical strings, a single string, and empty inputs.

### Directional scan

```go

package directionalscan

// SumWithSmallerLeftValuesSubtracted applies a local signed-sum rule: a value
// is subtracted when it is smaller than its immediate right neighbor and added
// otherwise. The running total must fit in an int.
func SumWithSmallerLeftValuesSubtracted(values []int) int {
	if len(values) == 0 {
		return 0
	}

	last := len(values) - 1
	total := values[last]
	rightValue := values[last]

	for index := last - 1; index >= 0; index-- {
		value := values[index]
		if value < rightValue {
			total -= value
		} else {
			total += value
		}
		rightValue = value
	}

	return total
}

```

**Cases to test:** Empty and single inputs, descending values that are
all added, one subtractive pair, and a mix of additions and subtractions.

### Sorted two-pointer pair search

```go

package twopointers

// PairSumSorted returns indices of two values whose sum equals target. Values
// must be sorted in ascending order, and relevant sums must fit in an int.
func PairSumSorted(values []int, target int) (left int, right int, ok bool) {
	left, right = 0, len(values)-1
	for left < right {
		sum := values[left] + values[right]
		switch {
		case sum < target:
			left++
		case sum > target:
			right--
		default:
			return left, right, true
		}
	}

	return 0, 0, false
}

```

**Cases to test:** A matching pair, negative values, the rule that
one position cannot be used twice, a missing pair, and empty input.

## Go techniques

These are Go-specific details that commonly affect otherwise correct
algorithms. Each rule below states the behavior that motivates it.

#### Keep input conversion outside the algorithm

An **online judge**, meaning the automated system that runs and checks a
solution, often supplies text or JSON (JavaScript Object Notation) data. The
algorithm itself should work with problem types such as `[]int`, `string`, or
`*Node`. Convert the external input to those types before the algorithm starts,
and convert the result back only after it finishes. This separation lets the
algorithm express the actual problem instead of mixing parsing with its main
loop.

#### Table-driven tests

A table-driven test stores named rows containing an input and expected result,
then runs the same checking loop over every row. Put ordinary examples and edge
cases in that table so they cannot accidentally use different test logic. If
parsing fails, report that boundary error rather than passing a partly
initialized value into the algorithm.

#### Sorting without unexpected mutation

Go sorting functions rearrange the supplied slice. This is called **mutating**
the input. If the caller expects the original order to remain unchanged, copy
the slice before sorting it. With `sort.Slice`, write a comparison that states
the intended order and any tie-breaking rule explicitly.

#### Map presence and sets

Reading a missing map key returns the value type's **zero value**: `0` for an
integer, `false` for a boolean, and `nil` for a pointer or slice. When zero could
also be a real stored value, read both results:

```go
value, found := entries[key]
```

`found` distinguishes a stored zero from a missing key. A set, which records
only whether keys exist, is conventionally written `map[T]struct{}`; the empty
struct uses no per-value data because only the key matters. Iterating over a Go
map does not promise any order, so sort the keys separately when output order
matters.

#### Priority queues

Use `container/heap` when the program repeatedly needs the next smallest,
largest, or most urgent item. The heap interface asks the program to provide a
comparison method named `Less`. By default, the item that `Less` considers
smallest rises to the top; reverse the comparison when the largest item should
come out first.

Some algorithms discover a better priority for an item already in the heap.
Dijkstra's algorithm can avoid searching for and editing the old entry: push a
new entry with the improved distance. When an older entry later reaches the
top, compare it with the currently stored best distance and discard it if they
differ. Such an outdated entry is called **stale**.

#### Bytes, runes, and visible characters

A Go string is stored as bytes. Basic Latin characters such as `A` use one
byte, but many Unicode characters use several bytes under UTF-8, Go's standard
text encoding. For example, the Go string literal `"\u00e9"` represents one
**rune**, Go's name for one Unicode code point, but has length `2` bytes. Use
byte indexing when the problem explicitly uses bytes or a restricted alphabet.
Range over the string or convert it to `[]rune` when Unicode code points are
the intended units. One user-visible symbol can contain multiple code points;
text-editor rules may therefore need a Unicode library that groups them as
visible characters rather than plain `[]rune` indexing.

#### Type parameters

A **type parameter** lets one implementation work with several value types.
For example, `Stack[T]` can store `int`, `string`, or another type while keeping
the same push and pop behavior. In `Stack[T any]`, `any` means that callers may
choose any value type for `T`. Use type parameters when the operations truly do
not depend on one concrete type.

#### Replacing recursion with a loop

When replacing recursion with a loop, explicitly store what each unfinished
call would have remembered: usually the current item, which choice comes next,
and any partial result needed after returning.

#### Slice storage and retained memory

Two slices may share one underlying array, and `append` may either reuse that
array or allocate a new one. Therefore a change through one slice may or may
not appear through another depending on whether an allocation occurred. A
small surviving slice can also keep a much larger backing array in memory.
Copy the small portion when it must live much longer than the large input.

Long-lived stacks, queues, and heaps should clear removed positions when those
positions contain pointers. Shortening a slice changes its visible length but
may leave the old pointer in the backing array, which can keep an otherwise
unreachable object from being reclaimed.

## Further-study checklist

The material above covers the reusable concepts emphasized by the repository's
problem set and contemporary general interview-preparation lists. Lower-priority
extensions for specialized interviews include strongly connected components,
all-pairs shortest paths, maximum flow, suffix arrays and suffix automata,
computational geometry beyond normalized slopes, and randomized balanced search
trees. These names are pointers for later study, not concepts assumed by this
guide. Learn them when a target role, contest, or interview loop calls for them;
they are not prerequisites for the core patterns above.
