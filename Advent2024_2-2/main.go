package main

import (
	"fmt"
	"project-layout/filesHandling"
	"project-layout/safety"
	"strconv"
	"strings"
)

func main() {
	l := filesHandling.ReadFile("advent1.txt")
	total := 0
	for _, v := range l {
		if safety.IsSafe(v) {
			total++
		} else if isSpecialSafe(v) {
			total++
		}
	}
	fmt.Println(total)
}
func isSpecialSafe(s string) bool {
	// Slice to hold the converted integers
	var nums []int

	// Split the input string by spaces and convert each part to an integer
	for _, numStr := range strings.Split(s, " ") {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			// If any conversion fails, return false
			return false
		}
		nums = append(nums, num)
	}
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		// Create a copy of the original list
		newNums := make([]int, len(nums))
		copy(newNums, nums)
		newNums = append(newNums[:i], newNums[i+1:]...)
		fmt.Println(newNums)
		st := intSliceToString(newNums)
		fmt.Println(st)
		if safety.IsSafe(st) {
			return true
		}
	}
	return false
}

func intSliceToString(numbers []int) string {
	s := ""
	for i, n := range numbers {
		if i < len(numbers)-1 {
			s += strconv.Itoa(n) + " "
		} else {
			s += strconv.Itoa(n)
		}
	}
	return s
}
