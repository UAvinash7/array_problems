/*

Problem Statement

Given a time in 12-hour AM/PM format, convert it to military (24-hour) time.

Note: - 12:00:00AM on a 12-hour clock is 00:00:00 on a 24-hour clock.
- 12:00:00PM on a 12-hour clock is 12:00:00 on a 24-hour clock.

Example
s = `12:01:00PM`

Return '12:01:00'.

s = `12:01:00AM`

Return '00:01:00'.

Function Description

Complete the timeConversion function with the following parameter(s):

string s: a time in 12 hour format

Returns

string: the time in 24 hour format

Input Format

A single string s that represents a time in 12-hour clock format (i.e.: hh:mm:ssAM or hh:mm:ssPM).

Constraints

All input times are valid

Sample Input 0

07:05:45PM

Sample Output 0

19:05:45


*/


// Solution

package main

import (
	"fmt"
	"strings"
	"strconv"
)

func timeConversion(s string) string {
	//format of string input HH:MM:SSAM or HH:MM:SSPM i.e., Hour:Minute:SecondAM/PM

	var findAMOrPM string = s[len(s)-2:]	// identifies whether it's AM or PM

	// spliting the format of 12 hours numerical value into hours, minutes and seconds.
	timeFormat := strings.Split(s[:len(s)-2], ":")
	hh := timeFormat[0]
	mm := timeFormat[1]
	ss := timeFormat[2]

	// converting the value of hour, minute and second variable from string into integer
	hour, _ := strconv.Atoi(hh)
	minute, _ := strconv.Atoi(mm)
	second, _ := strconv.Atoi(ss)

	if hour < 12 && findAMOrPM == "PM" {
		hour += 12
	}

	if hour == 12 && findAMOrPM == "AM" {
		hour = 0
	}
	
	return fmt.Sprintf("%02d:%02d:%02d%s\n", hour, minute, second, findAMOrPM)
	

}

func main() {
	
	// 12 hour format input
	var str1 string = "07:05:45AM"
	var str2 string = "07:05:45PM"
	var str3 string = "12:05:59AM"
	var str4 string = "13:41:54PM"
 
	// we will get 24 hour military format output
	fmt.Println(timeConversion(str1))
	fmt.Println(timeConversion(str2))
	fmt.Println(timeConversion(str3))
	fmt.Println(timeConversion(str4))


}