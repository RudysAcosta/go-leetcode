package main

import (
	"fmt"
	"leetcode/array"
)

func main() {
	// fmt.Println(array.SearchInsertPosition())

	// test := array.PlusOne([]int{9, 9, 9})

	// test := array.TwoSum([]int{3, 2, 4, 4}, 8)

	// fmt.Println(test)

	// matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// array.RotateImage(matrix)

	nums := []int{1, 2, 4, 2, 4, 4, 3}

	// fmt.Println(array.SquaresSorted(nums))

	fmt.Println(array.RemoveElement(nums, 4))

}
