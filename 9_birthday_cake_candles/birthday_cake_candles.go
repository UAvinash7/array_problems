/*

Problem Statement

You are in charge of the cake for a child's birthday. It will have one candle for each year of their total age. They will only be able to blow out the tallest of the candles. Your task is to count how many candles are the tallest.

Example

candles = [4, 4, 1, 3]

The tallest candles are 4 units high. There are 2 candles with this height, so the function should return 2.

Function Description

Complete the function birthdayCakeCandles with the following parameter(s):

int candles[n]: the candle heights

Returns

int: the number of candles that are tallest

Input Format

The first line contains a single integer, n, the size of candles[].
The second line contains n space-separated integers, where each integer i describes the height of candles[i].

Constraints

1 <= n <= 10 ^ 5
1 <= candles[i] <= 10 ^ 7

Sample Input 0
4
3 2 1 3

Sample Output 0
2

Explanation 0
Candle heights are [3, 2, 1, 3]. The tallest candles are 3 units, and there are 2 of them.

*/

// Solution

package main

import "fmt"

func birthdayCakeCandles(candles []int32) int32 {
	var max, count int32
	for i := 0; i < len(candles); i++ {
		if max < candles[i] {
			max = candles[i]
		}
	}
	for i := 0; i < len(candles); i++ {
		if max == candles[i] {
			count++
		}
	}
	return count
}

func main() {
	var size int
	fmt.Print("Enter the size of the array: ")
	fmt.Scanf("%d\n", &size)
	if size <= 0 {
		fmt.Print("Size cannot be less than or equal to zero")
		return
	}
	var inputArray = make([]int32, size)
	for i:= 0; i < size; i++ {
		fmt.Printf("Enter the value of %d th element: \n", i)
		fmt.Scanf("%d\n", &inputArray[i])
	}
	fmt.Println("Input Array:", inputArray)
	numberOfTallestCandles := birthdayCakeCandles(inputArray)
	fmt.Println("Number of tallest candles: ", numberOfTallestCandles)
}
