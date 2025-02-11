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

func main() {
	var n, temp, sum int
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d\n", &temp)
		sum += temp
	}
	fmt.Println(sum)
}


/*
Run the code by using command `go run simple_array_sum.go`
After that enter the value of variable n and hit enter and then followed by the value of array elements
And you will get the sum of the elements in the array as result i.e., sum variable.
*/
