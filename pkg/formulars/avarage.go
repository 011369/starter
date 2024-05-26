package formulars

func Mean(nums []float64) float64 {
	n := len(nums)
	var sum, mean float64
	for _, num := range nums {
		sum += num
	}
	mean = sum / float64(n)
	return mean
}
