/*
   Author: Patrich Paolo Datu
   Notes:
   This will now cover the maximum number of keys for input
   I am still unsure on what to do with the remaining message after reading a zero length segment. I just relied on this part "the message is decoded by translating the keys in the segments one-at-a-time into the header characters to which they have been mapped."

   TODOs:
   - Readline for input
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	array := [...]string{"TNM AEIOU",
		"0010101100011",
		"1010001001110110011",
		"11000",
		"$#**\\",
		"0100000101101100011100001000"}

	pointer := 0
	headers := []string{}
	messages := []string{}

	for i := 0; i < len(array); i += 1 {
		// fmt.Println(array[i])
		if len(headers) < (pointer + 1) {
			headers = append(headers, array[i])
			messages = append(messages, "")
		} else {
			messages[pointer] += array[i]
			if strings.HasSuffix(array[i], "000") {
				//strings.TrimSuffix(messages[pointer], "000")
				pointer += 1
			}
		}
	}
	for i := 0; i < len(headers); i += 1 {
		// fmt.Println("Evaluating: ", headers[i], " - ", messages[i])
		// Generate Map
		var twod [8][127]string // 0-7 length and 0-126 possible keys
		l := 0
		for j := 1; j < 8; j += 1 {
			for k := 0; k < (poweroftwo(j)-1) && l < len(headers[i]); k += 1 {
				twod[j][k] = string(headers[i][l])
				l++
			}
		}
		// fmt.Println("map: ", twod)
		// Loop
		first := 0
		for {
			// Get first 3 and evaluate length
			last := first + 3
			length := bintodec(messages[i][first:last])
			if length != 0 {
				// Loop until substring is equal to ones(length)
				first = last
				last = last + length
				for {
					// fmt.Println("Checking map: ", messages[i][first:last], twod[length][bintodec(messages[i][first:last])])
					if messages[i][first:last] == ones(length) {
						first = last
						break
					} else {
						fmt.Printf(twod[length][bintodec(messages[i][first:last])])
						first = last
						last = last + length
					}
				}
			} else {
				// Translating keys one at a time
				length = 1
				first = last
				last = last + 1
				if last < len(messages[i]) {
					for {
						// fmt.Println("Checking map: ", messages[i][first:last], twod[length][bintodec(messages[i][first:last])])
						if messages[i][first:last] == ones(length) {
							first = last
							break
						} else {
							fmt.Printf(twod[length][bintodec(messages[i][first:last])])
							first = last
							last = last + length
						}
					}
				} else {
					fmt.Println()
					break
				}
			}
		}
	}
}

func poweroftwo(count int) int {
	if count == 0 {
		return 1
	}
	return 2 * poweroftwo(count-1)
}

func ones(count int) string {
	ones := ""
	for i := 0; i < count; i += 1 {
		ones += "1"
	}
	return ones
}

func bintodec(str string) int {
	bin := 0
	for i := 0; i < len(str); i++ {
		bin = bin*2 + int(str[i]) - int('0')
	}
	//fmt.Println("Bin to Dec: ", str, bin)
	return bin
}
