package arraysslices

func Sum(a []int) int {
	var sum int

	for _, x := range a {
		sum += x
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sum []int
	for _, v := range numbersToSum {
		sum = append(sum, Sum(v))
	}
	return sum
}

func SumAllTrail(numbersToSum ...[]int) []int {
	var sum []int

	for _, v := range numbersToSum {
		if len(v) > 0 {
			sum = append(sum, Sum(v[1:]))
		} else {
			sum = append(sum, 0)
		}
	}

	return sum
}
