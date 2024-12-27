package safety

import (
	"strconv"
	"strings"
)

// IsSafe checks if a space-separated string of integers meets certain safety criteria.
func IsSafe(s string) bool {
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

	// Flags to track if the current number is bigger or smaller than the previous one
	var bigger, smaller bool

	// Iterate through the slice of integers
	for i, v := range nums {
		if i == 0 {
			// Skip the first element as there is no previous element to compare with
			continue
		}

		// Check if the current number is smaller than the previous one and within a range of 3
		if (v < nums[i-1]) && ((v + 3) >= nums[i-1]) {
			if bigger {
				// If the previous number was bigger, return false (invalid pattern)
				return false
			}
			smaller = true
			if i < len(nums)-1 {
				// Continue to the next iteration if not the last element
				continue
			}
		} else if (v > nums[i-1]) && ((v - 3) <= nums[i-1]) {
			// Check if the current number is bigger than the previous one and within a range of 3
			if smaller {
				// If the previous number was smaller, return false (invalid pattern)
				return false
			}
			bigger = true
			if i < len(nums)-1 {
				// Continue to the next iteration if not the last element
				continue
			}
		} else if ((v + 3) > nums[i-1]) || ((v - 3) < nums[i-1]) {
			// If the current number is outside the range of 3 from the previous number, return false
			return false
		}

		// Check if the sequence alternates correctly
		if !bigger && smaller {
			return true
		}
		if !smaller && bigger {
			return true
		}
	}

	// If the sequence does not meet the criteria, return false
	return false
}
