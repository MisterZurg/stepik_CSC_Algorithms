# Algorithmic Warm Up
## Programming Challenges
1. Fibonacci Number
2. Last Digit of Fibonacci Number
3. Huge Fibonacci Number
4. Last Digit of the Sum of Fibonacci Numbers
5. Last Digit of the Partial Sum of Fibonacci Numbers
6. Last Digit of the Sum of Squares of Fibonacci Numbers
7. Greatest Common Divisor
8. Least Common Multiple

## Summary of Algorithmic Ideas
**What solution fits into a second?**  
Modern computers perform about
$10^8 - 10^9$  basic operations per second. 
If your program contains a loop with `n` iterations and `n` can be as large as $10^14$,
it will run for a couple of days. In turn, this means that for a programming challenge where a naive solution takes so many steps, you need to come up with a different idea.

**Working with large integers.** 
If you need to compute the last digit of an integer
`m`, then, in many cases, you can avoid computing `m` explicitly: when computing it, take every intermediate step modulo `10`.
This will ensure that all intermediate values are small enough so that they fit into integer types (in programming languages with integer overflow) and that arithmetic operations with them are fast.

## Interview Questions
**Josephus**
> Find the position of the survivor in the vicious series of killings described by Flavius Josephus, the first-century historian and head of Jewish forces in Galilee.

**Range Sum Queries**
> Given an integer array and a set of ranges in it, compute the sum for each range.

