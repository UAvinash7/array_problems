/*

Problem Statement

Given a set of distinct integers, print the size of a maximal subset of S where the sum of any 2 numbers in S' is not evenly divisible by k.

Example

S = [19, 10, 12, 10, 24, 25, 22]
k = 4

One of the arrays that can be created is S'[0] = [10, 12, 25]. Another is S'[1] = [19, 22, 24]. After testing all permutations, the maximum length solution array has 3 elements.

Function Description

Complete the nonDivisibleSubset function in the editor below.

nonDivisibleSubset has the following parameter(s):

int S[n]: an array of integers
int k: the divisor

Returns

int: the length of the longest subset of S meeting the criteria

Input Format

The first line contains 2 space-separated integers, n and k, the number of values in S and the non factor.
The second line contains n space-separated integers, each an S[i], the unique values of the set.

Constraints

1 <= n <= 10 ^ 5
1 <= k <= 100
1 <= S[i] <= 10 ^ 9
All of the given numbers are distinct.

Sample Input

STDIN    	Function
-----    	--------
4 3      	S[] size n = 4, k = 3
1 7 2 4  	S = [1, 7, 2, 4]

Sample Output

3

Explanation

The sums of all permutations of two elements from S = {1, 7, 2, 4} are:

1 + 7 = 8
1 + 2 = 3
1 + 4 = 5
7 + 2 = 9
7 + 4 = 11
2 + 4 = 6

Only S' = {1, 7, 4} will not ever sum to a multiple of k = 3.

*/

// Solution

package main

import (
	"fmt"
	"math"
)

// nonDivisibleSubset finds the size of the largest possible subset of S
// where the sum of any two numbers in the subset is not evenly divisible by k.
func nonDivisibleSubset(k int32, s []int32) int32 {
	// Create a frequency map (or array) to store counts of remainders
	remainderCounts := make(map[int32]int32)
	for _, num := range s {
		remainder := num % k
		remainderCounts[remainder]++
	}

	var count int32 = 0

	// Case 1: Handle numbers with a remainder of 0
	// We can only include at most one number with a remainder of 0
	if remainderCounts[0] > 0 {
		count++
	}

	// Case 2: Iterate through remainder pairs
	// We only need to check up to k/2 because pairs (d, k-d) are symmetric
	for d := int32(1); d <= k/2; d++ {
		// If d is equal to k-d (only happens when k is even), it's the middle case
		if d == k-d {
			// Only one element with this remainder can be included
			if remainderCounts[d] > 0 {
				count++
			}
		} else {
			// For other remainders, pick the group with the maximum count
			// as you cannot pick from both d and k-d groups
			count += int32(math.Max(float64(remainderCounts[d]), float64(remainderCounts[k-d])))
		}
	}

	return count
}

func main() {
	// Example usage with sample input from HackerRank:
	// S = [1, 7, 2, 4], k = 3. Max subset is [1, 7, 4] with size 3.
	k := int32(3)
	s := []int32{1, 7, 2, 4}
	result := nonDivisibleSubset(k, s)
	fmt.Printf("Max Non-Divisible Subset Size: %d\n", result) // Output: 3

	// Another example: S = [19, 10, 12, 10, 24, 25, 22], k = 4. Max subset size 3.
	k2 := int32(4)
	s2 := []int32{19, 10, 12, 10, 24, 25, 22}
	result2 := nonDivisibleSubset(k2, s2)
	fmt.Printf("Max Non-Divisible Subset Size: %d\n", result2) // Output: 3
}
