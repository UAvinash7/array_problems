/*

Problem Statement

Given an array of integers, find the sum of its elements.

For example, if the array ar = [1, 2, 3], 1 + 2 + 3 = 6, so return 6.

Function Description

Complete the simpleArraySum function with the following parameter(s):

ar[n]: an array of integers

Returns
int: the sum of the array elements

Input Format

The first line contains an integer, n, denoting the size of the array.
The second line contains n space-separated integers representing the array's elements.

Constraints
0 < n, ar[i] <= 1000

Sample Input

STDIN           Function
-----           --------
6               ar[] size n = 6
1 2 3 4 10 11   ar = [1, 2, 3, 4, 10, 11]

Sample Output
31

Explanation

Print the sum of the array's elements: 1 + 2 + 3 + 4 + 10 + 11 = 31.

*/

package main

import "fmt"

func simpleSumArray(arr []int) int {
	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func main() {
	var n int
	fmt.Println("Enter the length of an array:")
	fmt.Scanf("%d\n", &n)
	var inputArray []int
	for i := 0; i < n; i++ {
		var temp int
		fmt.Printf("Enter the value of %d th element:\n", i)
		fmt.Scanf("%d\n", &temp)
		inputArray = append(inputArray, temp)
	}
	fmt.Println("Elements of input array is:", inputArray)
	fmt.Println("simple array sum value is:", simpleSumArray(inputArray))
}


/*
Run the code by using command `go run simple_array_sum.go`
After that enter the value of length of the array variable n and hit enter and then followed by the value of array elements
And you will get the complete array elements printed and the sum of the elements in the array as result i.e., sum variable.
*/
