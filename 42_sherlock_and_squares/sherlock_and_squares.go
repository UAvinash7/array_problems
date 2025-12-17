/*

Problem Statement

Watson likes to challenge Sherlock's math ability. He will provide a starting and ending value that describe a range of integers, inclusive of the endpoints. Sherlock must determine the number of square integers within that range.

Note: A square integer is an integer which is the square of an integer, e.g. 1, 4, 9, 16, 25.

Example

a = 24
b = 49

There are three square integers in the range: 25, 36 and 49. Return 3.

Function Description

Complete the squares function in the editor below. It should return an integer representing the number of square integers in the inclusive range from a to b.

squares has the following parameter(s):

int a: the lower range boundary
int b: the upper range boundary

Returns

int: the number of square integers in the range

Input Format

The first line contains q, the number of test cases.
Each of the next q lines contains two space-separated integers, a and b, the starting and ending integers in the ranges.

Constraints

1 <= q <= 100
1 <= a <= b <= 10 ^ 9

Sample Input

2
3 9
17 24

Sample Output

2
0

Explanation

Test Case #00: In range [3, 9], 4 and 9 are the two square integers.
Test Case #01: In range [17, 24], there are no square integers.

*/

// Solution

package main

import (
    "fmt"
    "math"
)

// squares function calculates the number of square integers in the range [a, b].
func squares(a int32, b int32) int32 {
    // Calculate the ceiling of the square root of 'a'
    sqrtA := math.Ceil(math.Sqrt(float64(a)))
    // Calculate the floor of the square root of 'b'
    sqrtB := math.Floor(math.Sqrt(float64(b)))

    // The number of perfect squares is the difference between the integer square roots + 1
    // (to make the range inclusive, e.g., integers in [3, 4] are 3 and 4, so 4-3+1 = 2)
    count := int32(sqrtB) - int32(sqrtA) + 1

    // If the count is negative (which can happen with certain floating point errors or edge cases if not handled correctly), return 0.
    if count < 0 {
        return 0
    }

    return count
}

func main() {
    // Example usage for a single test case (a=24, b=49)
    // The HackerRank problem requires reading multiple test cases from stdin.
    var a, b int32 = 24, 49
    result := squares(a, b)
    fmt.Println(result) // Output: 3 (squares are 25, 36, 49)

    // Another example (a=3, b=9)
    // Squares are 4 and 9
    var a2, b2 int32 = 3, 9
    result2 := squares(a2, b2)
    fmt.Println(result2) // Output: 2
}
