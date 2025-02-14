/*

Problem Statement

Given an array of integers, calculate the ratios of its elements that are positive, negative, and zero. Print the decimal value of each fraction on a new line with 6 places after the decimal.

Note: This challenge introduces precision problems. The test cases are scaled to six decimal places, though answers with absolute error of up to 10^(-4) are acceptable.

Example
arr = [1, 1, 0, -1, -1]

There are n = 5 elements: two positive, two negative and one zero. Their ratios are 2/5 = 0.400000, 2/5 = 0.400000 and 1/5 = 0.200000. Results are printed as:

0.400000
0.400000
0.200000

Function Description

Complete the plusMinus function with the following parameter(s):

int arr[n]: an array of integers

Print
Print the ratios of positive, negative and zero values in the array. Each value should be printed on a separate line with 6 digits after the decimal. The function should not return a value.

Input Format

The first line contains an integer, n, the size of the array.
The second line contains n space-separated integers that describe arr[n].

Constraints
0 < n <= 100
-100 <= arr[i] <= 100


Sample Input

STDIN           Function
-----           --------
6               arr[] size n = 6
-4 3 -9 0 4 1   arr = [-4, 3, -9, 0, 4, 1]

Sample Output

0.500000
0.333333
0.166667

Explanation

There are 3 positive numbers, 2 negative numbers, and 1 zero in the array.
The proportions of occurrence are positive: 3/6 = 0.500000, negative: 2/6 = 0.333333 and zeros: 16 = 0.166667.

*/

// Solution

package main

import "fmt"

func plusMinus(arr []int32) {
    // Write your code here
    var positiveNumber int = 0
    var zeroNumber int = 0
    var negativeNumber int = 0
    var positiveNumberRatio, negativeNumberRatio, zeroNumberRatio float64
    for i := 0; i < len(arr); i++ {
        if arr[i] > 0 {
            positiveNumber++
        } else if arr[i] < 0 {
            negativeNumber++
        } else {
            zeroNumber++
        }
    }
    positiveNumberRatio = float64(positiveNumber/len(arr))
    zeroNumberRatio = float64(zeroNumber/len(arr))
    negativeNumberRatio = float64(negativeNumber/len(arr))
    fmt.Printf("%.6f\n", positiveNumberRatio)
    fmt.Printf("%.6f\n", zeroNumberRatio)
    fmt.Printf("%.6f\n", negativeNumberRatio)

}

func main() {
	var arr = []int32{1,1,0,-1,-1}
	plusMinus(arr)
}