/*

Problem Statement

Staircase detail

This is a staircase of size n= 4:

   #
  ##
 ###
####

Its base and height are both equal to n. It is drawn using # symbols and spaces. The last line is not preceded by any spaces.

Write a program that prints a staircase of size n.

Function Description

Complete the staircase function with the following parameter(s):

int n: an integer

Print

Print a staircase as described above. No value should be returned.

Note: The last line is not preceded by spaces. All lines are right-aligned.

Input Format

A single integer, n, denoting the size of the staircase.

Constraints

 0 < n <= 100.

Sample Input

6

Sample Output

     #
    ##
   ###
  ####
 #####
######

Explanation

The staircase is right-aligned, composed of # symbols and spaces, and has a height and width of n = 6.

*/

// Solution

package main

import "fmt"

func staircase (n int32) {
   for i := int32(0); i < n; i++ {
      for j := i; j < n - 1; j++ {
         fmt.Print(" ")
      }
      for j := int32(0); j <= i; j++ {
         fmt.Print("#")
      }
      fmt.Println()
   }
}

func main() {
	var inputNumber int32
   fmt.Print("Enter the value of inputNumber: ")
   fmt.Scanf("%d\n", &inputNumber)
   if inputNumber <= 0 {
      fmt.Print("input number should not be less than or equal to zero")
      return
   }
   staircase(inputNumber)
}
