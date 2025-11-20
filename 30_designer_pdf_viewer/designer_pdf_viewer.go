/*

Problem Statement

When a contiguous block of text is selected in a PDF viewer, the selection is highlighted with a blue rectangle. In this PDF viewer, each word is highlighted independently. For example:

 _______________________________________________________________
|						Highlighted Text						|
|																|
|		 _______			 _______			 _______		|
|		|		|			|		|			|		|		|
|		|	abc	|			|	def	|			|	ghi	|		|
|		|_______|			|_______|			|_______|		|
|			Blue				Blue				Blue		|
			Color				Color				Color		|	
|																|
|_______________________________________________________________|																


There is a list of 26 character heights aligned by index to their letters. For example, 'a' is at index 0 and 'z' is at index 25. There will also be a string. Using the letter heights given, determine the area of the rectangle highlight in mm ^ 2 assuming all letters are 1 mm wide.

Example
 
h = [1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 1, 1, 5, 5, 1, 5, 2, 5, 5, 5, 5, 5, 5] word = ` torn `

The heights are  t = 2, o = 1, r = 1 and n = 1. The tallest letter is 2 high and there are 4 letters. The hightlighted area will be 2 * 4 = 8 mm ^ 2 so the answer is 8.

Function Description

Complete the designerPdfViewer function in the editor below.

designerPdfViewer has the following parameter(s):

int h[26]: the heights of each letter
string word: a string

Returns

int: the size of the highlighted area

Input Format

The first line contains 26 space-separated integers describing the respective heights of each consecutive lowercase English letter, ascii[a-z].
The second line contains a single word consisting of lowercase English alphabetic letters.

Constraints

1 <= h[?] <= 7, where ? is an English lowercase letter.
word contains no more than 10 letters.

Sample Input 0

1 3 1 3 1 4 1 3 2 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5
abc

Sample Output 0

9

Explanation 0

We are highlighting the word abc:

Letter heights are a = 1, b= 3 and c = 1. The tallest letter, b, is 3 mm high. The selection area for this word is 3 * 1 mm * 3 mm = 9 mm ^ 2.

Note: Recall that the width of each character is 1 mm.

Sample Input 1

1 3 1 3 1 4 1 3 2 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5 7
zaba

Sample Output 1

28

Explanation 1

The tallest letter in zaba is z at 7 mm. The selection area for this word is 4 * 1 mm * 7 mm = 28 mm ^ 2.

*/

// Solution

package main

import (
	"fmt"
)

// designerPdfViewer calculates the area of the highlighted word.
func designerPdfViewer(h []int32, word string) int32 {
	maxHeight := int32(0)
	for _, char := range word {
		// Calculate the index for the height array (0 for 'a', 1 for 'b', etc.)
		index := char - 'a'
		if h[index] > maxHeight {
			maxHeight = h[index]
		}
	}
	// Area is maxHeight * number of characters in the word
	return maxHeight * int33(len(word))
}

func main() {
	// Example usage:
	// Heights of 'a' through 'z'
	heights := []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7}
	word := "zebra"

	result := designerPdfViewer(heights, word)
	fmt.Println(result) // Expected output: 28 (tallest char 'z' has height 7, word length 5, so 7*4=28)
}