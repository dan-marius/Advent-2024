package main

import (
	"fmt"
	"project-layout/filesHandling"
	"regexp"
	"strconv"
)

func main() {

	// Read numbers from the file
	lines := filesHandling.ReadFile("advent1.txt")
	var sum int64
	for _, line := range lines {
		sum += captures(line)
	}
	fmt.Println(sum)
}
func captures(content string) int64 {
	// Content to be parsed.

	// Regex pattern captures "mul(number1,number2)"
	pattern := regexp.MustCompile(`mul\((?P<num1>\d{1,3}),(?P<num2>\d{1,3})\)`)

	// Template to convert "key: value" to "key=value" by
	// referencing the values captured by the regex pattern.
	template := "mul($num1, $num2) = %d\n"

	result := []byte{}
	sum := 0
	// For each match of the regex in the content.
	for _, submatches := range pattern.FindAllStringSubmatchIndex(content, -1) {
		// Extract the captured numbers
		num1Str := string(content[submatches[2]:submatches[3]])
		num2Str := string(content[submatches[4]:submatches[5]])

		// Convert strings to integers
		num1, _ := strconv.Atoi(num1Str)
		num2, _ := strconv.Atoi(num2Str)

		// Calculate the product
		product := num1 * num2
		sum += product
		// Apply the template with the calculated product
		result = pattern.ExpandString(result, fmt.Sprintf(template, product), content, submatches)
	}
	fmt.Println(string(result))
	return int64(sum)
}