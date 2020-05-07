# Chocolate Feast
Reference: [https://www.hackerrank.com/challenges/chocolate-feast/problem](https://www.hackerrank.com/challenges/chocolate-feast/problem)

Little Bobby loves chocolate. He frequently goes to his favorite **5 & 10** store, Penny Auntie, to buy them. They are having a promotion at Penny Auntie. If Bobby saves enough wrappers, he can turn them in for a free chocolate.

For example, Bobby has **n = 15** to spend on bars of chocolate that cost **c = 3** each. He can turn in **m = 2** wrappers to receive another bar. Initially, he buys **5** bars and has **5** wrappers after eating them. He turns in **4** of them, leaving him with **1**, for **2** more bars. After eating those two, he has **3** wrappers, turns in **2** leaving him with **1** wrapper and his new bar. Once he eats that one, he has **2** wrappers and turns them in for another bar. After eating that one, he only has **1** wrapper, and his feast ends. Overall, he has eaten **5 + 2 + 1 + 1 = 9** bars.

## Input Format

The first line contains an integer, **t**, denoting the number of test cases to analyze.

Each of the next lines contains three space-separated integers: **n**, **c**, and **m**. They represent money to spend, cost of a chocolate, and the number of wrappers he can turn in for a free chocolate.

## Constraints

```
1 ≤ t ≤ 1000
2 ≤ n ≤ 10^5
1 ≤ c ≤ n
2 ≤ m ≤ n
```

## Output Format

For each trip to Penny Auntie, print the total number of chocolates Bobby eats on a new line.

## Sample Input

```
3
10 2 5
12 4 4
6 2 2
```

## Sample Output

```
6
3
5
```

## Explanation

Bobby makes the following **3** trips to the store:

1. He spends his **10** dollars on **5** chocolates at **2** dollars apiece. He then eats them and exchanges all **5** wrappers to get **1** more. He eats **6** chocolates.
2. He spends his **12** dollars on **3** chocolates at **4** dollars apiece. He has **3** wrappers, but needs **4** to trade for his next chocolate. He eats **3** chocolates.
3. He spends **6** dollars on **3** chocolates at **2** dollars apiece. He then exchanges **2** of the **3** wrappers for **1** additional piece. Next, he uses his third leftover chocolate wrapper from his initial purchase with the wrapper from his trade-in to do a second trade-in for **1** more piece. At this point he has **1** wrapper left, which is not enough to perform another trade-in. He eats **5** chocolates.
