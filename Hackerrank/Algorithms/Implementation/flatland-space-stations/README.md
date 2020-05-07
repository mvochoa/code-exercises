# Flatland Space Stations
Reference: [https://www.hackerrank.com/challenges/flatland-space-stations/problem](https://www.hackerrank.com/challenges/flatland-space-stations/problem)

Flatland is a country with a number of cities, some of which have space stations. Cities are numbered consecutively and each has a road of $`1km`$ length connecting it to the next city. It is not a circular route, so the first city doesn't connect with the last city. Determine the maximum distance from any city to it's nearest space station.

For example, there are $`n = 3`$ cities and $`m = 1`$ of them has a space station, city $`1`$. They occur consecutively along a route. City $`2`$ is $`2 - 1 = 1`$ unit away and city $`3`$ is $`3 - 1 = 2`$ units away. City $`1`$ is $`0`$ units from its nearest space station as one is located there. The maximum distance is $`2`$.

## Input Format

The first line consists of two space-separated integers, $`n`$ and $`m`$.

The second line contains $`m`$ space-separated integers, the indices of each city having a space-station. These values are unordered and unique.

## Constraints

- $`1 \leq n \leq 10^{5}`$
- $`1 \leq m \leq n`$
- There will be at least $`1`$ city with a space station.
- No city has more than one space station.

## Output Format

Print an integer denoting the maximum distance that an astronaut in a Flatland city would need to travel to reach the nearest space station.

## Sample Input 0

```
5 2
0 4
```

## Sample Output 0

```
2
```

## Explanation 0

This sample corresponds to following graphic:

![](upload/graph.png)

The distance to the nearest space station for each city is listed below:

- $`c[0]`$ has distance - $`0 km`$, as it contains a space station.
- $`c[1]`$ has distance - $`1 km`$, to the space station in $`c[0]`$.
- $`c[2]`$ has distance - $`2 km`$, to the space stations in $`c[0]`$ and $`c[4]`$.
- $`c[3]`$ has distance - $`1 km`$, to the space station in $`c[4]`$.
- $`c[4]`$ has distance - $`0 km`$, as it contains a space station.

We then take $`max(0, 1, 2, 1, 0) = 2`$.

## Sample Input 1

```
6 6
0 1 2 4 3 5
```

## Sample Output 1

```
0
```

## Explanation 1

In this sample, $`n = m`$ so every city has space station and we print $`0`$ as our answer.
