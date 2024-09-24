package array

func TwoSum(nums []int, target int) []int {
	comp := make(map[int]int)
	for i, n := range nums {
		if index, found := comp[target-n]; found {
			return []int{index, i}
		}
		comp[n] = i
	}
	return nil
}
