/*

Problem Statement

You are choreographing a circus show with various animals. For one act, you are given two kangaroos on a number line ready to jump in the positive direction (i.e, toward positive infinity).

The first kangaroo starts at location x1 and moves at a rate of v1 meters per jump.
The second kangaroo starts at location x2 and moves at a rate of v2 meters per jump.
You have to figure out a way to get both kangaroos at the same location at the same time as part of the show. If it is possible, return YES, otherwise return NO.

Example
x1 = 2
v1 = 1
x2 = 1
v2 = 2

After one jump, they are both at x = 3, (x1 + v1 = 2 + 1 = 3, x2 + v2 = 1 + 2 = 3), so the answer is YES.

Function Description

Complete the function kangaroo in the editor below.

kangaroo has the following parameter(s):

int x1, int v1: starting position and jump distance for kangaroo 1
int x2, int v2: starting position and jump distance for kangaroo 2

Returns
string: either YES or NO

Input Format

A single line of four space-separated integers denoting the respective values of x1, v1, x2, and v2.

Constraints
0 <= x1 < x2 <= 10000
1 <= v1 <= 10000
1 <= v2 <= 10000

Sample Input 0
0 3 4 2

Sample Output 0
YES

Explanation 0

The two kangaroos jump through the following sequence of locations:

image

From the image, it is clear that the kangaroos meet at the same location (number 12 on the number line) after same number of jumps (4 jumps), and we print YES.

Sample Input 1
0 2 5 3
Sample Output 1
NO
Explanation 1

The second kangaroo has a starting location that is ahead (further to the right) of the first kangaroo's starting location (i.e., x2 > x1). Because the second kangaroo moves at a faster rate (meaning v2 > v1) and is already ahead of the first kangaroo, the first kangaroo will never be able to catch up. Thus, we print NO.

*/

// Solution


package main

import "fmt"

func kangaroo(x1 int32,  v1 int32, x2 int32, v2 int32) string {
	for x1 > x2 {
		return findResult(x1, v1, x2, v2)
	}
	return findResult(x2, v2, x1, v1)
}

func findResult (x1, v1, x2, v2 int32) string {
	for x1 > x2 {
		x1 = x1 + v1
		x2 = x2 + v2
		if x1 == x2 {
			return "YES"
		}
		fmt.Println("value of x1:", x1)
		fmt.Println("value of x2:", x2)
	}
	return "NO"
}

func main() {
	var x1, v1, x2, v2 int32
	fmt.Print("Enter the value of x1: ")
	fmt.Scanf("%d\n", &x1)
	fmt.Print("Enter the value of v1: ")
	fmt.Scanf("%d\n", &v1)
	fmt.Print("Enter the value of x2: ")
	fmt.Scanf("%d\n", &x2)
	fmt.Print("Enter the value of v2: ")
	fmt.Scanf("%d\n", &v2)
	result := kangaroo(x1, v1, x2, v2)
	fmt.Print("result: ", result)
}



// Another Solution

/*

package main

import "fmt"

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if v1 <= v2 {
		return "NO"
	} else if (x2 - x1) % (v1 - v2) == 0 {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	var x1, v1, x2, v2 int32
	fmt.Print("Enter the value of x1: ")
	fmt.Scanf("%d\n", &x1)
	fmt.Print("Enter the value of v1: ")
	fmt.Scanf("%d\n", &v1)
	fmt.Print("Enter the value of x2: ")
	fmt.Scanf("%d\n", &x2)
	fmt.Print("Enter the value of v2: ")
	fmt.Scanf("%d\n", &v2)
	result := kangaroo(x1, v1, x2, v2)
	fmt.Println("result:", result)
}

*/