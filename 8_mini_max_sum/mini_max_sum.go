/*

Problem Statement

Given five positive integers, find the minimum and maximum values that can be calculated by summing exactly four of the five integers. Then print the respective minimum and maximum values as a single line of two space-separated long integers.

Example

arr = [1, 3, 5, 7, 9]

The minimum sum is 1 + 3 + 5 + 7 = 16 and the maximum sum is 3 + 5 + 7 + 9 = 24. The function prints
16 24

Function Description

Complete the miniMaxSum function with the following parameter(s):

arr[5]: an array of 5 integers

Print

Print two space-separated integers on one line: the minimum sum and the maximum sum of 4 of 5 elements.No value should be returned.

Note For some languages, like C, C++, and Java, the sums may require that you use a long integer due to their size.

Input Format

A single line of five space-separated integers.

Constraints

1 <= arr[i] <= 10 ^ 9

Sample Input

1 2 3 4 5

Sample Output

10 14

Explanation

The numbers are 1, 2, 3, 4, and 5. Calculate the following sums using four of the five integers:

Sum everything except 1, the sum is 2 + 3 + 4 + 5 = 14.
Sum everything except 2, the sum is 1 + 3 + 4 + 5 = 13.
Sum everything except 3, the sum is 1 + 2 + 4 + 5 = 12.
Sum everything except 4, the sum is 1 + 2 + 3 + 5 = 11.
Sum everything except 5, the sum is 1 + 2 + 3 + 4 = 10.

Hints: Beware of integer overflow! Use a 64-bit integer to store the sums.

Need help to get started? Try the Solve Me First problem.

*/

// Solution

package main

import "fmt"

func main() {
	var size int
	fmt.Print("Enter the size of the array: ")
	fmt.Scanf("%d\n", &size)
	if size <= 0 {
		fmt.Print("Value of size should not be less than or equal to 0")
		return
	}
	var inputArray = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter the value for %d th element: \n ", i)
		fmt.Scanf("%d\n", &inputArray[i])
	}
	fmt.Println("Value of inputArray is:", inputArray)

	minimumValue, maximumValue := minMaxSum(inputArray)
	fmt.Printf("minimum sum value of array is: %d and maximum sum value of array is: %d\n", minimumValue, maximumValue)
}

func minMaxSum(arr []int) (int, int) {
var min, max int
for i := 0; i < len(arr) - 1; i++ {
	min += arr[i]
}
for i := 1; i < len(arr); i++ {
	max += arr[i]
}
return min, max
}