/*
   Author: Patrich Paolo Datu
   Notes:
   This solution only solves for a small set of keys(e.g. the sample inputs), up to 27 keys only.
   I am still unsure on what to do with the remaining message after reading a zero length segment. I just relied on this part "the message is decoded by translating the keys in the segments one-at-a-time into the header characters to which they have been mapped."

   TODOs:
   - Readline for input
   - create a finite 2d array for storing the map instead of using a hardcoded map, not relying on the string value but on the decimal value of the binary key when converted.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	hashmap := map[string]string{
		"0":     "",
		"00":    "",
		"01":    "",
		"10":    "",
		"000":   "",
		"001":   "",
		"010":   "",
		"011":   "",
		"100":   "",
		"101":   "",
		"110":   "",
		"0000":  "",
		"0001":  "",
		"0010":  "",
		"0011":  "",
		"0100":  "",
		"0101":  "",
		"0110":  "",
		"0111":  "",
		"1000":  "",
		"1001":  "",
		"1010":  "",
		"1011":  "",
		"1100":  "",
		"1101":  "",
		"1110":  "",
		"00000": ""}

	keys := [...]string{
		"0",
		"00",
		"01",
		"10",
		"000",
		"001",
		"010",
		"011",
		"100",
		"101",
		"110",
		"0000",
		"0001",
		"0010",
		"0011",
		"0100",
		"0101",
		"0110",
		"0111",
		"1000",
		"1001",
		"1010",
		"1011",
		"1100",
		"1101",
		"1110",
		"00000"}
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
				strings.TrimSuffix(messages[pointer], "000")
				pointer += 1
			}
		}
	}
	for i := 0; i < len(headers); i += 1 {
		// fmt.Println("Evaluating: ", headers[i], " - ", messages[i])
		// Generate Map
		for m := 0; m < len(headers[i]); m += 1 {
			hashmap[keys[m]] = string(headers[i][m])
		}

		// Loop
		first := 0
		for {
			// Get first 3 and evaluate length
			last := first + 3
			length := length(messages[i][first:last])
			if length != 0 {
				// Loop until substring is equal to ones(length)
				first = last
				last = last + length
				for {
					// fmt.Println("Checking map: ", messages[i][first:last], hashmap[messages[i][first:last]])
					if messages[i][first:last] == ones(length) {
						first = last
						break
					} else {
						fmt.Printf(hashmap[messages[i][first:last]])
						first = last
						last = last + length
					}
				}
			} else {
				// Translating keys one at a time
				first = last
				last = last + 1
				if last < len(messages[i]) {
					for {
						// fmt.Println("Checking map: ", messages[i][first:last], hashmap[messages[i][first:last]])
						if messages[i][first:last] == ones(1) {
							first = last
							break
						} else {
							fmt.Printf(hashmap[messages[i][first:last]])
							first = last
							last = last + 1
						}
					}
				} else {
					fmt.Printf("\n")
					break
				}
			}
		}
	}
}

func ones(count int) string {
	ones := ""
	for i := 0; i < count; i += 1 {
		ones += "1"
	}
	return ones
}

func length(str string) int {
	bin := 0
	for i := 0; i < 3; i++ {
		bin = bin*2 + int(str[i]) - int('0')
	}
	// fmt.Println("Length: ", str, bin)
	return bin
}
