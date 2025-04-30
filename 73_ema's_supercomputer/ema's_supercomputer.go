/*

Problem Statement

Ema built a quantum computer! Help her test its capabilities by solving the problem below.

Given a grid of size n * m, each cell in the grid is either good or bad.

A valid plus is defined here as the crossing of two segments (horizontal and vertical) of equal lengths. These lengths must be odd, and the middle cell of its horizontal segment must cross the middle cell of its vertical segment.

In the diagram below, the blue pluses are valid and the orange ones are not valid. 

pluseses.png

Find the two largest valid pluses that can be drawn on good cells in the grid, and return an integer denoting the maximum product of their areas. In the above diagrams, our largest pluses have areas of 5 and 9. The product of their areas is 5 * 9 = 45.

Note: The two pluses cannot overlap, and the product of their areas should be maximal.

Function Description

Complete the twoPluses function in the editor below. It should return an integer that represents the area of the two largest pluses.

twoPluses has the following parameter(s):

grid: an array of strings where each string represents a row and each character of the string represents a column of that row

Input Format

The first line contains two space-separated integers, n and m.

Each of the next n lines contains a string of m characters where each character is either G (good) or B (bad). These strings represent the rows of the grid. If the y th character in the x th line is G, then (x, y) is a good cell. Otherwise it's a bad cell.

Constraints

2 <= n <= 15
2 <= m <= 15

Output Format

Find 2 pluses that can be drawn on good cells of the grid, and return an integer denoting the maximum product of their areas.

Sample Input 0

5 6
GGGGGG
GBBBGB
GGGGGG
GGBBGB
GGGGGG

Sample Output 0

5

Sample Input 1

6 6
BGBBGB
GGGGGG
BGBBGB
GGGGGG
BGBBGB
BGBBGB

Sample Output 1

25

Explanation

Here are two possible solutions for Sample 0 (left) and Sample 1 (right): 

plusss.png

Explanation Key:

Green: good cell
Red: bad cell
Blue: possible pluses.

For the explanation below, we will refer to a plus of length i as P i.

Sample 0

There is enough good space to color one P 3 plus and one P 1 plus. Area(P 3) = 5 Units, and Area(P 1) = 1 Unit. The product of their areas is 5 * 1 = 5.

Sample 1

There is enough good space to color two P 3 pluses. Area(P 3) = 5 Units. The product of the areas of our two P 3 pluses is 5 * 5 = 25.

*/

// Solution

