/*

Problem Statement

A jail has a number of prisoners and a number of treats to pass out to them. Their jailer decides the fairest way to divide the treats is to seat the prisoners around a circular table in sequentially numbered chairs. A chair number will be drawn from a hat. Beginning with the prisoner in that chair, one candy will be handed to each prisoner sequentially around the table until all have been distributed.

The jailer is playing a little joke, though. The last piece of candy looks like all the others, but it tastes awful. Determine the chair number occupied by the prisoner who will receive that candy.

Example

n = 4
m = 6
s = 2

There are 4 prisoners, 6 pieces of candy and distribution starts at chair 2. The prisoners arrange themselves in seats numbered 1 to 4. Prisoners receive candy at positions 2, 3, 4, 1, 2, 3. The prisoner to be warned sits in chair number 3.

Function Description

Complete the saveThePrisoner function in the editor below. It should return an integer representing the chair number of the prisoner to warn.

saveThePrisoner has the following parameter(s):

int n: the number of prisoners
int m: the number of sweets
int s: the chair number to begin passing out sweets from

Returns

int: the chair number of the prisoner to warn

Input Format

The first line contains an integer, t, the number of test cases.
The next t lines each contain 3 space-separated integers:

n : the number of prisoners
m : the number of sweets
s : the chair number to start passing out treats at

Constraints

1 <= t <= 100
1 <= n <= 10 ^ 9
1 <= m <= 10 ^ 9
1 <= s <= n

Sample Input 0

2
5 2 1
5 2 2

Sample Output 0

2
3

Explanation 0

In the first query, there are n = 5 prisoners and m = 2 sweets. Distribution starts at seat number s = 1. Prisoners in seats numbered 1 and 2 get sweets. Warn prisoner 2.

In the second query, distribution starts at seat 2 so prisoners in seats 2 and 3 get sweets. Warn prisoner 3.

Sample Input 1

2
7 19 2
3 7 3

Sample Output 1

6
3

Explanation 1

In the first test case, there are n = 7 prisoners, m = 19 sweets and they are passed out starting at chair s = 2. The candies go all around twice and there are 5 more candies passed to each prisoner from seat 2 to seat 6.

In the second test case, there are n = 3 prisoners, m = 7 candies and they are passed out starting at seat s = 3. They go around twice, and there is one more to pass out to the prisoner at seat 3.

*/

// Solution

 package main

import (
	"fmt"
)

// saveThePrisoner determines which prisoner receives the last sweet.
func saveThePrisoner(n int32, m int32, s int32) int32 {
	// Calculate the ending position using modular arithmetic.
	// We use (s - 1) to convert the starting position to 0-based indexing.
	// The total steps are (s - 1) + m.
	// The modulo n handles the circular distribution.
	// We add 1 back to convert the result back to 1-based indexing.
	var result int32 = (s - 1 + m - 1) % n + 1
	return result
}

func main() {
	// Example usage
	n := int32(4) // 4 prisoners
	m := int32(6) // 6 sweets
	s := int32(2) // start at chair 2

	// Expected output: 3
	// Distribution sequence: 2, 3, 4, 1, 2, 3
	fmt.Printf("Number of prisoners: %d, sweets: %d, start: %d -> Last sweet to prisoner: %d\n", n, m, s, saveThePrisoner(n, m, s))

	n2 := int32(7) // 7 prisoners
	m2 := int32(19) // 19 sweets
	s2 := int32(2) // start at chair 2

	// Expected output: 6
	// Distribution sequence: 2, 3, 4, 5, 6, 7, 1...
	fmt.Printf("Number of prisoners: %d, sweets: %d, start: %d -> Last sweet to prisoner: %d\n", n2, m2, s2, saveThePrisoner(n2, m2, s2))

	n3 := int32(3) // 3 prisoners
	m3 := int32(7) // 7 sweets
	s3 := int32(3) // start at chair 3

	// Expected output: 2
	// Distribution sequence: 3, 1, 2, 3, 1, 2, 3...
	fmt.Printf("Number of prisoners: %d, sweets: %d, start: %d -> Last sweet to prisoner: %d\n", n3, m3, s3, saveThePrisoner(n3, m3, s3))
}
