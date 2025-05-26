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
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	//format of string input s = HH:MM:SSAM or HH:MM:SSPM i.e., Hour:Minute:SecondAM/PM
	
	var findAM_OR_PM, hour, minute, second string
	var inputString []string
	var hh, mm, ss int
	findAM_OR_PM = s[len(s)-2:]		// identifies whether it's AM or PM
	
	// spliting the format of 12 hours string value into hours, minutes and seconds.
	inputString = strings.Split(s[:len(s)-2], ":")
	hour = inputString[0]
	minute = inputString[1]
	second = inputString[2]
	
	// converting the value of hour, minute and second variable from string into integer
	hh, _ = strconv.Atoi(hour)
	mm, _ = strconv.Atoi(minute)
	ss, _ = strconv.Atoi(second)

	if hh < 12 && findAM_OR_PM == "PM" {
		hh = hh + 12
	}
	if hh == 12 && findAM_OR_PM == "AM" {
		hh = 0
	}
	return fmt.Sprintf("%02d:%02d:%02d", hh, mm, ss)
}

func main() {
	var testString1 = "07:05:45PM"
	fmt.Printf("time conversion of %s in 24 hour format is: %s\n", testString1, timeConversion(testString1))
	var testString2 = "12:40:22AM"
	fmt.Printf("time conversion of %s in 24 hour format is: %s\n", testString2, timeConversion(testString2))
	var testString3 = "06:40:03AM"
	fmt.Printf("time conversion of %s in 24 hour format is: %s\n", testString3, timeConversion(testString3))
	var testString4 = "12:05:39AM"
	fmt.Printf("time conversion of %s in 24 hour format is: %s\n", testString4, timeConversion(testString4))
	var testString5 = "12:45:54PM"
	fmt.Printf("time conversion of %s in 24 hour format is: %s\n", testString5, timeConversion(testString5))
	var testString6 = "02:34:50PM"
	fmt.Printf("time conversion of %s in 24 hour format is: %s\n", testString6, timeConversion(testString6))
	var testString7 = "04:59:59AM"
	fmt.Printf("time conversion of %s in 24 hour format is: %s\n", testString7, timeConversion(testString7))
	var testString8 = "04:59:59PM"
	fmt.Printf("time conversion of %s in 24 hour format is: %s\n", testString8, timeConversion(testString8))
	var testString9 = "12:00:00AM"
	fmt.Printf("time conversion of %s in 24 hour format is: %s\n", testString9, timeConversion(testString9))
	var testString10 = "11:59:59PM"
	fmt.Printf("time conversion of %s in 24 hour format is: %s\n", testString10, timeConversion(testString10))

	
}