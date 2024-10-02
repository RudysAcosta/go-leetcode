package array

import (
	"math"
)

func FindNumbersEvenNumber(nums []int) int {

	numEve := 0

	for _, v := range nums {
		cant := int(math.Floor(math.Log10(float64(v)))) + 1

		if cant%2 == 0 {
			numEve++
		}
	}
	return numEve
}
