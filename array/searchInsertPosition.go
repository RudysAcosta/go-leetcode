package array

func SearchInsertPosition() (index int) {
	nums := [4]int{3, 4, 5, 6}
	target := 2

	for i := 0; i < len(nums); i++ {
		if target <= nums[i] {
			return i
		}
	}

	return len(nums)
}
