package array

func DuplicateZeros(arr []int) {

	countZeros := 0
	n := len(arr)

	for i := 0; i < n-countZeros; i++ {
		if arr[i] == 0 {
			if i == n-countZeros-1 {
				arr[n-1] = 0
				n--
				break
			}
			countZeros++
		}
	}

	j := n - 1
	i := j - countZeros

	for i >= 0 {
		if arr[i] == 0 {
			arr[j] = 0
			j--
			arr[j] = 0
		} else {
			arr[j] = arr[i]
		}
		i--
		j--
	}
}
