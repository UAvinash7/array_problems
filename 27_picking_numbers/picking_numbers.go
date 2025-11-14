/*

Problem Statement

Given an array of integers, find the longest subarray where the absolute difference between any two elements is less than or equal to 1.

Example

a = [1, 1, 2, 2, 4, 4, 5, 5, 5]

There are two subarrays meeting the criterion: [1, 1, 2, 2] and [4, 4, 5, 5, 5]. The maximum length subarray has 5 elements.

Function Description

Complete the pickingNumbers function in the editor below.

pickingNumbers has the following parameter(s):

int a[n]: an array of integers

Returns

int: the length of the longest subarray that meets the criterion

Input Format

The first line contains a single integer n, the size of the array a.
The second line contains n space-separated integers, each an a[i].

Constraints

2 <= n <= 100
0 < a[i] < 100

The answer will be >= 2.

Sample Input 0

6
4 6 5 3 3 1

Sample Output 0

3

Explanation 0

We choose the following multiset of integers from the array: {4, 3, 3}. Each pair in the multiset has an absolute difference <= 1 (i.e., |4 - 3| = 1 and |3 - 3| = 0), so we print the number of chosen integers, 3, as our answer.

Sample Input 1

6
1 2 2 3 1 2

Sample Output 1

5

Explanation 1

We choose the following multiset of integers from the array: {1, 2, 2, 1, 2}. Each pair in the multiset has an absolute difference <= 1 (i.e., |1 - 2| = 1, |1 -1| = 0, and |2 - 2| = 0), so we print the number of chosen integers, 5, as our answer.

*/

// Solution

package main

import (
	"fmt"
	"sort"
)

func pickingNumbers(a []int) int {
	// Create a frequency map for the numbers
	freqMap := make(map[int]int)
	for _, num := range a {
		freqMap[num]++
	}

	maxLength := 0

	// Iterate through the numbers in the frequency map
	// We sort the keys to ensure we consider consecutive numbers correctly
	keys := make([]int, 0, len(freqMap))
	for k := range freqMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, num := range keys {
		// Calculate the length of the subarray containing 'num' and 'num+1'
		currentLength := freqMap[num] + freqMap[num+1] // freqMap[num+1] will be 0 if num+1 doesn't exist

		// Update maxLength if currentLength is greater
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

func main() {
	// Example usage
	arr1 := []int{4, 6, 5, 3, 3, 1}
	fmt.Println(pickingNumbers(arr1)) // Output: 3 (from [3, 3, 4] or [4, 5])

	arr2 := []int{1, 2, 2, 3, 1, 2}
	fmt.Println(pickingNumbers(arr2)) // Output: 5 (from [1, 1, 2, 2, 2])
}