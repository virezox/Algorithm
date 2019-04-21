package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MXVIII"))
}

/*
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer. Input is guaranteed to be within the range from 1 to 3999.
*/

// func romanToInt(s string) int {
//     r := 0
//     for i := 0; i < len(s); i++ {
//         switch s[i : i+1] {
//         case "M":
//             r += 1000
//         case "D":
//             r += 500
//         case "C":
//             if i+1 < len(s) {
//                 if s[i+1:i+2] == "D" || s[i+1:i+2] == "M" {
//                     r -= 100
//                 } else {
//                     r += 100
//                 }
//             } else {
//                 r += 100
//             }
//         case "L":
//             r += 50
//         case "X":
//             if i+1 < len(s) {
//                 if s[i+1:i+2] == "L" || s[i+1:i+2] == "C" {
//                     r -= 10
//                 } else {
//                     r += 10
//                 }
//             } else {
//                 r += 10
//             }
//         case "V":
//             r += 5
//         case "I":
//             if i+1 < len(s) {
//                 if s[i+1:i+2] == "V" || s[i+1:i+2] == "X" {
//                     r--
//                 } else {
//                     r++
//                 }
//             } else {
//                 r++
//             }
//         }
//     }
//     return r
// }

/*
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
*/

func romanToInt(s string) int {
	res := 0
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i := 0; i < len(s); i++ {
		if s[i:i+1] == "I" && i+1 < len(s) && (s[i+1:i+2] == "V" || s[i+1:i+2] == "X") ||
			s[i:i+1] == "X" && i+1 < len(s) && (s[i+1:i+2] == "L" || s[i+1:i+2] == "C") ||
			s[i:i+1] == "C" && i+1 < len(s) && (s[i+1:i+2] == "D" || s[i+1:i+2] == "M") {
			res -= m[s[i:i+1]]
		} else {
			res += m[s[i:i+1]]
		}
	}
	return res
}
