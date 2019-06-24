# A Very Big Sum
Reference: [https://www.hackerrank.com/challenges/a-very-big-sum/problem](https://www.hackerrank.com/challenges/a-very-big-sum/problem)

Calculate and print the sum of the elements in an array, keeping in mind that some of those integers may be quite large.

## Input Format

The first line of the input consists of an integer $`n`$.

The next line contains $`n`$ space-separated integers contained in the array.

## Constraints

- $`1 \leq n \leq 10`$
- $`0 \leq arr[i] \leq 10^{10}`$

## Output Format

Print the integer sum of the elements in the array.

## Sample Input

```
5
1000000001 1000000002 1000000003 1000000004 1000000005
```

Sample Output

```
5000000015
```

## Note

The range of the 32-bit integer is $`(-2^{31})\ to\ (2^{31} - 1)\ or\ [-2147483648,2147483647]`$.

When we add several integer values, the resulting sum might exceed the above range. You might need to use long long int in C/C++ or long data type in Java to store such sums.
