/*

Problem Statement

Given a square matrix, calculate the absolute difference between the sums of its diagonals.

For example, the square matrix  is shown below:

1 2 3
4 5 6
9 8 9

The left-to-right diagonal = 1 + 5 + 9 = 15.
The right-to-left diagonal = 3 + 5 + 9 = 17.
Their absolute difference is | 15 - 17 | = 2.

Function description

Complete the diagonalDifference function with the following parameter:

int arr[n][m]: a 2-D array of integers

Return

int: the absolute difference in sums along the diagonals

Input Format

The first line contains a single integer, n, the number of rows and columns in the square matrix arr.
Each of the next n lines describes a row, arr[i], and consists of n space-separated integers arr[i][j].

Constraints

-100 <= arr[i][j] <= 100

Sample Input

STDIN      Function
-----      --------
3           arr[][] sizes n = 3, m = 3
11 2 4     arr = [[11, 2, 4], [4, 5, 6], [10, 8, -12]]
4 5 6
10 8 -12

Sample Output

15

Explanation

The primary diagonal is:

11
   5
     -12
Sum across the primary diagonal: 11 + 5 + (-12) = 4.

The secondary diagonal is:

     4
   5
10
Sum across the secondary diagonal: 4 + 5 + 10 = 19

Difference: |4 - 19| = 15

Note: |x| is the absolute value of x.

*/

// Solution

package main

import (
	"fmt"
	"math"
)

func diagonalDifference(arr [][]int32) int32 {
	var differenceResult, resultLeftToRight, resultRightToLeft int32
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i  == j {
				resultLeftToRight += arr[i][j]
			}
			if i + j == len(arr)-1 {
				resultRightToLeft += arr[i][j]
			}
		}
	}
	differenceResult = int32(math.Abs(float64(resultLeftToRight) - float64(resultRightToLeft)))
	return differenceResult

}

func main() {
	var rows, cols int32
	fmt.Print("Enter the number of rows: ")
	fmt.Scanf("%d\n", &rows)
	if rows <= 0 {
		fmt.Print("Number of row should be greater than zero\n")
		return
	}
	fmt.Printf("Enter the number of columns: ")
	fmt.Scanf("%d\n", &cols)
	if cols <= 0 {
		fmt.Print("Number of columns should be greater than zero\n")
		return
	}
	var matrix = make([][]int32, rows)
	for i := int32(0); i < rows; i++ {
		matrix[i] = make([]int32, cols)
		for j := int32(0); j < cols; j++ {
			fmt.Printf("Enter the value of matrix[%d][%d] element: ", i, j)
			fmt.Scanf("%d\n", &matrix[i][j])
		}
	}
	fmt.Println("matrix:", matrix)
	result := diagonalDifference(matrix)
	fmt.Println("diagonal difference: ", result)
}