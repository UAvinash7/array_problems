/*

Problem Statement

There is a string, s, of lowercase English letters that is repeated infinitely many times. Given an integer, n, find and print the number of letter a's in the first n letters of the infinite string.

Example

s = 'abcac'
n = 10

The substring we consider is abcacabcac, the first 10 characters of the infinite string. There are 4 occurrences of a in the substring.

Function Description

Complete the repeatedString function in the editor below.

repeatedString has the following parameter(s):

s: a string to repeat
n: the number of characters to consider

Returns

int: the frequency of a in the substring

Input Format

The first line contains a single string, s.
The second line contains an integer, n.

Constraints

1 <= |s| <= 100
1 <= n <= 10 ^ 12
For 25% of the test cases, n <= 10 ^ 6.

Sample Input

Sample Input 0

aba
10

Sample Output 0

7

Explanation 0

The first n = 10 letters of the infinite string are abaabaabaa. Because there are 7 a's, we return 7.

Sample Input 1

a
1000000000000

Sample Output 1

1000000000000

Explanation 1

Because all of the first n = 1000000000000 letters of the infinite string are a, we return 1000000000000.

*/

// Solution

package main

import (
	"fmt"
	"strings"
)

// repeatedString calculates the number of 'a's in the first n characters of an infinite string s.
func repeatedString(s string, n int64) int64 {
	// Length of the original string
	sLen := int64(len(s))

	// Count the number of 'a's in the original string 's'
	countInS := int64(strings.Count(s, "a"))

	// Calculate how many full repetitions of 's' fit into 'n'
	fullRepeats := n / sLen

	// Calculate the number of characters remaining after the full repetitions
	remainder := n % sLen

	// Total count of 'a's from the full repetitions
	totalAs := fullRepeats * countInS

	// Count the 'a's in the remaining partial string
	for i := int64(0); i < remainder; i++ {
		if s[i] == 'a' {
			totalAs++
		}
	}

	return totalAs
}

func main() {
	fmt.Println("Output:", repeatedString("aba", 10))
}
