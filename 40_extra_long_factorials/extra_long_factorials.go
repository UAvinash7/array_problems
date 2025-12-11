/*

Problem Statement

The factorial of the integer n, written n!, is defined as:

n! = n * (n - 1) * (n - 2) * ...... * 3 * 2 * 1

Calculate and print the factorial of a given integer.

For example, if n = 30, we calculate 30 * 29 * 28 * ...... * 3 * 2 * 1 and get 265252859812191058636308480000000.

Function Description

Complete the extraLongFactorials function in the editor below. It should print the result and return.

extraLongFactorials has the following parameter(s):

n: an integer

Note: Factorials of n > 20 can't be stored even in a 64 - bit long long variable. Big integers must be used for such calculations. Languages like Java, Python, Ruby etc. can handle big integers, but we need to write additional code in C/C++ to handle huge values.

We recommend solving this challenge using BigIntegers.

Input Format

Input consists of a single integer n

Constraints

1 <= n <= 100

Output Format

Print the factorial of n.

Sample Input

25

Sample Output

2432902008176640000


Explanation

25! = 25 * 24 * 23 * ...... * 3 * 2 * 1

*/

// Solution

package main

import (
	"fmt"
	"math/big"
)

// extraLongFactorials calculates the factorial of n and prints the result.
func extraLongFactorials(n int32) {
	// Initialize a new BigInt with the value 1.
	result := big.NewInt(1)

	// Iterate from 1 to n, multiplying the result by the current number.
	// We use the Mul method for big integers.
	for i := int32(1); i <= n; i++ {
		// Create a new BigInt for the current multiplier.
		multiplier := big.NewInt(int64(i))
		// Multiply the current result by the multiplier.
		result.Mul(result, multiplier)
	}

	// Print the final result. The big.Int type handles large numbers automatically.
	fmt.Println(result)
}

func main() {
	var n int32
	// Read input from stdin.
	fmt.Scanf("%d\n", &n)

	// Call the function to compute and print the factorial.
	extraLongFactorials(n)
}
