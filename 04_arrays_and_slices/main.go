package arrays_and_slices

func Sum(arr []int) int {
	var sum int = 0

	for _, e := range arr {
		sum += e
	}

	return sum
}

func SumAll(arrs ...[]int) []int {
	var sums []int

	for _, numbers := range arrs {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(arrs ...[]int) []int {
	var sums []int
	for _, numbers := range arrs {
		var chunk []int
		if 0 == len(numbers) {
			chunk = []int{0}
		} else {
			chunk = numbers[1:]
		}
		sums = append(sums, Sum(chunk))
	}

	return sums
}
