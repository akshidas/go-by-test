package arraysslices

func Sum(a []int) int {
	var sum int

	for _, x := range a {
		sum += x
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {

	return nil
}
