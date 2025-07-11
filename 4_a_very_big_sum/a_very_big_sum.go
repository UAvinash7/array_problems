/*

Problem Statement

In this challenge, you need to calculate and print the sum of elements in an array, considering that some integers may be very large.

Function Description

Complete the aVeryBigSum function with the following parameter(s):

int ar[n]: an array of integers

Return

long: the sum of the array elements

Input Format

The first line of the input consists of an integer n.
The next line contains n space-separated integers contained in the array.

Output Format

Return the integer sum of the elements in the array.

Constraints

1 <= n <= 10
0 <= ar[i] <= 10 ^ (10)

Sample Input

STDIN                                                   Function
-----                                                   --------
5                                                       arr[] size n = 5
1000000001 1000000002 1000000003 1000000004 1000000005  arr[...]

Output

5000000015

Note:

The range of the 32-bit integer is (-2 ^ 31) to ((2 ^ 31) - 1) or [-2147483648, 2147483647].

When we add several integer values, the resulting sum might exceed the above range. You might need to use long int C/C++/Java to store such sums.

*/

// Solution

package main

import "fmt"

func aVeryBigSum(arr []int64) int64 {
	var sum int64
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func main() {
	var n int64
	fmt.Println("Enter the length of the array:")
	fmt.Scanf("%d\n", &n)
	if n <= 0 {
		fmt.Print("Length of array should be greater than zero\n")
		return
	}
	var inputArray []int64
	for i := int64(0); i < n; i++ {
		var temp int64
		fmt.Printf("Enter the value of %d th element:\n", i)
		fmt.Scanf("%d\n", &temp)
		inputArray = append(inputArray, temp)
	}
	fmt.Println("input array:", inputArray)
	result := aVeryBigSum(inputArray)
	fmt.Println("Very Big Sum Of Input Array is:", result)
}

/*

Run the code by using command `go run a_very_big_sum.go`
After that enter the value of variable n and hit enter and then followed by the value of array elements
And you will get the very big sum of the elements in the array as result i.e., sum variable.

*/
