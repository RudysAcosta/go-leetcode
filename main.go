package main

import (
	"fmt"
	"leetcode/array"
)

func main() {
	// fmt.Println(array.SearchInsertPosition())

	// test := array.PlusOne([]int{9, 9, 9})

	// test := array.TwoSum([]int{3, 2, 4, 4}, 8)

	test := array.MergeSortedArray([]int{2, 0}, 1, []int{1}, 1)

	fmt.Println(test)
}
