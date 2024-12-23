package safety

import (
	"strconv"
	"strings"
)

func IsSafe(s string) bool {
	var nums []int
	for _, numStr := range strings.Split(s, " ") {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return false
		}
		nums = append(nums, num)
	}
	var bigger, smaller bool
	for i, v := range nums {
		if i == 0 {
			continue
		}
		if (v < nums[i-1]) && ((v + 3) >= nums[i-1]) {
			if bigger {
				return false
			}
			smaller = true
			if i < len(nums)-1 {
				continue
			}
		} else if (v > nums[i-1]) && ((v - 3) <= nums[i-1]) {
			if smaller {
				return false
			}
			bigger = true
			if i < len(nums)-1 {
				continue
			}
		} else if ((v + 3) > nums[i-1]) || ((v - 3) < nums[i-1]) {
			return false
		}
		if !bigger && smaller {
			return true
		}
		if !smaller && bigger {
			return true
		}
	}
	return false
}
