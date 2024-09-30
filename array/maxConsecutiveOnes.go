package array

func MaxConsecutiveOnes(nums []int) int {
	max := 0
	j := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			j++
		} else {
			j = 0
		}

		if j > max {
			max = j
		}
	}

	return max
}
