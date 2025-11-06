/*

Problem Statement

A person wants to determine the most expensive computer keyboard and USB drive that can be purchased with a give budget. Given price lists for keyboards and USB drives and a budget, find the cost to buy them. If it is not possible to buy both items, return -1.

Example

b = 60
keyboards = [40, 50, 60]
drives = [5, 8, 12]

The person can buy a 40 keyboard + 12 USB drive = 52, or a 50 keyboard + 8 USB drive = 58. Choose the latter as the more expensive option and return 58.

Function Description

Complete the getMoneySpent function in the editor below.

getMoneySpent has the following parameter(s):

int keyboards[n]: the keyboard prices
int drives[m]: the drive prices
int b: the budget

Returns

int: the maximum that can be spent, or -1 if it is not possible to buy both items

Input Format

The first line contains three space-separated integers b, n, and m, the budget, the number of keyboard models and the number of USB drive models.
The second line contains n space-separated integers keyboard[i], the prices of each keyboard model.
The third line contains m space-separated integers drives, the prices of the USB drives.

Constraints

1 <= n, m <= 1000
1 <= b <= 10 ^ 6
The price of each item is in the inclusive range [1, 10 ^ 6].

Sample Input 0

10 2 3
3 1
5 2 8

Sample Output 0

9

Explanation 0

Buy the 2nd keyboard and the 3rd USB drive for a total cost of 8 + 1 = 9.

Sample Input 1

5 1 1
4
5

Sample Output 1

-1

Explanation 1

There is no way to buy one keyboard and one USB drive because 4 + 5 > 5, so return -1.

*/

// Solution

package main

import (
	"fmt"
	"sort"
)

func getMoneySpent(keyboards []int, drives []int, b int) int {
	maxMoneySpent := -1

	// Sort keyboards in ascending order to optimize the search
	sort.Ints(keyboards)
	// Sort drives in ascending order to optimize the search
	sort.Ints(drives)

	// Iterate through each keyboard price
	for _, keyboardPrice := range keyboards {
		// Iterate through each drive price
		for _, drivePrice := range drives {
			totalPrice := keyboardPrice + drivePrice

			// If the total price exceeds the budget, move to the next keyboard
			// (since both arrays are sorted, further drive prices with this keyboard
			// will also exceed the budget)
			if totalPrice > b {
				break
			}

			// If the total price is within budget and greater than the current max, update maxMoneySpent
			if totalPrice > maxMoneySpent {
				maxMoneySpent = totalPrice
			}
		}
	}

	return maxMoneySpent
}

func main() {
	// Example usage:
	// b: budget
	// keyboards: array of keyboard prices
	// drives: array of drive prices
	b := 60
	keyboards := []int{40, 50, 60}
	drives := []int{5, 8, 12}

	result := getMoneySpent(keyboards, drives, b)
	fmt.Println(result) // Output: 58

	b2 := 10
	keyboards2 := []int{3, 1}
	drives2 := []int{5, 2, 8}

	result2 := getMoneySpent(keyboards2, drives2, b2)
	fmt.Println(result2) // Output: 8

	b3 := 5
	keyboards3 := []int{4}
	drives3 := []int{5}

	result3 := getMoneySpent(keyboards3, drives3, b3)
	fmt.Println(result3) // Output: -1
}