/*

Problem Statement

There is a large pile of socks that must be paired by color. Given an array of integers representing the color of each sock, determine how many pairs of socks with matching colors there are.

Example

n = 7
ar = [1, 2, 1, 2, 1, 3, 2]

There is one pair of color 1 and one of color 2. There are three odd socks left, one of each color. The number of pairs is 2.

Function Description

Complete the sockMerchant function in the editor below.

sockMerchant has the following parameter(s):

int n: the number of socks in the pile
int ar[n]: the colors of each sock

Returns

int: the number of pairs

Input Format

The first line contains an integer n, the number of socks represented in ar.
The second line contains n space-separated integers, ar[i], the colors of the socks in the pile.

Constraints

1 <= n <= 100
1 <= ar[i] <= 100, where 0 <= i <= n

Sample Input

STDIN                       	Function
-----                       	--------
9                           	n = 9
10 20 20 10 10 30 50 10 20  	ar = [10, 20, 20, 10, 10, 30, 50, 10, 20]

Sample Output

3

Explanation

There are three pairs of socks.


*/

// Solution

package main

import "fmt"

func sockMerchant(n int32, ar []int32) int32 {
    // Create a map to store the frequency of each sock color
    sockCounts := make(map[int32]int32)

    // Iterate through the array of sock colors and count their frequencies
    for _, color := range ar {
        sockCounts[color]++
    }

    var pairs int32 = 0
    // Iterate through the sock counts and calculate the number of pairs
    for _, count := range sockCounts {
        pairs += count / 2 // Each pair requires two socks of the same color
    }

    return pairs
}

func main() {
    // Example usage:
    n := int32(9)
    ar := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
    result := sockMerchant(n, ar)
    fmt.Println(result) // Expected output: 3
}