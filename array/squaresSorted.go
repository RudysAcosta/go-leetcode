package array

import (
	"sort"
)

func SquaresSorted(nums []int) []int {

	squaresNums := []int{}

	for _, v := range nums {
		squaresNums = append(squaresNums, v*v)
	}

	sort.Ints(squaresNums)

	return nums
}

func SquaresSortedTwoPoint(nums []int) []int {

	n := len(nums)
	result := make([]int, n)

	left, right := 0, n-1
	pos := right

	for left <= right {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]

		if leftSquare > rightSquare {
			result[pos] = leftSquare
			left++
		} else {
			result[pos] = rightSquare
			right--
		}

		pos--
	}

	return result

}
