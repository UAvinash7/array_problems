/*

Problem Statement

Complete the function solveMeFirst to compute the sum of two integers.

Example
a = 7
b = 3

Return 10.

Function Description

Complete the solveMeFirst function with the following parameters:

int a: the first value
int b: the second value

Returns
- int: the sum of a and b

Constraints
1 <= a, b <= 1000

Sample Input
a = 2
b = 3

Sample Output
5

Explanation
2 + 3 = 5

*/

// Solution

package main

import "fmt"

func solveMeFirst(a uint32, b uint32) uint32 {
	return a + b
}

func main() {
	var a, b, res uint32
	fmt.Println("Enter first value:")
	fmt.Scanf("%d\n", &a)
	fmt.Println("Enter second value:")
	fmt.Scanf("%d\n", &b)
	res = solveMeFirst(a, b)
	fmt.Println("Sum of first and second value is:", res)
}

/*
Run the code by using command `go run solve_me_first.go`
After that enter the value of variable a and variable b
And you will get the sum of the variables a and b as result i.e., res variable.
*/
