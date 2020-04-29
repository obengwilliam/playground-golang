package arraysslices

func Sum(v []int) int {
	var sum = 0
	for _, num := range v {
		sum += num
	}
	return sum
}

func SumAll(arrays ...[]int) []int {
	var sum = []int{}
	for _, nums := range arrays {
		sum = append(sum, Sum(nums))
	}
	return sum
}

func SumAllTails(arrays ...[]int) []int {
	var sum = []int{}
	for _, nums := range arrays {
		if len(nums) == 0 {
			sum = append(sum, 0)
		} else {

			sum = append(sum, Sum(nums[1:]))
		}
	}
	return sum
}
